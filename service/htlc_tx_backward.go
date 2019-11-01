package service

import (
	"LightningOnOmni/bean"
	"LightningOnOmni/dao"
	"LightningOnOmni/tool"
	"encoding/json"
	"errors"
	"github.com/asdine/storm/q"
	"log"
	"sync"
)

type htlcBackwardTxManager struct {
	operationFlag sync.Mutex
}

// HTLC Reverse pass the R (Preimage R)
var HtlcBackwardTxService htlcBackwardTxManager

// SendRToPreviousNode
//
// Process type -46: Send R to Previous Node (middleman).
//  * R is <Preimage_R>
func (service *htlcBackwardTxManager) SendRToPreviousNode(msgData string,
	user bean.User) (data map[string]interface{}, previousNode string, err error) {

	// region Parse data inputed from [Carol] websocket client.
	if tool.CheckIsString(&msgData) == false {
		return nil, "", errors.New("empty json data")
	}

	reqData := &bean.HtlcSendR{}
	err = json.Unmarshal([]byte(msgData), reqData)
	if err != nil {
		log.Println(err.Error())
		return nil, "", err
	}
	// endregion

	// region Check data inputed from websocket client of Carol.
	if tool.CheckIsString(&reqData.RequestHash) == false {
		err = errors.New("empty request_hash")
		log.Println(err)
		return nil, "", err
	}

	if tool.CheckIsString(&reqData.ChannelAddressPrivateKey) == false {
		err = errors.New("channel_address_private_key is empty")
		log.Println(err)
		return nil, "", err
	}

	if tool.CheckIsString(&reqData.CurrHtlcTempAddressForHE1bPubKey) == false {
		err = errors.New("curr_htlc_temp_address_for_he1b_pub_key is empty")
		log.Println(err)
		return nil, "", err
	}

	if tool.CheckIsString(&reqData.CurrHtlcTempAddressForHE1bPrivateKey) == false {
		err = errors.New("curr_htlc_temp_address_for_he1b_private_key is empty")
		log.Println(err)
		return nil, "", err
	}
	// endregion

	// region Check out if the input R is correct.
	rAndHInfo := &dao.HtlcRAndHInfo{}
	err = db.Select(
		q.Eq("RequestHash", reqData.RequestHash),
		q.Eq("R", reqData.R), // R from websocket client of Carol
		q.Eq("CurrState", dao.NS_Finish)).First(rAndHInfo)

	if err != nil {
		log.Println(err.Error())
		return nil, "", err
	}
	// endregion

	// region Get peerId of previous node.
	htlcSingleHopPathInfo := dao.HtlcSingleHopPathInfo{}
	err = db.Select(q.Eq("HAndRInfoRequestHash",
		reqData.RequestHash)).First(&htlcSingleHopPathInfo)

	if err != nil {
		log.Println(err)
		return nil, "", err
	}

	// Currently solution is Alice to Bob to Carol.
	if htlcSingleHopPathInfo.CurrStep < 2 {
		return nil, "", errors.New("The transfer H has not completed yet.")
	} else if htlcSingleHopPathInfo.CurrStep > 3 {
		return nil, "", errors.New("The transfer R has completed.")
	}

	// If CurrStep = 2, that indicate the transfer H has completed.
	currChannelIndex := htlcSingleHopPathInfo.TotalStep - htlcSingleHopPathInfo.CurrStep - 1
	if currChannelIndex < -1 || currChannelIndex > len(htlcSingleHopPathInfo.ChannelIdArr) {
		return nil, "", errors.New("err channel id")
	}

	currChannel := &dao.ChannelInfo{}
	err = db.One("Id", htlcSingleHopPathInfo.ChannelIdArr[currChannelIndex], currChannel)
	if err != nil {
		log.Println(err.Error())
		return nil, "", err
	}

	if currChannel.PeerIdA != user.PeerId && currChannel.PeerIdB != user.PeerId {
		return nil, "", errors.New("error user.")
	}

	if user.PeerId == currChannel.PeerIdA {
		previousNode = currChannel.PeerIdB
	} else {
		previousNode = currChannel.PeerIdA
	}

	// Transfer H or R increase step.
	htlcSingleHopPathInfo.CurrStep += 1
	// endregion

	// region Save private key to memory.
	if currChannel.PeerIdB == user.PeerId {
		tempAddrPrivateKeyMap[currChannel.PubKeyB] = reqData.ChannelAddressPrivateKey
	} else {
		tempAddrPrivateKeyMap[currChannel.PubKeyA] = reqData.ChannelAddressPrivateKey
	}

	tempAddrPrivateKeyMap[reqData.CurrHtlcTempAddressForHE1bPubKey] = reqData.CurrHtlcTempAddressForHE1bPrivateKey

	// Save pubkey to database.
	htlcSingleHopPathInfo.BobCurrHtlcTempForHt1bPubKey = reqData.CurrHtlcTempAddressForHE1bPubKey
	err = db.Update(htlcSingleHopPathInfo)
	if err != nil {
		log.Println(err.Error())
		return nil, "", err
	}
	// endregion

	// Generate response message.
	// If no error, the response data is displayed in websocket client of previous node.
	// Otherwise, it is displayed in websocket client of myself.
	responseData := make(map[string]interface{})
	responseData["id"] = rAndHInfo.Id
	responseData["request_hash"] = rAndHInfo.RequestHash
	responseData["r"] = rAndHInfo.R

	return responseData, previousNode, nil
}

// SignGetR
//
// Process type -47: Bob (middleman) check out if R is correct.
//  * R is <Preimage_R>
func (service *htlcBackwardTxManager) SignGetR(msgData string, user bean.User) (
	data map[string]interface{}, targetUser string, err error) {

	// if tool.CheckIsString(&msgData) == false {
	// 	return nil, "", errors.New("empty json data")
	// }

	// data = make(map[string]interface{})
	// data["approval"] = requestData.Approval
	// data["request_hash"] = requestData.RequestHash
	// return data, rAndHInfo.SenderPeerId, nil

	// 只有每个通道的转账发送方才能去创建关于R的交易
	// 如果转账发送方是PeerIdA（Alice），也就是Alice转账给bob，那就创建HED1a:Alice这边直接给钱bob，不需要时间锁定;HE1b,HERD1b，Bob因为是收款方，他自己的钱需要RSMC的方式锁在通道
	// 如果转账发送方是PeerIdB（Bob），也就是Bob转账给Alice，那就创建HE1a,HERD1a：Alice作为收款方，她得到的钱就需要RSMC锁定;HED1b：bob是发送方，他这边给Alice的钱是不需要锁定
	// 锁定的，就是自己的钱，就是给自己设定限制，给对方承诺

	return nil, "", nil
}

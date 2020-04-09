package lightclient

import (
	"encoding/json"
	"errors"
	"obd/bean/enum"
	"obd/service"
)

func routerOfP2PNode(msgType enum.MsgType, data string, client *Client) (retData string, retErr error) {
	defaultErr := errors.New("fail to deal msg in the inter node")
	status := false
	switch msgType {
	case enum.MsgType_ChannelOpen_N32:
		err := service.ChannelService.BeforeBobOpenChannelAtBobSide(data, client.User)
		if err == nil {
			status = true
		} else {
			defaultErr = err
		}
	case enum.MsgType_ChannelAccept_N33:
		_, err := service.ChannelService.AfterBobAcceptChannelAtAliceSide(data, client.User)
		if err == nil {
			status = true
		} else {
			defaultErr = err
		}
	case enum.MsgType_FundingCreate_BtcFundingCreated_N3400:
		_, err := service.FundingTransactionService.BeforeBobSignBtcFundingAtBobSide(data, client.User)
		if err == nil {
			status = true
		} else {
			defaultErr = err
		}
	case enum.MsgType_FundingSign_BtcSign_N3500:
		_, err := service.FundingTransactionService.AfterBobSignBtcFundingAtAliceSide(data, client.User)
		if err == nil {
			status = true
		} else {
			defaultErr = err
		}
	case enum.MsgType_FundingCreate_AssetFundingCreated_N34:
		_, err := service.FundingTransactionService.BeforeBobSignOmniFundingAtBobSide(data, client.User)
		if err == nil {
			status = true
		} else {
			defaultErr = err
		}
	case enum.MsgType_FundingSign_AssetFundingSigned_N35:
		node, err := service.FundingTransactionService.AfterBobSignOmniFundingAtAilceSide(data, client.User)
		if err == nil {
			status = true
			retData, _ := json.Marshal(node)
			return string(retData), nil
		} else {
			defaultErr = err
		}
	case enum.MsgType_CommitmentTx_CommitmentTransactionCreated_N351:
		node, err := service.CommitmentTxSignedService.BeforeBobSignCommitmentTranctionAtBobSide(data, client.User)
		if err == nil {
			status = true
			retData, _ := json.Marshal(node)
			return string(retData), nil
		}
		defaultErr = err
	case enum.MsgType_CommitmentTxSigned_ToAliceSign_N353:
		node, needNoticeAlice, err := service.CommitmentTxService.AfterBobSignCommitmentTranctionAtAliceSide(data, client.User)
		if needNoticeAlice {
			retAliceData := ""
			aliceDataStatus := false
			if err != nil {
				retAliceData = err.Error()
			} else {
				aliceDataStatus = true
				tempData, _ := json.Marshal(node["aliceData"])
				retAliceData = string(tempData)
			}
			client.sendToMyself(enum.MsgType_CommitmentTxSigned_RevokeAndAcknowledgeCommitmentTransaction_N352, aliceDataStatus, string(retAliceData))
		}
		if err == nil {
			status = true
			retBobData, _ := json.Marshal(node["bobData"])
			return string(retBobData), nil
		}
		defaultErr = err
	case enum.MsgType_CommitmentTxSigned_SecondToBobSign_N354:
		node, err := service.CommitmentTxSignedService.AfterAliceSignCommitmentTranctionAtBobSide(data, client.User)
		if err == nil {
			status = true
			retData, _ := json.Marshal(node)
			client.sendToMyself(enum.MsgType_CommitmentTxSigned_RevokeAndAcknowledgeCommitmentTransaction_N352, true, string(retData))
			return string(retData), nil
		}
		defaultErr = err
	case enum.MsgType_CloseChannelRequest_N38:
		node, err := service.ChannelService.BeforeBobSignCloseChannelAtBobSide(data, *client.User)
		if err == nil {
			status = true
			retData, _ := json.Marshal(node)
			return string(retData), nil
		}
		defaultErr = err
	case enum.MsgType_CloseChannelSign_N39:
		node, err := service.ChannelService.AfterBobSignCloseChannelAtAliceSide(data, *client.User)
		if err == nil {
			status = true
			retData, _ := json.Marshal(node)
			return string(retData), nil
		}
		defaultErr = err
	case enum.MsgType_HTLC_AddHTLC_N40:
		node, err := service.HtlcForwardTxService.BeforeBobSignPayerAddHtlcAtBobSide(data, *client.User)
		if err == nil {
			status = true
			retData, _ := json.Marshal(node)
			return string(retData), nil
		}
		defaultErr = err
	default:
		status = true
	}
	if status {
		defaultErr = nil
	}
	return "", defaultErr
}

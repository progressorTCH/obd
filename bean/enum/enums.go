package enum

type SendTargetType int

const (
	SendTargetType_SendToNone     SendTargetType = -1
	SendTargetType_SendToAll      SendTargetType = 0
	SendTargetType_SendToSomeone  SendTargetType = 1
	SendTargetType_SendToExceptMe SendTargetType = 2
)

type MsgType int

const (
	MsgType_Error_0      MsgType = 0
	MsgType_UserLogin_1  MsgType = 1
	MsgType_UserLogout_2 MsgType = 2

	MsgType_Core_GetNewAddress_1001                    MsgType = 1001
	MsgType_Core_GetMiningInfo_1002                    MsgType = 1002
	MsgType_Core_GetNetworkInfo_1003                   MsgType = 1003
	MsgType_Core_SignMessageWithPrivKey_1004           MsgType = 1004
	MsgType_Core_VerifyMessage_1005                    MsgType = 1005
	MsgType_Core_DumpPrivKey_1006                      MsgType = 1006
	MsgType_Core_ListUnspent_1007                      MsgType = 1007
	MsgType_Core_BalanceByAddress_1008                 MsgType = 1008
	MsgType_Core_FundingBTC_1009                       MsgType = 1009
	MsgType_Core_BtcCreateMultiSig_1010                MsgType = 1010
	MsgType_Core_Btc_ImportPrivKey_1011                MsgType = 1011
	MsgType_Core_Omni_Getbalance_1200                  MsgType = 1200
	MsgType_Core_Omni_CreateNewTokenFixed_1201         MsgType = 1201
	MsgType_Core_Omni_CreateNewTokenManaged_1202       MsgType = 1202
	MsgType_Core_Omni_GrantNewUnitsOfManagedToken_1203 MsgType = 1203
	MsgType_Core_Omni_RevokeUnitsOfManagedToken_1204   MsgType = 1204
	MsgType_Core_Omni_ListProperties_1205              MsgType = 1205
	MsgType_Core_Omni_GetTransaction_1206              MsgType = 1206

	MsgType_Core_Omni_FundingAsset_2001 MsgType = 2001

	MsgType_GetMnemonic_101                MsgType = 101
	MsgType_Mnemonic_CreateAddress_N200    MsgType = -200
	MsgType_Mnemonic_GetAddressByIndex_201 MsgType = -201

	MsgType_ChannelOpen_N32                   MsgType = -32
	MsgType_ChannelOpen_ItemByTempId_N3201    MsgType = -3201
	MsgType_ChannelOpen_AllItem_N3202         MsgType = -3202
	MsgType_ChannelOpen_Count_N3203           MsgType = -3203
	MsgType_ChannelOpen_DelItemByTempId_N3204 MsgType = -3204
	MsgType_ForceCloseChannel_N3205           MsgType = -3205
	MsgType_GetChannelInfoByChanId_N3206      MsgType = -3206
	MsgType_GetChannelInfoByChanId_N3207      MsgType = -3207

	MsgType_ChannelAccept_N33 MsgType = -33

	MsgType_FundingCreate_AssetFundingCreated_N34 MsgType = -34
	MsgType_FundingCreate_BtcCreate_N3400         MsgType = -3400
	MsgType_FundingCreate_ItemByTempId_N3401      MsgType = -3401
	MsgType_FundingCreate_ItemById_N3402          MsgType = -3402
	MsgType_FundingCreate_ALlItem_N3403           MsgType = -3403
	MsgType_FundingCreate_Count_N3404             MsgType = -3404
	MsgType_FundingCreate_DelById_N3405           MsgType = -3405

	MsgType_FundingSign_AssetFundingSigned_N35 MsgType = -35
	MsgType_FundingSign_BtcSign_N3500          MsgType = -3500

	MsgType_CommitmentTx_CommitmentTransactionCreated_N351 MsgType = -351
	MsgType_CommitmentTx_ItemsByChanId_N35101              MsgType = -35101
	MsgType_CommitmentTx_ItemById_N35102                   MsgType = -35102
	MsgType_CommitmentTx_Count_N35103                      MsgType = -35103
	MsgType_CommitmentTx_LatestCommitmentTxByChanId_N35104 MsgType = -35104
	MsgType_CommitmentTx_LatestRDByChanId_N35105           MsgType = -35105
	MsgType_CommitmentTx_LatestBRByChanId_N35106           MsgType = -35106
	MsgType_SendBreachRemedyTransaction_N35107             MsgType = -35107
	MsgType_CommitmentTx_AllRDByChanId_N35108              MsgType = -35108
	MsgType_CommitmentTx_AllBRByChanId_N35109              MsgType = -35109
	MsgType_CommitmentTx_GetBroadcastCommitmentTx_N35110   MsgType = -35110
	MsgType_CommitmentTx_GetBroadcastRDTx_N35111           MsgType = -35111
	MsgType_CommitmentTx_GetBroadcastBRTx_N35112           MsgType = -35112

	MsgType_CommitmentTxSigned_RevokeAndAcknowledgeCommitmentTransaction_N352 MsgType = -352
	MsgType_CommitmentTxSigned_ItemByChanId_N35201                            MsgType = -35201
	MsgType_CommitmentTxSigned_ItemById_N35202                                MsgType = -35202
	MsgType_CommitmentTxSigned_Count_N35203                                   MsgType = -35203

	MsgType_GetBalanceRequest_N353 MsgType = -353
	MsgType_GetBalanceRespond_N354 MsgType = -354

	MsgType_CloseChannelRequest_N38 MsgType = -38
	MsgType_CloseChannelSign_N39    MsgType = -39

	MsgType_HTLC_Invoice_N4003              MsgType = -4003
	MsgType_HTLC_AddHTLC_N40                MsgType = -40
	MsgType_HTLC_CreatedRAndHInfoList_N4001 MsgType = -4001
	MsgType_HTLC_CreatedRAndHInfoItem_N4002 MsgType = -4002
	MsgType_HTLC_AddHTLCSigned_N41          MsgType = -41
	MsgType_HTLC_SignedRAndHInfoList_N4101  MsgType = -4101
	MsgType_HTLC_SignedRAndHInfoItem_N4102  MsgType = -4102
	MsgType_HTLC_GetRFromLCommitTx_N4103    MsgType = -4103
	MsgType_HTLC_GetPathInfoByH_N4104       MsgType = -4104
	MsgType_HTLC_GetRInfoByHOfOwner_N4105   MsgType = -4105
	MsgType_HTLC_FindPathAndSendH_N42       MsgType = -42
	MsgType_HTLC_SendH_N43                  MsgType = -43
	MsgType_HTLC_SignGetH_N44               MsgType = -44
	MsgType_HTLC_CreateCommitmentTx_N45     MsgType = -45
	MsgType_HTLC_SendR_N46                  MsgType = -46
	MsgType_HTLC_VerifyR_N47                MsgType = -47
	MsgType_HTLC_RequestCloseCurrTx_N48     MsgType = -48
	MsgType_HTLC_CloseSigned_N49            MsgType = -49

	//https://github.com/LightningOnOmnilayer/Omni-BOLT-spec/blob/master/OmniBOLT-05-Atomic-Swap-among-Channels.md
	MsgType_Atomic_Swap_N80        MsgType = -80
	MsgType_Atomic_Swap_Accept_N81 MsgType = -81
)

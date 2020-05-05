package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

var ValidCreateProjectMsg = MsgCreateProject{
	SignBytes:  "SampleSignBytes",
	TxHash:     "SampleTxBytes",
	SenderDid:  "SenderDid",
	ProjectDid: "ProjectDid",
	PubKey:     "PubKey",
	Data: ProjectDoc{
		NodeDid:              "nodeDid",
		RequiredClaims:       "3",
		EvaluatorPayPerClaim: "2",
		ServiceEndpoint:      "https://google.co.in",
		CreatedOn:            "time1",
		CreatedBy:            "time2",
		Status:               "CREATED",
	},
}

var ValidUpdateProjectMsg = MsgCreateProject{
	SignBytes:  "UpdatedSignBytes",
	TxHash:     "UpdatedTxBytes",
	SenderDid:  "SenderDid",
	ProjectDid: "ProjectDid",
	PubKey:     "PubKey",
	Data: ProjectDoc{
		NodeDid:              "nodeDid",
		RequiredClaims:       "3",
		EvaluatorPayPerClaim: "2",
		ServiceEndpoint:      "https://google.co.in",
		CreatedOn:            "time1",
		CreatedBy:            "time2",
		Status:               "PENDING",
	},
}

var ValidWithdrawalInfo = WithdrawalInfo{
	ActionID:     "1",
	ProjectDid:   "6iftm1hHdaU6LJGKayRMev",
	RecipientDid: "6iftm1hHdaU6LJGKayRMev",
	Amount:       sdk.NewCoin(ixo.IxoNativeToken, sdk.NewInt(10)),
}

var (
	ValidAddress1, _ = sdk.AccAddressFromHex("0F6A8D732716BA24B213D7C28984FBE1248D009D")
	ValidAccId1      = InternalAccountID(ValidAddress1.String())
)

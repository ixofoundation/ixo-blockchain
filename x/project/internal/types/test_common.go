package types

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
)

var (
	ProjectDid       = "ProjectDid"
	ValidAddress1, _ = sdk.AccAddressFromHex("0F6A8D732716BA24B213D7C28984FBE1248D009D")
	ValidAccId1      = InternalAccountID(ValidAddress1.String())
)

var validProjectData = struct {
	NodeDid              string
	RequiredClaims       string
	EvaluatorPayPerClaim string
	ServiceEndpoint      string
	CreatedOn            string
	CreatedBy            string
}{
	NodeDid:              "nodeDid",
	RequiredClaims:       "3",
	EvaluatorPayPerClaim: "2",
	ServiceEndpoint:      "https://google.co.in",
	CreatedOn:            "time1",
	CreatedBy:            "time2",
}

var ValidProjectDoc = ProjectDoc{
	TxHash:     "SampleTxHash",
	SenderDid:  "SenderDid",
	ProjectDid: ProjectDid,
	PubKey:     "PubKey",
	Status:     "CREATED",
	Data:       MustMarshalJson(validProjectData),
}

var ValidUpdatedProjectDoc = ProjectDoc{
	TxHash:     "SampleTxHash",
	SenderDid:  "SenderDid",
	ProjectDid: ProjectDid,
	PubKey:     "PubKey",
	Status:     "PENDING", // Updated
	Data:       MustMarshalJson(validProjectData),
}

var ValidCreateProjectMsg = MsgCreateProject{
	TxHash:     "SampleTxBytes",
	SenderDid:  "SenderDid",
	ProjectDid: ProjectDid,
	PubKey:     "PubKey",
	Data:       MustMarshalJson(validProjectData),
}

var ValidWithdrawalInfo = WithdrawalInfo{
	ActionID:     "1",
	ProjectDid:   "6iftm1hHdaU6LJGKayRMev",
	RecipientDid: "6iftm1hHdaU6LJGKayRMev",
	Amount:       sdk.NewCoin(ixo.IxoNativeToken, sdk.NewInt(10)),
}

func MustMarshalJson(v interface{}) []byte {
	bz, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return bz
}

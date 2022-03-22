package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ixotypes "github.com/ixofoundation/ixo-blockchain/lib/types"
)

var (
	ProjectDid    = "did:ixo:U7GK8p8rVhJMKhBVRCJJ8c"
	UserDid       = "did:ixo:4XJLBfGtWSGKSz4BeRxdun"
	ProjectPubKey = "FmwNAfvV2xEqHwszrVJVBR3JgQ8AFCQEVzo1p6x4L8VW"

	ValidAddress1, _ = sdk.AccAddressFromHex("0F6A8D732716BA24B213D7C28984FBE1248D009D")
	ValidAccId1      = InternalAccountID(ValidAddress1.String())
)

var validProjectData = struct {
	Field1 string
	Field2 string
	FieldN string
	Fees   ProjectFeesMap `json:"fees" yaml:"fees"`
}{
	Field1: "field 1 value",
	Field2: "field 2 value",
	FieldN: "field N value",
	Fees: ProjectFeesMap{
		Context: "",
		Items:   nil,
	},
}

var ValidProjectDoc = ProjectDoc{
	TxHash:     "DummyTxHash",
	SenderDid:  UserDid,
	ProjectDid: ProjectDid,
	PubKey:     ProjectPubKey,
	Status:     string(CreatedProject),
	Data:       MustMarshalJson(validProjectData),
}

var ValidUpdatedProjectDoc = ProjectDoc{
	TxHash:     "DummyTxHash",
	SenderDid:  UserDid,
	ProjectDid: ProjectDid,
	PubKey:     ProjectPubKey,
	Status:     string(PendingStatus),
	Data:       MustMarshalJson(validProjectData),
}

var ValidCreateProjectMsg = MsgCreateProject{
	TxHash:     "DummyTxHash",
	SenderDid:  UserDid,
	ProjectDid: ProjectDid,
	PubKey:     ProjectPubKey,
	Data:       MustMarshalJson(validProjectData),
}

var ValidWithdrawalInfo = WithdrawalInfoDoc{
	ProjectDid:   ProjectDid,
	RecipientDid: UserDid,
	Amount:       sdk.NewCoin(ixotypes.IxoNativeToken, sdk.NewInt(10)),
}

func MustMarshalJson(v interface{}) []byte {
	bz, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return bz
}

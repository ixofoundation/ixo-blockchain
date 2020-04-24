package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-cosmos/x/ixo/sovrin"
)

func NewMsgSend(toDid string, amount sdk.Coins, senderDid sovrin.SovrinDid) MsgSend {
	return MsgSend{
		SignBytes: "",
		PubKey:    senderDid.VerifyKey,
		FromDid:   senderDid.Did,
		ToDid:     toDid,
		Amount:    amount,
	}
}

func CheckNotEmpty(value string, name string) (valid bool, err sdk.Error) {
	if len(value) == 0 {
		return false, sdk.ErrUnknownRequest(name + " is empty.")
	} else {
		return true, nil
	}
}

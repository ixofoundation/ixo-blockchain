package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"strings"
)

func NewMsgCreateBond(senderDid did.Did, bondDoc BondDoc, bondDid did.IxoDid) MsgCreateBond {
	return MsgCreateBond{
		TxHash:    "",
		SenderDid: senderDid,
		BondDid:   bondDid.Did,
		PubKey:    bondDid.VerifyKey,
		Data:      bondDoc,
	}
}

func NewMsgUpdateBondStatus(senderDid did.Did, updateBondStatusDoc UpdateBondStatusDoc, bondDid did.IxoDid) MsgUpdateBondStatus {
	return MsgUpdateBondStatus{
		SenderDid: senderDid,
		BondDid:   bondDid.Did,
		Data:      updateBondStatusDoc,
	}
}

func CheckNotEmpty(value string, name string) (valid bool, err sdk.Error) {
	if strings.TrimSpace(value) == "" {
		return false, sdk.ErrUnknownRequest(name + " is empty.")
	} else {
		return true, nil
	}
}

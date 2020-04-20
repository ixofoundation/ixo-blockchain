package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-cosmos/x/ixo/sovrin"
)

func NewMsgCreateBond(bondDoc BondDoc, bondDid sovrin.SovrinDid) MsgCreateBond {
	return MsgCreateBond{
		SignBytes: "",
		TxHash:    "",
		SenderDid: "",
		BondDid:   bondDid.Did,
		PubKey:    bondDid.VerifyKey,
		Data:      bondDoc,
	}
}

func NewMsgUpdateBondStatus(txHash string, senderDid string, updateBondStatusDoc UpdateBondStatusDoc, projectDid sovrin.SovrinDid) MsgUpdateBondStatus {
	return MsgUpdateBondStatus{
		SignBytes: "",
		TxHash:    txHash,
		SenderDid: senderDid,
		BondDid:   projectDid.Did,
		Data:      updateBondStatusDoc,
	}
}

func CheckNotEmpty(value string, name string) (valid bool, err sdk.Error) {
	if len(value) == 0 {
		return false, sdk.ErrUnknownRequest(name + " is empty.")
	} else {
		return true, nil
	}
}

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"strings"
)

func NewMsgSend(toDid did.Did, amount sdk.Coins, senderDid did.Did) MsgSend {
	return MsgSend{
		FromDid: senderDid,
		ToDid:   toDid,
		Amount:  amount,
	}
}

func NewMsgOracleTransfer(fromDid, toDid did.Did, amount sdk.Coins,
	oracleDid did.Did, proof string) MsgOracleTransfer {
	return MsgOracleTransfer{
		OracleDid: oracleDid,
		FromDid:   fromDid,
		ToDid:     toDid,
		Amount:    amount,
		Proof:     proof,
	}
}

func NewMsgOracleMint(toDid did.Did, amount sdk.Coins,
	oracleDid did.Did, proof string) MsgOracleMint {
	return MsgOracleMint{
		OracleDid: oracleDid,
		ToDid:     toDid,
		Amount:    amount,
		Proof:     proof,
	}
}

func NewMsgOracleBurn(fromDid did.Did, amount sdk.Coins,
	oracleDid did.Did, proof string) MsgOracleBurn {
	return MsgOracleBurn{
		OracleDid: oracleDid,
		FromDid:   fromDid,
		Amount:    amount,
		Proof:     proof,
	}
}

func CheckNotEmpty(value string, name string) (valid bool, err sdk.Error) {
	if strings.TrimSpace(value) == "" {
		return false, sdk.ErrUnknownRequest(name + " is empty.")
	} else {
		return true, nil
	}
}

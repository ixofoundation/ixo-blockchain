package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"strings"

	"github.com/ixofoundation/ixo-blockchain/x/ixo/sovrin"
)

func NewMsgSend(toDid ixo.Did, amount sdk.Coins, senderDid sovrin.SovrinDid) MsgSend {
	return MsgSend{
		PubKey:  senderDid.VerifyKey,
		FromDid: senderDid.Did,
		ToDid:   toDid,
		Amount:  amount,
	}
}

func NewMsgOracleTransfer(fromDid, toDid ixo.Did, amount sdk.Coins,
	oracleDid sovrin.SovrinDid, proof string) MsgOracleTransfer {
	return MsgOracleTransfer{
		PubKey:    oracleDid.VerifyKey,
		OracleDid: oracleDid.Did,
		FromDid:   fromDid,
		ToDid:     toDid,
		Amount:    amount,
		Proof:     proof,
	}
}

func NewMsgOracleMint(toDid ixo.Did, amount sdk.Coins,
	oracleDid sovrin.SovrinDid, proof string) MsgOracleMint {
	return MsgOracleMint{
		PubKey:    oracleDid.VerifyKey,
		OracleDid: oracleDid.Did,
		ToDid:     toDid,
		Amount:    amount,
		Proof:     proof,
	}
}

func NewMsgOracleBurn(fromDid ixo.Did, amount sdk.Coins,
	oracleDid sovrin.SovrinDid, proof string) MsgOracleBurn {
	return MsgOracleBurn{
		PubKey:    oracleDid.VerifyKey,
		OracleDid: oracleDid.Did,
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

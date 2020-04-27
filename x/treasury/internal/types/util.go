package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"

	"github.com/ixofoundation/ixo-cosmos/x/ixo/sovrin"
)

func NewMsgSend(toDid ixo.Did, amount sdk.Coins, senderDid sovrin.SovrinDid) MsgSend {
	return MsgSend{
		SignBytes: "",
		PubKey:    senderDid.VerifyKey,
		FromDid:   senderDid.Did,
		ToDid:     toDid,
		Amount:    amount,
	}
}

func NewMsgSendOnBehalfOf(fromDid, toDid ixo.Did, amount sdk.Coins, oracleDid sovrin.SovrinDid) MsgOracleTransfer {
	return MsgOracleTransfer{
		SignBytes: "",
		PubKey:    oracleDid.VerifyKey,
		OracleDid: oracleDid.Did,
		FromDid:   fromDid,
		ToDid:     toDid,
		Amount:    amount,
	}
}

func NewMsgMint(toDid ixo.Did, amount sdk.Coins, oracleDid sovrin.SovrinDid) MsgOracleMint {
	return MsgOracleMint{
		SignBytes: "",
		PubKey:    oracleDid.VerifyKey,
		OracleDid: oracleDid.Did,
		ToDid:     toDid,
		Amount:    amount,
	}
}

func NewMsgBurn(fromDid ixo.Did, amount sdk.Coins, oracleDid sovrin.SovrinDid) MsgOracleBurn {
	return MsgOracleBurn{
		SignBytes: "",
		PubKey:    oracleDid.VerifyKey,
		OracleDid: oracleDid.Did,
		FromDid:   fromDid,
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

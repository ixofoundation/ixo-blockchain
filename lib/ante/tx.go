package ante

import (
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
)

type IxoFeeTx struct {
	signing.SigVerifiableTx
}

func (tx *IxoFeeTx) FeePayer() []IidTxMsg {
	var msgs []IxoFeeTx

	return msgs
}

package ante

import "github.com/cosmos/cosmos-sdk/x/auth/signing"

type IidTx struct {
	signing.SigVerifiableTx
}

func(tx *IidTx) GetIidControllers() []IidTxMsg {
	var msgs []IidMsg

	for _, txMsg := range tx.GetMsgs() {
		iidMsg, ok := txMsg.(IidTxMsg);
		if !ok {
				continue
		}
		append(msgs, iidMsg)
	}
}

func(tx *IidTx) VerifyIidControllersAgainstSigniture() error {

}
VerifyIidControllersAgainstSigniture() error

type IidTxMsg interface {
	GetIidController() string
}

package ante

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	iidkeeper "github.com/ixofoundation/ixo-blockchain/x/iid/keeper"
)

type IidTx struct {
	signing.SigVerifiableTx
}

func (tx *IidTx) GetIidControllers() []IidTxMsg {
	var msgs []IidTxMsg

	for _, txMsg := range tx.GetMsgs() {
		iidMsg, ok := txMsg.(IidTxMsg)
		if !ok {
			continue
		}
		msgs = append(msgs, iidMsg)
	}

	return msgs
}

func (tx *IidTx) VerifyIidControllersAgainstSigniture(ctx sdk.Context, iidKeeper iidkeeper.Keeper) error {

	pubKeys, err := tx.GetPubKeys()
	if err != nil {
		return sdkerrors.Wrap(err, "Tx must be a IIDTx")
	}

	// for _, pk := range pubKeys {

	// 	if

	// }

	iidHasPubKey := false

	for _, iidMsg := range tx.GetIidControllers() {
		iid := iidMsg.GetIidController()
		iidDoc, exists := iidKeeper.GetDidDocument(ctx, []byte(iid))
		if !exists {
			return sdkerrors.Wrap(errors.New("iid not found"), "iid not found")
		}

		for _, pk := range pubKeys {
			if iidDoc.HasPublicKey(pk) {
				iidHasPubKey = true
			}
		}
	}

	if !iidHasPubKey {
		return sdkerrors.Wrap(errors.New("iid does not match public key in signiture"), "iid does not match public key in signiture")
	}

	return nil
}

type IidTxMsg interface {
	GetIidController() string
}

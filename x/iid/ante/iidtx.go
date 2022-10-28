package ante

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	iidkeeper "github.com/ixofoundation/ixo-blockchain/x/iid/keeper"
)

// type IidTx struct {
// 	signing.SigVerifiableTx
// }

func GetIidControllers(tx signing.SigVerifiableTx) []IidTxMsg {
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

func VerifyIidControllersAgainstSigniture(tx signing.SigVerifiableTx, ctx sdk.Context, iidKeeper iidkeeper.Keeper) error {

	pubKeys, err := tx.GetPubKeys()
	if err != nil {
		return sdkerrors.Wrap(err, "Tx must be a IIDTx")
	}

	iidHasPubKey := false
	controllers := GetIidControllers(tx)

	for _, iidMsg := range controllers {
		iid := iidMsg.GetIidController()
		iidDoc, exists := iidKeeper.GetDidDocument(ctx, []byte(iid))

		if !exists {
			return sdkerrors.Wrap(errors.New("iid not found"), "iid not found")
		}

		for _, pk := range pubKeys {
			if iidHasPubKey = iidDoc.HasPublicKey(pk); iidHasPubKey {
				break
			}
		}
	}

	if !iidHasPubKey && len(controllers) > 0 {
		return sdkerrors.Wrap(errors.New("iid does not match public key in signiture"), "iid does not match public key in signiture")
	}

	return nil
}

type IidTxMsg interface {
	GetIidController() string
}

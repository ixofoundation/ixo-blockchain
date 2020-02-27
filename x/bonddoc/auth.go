package bonddoc

import (
	"github.com/btcsuite/btcutil/base58"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-cosmos/x/bonddoc/internal/types"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

func NewAnteHandler(bonddocKeeper Keeper) sdk.AnteHandler {
	return func(ctx sdk.Context, tx sdk.Tx, simulate bool) (_ sdk.Context, _ sdk.Result, abort bool) {

		ixoTx, ok := tx.(ixo.IxoTx)
		if !ok {
			return ctx, sdk.ErrInternal("tx must be ixo.IxoTx").Result(), true
		}

		msg := ixoTx.GetMsgs()[0]
		bondMsg := msg.(types.BondMsg)
		pubKey := [32]byte{}

		if bondMsg.IsNewDid() {
			createBondMsg := msg.(types.CreateBondMsg)
			copy(pubKey[:], base58.Decode(createBondMsg.GetPubKey()))

		} else {
			bondDid := ixo.Did(msg.GetSigners()[0])
			bondDoc, err := bonddocKeeper.GetBondDoc(ctx, bondDid)
			if err != nil {
				return ctx, sdk.ErrInternal("bond did not found").Result(), false
			}

			copy(pubKey[:], base58.Decode(bondDoc.GetPubKey()))
		}

		var sigs = ixoTx.GetSignatures()
		if len(sigs) != 1 {
			return ctx,
				sdk.ErrUnauthorized("there can only be one signer").Result(),
				true
		}
		res := ixo.VerifySignature(msg, pubKey, sigs[0])

		if !res {
			return ctx, sdk.ErrInternal("Signature Verification failed").Result(), true
		}

		return ctx, sdk.Result{}, false // continue...

	}
}

package did

import (
	"github.com/btcsuite/btcutil/base58"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

func NewAnteHandler(didKeeper Keeper) sdk.AnteHandler {
	return func(
		ctx sdk.Context, tx sdk.Tx,
	) (_ sdk.Context, _ sdk.Result, abort bool) {

		// This always be a IxoTx
		ixoTx, ok := tx.(ixo.IxoTx)
		if !ok {
			return ctx, sdk.ErrInternal("tx must be ixo.IxoTx").Result(), true
		}

		msg := ixoTx.GetMsgs()[0]
		didMsg := msg.(DidMsg)
		pubKey := [32]byte{}

		if didMsg.IsNewDid() {
			addDidMsg := didMsg.(AddDidMsg)
			copy(pubKey[:], base58.Decode(addDidMsg.DidDoc.PubKey))
		} else {
			did := ixo.Did(msg.GetSigners()[0])
			didDoc := didKeeper.GetDidDoc(ctx, did)
			if didDoc == nil {
				return ctx,
					sdk.ErrUnauthorized("Issuer did not found").Result(),
					true
			}
			copy(pubKey[:], base58.Decode(didDoc.GetPubKey()))
		}

		// Assert that there are signatures.
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

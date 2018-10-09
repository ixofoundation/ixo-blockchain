package project

import (
	"github.com/btcsuite/btcutil/base58"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-cosmos/x/did"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

func NewAnteHandler(projectKeeper Keeper, didKeeper did.Keeper) sdk.AnteHandler {
	return func(
		ctx sdk.Context, tx sdk.Tx,
	) (_ sdk.Context, _ sdk.Result, abort bool) {

		// This always be a IxoTx
		ixoTx, ok := tx.(ixo.IxoTx)
		if !ok {
			return ctx, sdk.ErrInternal("tx must be ixo.IxoTx").Result(), true
		}

		msg := ixoTx.GetMsgs()[0]
		projectMsg := msg.(ProjectMsg)
		pubKey := [32]byte{}

		if projectMsg.IsNewDid() {
			createProjectMsg := msg.(CreateProjectMsg)
			//Get public key from payload
			copy(pubKey[:], base58.Decode(createProjectMsg.GetPubKey()))

		} else {
			projectDid := ixo.Did(msg.GetSigners()[0])
			// Get Project Doc
			projectDoc, found := projectKeeper.GetProjectDoc(ctx, projectDid)
			if !found {
				return ctx, sdk.ErrInternal("project did not found").Result(), true
			}
			copy(pubKey[:], base58.Decode(projectDoc.GetPubKey()))
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

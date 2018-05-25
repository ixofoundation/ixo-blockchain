package project

import (
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-cosmos/x/did"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

func NewAnteHandler(projectMapper SealedProjectMapper, didMapper did.SealedDidMapper) sdk.AnteHandler {
	return func(
		ctx sdk.Context, tx sdk.Tx,
	) (_ sdk.Context, _ sdk.Result, abort bool) {

		// This always be a IxoTx
		_, ok := tx.(ixo.IxoTx)
		if !ok {
			return ctx, sdk.ErrInternal("tx must be ixo.IxoTx").Result(), true
		}

		msg := tx.GetMsg()
		projectMsg := msg.(ProjectMsg)
		fmt.Println("Auth: check")
		pubKey := [32]byte{}

		if projectMsg.IsPegMsg() {
			pegDid := ixo.Did(msg.GetSigners()[0])
			didDoc := didMapper.GetDidDoc(ctx, pegDid)
			copy(pubKey[:], base58.Decode(didDoc.GetPubKey()))
		} else {
			if projectMsg.IsNewDid() {
				createProjectMsg := msg.(CreateProjectMsg)
				//Get public key from payload
				copy(pubKey[:], base58.Decode(createProjectMsg.ProjectDoc.PubKey))

			} else {
				projectDid := ixo.Did(msg.GetSigners()[0])
				// Get Project Doc
				projectDoc, found := projectMapper.GetProjectDoc(ctx, projectDid)
				if !found {
					return ctx, sdk.ErrInternal("project did not found").Result(), true
				}
				copy(pubKey[:], base58.Decode(projectDoc.GetPubKey()))
			}
		}
		// Assert that there are signatures.
		var sigs = tx.GetSignatures()
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

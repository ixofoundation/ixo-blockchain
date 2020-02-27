package bonds

import (
	"github.com/btcsuite/btcutil/base58"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-cosmos/x/bonds/internal/types"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

func NewAnteHandler(bondsKeeper Keeper) sdk.AnteHandler {
	return func(ctx sdk.Context, tx sdk.Tx, simulate bool) (_ sdk.Context, _ sdk.Result, abort bool) {

		ixoTx, ok := tx.(ixo.IxoTx)
		if !ok {
			return ctx, sdk.ErrInternal("tx must be ixo.IxoTx").Result(), true
		}

		msg := ixoTx.GetMsgs()[0]
		bondsMsg := msg.(types.BondsMsg)
		pubKey := [32]byte{}

		if bondsMsg.IsNewDid() {
			createBondMsg := msg.(types.MsgCreateBond)
			copy(pubKey[:], base58.Decode(createBondMsg.PubKey))
		} else {
			bondDid := ixo.Did(msg.GetSigners()[0])
			bond, found := bondsKeeper.GetBond(ctx, bondDid)
			if !found {
				return ctx, sdk.ErrInternal("bond not found").Result(), true
			}

			copy(pubKey[:], base58.Decode(bond.PubKey))
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

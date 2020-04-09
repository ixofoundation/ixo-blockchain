package bonds

import (
	"github.com/btcsuite/btcutil/base58"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-cosmos/x/bonds/internal/types"
	"github.com/ixofoundation/ixo-cosmos/x/did"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

func NewAnteHandler(bondsKeeper Keeper, didKeeper did.Keeper) sdk.AnteHandler {
	return func(ctx sdk.Context, tx sdk.Tx, simulate bool) (_ sdk.Context, _ sdk.Result, abort bool) {

		ixoTx, ok := tx.(ixo.IxoTx)
		if !ok {
			return ctx, sdk.ErrInternal("tx must be ixo.IxoTx").Result(), true
		}

		msg := ixoTx.GetMsgs()[0]
		pubKey := [32]byte{}
		var senderDid ixo.Did

		// Get signer PubKey and sender DID
		switch msg := msg.(type) {
		case types.MsgCreateBond:
			senderDid = msg.CreatorDid
			copy(pubKey[:], base58.Decode(msg.PubKey))
		case types.MsgEditBond:
			senderDid = msg.EditorDid
			bondDid := ixo.Did(msg.GetSigners()[0])
			bond, found := bondsKeeper.GetBond(ctx, bondDid)
			if !found {
				return ctx, sdk.ErrInternal("bond not found").Result(), true
			}
			copy(pubKey[:], base58.Decode(bond.PubKey))
		case types.MsgBuy:
			senderDid = msg.BuyerDid
			copy(pubKey[:], base58.Decode(msg.PubKey))
		case types.MsgSell:
			senderDid = msg.SellerDid
			copy(pubKey[:], base58.Decode(msg.PubKey))
		case types.MsgSwap:
			senderDid = msg.SwapperDid
			copy(pubKey[:], base58.Decode(msg.PubKey))
		default:
			panic("Unrecognized message type")
		}

		// Check that sender's DID is ledgered
		senderDidDoc, _ := didKeeper.GetDidDoc(ctx, senderDid)
		if senderDidDoc == nil {
			return ctx,
				sdk.ErrUnauthorized("Sender did not found").Result(),
				true
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

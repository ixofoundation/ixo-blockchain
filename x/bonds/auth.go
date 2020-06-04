package bonds

import (
	"github.com/btcsuite/btcutil/base58"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/internal/types"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
)

func GetPubKeyGetter(keeper Keeper, didKeeper did.Keeper) ixo.PubKeyGetter {
	return func(ctx sdk.Context, msg sdk.Msg) ([32]byte, sdk.Result) {

		// Get signer PubKey and sender DID
		var pubKey [32]byte
		var senderDid ixo.Did
		switch msg := msg.(type) {
		case types.MsgCreateBond:
			senderDid = msg.CreatorDid
			copy(pubKey[:], base58.Decode(msg.PubKey))
		case types.MsgEditBond:
			senderDid = msg.EditorDid
			bondDid := ixo.Did(msg.GetSigners()[0])
			bond, found := keeper.GetBond(ctx, bondDid)
			if !found {
				return pubKey, sdk.ErrInternal("bond not found").Result()
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
			return pubKey, sdk.ErrUnauthorized("Sender did not found").Result()
		}

		return pubKey, sdk.Result{}
	}
}

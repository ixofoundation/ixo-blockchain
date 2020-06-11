package bonddoc

import (
	"github.com/btcsuite/btcutil/base58"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-blockchain/x/bonddoc/internal/types"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
)

func GetPubKeyGetter(keeper Keeper) ixo.PubKeyGetter {
	return func(ctx sdk.Context, msg ixo.IxoMsg) ([32]byte, sdk.Result) {

		// Get signer PubKey
		var pubKey [32]byte
		switch msg := msg.(type) {
		case types.MsgCreateBond:
			copy(pubKey[:], base58.Decode(msg.GetPubKey()))
		case types.MsgUpdateBondStatus:
			bondDid := msg.GetSignerDid()
			bondDoc, err := keeper.GetBondDoc(ctx, bondDid)
			if err != nil {
				return pubKey, sdk.ErrInternal("bond did not found").Result()
			}
			copy(pubKey[:], base58.Decode(bondDoc.GetPubKey()))
		default:
			return pubKey, sdk.ErrUnknownRequest("No match for message type.").Result()
		}
		return pubKey, sdk.Result{}
	}
}

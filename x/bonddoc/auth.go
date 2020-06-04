package bonddoc

import (
	"github.com/btcsuite/btcutil/base58"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-blockchain/x/bonddoc/internal/types"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
)

func GetPubKeyGetter(keeper Keeper) ixo.PubKeyGetter {
	return func(ctx sdk.Context, msg sdk.Msg) ([32]byte, sdk.Result) {
		// Message must be a BondMsg
		bondMsg := msg.(types.BondMsg)

		// Get signer PubKey
		var pubKey [32]byte
		if bondMsg.IsNewDid() {
			createBondMsg := msg.(types.MsgCreateBond)
			copy(pubKey[:], base58.Decode(createBondMsg.GetPubKey()))
		} else {
			bondDid := ixo.Did(msg.GetSigners()[0])
			bondDoc, err := keeper.GetBondDoc(ctx, bondDid)
			if err != nil {
				return pubKey, sdk.ErrInternal("bond did not found").Result()
			}
			copy(pubKey[:], base58.Decode(bondDoc.GetPubKey()))
		}
		return pubKey, sdk.Result{}
	}
}

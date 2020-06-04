package did

import (
	"github.com/btcsuite/btcutil/base58"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/did/internal/types"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
)

func GetPubKeyGetter(keeper Keeper) ixo.PubKeyGetter {
	return func(ctx sdk.Context, msg sdk.Msg) ([32]byte, sdk.Result) {
		// Message must be a DidMsg
		didMsg := msg.(types.DidMsg)

		// Get signer PubKey
		var pubKey [32]byte
		if didMsg.IsNewDid() {
			addDidMsg := didMsg.(types.MsgAddDid)
			copy(pubKey[:], base58.Decode(addDidMsg.DidDoc.PubKey))
		} else {
			did := ixo.Did(msg.GetSigners()[0])
			didDoc, _ := keeper.GetDidDoc(ctx, did)
			if didDoc == nil {
				return pubKey, sdk.ErrUnauthorized("Issuer did not found").Result()
			}

			copy(pubKey[:], base58.Decode(didDoc.GetPubKey()))
		}
		return pubKey, sdk.Result{}
	}
}

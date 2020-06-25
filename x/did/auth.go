package did

import (
	"github.com/btcsuite/btcutil/base58"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/did/internal/types"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"github.com/tendermint/tendermint/crypto"
	ed25519Keys "github.com/tendermint/tendermint/crypto/ed25519"
)

func GetPubKeyGetter(keeper Keeper) ixo.PubKeyGetter {
	return func(ctx sdk.Context, msg ixo.IxoMsg) (pubKey crypto.PubKey, res sdk.Result) {

		// Get signer PubKey
		var pubKeyRaw [32]byte
		switch msg := msg.(type) {
		case types.MsgAddDid:
			copy(pubKeyRaw[:], base58.Decode(msg.DidDoc.PubKey))
		default:
			// For the remaining messages, the did is the signer
			didDoc, _ := keeper.GetDidDoc(ctx, msg.GetSignerDid())
			if didDoc == nil {
				return pubKey, sdk.ErrUnauthorized("Issuer did not found").Result()
			}
			copy(pubKeyRaw[:], base58.Decode(didDoc.GetPubKey()))
		}
		return ed25519Keys.PubKeyEd25519(pubKeyRaw), sdk.Result{}
	}
}

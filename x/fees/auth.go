package fees

import (
	"github.com/btcsuite/btcutil/base58"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"github.com/ixofoundation/ixo-blockchain/x/fees/internal/types"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
)

func GetPubKeyGetter(didKeeper did.Keeper) ixo.PubKeyGetter {
	return func(ctx sdk.Context, msg sdk.Msg) ([32]byte, sdk.Result) {
		// Message must be a FeesMsg
		feesMsg := msg.(types.FeesMsg)
		pubKey := [32]byte{}
		copy(pubKey[:], base58.Decode(feesMsg.GetPubKey()))

		// Check that sender's DID is ledgered
		senderDidDoc, _ := didKeeper.GetDidDoc(ctx, feesMsg.GetSenderDid())
		if senderDidDoc == nil {
			return pubKey, sdk.ErrUnauthorized("Sender did not found").Result()
		}

		return pubKey, sdk.Result{}
	}
}

package treasury

import (
	"github.com/btcsuite/btcutil/base58"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-blockchain/x/did"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"github.com/ixofoundation/ixo-blockchain/x/treasury/internal/types"
)

func GetPubKeyGetter(didKeeper did.Keeper) ixo.PubKeyGetter {
	return func(ctx sdk.Context, msg sdk.Msg) ([32]byte, sdk.Result) {
		// Message must be a TreasuryMsg
		treasuryMsg := msg.(types.TreasuryMessage)

		// Get signer PubKey
		var pubKey [32]byte
		copy(pubKey[:], base58.Decode(treasuryMsg.GetPubKey()))

		// Check that sender's DID is ledgered
		senderDidDoc, _ := didKeeper.GetDidDoc(ctx, treasuryMsg.GetSenderDid())
		if senderDidDoc == nil {
			return pubKey, sdk.ErrUnauthorized("Sender did not found").Result()
		}

		return pubKey, sdk.Result{}
	}
}

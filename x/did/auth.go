package did

import (
	"encoding/hex"
	"github.com/btcsuite/btcutil/base58"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

var (
	// simulation pubkey to estimate gas consumption
	simEd25519Pubkey ed25519.PubKey //PubKeyEd25519
)

func init() {
	// This decodes a valid hex string into a ed25519Pubkey for use in transaction simulation
	bz, _ := hex.DecodeString("035AD6810A47F073553FF30D2FCC7E0D3B1C0B74B61A1AAA2582344037151E14")
	copy(simEd25519Pubkey[:], bz)
}

func NewDefaultPubKeyGetter(keeper Keeper) ixo.PubKeyGetter {
	return func(ctx sdk.Context, msg ixo.IxoMsg) (pubKey crypto.PubKey, err error) {

		signerDidDoc, err := keeper.GetDidDoc(ctx, msg.GetSignerDid())
		if err != nil {
			return pubKey, sdkerrors.Wrap(ErrInvalidDid, "signer DID not found")
		}

		var pubKeyRaw ed25519.PubKey
		copy(pubKeyRaw[:], base58.Decode(signerDidDoc.GetPubKey()))
		return pubKeyRaw, nil
	}
}

func NewModulePubKeyGetter(keeper Keeper) ixo.PubKeyGetter {
	return func(ctx sdk.Context, msg ixo.IxoMsg) (pubKey crypto.PubKey, err error) {

		// MsgAddDid: pubkey from msg since user DID does not exist yet
		// Other: signer DID exists, so get pubkey from did module

		var pubKeyEd25519 ed25519.PubKey
		switch msg := msg.(type) {
		case MsgAddDid:
			copy(pubKeyEd25519[:], base58.Decode(msg.PubKey))
		default:
			return NewDefaultPubKeyGetter(keeper)(ctx, msg)
		}
		return pubKeyEd25519, nil
	}
}
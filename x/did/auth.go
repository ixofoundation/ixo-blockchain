package did

import (
	"encoding/hex"
	"github.com/btcsuite/btcutil/base58"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519" //"github.com/tendermint/tendermint/crypto/ed25519"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ixofoundation/ixo-blockchain/x/did/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/did/internal/types"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
)

var (
	// simulation pubkey to estimate gas consumption
	simEd25519Pubkey ed25519.PubKey //PubKeyEd25519
)

func init() {
	// This decodes a valid hex string into a ed25519Pubkey for use in transaction simulation
	bz, _ := hex.DecodeString("035AD6810A47F073553FF30D2FCC7E0D3B1C0B74B61A1AAA2582344037151E14")
	copy(simEd25519Pubkey.Key[:], bz)
}

func NewDefaultPubKeyGetter(keeper keeper.Keeper) ixo.PubKeyGetter {
	return func(ctx sdk.Context, msg ixo.IxoMsg) (pubKey cryptotypes.PubKey, err error) {

		signerDidDoc, err := keeper.GetDidDoc(ctx, msg.GetSignerDid())
		if err != nil {
			return pubKey, sdkerrors.Wrap(types.ErrInvalidDid, "signer DID not found")
		}

		var pubKeyRaw ed25519.PubKey
		pubKeyRaw.Key = base58.Decode(signerDidDoc.GetPubKey())
		//copy(pubKeyRaw.Key[:], base58.Decode(signerDidDoc.GetPubKey()))
		return &pubKeyRaw, nil
	}
}

func NewModulePubKeyGetter(keeper keeper.Keeper) ixo.PubKeyGetter {
	return func(ctx sdk.Context, msg ixo.IxoMsg) (pubKey cryptotypes.PubKey, err error) {

		// MsgAddDid: pubkey from msg since user DID does not exist yet
		// Other: signer DID exists, so get pubkey from did module

		var pubKeyEd25519 ed25519.PubKey
		switch msg := msg.(type) {
		case *types.MsgAddDid:
			pubKeyEd25519.Key = base58.Decode(msg.PubKey)
		default:
			return NewDefaultPubKeyGetter(keeper)(ctx, msg)
		}
		return &pubKeyEd25519, nil
	}
}

package did

import (
	"encoding/hex"
	"github.com/btcsuite/btcutil/base58"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

var (
	// simulation pubkey to estimate gas consumption
	simEd25519Pubkey ed25519.PubKeyEd25519
)

func init() {
	// This decodes a valid hex string into a ed25519Pubkey for use in transaction simulation
	bz, _ := hex.DecodeString("035AD6810A47F073553FF30D2FCC7E0D3B1C0B74B61A1AAA2582344037151E14")
	copy(simEd25519Pubkey[:], bz)
}

func GetPubKeyGetter(keeper Keeper) ixo.PubKeyGetter {
	return func(ctx sdk.Context, msg ixo.IxoMsg) (pubKey crypto.PubKey, err error) {

		// Get signer PubKey
		var pubKeyEd25519 ed25519.PubKeyEd25519
		switch msg := msg.(type) {
		case MsgAddDid:
			copy(pubKeyEd25519[:], base58.Decode(msg.PubKey))
		default:
			// For the remaining messages, the did is the signer
			didDoc, _ := keeper.GetDidDoc(ctx, msg.GetSignerDid())
			if didDoc == nil {
				return pubKey, sdkerrors.Wrap(ErrInvalidDid, "issuer did not found")

			}
			copy(pubKeyEd25519[:], base58.Decode(didDoc.GetPubKey()))
		}
		return pubKeyEd25519, nil
	}
}

func getAddDidSignBytes(chainID string, tx auth.StdTx, acc exported.Account, genesis bool) []byte {
	var accNum uint64
	if !genesis {
		// Fixed account number used so that sign bytes do not depend on it
		accNum = uint64(0)
	}

	return auth.StdSignBytes(
		chainID, accNum, acc.GetSequence(), tx.Fee, tx.Msgs, tx.Memo,
	)
}

func NewAddDidAnteHandler(ak auth.AccountKeeper, supplyKeeper supply.Keeper, pubKeyGetter ixo.PubKeyGetter) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		//ante.NewMempoolFeeDecorator(),
		ante.NewValidateBasicDecorator(),
		ante.NewValidateMemoDecorator(ak),
		ante.NewConsumeGasForTxSizeDecorator(ak),
		NewSetPubKeyDecorator(ak, pubKeyGetter), // SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewValidateSigCountDecorator(ak),
		ixo.NewDeductFeeDecorator(ak, supplyKeeper, pubKeyGetter),
		//ixo.NewSigGasConsumeDecorator(ak, sigGasConsumer, pubKeyGetter),
		NewSigVerificationDecorator(ak, pubKeyGetter),
		ixo.NewIncrementSequenceDecorator(ak, pubKeyGetter), // innermost AnteDecorator
	)
}

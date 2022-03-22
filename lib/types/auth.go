package types

import (
	"encoding/hex"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/crypto/types/multisig"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
	"github.com/spf13/pflag"
)

var (
	expectedMinGasPrices       = "0.025" + IxoNativeToken
	approximationGasAdjustment = float64(1.5)
	fees                       = "1000000" + IxoNativeToken
	// TODO: parameterise (or remove) hard-coded gas prices and adjustments

	// simulation signature values used to estimate gas consumption
	simEd25519Pubkey ed25519.PubKey
	simEd25519Sig    [ed25519.SignatureSize]byte
)

func init() {
	// This decodes a valid hex string into a ed25519Pubkey for use in transaction simulation
	bz, _ := hex.DecodeString("035AD6810A47F073553FF30D2FCC7E0D3B1C0B74B61A1AAA2582344037151E14")
	copy(simEd25519Pubkey.Key[:], bz)
}

type PubKeyGetter func(ctx sdk.Context, msg IxoMsg) (cryptotypes.PubKey, error)

func NewDefaultAnteHandler(ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
	sigGasConsumer ante.SignatureVerificationGasConsumer, pubKeyGetter PubKeyGetter,
	signModeHandler authsigning.SignModeHandler) sdk.AnteHandler {

	// Refer to inline documentation in app/app.go for introduction to why we
	// need a custom ixo AnteHandler. Below, we will discuss the differences
	// between the default Cosmos AnteHandler and our custom ixo AnteHandler.
	//
	// It is clear below that our custom AnteHandler is not completely custom.
	// It uses various functions from the Cosmos ante module. However, it also
	// uses customised decorators, without adding completely new decorators.
	// Below we present the differences in the customised decorators.
	//
	// In general:
	// - Enforces messages to be of type IxoMsg, to be used with pubKeyGetter.
	// - Does not allow for multiple messages (to be added in the future).
	// - Does not allow for multiple signatures (to be added in the future).
	//
	// NewSetPubKeyDecorator: as opposed to the Cosmos version...
	// - Gets signer pubkey from pubKeyGetter argument instead of tx signatures.
	// - Gets signer address from pubkey instead of the messages' GetSigners().
	// - Uses simEd25519Pubkey instead of simSecp256k1Pubkey for simulations.
	//
	// NewDeductFeeDecorator:
	// - Gets fee payer address from the pubkey obtained from pubKeyGetter
	//   instead of from the first message's GetSigners() function.
	//
	// NewSigGasConsumeDecorator:
	// - Gets the only signer address from the pubkey obtained from pubKeyGetter
	//   instead of from the messages' GetSigners() function.
	// - Uses simEd25519Pubkey instead of simSecp256k1Pubkey for simulations.
	//
	// NewSigVerificationDecorator:
	// - Gets the only signer address and account from the pubkey obtained from
	//   pubKeyGetter instead of from the messages' GetSigners() function.
	//
	// NewIncrementSequenceDecorator:
	// - Gets the only signer address from the pubkey obtained from pubKeyGetter
	//   instead of from the messages' GetSigners() function.

	return sdk.ChainAnteDecorators(
		ante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		ante.NewMempoolFeeDecorator(),
		ante.NewValidateBasicDecorator(),
		ante.NewValidateMemoDecorator(ak),
		ante.NewConsumeGasForTxSizeDecorator(ak),
		NewSetPubKeyDecorator(ak, pubKeyGetter), // SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewValidateSigCountDecorator(ak),
		NewDeductFeeDecorator(ak, bk, pubKeyGetter),
		NewSigGasConsumeDecorator(ak, sigGasConsumer, pubKeyGetter),
		NewSigVerificationDecorator(ak, signModeHandler, pubKeyGetter),
		NewIncrementSequenceDecorator(ak, pubKeyGetter), // innermost AnteDecorator
	)
}

func GenerateOrBroadcastTxCLI(clientCtx client.Context, flagSet *pflag.FlagSet, ixoDid exported.IxoDid, msg sdk.Msg) error {
	txf := tx.NewFactoryCLI(clientCtx, flagSet)
	return GenerateOrBroadcastTxWithFactory(clientCtx, txf, ixoDid, msg)
}

func GenerateOrBroadcastTxWithFactory(clientCtx client.Context, txf tx.Factory, ixoDid exported.IxoDid, msg sdk.Msg) error {
	if clientCtx.GenerateOnly {
		return tx.GenerateTx(clientCtx, txf, msg) // like old PrintUnsignedStdTx
	}

	return BroadcastTx(clientCtx, txf, ixoDid, msg)
}

// TODO: This function was a copy and paste of the the internal function which Cosmos made private.
// TODO: Function looked 99% the same so just wrapped the official implementation.
func BroadcastTx(clientCtx client.Context, txf tx.Factory, ixoDid exported.IxoDid, msg sdk.Msg) error {
	return tx.BroadcastTx(clientCtx, txf)
}

func checkMultipleSigners(mode signing.SignMode, tx authsigning.Tx) error {
	if mode == signing.SignMode_SIGN_MODE_DIRECT &&
		len(tx.GetSigners()) > 1 {
		return sdkerrors.Wrap(sdkerrors.ErrNotSupported, "Signing in DIRECT mode is only supported for transactions with one signer only")
	}
	return nil
}

/*
!! DONT think this is needed anymore.
func Sign(txf tx.Factory, clientCtx client.Context, txBuilder client.TxBuilder, overwriteSig bool, ixoDid exported.IxoDid) error {
	var privateKey ed25519.PrivKey
	privateKey.Key = append(base58.Decode(ixoDid.Secret.SignKey), base58.Decode(ixoDid.VerifyKey)...)

	signMode := txf.SignMode()
	if signMode == signing.SignMode_SIGN_MODE_UNSPECIFIED {
		// use the SignModeHandler's default mode if unspecified
		signMode = clientCtx.TxConfig.SignModeHandler().DefaultMode()
	}
	err := checkMultipleSigners(signMode, txBuilder.GetTx())
	if err != nil {
		return err
	}

	signerData := authsigning.SignerData{
		ChainID:       txf.ChainID(),
		AccountNumber: txf.AccountNumber(),
		Sequence:      txf.Sequence(),
	}

	// For SIGN_MODE_DIRECT, calling SetSignatures calls setSignerInfos on
	// TxBuilder under the hood, and SignerInfos is needed to generated the
	// sign bytes. This is the reason for setting SetSignatures here, with a
	// nil signature.
	//
	// Note: this line is not needed for SIGN_MODE_LEGACY_AMINO, but putting it
	// also doesn't affect its generated sign bytes, so for code's simplicity
	// sake, we put it here.
	sigData := signing.SingleSignatureData{
		SignMode:  signMode,
		Signature: nil,
	}
	sig := signing.SignatureV2{
		PubKey:   privateKey.PubKey(),
		Data:     &sigData,
		Sequence: txf.Sequence(),
	}
	var prevSignatures []signing.SignatureV2
	if !overwriteSig {
		prevSignatures, err = txBuilder.GetTx().GetSignaturesV2()
		if err != nil {
			return err
		}
	}
	if err := txBuilder.SetSignatures(sig); err != nil {
		return err
	}

	bytesToSign, err := clientCtx.TxConfig.SignModeHandler().GetSignBytes(signMode, signerData, txBuilder.GetTx())
	if err != nil {
		return err
	}

	sigBytes, err := privateKey.Sign(bytesToSign)
	if err != nil {
		return err
	}

	// Construct the SignatureV2 struct
	sigData = signing.SingleSignatureData{
		SignMode:  signMode,
		Signature: sigBytes,
	}
	sig = signing.SignatureV2{
		PubKey:   privateKey.PubKey(),
		Data:     &sigData,
		Sequence: txf.Sequence(),
	}

	if overwriteSig {
		return txBuilder.SetSignatures(sig)
	}
	prevSignatures = append(prevSignatures, sig)
	return txBuilder.SetSignatures(prevSignatures...)
}
*/

// Identical to DefaultSigVerificationGasConsumer, but with ed25519 allowed
func IxoSigVerificationGasConsumer(
	meter sdk.GasMeter, sig signing.SignatureV2, params authtypes.Params,
) error {
	pubkey := sig.PubKey
	switch pubkey := pubkey.(type) {
	case *ed25519.PubKey:
		meter.ConsumeGas(params.SigVerifyCostED25519, "ante verify: ed25519")
		return nil

	case *secp256k1.PubKey:
		meter.ConsumeGas(params.SigVerifyCostSecp256k1, "ante verify: secp256k1")
		return nil

	case multisig.PubKey:
		multisignature, ok := sig.Data.(*signing.MultiSignatureData)
		if !ok {
			return fmt.Errorf("expected %T, got, %T", &signing.MultiSignatureData{}, sig.Data)
		}
		err := ante.ConsumeMultisignatureVerificationGas(meter, multisignature, pubkey, params, sig.Sequence)
		if err != nil {
			return err
		}
		return nil

	default:
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidPubKey, "unrecognized public key type: %T", pubkey)
	}
}

/*
!! DONT Think this is need anymore.
func SignAndBroadcastTxFromStdSignMsg(clientCtx client.Context,
	msg sdk.Msg, ixoDid exported.IxoDid, flagSet *pflag.FlagSet) (*sdk.TxResponse, error) {

	txf := tx.NewFactoryCLI(clientCtx, flagSet)
	txf = txf.WithFees(fees).WithGasPrices("").WithGas(0)

	tx, err := tx.BuildUnsignedTx(txf, msg)
	if err != nil {
		return nil, err
	}

	if !clientCtx.SkipConfirm {
		out, err := clientCtx.TxConfig.TxJSONEncoder()(tx.GetTx())
		if err != nil {
			return nil, err
		}

		_, _ = fmt.Fprintf(os.Stderr, "%s\n\n", out)

		buf := bufio.NewReader(os.Stdin)
		ok, err := input.GetConfirmation("confirm transaction before signing and broadcasting", buf, os.Stderr)

		if err != nil || !ok {
			_, _ = fmt.Fprintf(os.Stderr, "%s\n", "cancelled transaction")
			return nil, err
		}
	}

	err = Sign(txf, clientCtx, tx, true, ixoDid)
	if err != nil {
		return nil, err
	}

	txBytes, err := clientCtx.TxConfig.TxEncoder()(tx.GetTx())
	if err != nil {
		return nil, err
	}

	// broadcast to a Tendermint node
	res, err := clientCtx.BroadcastTx(txBytes)
	if err != nil {
		return &sdk.TxResponse{}, err
	}

	return res, nil
}
*/

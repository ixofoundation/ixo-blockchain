package types

import (
	"bufio"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/cosmos/cosmos-sdk/client"
	//"github.com/cosmos/cosmos-sdk/client/context"
	//"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/input"
	"github.com/cosmos/cosmos-sdk/client/tx"
	//"github.com/cosmos/cosmos-sdk/codec"
	//"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	//"github.com/cosmos/cosmos-sdk/x/auth"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	//"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/spf13/pflag"

	//"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
	//"github.com/spf13/viper"
	"github.com/tendermint/ed25519"
	"github.com/tendermint/tendermint/crypto"
	ed25519tm "github.com/tendermint/tendermint/crypto/ed25519"
	//"github.com/tendermint/tendermint/crypto/multisig"
	"github.com/cosmos/cosmos-sdk/crypto/types/multisig"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"os"
)

var (
	expectedMinGasPrices       = "0.025" + IxoNativeToken
	approximationGasAdjustment = float64(1.5)
	// TODO: parameterise (or remove) hard-coded gas prices and adjustments

	// simulation signature values used to estimate gas consumption
	simEd25519Pubkey ed25519tm.PubKey
	simEd25519Sig    [ed25519.SignatureSize]byte
)

func init() {
	// This decodes a valid hex string into a ed25519Pubkey for use in transaction simulation
	bz, _ := hex.DecodeString("035AD6810A47F073553FF30D2FCC7E0D3B1C0B74B61A1AAA2582344037151E14")
	copy(simEd25519Pubkey[:], bz)
}

type PubKeyGetter func(ctx sdk.Context, msg IxoMsg) (crypto.PubKey, error)

func NewDefaultAnteHandler(ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
	sigGasConsumer ante.SignatureVerificationGasConsumer, pubKeyGetter PubKeyGetter,
	signModeHandler authsigning.SignModeHandler,) sdk.AnteHandler {
	//ak auth.AccountKeeper, supplyKeeper supply.Keeper, sigGasConsumer ante.SignatureVerificationGasConsumer, pubKeyGetter PubKeyGetter) sdk.AnteHandler {

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

//TODO uncomment
//func ApproximateFeeForTx(cliCtx context.CLIContext, tx auth.StdTx, chainId string) (auth.StdFee, error) {
//
//	// Set up a transaction builder
//	cdc := cliCtx.Codec
//	txEncoder := auth.DefaultTxEncoder
//	gasAdjustment := approximationGasAdjustment
//	fees := sdk.NewCoins(sdk.NewCoin(IxoNativeToken, sdk.OneInt()))
//	txBldr := auth.NewTxBuilder(txEncoder(cdc), 0, 0, 0, gasAdjustment, true, chainId, tx.Memo, fees, nil)
//
//	// Approximate gas consumption
//	txBldr, err := utils.EnrichWithGas(txBldr, cliCtx, tx.Msgs)
//	if err != nil {
//		return auth.StdFee{}, err
//	}
//
//	// Clear fees and set gas-prices to deduce updated fee = (gas * gas-prices)
//	signMsg, err := txBldr.WithFees("").WithGasPrices(expectedMinGasPrices).BuildSignMsg(tx.Msgs)
//	if err != nil {
//		return auth.StdFee{}, err
//	}
//
//	return signMsg.Fee, nil
//}

func GenerateOrBroadcastTxCLI(clientCtx client.Context, flagSet *pflag.FlagSet, ixoDid exported.IxoDid, msg sdk.Msg,) error {
	txf := tx.NewFactoryCLI(clientCtx, flagSet)
	return GenerateOrBroadcastTxWithFactory(clientCtx, txf, ixoDid, msg)
}

func GenerateOrBroadcastTxWithFactory(clientCtx client.Context, txf tx.Factory, ixoDid exported.IxoDid, msg sdk.Msg) error {
	if clientCtx.GenerateOnly {
		return tx.GenerateTx(clientCtx, txf, msg) // like old PrintUnsignedStdTx
	}

	return BroadcastTx(clientCtx, txf, ixoDid, msg)
}

//func GenerateOrBroadcastMsgs(cliCtx context.CLIContext, msg sdk.Msg, ixoDid exported.IxoDid) error {
//	msgs := []sdk.Msg{msg}
//	txBldr := auth.NewTxBuilderFromCLI(cliCtx.Input).WithTxEncoder(utils.GetTxEncoder(cliCtx.Codec))
//
//	if cliCtx.GenerateOnly {
//		return utils.PrintUnsignedStdTx(txBldr, cliCtx, msgs)
//	}
//
//	return CompleteAndBroadcastTxCLI(txBldr, cliCtx, msgs, ixoDid)
//}

// GasEstimateResponse defines a response definition for tx gas estimation.
type GasEstimateResponse struct {
	GasEstimate uint64 `json:"gas_estimate" yaml:"gas_estimate"`
}

// like CompleteAndBroadcastTxCLI (copied from cosmos-sdk)
func BroadcastTx(clientCtx client.Context, txf tx.Factory, ixoDid exported.IxoDid, msg sdk.Msg) error {
	txf, err := tx.PrepareFactory(clientCtx, txf)
	if err != nil {
		return err
	}

	if txf.SimulateAndExecute() || clientCtx.Simulate {
		_, adjusted, err := tx.CalculateGas(clientCtx.QueryWithData, txf, msg)
		if err != nil {
			return err
		}

		txf = txf.WithGas(adjusted)
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", GasEstimateResponse{GasEstimate: txf.Gas()})
	}

	if clientCtx.Simulate {
		return nil
	}

	tx, err := tx.BuildUnsignedTx(txf, msg) //like old BuildSignMsg
	if err != nil {
		return err
	}

	if !clientCtx.SkipConfirm {
		out, err := clientCtx.TxConfig.TxJSONEncoder()(tx.GetTx())
		if err != nil {
			return err
		}

		_, _ = fmt.Fprintf(os.Stderr, "%s\n\n", out)

		buf := bufio.NewReader(os.Stdin)
		ok, err := input.GetConfirmation("confirm transaction before signing and broadcasting", buf, os.Stderr)

		if err != nil || !ok {
			_, _ = fmt.Fprintf(os.Stderr, "%s\n", "cancelled transaction")
			return err
		}
	}

	err = Sign(txf, clientCtx.GetFromName(), tx, ixoDid, msg) //like old local BuildAndSign
	if err != nil {
		return err
	}

	txBytes, err := clientCtx.TxConfig.TxEncoder()(tx.GetTx())
	if err != nil {
		return err
	}

	// broadcast to a Tendermint node
	res, err := clientCtx.BroadcastTx(txBytes)
	if err != nil {
		return err
	}

	return clientCtx.PrintOutput(res)
}

//func CompleteAndBroadcastTxCLI(txBldr auth.TxBuilder, cliCtx context.CLIContext, msgs []sdk.Msg, ixoDid exported.IxoDid) error {
//	txBldr, err := utils.PrepareTxBuilder(txBldr, cliCtx)
//	if err != nil {
//		return err
//	}
//
//	if txBldr.SimulateAndExecute() || cliCtx.Simulate {
//		txBldr, err = utils.EnrichWithGas(txBldr, cliCtx, msgs)
//		if err != nil {
//			return err
//		}
//
//		gasEst := utils.GasEstimateResponse{GasEstimate: txBldr.Gas()}
//		_, _ = fmt.Fprintf(os.Stderr, "%s\n", gasEst.String())
//	}
//
//	if cliCtx.Simulate {
//		return nil
//	}
//
//	if !cliCtx.SkipConfirm {
//		stdSignMsg, err := txBldr.BuildSignMsg(msgs) //like new BuildUnsignedTx
//		if err != nil {
//			return err
//		}
//
//		var json []byte
//		if viper.GetBool(flags.FlagIndentResponse) {
//			json, err = cliCtx.Codec.MarshalJSONIndent(stdSignMsg, "", "  ")
//			if err != nil {
//				panic(err)
//			}
//		} else {
//			json = cliCtx.Codec.MustMarshalJSON(stdSignMsg)
//		}
//
//		_, _ = fmt.Fprintf(os.Stderr, "%s\n\n", json)
//
//		buf := bufio.NewReader(os.Stdin)
//		ok, err := input.GetConfirmation("confirm transaction before signing and broadcasting", buf)
//		if err != nil || !ok {
//			_, _ = fmt.Fprintf(os.Stderr, "%s\n", "cancelled transaction")
//			return err
//		}
//	}
//
//	// build and sign the transaction
//	txBytes, err := BuildAndSign(txBldr, cliCtx, msgs, ixoDid)
//	if err != nil {
//		return err
//	}
//
//	// broadcast to a Tendermint node
//	res, err := cliCtx.BroadcastTx(txBytes)
//	if err != nil {
//		return err
//	}
//
//	return cliCtx.PrintOutput(res)
//}

// copied from cosmos-sdk
func Sign(txf tx.Factory, name string, txBuilder client.TxBuilder, ixoDid exported.IxoDid, msg sdk.Msg) error {
	if txf.Keybase() == nil {
		return errors.New("keybase must be set prior to signing a transaction")
	}

	signMode := txf.SignMode()
	if signMode == signing.SignMode_SIGN_MODE_UNSPECIFIED {
		// use the SignModeHandler's default mode if unspecified
		signMode = signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON //TODO - no access to txf.txConfig.SignModeHandler().DefaultMode()
	}

	//key, err := txf.keybase.Key(name)
	//if err != nil {
	//	return err
	//}

	var privateKey ed25519tm.PrivKey
	copy(privateKey[:], base58.Decode(ixoDid.Secret.SignKey))
	copy(privateKey[32:], base58.Decode(ixoDid.VerifyKey))

	pubKey := privateKey.PubKey() //key.GetPubKey()
	//signerData := authsigning.SignerData{
	//	ChainID:       txf.ChainID(),
	//	AccountNumber: txf.AccountNumber(),
	//	Sequence:      txf.Sequence(),
	//}

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
	sigv2 := signing.SignatureV2{
		PubKey:   pubKey,
		Data:     &sigData,
		Sequence: txf.Sequence(),
	}
	if err := txBuilder.SetSignatures(sigv2); err != nil {
		return err
	}

	// Generate the bytes to be signed.
	//signBytes, err := txf.txConfig.SignModeHandler().GetSignBytes(signMode, signerData, txBuilder.GetTx())
	//if err != nil {
	//	return err
	//}

	sigBytes, err := privateKey.Sign(msg.GetSignBytes())
	if err != nil {
		return err
	} // WHERE TO USE?

	// Sign those bytes
	//sigBytes, _, err := txf.Keybase().Sign(name, signBytes)
	//if err != nil {
	//	return err
	//}

	// Construct the SignatureV2 struct
	sigData = signing.SingleSignatureData{
		SignMode:  signMode,
		Signature: sigBytes,
	}
	sigv2 = signing.SignatureV2{
		PubKey:   pubKey,
		Data:     &sigData,
		Sequence: txf.Sequence(),
	}

	// And here the tx is populated with the signature
	return txBuilder.SetSignatures(sigv2)
}

//func Sign(cliCtx context.CLIContext, msg auth.StdSignMsg, ixoDid exported.IxoDid) ([]byte, error) {
//	var privateKey ed25519tm.PrivKeyEd25519
//	copy(privateKey[:], base58.Decode(ixoDid.Secret.SignKey))
//	copy(privateKey[32:], base58.Decode(ixoDid.VerifyKey))
//
//	sig, err := MakeSignature(msg.Bytes(), privateKey)
//	if err != nil {
//		return nil, err
//	}
//
//	encoder := utils.GetTxEncoder(cliCtx.Codec)
//	return encoder(auth.NewStdTx(msg.Msgs, msg.Fee, []auth.StdSignature{sig}, msg.Memo))
//}
//
//func BuildAndSign(txBldr auth.TxBuilder, ctx context.CLIContext,
//	msgs []sdk.Msg, ixoDid exported.IxoDid) ([]byte, error) {
//	msg, err := txBldr.BuildSignMsg(msgs)
//	if err != nil {
//		return nil, err
//	}
//
//	return Sign(ctx, msg, ixoDid)
//}
//
//func SignAndBroadcastTxFromStdSignMsg(cliCtx context.CLIContext,
//	msg auth.StdSignMsg, ixoDid exported.IxoDid) (sdk.TxResponse, error) {
//
//	// sign the transaction
//	txBytes, err := Sign(cliCtx, msg, ixoDid)
//	if err != nil {
//		return sdk.TxResponse{}, err
//	}
//
//	// broadcast to a Tendermint node
//	res, err := cliCtx.BroadcastTx(txBytes)
//	if err != nil {
//		return sdk.TxResponse{}, err
//	}
//	return res, nil
//}
//
//func MakeSignature(signBytes []byte,
//	privateKey ed25519tm.PrivKeyEd25519) (auth.StdSignature, error) {
//	sig, err := privateKey.Sign(signBytes)
//	if err != nil {
//		return auth.StdSignature{}, err
//	}
//
//	return auth.StdSignature{
//		PubKey:    privateKey.PubKey(),
//		Signature: sig,
//	}, nil
//}

// Identical to DefaultSigVerificationGasConsumer, but with ed25519 allowed
//func IxoSigVerificationGasConsumer(
//	meter sdk.GasMeter, sig []byte, pubkey crypto.PubKey, params authtypes.Params,
//) error {
//	switch pubkey := pubkey.(type) {
//	case ed25519tm.PubKey:
//		meter.ConsumeGas(params.SigVerifyCostED25519, "ante verify: ed25519")
//		return nil
//
//	case secp256k1.PubKey:
//		meter.ConsumeGas(params.SigVerifyCostSecp256k1, "ante verify: secp256k1")
//		return nil
//
//	case multisig.PubKeyMultisigThreshold:
//		var multisignature multisig.Multisignature
//		codec.Cdc.MustUnmarshalBinaryBare(sig, &multisignature)
//
//		ante.ConsumeMultisignatureVerificationGas(meter, multisignature, pubkey, params)
//		return nil
//
//	default:
//		return sdkerrors.Wrapf(sdkerrors.ErrInvalidPubKey, "unrecognized public key type: %T", pubkey)
//	}
//}

// Identical to DefaultSigVerificationGasConsumer, but with ed25519 allowed
func IxoSigVerificationGasConsumer(
	meter sdk.GasMeter, sig signing.SignatureV2, params authtypes.Params,
) error {
	pubkey := sig.PubKey
	switch pubkey := pubkey.(type) {
	case *ed25519tm.PubKey:
		meter.ConsumeGas(params.SigVerifyCostED25519, "ante verify: ed25519")
		return sdkerrors.Wrap(sdkerrors.ErrInvalidPubKey, "ED25519 public keys are unsupported")

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

//func consumeSimSigGas(gasmeter sdk.GasMeter, pubkey crypto.PubKey, sig auth.StdSignature, params auth.Params) {
//	simSig := auth.StdSignature{PubKey: pubkey}
//	if len(sig.Signature) == 0 {
//		simSig.Signature = simEd25519Sig[:]
//	}
//
//	sigBz := ModuleCdc.MustMarshalBinaryLengthPrefixed(simSig)
//	cost := sdk.Gas(len(sigBz) + 6)
//
//	// If the pubkey is a multi-signature pubkey, then we estimate for the maximum
//	// number of signers.
//	if _, ok := pubkey.(multisig.PubKeyMultisigThreshold); ok {
//		cost *= params.TxSigLimit
//	}
//
//	gasmeter.ConsumeGas(params.TxSizeCostPerByte*cost, "txSize")
//}

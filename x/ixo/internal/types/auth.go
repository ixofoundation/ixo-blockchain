package types

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/input"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	authexported "github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
	"github.com/spf13/viper"
	"github.com/tendermint/ed25519"
	"github.com/tendermint/tendermint/crypto"
	ed25519tm "github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/crypto/multisig"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"os"
)

var (
	expectedMinGasPrices       = "0.025" + IxoNativeToken
	approximationGasAdjustment = float64(1.5)
	// TODO: parameterise (or remove) hard-coded gas prices and adjustments

	// simulation signature values used to estimate gas consumption
	simEd25519Pubkey ed25519tm.PubKeyEd25519
	simEd25519Sig    [ed25519.SignatureSize]byte
)

func init() {
	// This decodes a valid hex string into a ed25519Pubkey for use in transaction simulation
	bz, _ := hex.DecodeString("035AD6810A47F073553FF30D2FCC7E0D3B1C0B74B61A1AAA2582344037151E14")
	copy(simEd25519Pubkey[:], bz)
}

type PubKeyGetter func(ctx sdk.Context, msg IxoMsg) (crypto.PubKey, error)

func NewDefaultPubKeyGetter(didKeeper DidKeeper) PubKeyGetter {
	return func(ctx sdk.Context, msg IxoMsg) (pubKey crypto.PubKey, res error) {

		signerDidDoc, err := didKeeper.GetDidDoc(ctx, msg.GetSignerDid())
		if err != nil {
			return pubKey, err
		}

		var pubKeyRaw ed25519tm.PubKeyEd25519
		copy(pubKeyRaw[:], base58.Decode(signerDidDoc.GetPubKey()))
		return pubKeyRaw, nil
	}
}

func consumeSimSigGas(gasmeter sdk.GasMeter, pubkey crypto.PubKey, sig auth.StdSignature, params auth.Params) {
	simSig := auth.StdSignature{PubKey: pubkey}
	if len(sig.Signature) == 0 {
		simSig.Signature = simEd25519Sig[:]
	}

	sigBz := ModuleCdc.MustMarshalBinaryLengthPrefixed(simSig)
	cost := sdk.Gas(len(sigBz) + 6)

	// If the pubkey is a multi-signature pubkey, then we estimate for the maximum
	// number of signers.
	if _, ok := pubkey.(multisig.PubKeyMultisigThreshold); ok {
		cost *= params.TxSigLimit
	}

	gasmeter.ConsumeGas(params.TxSizeCostPerByte*cost, "txSize")
}

func ProcessSig(
	ctx sdk.Context, acc authexported.Account, sig auth.StdSignature, signBytes []byte, simulate bool, params auth.Params,
) (updatedAcc authexported.Account, res error) {

	var pubKey crypto.PubKey
	pubKey = acc.GetPubKey()
	err := acc.SetPubKey(pubKey)
	if err != nil {
		return nil, sdkerrors.Wrap(ErrInternal, "setting PubKey on signer's account")
	}

	if simulate {
		// Simulated txs should not contain a signature and are not required to
		// contain a pubkey, so we must account for tx size of including a
		// StdSignature (Amino encoding) and simulate gas consumption
		// (assuming an ED25519 simulation key).
		consumeSimSigGas(ctx.GasMeter(), pubKey, sig, params)
	}

	// Consume signature gas
	ctx.GasMeter().ConsumeGas(params.SigVerifyCostED25519, "ante verify: ed25519")

	// Verify signature
	if !simulate && !pubKey.VerifyBytes(signBytes, sig.Signature) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Signature Verification failed")
	}

	if err := acc.SetSequence(acc.GetSequence() + 1); err != nil {
		panic(err)
	}

	return acc, res
}

func getSignBytes(chainID string, tx auth.StdTx, acc authexported.Account, genesis bool) []byte {
	var accNum uint64
	if !genesis {
		accNum = acc.GetAccountNumber()
	}

	return auth.StdSignBytes(
		chainID, accNum, acc.GetSequence(), tx.Fee, tx.Msgs, tx.Memo,
	)
}

func NewDefaultAnteHandler(ak auth.AccountKeeper, sk supply.Keeper, pubKeyGetter PubKeyGetter) sdk.AnteHandler {
	return func(
		ctx sdk.Context, tx sdk.Tx, simulate bool,
	) (newCtx sdk.Context, res error) {

		if addr := sk.GetModuleAddress(auth.FeeCollectorName); addr == nil {
			panic(fmt.Sprintf("%s module account has not been set", auth.FeeCollectorName))
		}

		// all transactions must be of type auth.StdTx
		stdTx, ok := tx.(auth.StdTx)
		if !ok {
			// Set a gas meter with limit 0 as to prevent an infinite gas meter attack
			// during runTx.
			newCtx = auth.SetGasMeter(simulate, ctx, 0)
			return newCtx, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "tx must be auth.StdTx")
		}

		params := ak.GetParams(ctx)

		newCtx = auth.SetGasMeter(simulate, ctx, stdTx.Fee.Gas)

		// AnteHandlers must have their own defer/recover in order for the BaseApp
		// to know how much gas was used! This is because the GasMeter is created in
		// the AnteHandler, but if it panics the context won't be set properly in
		// runTx's recover call.
		defer func() {
			if r := recover(); r != nil {
				switch rType := r.(type) {
				//30.0 doesn't have these below mentioned funnctions ..need to check with miguel
				case sdk.ErrorOutOfGas:
					log := fmt.Sprintf(
						"out of gas in location: %v; gasWanted: %d, gasUsed: %d",
						rType.Descriptor, stdTx.Fee.Gas, newCtx.GasMeter().GasConsumed(),
					)
					res = sdkerrors.Wrap(sdkerrors.ErrOutOfGas, log)
				default:
					panic(r)
				}
			}
		}()

		if err := tx.ValidateBasic(); err != nil {
			return newCtx, err
		}

		// Number of messages in the tx must be 1
		if len(tx.GetMsgs()) != 1 {
			return ctx, sdkerrors.Wrap(ErrInternal, "number of messages must be 1")
		}

		newCtx.GasMeter().ConsumeGas(params.TxSizeCostPerByte*sdk.Gas(len(newCtx.TxBytes())), "txSize")

		// all messages must be of type IxoMsg
		msg, ok := stdTx.GetMsgs()[0].(IxoMsg)
		if !ok {
			return newCtx, sdkerrors.Wrap(ErrInternal, "msg must be ixo.IxoMsg")
		}

		// Get pubKey
		pubKey, res := pubKeyGetter(ctx, msg)
		if res != nil {
			return newCtx, res
		}

		// fetch first (and only) signer, who's going to pay the fees
		signerAddr := sdk.AccAddress(pubKey.Address())
		signerAcc, res := auth.GetSignerAcc(newCtx, ak, signerAddr)
		if res != nil {
			return newCtx, res
		}

		// deduct the fees
		if !stdTx.Fee.Amount.IsZero() {
			res = auth.DeductFees(sk, newCtx, signerAcc, stdTx.Fee.Amount)
			if res != nil {
				return newCtx, res
			}

			// reload the account as fees have been deducted
			signerAcc = ak.GetAccount(newCtx, signerAcc.GetAddress())
		}

		// check signature, return account with incremented nonce
		ixoSig := auth.StdSignature{PubKey: pubKey, Signature: stdTx.GetSignatures()[0]}
		isGenesis := ctx.BlockHeight() == 0
		signBytes := getSignBytes(newCtx.ChainID(), stdTx, signerAcc, isGenesis)
		signerAcc, res = ProcessSig(newCtx, signerAcc, ixoSig, signBytes, simulate, params)
		if res != nil {
			return newCtx, res
		}

		ak.SetAccount(newCtx, signerAcc)

		return newCtx, nil // continue...
	}
}

func ApproximateFeeForTx(cliCtx context.CLIContext, tx auth.StdTx, chainId string) (auth.StdFee, error) {

	// Set up a transaction builder
	cdc := cliCtx.Codec
	txEncoder := auth.DefaultTxEncoder
	gasAdjustment := approximationGasAdjustment
	fees := sdk.NewCoins(sdk.NewCoin(IxoNativeToken, sdk.OneInt()))
	txBldr := auth.NewTxBuilder(txEncoder(cdc), 0, 0, 0, gasAdjustment, true, chainId, tx.Memo, fees, nil)

	// Approximate gas consumption
	txBldr, err := utils.EnrichWithGas(txBldr, cliCtx, tx.Msgs)
	if err != nil {
		return auth.StdFee{}, err
	}

	// Clear fees and set gas-prices to deduce updated fee = (gas * gas-prices)
	signMsg, err := txBldr.WithFees("").WithGasPrices(expectedMinGasPrices).BuildSignMsg(tx.Msgs)
	if err != nil {
		return auth.StdFee{}, err
	}

	return signMsg.Fee, nil
}

func GenerateOrBroadcastMsgs(cliCtx context.CLIContext, msg sdk.Msg, ixoDid exported.IxoDid) error {
	msgs := []sdk.Msg{msg}
	txBldr := auth.NewTxBuilderFromCLI(cliCtx.Input).WithTxEncoder(utils.GetTxEncoder(cliCtx.Codec))

	if cliCtx.GenerateOnly {
		return utils.PrintUnsignedStdTx(txBldr, cliCtx, msgs)
	}

	return CompleteAndBroadcastTxCLI(txBldr, cliCtx, msgs, ixoDid)
}

func Sign(cliCtx context.CLIContext, msg auth.StdSignMsg,
	ixoDid exported.IxoDid) ([]byte, error) {
	if len(msg.Msgs) != 1 {
		panic("expected one message")
	}

	var privateKey ed25519tm.PrivKeyEd25519
	copy(privateKey[:], base58.Decode(ixoDid.Secret.SignKey))
	copy(privateKey[32:], base58.Decode(ixoDid.VerifyKey))

	sig, err := MakeSignature(msg.Bytes(), privateKey)
	if err != nil {
		return nil, err
	}

	encoder := utils.GetTxEncoder(cliCtx.Codec)
	return encoder(auth.NewStdTx(msg.Msgs, msg.Fee, []auth.StdSignature{sig}, msg.Memo))
}

func BuildAndSign(txBldr auth.TxBuilder, ctx context.CLIContext,
	msgs []sdk.Msg, ixoDid exported.IxoDid) ([]byte, error) {
	msg, err := txBldr.BuildSignMsg(msgs)
	if err != nil {
		return nil, err
	}

	return Sign(ctx, msg, ixoDid)
}

func CompleteAndBroadcastTxCLI(txBldr auth.TxBuilder, cliCtx context.CLIContext, msgs []sdk.Msg, ixoDid exported.IxoDid) error {
	txBldr, err := utils.PrepareTxBuilder(txBldr, cliCtx)
	if err != nil {
		return err
	}

	//fromName := cliCtx.GetFromName()

	if txBldr.SimulateAndExecute() || cliCtx.Simulate {
		txBldr, err = utils.EnrichWithGas(txBldr, cliCtx, msgs)
		if err != nil {
			return err
		}

		gasEst := utils.GasEstimateResponse{GasEstimate: txBldr.Gas()}
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", gasEst.String())
	}

	if cliCtx.Simulate {
		return nil
	}

	if !cliCtx.SkipConfirm {
		stdSignMsg, err := txBldr.BuildSignMsg(msgs)
		if err != nil {
			return err
		}

		var json []byte
		if viper.GetBool(flags.FlagIndentResponse) {
			json, err = cliCtx.Codec.MarshalJSONIndent(stdSignMsg, "", "  ")
			if err != nil {
				panic(err)
			}
		} else {
			json = cliCtx.Codec.MustMarshalJSON(stdSignMsg)
		}

		_, _ = fmt.Fprintf(os.Stderr, "%s\n\n", json)

		buf := bufio.NewReader(os.Stdin)
		ok, err := input.GetConfirmation("confirm transaction before signing and broadcasting", buf)
		if err != nil || !ok {
			_, _ = fmt.Fprintf(os.Stderr, "%s\n", "cancelled transaction")
			return err
		}
	}

	//passphrase, err := keys.GetPassphrase(fromName)
	//if err != nil {
	//	return err
	//}

	// build and sign the transaction
	txBytes, err := BuildAndSign(txBldr, cliCtx, msgs, ixoDid)
	if err != nil {
		return err
	}

	// broadcast to a Tendermint node
	res, err := cliCtx.BroadcastTx(txBytes)
	if err != nil {
		return err
	}

	return cliCtx.PrintOutput(res)
}

func SignAndBroadcastTxFromStdSignMsg(cliCtx context.CLIContext,
	msg auth.StdSignMsg, ixoDid exported.IxoDid) (sdk.TxResponse, error) {

	// sign the transaction
	txBytes, err := Sign(cliCtx, msg, ixoDid)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	// broadcast to a Tendermint node
	res, err := cliCtx.BroadcastTx(txBytes)
	if err != nil {
		return sdk.TxResponse{}, err
	}
	return res, nil
}

func MakeSignature(signBytes []byte,
	privateKey ed25519tm.PrivKeyEd25519) (auth.StdSignature, error) {
	sig, err := privateKey.Sign(signBytes)
	if err != nil {
		return auth.StdSignature{}, err
	}

	return auth.StdSignature{
		PubKey:    privateKey.PubKey(),
		Signature: sig,
	}, nil
}

// Identical to DefaultSigVerificationGasConsumer, but with ed25519 allowed
func IxoSigVerificationGasConsumer(
	meter sdk.GasMeter, sig []byte, pubkey crypto.PubKey, params auth.Params,
) error {
	switch pubkey := pubkey.(type) {
	case ed25519tm.PubKeyEd25519:
		meter.ConsumeGas(params.SigVerifyCostED25519, "ante verify: ed25519")
		return nil

	case secp256k1.PubKeySecp256k1:
		meter.ConsumeGas(params.SigVerifyCostSecp256k1, "ante verify: secp256k1")
		return nil

	case multisig.PubKeyMultisigThreshold:
		var multisignature multisig.Multisignature
		codec.Cdc.MustUnmarshalBinaryBare(sig, &multisignature)

		consumeMultisignatureVerificationGas(meter, multisignature, pubkey, params)
		return nil

	default:
		return sdkerrors.Wrap(sdkerrors.ErrInvalidPubKey, "unrecognized public key type")
	}
}

func consumeMultisignatureVerificationGas(meter sdk.GasMeter,
	sig multisig.Multisignature, pubkey multisig.PubKeyMultisigThreshold,
	params auth.Params) {

	size := sig.BitArray.Size()
	sigIndex := 0
	for i := 0; i < size; i++ {
		if sig.BitArray.GetIndex(i) {
			IxoSigVerificationGasConsumer(meter, sig.Sigs[sigIndex], pubkey.PubKeys[i], params)
			sigIndex++
		}
	}
}

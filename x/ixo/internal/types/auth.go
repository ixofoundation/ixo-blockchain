package types

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/input"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
	"github.com/spf13/viper"
	"github.com/tendermint/ed25519"
	"github.com/tendermint/tendermint/crypto"
	ed25519Keys "github.com/tendermint/tendermint/crypto/ed25519"
	"os"
	"time"
)

var (
	expectedMinGasPrices       = "0.025" + IxoNativeToken
	approximationGasAdjustment = float64(1.5)
	// TODO: parameterise (or remove) hard-coded gas prices and adjustments
)

type PubKeyGetter func(ctx sdk.Context, msg IxoMsg) (crypto.PubKey, sdk.Result)

func NewDefaultPubKeyGetter(didKeeper DidKeeper) PubKeyGetter {
	return func(ctx sdk.Context, msg IxoMsg) (pubKey crypto.PubKey, res sdk.Result) {

		signerDidDoc, err := didKeeper.GetDidDoc(ctx, msg.GetSignerDid())
		if err != nil {
			return pubKey, err.Result()
		}

		var pubKeyRaw [32]byte
		copy(pubKeyRaw[:], base58.Decode(signerDidDoc.GetPubKey()))
		return ed25519Keys.PubKeyEd25519(pubKeyRaw), sdk.Result{}
	}
}

// ProcessPubKey verifies that the given account address matches that of the
// IxoSignature. In addition, it will set the public key of the account if it
// has not been set.
func processPubKey(acc auth.Account, pubKey crypto.PubKey, simulate bool) (crypto.PubKey, sdk.Result) {
	// If pubkey is not known for account, set it from the StdSignature.
	accPubKey := acc.GetPubKey()
	if simulate {
		// In simulate mode the transaction comes with no signatures, thus if the
		// account's pubkey is nil, both signature verification and gasKVStore.Set()
		// shall consume enough gas to process an ed25519 pubkey
		if accPubKey == nil {
			return ed25519Keys.PubKeyEd25519{}, sdk.Result{}
		}

		return accPubKey, sdk.Result{}
	}

	if accPubKey == nil {
		accPubKey = pubKey
		if accPubKey == nil {
			return nil, sdk.ErrInvalidPubKey("PubKey not found").Result()
		}

		// TODO: uncomment as soon as this is true for pubkeys and accounts
		// Note: once the below is uncommented, this function can actually be
		//       removed and replaced by use of the same one in Cosmos SDK
		//if !bytes.Equal(accPubKey.Address(), acc.GetAddress()) {
		//	return nil, sdk.ErrInvalidPubKey(
		//		fmt.Sprintf("PubKey does not match Signer address %s", acc.GetAddress())).Result()
		//}
	}

	return accPubKey, sdk.Result{}
}

func ProcessSig(ctx sdk.Context, acc auth.Account, sig auth.StdSignature, signBytes []byte,
	simulate bool, params auth.Params) (updatedAcc auth.Account, res sdk.Result) {

	pubKey, res := processPubKey(acc, sig.PubKey, simulate)
	if !res.IsOK() {
		return nil, res
	}

	err := acc.SetPubKey(pubKey)
	if err != nil {
		return nil, sdk.ErrInternal("setting PubKey on signer's account").Result()
	}

	if simulate {
		// Simulated txs should not contain a signature and are not required to
		// contain a pubkey, so we must account for tx size of including an
		// IxoSignature and simulate gas consumption (assuming ED25519 key).
		//consumeSimSigGas(ctx.GasMeter(), sig, params)

		// NOTE: this is not the case in the ixo blockchain. The IxoSignature
		// will be blank but still count towards the transaction size given
		// that it uses a fixed length byte array [64]byte as the sig value.
	}

	// Consume signature gas
	ctx.GasMeter().ConsumeGas(params.SigVerifyCostED25519, "ante verify: ed25519")

	// Verify signature
	if !simulate && !pubKey.VerifyBytes(signBytes, sig.Signature) {
		return nil, sdk.ErrUnauthorized("Signature Verification failed").Result()
	}

	if err := acc.SetSequence(acc.GetSequence() + 1); err != nil {
		panic(err)
	}

	return acc, res
}

func getSignBytes(chainID string, ixoTx IxoTx, acc auth.Account, genesis bool) []byte {
	var accNum uint64
	if !genesis {
		accNum = acc.GetAccountNumber()
	}

	return auth.StdSignBytes(
		chainID, accNum, acc.GetSequence(), ixoTx.Fee, ixoTx.Msgs, ixoTx.Memo,
	)
}

func NewDefaultAnteHandler(ak auth.AccountKeeper, sk supply.Keeper, pubKeyGetter PubKeyGetter) sdk.AnteHandler {
	return func(
		ctx sdk.Context, tx sdk.Tx, simulate bool,
	) (newCtx sdk.Context, res sdk.Result, abort bool) {

		if addr := sk.GetModuleAddress(auth.FeeCollectorName); addr == nil {
			panic(fmt.Sprintf("%s module account has not been set", auth.FeeCollectorName))
		}

		// all transactions must be of type ixo.IxoTx
		ixoTx, ok := tx.(IxoTx)
		if !ok {
			// Set a gas meter with limit 0 as to prevent an infinite gas meter attack
			// during runTx.
			newCtx = auth.SetGasMeter(simulate, ctx, 0)
			return newCtx, sdk.ErrInternal("tx must be ixo.IxoTx").Result(), true
		}

		params := ak.GetParams(ctx)

		// Ensure that the provided fees meet a minimum threshold for the validator,
		// if this is a CheckTx. This is only for local mempool purposes, and thus
		// is only ran on check tx.
		if ctx.IsCheckTx() && !simulate {
			res := auth.EnsureSufficientMempoolFees(ctx, ixoTx.Fee)
			if !res.IsOK() {
				return newCtx, res, true
			}
		}

		newCtx = auth.SetGasMeter(simulate, ctx, ixoTx.Fee.Gas)

		// AnteHandlers must have their own defer/recover in order for the BaseApp
		// to know how much gas was used! This is because the GasMeter is created in
		// the AnteHandler, but if it panics the context won't be set properly in
		// runTx's recover call.
		defer func() {
			if r := recover(); r != nil {
				switch rType := r.(type) {
				case sdk.ErrorOutOfGas:
					log := fmt.Sprintf(
						"out of gas in location: %v; gasWanted: %d, gasUsed: %d",
						rType.Descriptor, ixoTx.Fee.Gas, newCtx.GasMeter().GasConsumed(),
					)
					res = sdk.ErrOutOfGas(log).Result()

					res.GasWanted = ixoTx.Fee.Gas
					res.GasUsed = newCtx.GasMeter().GasConsumed()
					abort = true
				default:
					panic(r)
				}
			}
		}()

		if err := tx.ValidateBasic(); err != nil {
			return newCtx, err.Result(), true
		}

		newCtx.GasMeter().ConsumeGas(params.TxSizeCostPerByte*sdk.Gas(len(newCtx.TxBytes())), "txSize")

		if res := auth.ValidateMemo(auth.StdTx{Memo: ixoTx.Memo}, params); !res.IsOK() {
			return newCtx, res, true
		}

		// fetch first (and only) signer, who's going to pay the fees
		signerAddr := ixoTx.GetSigner()
		signerAcc, res := auth.GetSignerAcc(newCtx, ak, signerAddr)
		if !res.IsOK() {
			return newCtx, res, true
		}

		// deduct the fees
		if !ixoTx.Fee.Amount.IsZero() {
			res = auth.DeductFees(sk, newCtx, signerAcc, ixoTx.Fee.Amount)
			if !res.IsOK() {
				return newCtx, res, true
			}

			// reload the account as fees have been deducted
			signerAcc = ak.GetAccount(newCtx, signerAcc.GetAddress())
		}

		// all messages must be of type IxoMsg
		msg, ok := ixoTx.GetMsgs()[0].(IxoMsg)
		if !ok {
			return newCtx, sdk.ErrInternal("msg must be ixo.IxoMsg").Result(), true
		}

		// Get pubKey
		pubKey, res := pubKeyGetter(ctx, msg)
		if !res.IsOK() {
			return newCtx, res, true
		}

		// check signature, return account with incremented nonce
		ixoSig := auth.StdSignature{PubKey: pubKey, Signature: ixoTx.GetSignatures()[0].SignatureValue[:]}
		isGenesis := ctx.BlockHeight() == 0
		signBytes := getSignBytes(newCtx.ChainID(), ixoTx, signerAcc, isGenesis)
		signerAcc, res = ProcessSig(newCtx, signerAcc, ixoSig, signBytes, simulate, params)
		if !res.IsOK() {
			return newCtx, res, true
		}

		ak.SetAccount(newCtx, signerAcc)

		return newCtx, sdk.Result{GasWanted: ixoTx.Fee.Gas}, false // continue...
	}
}

func signAndBroadcast(ctx context.CLIContext, msg auth.StdSignMsg,
	ixoDid exported.IxoDid) (sdk.TxResponse, error) {
	if len(msg.Msgs) != 1 {
		panic("expected one message")
	}

	privKey := [64]byte{}
	copy(privKey[:], base58.Decode(ixoDid.Secret.SignKey))
	copy(privKey[32:], base58.Decode(ixoDid.VerifyKey))

	signature := SignIxoMessage(msg.Bytes(), privKey)
	tx := NewIxoTxSingleMsg(msg.Msgs[0], msg.Fee, signature, msg.Memo)

	bz, err := ctx.Codec.MarshalJSON(tx)
	if err != nil {
		return sdk.TxResponse{}, fmt.Errorf("Could not marshall tx to binary. Error: %s", err.Error())
	}

	res, err := ctx.BroadcastTx(bz)
	if err != nil {
		return sdk.TxResponse{}, fmt.Errorf("Could not broadcast tx. Error: %s", err.Error())
	}

	return res, nil
}

func simulateMsgs(txBldr auth.TxBuilder, cliCtx context.CLIContext, msgs []sdk.Msg) (estimated, adjusted uint64, err error) {
	// Build the transaction
	stdSignMsg, err := txBldr.BuildSignMsg(msgs)
	if err != nil {
		return
	}

	// Signature set to a blank signature
	signature := IxoSignature{}
	signature.Created = signature.Created.Add(1) // maximizes signature length
	tx := NewIxoTxSingleMsg(
		stdSignMsg.Msgs[0], stdSignMsg.Fee, signature, stdSignMsg.Memo)

	bz, err := cliCtx.Codec.MarshalJSON(tx)
	if err != nil {
		err = fmt.Errorf("Could not marshall tx to binary. Error: %s", err.Error())
		return
	}

	estimated, adjusted, err = utils.CalculateGas(
		cliCtx.QueryWithData, cliCtx.Codec, bz, txBldr.GasAdjustment())
	return
}

func enrichWithGas(txBldr auth.TxBuilder, cliCtx context.CLIContext, msgs []sdk.Msg) (auth.TxBuilder, error) {
	_, adjusted, err := simulateMsgs(txBldr, cliCtx, msgs)
	if err != nil {
		return txBldr, err
	}

	return txBldr.WithGas(adjusted), nil
}

func ApproximateFeeForTx(cliCtx context.CLIContext, tx IxoTx, chainId string) (auth.StdFee, error) {

	// Set up a transaction builder
	cdc := cliCtx.Codec
	txEncoder := auth.DefaultTxEncoder
	gasAdjustment := approximationGasAdjustment
	fees := sdk.NewCoins(sdk.NewCoin(IxoNativeToken, sdk.OneInt()))
	txBldr := auth.NewTxBuilder(txEncoder(cdc), 0, 0, 0, gasAdjustment, true, chainId, tx.Memo, fees, nil)

	// Approximate gas consumption
	txBldr, err := enrichWithGas(txBldr, cliCtx, tx.Msgs)
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
	txBldr := auth.NewTxBuilderFromCLI()

	if cliCtx.GenerateOnly {
		return utils.PrintUnsignedStdTx(txBldr, cliCtx, msgs)
	}

	return CompleteAndBroadcastTxCLI(txBldr, cliCtx, msgs, ixoDid)
}

func CompleteAndBroadcastTxCLI(txBldr auth.TxBuilder, cliCtx context.CLIContext, msgs []sdk.Msg, ixoDid exported.IxoDid) error {
	txBldr, err := utils.PrepareTxBuilder(txBldr, cliCtx)
	if err != nil {
		return err
	}

	//fromName := cliCtx.GetFromName()

	if txBldr.SimulateAndExecute() || cliCtx.Simulate {
		txBldr, err = enrichWithGas(txBldr, cliCtx, msgs)
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

	// Build the transaction
	stdSignMsg, err := txBldr.BuildSignMsg(msgs)
	if err != nil {
		return err
	}

	// Sign and broadcast to a Tendermint node
	res, err := signAndBroadcast(cliCtx, stdSignMsg, ixoDid)
	if err != nil {
		return err
	}

	fmt.Println(res.String())
	fmt.Printf("Committed at block %d. Hash: %s\n", res.Height, res.TxHash)
	return nil
}

func CompleteAndBroadcastTxRest(cliCtx context.CLIContext, msg sdk.Msg, ixoDid exported.IxoDid) ([]byte, error) {

	// TODO: implement using txBldr or just remove function completely (ref: #123)

	// Construct dummy tx and approximate and set fee
	tx := NewIxoTxSingleMsg(msg, auth.StdFee{}, IxoSignature{}, "")
	chainId := viper.GetString(flags.FlagChainID)
	fee, err := ApproximateFeeForTx(cliCtx, tx, chainId)
	if err != nil {
		return nil, err
	}

	// Construct sign message
	stdSignMsg := auth.StdSignMsg{
		Fee:  fee,
		Msgs: []sdk.Msg{msg},
		Memo: "",
	}

	// Sign and broadcast to a Tendermint node
	res, err := signAndBroadcast(cliCtx, stdSignMsg, ixoDid)
	if err != nil {
		return nil, err
	}

	output, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return nil, err
	}
	return output, nil
}

func SignAndBroadcastTxFromStdSignMsg(cliCtx context.CLIContext,
	msg auth.StdSignMsg, ixoDid exported.IxoDid) (sdk.TxResponse, error) {
	return signAndBroadcast(cliCtx, msg, ixoDid)
}

func SignIxoMessage(signBytes []byte, privKey [64]byte) IxoSignature {

	signatureBytes := ed25519.Sign(&privKey, signBytes)
	signature := *signatureBytes

	return NewIxoSignature(time.Now(), signature)
}

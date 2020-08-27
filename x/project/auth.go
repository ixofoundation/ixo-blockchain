package project

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"github.com/ixofoundation/ixo-blockchain/x/project/internal/types"
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

func GetPubKeyGetter(keeper Keeper, didKeeper did.Keeper) ixo.PubKeyGetter {
	return func(ctx sdk.Context, msg ixo.IxoMsg) (pubKey crypto.PubKey, err error) {

		// Get signer PubKey
		var pubKeyEd25519 ed25519.PubKeyEd25519
		switch msg := msg.(type) {
		case MsgCreateProject:
			copy(pubKeyEd25519[:], base58.Decode(msg.GetPubKey()))
		case MsgWithdrawFunds:
			signerDid := msg.GetSignerDid()
			signerDoc, _ := didKeeper.GetDidDoc(ctx, signerDid)
			if signerDoc == nil {
				return pubKey, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "signer did not found")
			}
			copy(pubKeyEd25519[:], base58.Decode(signerDoc.GetPubKey()))
		default:
			// For the remaining messages, the project is the signer
			projectDoc, err := keeper.GetProjectDoc(ctx, msg.GetSignerDid())
			if err != nil {
				return pubKey, sdkerrors.Wrap(types.ErrInternal, "project did not found")
			}
			copy(pubKeyEd25519[:], base58.Decode(projectDoc.GetPubKey()))
		}
		return pubKeyEd25519, nil
	}
}

// Identical to Cosmos DeductFees function, but tokens sent to project account
func deductProjectFundingFees(bankKeeper bank.Keeper, ctx sdk.Context,
	acc exported.Account, projectAddr sdk.AccAddress, fees sdk.Coins) error {
	blockTime := ctx.BlockHeader().Time
	coins := acc.GetCoins()

	if !fees.IsValid() {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "invalid fee amount %s", fees)
	}

	// verify the account has enough funds to pay for fees
	_, hasNeg := coins.SafeSub(fees)
	if hasNeg {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "insufficient funds to pay for fees %s < %s", coins, fees)
	}

	// Validate the account has enough "spendable" coins as this will cover cases
	// such as vesting accounts.
	spendableCoins := acc.SpendableCoins(blockTime)
	if _, hasNeg := spendableCoins.SafeSub(fees); hasNeg {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "insufficient funds to pay for fees; %s < %s", spendableCoins, fees)
	}

	err := bankKeeper.SendCoins(ctx, acc.GetAddress(), projectAddr, fees)
	if err != nil {
		return err
	}

	return nil
}

func getProjectCreationSignBytes(chainID string, tx auth.StdTx, acc exported.Account, genesis bool) []byte {
	var accNum uint64
	if !genesis {
		// Fixed account number used so that sign bytes do not depend on it
		accNum = uint64(0)
	}

	return auth.StdSignBytes(
		chainID, accNum, acc.GetSequence(), tx.Fee, tx.Msgs, tx.Memo,
	)
}

// SetUpContextDecorator sets the GasMeter in the Context and wraps the next AnteHandler with a defer clause
// to recover from any downstream OutOfGas panics in the AnteHandler chain to return an error with information
// on gas provided and gas used.
// CONTRACT: Must be first decorator in the chain
// CONTRACT: Tx must implement GasTx interface
type SetUpContextDecorator struct{}

func NewSetUpContextDecorator() SetUpContextDecorator {
	return SetUpContextDecorator{}
}

func (sud SetUpContextDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	// all transactions must implement GasTx
	gasTx, ok := tx.(ante.GasTx)
	if !ok {
		// Set a gas meter with limit 0 as to prevent an infinite gas meter attack
		// during runTx.
		newCtx = ante.SetGasMeter(simulate, ctx, 0)
		return newCtx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be GasTx")
	}

	// Addding of DID uses an infinite gas meter
	newCtx = ctx.WithGasMeter(sdk.NewInfiniteGasMeter())

	// Decorator will catch an OutOfGasPanic caused in the next antehandler
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
					rType.Descriptor, gasTx.GetGas(), newCtx.GasMeter().GasConsumed())

				err = sdkerrors.Wrap(sdkerrors.ErrOutOfGas, log)
			default:
				panic(r)
			}
		}
	}()

	return next(newCtx, tx, simulate)
}

type SetPubKeyDecorator struct {
	ak  keeper.AccountKeeper
	pkg ixo.PubKeyGetter
}

func NewSetPubKeyDecorator(ak keeper.AccountKeeper, pkg ixo.PubKeyGetter) SetPubKeyDecorator {
	return SetPubKeyDecorator{
		ak:  ak,
		pkg: pkg,
	}
}

func (spkd SetPubKeyDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	_, ok := tx.(ante.SigVerifiableTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "invalid tx type")
	}

	// message must be of type MsgCreateProject
	msg, ok := tx.GetMsgs()[0].(MsgCreateProject)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "msg must be MsgCreateProject")
	}

	// Get project pubKey
	pubKey, err := spkd.pkg(ctx, msg)
	if err != nil {
		return ctx, err
	}

	// Fetch signer (project itself). Account expected to not exist
	signerAddr := sdk.AccAddress(pubKey.Address())
	_, err = auth.GetSignerAcc(ctx, spkd.ak, signerAddr)
	if err == nil {
		return ctx, fmt.Errorf("expected project account to not exist")
	}

	// Create signer's account
	signerAcc := spkd.ak.NewAccountWithAddress(ctx, signerAddr)
	spkd.ak.SetAccount(ctx, signerAcc)

	pubkeys := []crypto.PubKey{pubKey}
	signers := []sdk.AccAddress{signerAddr}

	for i, pk := range pubkeys {
		// PublicKey was omitted from slice since it has already been set in context
		if pk == nil {
			if !simulate {
				continue
			}
			pk = simEd25519Pubkey
		}
		// Only make check if simulate=false
		if !simulate && !bytes.Equal(pk.Address(), signers[i]) {
			return ctx, sdkerrors.Wrapf(sdkerrors.ErrInvalidPubKey,
				"pubKey does not match signer address %s with signer index: %d", signers[i], i)
		}

		acc, err := ante.GetSignerAcc(ctx, spkd.ak, signers[i])
		if err != nil {
			return ctx, err
		}
		// account already has pubkey set,no need to reset
		if acc.GetPubKey() != nil {
			continue
		}
		err = acc.SetPubKey(pk)
		if err != nil {
			return ctx, sdkerrors.Wrap(sdkerrors.ErrInvalidPubKey, err.Error())
		}
		spkd.ak.SetAccount(ctx, acc)
	}

	return next(ctx, tx, simulate)
}

// DeductFeeDecorator deducts fees from the first signer of the tx
// If the first signer does not have the funds to pay for the fees, return with InsufficientFunds error
// Call next AnteHandler if fees successfully deducted
// CONTRACT: Tx must implement FeeTx interface to use DeductFeeDecorator
type DeductFeeDecorator struct {
	ak           keeper.AccountKeeper
	supplyKeeper authtypes.SupplyKeeper
	bk           bank.Keeper
	didKeeper    did.Keeper
	pkg          ixo.PubKeyGetter
}

func NewDeductFeeDecorator(ak keeper.AccountKeeper, sk authtypes.SupplyKeeper,
	bk bank.Keeper, didKeeper did.Keeper, pkg ixo.PubKeyGetter) DeductFeeDecorator {
	return DeductFeeDecorator{
		ak:           ak,
		supplyKeeper: sk,
		bk:           bk,
		didKeeper:    didKeeper,
		pkg:          pkg,
	}
}

func (dfd DeductFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	_, ok := tx.(ante.FeeTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	if addr := dfd.supplyKeeper.GetModuleAddress(auth.FeeCollectorName); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", auth.FeeCollectorName))
	}

	// all messages must be of type MsgCreateProject
	msg, ok := tx.GetMsgs()[0].(MsgCreateProject)
	if !ok {
		return ctx, fmt.Errorf("msg must be MsgCreateProject")
	}

	// Get pubKey
	pubKey, err := dfd.pkg(ctx, msg)
	if err != nil {
		return ctx, err
	}

	// fetch first (and only) signer, who's going to pay the fees
	feePayer := sdk.AccAddress(pubKey.Address())
	feePayerAcc := dfd.ak.GetAccount(ctx, feePayer)

	if feePayerAcc == nil {
		return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "fee payer address: %s does not exist", feePayer)
	}

	// confirm that fee is the exact amount expected
	stdTx, ok := tx.(auth.StdTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "invalid transaction type")
	}
	expectedTotalFee := sdk.NewCoins(sdk.NewCoin(
		ixo.IxoNativeToken, sdk.NewInt(MsgCreateProjectTotalFee)))
	if !stdTx.Fee.Amount.IsEqual(expectedTotalFee) {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "invalid fee")
	}

	// Calculate transaction fee and project funding
	transactionFee := sdk.NewCoins(sdk.NewCoin(
		ixo.IxoNativeToken, sdk.NewInt(MsgCreateProjectTransactionFee)))
	projectFunding := expectedTotalFee.Sub(transactionFee) // panics if negative result

	// deduct the fees
	if !stdTx.Fee.Amount.IsZero() {
		// fetch fee payer account
		feePayerDidDoc, err := dfd.didKeeper.GetDidDoc(ctx, msg.SenderDid)
		if err != nil {
			return ctx, err
		}
		feePayerAcc, err := auth.GetSignerAcc(ctx, dfd.ak, feePayerDidDoc.Address())
		if err != nil {
			return ctx, err
		}

		err = auth.DeductFees(dfd.supplyKeeper, ctx, feePayerAcc, transactionFee)
		if err != nil {
			return ctx, err
		}

		projectAddr := sdk.AccAddress(pubKey.Address())
		err = deductProjectFundingFees(dfd.bk, ctx, feePayerAcc, projectAddr, projectFunding)
		if err != nil {
			return ctx, err
		}

		// reload the account as fees have been deducted
		feePayerAcc = dfd.ak.GetAccount(ctx, feePayerAcc.GetAddress())
	}

	return next(ctx, tx, simulate)
}

// Verify all signatures for a tx and return an error if any are invalid. Note,
// the SigVerificationDecorator decorator will not get executed on ReCheck.
//
// CONTRACT: Pubkeys are set in context for all signers before this decorator runs
// CONTRACT: Tx must implement SigVerifiableTx interface
type SigVerificationDecorator struct {
	ak  keeper.AccountKeeper
	pkg ixo.PubKeyGetter
}

func NewSigVerificationDecorator(ak keeper.AccountKeeper, pkg ixo.PubKeyGetter) SigVerificationDecorator {
	return SigVerificationDecorator{
		ak:  ak,
		pkg: pkg,
	}
}

func (svd SigVerificationDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	// no need to verify signatures on recheck tx
	if ctx.IsReCheckTx() {
		return next(ctx, tx, simulate)
	}
	sigTx, ok := tx.(ante.SigVerifiableTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "invalid transaction type")
	}

	// message must be of type MsgCreateProject
	msg, ok := tx.GetMsgs()[0].(MsgCreateProject)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "msg must be ixo.MsgCreateProject")
	}

	// Get did pubKey
	pubKey, err := svd.pkg(ctx, msg)
	if err != nil {
		return ctx, err
	}

	// Fetch signer (account underlying DID). Account expected to not exist
	signerAddr := sdk.AccAddress(pubKey.Address())
	signerAcc, err := ante.GetSignerAcc(ctx, svd.ak, signerAddr)
	if err != nil {
		return ctx, err
	}

	// check signature, return account with incremented nonce
	stdTx, ok := tx.(auth.StdTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "invalid transaction type")
	}
	ixoSig := auth.StdSignature{PubKey: pubKey, Signature: sigTx.GetSignatures()[0]}
	isGenesis := ctx.BlockHeight() == 0
	params := svd.ak.GetParams(ctx)
	signBytes := getProjectCreationSignBytes(ctx.ChainID(), stdTx, signerAcc, isGenesis)
	signerAcc, err = ixo.ProcessSig(ctx, signerAcc, ixoSig, signBytes, simulate, params)
	if err != nil {
		return ctx, err
	}

	svd.ak.SetAccount(ctx, signerAcc)

	return next(ctx, tx, simulate)
}

func NewProjectCreationAnteHandler(ak auth.AccountKeeper, supplyKeeper supply.Keeper,
	bk bank.Keeper, didKeeper did.Keeper, pubKeyGetter ixo.PubKeyGetter) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		//ante.NewMempoolFeeDecorator(),
		ante.NewValidateBasicDecorator(),
		ante.NewValidateMemoDecorator(ak),
		ante.NewConsumeGasForTxSizeDecorator(ak),
		NewSetPubKeyDecorator(ak, pubKeyGetter), // SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewValidateSigCountDecorator(ak),
		NewDeductFeeDecorator(ak, supplyKeeper, bk, didKeeper, pubKeyGetter),
		//ixo.NewSigGasConsumeDecorator(ak, sigGasConsumer, pubKeyGetter),
		NewSigVerificationDecorator(ak, pubKeyGetter),
		ixo.NewIncrementSequenceDecorator(ak, pubKeyGetter), // innermost AnteDecorator
	)
}

package project

import (
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/spf13/viper"

	"github.com/ixofoundation/ixo-blockchain/x/did"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
)

const (
	MsgCreateProjectFee            = int64(100000)
	MsgCreateProjectTransactionFee = int64(10000)
)

func GetPubKeyGetter(keeper Keeper, didKeeper did.Keeper) ixo.PubKeyGetter {
	return func(ctx sdk.Context, msg ixo.IxoMsg) ([32]byte, sdk.Result) {

		// Get signer PubKey
		var pubKey [32]byte
		switch msg := msg.(type) {
		case MsgCreateProject:
			copy(pubKey[:], base58.Decode(msg.GetPubKey()))
		case MsgUpdateProjectStatus:
			projectDoc, err := keeper.GetProjectDoc(ctx, msg.ProjectDid)
			if err != nil {
				return pubKey, sdk.ErrInternal("project did not found").Result()
			}
			copy(pubKey[:], base58.Decode(projectDoc.GetPubKey()))
		case MsgCreateAgent:
			projectDoc, err := keeper.GetProjectDoc(ctx, msg.ProjectDid)
			if err != nil {
				return pubKey, sdk.ErrInternal("project did not found").Result()
			}
			copy(pubKey[:], base58.Decode(projectDoc.GetPubKey()))
		case MsgUpdateAgent:
			projectDoc, err := keeper.GetProjectDoc(ctx, msg.ProjectDid)
			if err != nil {
				return pubKey, sdk.ErrInternal("project did not found").Result()
			}
			copy(pubKey[:], base58.Decode(projectDoc.GetPubKey()))
		case MsgCreateClaim:
			projectDoc, err := keeper.GetProjectDoc(ctx, msg.ProjectDid)
			if err != nil {
				return pubKey, sdk.ErrInternal("project did not found").Result()
			}
			copy(pubKey[:], base58.Decode(projectDoc.GetPubKey()))
		case MsgCreateEvaluation:
			projectDoc, err := keeper.GetProjectDoc(ctx, msg.ProjectDid)
			if err != nil {
				return pubKey, sdk.ErrInternal("project did not found").Result()
			}
			copy(pubKey[:], base58.Decode(projectDoc.GetPubKey()))
		case MsgWithdrawFunds:
			didDoc, _ := didKeeper.GetDidDoc(ctx, msg.Data.RecipientDid)
			if didDoc == nil {
				return pubKey, sdk.ErrUnauthorized("signer did not found").Result()
			}
			copy(pubKey[:], base58.Decode(didDoc.GetPubKey()))
		default:
			return pubKey, sdk.ErrUnknownRequest("No match for message type.").Result()
		}
		return pubKey, sdk.Result{}
	}
}

// Identical to Cosmos DeductFees function, but tokens sent to project account
func deductProjectFundingFees(bankKeeper bank.Keeper, ctx sdk.Context, acc auth.Account, projectDid ixo.Did, fees sdk.Coins) sdk.Result {
	blockTime := ctx.BlockHeader().Time
	coins := acc.GetCoins()

	if !fees.IsValid() {
		return sdk.ErrInsufficientFee(fmt.Sprintf("invalid fee amount: %s", fees)).Result()
	}

	// verify the account has enough funds to pay for fees
	_, hasNeg := coins.SafeSub(fees)
	if hasNeg {
		return sdk.ErrInsufficientFunds(
			fmt.Sprintf("insufficient funds to pay for fees; %s < %s", coins, fees),
		).Result()
	}

	// Validate the account has enough "spendable" coins as this will cover cases
	// such as vesting accounts.
	spendableCoins := acc.SpendableCoins(blockTime)
	if _, hasNeg := spendableCoins.SafeSub(fees); hasNeg {
		return sdk.ErrInsufficientFunds(
			fmt.Sprintf("insufficient funds to pay for fees; %s < %s", spendableCoins, fees),
		).Result()
	}

	projectAddr := ixo.DidToAddr(projectDid)
	err := bankKeeper.SendCoins(ctx, acc.GetAddress(), projectAddr, fees)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{}
}

func getProjectCreationSignBytes(chainID string, ixoTx ixo.IxoTx, acc auth.Account, genesis bool) []byte {
	var accNum uint64
	if !genesis {
		// Fixed account number used so that sign bytes do not depend on it
		accNum = uint64(0)
	}

	return auth.StdSignBytes(
		chainID, accNum, acc.GetSequence(), ixoTx.Fee, ixoTx.Msgs, ixoTx.Memo,
	)
}

func GetProjectCreationStdSignMsg(msgs []sdk.Msg) auth.StdSignMsg {
	chainID := viper.GetString(flags.FlagChainID)
	accNum, accSeq := uint64(0), uint64(0)
	fee := auth.NewStdFee(0, sdk.NewCoins(sdk.NewCoin(
		ixo.IxoNativeToken, sdk.NewInt(MsgCreateProjectFee))))
	memo := viper.GetString(flags.FlagMemo)

	return auth.StdSignMsg{
		ChainID:       chainID,
		AccountNumber: accNum,
		Sequence:      accSeq,
		Fee:           fee,
		Msgs:          msgs,
		Memo:          memo,
	}
}

func NewProjectCreationAnteHandler(ak auth.AccountKeeper, sk supply.Keeper,
	bk bank.Keeper, pubKeyGetter ixo.PubKeyGetter) sdk.AnteHandler {
	return func(
		ctx sdk.Context, tx sdk.Tx, simulate bool,
	) (newCtx sdk.Context, res sdk.Result, abort bool) {

		if addr := sk.GetModuleAddress(auth.FeeCollectorName); addr == nil {
			panic(fmt.Sprintf("%s module account has not been set", auth.FeeCollectorName))
		}

		// all transactions must be of type ixo.IxoTx
		ixoTx, ok := tx.(ixo.IxoTx)
		if !ok {
			// Set a gas meter with limit 0 as to prevent an infinite gas meter attack
			// during runTx.
			newCtx = auth.SetGasMeter(simulate, ctx, 0)
			return newCtx, sdk.ErrInternal("tx must be ixo.IxoTx").Result(), true
		}

		params := ak.GetParams(ctx)

		// Project creation uses an infinite gas meter
		newCtx = ctx.WithGasMeter(sdk.NewInfiniteGasMeter())

		if err := tx.ValidateBasic(); err != nil {
			return newCtx, err.Result(), true
		}

		newCtx.GasMeter().ConsumeGas(params.TxSizeCostPerByte*sdk.Gas(len(newCtx.TxBytes())), "txSize")

		if res := auth.ValidateMemo(auth.StdTx{Memo: ixoTx.Memo}, params); !res.IsOK() {
			return newCtx, res, true
		}

		// message must be of type MsgCreateProject
		msg, ok := ixoTx.GetMsgs()[0].(MsgCreateProject)
		if !ok {
			return newCtx, sdk.ErrInternal("msg must be MsgCreateProject").Result(), true
		}

		// Fetch signer (project itself). Account expected to not exist
		signerAddr := ixoTx.GetSigner()
		signerAcc, res := auth.GetSignerAcc(newCtx, ak, signerAddr)
		if res.IsOK() {
			return newCtx, sdk.ErrInternal("expected project account to not exist").Result(), true
		}

		// confirm that fee is the exact amount expected
		expectedTotalFee := sdk.NewCoins(sdk.NewCoin(
			ixo.IxoNativeToken, sdk.NewInt(MsgCreateProjectFee)))
		if !ixoTx.Fee.Amount.IsEqual(expectedTotalFee) {
			return newCtx, sdk.ErrInvalidCoins("invalid fee").Result(), true
		}

		// Calculate transaction fee and project funding
		transactionFee := sdk.NewCoins(sdk.NewCoin(ixo.IxoNativeToken, sdk.NewInt(MsgCreateProjectTransactionFee)))
		projectFunding := expectedTotalFee.Sub(transactionFee) // panics if negative result

		// deduct the fees
		if !ixoTx.Fee.Amount.IsZero() {
			// fetch fee payer
			feePayerAddr := ixo.DidToAddr(msg.SenderDid)
			feePayerAcc, res := auth.GetSignerAcc(ctx, ak, feePayerAddr)
			if !res.IsOK() {
				return newCtx, res, true
			}

			res = auth.DeductFees(sk, newCtx, feePayerAcc, transactionFee)
			if !res.IsOK() {
				return newCtx, res, true
			}

			res = deductProjectFundingFees(bk, newCtx, feePayerAcc, msg.ProjectDid, projectFunding)
			if !res.IsOK() {
				return newCtx, res, true
			}

			// reload the account as fees have been deducted
			feePayerAcc = ak.GetAccount(newCtx, feePayerAcc.GetAddress())
		}

		// Get pubKey
		pubKey, res := pubKeyGetter(ctx, msg)
		if !res.IsOK() {
			return newCtx, res, true
		}

		// Fetch signer account (project itself); create if it does not exist
		signerAcc, res = auth.GetSignerAcc(ctx, ak, signerAddr)
		if !res.IsOK() {
			signerAcc = ak.NewAccountWithAddress(ctx, signerAddr)
			ak.SetAccount(ctx, signerAcc)
		}

		// check signature, return account with incremented nonce
		ixoSig := ixoTx.GetSignatures()[0]
		isGenesis := ctx.BlockHeight() == 0
		signBytes := getProjectCreationSignBytes(newCtx.ChainID(), ixoTx, signerAcc, isGenesis)
		signerAcc, res = ixo.ProcessSig(newCtx, signerAcc, signBytes, pubKey, ixoSig, simulate, params)
		if !res.IsOK() {
			return newCtx, res, true
		}

		ak.SetAccount(newCtx, signerAcc)

		return newCtx, sdk.Result{GasWanted: ixoTx.Fee.Gas}, false // continue...
	}
}

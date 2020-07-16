package project

import (
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

func GetPubKeyGetter(keeper Keeper, didKeeper did.Keeper) ixo.PubKeyGetter {
	return func(ctx sdk.Context, msg ixo.IxoMsg) (pubKey crypto.PubKey, res sdk.Result) {

		// Get signer PubKey
		var pubKeyEd25519 ed25519.PubKeyEd25519
		switch msg := msg.(type) {
		case MsgCreateProject:
			copy(pubKeyEd25519[:], base58.Decode(msg.GetPubKey()))
		case MsgWithdrawFunds:
			signerDid := msg.GetSignerDid()
			signerDoc, _ := didKeeper.GetDidDoc(ctx, signerDid)
			if signerDoc == nil {
				return pubKey, sdk.ErrUnauthorized("signer did not found").Result()
			}
			copy(pubKeyEd25519[:], base58.Decode(signerDoc.GetPubKey()))
		default:
			// For the remaining messages, the project is the signer
			projectDoc, err := keeper.GetProjectDoc(ctx, msg.GetSignerDid())
			if err != nil {
				return pubKey, sdk.ErrInternal("project did not found").Result()
			}
			copy(pubKeyEd25519[:], base58.Decode(projectDoc.GetPubKey()))
		}
		return pubKeyEd25519, sdk.Result{}
	}
}

// Identical to Cosmos DeductFees function, but tokens sent to project account
func deductProjectFundingFees(bankKeeper bank.Keeper, ctx sdk.Context,
	acc auth.Account, projectAddr sdk.AccAddress, fees sdk.Coins) sdk.Result {
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

	err := bankKeeper.SendCoins(ctx, acc.GetAddress(), projectAddr, fees)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{}
}

func getProjectCreationSignBytes(chainID string, tx auth.StdTx, acc auth.Account, genesis bool) []byte {
	var accNum uint64
	if !genesis {
		// Fixed account number used so that sign bytes do not depend on it
		accNum = uint64(0)
	}

	return auth.StdSignBytes(
		chainID, accNum, acc.GetSequence(), tx.Fee, tx.Msgs, tx.Memo,
	)
}

func NewProjectCreationAnteHandler(ak auth.AccountKeeper, sk supply.Keeper,
	bk bank.Keeper, didKeeper did.Keeper,
	pubKeyGetter ixo.PubKeyGetter) sdk.AnteHandler {
	// TODO: ensure in all ante handlers not allowing multiple messages
	return func(
		ctx sdk.Context, tx sdk.Tx, simulate bool,
	) (newCtx sdk.Context, res sdk.Result, abort bool) {

		if addr := sk.GetModuleAddress(auth.FeeCollectorName); addr == nil {
			panic(fmt.Sprintf("%s module account has not been set", auth.FeeCollectorName))
		}

		// all transactions must be of type auth.StdTx
		stdTx, ok := tx.(auth.StdTx)
		if !ok {
			// Set a gas meter with limit 0 as to prevent an infinite gas meter attack
			// during runTx.
			newCtx = auth.SetGasMeter(simulate, ctx, 0)
			return newCtx, sdk.ErrInternal("tx must be auth.StdTx").Result(), true
		}

		params := ak.GetParams(ctx)

		// Project creation uses an infinite gas meter
		newCtx = ctx.WithGasMeter(sdk.NewInfiniteGasMeter())

		if err := tx.ValidateBasic(); err != nil {
			return newCtx, err.Result(), true
		}

		newCtx.GasMeter().ConsumeGas(params.TxSizeCostPerByte*sdk.Gas(len(newCtx.TxBytes())), "txSize")

		if res := auth.ValidateMemo(auth.StdTx{Memo: stdTx.Memo}, params); !res.IsOK() {
			return newCtx, res, true
		}

		// message must be of type MsgCreateProject
		msg, ok := stdTx.GetMsgs()[0].(MsgCreateProject)
		if !ok {
			return newCtx, sdk.ErrInternal("msg must be MsgCreateProject").Result(), true
		}

		// Get project pubKey
		projectPubKey, res := pubKeyGetter(ctx, msg)
		if !res.IsOK() {
			return newCtx, res, true
		}

		// Fetch signer (project itself). Account expected to not exist
		signerAddr := sdk.AccAddress(projectPubKey.Address())
		_, res = auth.GetSignerAcc(newCtx, ak, signerAddr)
		if res.IsOK() {
			return newCtx, sdk.ErrInternal("expected project account to not exist").Result(), true
		}

		// confirm that fee is the exact amount expected
		expectedTotalFee := sdk.NewCoins(sdk.NewCoin(
			ixo.IxoNativeToken, sdk.NewInt(MsgCreateProjectFee)))
		if !stdTx.Fee.Amount.IsEqual(expectedTotalFee) {
			return newCtx, sdk.ErrInvalidCoins("invalid fee").Result(), true
		}

		// Calculate transaction fee and project funding
		transactionFee := sdk.NewCoins(sdk.NewCoin(ixo.IxoNativeToken, sdk.NewInt(MsgCreateProjectTransactionFee)))
		projectFunding := expectedTotalFee.Sub(transactionFee) // panics if negative result

		// deduct the fees
		if !stdTx.Fee.Amount.IsZero() {
			// fetch fee payer account
			feePayerDidDoc, err := didKeeper.GetDidDoc(ctx, msg.SenderDid)
			if err != nil {
				return newCtx, err.Result(), true
			}
			feePayerAcc, res := auth.GetSignerAcc(ctx, ak, feePayerDidDoc.Address())
			if !res.IsOK() {
				return newCtx, res, true
			}

			res = auth.DeductFees(sk, newCtx, feePayerAcc, transactionFee)
			if !res.IsOK() {
				return newCtx, res, true
			}

			projectAddr := sdk.AccAddress(projectPubKey.Address())
			res = deductProjectFundingFees(bk, newCtx, feePayerAcc, projectAddr, projectFunding)
			if !res.IsOK() {
				return newCtx, res, true
			}

			// reload the account as fees have been deducted
			feePayerAcc = ak.GetAccount(newCtx, feePayerAcc.GetAddress())
		}

		// Fetch signer account (project itself); create if it does not exist
		signerAcc, res := auth.GetSignerAcc(ctx, ak, signerAddr)
		if !res.IsOK() {
			signerAcc = ak.NewAccountWithAddress(ctx, signerAddr)
			ak.SetAccount(ctx, signerAcc)
		}

		// check signature, return account with incremented nonce
		ixoSig := stdTx.GetSignatures()[0]
		isGenesis := ctx.BlockHeight() == 0
		signBytes := getProjectCreationSignBytes(newCtx.ChainID(), stdTx, signerAcc, isGenesis)
		signerAcc, res = ixo.ProcessSig(newCtx, signerAcc, ixoSig, signBytes, simulate, params)
		if !res.IsOK() {
			return newCtx, res, true
		}

		ak.SetAccount(newCtx, signerAcc)

		return newCtx, sdk.Result{GasWanted: stdTx.Fee.Gas}, false // continue...
	}
}

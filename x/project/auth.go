package project

import (
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"github.com/ixofoundation/ixo-blockchain/x/project/internal/types"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

func GetPubKeyGetter(keeper Keeper, didKeeper did.Keeper) ixo.PubKeyGetter {
	return func(ctx sdk.Context, msg ixo.IxoMsg) (pubKey crypto.PubKey, res error) {

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
		return sdkerrors.Wrap(sdkerrors.ErrInsufficientFee, "invalid fee amount")
	}

	// verify the account has enough funds to pay for fees
	_, hasNeg := coins.SafeSub(fees)
	if hasNeg {
		return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "insufficient funds to pay for fees")
	}

	// Validate the account has enough "spendable" coins as this will cover cases
	// such as vesting accounts.
	spendableCoins := acc.SpendableCoins(blockTime)
	if _, hasNeg := spendableCoins.SafeSub(fees); hasNeg {
		return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "insufficient funds to pay for fees")
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

func NewProjectCreationAnteHandler(ak auth.AccountKeeper, sk supply.Keeper,
	bk bank.Keeper, didKeeper did.Keeper,
	pubKeyGetter ixo.PubKeyGetter) sdk.AnteHandler {
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
			return newCtx, sdkerrors.Wrap(types.ErrInternal, "tx must be auth.StdTx")
		}

		params := ak.GetParams(ctx)

		// Project creation uses an infinite gas meter
		newCtx = ctx.WithGasMeter(sdk.NewInfiniteGasMeter())

		if err := tx.ValidateBasic(); err != nil {
			return newCtx, err
		}

		// Number of messages in the tx must be 1
		if len(tx.GetMsgs()) != 1 {
			return ctx, sdkerrors.Wrap(types.ErrInternal, "number of messages must be 1")

		}

		newCtx.GasMeter().ConsumeGas(params.TxSizeCostPerByte*sdk.Gas(len(newCtx.TxBytes())), "txSize")

		//if res := exported.ValidateMemo(auth.StdTx{Memo: stdTx.Memo}, params); res != nil {
		//	return newCtx, res
		//}

		// message must be of type MsgCreateProject
		msg, ok := stdTx.GetMsgs()[0].(MsgCreateProject)
		if !ok {
			return newCtx, sdkerrors.Wrap(types.ErrInternal, "msg must be MsgCreateProject")
		}

		// Get project pubKey
		projectPubKey, res := pubKeyGetter(ctx, msg)
		if res != nil {
			return newCtx, res
		}

		// Fetch signer (project itself). Account expected to not exist
		signerAddr := sdk.AccAddress(projectPubKey.Address())
		_, res = auth.GetSignerAcc(newCtx, ak, signerAddr)
		if res != nil {
			return newCtx, sdkerrors.Wrap(types.ErrInternal, "expected project account to not exist")

		}

		// confirm that fee is the exact amount expected
		expectedTotalFee := sdk.NewCoins(sdk.NewCoin(
			ixo.IxoNativeToken, sdk.NewInt(MsgCreateProjectTotalFee)))
		if !stdTx.Fee.Amount.IsEqual(expectedTotalFee) {
			return newCtx, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "invalid fee")

		}

		// Calculate transaction fee and project funding
		transactionFee := sdk.NewCoins(sdk.NewCoin(
			ixo.IxoNativeToken, sdk.NewInt(MsgCreateProjectTransactionFee)))
		projectFunding := expectedTotalFee.Sub(transactionFee) // panics if negative result

		// deduct the fees
		if !stdTx.Fee.Amount.IsZero() {
			// fetch fee payer account
			feePayerDidDoc, err := didKeeper.GetDidDoc(ctx, msg.SenderDid)
			if err != nil {
				return newCtx, err
			}
			feePayerAcc, res := auth.GetSignerAcc(ctx, ak, feePayerDidDoc.Address())
			if res != nil {
				return newCtx, res
			}

			res = auth.DeductFees(sk, newCtx, feePayerAcc, transactionFee)
			if res != nil {
				return newCtx, res
			}

			projectAddr := sdk.AccAddress(projectPubKey.Address())
			res = deductProjectFundingFees(bk, newCtx, feePayerAcc, projectAddr, projectFunding)
			if res != nil {
				return newCtx, res
			}

			// reload the account as fees have been deducted
			feePayerAcc = ak.GetAccount(newCtx, feePayerAcc.GetAddress())
		}

		// Fetch signer account (project itself); create if it does not exist
		signerAcc, res := auth.GetSignerAcc(ctx, ak, signerAddr)
		if res != nil {
			signerAcc = ak.NewAccountWithAddress(ctx, signerAddr)
			ak.SetAccount(ctx, signerAcc)
		}

		// check signature, return account with incremented nonce
		ixoSig := stdTx.GetSignatures()[0]
		isGenesis := ctx.BlockHeight() == 0
		signBytes := getProjectCreationSignBytes(newCtx.ChainID(), stdTx, signerAcc, isGenesis)
		signerAcc, res = ixo.ProcessSig(newCtx, signerAcc, ixoSig, signBytes, simulate, params)
		if res != nil {
			return newCtx, res
		}

		ak.SetAccount(newCtx, signerAcc)

		return newCtx, nil // continue...
	}
}

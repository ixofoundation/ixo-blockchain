package ante

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/ixofoundation/ixo-blockchain/x/project/types"
)

type IxoFeeHandlerDecorator struct {
	iidKeeper iidkeeper.Keeper
}

func NewIxoFeeHandlerDecorator(iidKeeper iidkeeper.Keeper, defaultFeeHandler ante.DeductFeeDecorator) IxoFeeHandlerDecorator {
	return IxoFeeHandlerDecorator{iidKeeper: iidKeeper}
}

func (dec IxoFeeHandlerDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	iidTx, ok := tx.(IidTx)
	if !ok {
		// return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a IIDTx")
		return next(ctx, tx, simulate)
	}

	if err := iidTx.VerifyIidControllersAgainstSigniture(ctx, dec.iidKeeper); err != nil {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a signedTx")
	}

	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	if addr := dfd.ak.GetModuleAddress(authtypes.FeeCollectorName); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", authtypes.FeeCollectorName))
	}

	// all messages must be of type MsgCreateProject
	msg, ok := tx.GetMsgs()[0].(*types.MsgCreateProject)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "msg must be MsgCreateProject")
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
	expectedTotalFee := sdk.NewCoins(sdk.NewCoin(
		ixotypes.IxoNativeToken, sdk.NewInt(types.MsgCreateProjectTotalFee)))
	if !feeTx.GetFee().IsEqual(expectedTotalFee) {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "invalid fee")
	}

	// Calculate transaction fee and project funding
	transactionFee := sdk.NewCoins(sdk.NewCoin(
		ixotypes.IxoNativeToken, sdk.NewInt(types.MsgCreateProjectTransactionFee)))
	projectFunding := expectedTotalFee.Sub(transactionFee) // panics if negative result

	// deduct the fees
	if !feeTx.GetFee().IsZero() {
		// fetch fee payer account
		feePayerDidDoc, err := dfd.didKeeper.GetDidDoc(ctx, msg.SenderDid)
		if err != nil {
			return ctx, err
		}
		feePayerAcc, err := ante.GetSignerAcc(ctx, dfd.ak, feePayerDidDoc.Address())
		if err != nil {
			return ctx, err
		}

		err = ante.DeductFees(dfd.bk, ctx, feePayerAcc, transactionFee)
		if err != nil {
			return ctx, err
		}

		projectAddr := sdk.AccAddress(pubKey.Address())
		err = deductProjectFundingFees(dfd.bk, ctx, feePayerAcc, projectAddr, projectFunding)
		if err != nil {
			return ctx, err
		}
	}

	return next(ctx, tx, simulate)
}

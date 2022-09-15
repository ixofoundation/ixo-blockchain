package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	iidkeeper "github.com/ixofoundation/ixo-blockchain/x/iid/keeper"
)

// MempoolFeeDecorator will check if the transaction's fee is at least as large
// as the local validator's minimum gasFee (defined in validator config).
// If fee is too low, decorator returns error and tx is rejected from mempool.
// Note this only applies when ctx.CheckTx = true
// If fee is high enough or not CheckTx, then call next AnteHandler
// CONTRACT: Tx must implement FeeTx to use MempoolFeeDecorator
type IidResolutionDecorator struct {
	iidKeeper iidkeeper.Keeper
}

func NewIidResolutionDecorator() IidResolutionDecorator {
	return IidResolutionDecorator{}
}

func (dec IidResolutionDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	iidTx, ok := tx.(IidTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a IIDTx")
	}

	if err := iidTx.VerifyIidControllersAgainstSigniture(ctx, dec.iidKeeper); err != nil {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a signedTx")
	}

	return next(ctx, tx, simulate)
}

// DeductFeeDecorator deducts fees from the first signer of the tx
// If the first signer does not have the funds to pay for the fees, return with InsufficientFunds error
// Call next AnteHandler if fees successfully deducted
// CONTRACT: Tx must implement FeeTx interface to use DeductFeeDecorator
// type DeductFeeDecorator struct {
// 	ak         AccountKeeper
// 	bankKeeper types.BankKeeper
// }

// func NewDeductFeeDecorator(ak AccountKeeper, bk types.BankKeeper) DeductFeeDecorator {
// 	return DeductFeeDecorator{
// 		ak:         ak,
// 		bankKeeper: bk,
// 	}
// }

// func (dfd DeductFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
// 	feeTx, ok := tx.(sdk.FeeTx)
// 	if !ok {
// 		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
// 	}

// 	if addr := dfd.ak.GetModuleAddress(types.FeeCollectorName); addr == nil {
// 		panic(fmt.Sprintf("%s module account has not been set", types.FeeCollectorName))
// 	}

// 	feePayer := feeTx.FeePayer()
// 	feePayerAcc := dfd.ak.GetAccount(ctx, feePayer)

// 	if feePayerAcc == nil {
// 		return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "fee payer address: %s does not exist", feePayer)
// 	}

// 	// deduct the fees
// 	if !feeTx.GetFee().IsZero() {
// 		err = DeductFees(dfd.bankKeeper, ctx, feePayerAcc, feeTx.GetFee())
// 		if err != nil {
// 			return ctx, err
// 		}
// 	}

// 	return next(ctx, tx, simulate)
// }

// // DeductFees deducts fees from the given account.
// func DeductFees(bankKeeper types.BankKeeper, ctx sdk.Context, acc types.AccountI, fees sdk.Coins) error {
// 	if !fees.IsValid() {
// 		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "invalid fee amount: %s", fees)
// 	}

// 	err := bankKeeper.SendCoinsFromAccountToModule(ctx, acc.GetAddress(), types.FeeCollectorName, fees)
// 	if err != nil {
// 		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, err.Error())
// 	}

// 	return nil
// }

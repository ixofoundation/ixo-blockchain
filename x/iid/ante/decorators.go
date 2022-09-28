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

func NewIidResolutionDecorator(iidKeeper iidkeeper.Keeper) IidResolutionDecorator {
	return IidResolutionDecorator{iidKeeper: iidKeeper}
}

func (dec IidResolutionDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	iidTx, ok := tx.(IidTx)
	if !ok {
		// return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a IIDTx")
		return next(ctx, tx, simulate)
	}

	if err := iidTx.VerifyIidControllersAgainstSigniture(ctx, dec.iidKeeper); err != nil {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a signedTx")
	}

	return next(ctx, tx, simulate)
}

type IidCapabilityVerificationDectorator struct {
	// iidKeeper iidkeeper.Keeper
}

func NewIidCapabilityVerificationDectorator() IidCapabilityVerificationDectorator {
	return IidCapabilityVerificationDectorator{}
}

func (dec IidCapabilityVerificationDectorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	_, ok := tx.(IidTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a IIDTx")
	}

	// if err := iidTx.VerifyIidControllersAgainstSigniture(ctx, dec.iidKeeper); err != nil {
	// 	return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a signedTx")
	// }

	return next(ctx, tx, simulate)
}

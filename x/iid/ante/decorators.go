package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	iidkeeper "github.com/ixofoundation/ixo-blockchain/x/iid/keeper"
)

type IidResolutionDecorator struct {
	iidKeeper iidkeeper.Keeper
}

func NewIidResolutionDecorator(iidKeeper iidkeeper.Keeper) IidResolutionDecorator {
	return IidResolutionDecorator{iidKeeper: iidKeeper}
}

func (dec IidResolutionDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	iidTx, ok := tx.(signing.SigVerifiableTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a IIDTx")
	}

	if err := VerifyIidControllersAgainstSigniture(iidTx, ctx, dec.iidKeeper); err != nil {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, err.Error())
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
	// _, ok := tx.(IidTx)
	// if !ok {
	// 	return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a IIDTx")
	// }

	// if err := iidTx.VerifyIidControllersAgainstSigniture(ctx, dec.iidKeeper); err != nil {
	// 	return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a signedTx")
	// }

	return next(ctx, tx, simulate)
}

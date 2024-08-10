package ante

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	iidkeeper "github.com/ixofoundation/ixo-blockchain/v3/x/iid/keeper"
)

type IidResolutionDecorator struct {
	iidKeeper iidkeeper.Keeper
}

func NewIidResolutionDecorator(iidKeeper iidkeeper.Keeper) IidResolutionDecorator {
	return IidResolutionDecorator{iidKeeper: iidKeeper}
}

func (dec IidResolutionDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	sigTx, ok := tx.(signing.SigVerifiableTx)
	if !ok {
		return ctx, errorsmod.Wrap(sdkerrors.ErrTxDecode, "invalid tx type")
	}

	if err := VerifyIidControllersAgainstSignature(sigTx, ctx, dec.iidKeeper); err != nil {
		return ctx, errorsmod.Wrap(sdkerrors.ErrUnauthorized, err.Error())
	}

	return next(ctx, tx, simulate)
}

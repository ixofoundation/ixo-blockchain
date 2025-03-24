package ante

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	iidkeeper "github.com/ixofoundation/ixo-blockchain/v4/x/iid/keeper"
)

type IidResolutionDecorator struct {
	iidKeeper iidkeeper.Keeper
	cdc       codec.Codec
}

func NewIidResolutionDecorator(iidKeeper iidkeeper.Keeper, cdc codec.Codec) IidResolutionDecorator {
	return IidResolutionDecorator{iidKeeper: iidKeeper, cdc: cdc}
}

func (dec IidResolutionDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	sigTx, ok := tx.(signing.SigVerifiableTx)
	if !ok {
		return ctx, errorsmod.Wrap(sdkerrors.ErrTxDecode, "invalid tx type")
	}

	if err := VerifyIidControllersAgainstSignature(sigTx, ctx, dec.iidKeeper, dec.cdc); err != nil {
		return ctx, errorsmod.Wrap(sdkerrors.ErrUnauthorized, err.Error())
	}

	return next(ctx, tx, simulate)
}

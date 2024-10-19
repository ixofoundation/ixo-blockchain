package keeper

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"

	"github.com/ixofoundation/ixo-blockchain/v4/x/token/types"
)

// NewTokenProposalHandler creates a new governance Handler for a ParamChangeProposal
func NewTokenProposalHandler(k Keeper) govtypesv1.Handler {
	return func(ctx sdk.Context, content govtypesv1.Content) error {
		switch c := content.(type) {
		case *types.SetTokenContractCodes:
			return k.handleTokenParameterChangeProposal(ctx, c)

		default:
			return errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized param proposal content type: %T", c)
		}
	}
}

func (k Keeper) handleTokenParameterChangeProposal(ctx sdk.Context, p *types.SetTokenContractCodes) error {
	var xx types.Params
	k.ParamSpace.GetParamSetIfExists(ctx, &xx)

	xx.Ixo1155ContractCode = p.Ixo1155ContractCode

	k.ParamSpace.SetParamSet(ctx, &xx)

	return nil
}

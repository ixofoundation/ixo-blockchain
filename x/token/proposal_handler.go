package token

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/ixofoundation/ixo-blockchain/x/token/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/token/types"
)

const (
	TokenNftContractName   = "token_nft"
	TokenNftContractSymbol = "token"
)

// NewParamChangeProposalHandler creates a new governance Handler for a ParamChangeProposal
func NewTokenParamChangeProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.SetTokenContractCodes:
			return handleTokenParameterChangeProposal(ctx, k, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized param proposal content type: %T", c)
		}
	}
}

func handleTokenParameterChangeProposal(ctx sdk.Context, k keeper.Keeper, p *types.SetTokenContractCodes) error {
	var xx types.Params
	k.ParamSpace.GetParamSetIfExists(ctx, &xx)

	xx.Cw20ContractCode = strconv.FormatUint(p.Cw20ContractCode, 10)
	xx.Cw721ContractCode = strconv.FormatUint(p.Cw721ContractCode, 10)
	xx.Ixo1155ContractCode = strconv.FormatUint(p.Ixo1155ContractCode, 10)

	k.ParamSpace.SetParamSet(ctx, &xx)

	return nil
}

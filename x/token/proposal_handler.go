package token

import (
	"fmt"

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
		case *types.InitializeTokenContract:
			return handleTokenParameterChangeProposal(ctx, k, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized param proposal content type: %T", c)
		}
	}
}

func handleTokenParameterChangeProposal(ctx sdk.Context, k keeper.Keeper, p *types.InitializeTokenContract) error {
	fmt.Printf("token propsal handeler =============\n%+v\n", *p)
	fmt.Println("Supspace", k.ParamSpace.Name(), k.ParamSpace.HasKeyTable())
	var xx types.Params
	k.ParamSpace.GetParamSetIfExists(ctx, &xx)
	fmt.Printf("%+v\n", xx)

	return nil
}

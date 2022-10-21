package entity

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/ixofoundation/ixo-blockchain/x/entity/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/entity/types"
)

// NewParamChangeProposalHandler creates a new governance Handler for a ParamChangeProposal
func NewEntityParamChangeProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.Params:
			return handleEntityParameterChangeProposal(ctx, k, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized param proposal content type: %T", c)
		}
	}
}

func handleEntityParameterChangeProposal(ctx sdk.Context, k keeper.Keeper, p *types.Params) error {
	fmt.Printf("propsal handeler =============\n%+v\n", *p)
	fmt.Println("Supspace", k.ParamSpace.Name(), k.ParamSpace.HasKeyTable())
	var xx types.Params
	k.ParamSpace.GetParamSetIfExists(ctx, &xx)
	fmt.Printf("%+v\n", xx)

	k.ParamSpace.SetParamSet(ctx, p)

	// for _, c := range p.Changes {
	// 	ss, ok := k.GetParams()
	// 	if !ok {
	// 		return sdkerrors.Wrap(proposal.ErrUnknownSubspace, c.Subspace)
	// 	}

	// 	// k.Logger(ctx).Info(
	// 	// 	fmt.Sprintf("attempt to set new parameter value; key: %s, value: %s", c.Key, c.Value),
	// 	// )

	// 	if err := ss.Update(ctx, []byte(c.Key), []byte(c.Value)); err != nil {
	// 		return sdkerrors.Wrapf(proposal.ErrSettingParameter, "key: %s, value: %s, err: %s", c.Key, c.Value, err.Error())
	// 	}
	// }

	return nil
}

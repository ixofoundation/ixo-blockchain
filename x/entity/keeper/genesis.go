package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/v4/x/entity/types"
)

// InitGenesis initializes the x/entity module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, gs types.GenesisState) {
	// Initialise params
	k.SetParams(ctx, &gs.Params)

	for _, entity := range gs.Entities {
		k.SetEntity(ctx, []byte(entity.Id), entity)
	}
}

// ExportGenesis returns the x/entity module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	params := k.GetParams(ctx)
	entities := k.GetAllEntity(ctx)

	return &types.GenesisState{
		Entities: entities,
		Params:   params,
	}
}

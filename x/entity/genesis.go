package entity

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/v3/x/entity/keeper"
	"github.com/ixofoundation/ixo-blockchain/v3/x/entity/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, gs *types.GenesisState) []abci.ValidatorUpdate {
	// Initialise params
	k.SetParams(ctx, &gs.Params)

	for _, entity := range gs.Entities {
		k.SetEntity(ctx, []byte(entity.Id), entity)
	}

	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	params := k.GetParams(ctx)
	entities := k.GetAllEntity(ctx)

	return &types.GenesisState{
		Entities: entities,
		Params:   params,
	}
}

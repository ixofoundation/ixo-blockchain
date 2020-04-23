package oracles

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis new oracles genesis
func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	keeper.SetOracles(ctx, data.Oracles)
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper Keeper) GenesisState {
	oracles := keeper.GetOracles(ctx)
	return NewGenesisState(oracles)
}

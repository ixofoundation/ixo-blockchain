package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/v5/x/token/types"
)

// InitGenesis initializes the x/token module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, gs types.GenesisState) {
	// save tokens to the store
	for _, t := range gs.Tokens {
		k.SetToken(ctx, t)
	}
	// save token properties to the store
	for _, tp := range gs.TokenProperties {
		k.SetTokenProperties(ctx, tp)
	}

	// Initialise params
	k.SetParams(ctx, &gs.Params)
}

// ExportGenesis returns the x/token module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	tokens := k.GetTokens(ctx)

	tokenProperties := k.GetTokenPropertiesAll(ctx)

	params := k.GetParams(ctx)

	return &types.GenesisState{
		Params:          params,
		Tokens:          tokens,
		TokenProperties: tokenProperties,
	}
}

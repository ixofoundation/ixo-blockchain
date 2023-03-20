package token

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/token/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/token/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func InitGenesis(ctx sdk.Context, k keeper.Keeper, gs *types.GenesisState) []abci.ValidatorUpdate {
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

	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	tokens := k.GetTokens(ctx)

	tokenProperties := k.GetTokenPropertiesAll(ctx)

	params := k.GetParams(ctx)

	return &types.GenesisState{
		Params:          params,
		Tokens:          tokens,
		TokenProperties: tokenProperties,
	}
}

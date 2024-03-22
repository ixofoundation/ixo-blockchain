package claims

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/v3/x/claims/keeper"
	"github.com/ixofoundation/ixo-blockchain/v3/x/claims/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, gs *types.GenesisState) []abci.ValidatorUpdate {
	// Initialise params
	k.SetParams(ctx, &gs.Params)

	// save collections to the store
	for _, c := range gs.Collections {
		k.SetCollection(ctx, c)
	}

	// save claims to the store
	for _, c := range gs.Claims {
		k.SetClaim(ctx, c)
	}

	// save disputes to the store
	for _, d := range gs.Disputes {
		k.SetDispute(ctx, d)
	}

	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	params := k.GetParams(ctx)

	collections := k.GetCollections(ctx)

	claims := k.GetClaims(ctx)

	disputes := k.GetDisputes(ctx)

	return &types.GenesisState{
		Params:      params,
		Collections: collections,
		Disputes:    disputes,
		Claims:      claims,
	}
}

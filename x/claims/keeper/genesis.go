package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/v4/x/claims/types"
)

// InitGenesis initializes the x/claims module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, gs types.GenesisState) {
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

	// save intents to the store
	for _, i := range gs.Intents {
		k.SetIntent(ctx, i)
	}
}

// ExportGenesis returns the x/claims module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	params := k.GetParams(ctx)

	collections := k.GetCollections(ctx)
	claims := k.GetClaims(ctx)
	disputes := k.GetDisputes(ctx)
	intents := k.GetIntents(ctx)

	return &types.GenesisState{
		Params:      params,
		Collections: collections,
		Disputes:    disputes,
		Claims:      claims,
		Intents:     intents,
	}
}

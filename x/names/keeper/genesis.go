package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-blockchain/v7/x/names/types"
)

// InitGenesis writes namespaces and name records to state. Genesis validation
// is performed by GenesisState.Validate before this is called.
func (k Keeper) InitGenesis(ctx sdk.Context, gs types.GenesisState) {
	for _, ns := range gs.Namespaces {
		k.SetNamespace(ctx, ns)
	}
	for _, r := range gs.Names {
		k.SetNameRecord(ctx, r)
	}
}

// ExportGenesis returns the full namespaces + names state.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return &types.GenesisState{
		Namespaces: k.GetAllNamespaces(ctx),
		Names:      k.GetAllNames(ctx),
	}
}

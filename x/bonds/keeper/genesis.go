package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/v3/x/bonds/types"
)

// InitGenesis initializes the x/bonds module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) {
	// Initialise bonds
	for _, b := range data.Bonds {
		k.SetBond(ctx, b.BondDid, b)
		k.SetBondDid(ctx, b.Token, b.BondDid)
	}

	// Initialise batches
	for _, b := range data.Batches {
		k.SetBatch(ctx, b.BondDid, b)
	}

	// Initialise params
	k.SetParams(ctx, data.Params)
}

// ExportGenesis returns the x/bonds module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	// Export bonds and batches
	var bonds []types.Bond
	var batches []types.Batch
	iterator := k.GetBondIterator(ctx)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		bond := k.MustGetBondByKey(ctx, iterator.Key())
		batch := k.MustGetBatch(ctx, bond.BondDid)
		bonds = append(bonds, bond)
		batches = append(batches, batch)
	}

	// Export params
	params := k.GetParams(ctx)

	return &types.GenesisState{
		Bonds:   bonds,
		Batches: batches,
		Params:  params,
	}
}

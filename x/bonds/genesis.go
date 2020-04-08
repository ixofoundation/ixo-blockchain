package bonds

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	// Initialise bonds
	for _, b := range data.Bonds {
		keeper.SetBond(ctx, b.BondDid, b)
		keeper.SetBondDid(ctx, b.Token, b.BondDid)
	}

	// Initialise batches
	for _, b := range data.Batches {
		keeper.SetBatch(ctx, b.BondDid, b)
	}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	// Export bonds and batches
	var bonds []Bond
	var batches []Batch

	iterator := k.GetBondIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		bond := k.MustGetBondByKey(ctx, iterator.Key())
		batch := k.MustGetBatch(ctx, bond.BondDid)

		bonds = append(bonds, bond)
		batches = append(batches, batch)
	}

	return GenesisState{
		Bonds:   bonds,
		Batches: batches,
	}
}

package bonds

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/types"
)

func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data types.GenesisState) {
	// Initialise bonds
	for _, b := range data.Bonds {
		keeper.SetBond(ctx, b.BondDid, b)
		keeper.SetBondDid(ctx, b.Token, b.BondDid)
	}

	// Initialise batches
	for _, b := range data.Batches {
		keeper.SetBatch(ctx, b.BondDid, b)
	}

	// Initialise params
	keeper.SetParams(ctx, data.Params)
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	// Export bonds and batches
	var bonds []types.Bond
	var batches []types.Batch
	iterator := k.GetBondIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		bond := k.MustGetBondByKey(ctx, iterator.Key())
		batch := k.MustGetBatch(ctx, bond.BondDid)
		bonds = append(bonds, bond)
		batches = append(batches, batch)
	}

	// Export params
	params := k.GetParams(ctx)

	return types.GenesisState{
		Bonds:   bonds,
		Batches: batches,
		Params:  params,
	}
}

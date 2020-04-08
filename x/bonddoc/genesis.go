package bonddoc

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	// Initialise bond docs
	for _, b := range data.BondDocs {
		keeper.SetBondDoc(ctx, &b)
	}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	// Export bond docs
	var bondDocs []MsgCreateBond

	iterator := k.GetBondDocIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		bondDoc := k.MustGetBondDocByKey(ctx, iterator.Key())

		bondDocs = append(bondDocs, *bondDoc.(*MsgCreateBond))
	}

	return GenesisState{
		BondDocs: bondDocs,
	}
}

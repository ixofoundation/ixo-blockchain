package fiat

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	for _, fiatPeg := range data.FiatPegWallet {
		keeper.SetFiatPeg(ctx, fiatPeg)
	}
}

func ExportGenesis(ctx sdk.Context, keeper Keeper) (data GenesisState) {
	fiatPegs := keeper.GetFiatPegs(ctx)

	return GenesisState{fiatPegs}
}

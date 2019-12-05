package fiat

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	for _, fiatAccount := range data.FiatAccount {
		keeper.SetFiatAccount(ctx, fiatAccount)
	}
}

func ExportGenesis(ctx sdk.Context, keeper Keeper) (data GenesisState) {
	fiatAccounts := keeper.GetFiatAccounts(ctx)

	return GenesisState{fiatAccounts}
}

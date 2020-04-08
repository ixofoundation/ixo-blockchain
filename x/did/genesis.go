package did

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	// Initialise did docs
	for _, d := range data.DidDocs {
		keeper.AddDidDoc(ctx, d)
	}

	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, keeper Keeper) (data GenesisState) {
	return GenesisState{
		DidDocs: keeper.GetAllDidDocs(ctx),
	}
}

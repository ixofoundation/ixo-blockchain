package did

import (
	"github.com/cosmos/cosmos-sdk/types"
	abciTypes "github.com/tendermint/tendermint/abci/types"
)

func InitGenesis(ctx types.Context, keeper Keeper, data GenesisState) []abciTypes.ValidatorUpdate {
	return []abciTypes.ValidatorUpdate{}
}

func ExportGenesis(ctx types.Context, keeper Keeper) (data GenesisState) {
	return GenesisState{}
}

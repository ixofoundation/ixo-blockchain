package claims

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/claims/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/claims/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState *types.GenesisState) []abci.ValidatorUpdate {
	k.SetParams(ctx, &genState.Params)

	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesisState()
	genesis.Params = k.GetParams(ctx)

	return genesis
}

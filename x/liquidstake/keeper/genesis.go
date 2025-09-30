package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/v6/x/liquidstake/types"
)

// InitGenesis initializes the liquidstake module's state from a given genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	if err := types.ValidateGenesis(genState); err != nil {
		panic(err)
	}

	// init to prevent nil slice, []types.WhitelistedValidator(nil)
	if genState.Params.WhitelistedValidators == nil || len(genState.Params.WhitelistedValidators) == 0 {
		genState.Params.WhitelistedValidators = []types.WhitelistedValidator{}
	}

	// Initialise params
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}

	// Save liquid validators to the store
	for _, lv := range genState.LiquidValidators {
		k.SetLiquidValidator(ctx, lv)
	}

	// Ensure that the module account exists
	moduleAcc := k.accountKeeper.GetModuleAccount(ctx, types.ModuleName)
	if moduleAcc == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}
}

// ExportGenesis returns the liquidstake module's genesis state.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	params := k.GetParams(ctx)

	// init to prevent nil slice, []types.WhitelistedValidator(nil)
	if params.WhitelistedValidators == nil || len(params.WhitelistedValidators) == 0 {
		params.WhitelistedValidators = []types.WhitelistedValidator{}
	}

	liquidValidators := k.GetAllLiquidValidators(ctx)
	return types.NewGenesisState(params, liquidValidators)
}

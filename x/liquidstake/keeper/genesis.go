package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-blockchain/v8/x/liquidstake/types"
)

// InitGenesis applies the multi-pool genesis state: writes ModuleParams,
// every Pool, and every per-pool LiquidValidator. Validation has already
// been performed by types.ValidateGenesis (called by the AppModule wrapper).
func (k Keeper) InitGenesis(ctx sdk.Context, genState types.GenesisState) {
	if err := types.ValidateGenesis(genState); err != nil {
		panic(err)
	}
	if err := k.SetModuleParams(ctx, genState.ModuleParams); err != nil {
		panic(err)
	}
	for _, p := range genState.Pools {
		k.SetPool(ctx, p)
	}
	for _, group := range genState.PoolLiquidValidators {
		for _, lv := range group.LiquidValidators {
			k.SetLiquidValidator(ctx, group.PoolId, lv)
		}
	}

	// Sanity check: the module account must exist (set up via maccPerms).
	if k.accountKeeper.GetModuleAccount(ctx, types.ModuleName) == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}
}

// ExportGenesis serialises the full multi-pool state. Per-pool validators
// are grouped under PoolLiquidValidators so the genesis file remains tidy.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	pools := k.GetAllPools(ctx)
	poolValidators := make([]types.PoolLiquidValidators, 0, len(pools))
	for _, p := range pools {
		lvs := k.GetAllLiquidValidatorsForPool(ctx, p.PoolId)
		// Skip empty groups to keep export concise; consumers may always
		// treat a missing entry as "no liquid validators".
		if len(lvs) == 0 {
			continue
		}
		poolValidators = append(poolValidators, types.PoolLiquidValidators{
			PoolId:           p.PoolId,
			LiquidValidators: lvs,
		})
	}
	return types.NewGenesisState(k.GetModuleParams(ctx), pools, poolValidators)
}

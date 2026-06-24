package liquidstake

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-blockchain/v8/x/liquidstake/keeper"
	"github.com/ixofoundation/ixo-blockchain/v8/x/liquidstake/types"
)

// BeginBlock refreshes every pool's LiquidValidator set so newly whitelisted
// validators become eligible for delegation. Heavy work (rebalancing,
// autocompounding) is gated to epoch hooks instead of running every block.
//
// Per-pool failures are logged inside UpdateLiquidValidatorSet; one bad
// pool must not stall the chain.
func BeginBlock(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	if k.IsModulePaused(ctx) {
		return
	}
	pools := k.GetAllPools(ctx)
	for _, p := range pools {
		if p.Paused {
			continue
		}
		_ = k.UpdateLiquidValidatorSet(ctx, p, false)
	}
}

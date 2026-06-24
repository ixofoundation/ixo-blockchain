package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	epochstypes "github.com/ixofoundation/ixo-blockchain/v8/x/epochs/types"
	"github.com/ixofoundation/ixo-blockchain/v8/x/liquidstake/types"
)

// AfterEpochEnd is intentionally a no-op for the liquidstake module; all
// per-epoch work runs in BeforeEpochStart so observers can read the
// post-rebalance state immediately after the epoch boundary.
func (k Keeper) AfterEpochEnd(ctx sdk.Context, epochIdentifier string, epochNumber int64) error {
	return nil
}

// BeforeEpochStart is the per-epoch entrypoint. On the autocompound epoch,
// every non-paused pool sweeps its rewards (fee + weighted distribution +
// re-stake). On the rebalance epoch, every non-paused pool refreshes its
// LiquidValidator set and rebalances toward target weights.
//
// Iteration is sequential by pool_id; one pool's failure does not abort
// processing of the others (failures are logged inside each helper).
func (k Keeper) BeforeEpochStart(ctx sdk.Context, epochIdentifier string, _ int64) error {
	if k.IsModulePaused(ctx) {
		return nil
	}

	// Snapshot pool list so a hook running mid-iteration cannot mutate the
	// store under us. Pool counts are bounded (single digits expected).
	//
	// AutocompoundEpoch and RebalanceEpoch are normally distinct identifiers
	// (production: "hour" / "day"). For local testing we collapse them both
	// to "2min", in which case both branches must fire on the same tick —
	// hence two separate ifs rather than a switch (Go forbids duplicate
	// case constants).
	pools := k.GetAllPools(ctx)
	for _, p := range pools {
		if p.Paused {
			continue
		}
		if epochIdentifier == types.AutocompoundEpoch {
			k.AutocompoundStakingRewards(ctx, p, p.WhitelistedValsMap())
		}
		if epochIdentifier == types.RebalanceEpoch {
			_ = k.UpdateLiquidValidatorSet(ctx, p, true)
		}
	}
	return nil
}

// ----------------------------------------------------------------------------
// EpochHooks adapter
// ----------------------------------------------------------------------------

type Hooks struct {
	k Keeper
}

var _ epochstypes.EpochHooks = Hooks{}

func (k Keeper) Hooks() Hooks { return Hooks{k} }

func (Hooks) GetModuleName() string { return types.ModuleName }

func (h Hooks) BeforeEpochStart(ctx sdk.Context, epochIdentifier string, epochNumber int64) error {
	return h.k.BeforeEpochStart(ctx, epochIdentifier, epochNumber)
}

func (h Hooks) AfterEpochEnd(ctx sdk.Context, epochIdentifier string, epochNumber int64) error {
	return h.k.AfterEpochEnd(ctx, epochIdentifier, epochNumber)
}

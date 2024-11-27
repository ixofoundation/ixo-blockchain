package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	epochstypes "github.com/ixofoundation/ixo-blockchain/v3/x/epochs/types"
	"github.com/ixofoundation/ixo-blockchain/v3/x/liquidstake/types"
)

// AfterEpochEnd is a hook which is executed after the end of an epoch. It is a no-op for liquidstake module.
func (k Keeper) AfterEpochEnd(ctx sdk.Context, epochIdentifier string, epochNumber int64) error {
	// no-op
	return nil
}

// BeforeEpochStart is a hook which is executed before the start of an epoch.
func (k Keeper) BeforeEpochStart(ctx sdk.Context, epochIdentifier string, _ int64) error {
	params := k.GetParams(ctx)

	if !params.ModulePaused {
		// Update the liquid validator set at the start of each epoch
		if epochIdentifier == types.AutocompoundEpoch {
			k.AutocompoundStakingRewards(ctx, types.GetWhitelistedValsMap(params.WhitelistedValidators))
		}

		if epochIdentifier == types.RebalanceEpoch {
			// return value of UpdateLiquidValidatorSet is useful only in testing
			_ = k.UpdateLiquidValidatorSet(ctx, true)
		}
	}

	return nil
}

// ___________________________________________________________________________________________________

// Hooks wrapper struct for liquidstake keeper.
type Hooks struct {
	k Keeper
}

var _ epochstypes.EpochHooks = Hooks{}

func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

// GetModuleName implements types.EpochHooks.
func (Hooks) GetModuleName() string {
	return types.ModuleName
}

func (h Hooks) BeforeEpochStart(ctx sdk.Context, epochIdentifier string, epochNumber int64) error {
	return h.k.BeforeEpochStart(ctx, epochIdentifier, epochNumber)
}

func (h Hooks) AfterEpochEnd(ctx sdk.Context, epochIdentifier string, epochNumber int64) error {
	return h.k.AfterEpochEnd(ctx, epochIdentifier, epochNumber)
}

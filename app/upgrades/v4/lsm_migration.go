package v4

import (
	"context"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// keeper contains the staking keeper functions required for the migration
type keeper interface {
	GetAllDelegations(ctx context.Context) ([]stakingtypes.Delegation, error)
	GetAllValidators(ctx context.Context) ([]stakingtypes.Validator, error)
	SetDelegation(ctx context.Context, delegation stakingtypes.Delegation) error
	SetValidator(ctx context.Context, validator stakingtypes.Validator) error
	RefreshTotalLiquidStaked(ctx context.Context) error
}

// Peforms migration for adding LSM support after v0.45.16-ics, since it was removed in cosmos-sdk-lsm v0.50:
//   - Setting each validator's ValidatorBondShares and LiquidShares to 0
//   - Setting each delegation's ValidatorBond field to false
//   - Calculating the total liquid staked by summing the delegations from ICA accounts
func migrateCosmosLSMModule(ctx sdk.Context, k keeper) error {

	ctx.Logger().Info("Staking LSM Migration: Migrating validators")
	MigrateValidators(ctx, k)

	ctx.Logger().Info("Staking LSM Migration: Migrating delegations")
	MigrateDelegations(ctx, k)

	// This is not needed
	// ctx.Logger().Info("Staking LSM Migration: Migrating UBD entries")
	// if err := MigrateUBDEntries(ctx, store, cdc); err != nil {
	// 	return err
	// }

	ctx.Logger().Info("Staking LSM Migration: Calculating total liquid staked")
	return k.RefreshTotalLiquidStaked(ctx)
}

// Set each validator's ValidatorBondShares and LiquidShares to 0
func MigrateValidators(ctx sdk.Context, k keeper) {
	validators, err := k.GetAllValidators(ctx)
	if err != nil {
		panic(err)
	}
	for _, validator := range validators {
		validator.ValidatorBondShares = math.LegacyZeroDec()
		validator.LiquidShares = math.LegacyZeroDec()
		k.SetValidator(ctx, validator)
	}
}

// Set each delegation's ValidatorBond field to false
func MigrateDelegations(ctx sdk.Context, k keeper) {
	delegations, err := k.GetAllDelegations(ctx)
	if err != nil {
		panic(err)
	}
	for _, delegation := range delegations {
		delegation.ValidatorBond = false
		k.SetDelegation(ctx, delegation)
	}
}

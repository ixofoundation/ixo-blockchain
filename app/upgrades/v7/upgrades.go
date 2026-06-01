package v7

import (
	"context"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/ixofoundation/ixo-blockchain/v7/app/keepers"
)

// CreateUpgradeHandler builds the v7 upgrade handler.
//
// Order of operations:
//
//  1. RunMigrations(fromVM) — gives every module a chance to apply its own
//     ConsensusVersion bump migrations. The liquidstake module's
//     ConsensusVersion is bumped from 1 to 2 in this release; we bypass the
//     module's auto-migration path because we want bespoke control over the
//     legacy → multi-pool reshape (see migrateLiquidStakeToMultiPool).
//  2. migrateLiquidStakeToMultiPool — reads legacy Params, writes
//     ModuleParams + the migrated "zero" Pool, re-keys all per-validator
//     records under the new per-pool prefix.
//
// We force fromVM[liquidstake] = 2 BEFORE RunMigrations so the module
// manager's automatic migrations do not touch the store; our explicit
// migration owns that responsibility.
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	appKeepers keepers.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(c context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		ctx := sdk.UnwrapSDKContext(c)
		ctx.Logger().Info("🚀 executing Ixo " + UpgradeName + " upgrade 🚀")

		// Take ownership of the liquidstake module's migration so the module
		// manager doesn't try to run an auto-migration with no registered
		// MigrationHandler for v1 → v2.
		const liquidStakeModule = "liquidstake"
		fromVM[liquidStakeModule] = 2

		ctx.Logger().Info("Run module migrations")
		newVM, err := mm.RunMigrations(c, configurator, fromVM)
		if err != nil {
			return nil, err
		}

		ctx.Logger().Info("Migrate liquidstake module to multi-pool layout")
		if err := migrateLiquidStakeToMultiPool(ctx, appKeepers.LiquidStakeKeeper); err != nil {
			return nil, err
		}

		return newVM, nil
	}
}

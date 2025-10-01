package v5

import (
	"context"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/ixofoundation/ixo-blockchain/v6/app/keepers"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	keepers keepers.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(context context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		ctx := sdk.UnwrapSDKContext(context)
		ctx.Logger().Info("🚀 executing Ixo " + UpgradeName + " upgrade 🚀")

		// -------------------------------------------------
		// Run migrations before applying any other state changes.
		// NOTE: DO NOT PUT ANY STATE CHANGES BEFORE RunMigrations().
		// -------------------------------------------------
		ctx.Logger().Info("Run migrations")
		migrations, err := mm.RunMigrations(context, configurator, fromVM)
		if err != nil {
			return nil, err
		}

		return migrations, nil
	}
}

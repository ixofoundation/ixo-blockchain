package v2

import (
	"context"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/ixofoundation/ixo-blockchain/v5/app/keepers"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	_ keepers.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(context context.Context, plan upgradetypes.Plan, _ module.VersionMap) (module.VersionMap, error) {
		ctx := sdk.UnwrapSDKContext(context)
		ctx.Logger().Info("🚀 executing Ixo v2 upgrade 🚀")

		// 1st-time running in-store migrations, setting fromVersion to
		// avoid running InitGenesis and migrations.
		fromVM := map[string]uint64{
			"auth":               2,
			"authz":              1,
			"bank":               2,
			"bonds":              1,
			"capability":         1,
			"claims":             1,
			"crisis":             1,
			"distribution":       2,
			"entity":             1,
			"evidence":           1,
			"feegrant":           1,
			"feeibc":             1,
			"genutil":            1,
			"gov":                2,
			"ibc":                2,
			"iid":                1,
			"interchainaccounts": 1,
			"intertx":            1,
			"mint":               1,
			"params":             1,
			"slashing":           2,
			"staking":            2,
			"token":              1,
			"transfer":           2,
			"upgrade":            1,
			"vesting":            1,
			"wasm":               2,
		}

		// Run migrations before applying any other state changes.
		migrations, err := mm.RunMigrations(ctx, configurator, fromVM)
		return migrations, err
	}
}

package v2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, _ module.VersionMap) (module.VersionMap, error) {
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

		// ctx.Logger().Info("set ICQKeeper params")
		// setICQParams(ctx, keepers.ICQKeeper)

		// ctx.Logger().Info("update ICAHostKeeper params to allow all messages")
		// setICAHostParams(ctx, &keepers.ICAHostKeeper)

		return migrations, err
	}
}

// TODO pass keepers through with cosmos sdk upgrade and implement this, not needed for now
// func setICQParams(ctx sdk.Context, icqKeeper *icqkeeper.Keeper) {
// 	icqparams := icqtypes.DefaultParams()
// 	icqparams.AllowQueries = wasmbinding.GetStargateWhitelistedPaths()
// 	// Adding SmartContractState query to allowlist
// 	icqparams.AllowQueries = append(icqparams.AllowQueries, "/cosmwasm.wasm.v1.Query/SmartContractState")
// 	icqKeeper.SetParams(ctx, icqparams)
// }

// func setICAHostParams(ctx sdk.Context, icahostkeeper *icahostkeeper.Keeper) {
// 	icahostkeeper.SetParams(ctx, icahosttypes.Params{
// 		HostEnabled:   true,
// 		AllowMessages: []string{"*"},
// 	})
// }

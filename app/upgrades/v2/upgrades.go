package v2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	icqkeeper "github.com/cosmos/ibc-apps/modules/async-icq/v4/keeper"
	icqtypes "github.com/cosmos/ibc-apps/modules/async-icq/v4/types"
	icahostkeeper "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/host/keeper"
	icahosttypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/host/types"
	"github.com/ixofoundation/ixo-blockchain/app/keepers"
	"github.com/ixofoundation/ixo-blockchain/wasmbinding"
	packetforwardtypes "github.com/strangelove-ventures/packet-forward-middleware/v4/router/types"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	keepers *keepers.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		ctx.Logger().Info("🚀 executing Ixo v2 upgrade 🚀")

		// Run migrations before applying any other state changes.
		migrations, err := mm.RunMigrations(ctx, configurator, fromVM)

		ctx.Logger().Info("set PacketForwardKeeper params")
		keepers.PacketForwardKeeper.SetParams(ctx, packetforwardtypes.DefaultParams())

		ctx.Logger().Info("set ICQKeeper params")
		setICQParams(ctx, keepers.ICQKeeper)

		ctx.Logger().Info("set ICAHostKeeper params to allow all messages")
		setICAHostParams(ctx, &keepers.ICAHostKeeper)

		return migrations, err
	}
}

func setICQParams(ctx sdk.Context, icqKeeper *icqkeeper.Keeper) {
	icqparams := icqtypes.DefaultParams()
	icqparams.AllowQueries = wasmbinding.GetStargateWhitelistedPaths()
	// Adding SmartContractState query to allowlist
	icqparams.AllowQueries = append(icqparams.AllowQueries, "/cosmwasm.wasm.v1.Query/SmartContractState")
	icqKeeper.SetParams(ctx, icqparams)
}

func setICAHostParams(ctx sdk.Context, icahostkeeper *icahostkeeper.Keeper) {
	icahostkeeper.SetParams(ctx, icahosttypes.Params{
		HostEnabled:   true,
		AllowMessages: []string{"*"},
	})
}

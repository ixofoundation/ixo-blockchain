package v2

import (
	"fmt"

	"github.com/CosmWasm/wasmd/x/wasm"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	icqkeeper "github.com/cosmos/ibc-apps/modules/async-icq/v4/keeper"
	icqtypes "github.com/cosmos/ibc-apps/modules/async-icq/v4/types"
	icacontrollertypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/controller/types"
	icahostkeeper "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/host/keeper"
	icahosttypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/host/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v4/modules/apps/29-fee/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	ibchost "github.com/cosmos/ibc-go/v4/modules/core/24-host"
	intertxtypes "github.com/cosmos/interchain-accounts/x/inter-tx/types"
	"github.com/ixofoundation/ixo-blockchain/app/keepers"
	"github.com/ixofoundation/ixo-blockchain/wasmbinding"
	bondstypes "github.com/ixofoundation/ixo-blockchain/x/bonds/types"
	claimsmoduletypes "github.com/ixofoundation/ixo-blockchain/x/claims/types"
	entitytypes "github.com/ixofoundation/ixo-blockchain/x/entity/types"
	iidtypes "github.com/ixofoundation/ixo-blockchain/x/iid/types"
	tokentypes "github.com/ixofoundation/ixo-blockchain/x/token/types"
	packetforwardtypes "github.com/strangelove-ventures/packet-forward-middleware/v4/router/types"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	keepers *keepers.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		ctx.Logger().Info("🚀 executing Ixo v2 upgrade 🚀")

		// 1st-time running in-store migrations, using 1 as fromVersion to
		// avoid running InitGenesis.
		fromVM = map[string]uint64{
			authtypes.StoreKey:          1,
			banktypes.StoreKey:          1,
			stakingtypes.StoreKey:       1,
			minttypes.StoreKey:          1,
			distrtypes.StoreKey:         1,
			slashingtypes.StoreKey:      1,
			govtypes.StoreKey:           1,
			paramstypes.StoreKey:        1,
			ibchost.StoreKey:            1,
			upgradetypes.StoreKey:       1,
			evidencetypes.StoreKey:      1,
			ibctransfertypes.StoreKey:   1,
			capabilitytypes.StoreKey:    1,
			authzkeeper.StoreKey:        1,
			feegrant.StoreKey:           1,
			wasm.StoreKey:               1,
			icahosttypes.StoreKey:       1,
			icacontrollertypes.StoreKey: 1,
			intertxtypes.StoreKey:       1,
			ibcfeetypes.StoreKey:        1,
			iidtypes.StoreKey:           1,
			bondstypes.StoreKey:         1,
			entitytypes.StoreKey:        1,
			tokentypes.StoreKey:         1,
			claimsmoduletypes.StoreKey:  1,
		}

		// Run migrations before applying any other state changes.
		ctx.Logger().Info(fmt.Sprintf("pre migrate version map: %v", fromVM))
		migrations, err := mm.RunMigrations(ctx, configurator, fromVM)
		ctx.Logger().Info(fmt.Sprintf("post migrate version map: %v", migrations))

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

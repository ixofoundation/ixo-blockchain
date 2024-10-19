package v4

import (
	"context"

	wasmv2 "github.com/CosmWasm/wasmd/x/wasm/migrations/v2"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"

	cmttypes "github.com/cometbft/cometbft/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	packetforwardtypes "github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v8/packetforward/types"
	icqtypes "github.com/cosmos/ibc-apps/modules/async-icq/v8/types"
	icacontrollertypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/types"
	icahosttypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	ibcconnectiontypes "github.com/cosmos/ibc-go/v8/modules/core/03-connection/types"
	ibcexported "github.com/cosmos/ibc-go/v8/modules/core/exported"

	// Local
	"github.com/ixofoundation/ixo-blockchain/v4/app/keepers"
	"github.com/ixofoundation/ixo-blockchain/v4/lib/ixo"
	"github.com/ixofoundation/ixo-blockchain/v4/wasmbinding"
	bondstypes "github.com/ixofoundation/ixo-blockchain/v4/x/bonds/types"
	claimsmoduletypes "github.com/ixofoundation/ixo-blockchain/v4/x/claims/types"
	entitytypes "github.com/ixofoundation/ixo-blockchain/v4/x/entity/types"
	smartaccounttypes "github.com/ixofoundation/ixo-blockchain/v4/x/smart-account/types"
	tokentypes "github.com/ixofoundation/ixo-blockchain/v4/x/token/types"

	// SDK v47 modules
	"cosmossdk.io/math"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	keepers keepers.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(context context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		ctx := sdk.UnwrapSDKContext(context)
		ctx.Logger().Info("🚀 executing Ixo " + UpgradeName + "upgrade 🚀")

		// -------------------------------------------------
		// Migrate Params
		// -------------------------------------------------
		ctx.Logger().Info("Migrate params")
		baseAppLegacySS := keepers.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramstypes.ConsensusParamsKeyTable())
		// https://github.com/cosmos/cosmos-sdk/pull/12363/files
		// Set param key table for params module migration
		for _, subspace := range keepers.ParamsKeeper.GetSubspaces() {
			subspace := subspace

			var keyTable paramstypes.KeyTable
			switch subspace.Name() {
			// sdk
			case authtypes.ModuleName:
				keyTable = authtypes.ParamKeyTable() //nolint:staticcheck
			case banktypes.ModuleName:
				keyTable = banktypes.ParamKeyTable() //nolint:staticcheck
			case stakingtypes.ModuleName:
				keyTable = stakingtypes.ParamKeyTable() //nolint:staticcheck
			case minttypes.ModuleName:
				keyTable = minttypes.ParamKeyTable() //nolint:staticcheck
			case distrtypes.ModuleName:
				keyTable = distrtypes.ParamKeyTable() //nolint:staticcheck
			case slashingtypes.ModuleName:
				keyTable = slashingtypes.ParamKeyTable() //nolint:staticcheck
			case govtypes.ModuleName:
				keyTable = govv1.ParamKeyTable() //nolint:staticcheck
			case crisistypes.ModuleName:
				keyTable = crisistypes.ParamKeyTable() //nolint:staticcheck

			// ibc types
			case ibcexported.ModuleName:
				keyTable = ibcclienttypes.ParamKeyTable()
				keyTable.RegisterParamSet(&ibcconnectiontypes.Params{})
			case ibctransfertypes.ModuleName:
				keyTable = ibctransfertypes.ParamKeyTable() //nolint:staticcheck
			case icahosttypes.SubModuleName:
				keyTable = icahosttypes.ParamKeyTable() //nolint:staticcheck
			case icacontrollertypes.SubModuleName:
				keyTable = icacontrollertypes.ParamKeyTable() //nolint:staticcheck
			case icqtypes.ModuleName:
				keyTable = icqtypes.ParamKeyTable() //nolint:staticcheck
			case packetforwardtypes.ModuleName:
				keyTable = packetforwardtypes.ParamKeyTable() //nolint:staticcheck
			// case ibchookstypes.ModuleName:
			// 	keyTable = ibchookstypes.ParamKeyTable() //nolint:staticcheck

			// wasm
			case wasmtypes.ModuleName:
				keyTable = wasmv2.ParamKeyTable() //nolint:staticcheck

			// ixo modules
			case bondstypes.ModuleName:
				keyTable = bondstypes.ParamKeyTable() //nolint:staticcheck
			case claimsmoduletypes.ModuleName:
				keyTable = claimsmoduletypes.ParamKeyTable() //nolint:staticcheck
			case entitytypes.ModuleName:
				keyTable = entitytypes.ParamKeyTable() //nolint:staticcheck
			// epochs doesn't have params
			// iidtypes doesn't have params
			case tokentypes.ModuleName:
				keyTable = tokentypes.ParamKeyTable() //nolint:staticcheck
			case smartaccounttypes.ModuleName:
				keyTable = smartaccounttypes.ParamKeyTable() //nolint:staticcheck

			default:
				continue
			}

			if !subspace.HasKeyTable() {
				subspace.WithKeyTable(keyTable)
			}
		}

		// Migrate Tendermint consensus parameters from x/params module to a deprecated x/consensus module.
		// The old params module is required to still be imported in your app.go in order to handle this migration.
		err := baseapp.MigrateParams(ctx, baseAppLegacySS, keepers.ConsensusParamsKeeper.ParamsStore)
		if err != nil {
			return nil, err
		}

		// Remove "mint" from fromVM since have custom mint module and want to run init genesis for it
		delete(fromVM, minttypes.ModuleName)

		// -------------------------------------------------
		// Run migrations before applying any other state changes.
		// NOTE: DO NOT PUT ANY STATE CHANGES BEFORE RunMigrations().
		// -------------------------------------------------
		ctx.Logger().Info("Run migrations")
		migrations, err := mm.RunMigrations(ctx, configurator, fromVM)
		if err != nil {
			return nil, err
		}

		// -------------------------------------------------
		// Set proposal param:
		// -------------------------------------------------
		ctx.Logger().Info("Set expedited proposal params")
		govParams, err := keepers.GovKeeper.Params.Get(ctx)
		if err != nil {
			return nil, err
		}
		// normal proposal deposit is 10k ixo, make expedited proposal deposit 3x
		govParams.ExpeditedMinDeposit = sdk.NewCoins(sdk.NewCoin(ixo.IxoNativeToken, math.NewInt(ExpeditedProposalDeposit)))
		govParams.MinInitialDepositRatio = MinInitialDepositRatio
		err = keepers.GovKeeper.Params.Set(ctx, govParams)
		if err != nil {
			return nil, err
		}

		// -------------------------------------------------
		// Set consensus params:
		// -------------------------------------------------
		ctx.Logger().Info("Set consensus params")
		defaultConsensusParams := cmttypes.DefaultConsensusParams().ToProto()
		defaultConsensusParams.Block.MaxBytes = BlockMaxBytes // previously 5000000
		defaultConsensusParams.Block.MaxGas = BlockMaxGas     // unchanged
		err = keepers.ConsensusParamsKeeper.ParamsStore.Set(ctx, defaultConsensusParams)
		if err != nil {
			return nil, err
		}

		// -------------------------------------------------
		// Set the authenticator params in the store
		// -------------------------------------------------
		ctx.Logger().Info("Set authenticator params")
		authenticatorParams := keepers.SmartAccountKeeper.GetParams(ctx)
		authenticatorParams.MaximumUnauthenticatedGas = MaximumUnauthenticatedGas
		authenticatorParams.IsSmartAccountActive = IsSmartAccountActive
		keepers.SmartAccountKeeper.SetParams(ctx, authenticatorParams)

		// -------------------------------------------------
		// Set the ICQ params in the store
		// -------------------------------------------------
		ctx.Logger().Info("Set ICQKeeper params")
		icqparams := icqtypes.DefaultParams()
		icqparams.AllowQueries = wasmbinding.GetStargateWhitelistedPaths()
		// Adding SmartContractState query to allowlist
		icqparams.AllowQueries = append(icqparams.AllowQueries, "/cosmwasm.wasm.v1.Query/SmartContractState")
		err = keepers.ICQKeeper.SetParams(ctx, icqparams)
		if err != nil {
			return nil, err
		}

		// -------------------------------------------------
		// Set the ICA Host params in the store
		// -------------------------------------------------
		ctx.Logger().Info("Set ICAHostKeeper params")
		// Allow all messages
		hostParams := icahosttypes.Params{
			HostEnabled:   true,
			AllowMessages: []string{"*"},
		}
		keepers.ICAHostKeeper.SetParams(ctx, hostParams)

		return migrations, nil
	}
}

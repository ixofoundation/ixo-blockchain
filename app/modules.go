package app

import (
	"cosmossdk.io/x/evidence"
	evidencetypes "cosmossdk.io/x/evidence/types"
	"cosmossdk.io/x/feegrant"
	feegrantmodule "cosmossdk.io/x/feegrant/module"
	"cosmossdk.io/x/upgrade"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/consensus"
	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/mint"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	sdkparams "github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	packetforward "github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v8/packetforward"
	packetforwardtypes "github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v8/packetforward/types"
	icq "github.com/cosmos/ibc-apps/modules/async-icq/v8"
	icqtypes "github.com/cosmos/ibc-apps/modules/async-icq/v8/types"
	ibchooks "github.com/cosmos/ibc-apps/modules/ibc-hooks/v8"
	ibchookstypes "github.com/cosmos/ibc-apps/modules/ibc-hooks/v8/types"
	"github.com/cosmos/ibc-go/modules/capability"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	ica "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts"
	icatypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/types"
	ibcfee "github.com/cosmos/ibc-go/v8/modules/apps/29-fee"
	ibcfeetypes "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/types"
	"github.com/cosmos/ibc-go/v8/modules/apps/transfer"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v8/modules/core"
	ibchost "github.com/cosmos/ibc-go/v8/modules/core/exported"
	ibctm "github.com/cosmos/ibc-go/v8/modules/light-clients/07-tendermint"
	"github.com/ixofoundation/ixo-blockchain/v3/x/bonds"
	bondstypes "github.com/ixofoundation/ixo-blockchain/v3/x/bonds/types"
	claimsmodule "github.com/ixofoundation/ixo-blockchain/v3/x/claims"
	claimsmoduletypes "github.com/ixofoundation/ixo-blockchain/v3/x/claims/types"
	entitymodule "github.com/ixofoundation/ixo-blockchain/v3/x/entity"
	entityclient "github.com/ixofoundation/ixo-blockchain/v3/x/entity/client"
	entitytypes "github.com/ixofoundation/ixo-blockchain/v3/x/entity/types"
	iidmodule "github.com/ixofoundation/ixo-blockchain/v3/x/iid"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v3/x/iid/types"
	smartaccount "github.com/ixofoundation/ixo-blockchain/v3/x/smart-account"
	smartaccounttypes "github.com/ixofoundation/ixo-blockchain/v3/x/smart-account/types"
	tokenmodule "github.com/ixofoundation/ixo-blockchain/v3/x/token"
	tokenclient "github.com/ixofoundation/ixo-blockchain/v3/x/token/client"
	tokentypes "github.com/ixofoundation/ixo-blockchain/v3/x/token/types"
)

// moduleAccountPermissions defines module account permissions
var moduleAccountPermissions = map[string][]string{
	// Standard Cosmos module accounts
	authtypes.FeeCollectorName:     nil,
	distrtypes.ModuleName:          nil,
	minttypes.ModuleName:           {authtypes.Minter},
	stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
	stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
	govtypes.ModuleName:            {authtypes.Burner},
	ibctransfertypes.ModuleName:    {authtypes.Minter, authtypes.Burner},
	ibcfeetypes.ModuleName:         nil,
	icatypes.ModuleName:            nil,
	icqtypes.ModuleName:            nil,
	wasmtypes.ModuleName:           {authtypes.Burner},

	// Custom ixo module accounts
	bondstypes.BondsMintBurnAccount:       {authtypes.Minter, authtypes.Burner},
	bondstypes.BatchesIntermediaryAccount: nil,
	bondstypes.BondsReserveAccount:        nil,
	smartaccounttypes.ModuleName:          nil,
}

// appModules return modules to initialize module manager.
func appModules(
	app *IxoApp,
	appCodec codec.Codec,
	txConfig client.TxConfig,
	skipGenesisInvariants bool,
) []module.AppModule {
	return []module.AppModule{
		// Standard Cosmos AppModules
		genutil.NewAppModule(app.AccountKeeper, app.StakingKeeper, app.BaseApp, txConfig),
		auth.NewAppModule(appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts, app.GetSubspace(authtypes.ModuleName)),
		vesting.NewAppModule(app.AccountKeeper, app.BankKeeper),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper, app.GetSubspace(banktypes.ModuleName)),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper, false),
		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper, app.GetSubspace(govtypes.ModuleName)),
		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper, nil, app.GetSubspace(minttypes.ModuleName)),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper, app.GetSubspace(slashingtypes.ModuleName), app.interfaceRegistry),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper, app.GetSubspace(distrtypes.ModuleName)),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper, app.GetSubspace(stakingtypes.ModuleName)),
		upgrade.NewAppModule(app.UpgradeKeeper, app.AccountKeeper.AddressCodec()),
		evidence.NewAppModule(app.EvidenceKeeper),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		authzmodule.NewAppModule(appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		ibc.NewAppModule(app.IBCKeeper),
		ibctm.NewAppModule(),
		sdkparams.NewAppModule(app.ParamsKeeper),
		consensus.NewAppModule(appCodec, app.ConsensusParamsKeeper),
		wasm.NewAppModule(appCodec, &app.WasmKeeper, app.StakingKeeper, app.AccountKeeper, app.BankKeeper, app.BaseApp.MsgServiceRouter(), app.GetSubspace(wasmtypes.ModuleName)),
		ibcfee.NewAppModule(app.IBCFeeKeeper),
		ica.NewAppModule(&app.ICAControllerKeeper, &app.ICAHostKeeper),
		icq.NewAppModule(app.ICQKeeper, app.GetSubspace(icqtypes.ModuleName)),
		packetforward.NewAppModule(app.PacketForwardKeeper, app.GetSubspace(packetforwardtypes.ModuleName)),
		crisis.NewAppModule(app.CrisisKeeper, skipGenesisInvariants, app.GetSubspace(crisistypes.ModuleName)),
		transfer.NewAppModule(app.TransferKeeper),
		ibchooks.NewAppModule(app.AccountKeeper),

		// Custom ixo AppModules
		iidmodule.NewAppModule(app.appCodec, app.IidKeeper),
		bonds.NewAppModule(app.BondsKeeper, app.AccountKeeper),
		entitymodule.NewAppModule(app.EntityKeeper),
		tokenmodule.NewAppModule(app.TokenKeeper),
		claimsmodule.NewAppModule(app.ClaimsKeeper),
		smartaccount.NewAppModule(appCodec, *app.SmartAccountKeeper),
	}
}

// ModuleBasics defines the module BasicManager that is in charge of setting up basic,
// non-dependant module elements, such as codec registration
// and genesis verification.
func newBasicManagerFromManager(app *IxoApp) module.BasicManager {
	basicManager := module.NewBasicManagerFromManager(
		app.ModuleManager,
		map[string]module.AppModuleBasic{
			genutiltypes.ModuleName: genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),
			govtypes.ModuleName: gov.NewAppModuleBasic(
				[]govclient.ProposalHandler{
					paramsclient.ProposalHandler,
					entityclient.ProposalHandler,
					tokenclient.ProposalHandler,
				},
			),
		})
	basicManager.RegisterLegacyAminoCodec(app.legacyAmino)
	basicManager.RegisterInterfaces(app.interfaceRegistry)
	return basicManager
}

// TODO: uncomment while no tests, re-add when tests are added
// // simulationModules returns modules for simulation manager
// // define the order of the modules for deterministic simulationss
// func simulationModules(
// 	app *IxoApp,
// 	encodingConfig appparams.EncodingConfig,
// 	_ bool,
// ) []module.AppModuleSimulation {
// 	appCodec := encodingConfig.Marshaler

// 	return []module.AppModuleSimulation{
// 		auth.NewAppModule(appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts),
// 		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
// 		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
// 		gov.NewAppModule(appCodec, *app.GovKeeper, app.AccountKeeper, app.BankKeeper),
// 		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper),
// 		staking.NewAppModule(appCodec, *app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
// 		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
// 		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
// 		sdkparams.NewAppModule(app.ParamsKeeper),
// 		wasm.NewAppModule(appCodec, &app.WasmKeeper, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
// 		authzmodule.NewAppModule(appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
// 		evidence.NewAppModule(app.EvidenceKeeper),
// 		ibc.NewAppModule(app.IBCKeeper),
// 		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
// 		transfer.NewAppModule(app.TransferKeeper),
// 	}
// }

// OrderBeginBlockers returns the order of BeginBlockers, by module name.
func OrderBeginBlockers() []string {
	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	return []string{

		// Next 7 is Staking ordering, dont change
		minttypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		evidencetypes.ModuleName,
		stakingtypes.ModuleName,
		authtypes.ModuleName,
		ibchost.ModuleName,

		banktypes.ModuleName,
		genutiltypes.ModuleName,
		crisistypes.ModuleName,
		paramstypes.ModuleName,
		capabilitytypes.ModuleName,
		govtypes.ModuleName,
		ibctransfertypes.ModuleName,
		vestingtypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		icatypes.ModuleName,
		ibcfeetypes.ModuleName,
		wasmtypes.ModuleName,
		packetforwardtypes.ModuleName,
		icqtypes.ModuleName,
		consensusparamtypes.ModuleName,
		ibchookstypes.ModuleName,

		// Custom ixo modules
		smartaccounttypes.ModuleName,
		bondstypes.ModuleName,
		iidtypes.ModuleName,
		entitytypes.ModuleName,
		tokentypes.ModuleName,
		claimsmoduletypes.ModuleName,
	}
}

// OrderEndBlockers returns EndBlockers (crisis, govtypes, staking) with no relative order.
func OrderEndBlockers() []string {
	return []string{
		crisistypes.ModuleName,
		// Standard Cosmos modules
		// Staking must be after gov.
		govtypes.ModuleName,
		stakingtypes.ModuleName,

		genutiltypes.ModuleName,
		feegrant.ModuleName,
		distrtypes.ModuleName,
		evidencetypes.ModuleName,
		banktypes.ModuleName,
		upgradetypes.ModuleName,
		ibchost.ModuleName,
		paramstypes.ModuleName,
		authtypes.ModuleName,
		minttypes.ModuleName,
		vestingtypes.ModuleName,
		capabilitytypes.ModuleName,
		slashingtypes.ModuleName,
		ibctransfertypes.ModuleName,
		authz.ModuleName,
		icatypes.ModuleName,
		ibcfeetypes.ModuleName,
		wasmtypes.ModuleName,
		packetforwardtypes.ModuleName,
		icqtypes.ModuleName,
		consensusparamtypes.ModuleName,
		ibchookstypes.ModuleName,

		// Custom ixo modules
		smartaccounttypes.ModuleName,
		iidtypes.ModuleName,
		entitytypes.ModuleName,
		tokentypes.ModuleName,
		bondstypes.ModuleName,
		claimsmoduletypes.ModuleName,
	}
}

// OrderInitGenesis returns module names in order for init genesis calls.
func OrderInitGenesis() []string {
	return []string{
		// Standard Cosmos modules
		// Capability module must occur first so that it can initialize any capabilities
		// so that other modules that want to create or claim capabilities afterwards in InitChain
		// can do so safely.
		capabilitytypes.ModuleName,

		// NOTE: The genutils module must occur after staking so that pools are
		// properly initialized with tokens from genesis accounts.
		// NOTE: The genutils module must also occur after auth so that it can access the params from auth.
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		slashingtypes.ModuleName,
		govtypes.ModuleName,
		minttypes.ModuleName,
		crisistypes.ModuleName,
		ibchost.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		ibctransfertypes.ModuleName,
		consensusparamtypes.ModuleName,
		upgradetypes.ModuleName,
		paramstypes.ModuleName,
		vestingtypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		icatypes.ModuleName,
		ibcfeetypes.ModuleName,
		// wasm after ibc transfer
		wasmtypes.ModuleName,
		icqtypes.ModuleName,
		packetforwardtypes.ModuleName,
		ibchookstypes.ModuleName,

		// Custom ixo modules
		smartaccounttypes.ModuleName,
		iidtypes.ModuleName,
		bondstypes.ModuleName,
		tokentypes.ModuleName,
		entitytypes.ModuleName,
		claimsmoduletypes.ModuleName,
	}
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *IxoApp) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

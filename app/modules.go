package app

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmclient "github.com/CosmWasm/wasmd/x/wasm/client"
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
	"github.com/cosmos/cosmos-sdk/x/capability"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
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
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	icq "github.com/cosmos/ibc-apps/modules/async-icq/v4"
	icqtypes "github.com/cosmos/ibc-apps/modules/async-icq/v4/types"
	ica "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts"
	icatypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/types"
	ibcfee "github.com/cosmos/ibc-go/v4/modules/apps/29-fee"
	ibcfeetypes "github.com/cosmos/ibc-go/v4/modules/apps/29-fee/types"
	"github.com/cosmos/ibc-go/v4/modules/apps/transfer"
	ibctransfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v4/modules/core"
	ibcclientclient "github.com/cosmos/ibc-go/v4/modules/core/02-client/client"
	ibchost "github.com/cosmos/ibc-go/v4/modules/core/24-host"
	intertx "github.com/cosmos/interchain-accounts/x/inter-tx"
	intertxtypes "github.com/cosmos/interchain-accounts/x/inter-tx/types"
	appparams "github.com/ixofoundation/ixo-blockchain/v2/app/params"
	"github.com/ixofoundation/ixo-blockchain/v2/x/bonds"
	bondstypes "github.com/ixofoundation/ixo-blockchain/v2/x/bonds/types"
	claimsmodule "github.com/ixofoundation/ixo-blockchain/v2/x/claims"
	claimsmoduletypes "github.com/ixofoundation/ixo-blockchain/v2/x/claims/types"
	entitymodule "github.com/ixofoundation/ixo-blockchain/v2/x/entity"
	entityclient "github.com/ixofoundation/ixo-blockchain/v2/x/entity/client"
	entitytypes "github.com/ixofoundation/ixo-blockchain/v2/x/entity/types"
	iidmodule "github.com/ixofoundation/ixo-blockchain/v2/x/iid"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v2/x/iid/types"
	tokenmodule "github.com/ixofoundation/ixo-blockchain/v2/x/token"
	tokenclient "github.com/ixofoundation/ixo-blockchain/v2/x/token/client"
	tokentypes "github.com/ixofoundation/ixo-blockchain/v2/x/token/types"
	packetforward "github.com/strangelove-ventures/packet-forward-middleware/v4/router"
	packetforwardtypes "github.com/strangelove-ventures/packet-forward-middleware/v4/router/types"
)

// module account permissions
var maccPerms = map[string][]string{
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
	wasm.ModuleName:                {authtypes.Burner},

	// Custom ixo module accounts
	bondstypes.BondsMintBurnAccount:       {authtypes.Minter, authtypes.Burner},
	bondstypes.BatchesIntermediaryAccount: nil,
	bondstypes.BondsReserveAccount:        nil,
}

// ModuleBasics defines the module BasicManager is in charge of setting up basic,
// non-dependant module elements, such as codec registration
// and genesis verification.
var ModuleBasics = module.NewBasicManager(
	// Standard Cosmos modules
	auth.AppModuleBasic{},
	genutil.AppModuleBasic{},
	authzmodule.AppModuleBasic{},
	bank.AppModuleBasic{},
	capability.AppModuleBasic{},
	staking.AppModuleBasic{},
	mint.AppModuleBasic{},
	distr.AppModuleBasic{},
	sdkparams.AppModuleBasic{},
	crisis.AppModuleBasic{},
	slashing.AppModuleBasic{},
	ibc.AppModuleBasic{},
	upgrade.AppModuleBasic{},
	evidence.AppModuleBasic{},
	transfer.AppModuleBasic{},
	vesting.AppModuleBasic{},
	feegrantmodule.AppModuleBasic{},
	gov.NewAppModuleBasic(
		append(
			wasmclient.ProposalHandlers,
			paramsclient.ProposalHandler,
			distrclient.ProposalHandler,
			upgradeclient.ProposalHandler,
			upgradeclient.CancelProposalHandler,
			entityclient.ProposalHandler,
			tokenclient.ProposalHandler,
			ibcclientclient.UpdateClientProposalHandler,
			ibcclientclient.UpgradeProposalHandler,
		)...,
	),
	wasm.AppModuleBasic{},
	ica.AppModuleBasic{},
	intertx.AppModuleBasic{},
	ibcfee.AppModuleBasic{},
	icq.AppModuleBasic{},
	packetforward.AppModuleBasic{},

	// Custom ixo modules
	iidmodule.AppModuleBasic{},
	bonds.AppModuleBasic{},
	entitymodule.AppModuleBasic{},
	tokenmodule.AppModuleBasic{},
	claimsmodule.AppModuleBasic{},
)

// appModules return modules to initialize module manager.
func appModules(
	app *IxoApp,
	encodingConfig appparams.EncodingConfig,
	skipGenesisInvariants bool,
) []module.AppModule {
	appCodec := encodingConfig.Marshaler

	return []module.AppModule{
		// Standard Cosmos AppModules
		genutil.NewAppModule(
			app.AccountKeeper, app.StakingKeeper, app.BaseApp.DeliverTx,
			encodingConfig.TxConfig,
		),
		auth.NewAppModule(appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts),
		vesting.NewAppModule(app.AccountKeeper, app.BankKeeper),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
		crisis.NewAppModule(&app.CrisisKeeper, skipGenesisInvariants),
		gov.NewAppModule(appCodec, *app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		staking.NewAppModule(appCodec, *app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		upgrade.NewAppModule(app.UpgradeKeeper),
		wasm.NewAppModule(appCodec, &app.WasmKeeper, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		ibc.NewAppModule(app.IBCKeeper),
		ica.NewAppModule(&app.ICAControllerKeeper, &app.ICAHostKeeper),
		intertx.NewAppModule(appCodec, app.InterTxKeeper),
		sdkparams.NewAppModule(app.ParamsKeeper),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		authzmodule.NewAppModule(appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		transfer.NewAppModule(app.TransferKeeper),
		ibcfee.NewAppModule(app.IBCFeeKeeper),
		packetforward.NewAppModule(app.PacketForwardKeeper),
		icq.NewAppModule(*app.ICQKeeper),

		// Custom ixo AppModules
		iidmodule.NewAppModule(app.appCodec, app.IidKeeper),
		bonds.NewAppModule(app.BondsKeeper, app.AccountKeeper),
		entitymodule.NewAppModule(app.EntityKeeper),
		tokenmodule.NewAppModule(app.TokenKeeper),
		claimsmodule.NewAppModule(app.ClaimsKeeper),
	}
}

// simulationModules returns modules for simulation manager
// define the order of the modules for deterministic simulationss
func simulationModules(
	app *IxoApp,
	encodingConfig appparams.EncodingConfig,
	_ bool,
) []module.AppModuleSimulation {
	appCodec := encodingConfig.Marshaler

	return []module.AppModuleSimulation{
		auth.NewAppModule(appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
		gov.NewAppModule(appCodec, *app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper),
		staking.NewAppModule(appCodec, *app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		sdkparams.NewAppModule(app.ParamsKeeper),
		wasm.NewAppModule(appCodec, &app.WasmKeeper, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		authzmodule.NewAppModule(appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		evidence.NewAppModule(app.EvidenceKeeper),
		ibc.NewAppModule(app.IBCKeeper),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		transfer.NewAppModule(app.AppKeepers.TransferKeeper),
	}
}

// orderBeginBlockers returns the order of BeginBlockers, by module name.
func orderBeginBlockers() []string {
	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	return []string{
		// Standard Cosmos modules
		// Upgrades should be run VERY first
		upgradetypes.ModuleName,
		minttypes.ModuleName,
		// Next 4 is Staking ordering
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		evidencetypes.ModuleName,
		stakingtypes.ModuleName,
		ibchost.ModuleName,
		banktypes.ModuleName,
		genutiltypes.ModuleName,
		crisistypes.ModuleName,
		paramstypes.ModuleName,
		authtypes.ModuleName,
		capabilitytypes.ModuleName,
		govtypes.ModuleName,
		ibctransfertypes.ModuleName,
		vestingtypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		icatypes.ModuleName,
		ibcfeetypes.ModuleName,
		intertxtypes.ModuleName,
		wasm.ModuleName,
		packetforwardtypes.ModuleName,
		icqtypes.ModuleName,

		// Custom ixo modules
		bondstypes.ModuleName,
		iidtypes.ModuleName,
		entitytypes.ModuleName,
		tokentypes.ModuleName,
		claimsmoduletypes.ModuleName,
	}
}

// orderEndBlockers returns EndBlockers (crisis, govtypes, staking) with no relative order.
func orderEndBlockers() []string {
	return []string{
		// Standard Cosmos modules
		crisistypes.ModuleName,
		// Staking must be after gov.
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		distrtypes.ModuleName,
		evidencetypes.ModuleName,
		banktypes.ModuleName,
		upgradetypes.ModuleName,
		ibchost.ModuleName,
		paramstypes.ModuleName,
		authtypes.ModuleName,
		minttypes.ModuleName,
		genutiltypes.ModuleName,
		vestingtypes.ModuleName,
		capabilitytypes.ModuleName,
		slashingtypes.ModuleName,
		ibctransfertypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		icatypes.ModuleName,
		ibcfeetypes.ModuleName,
		intertxtypes.ModuleName,
		wasm.ModuleName,
		packetforwardtypes.ModuleName,
		icqtypes.ModuleName,

		// Custom ixo modules
		iidtypes.ModuleName,
		entitytypes.ModuleName,
		tokentypes.ModuleName,
		bondstypes.ModuleName,
		claimsmoduletypes.ModuleName,
	}
}

// orderInitGenesis returns module names in order for init genesis calls.
func orderInitBlockers() []string {
	return []string{
		// Standard Cosmos modules
		// Capability module must occur first so that it can initialize any capabilities
		// so that other modules that want to create or claim capabilities afterwards in InitChain
		// can do so safely.
		capabilitytypes.ModuleName,
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
		upgradetypes.ModuleName,
		paramstypes.ModuleName,
		vestingtypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		icatypes.ModuleName,
		ibcfeetypes.ModuleName,
		intertxtypes.ModuleName,
		// wasm after ibc transfer
		wasm.ModuleName,
		icqtypes.ModuleName,
		packetforwardtypes.ModuleName,

		// Custom ixo modules
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

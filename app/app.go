package app

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	srvconfig "github.com/cosmos/cosmos-sdk/server/config"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/gorilla/mux"
	"github.com/ixofoundation/ixo-blockchain/app/params"
	"github.com/ixofoundation/ixo-blockchain/x/bonds"
	bondskeeper "github.com/ixofoundation/ixo-blockchain/x/bonds/keeper"
	bondstypes "github.com/ixofoundation/ixo-blockchain/x/bonds/types"
	"github.com/ixofoundation/ixo-blockchain/x/payments"
	paymentskeeper "github.com/ixofoundation/ixo-blockchain/x/payments/keeper"
	paymentstypes "github.com/ixofoundation/ixo-blockchain/x/payments/types"
	"github.com/ixofoundation/ixo-blockchain/x/project"
	"github.com/rakyll/statik/fs"
	"net/http"

	//"github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/capability"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"io"
	"os"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authrest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	ibctransferkeeper "github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer/keeper"
	ibctransfertypes "github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer/types"
	ibc "github.com/cosmos/cosmos-sdk/x/ibc/core"
	ibcclient "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client"
	porttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/05-port/types"
	ibchost "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	ibckeeper "github.com/cosmos/cosmos-sdk/x/ibc/core/keeper"
	"github.com/cosmos/cosmos-sdk/x/mint"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	sdkparams "github.com/cosmos/cosmos-sdk/x/params"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	//simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	//"github.com/cosmos/cosmos-sdk/x/supply"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	dbm "github.com/tendermint/tm-db"

	//TODO uncomment
	//"github.com/ixofoundation/ixo-blockchain/x/bonds"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	ixotypes "github.com/ixofoundation/ixo-blockchain/x/ixo/types"
	projectkeeper "github.com/ixofoundation/ixo-blockchain/x/project/keeper"
	//"github.com/ixofoundation/ixo-blockchain/x/payments"
	projecttypes "github.com/ixofoundation/ixo-blockchain/x/project/types"
	//"github.com/ixofoundation/ixo-blockchain/x/treasury"
)

const (
	appName              = "ixoApp"
	Bech32MainPrefix     = "ixo"
	Bech32PrefixAccAddr  = Bech32MainPrefix
	Bech32PrefixAccPub   = Bech32MainPrefix + sdk.PrefixPublic
	Bech32PrefixValAddr  = Bech32MainPrefix + sdk.PrefixValidator + sdk.PrefixOperator
	Bech32PrefixValPub   = Bech32MainPrefix + sdk.PrefixValidator + sdk.PrefixOperator + sdk.PrefixPublic
	Bech32PrefixConsAddr = Bech32MainPrefix + sdk.PrefixValidator + sdk.PrefixConsensus
	Bech32PrefixConsPub  = Bech32MainPrefix + sdk.PrefixValidator + sdk.PrefixConsensus + sdk.PrefixPublic
)

var (
	// default home directories for ixocli
	//DefaultCLIHome = os.ExpandEnv("$HOME/.ixocli")

	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome = os.ExpandEnv("$HOME/.ixod")

	// ModuleBasics defines the module BasicManager which is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		// Standard Cosmos modules
		auth.AppModuleBasic{},
		//supply.AppModuleBasic{}, //All `x/supply` types and APIs have been moved to `x/bank`
		genutil.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(
			paramsclient.ProposalHandler, distrclient.ProposalHandler, upgradeclient.ProposalHandler, upgradeclient.CancelProposalHandler,
		),
		sdkparams.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		ibc.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		transfer.AppModuleBasic{},
		vesting.AppModuleBasic{},

		// Custom ixo modules
		did.AppModuleBasic{}, //TODO uncomment rest of ixo modules
		payments.AppModuleBasic{},
		project.AppModuleBasic{},
		bonds.AppModuleBasic{},
		//treasury.AppModuleBasic{},
		//oracles.AppModuleBasic{},
	)

	// Module account permissions
	maccPerms = map[string][]string{
		// Standard Cosmos module accounts
		authtypes.FeeCollectorName:     nil,
		distrtypes.ModuleName:          nil,
		minttypes.ModuleName:           {authtypes.Minter},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:            {authtypes.Burner},
		ibctransfertypes.ModuleName: 	{authtypes.Minter, authtypes.Burner},

		// Custom ixo module accounts
		//TODO uncomment ixo modules
		bondstypes.BondsMintBurnAccount:       {authtypes.Minter, authtypes.Burner},
		bondstypes.BatchesIntermediaryAccount: nil,
		bondstypes.BondsReserveAccount:        nil,
		//treasury.ModuleName:              {authtypes.Minter, authtypes.Burner},
		paymentstypes.PayRemainderPool:        nil,
	}

	// module accounts that are allowed to receive tokens
	allowedReceivingModAcc = map[string]bool{
		distrtypes.ModuleName: true,
	}

	// Reserved payments module ID prefixes
	paymentsReservedIdPrefixes = []string{
		projecttypes.ModuleName,
	}
)

//// MakeCodec - custom tx codec
//func MakeCodec() *codec.LegacyAmino { //.Codec {
//	var cdc = codec.NewLegacyAmino() //codec.New()
//
//	// Register standard Cosmos codecs
//	ModuleBasics.RegisterLegacyAminoCodec(cdc)
//	vestingtypes.RegisterLegacyAminoCodec(cdc)
//	sdk.RegisterLegacyAminoCodec(cdc)
//	crypto.RegisterCrypto(cdc)
//	//codec.RegisterEvidences(cdc)
//
//	// Register ixo codec
//	ixo.RegisterCodec(cdc)
//
//	return cdc
//}

// MakeCodecs constructs the *std.Codec and *codec.LegacyAmino instances used by
// ixoapp. It is useful for tests and clients who do not want to construct the
// full ixoapp.
func MakeCodecs() (codec.Marshaler, *codec.LegacyAmino) {
	config := MakeEncodingConfig()
	return config.Marshaler, config.Amino
	// TODO register ixo codec?
}

// Verify app interface at compile time
var _ simapp.App = (*ixoApp)(nil)
var _ servertypes.Application = (*ixoApp)(nil)

// Extended ABCI application
type ixoApp struct {
	*bam.BaseApp      `json:"_bam_base_app,omitempty"`
	legacyAmino       *codec.LegacyAmino      `json:"legacy_amino,omitempty"`
	appCodec          codec.Marshaler         `json:"app_codec,omitempty"`
	interfaceRegistry types.InterfaceRegistry `json:"interface_registry,omitempty"`

	invCheckPeriod uint `json:"inv_check_period,omitempty"`

	// keys to access the substores
	keys    map[string]*sdk.KVStoreKey        `json:"keys,omitempty"`
	tkeys   map[string]*sdk.TransientStoreKey `json:"tkeys,omitempty"`
	memKeys map[string]*sdk.MemoryStoreKey    `json:"mem_keys,omitempty"`

	// keepers
	AccountKeeper    authkeeper.AccountKeeper `json:"account_keeper"`
	BankKeeper       bankkeeper.Keeper        `json:"bank_keeper,omitempty"`
	CapabilityKeeper *capabilitykeeper.Keeper `json:"capability_keeper,omitempty"`
	StakingKeeper    stakingkeeper.Keeper     `json:"staking_keeper"`
	SlashingKeeper   slashingkeeper.Keeper    `json:"slashing_keeper"`
	MintKeeper       mintkeeper.Keeper        `json:"mint_keeper"`
	DistrKeeper      distrkeeper.Keeper       `json:"distr_keeper"`
	GovKeeper        govkeeper.Keeper         `json:"gov_keeper"`
	CrisisKeeper     crisiskeeper.Keeper      `json:"crisis_keeper"`
	UpgradeKeeper    upgradekeeper.Keeper     `json:"upgrade_keeper"`
	ParamsKeeper     paramskeeper.Keeper      `json:"params_keeper"`
	IBCKeeper        *ibckeeper.Keeper        `json:"ibc_keeper,omitempty"` // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	EvidenceKeeper   evidencekeeper.Keeper    `json:"evidence_keeper"`
	TransferKeeper   ibctransferkeeper.Keeper `json:"transfer_keeper"`

	// make scoped keepers public for test purposes
	ScopedIBCKeeper      capabilitykeeper.ScopedKeeper `json:"scoped_ibc_keeper"`
	ScopedTransferKeeper capabilitykeeper.ScopedKeeper `json:"scoped_transfer_keeper"`

	// Custom ixo keepers
	didKeeper did.Keeper `json:"did_keeper"`
	//TODO uncomment rest of ixo modules
	paymentsKeeper paymentskeeper.Keeper `json:"payments_keeper,omitempty"`
	projectKeeper  projectkeeper.Keeper  `json:"project_keeper"`
	bondsKeeper    bondskeeper.Keeper    `json:"bonds_keeper"`
	//oraclesKeeper  oracles.Keeper
	//treasuryKeeper treasury.Keeper

	// the module manager
	mm *module.Manager `json:"mm,omitempty"`

	// simulation manager
	sm *module.SimulationManager `json:"sm,omitempty"`
}

// TODO(?) Implement functions for servertypes.Application interface
// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *ixoApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig srvconfig.APIConfig) {
	//panic("implement me")

	clientCtx := apiSvr.ClientCtx
	rpc.RegisterRoutes(clientCtx, apiSvr.Router)
	// Register legacy tx routes.
	authrest.RegisterTxRoutes(clientCtx, apiSvr.Router)
	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register legacy and grpc-gateway routes for all modules.
	ModuleBasics.RegisterRESTRoutes(clientCtx, apiSvr.Router)
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// register swagger API from root so that other applications can override easily
	if apiConfig.Swagger {
		RegisterSwaggerAPI(clientCtx, apiSvr.Router)
	}
}

// RegisterSwaggerAPI registers swagger route with API Server
func RegisterSwaggerAPI(ctx client.Context, rtr *mux.Router) {
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}

	staticServer := http.FileServer(statikFS)
	rtr.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", staticServer))
}

//func (app *ixoApp) RegisterGRPCServer(g interface{}) {
//	panic("implement me")
//}

//func (app *ixoApp) RegisterGRPCServer(grpc.Server) {
//	panic("implement me")
//}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *ixoApp) RegisterTxService(clientCtx client.Context) {
	//panic("implement me")
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *ixoApp) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.interfaceRegistry)
}

// NewIxoApp returns a reference to an initialized IxoApp.
func NewIxoApp(
	logger log.Logger, db dbm.DB, traceStore io.Writer, loadLatest bool, skipUpgradeHeights map[int64]bool,
	homePath string, invCheckPeriod uint, encodingConfig params.EncodingConfig,
	appOpts servertypes.AppOptions, baseAppOptions ...func(*bam.BaseApp),
) *ixoApp {

	// TODO: Remove cdc in favor of appCodec once all modules are migrated.
	appCodec := encodingConfig.Marshaler
	legacyAmino := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	bApp := bam.NewBaseApp(appName, logger, db, encodingConfig.TxConfig.TxDecoder(), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetAppVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)

	keys := sdk.NewKVStoreKeys(
		// Standard Cosmos store keys
		authtypes.StoreKey, banktypes.StoreKey, stakingtypes.StoreKey,
		minttypes.StoreKey, distrtypes.StoreKey, slashingtypes.StoreKey,
		govtypes.StoreKey, paramstypes.StoreKey, ibchost.StoreKey, upgradetypes.StoreKey,
		evidencetypes.StoreKey, ibctransfertypes.StoreKey, capabilitytypes.StoreKey,

		// Custom ixo store keys
		// TODO uncomment ixo modules store keys
		did.StoreKey, paymentstypes.StoreKey,
		projecttypes.StoreKey,
		bondstypes.StoreKey,
		//treasury.StoreKey, oracles.StoreKey*/
	)
	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	app := &ixoApp{
		BaseApp:           bApp,
		legacyAmino:       legacyAmino,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		invCheckPeriod:    invCheckPeriod,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
	}

	// init params keeper and subspaces
	app.ParamsKeeper = initParamsKeeper(appCodec, legacyAmino, keys[paramstypes.StoreKey], tkeys[paramstypes.TStoreKey])

	// set the BaseApp's parameter store
	bApp.SetParamStore(app.ParamsKeeper.Subspace(bam.Paramspace).WithKeyTable(paramskeeper.ConsensusParamsKeyTable()))

	// add capability keeper and ScopeToModule for ibc module
	app.CapabilityKeeper = capabilitykeeper.NewKeeper(appCodec, keys[capabilitytypes.StoreKey], memKeys[capabilitytypes.MemStoreKey])
	scopedIBCKeeper := app.CapabilityKeeper.ScopeToModule(ibchost.ModuleName)
	scopedTransferKeeper := app.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)

	// add keepers (for standard Cosmos modules)
	app.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec, keys[authtypes.StoreKey], app.GetSubspace(authtypes.ModuleName), authtypes.ProtoBaseAccount, maccPerms,
	)
	app.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec, keys[banktypes.StoreKey], app.AccountKeeper, app.GetSubspace(banktypes.ModuleName), app.BlockedAddrs(),
	)
	stakingKeeper := stakingkeeper.NewKeeper(
		appCodec, keys[stakingtypes.StoreKey], app.AccountKeeper, app.BankKeeper, app.GetSubspace(stakingtypes.ModuleName),
	)
	app.MintKeeper = mintkeeper.NewKeeper(
		appCodec, keys[minttypes.StoreKey], app.GetSubspace(minttypes.ModuleName), &stakingKeeper,
		app.AccountKeeper, app.BankKeeper, authtypes.FeeCollectorName,
	)
	app.DistrKeeper = distrkeeper.NewKeeper(
		appCodec, keys[distrtypes.StoreKey], app.GetSubspace(distrtypes.ModuleName), app.AccountKeeper, app.BankKeeper,
		&stakingKeeper, authtypes.FeeCollectorName, app.ModuleAccountAddrs(),
	)
	app.SlashingKeeper = slashingkeeper.NewKeeper(
		appCodec, keys[slashingtypes.StoreKey], &stakingKeeper, app.GetSubspace(slashingtypes.ModuleName),
	)
	app.CrisisKeeper = crisiskeeper.NewKeeper(
		app.GetSubspace(crisistypes.ModuleName), invCheckPeriod, app.BankKeeper, authtypes.FeeCollectorName,
	)
	app.UpgradeKeeper = upgradekeeper.NewKeeper(skipUpgradeHeights, keys[upgradetypes.StoreKey], appCodec, homePath)

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	app.StakingKeeper = *stakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(app.DistrKeeper.Hooks(), app.SlashingKeeper.Hooks()),
	)

	// Create IBC Keeper
	app.IBCKeeper = ibckeeper.NewKeeper(
		appCodec, keys[ibchost.StoreKey], app.GetSubspace(ibchost.ModuleName), app.StakingKeeper, scopedIBCKeeper,
	)

	// register the proposal types
	govRouter := govtypes.NewRouter()
	govRouter.AddRoute(govtypes.RouterKey, govtypes.ProposalHandler).
		AddRoute(paramproposal.RouterKey, sdkparams.NewParamChangeProposalHandler(app.ParamsKeeper)).
		AddRoute(distrtypes.RouterKey, distr.NewCommunityPoolSpendProposalHandler(app.DistrKeeper)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.UpgradeKeeper)).
		AddRoute(ibchost.RouterKey, ibcclient.NewClientUpdateProposalHandler(app.IBCKeeper.ClientKeeper))
	app.GovKeeper = govkeeper.NewKeeper(
		appCodec, keys[govtypes.StoreKey], app.GetSubspace(govtypes.ModuleName), app.AccountKeeper, app.BankKeeper,
		&stakingKeeper, govRouter,
	)

	// Create Transfer Keepers
	app.TransferKeeper = ibctransferkeeper.NewKeeper(
		appCodec, keys[ibctransfertypes.StoreKey], app.GetSubspace(ibctransfertypes.ModuleName),
		app.IBCKeeper.ChannelKeeper, &app.IBCKeeper.PortKeeper,
		app.AccountKeeper, app.BankKeeper, scopedTransferKeeper,
	)
	transferModule := transfer.NewAppModule(app.TransferKeeper)

	// Create static IBC router, add transfer route, then set and seal it
	ibcRouter := porttypes.NewRouter()
	ibcRouter.AddRoute(ibctransfertypes.ModuleName, transferModule)
	app.IBCKeeper.SetRouter(ibcRouter)

	// create evidence keeper with router
	evidenceKeeper := evidencekeeper.NewKeeper(
		appCodec, keys[evidencetypes.StoreKey], &app.StakingKeeper, app.SlashingKeeper,
	)
	// If evidence needs to be handled for the app, set routes in router here and seal
	app.EvidenceKeeper = *evidenceKeeper

	/****  Module Options ****/

	// NOTE: we may consider parsing `appOpts` inside module constructors. For the moment
	// we prefer to be more strict in what arguments the modules expect.
	//var skipGenesisInvariants = cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))
	var skipGenesisInvariants = false
	opt := appOpts.Get(crisis.FlagSkipGenesisInvariants)
	if opt, ok := opt.(bool); ok {
		skipGenesisInvariants = opt
	}

	//TODO replace below keepers (for custom ixo modules)
	// app.didKeeper = did.NewKeeper(app.cdc, keys[did.StoreKey])
	//app.paymentsKeeper = payments.NewKeeper(app.cdc, keys[payments.StoreKey],
	//	app.BankKeeper, app.didKeeper, paymentsReservedIdPrefixes)
	//app.projectKeeper = project.NewKeeper(app.cdc, keys[project.StoreKey], app.subspaces[project.ModuleName],
	//	app.AccountKeeper, app.didKeeper, app.paymentsKeeper)
	//app.bondsKeeper = bonds.NewKeeper(app.BankKeeper, app.SupplyKeeper, app.AccountKeeper,
	//	app.StakingKeeper, app.didKeeper, keys[bonds.StoreKey], app.subspaces[bonds.ModuleName], app.cdc)
	//app.oraclesKeeper = oracles.NewKeeper(app.cdc, keys[oracles.StoreKey])
	//app.treasuryKeeper = treasury.NewKeeper(app.cdc, keys[treasury.StoreKey], app.BankKeeper,
	//	app.oraclesKeeper, app.SupplyKeeper, app.didKeeper)

	app.didKeeper = did.NewKeeper(app.appCodec, keys[did.StoreKey]) // not what Cosmos uses because keeper is different
	app.bondsKeeper = bondskeeper.NewKeeper(app.BankKeeper, app.AccountKeeper, app.StakingKeeper, app.didKeeper,
		keys[bondstypes.StoreKey], app.GetSubspace(bondstypes.ModuleName), app.appCodec)
	app.projectKeeper = projectkeeper.NewKeeper(app.appCodec, keys[projecttypes.StoreKey],
		app.GetSubspace(projecttypes.ModuleName), app.AccountKeeper, app.didKeeper, app.paymentsKeeper)
	app.paymentsKeeper = paymentskeeper.NewKeeper(app.appCodec, keys[paymentstypes.StoreKey],
		app.BankKeeper, app.didKeeper, paymentsReservedIdPrefixes)
	// TODO add the rest of ixo modules keeper

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.
	app.mm = module.NewManager(
		// Standard Cosmos appmodules
		genutil.NewAppModule(
			app.AccountKeeper, app.StakingKeeper, app.BaseApp.DeliverTx,
			encodingConfig.TxConfig,
		),
		auth.NewAppModule(appCodec, app.AccountKeeper, nil),
		vesting.NewAppModule(app.AccountKeeper, app.BankKeeper),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
		crisis.NewAppModule(&app.CrisisKeeper, skipGenesisInvariants),
		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		upgrade.NewAppModule(app.UpgradeKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		ibc.NewAppModule(app.IBCKeeper),
		sdkparams.NewAppModule(app.ParamsKeeper),
		transferModule,

		// Custom ixo AppModules
		did.NewAppModule(app.didKeeper), //TODO uncomment rest of ixo modules
		payments.NewAppModule(app.paymentsKeeper, app.BankKeeper),
		project.NewAppModule(app.projectKeeper, app.paymentsKeeper, app.BankKeeper),
		bonds.NewAppModule(app.bondsKeeper, app.AccountKeeper),
		//treasury.NewAppModule(app.treasuryKeeper),
		//oracles.NewAppModule(app.oraclesKeeper),
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	app.mm.SetOrderBeginBlockers(
		// Standard Cosmos modules
		upgradetypes.ModuleName, minttypes.ModuleName, distrtypes.ModuleName, slashingtypes.ModuleName,
		evidencetypes.ModuleName, stakingtypes.ModuleName, ibchost.ModuleName,
		// Custom ixo modules
		//TODO uncomment rest of ixo modules
		bondstypes.ModuleName,
	)
	app.mm.SetOrderEndBlockers(
		// Standard Cosmos modules
		crisistypes.ModuleName, govtypes.ModuleName, stakingtypes.ModuleName,
		// Custom ixo modules
		//TODO uncomment rest of ixo modules
		bondstypes.ModuleName,
		paymentstypes.ModuleName,
	)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	app.mm.SetOrderInitGenesis(
		// Standard Cosmos modules
		capabilitytypes.ModuleName, authtypes.ModuleName, banktypes.ModuleName, distrtypes.ModuleName, stakingtypes.ModuleName,
		slashingtypes.ModuleName, govtypes.ModuleName, minttypes.ModuleName, crisistypes.ModuleName,
		ibchost.ModuleName, genutiltypes.ModuleName, evidencetypes.ModuleName, ibctransfertypes.ModuleName,
		// Custom ixo modules
		//TODO uncomment rest of ixo modules
		did.ModuleName,
		projecttypes.ModuleName,
		paymentstypes.ModuleName,
		bondstypes.ModuleName, //treasury.ModuleName, oracles.ModuleName,
	)

	app.mm.RegisterInvariants(&app.CrisisKeeper)
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), encodingConfig.Amino)
	app.mm.RegisterServices(module.NewConfigurator(app.MsgServiceRouter(), app.GRPCQueryRouter()))

	// add test gRPC service for testing gRPC queries in isolation
	testdata.RegisterQueryServer(app.GRPCQueryRouter(), testdata.QueryImpl{})

	// create the simulation manager and define the order of the modules for deterministic simulations
	//
	// NOTE: this is not required apps that don't use the simulator for fuzz testing
	// transactions
	app.sm = module.NewSimulationManager(
		auth.NewAppModule(appCodec, app.AccountKeeper, authsims.RandomGenesisAccounts),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		sdkparams.NewAppModule(app.ParamsKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		ibc.NewAppModule(app.IBCKeeper),
		transferModule,
	)

	app.sm.RegisterStoreDecoders()

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetAnteHandler(NewIxoAnteHandler(app, encodingConfig)) //TODO encodingConfig
	app.SetEndBlocker(app.EndBlocker)

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(err.Error())
		}

		// Initialize and seal the capability keeper so all persistent capabilities
		// are loaded in-memory and prevent any further modules from creating scoped
		// sub-keepers.
		// This must be done during creation of baseapp rather than in InitChain so
		// that in-memory capabilities get regenerated on app restart.
		// Note that since this reads from the store, we can only perform it when
		// `loadLatest` is set to true.
		ctx := app.BaseApp.NewUncachedContext(true, tmproto.Header{})
		app.CapabilityKeeper.InitializeAndSeal(ctx)
	}

	app.ScopedIBCKeeper = scopedIBCKeeper
	app.ScopedTransferKeeper = scopedTransferKeeper

	return app
}

// Name returns the name of the App
func (app *ixoApp) Name() string { return app.BaseApp.Name() }

// LegacyAmino returns ixoApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *ixoApp) LegacyAmino() *codec.LegacyAmino {
	return codec.NewLegacyAmino()
}

// BeginBlocker application updates every begin block
func (app *ixoApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker application updates every end block
func (app *ixoApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

// InitChainer application update at chain initialization
func (app *ixoApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	app.legacyAmino.MustUnmarshalJSON(req.AppStateBytes, &genesisState) //TODO change
	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

// LoadHeight loads a particular height
func (app *ixoApp) LoadHeight(height int64) error {
	return app.LoadVersion(height) //, app.keys[bam.MainStoreKey])
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *ixoApp) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// SimulationManager implements the SimulationApp interface
func (app *ixoApp) SimulationManager() *module.SimulationManager {
	return app.sm
}

// BlockedAddrs returns all the app's module account addresses black listed for receiving tokens.
func (app *ixoApp) BlockedAddrs() map[string]bool {
	blockedAddrs := make(map[string]bool)
	for acc := range maccPerms {
		blockedAddrs[authtypes.NewModuleAddress(acc).String()] = !allowedReceivingModAcc[acc]
	}

	return blockedAddrs
}

// AppCodec returns ixoApp's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *ixoApp) AppCodec() codec.Marshaler {
	return app.appCodec
}

// InterfaceRegistry returns ixoApp's InterfaceRegistry
func (app *ixoApp) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *ixoApp) GetKey(storeKey string) *sdk.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *ixoApp) GetTKey(storeKey string) *sdk.TransientStoreKey {
	return app.tkeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
func (app *ixoApp) GetMemKey(storeKey string) *sdk.MemoryStoreKey {
	return app.memKeys[storeKey]
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *ixoApp) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}
	return dupMaccPerms
}

func NewIxoAnteHandler(app *ixoApp, encodingConfig params.EncodingConfig) sdk.AnteHandler { //TODO encodingConfig

	// The AnteHandler needs to get the signer's pubkey to verify signatures,
	// charge gas fees (to the corresponding address), and for other purposes.
	//
	// The default Cosmos AnteHandler fetches a signer address' pubkey from the
	// GetPubKey() function after querying the account from the account keeper.
	//
	// In the case of ixo, since signers are DIDs rather than addresses, we get
	// the DID Doc containing the pubkey from the did/project module (depending
	// if signer is a user or a project, respectively).
	//
	// This is what PubKeyGetters achieve.
	//
	// To get a pubkey from the did/project, the did/project must have been
	// created. But during the did/project creation, we also need the pubkeys,
	// which we cannot get because the did/project does not even exist yet.
	// For this purpose, a special didPubKeyGetter and projectPubkeyGetter were
	// created, which get the pubkey from the did/project creation msg itself,
	// given that the pubkey is one of the parameters in such messages.
	//
	// - did module msgs are signed by did module DIDs
	// - project module msgs are signed by project module DIDs (a.k.a projects)
	// - [[default]] remaining ixo module msgs are signed by did module DIDs
	//
	// A special case in the project module is the MsgWithdrawFunds message,
	// which is a project module message signed by a did module DID (instead
	// of a project DID). The project module PubKeyGetter deals with this
	// inconsistency by using the did module pubkey getter for MsgWithdrawFunds.

	//TODO uncomment
	defaultPubKeyGetter := did.NewDefaultPubKeyGetter(app.didKeeper)
	didPubKeyGetter := did.NewModulePubKeyGetter(app.didKeeper)
	//TODO uncomment ixo module
	projectPubKeyGetter := project.NewModulePubKeyGetter(app.projectKeeper, app.didKeeper)

	// Since we have parameterised pubkey getters, we can use the same default
	// ixo AnteHandler (ixo.NewDefaultAnteHandler) for all three pubkey getters
	// instead of having to implement three unique AnteHandlers.

	//TODO uncomment
	defaultIxoAnteHandler := ixotypes.NewDefaultAnteHandler(
		app.AccountKeeper, app.BankKeeper, ixotypes.IxoSigVerificationGasConsumer,
		defaultPubKeyGetter, encodingConfig.TxConfig.SignModeHandler())
	didAnteHandler := ixotypes.NewDefaultAnteHandler(
		app.AccountKeeper, app.BankKeeper, ixotypes.IxoSigVerificationGasConsumer,
		didPubKeyGetter, encodingConfig.TxConfig.SignModeHandler())
	//TODO uncomment ixo module
	projectAnteHandler := ixotypes.NewDefaultAnteHandler(
		app.AccountKeeper, app.BankKeeper, ixotypes.IxoSigVerificationGasConsumer,
		projectPubKeyGetter, encodingConfig.TxConfig.SignModeHandler())

	//defaultIxoAnteHandler := ixotypes.NewDefaultAnteHandler(
	//	app.AccountKeeper, app.BankKeeper, ixotypes.IxoSigVerificationGasConsumer,
	//	defaultPubKeyGetter, encodingConfig.TxConfig.SignModeHandler())
	//didAnteHandler := ixo.NewDefaultAnteHandler(
	//	app.AccountKeeper, app.SupplyKeeper, ixo.IxoSigVerificationGasConsumer, didPubKeyGetter)
	//projectAnteHandler := ixo.NewDefaultAnteHandler(
	//	app.AccountKeeper, app.SupplyKeeper, ixo.IxoSigVerificationGasConsumer, projectPubKeyGetter)

	// The default Cosmos AnteHandler is still used for standard Cosmos messages
	// implemented in standard Cosmos modules (bank, gov, etc.). The only change
	// is that we use an IxoSigVerificationGasConsumer instead of the default
	// one, since the default does not allow ED25519 signatures. Thus, like this
	// we enable ED25519 (as well as Secp) signing of standard Cosmos messages.

	cosmosAnteHandler := authante.NewAnteHandler(
		app.AccountKeeper, app.BankKeeper, ixotypes.IxoSigVerificationGasConsumer, encodingConfig.TxConfig.SignModeHandler())
			//authante.NewAnteHandler(
		//app.AccountKeeper, app.SupplyKeeper, ixo.IxoSigVerificationGasConsumer)

	// In the case of project creation, besides having a custom pubkey getter,
	// we also have to use a custom project creation AnteHandler. Recall that
	// one of the purposes of getting the pubkey is to charge gas fees. So we
	// expect the signer to have enough tokens to pay for gas fees. Typically,
	// these are sent to the signer before the signer signs their first message.
	//
	// However, in the case of a project, we cannot send tokens to it before its
	// creation since we do not know the project DID (and thus where to send the
	// tokens) until exactly before its creation (when project creation is done
	// through ixo-cellnode). The project however does have an original creator!
	//
	// Thus, the gas fees in the case of project creation are instead charged
	// to the original creator, which is pointed out in the project doc. For
	// this purpose, a custom project creation AnteHandler had to be created.

	//TODO uncomment ixo module
	projectCreationAnteHandler := project.NewProjectCreationAnteHandler(
		app.AccountKeeper, app.BankKeeper, app.didKeeper, encodingConfig.TxConfig.SignModeHandler(),
		projectPubKeyGetter)

	return func(ctx sdk.Context, tx sdk.Tx, simulate bool) (_ sdk.Context, err error) {
		// Route message based on ixo module router key
		// Otherwise, route to Cosmos ante handler
		msg := tx.GetMsgs()[0]
		switch msg.Route() {
		case did.RouterKey:
			return didAnteHandler(ctx, tx, simulate)
		//TODO uncomment rest of ixo modules
		case projecttypes.RouterKey:
			switch msg.Type() {
			case projecttypes.TypeMsgCreateProject:
				return projectCreationAnteHandler(ctx, tx, simulate)
			default:
				return projectAnteHandler(ctx, tx, simulate)
			}
		case bondstypes.RouterKey:
			return defaultIxoAnteHandler(ctx, tx, simulate) //fallthrough
		//case treasury.RouterKey:
		//	fallthrough
		case paymentstypes.RouterKey:
			return defaultIxoAnteHandler(ctx, tx, simulate)
		default:
			return cosmosAnteHandler(ctx, tx, simulate)
		}
	}
}

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(appCodec codec.BinaryMarshaler, legacyAmino *codec.LegacyAmino, key, tkey sdk.StoreKey) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	// init params keeper and subspaces (for standard Cosmos modules)
	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(minttypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govtypes.ParamKeyTable())
	paramsKeeper.Subspace(crisistypes.ModuleName)
	paramsKeeper.Subspace(ibctransfertypes.ModuleName)
	paramsKeeper.Subspace(ibchost.ModuleName)

	// init params keeper and subspaces (for custom ixo modules
	//TODO uncomment rest of ixo modules
	paramsKeeper.Subspace(projecttypes.ModuleName)
	paramsKeeper.Subspace(bondstypes.ModuleName)

	return paramsKeeper
}
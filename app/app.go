package app

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"

	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	reflectionv1 "cosmossdk.io/api/cosmos/reflection/v1"
	"cosmossdk.io/client/v2/autocli"
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/log"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	abci "github.com/cometbft/cometbft/abci/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
	nodeservice "github.com/cosmos/cosmos-sdk/client/grpc/node"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	runtimeservices "github.com/cosmos/cosmos-sdk/runtime/services"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/server/api"
	srvconfig "github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	sigtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authcodec "github.com/cosmos/cosmos-sdk/x/auth/codec"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	txmodule "github.com/cosmos/cosmos-sdk/x/auth/tx/config"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/gorilla/mux"
	"github.com/ixofoundation/ixo-blockchain/v3/app/keepers"
	"github.com/ixofoundation/ixo-blockchain/v3/app/params"
	"github.com/ixofoundation/ixo-blockchain/v3/app/upgrades"
	v2 "github.com/ixofoundation/ixo-blockchain/v3/app/upgrades/v2"
	v3 "github.com/ixofoundation/ixo-blockchain/v3/app/upgrades/v3"
	"github.com/ixofoundation/ixo-blockchain/v3/docs"
	"github.com/ixofoundation/ixo-blockchain/v3/lib/ixo"
	"github.com/spf13/cast"
)

const (
	appName = "IxoApp"
)

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	// module account permissions
	maccPerms = moduleAccountPermissions

	// scheduled upgrades and forks
	Upgrades = []upgrades.Upgrade{v2.Upgrade, v3.Upgrade}
	Forks    = []upgrades.Fork{}

	// EmptyWasmOpts defines a type alias for a list of wasm options.
	EmptyWasmOpts []wasmkeeper.Option

	// Verify app interface at compile time
	_ runtime.AppI            = (*IxoApp)(nil)
	_ servertypes.Application = (*IxoApp)(nil)
)

// init sets DefaultNodeHome to default ixod install location.
func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, ".ixod")
}

// IxoApp extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type IxoApp struct {
	homePath       string
	invCheckPeriod uint

	*baseapp.BaseApp
	keepers.AppKeepers

	legacyAmino       *codec.LegacyAmino
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry
	txConfig          client.TxConfig

	// the module manager
	ModuleManager      *module.Manager
	BasicModuleManager module.BasicManager

	// simulation manager
	sm *module.SimulationManager

	// module configurator
	configurator module.Configurator
}

// NewIxoApp returns a reference to an initialized IxoApp.
func NewIxoApp(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	skipUpgradeHeights map[int64]bool,
	homePath string,
	appOpts servertypes.AppOptions,
	wasmOpts []wasmkeeper.Option,
	baseAppOptions ...func(*baseapp.BaseApp),
) *IxoApp {
	encodingConfig := params.MakeEncodingConfig()
	appCodec := encodingConfig.Codec
	legacyAmino := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry
	txConfig := encodingConfig.TxConfig

	std.RegisterLegacyAminoCodec(legacyAmino)
	std.RegisterInterfaces(interfaceRegistry)

	// App Opts
	skipGenesisInvariants := cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))
	invCheckPeriod := cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod))

	bApp := baseapp.NewBaseApp(appName, logger, db, txConfig.TxDecoder(), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)
	bApp.SetTxEncoder(txConfig.TxEncoder())

	app := &IxoApp{
		BaseApp:           bApp,
		legacyAmino:       legacyAmino,
		txConfig:          txConfig,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		invCheckPeriod:    invCheckPeriod,
		homePath:          homePath,
	}

	wasmDir := filepath.Join(homePath, "wasm")
	wasmConfig, err := wasm.ReadWasmConfig(appOpts)
	// Uncomment this for debugging contracts. In the future this could be made into a param passed by the tests
	// wasmConfig.ContractDebugMode = true
	if err != nil {
		panic(fmt.Sprintf("error while reading wasm config: %s", err))
	}

	app.AppKeepers = keepers.NewAppKeepers(
		appCodec,
		interfaceRegistry,
		legacyAmino,
		bApp,
		maccPerms,
		invCheckPeriod,
		skipUpgradeHeights,
		homePath,
		wasmDir,
		wasmConfig,
		wasmOpts,
		app.BlockedAddresses(),
		logger,
		appOpts,
	)
	app.SetupHooks()

	/****  Module Options ****/
	// -----------------------------------------------------
	// NOTE: All module / keeper changes should happen prior to this module.NewManager line being called.
	// However in the event any changes do need to happen after this call, ensure that that keeper
	// is only passed in its keeper form (not de-ref'd anywhere)
	//
	// Generally NewAppModule will require the keeper that module defines to be passed in as an exact struct,
	// but should take in every other keeper as long as it matches a certain interface. (So no need to be de-ref'd)
	//
	// Any time a module requires a keeper de-ref'd that's not its native one,
	// its code-smell and should probably change.
	app.ModuleManager = module.NewManager(appModules(app, appCodec, txConfig, skipGenesisInvariants)...)

	// BasicModuleManager defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration and genesis verification.
	// By default it is composed of all the module from the module manager.
	// Additionally, app module basics can be overwritten by passing them as argument.
	// Notice I have to override the gov ModuleBasic with all the custom proposal handers, otherwise we lose them in the CLI.
	app.BasicModuleManager = newBasicManagerFromManager(app)

	// optional: enable sign mode textual by overwriting the default tx config (after setting the bank keeper)
	enabledSignModes := append(authtx.DefaultSignModes, sigtypes.SignMode_SIGN_MODE_TEXTUAL)
	txConfigOpts := authtx.ConfigOptions{
		EnabledSignModes:           enabledSignModes,
		TextualCoinMetadataQueryFn: txmodule.NewBankKeeperCoinMetadataQueryFn(app.BankKeeper),
	}
	txConfig, err = authtx.NewTxConfigWithOptions(
		appCodec,
		txConfigOpts,
	)
	if err != nil {
		panic(err)
	}
	app.txConfig = txConfig

	// Upgrades from v0.50.x onwards happen in pre block
	app.ModuleManager.SetOrderPreBlockers(upgradetypes.ModuleName)
	// Tell the app's module manager how to set the order of BeginBlockers, which are run at the beginning of every block.
	app.ModuleManager.SetOrderBeginBlockers(OrderBeginBlockers()...)
	// Tell the app's module manager how to set the order of EndBlockers, which are run at the end of every block.
	app.ModuleManager.SetOrderEndBlockers(OrderEndBlockers()...)
	// Tell the app's module manager how to set the order of InitGenesis, which are run genesis initialization.
	app.ModuleManager.SetOrderInitGenesis(OrderInitGenesis()...)
	// Uncomment if want to set a custom order for export genesis
	// app.ModuleManager.SetOrderExportGenesis(order...)
	// Uncomment if you want to set a custom migration order here.
	// app.mm.SetOrderMigrations(custom order)

	app.ModuleManager.RegisterInvariants(app.CrisisKeeper)
	app.configurator = module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	err = app.ModuleManager.RegisterServices(app.configurator)
	if err != nil {
		panic(err)
	}

	autocliv1.RegisterQueryServer(app.GRPCQueryRouter(), runtimeservices.NewAutoCLIQueryService(app.ModuleManager.Modules))

	reflectionSvc := getReflectionService()
	reflectionv1.RegisterReflectionServiceServer(app.GRPCQueryRouter(), reflectionSvc)

	// initialize stores
	app.MountKVStores(app.GetKVStoreKey())
	app.MountTransientStores(app.GetTransientStoreKey())
	app.MountMemoryStores(app.GetMemoryStoreKey())

	antihandler := NewIxoAnteHandler(HandlerOptions{
		HandlerOptions: ante.HandlerOptions{
			AccountKeeper:   app.AccountKeeper,
			BankKeeper:      app.BankKeeper,
			FeegrantKeeper:  app.FeeGrantKeeper,
			SignModeHandler: txConfig.SignModeHandler(),
			SigGasConsumer:  ixo.IxoSigVerificationGasConsumer,
		},
		IidKeeper:          app.IidKeeper,
		EntityKeeper:       app.EntityKeeper,
		WasmConfig:         wasmConfig,
		IBCKeeper:          app.IBCKeeper,
		TxCounterStoreKey:  runtime.NewKVStoreService(app.GetKey(wasmtypes.StoreKey)),
		appCodec:           app.appCodec,
		smartAccountKeeper: app.SmartAccountKeeper,
	})

	posthandler := NewPostHandler(appCodec, app.SmartAccountKeeper, app.AccountKeeper, encodingConfig.TxConfig.SignModeHandler())

	// initialize BaseApp
	app.SetAnteHandler(antihandler)
	app.SetPostHandler(posthandler)
	app.SetInitChainer(app.InitChainer)
	app.SetPreBlocker(app.PreBlocker)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)
	app.SetPrecommiter(app.Precommitter)
	app.SetPrepareCheckStater(app.PrepareCheckStater)

	// must be before Loading version
	// Register snapshot extensions to enable state-sync for wasm.
	if manager := app.SnapshotManager(); manager != nil {
		err := manager.RegisterExtensions(
			wasmkeeper.NewWasmSnapshotter(app.CommitMultiStore(), &app.AppKeepers.WasmKeeper),
		)
		if err != nil {
			panic("failed to register snapshot extension: " + err.Error())
		}
	}

	// setupUpgradeHandlers is used for registering any on-chain upgrades.
	// Make sure it's called after `app.ModuleManager` and `app.configurator` are set.
	app.setupUpgradeHandlers()
	app.setupUpgradeStoreLoaders()

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			panic(fmt.Errorf("error loading last version: %w", err))
		}

		ctx := app.BaseApp.NewUncachedContext(true, cmtproto.Header{})
		// Initialize pinned codes in wasmvm as they are not persisted there
		if err := app.AppKeepers.WasmKeeper.InitializePinnedCodes(ctx); err != nil {
			panic(fmt.Errorf("failed initialize pinned codes %s", err))
		}
	}

	// create the simulation manager and define the order of the modules for deterministic simulations
	app.sm = module.NewSimulationManager(simulationModules(app, appCodec, skipGenesisInvariants)...)
	app.sm.RegisterStoreDecoders()

	return app
}

// we cache the reflectionService to save us time within tests.
var cachedReflectionService *runtimeservices.ReflectionService = nil

func getReflectionService() *runtimeservices.ReflectionService {
	if cachedReflectionService != nil {
		return cachedReflectionService
	}
	reflectionSvc, err := runtimeservices.NewReflectionService()
	if err != nil {
		panic(err)
	}
	cachedReflectionService = reflectionSvc
	return reflectionSvc
}

//------------------------------------------------------------------------------
// Implement `ixoapptypes.App` interface for IxoApp
//------------------------------------------------------------------------------

// Name returns the name of the App
func (app *IxoApp) Name() string { return app.BaseApp.Name() }

// GetBaseApp returns the base app of the application
func (app *IxoApp) GetBaseApp() *baseapp.BaseApp { return app.BaseApp }

// PreBlocker application updates before each begin block.
func (app *IxoApp) PreBlocker(ctx sdk.Context, _ *abci.RequestFinalizeBlock) (*sdk.ResponsePreBlock, error) {
	return app.ModuleManager.PreBlock(ctx)
}

// BeginBlocker application updates every begin block
func (app *IxoApp) BeginBlocker(ctx sdk.Context) (sdk.BeginBlock, error) {
	app.beginBlockForks(ctx)
	return app.ModuleManager.BeginBlock(ctx)
}

// EndBlocker application updates every end block
func (app *IxoApp) EndBlocker(ctx sdk.Context) (sdk.EndBlock, error) {
	return app.ModuleManager.EndBlock(ctx)
}

// Precommitter application updates before the commital of a block after all transactions have been delivered.
func (app *IxoApp) Precommitter(ctx sdk.Context) {
	if err := app.ModuleManager.Precommit(ctx); err != nil {
		panic(err)
	}
}

func (app *IxoApp) PrepareCheckStater(ctx sdk.Context) {
	if err := app.ModuleManager.PrepareCheckState(ctx); err != nil {
		panic(err)
	}
}

// InitChainer application update at chain initialization
func (app *IxoApp) InitChainer(ctx sdk.Context, req *abci.RequestInitChain) (*abci.ResponseInitChain, error) {
	var genesisState GenesisState
	if err := json.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}

	err := app.UpgradeKeeper.SetModuleVersionMap(ctx, app.ModuleManager.GetVersionMap())
	if err != nil {
		panic(err)
	}

	return app.ModuleManager.InitGenesis(ctx, app.appCodec, genesisState)
}

// LoadHeight loads a particular height
func (app *IxoApp) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// LegacyAmino returns IxoApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *IxoApp) LegacyAmino() *codec.LegacyAmino {
	return app.legacyAmino
}

// AppCodec returns IxoApp's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *IxoApp) AppCodec() codec.Codec {
	return app.appCodec
}

// InterfaceRegistry returns IxoApp's InterfaceRegistry
func (app *IxoApp) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// TxConfig returns SimApp's TxConfig
func (app *IxoApp) TxConfig() client.TxConfig {
	return app.txConfig
}

// SimulationManager implements the SimulationApp interface
func (app *IxoApp) SimulationManager() *module.SimulationManager {
	return app.sm
}

// AutoCliOpts returns the autocli options for the app.
func (app *IxoApp) AutoCliOpts() autocli.AppOptions {
	modules := make(map[string]appmodule.AppModule, 0)
	for _, m := range app.ModuleManager.Modules {
		if moduleWithName, ok := m.(module.HasName); ok {
			moduleName := moduleWithName.Name()
			if appModule, ok := moduleWithName.(appmodule.AppModule); ok {
				modules[moduleName] = appModule
			}
		}
	}

	return autocli.AppOptions{
		Modules:               modules,
		ModuleOptions:         runtimeservices.ExtractAutoCLIOptions(app.ModuleManager.Modules),
		AddressCodec:          authcodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix()),
		ValidatorAddressCodec: authcodec.NewBech32Codec(sdk.GetConfig().GetBech32ValidatorAddrPrefix()),
		ConsensusAddressCodec: authcodec.NewBech32Codec(sdk.GetConfig().GetBech32ConsensusAddrPrefix()),
	}
}

// DefaultGenesis returns a default genesis from the registered AppModuleBasic's.
func (app *IxoApp) DefaultGenesis() map[string]json.RawMessage {
	return app.BasicModuleManager.DefaultGenesis(app.appCodec)
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *IxoApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig srvconfig.APIConfig) {
	clientCtx := apiSvr.ClientCtx

	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register new tendermint queries routes from grpc-gateway.
	cmtservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register node gRPC service for grpc-gateway.
	nodeservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register legacy and grpc-gateway routes for all modules.
	app.BasicModuleManager.RegisterGRPCGatewayRoutes(
		clientCtx,
		apiSvr.GRPCGatewayRouter,
	)

	// register swagger API from root so that other applications can override easily
	if err := RegisterSwaggerAPI(clientCtx, apiSvr.Router, apiConfig.Swagger); err != nil {
		panic(err)
	}
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *IxoApp) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *IxoApp) RegisterTendermintService(clientCtx client.Context) {
	cmtApp := server.NewCometABCIWrapper(app)
	cmtservice.RegisterTendermintService(
		clientCtx,
		app.BaseApp.GRPCQueryRouter(),
		app.interfaceRegistry,
		cmtApp.Query,
	)
}

// RegisterNodeService registers the node gRPC Query service.
func (app *IxoApp) RegisterNodeService(clientCtx client.Context, cfg srvconfig.Config) {
	nodeservice.RegisterNodeService(clientCtx, app.GRPCQueryRouter(), cfg)
}

func RegisterSwaggerAPI(_ client.Context, rtr *mux.Router, swaggerEnabled bool) error {
	if !swaggerEnabled {
		return nil
	}

	root, err := fs.Sub(docs.SwaggerUI, "swagger-ui")
	if err != nil {
		return err
	}

	staticServer := http.FileServer(http.FS(root))
	rtr.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", staticServer))

	return nil
}

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}
	return dupMaccPerms
}

//------------------------------------------------------------------------------
// Upgrades and forks
//------------------------------------------------------------------------------

// configure store loader that checks if version == upgradeHeight and applies store upgrades
func (app *IxoApp) setupUpgradeStoreLoaders() {
	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk: %s", err))
	}

	if app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		return
	}

	for _, upgrade := range Upgrades {
		if upgradeInfo.Name == upgrade.UpgradeName {
			app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &upgrade.StoreUpgrades))
		}
	}
}

func (app *IxoApp) setupUpgradeHandlers() {
	for _, upgrade := range Upgrades {
		app.UpgradeKeeper.SetUpgradeHandler(
			upgrade.UpgradeName,
			upgrade.CreateUpgradeHandler(app.ModuleManager, app.configurator),
		)
	}
}

// BeginBlockForks is intended to be ran in a chain upgrade.
func (app *IxoApp) beginBlockForks(ctx sdk.Context) {
	for _, fork := range Forks {
		if ctx.BlockHeight() == fork.UpgradeHeight {
			fork.BeginForkLogic(ctx)
			return
		}
	}
}

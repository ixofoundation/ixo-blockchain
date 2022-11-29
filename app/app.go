package app

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/cosmos/cosmos-sdk/x/authz"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmclient "github.com/CosmWasm/wasmd/x/wasm/client"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	srvconfig "github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authrest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/mint"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	sdkparams "github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/cosmos/ibc-go/v3/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v3/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v3/modules/core"
	ibcclient "github.com/cosmos/ibc-go/v3/modules/core/02-client"
	ibcclienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	porttypes "github.com/cosmos/ibc-go/v3/modules/core/05-port/types"
	ibchost "github.com/cosmos/ibc-go/v3/modules/core/24-host"
	ibckeeper "github.com/cosmos/ibc-go/v3/modules/core/keeper"
	"github.com/gorilla/mux"
	"github.com/ixofoundation/ixo-blockchain/app/params"
	"github.com/ixofoundation/ixo-blockchain/client/tx"
	libixo "github.com/ixofoundation/ixo-blockchain/lib/ixo"
	"github.com/ixofoundation/ixo-blockchain/x/bonds"
	bondskeeper "github.com/ixofoundation/ixo-blockchain/x/bonds/keeper"
	bondstypes "github.com/ixofoundation/ixo-blockchain/x/bonds/types"

	// this line is used by starport scaffolding # stargate/app/moduleImport
	entitymodule "github.com/ixofoundation/ixo-blockchain/x/entity"
	entityclient "github.com/ixofoundation/ixo-blockchain/x/entity/client"
	entitykeeper "github.com/ixofoundation/ixo-blockchain/x/entity/keeper"
	entitytypes "github.com/ixofoundation/ixo-blockchain/x/entity/types"

	tokenmodule "github.com/ixofoundation/ixo-blockchain/x/token"
	tokenclient "github.com/ixofoundation/ixo-blockchain/x/token/client"
	tokenkeeper "github.com/ixofoundation/ixo-blockchain/x/token/keeper"
	tokentypes "github.com/ixofoundation/ixo-blockchain/x/token/types"

	iidmodule "github.com/ixofoundation/ixo-blockchain/x/iid"
	iidmodulekeeper "github.com/ixofoundation/ixo-blockchain/x/iid/keeper"
	iidtypes "github.com/ixofoundation/ixo-blockchain/x/iid/types"
	"github.com/ixofoundation/ixo-blockchain/x/payments"
	paymentskeeper "github.com/ixofoundation/ixo-blockchain/x/payments/keeper"
	paymentstypes "github.com/ixofoundation/ixo-blockchain/x/payments/types"
	"github.com/ixofoundation/ixo-blockchain/x/project"
	projectkeeper "github.com/ixofoundation/ixo-blockchain/x/project/keeper"
	projecttypes "github.com/ixofoundation/ixo-blockchain/x/project/types"
	"github.com/rakyll/statik/fs"
	"github.com/spf13/cast"
	abci "github.com/tendermint/tendermint/abci/types"
	tmjson "github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
)

const (
	appName              = "IxoApp"
	Bech32MainPrefix     = "ixo"
	Bech32PrefixAccAddr  = Bech32MainPrefix
	Bech32PrefixAccPub   = Bech32MainPrefix + sdk.PrefixPublic
	Bech32PrefixValAddr  = Bech32MainPrefix + sdk.PrefixValidator + sdk.PrefixOperator
	Bech32PrefixValPub   = Bech32MainPrefix + sdk.PrefixValidator + sdk.PrefixOperator + sdk.PrefixPublic
	Bech32PrefixConsAddr = Bech32MainPrefix + sdk.PrefixValidator + sdk.PrefixConsensus
	Bech32PrefixConsPub  = Bech32MainPrefix + sdk.PrefixValidator + sdk.PrefixConsensus + sdk.PrefixPublic
)

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome = os.ExpandEnv("$HOME/.ixod")

	// ModuleBasics defines the module BasicManager which is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
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
		// this line is used by starport scaffolding # stargate/app/moduleBasic
		iidmodule.AppModuleBasic{},
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
			)...,
		),
		wasm.AppModuleBasic{},
		// Custom ixo modules
		bonds.AppModuleBasic{},
		payments.AppModuleBasic{},
		project.AppModuleBasic{},
		entitymodule.AppModuleBasic{},
		tokenmodule.AppModuleBasic{},
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
		ibctransfertypes.ModuleName:    {authtypes.Minter, authtypes.Burner},
		wasm.ModuleName:                {authtypes.Burner},

		// Custom ixo module accounts
		bondstypes.BondsMintBurnAccount:       {authtypes.Minter, authtypes.Burner},
		bondstypes.BatchesIntermediaryAccount: nil,
		bondstypes.BondsReserveAccount:        nil,
		paymentstypes.PayRemainderPool:        nil,
	}

	// Module accounts that are allowed to receive tokens
	allowedReceivingModAcc = map[string]bool{
		distrtypes.ModuleName: true,
	}

	// Reserved payments module ID prefixes
	paymentsReservedIdPrefixes = []string{
		projecttypes.ModuleName,
	}
)

var (
	//NodeDir      = ".ixod"
	//Bech32Prefix = "ixo"

	// If EnabledSpecificProposals is "", and this is "true", then enable all x/wasm proposals.
	// If EnabledSpecificProposals is "", and this is not "true", then disable all x/wasm proposals.
	ProposalsEnabled = "false"
	// If set to non-empty string it must be comma-separated list of values that are all a subset
	// of "EnableAllProposals" (takes precedence over ProposalsEnabled)
	// https://github.com/CosmWasm/wasmd/blob/02a54d33ff2c064f3539ae12d75d027d9c665f05/x/wasm/internal/types/proposal.go#L28-L34
	EnableSpecificProposals = ""
)

// Verify app interface at compile time
var _ simapp.App = (*IxoApp)(nil)
var _ servertypes.Application = (*IxoApp)(nil)

func GetEnabledProposals() []wasm.ProposalType {
	if EnableSpecificProposals == "" {
		if ProposalsEnabled == "true" {
			return wasm.EnableAllProposals
		}
		return wasm.DisableAllProposals
	}
	chunks := strings.Split(EnableSpecificProposals, ",")
	proposals, err := wasm.ConvertToProposals(chunks)
	if err != nil {
		panic(err)
	}
	return proposals
}

// Extended ABCI application
type IxoApp struct {
	*baseapp.BaseApp  `json:"_bam_base_app,omitempty"`
	legacyAmino       *codec.LegacyAmino      `json:"legacy_amino,omitempty"`
	appCodec          codec.Codec             `json:"app_codec,omitempty"`
	interfaceRegistry types.InterfaceRegistry `json:"interface_registry,omitempty"`

	invCheckPeriod uint `json:"inv_check_period,omitempty"`

	// keys to access the substores
	keys    map[string]*sdk.KVStoreKey        `json:"keys,omitempty"`
	tkeys   map[string]*sdk.TransientStoreKey `json:"tkeys,omitempty"`
	memKeys map[string]*sdk.MemoryStoreKey    `json:"mem_keys,omitempty"`

	// keepers
	AuthzKeeper      authzkeeper.Keeper       `json:"authz_keeper"`
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
	FeeGrantKeeper   feegrantkeeper.Keeper    `json:"feegrant_keeper"`
	WasmKeeper       wasm.Keeper              `json:"wasm_keeper"`

	// make scoped keepers public for test purposes
	ScopedIBCKeeper      capabilitykeeper.ScopedKeeper `json:"scoped_ibc_keeper"`
	ScopedTransferKeeper capabilitykeeper.ScopedKeeper `json:"scoped_transfer_keeper"`
	scopedWasmKeeper     capabilitykeeper.ScopedKeeper

	// Custom ixo keepers
	// this line is used by starport scaffolding # stargate/app/keeperDeclaration

	IidKeeper      iidmodulekeeper.Keeper `json:"iid_keeper"`
	EntityKeeper   entitykeeper.Keeper    `json:"entity_keeper"`
	TokenKeeper    tokenkeeper.Keeper     `json:"token_keeper"`
	BondsKeeper    bondskeeper.Keeper     `json:"bonds_keeper"`
	PaymentsKeeper paymentskeeper.Keeper  `json:"payments_keeper,omitempty"`
	ProjectKeeper  projectkeeper.Keeper   `json:"project_keeper"`

	// the module manager
	mm *module.Manager `json:"mm,omitempty"`

	// simulation manager
	sm *module.SimulationManager `json:"sm,omitempty"`
}

// NewIxoApp returns a reference to an initialized IxoApp.
func NewIxoApp(
	logger log.Logger, db dbm.DB, traceStore io.Writer, loadLatest bool, skipUpgradeHeights map[int64]bool,
	homePath string, invCheckPeriod uint, encodingConfig params.EncodingConfig, enabledProposals []wasm.ProposalType,
	appOpts servertypes.AppOptions, wasmOpts []wasm.Option, baseAppOptions ...func(*baseapp.BaseApp),
) *IxoApp {

	appCodec := encodingConfig.Marshaler
	legacyAmino := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	bApp := baseapp.NewBaseApp(appName, logger, db, encodingConfig.TxConfig.TxDecoder(), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)

	keys := sdk.NewKVStoreKeys(
		// Standard Cosmos store keys
		authtypes.StoreKey, banktypes.StoreKey, stakingtypes.StoreKey,
		minttypes.StoreKey, distrtypes.StoreKey, slashingtypes.StoreKey,
		govtypes.StoreKey, paramstypes.StoreKey, ibchost.StoreKey, upgradetypes.StoreKey,
		evidencetypes.StoreKey, ibctransfertypes.StoreKey, capabilitytypes.StoreKey,
		authzkeeper.StoreKey, feegrant.StoreKey,
		// this line is used by starport scaffolding # stargate/app/storeKey
		wasm.StoreKey,
		iidtypes.StoreKey,
		// Custom ixo store keys
		bondstypes.StoreKey,
		paymentstypes.StoreKey, projecttypes.StoreKey,
		entitytypes.StoreKey,
		tokentypes.StoreKey,
	)
	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	app := &IxoApp{
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
	bApp.SetParamStore(app.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramskeeper.ConsensusParamsKeyTable()))

	// add capability keeper and ScopeToModule for ibc module
	app.CapabilityKeeper = capabilitykeeper.NewKeeper(appCodec, keys[capabilitytypes.StoreKey], memKeys[capabilitytypes.MemStoreKey])
	scopedIBCKeeper := app.CapabilityKeeper.ScopeToModule(ibchost.ModuleName)
	scopedTransferKeeper := app.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	scopedWasmKeeper := app.CapabilityKeeper.ScopeToModule(wasm.ModuleName)

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

	app.AuthzKeeper = authzkeeper.NewKeeper(
		keys[authzkeeper.StoreKey], appCodec, app.BaseApp.MsgServiceRouter(),
	)

	app.CrisisKeeper = crisiskeeper.NewKeeper(
		app.GetSubspace(crisistypes.ModuleName), invCheckPeriod, app.BankKeeper, authtypes.FeeCollectorName,
	)

	app.FeeGrantKeeper = feegrantkeeper.NewKeeper(appCodec, keys[feegrant.StoreKey], app.AccountKeeper)
	// NewKeeper constructs an upgrade Keeper which requires the following arguments:
	// skipUpgradeHeights - map of heights to skip an upgrade
	// storeKey - a store key with which to access upgrade's store
	// cdc - the app-wide binary codec
	// homePath - root directory of the application's config
	app.UpgradeKeeper = upgradekeeper.NewKeeper(skipUpgradeHeights, keys[upgradetypes.StoreKey], appCodec, homePath, nil)

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	app.StakingKeeper = *stakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(app.DistrKeeper.Hooks(), app.SlashingKeeper.Hooks()),
	)

	// Create IBC Keeper
	app.IBCKeeper = ibckeeper.NewKeeper(
		appCodec, keys[ibchost.StoreKey], app.GetSubspace(ibchost.ModuleName), app.StakingKeeper, app.UpgradeKeeper, scopedIBCKeeper,
	)

	// Create Transfer Keepers
	app.TransferKeeper = ibctransferkeeper.NewKeeper(
		appCodec,
		keys[ibctransfertypes.StoreKey],
		app.GetSubspace(ibctransfertypes.ModuleName),
		app.IBCKeeper.ChannelKeeper,
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		scopedTransferKeeper,
	)
	transferModule := transfer.NewAppModule(app.TransferKeeper)
	transferIBCModule := transfer.NewIBCModule(app.TransferKeeper)

	app.IidKeeper = *iidmodulekeeper.NewKeeper(
		appCodec,
		keys[iidtypes.StoreKey],
		keys[iidtypes.MemStoreKey],
	)

	// this line is used by starport scaffolding # stargate/app/keeperDefinition

	// create evidence keeper with router
	evidenceKeeper := evidencekeeper.NewKeeper(
		appCodec, keys[evidencetypes.StoreKey], &app.StakingKeeper, app.SlashingKeeper,
	)
	// If evidence needs to be handled for the app, set routes in router here and seal
	app.EvidenceKeeper = *evidenceKeeper

	wasmDir := filepath.Join(homePath, "wasm")
	wasmConfig, err := wasm.ReadWasmConfig(appOpts)
	if err != nil {
		panic(fmt.Sprintf("error while reading wasm config: %s", err))
	}

	// The last arguments can contain custom message handlers, and custom query handlers,
	// if we want to allow any custom callbacks
	supportedFeatures := "iterator,staking,stargate"
	app.WasmKeeper = wasm.NewKeeper(
		appCodec,
		keys[wasm.StoreKey],
		app.GetSubspace(wasm.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		app.StakingKeeper,
		app.DistrKeeper,
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		scopedWasmKeeper,
		app.TransferKeeper,
		app.MsgServiceRouter(),
		app.GRPCQueryRouter(),
		wasmDir,
		wasmConfig,
		supportedFeatures,
		wasmOpts...,
	)

	//app.IBCKeeper.SetRouter(ibcRouter)

	/****  Module Options ****/

	// NOTE: we may consider parsing `appOpts` inside module constructors. For the moment
	// we prefer to be more strict in what arguments the modules expect.
	var skipGenesisInvariants = cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))

	// add keepers (for custom ixo modules)
	app.BondsKeeper = bondskeeper.NewKeeper(app.BankKeeper, app.AccountKeeper, app.StakingKeeper, app.IidKeeper,
		keys[bondstypes.StoreKey], app.GetSubspace(bondstypes.ModuleName), app.appCodec)
	app.PaymentsKeeper = paymentskeeper.NewKeeper(app.appCodec, keys[paymentstypes.StoreKey],
		app.BankKeeper, app.IidKeeper, paymentsReservedIdPrefixes)
	app.ProjectKeeper = projectkeeper.NewKeeper(app.appCodec, keys[projecttypes.StoreKey],
		app.GetSubspace(projecttypes.ModuleName), app.AccountKeeper, app.IidKeeper, app.PaymentsKeeper)

	app.EntityKeeper = entitykeeper.NewKeeper(
		appCodec,
		keys[entitytypes.StoreKey],
		keys[entitytypes.MemStoreKey],
		app.IidKeeper,
		app.WasmKeeper,
		app.AccountKeeper,
		app.GetSubspace(entitytypes.ModuleName),
	)

	app.TokenKeeper = tokenkeeper.NewKeeper(
		appCodec,
		keys[tokentypes.StoreKey],
		app.IidKeeper,
		app.WasmKeeper,
		app.AccountKeeper,
		app.AuthzKeeper,
		app.GetSubspace(tokentypes.ModuleName),
	)

	// register the proposal types
	govRouter := govtypes.NewRouter()
	govRouter.AddRoute(govtypes.RouterKey, govtypes.ProposalHandler).
		AddRoute(paramproposal.RouterKey, sdkparams.NewParamChangeProposalHandler(app.ParamsKeeper)).
		AddRoute(entitytypes.RouterKey, entitymodule.NewEntityParamChangeProposalHandler(app.EntityKeeper)).
		AddRoute(tokentypes.RouterKey, tokenmodule.NewTokenParamChangeProposalHandler(app.TokenKeeper)).
		AddRoute(distrtypes.RouterKey, distr.NewCommunityPoolSpendProposalHandler(app.DistrKeeper)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.UpgradeKeeper)).
		AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(app.IBCKeeper.ClientKeeper))
	app.GovKeeper = govkeeper.NewKeeper(
		appCodec, keys[govtypes.StoreKey], app.GetSubspace(govtypes.ModuleName), app.AccountKeeper, app.BankKeeper,
		&stakingKeeper, govRouter,
	)

	// Create static IBC router, add transfer route, then set and seal it
	ibcRouter := porttypes.NewRouter()
	ibcRouter.AddRoute(ibctransfertypes.ModuleName, transferIBCModule)
	if len(enabledProposals) != 0 {
		govRouter.AddRoute(wasm.RouterKey, wasm.NewWasmProposalHandler(app.WasmKeeper, enabledProposals))
	}
	ibcRouter.AddRoute(wasm.ModuleName, wasm.NewIBCHandler(app.WasmKeeper, app.IBCKeeper.ChannelKeeper))
	app.IBCKeeper.SetRouter(ibcRouter)

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.
	app.mm = module.NewManager(
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
		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		upgrade.NewAppModule(app.UpgradeKeeper),
		wasm.NewAppModule(appCodec, &app.WasmKeeper, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		ibc.NewAppModule(app.IBCKeeper),
		sdkparams.NewAppModule(app.ParamsKeeper),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		authzmodule.NewAppModule(appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		transferModule,

		// Custom ixo AppModules
		// this line is used by starport scaffolding # stargate/app/appModule
		iidmodule.NewAppModule(app.appCodec, app.IidKeeper, &app.WasmKeeper),
		bonds.NewAppModule(app.BondsKeeper, app.AccountKeeper),
		payments.NewAppModule(app.PaymentsKeeper, app.BankKeeper),
		project.NewAppModule(app.ProjectKeeper, app.PaymentsKeeper, app.BankKeeper),
		entitymodule.NewAppModule(app.EntityKeeper),
		tokenmodule.NewAppModule(app.TokenKeeper),
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	app.mm.SetOrderBeginBlockers(
		// Standard Cosmos modules
		upgradetypes.ModuleName, minttypes.ModuleName, distrtypes.ModuleName, slashingtypes.ModuleName,
		evidencetypes.ModuleName, stakingtypes.ModuleName, ibchost.ModuleName, banktypes.ModuleName,
		paymentstypes.ModuleName, genutiltypes.ModuleName, crisistypes.ModuleName,
		paramstypes.ModuleName, authtypes.ModuleName, capabilitytypes.ModuleName,
		govtypes.ModuleName, ibctransfertypes.ModuleName, vestingtypes.ModuleName,
		authz.ModuleName, feegrant.ModuleName, wasm.ModuleName,

		// Custom ixo modules
		projecttypes.ModuleName,
		bondstypes.ModuleName,
		iidtypes.ModuleName,
		entitytypes.ModuleName,
		tokentypes.ModuleName,
	)
	app.mm.SetOrderEndBlockers(
		// Custom ixo modules

		// Standard Cosmos modules
		crisistypes.ModuleName, govtypes.ModuleName, stakingtypes.ModuleName,
		distrtypes.ModuleName, evidencetypes.ModuleName, iidtypes.ModuleName, banktypes.ModuleName,
		upgradetypes.ModuleName, ibchost.ModuleName, paramstypes.ModuleName, authtypes.ModuleName,
		minttypes.ModuleName, projecttypes.ModuleName, genutiltypes.ModuleName, vestingtypes.ModuleName,
		capabilitytypes.ModuleName, slashingtypes.ModuleName, ibctransfertypes.ModuleName,
		authz.ModuleName, feegrant.ModuleName, wasm.ModuleName,

		entitytypes.ModuleName, tokentypes.ModuleName,
		bondstypes.ModuleName, paymentstypes.ModuleName,
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
		upgradetypes.ModuleName, paramstypes.ModuleName, vestingtypes.ModuleName, authz.ModuleName,
		feegrant.ModuleName, wasm.ModuleName,

		// Custom ixo modules
		// this line is used by starport scaffolding # stargate/app/initGenesis
		iidtypes.ModuleName, bondstypes.ModuleName,
		paymentstypes.ModuleName, projecttypes.ModuleName, wasm.ModuleName,
		tokentypes.ModuleName, entitytypes.ModuleName,
	)

	ModuleBasics.RegisterInterfaces(app.interfaceRegistry)
	app.mm.RegisterInvariants(&app.CrisisKeeper)
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), encodingConfig.Amino)
	app.mm.RegisterServices(module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter()))

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
		wasm.NewAppModule(appCodec, &app.WasmKeeper, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		authzmodule.NewAppModule(appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		evidence.NewAppModule(app.EvidenceKeeper),
		ibc.NewAppModule(app.IBCKeeper),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		transferModule,
	)

	app.sm.RegisterStoreDecoders()

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	// initialize BaseApp
	ixoAnteHandler, err := IxoAnteHandler(HandlerOptions{
		AccountKeeper:     app.AccountKeeper,
		BankKeeper:        app.BankKeeper,
		FeegrantKeeper:    app.FeeGrantKeeper,
		IidKeeper:         app.IidKeeper,
		ProjectKeeper:     app.ProjectKeeper,
		wasmConfig:        wasmConfig,
		txCounterStoreKey: keys[wasm.StoreKey],

		SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
		SigGasConsumer:  libixo.IxoSigVerificationGasConsumer,
	})
	if err != nil {
		panic(err)
	}

	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetAnteHandler(ixoAnteHandler)
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
		app.CapabilityKeeper.InitMemStore(ctx)
		app.CapabilityKeeper.Seal()
	}

	app.ScopedIBCKeeper = scopedIBCKeeper
	app.ScopedTransferKeeper = scopedTransferKeeper
	app.scopedWasmKeeper = scopedWasmKeeper

	return app
}

// MakeCodecs constructs the *std.Codec and *codec.LegacyAmino instances used by
// ixoapp. It is useful for tests and clients who do not want to construct the
// full ixoapp.
func MakeCodecs() (codec.Codec, *codec.LegacyAmino) {
	config := MakeTestEncodingConfig()
	return config.Marshaler, config.Amino
}

// Name returns the name of the App
func (app *IxoApp) Name() string { return app.BaseApp.Name() }

// BeginBlocker application updates every begin block
func (app *IxoApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker application updates every end block
func (app *IxoApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

// InitChainer application update at chain initialization
func (app *IxoApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	if err := tmjson.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}
	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

// LoadHeight loads a particular height
func (app *IxoApp) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *IxoApp) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// BlockedAddrs returns all the app's module account addresses black listed for receiving tokens.
func (app *IxoApp) BlockedAddrs() map[string]bool {
	blockedAddrs := make(map[string]bool)
	for acc := range maccPerms {
		blockedAddrs[authtypes.NewModuleAddress(acc).String()] = !allowedReceivingModAcc[acc]
	}

	return blockedAddrs
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

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *IxoApp) GetKey(storeKey string) *sdk.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *IxoApp) GetTKey(storeKey string) *sdk.TransientStoreKey {
	return app.tkeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
func (app *IxoApp) GetMemKey(storeKey string) *sdk.MemoryStoreKey {
	return app.memKeys[storeKey]
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *IxoApp) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// SimulationManager implements the SimulationApp interface
func (app *IxoApp) SimulationManager() *module.SimulationManager {
	return app.sm
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *IxoApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig srvconfig.APIConfig) {
	//panic("implement me")

	clientCtx := apiSvr.ClientCtx
	rpc.RegisterRoutes(clientCtx, apiSvr.Router)
	// Register legacy tx routes.
	authrest.RegisterTxRoutes(clientCtx, apiSvr.Router)
	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register client tx routes.
	tx.RegisterTxRoutes(clientCtx, apiSvr.Router)

	// Register legacy and grpc-gateway routes for all modules.
	ModuleBasics.RegisterRESTRoutes(clientCtx, apiSvr.Router)
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// register swagger API from root so that other applications can override easily
	if apiConfig.Swagger {
		RegisterSwaggerAPI(clientCtx, apiSvr.Router)
	}
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *IxoApp) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *IxoApp) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.interfaceRegistry)
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

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}
	return dupMaccPerms
}

type KVStoreKey struct {
	name string
}

// func NewIxoAnteHandler(app *IxoApp, encodingConfig params.EncodingConfig, wasmConfig WasmTypes.WasmConfig, key sdk.StoreKey) sdk.AnteHandler {

// 	// The AnteHandler needs to get the signer's pubkey to verify signatures,
// 	// charge gas fees (to the corresponding address), and for other purposes.
// 	//
// 	// The default Cosmos AnteHandler fetches a signer address' pubkey from the
// 	// GetPubKey() function after querying the account from the account keeper.
// 	//
// 	// In the case of ixo, since signers are DIDs rather than addresses, we get
// 	// the DID Doc containing the pubkey from the did/project module (depending
// 	// if signer is a user or a project, respectively).
// 	//
// 	// This is what PubKeyGetters achieve.
// 	//
// 	// To get a pubkey from the did/project, the did/project must have been
// 	// created. But during the did/project creation, we also need the pubkeys,
// 	// which we cannot get because the did/project does not even exist yet.
// 	// For this purpose, a special didPubKeyGetter and projectPubkeyGetter were
// 	// created, which get the pubkey from the did/project creation msg itself,
// 	// given that the pubkey is one of the parameters in such messages.
// 	//
// 	// - did module msgs are signed by did module DIDs
// 	// - project module msgs are signed by project module DIDs (a.k.a projects)
// 	// - [[default]] remaining ixo module msgs are signed by did module DIDs
// 	//
// 	// A special case in the project module is the MsgWithdrawFunds message,
// 	// which is a project module message signed by a did module DID (instead
// 	// of a project DID). The project module PubKeyGetter deals with this
// 	// inconsistency by using the did module pubkey getter for MsgWithdrawFunds.

// 	defaultPubKeyGetter := iid.NewDefaultPubKeyGetter(app.IidKeeper)
// 	iidPubKeyGetter := iid.NewModulePubKeyGetter(app.IidKeeper)
// 	projectPubKeyGetter := project.NewModulePubKeyGetter(app.ProjectKeeper, app.IidKeeper)

// 	// Since we have parameterised pubkey getters, we can use the same default
// 	// ixo AnteHandler (ixo.NewDefaultAnteHandler) for all three pubkey getters
// 	// instead of having to implement three unique AnteHandlers.

// 	defaultIxoAnteHandler := ixotypes.NewDefaultAnteHandler(
// 		app.AccountKeeper, app.BankKeeper, ixotypes.IxoSigVerificationGasConsumer,
// 		defaultPubKeyGetter, encodingConfig.TxConfig.SignModeHandler(), key, app.IBCKeeper,
// 		wasmConfig)
// 	iidAnteHandler := ixotypes.NewDefaultAnteHandler(
// 		app.AccountKeeper, app.BankKeeper, ixotypes.IxoSigVerificationGasConsumer,
// 		iidPubKeyGetter, encodingConfig.TxConfig.SignModeHandler(), key, app.IBCKeeper,
// 		wasmConfig)
// 	projectAnteHandler := ixotypes.NewDefaultAnteHandler(
// 		app.AccountKeeper, app.BankKeeper, ixotypes.IxoSigVerificationGasConsumer,
// 		projectPubKeyGetter, encodingConfig.TxConfig.SignModeHandler(), key, app.IBCKeeper,
// 		wasmConfig)

// 	// The default Cosmos AnteHandler is still used for standard Cosmos messages
// 	// implemented in standard Cosmos modules (bank, gov, etc.). The only change
// 	// is that we use an IxoSigVerificationGasConsumer instead of the default
// 	// one, since the default does not allow ED25519 signatures. Thus, like this
// 	// we enable ED25519 (as well as Secp) signing of standard Cosmos messages.

// 	options := authante.HandlerOptions{
// 		AccountKeeper:   app.AccountKeeper,
// 		BankKeeper:      app.BankKeeper,
// 		FeegrantKeeper:  app.FeeGrantKeeper,
// 		SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
// 		SigGasConsumer:  ixotypes.IxoSigVerificationGasConsumer,
// 	}

// 	cosmosAnteHandler, err := authante.NewAnteHandler(options)

// 	if err != nil {
// 		panic(sdkerrors.Wrap(err, "could not create Cosmos AnteHandler"))
// 	}

// 	// In the case of project creation, besides having a custom pubkey getter,
// 	// we also have to use a custom project creation AnteHandler. Recall that
// 	// one of the purposes of getting the pubkey is to charge gas fees. So we
// 	// expect the signer to have enough tokens to pay for gas fees. Typically,
// 	// these are sent to the signer before the signer signs their first message.
// 	//
// 	// However, in the case of a project, we cannot send tokens to it before its
// 	// creation since we do not know the project DID (and thus where to send the
// 	// tokens) until exactly before its creation (when project creation is done
// 	// through ixo-cellnode). The project however does have an original creator!
// 	//
// 	// Thus, the gas fees in the case of project creation are instead charged
// 	// to the original creator, which is pointed out in the project doc. For
// 	// this purpose, a custom project creation AnteHandler had to be created.

// 	projectCreationAnteHandler := project.NewProjectCreationAnteHandler(
// 		app.AccountKeeper, app.BankKeeper, app.IidKeeper,
// 		encodingConfig.TxConfig.SignModeHandler(), projectPubKeyGetter)

// 	// TODO: Routing https://docs.cosmos.network/v0.44/building-modules/msg-services.html#amino-legacymsgs
// 	return func(ctx sdk.Context, tx sdk.Tx, simulate bool) (_ sdk.Context, err error) {
// 		// Route message based on ixo module router key
// 		// Otherwise, route to Cosmos ante handler
// 		msg := tx.GetMsgs()[0].(legacytx.LegacyMsg)

// 		switch msg.Route() {
// 		case iidtypes.RouterKey:
// 			return iidAnteHandler(ctx, tx, simulate)
// 		case projecttypes.RouterKey:
// 			switch msg.Type() {
// 			case projecttypes.TypeMsgCreateProject:
// 				return projectCreationAnteHandler(ctx, tx, simulate)
// 			default:
// 				return projectAnteHandler(ctx, tx, simulate)
// 			}

// 		case bondstypes.RouterKey:
// 			fallthrough
// 		case paymentstypes.RouterKey:
// 			return defaultIxoAnteHandler(ctx, tx, simulate)
// 		default:
// 			return cosmosAnteHandler(ctx, tx, simulate)
// 		}
// 	}
// }

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey sdk.StoreKey) paramskeeper.Keeper {
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
	paramsKeeper.Subspace(wasm.ModuleName)
	// init params keeper and subspaces (for custom ixo modules)
	// this line is used by starport scaffolding # stargate/app/paramSubspace
	paramsKeeper.Subspace(iidtypes.ModuleName)
	paramsKeeper.Subspace(bondstypes.ModuleName)
	paramsKeeper.Subspace(projecttypes.ModuleName)
	paramsKeeper.Subspace(entitytypes.ModuleName)
	paramsKeeper.Subspace(tokentypes.ModuleName)

	return paramsKeeper
}

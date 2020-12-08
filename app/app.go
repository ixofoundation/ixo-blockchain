package app

import (
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	"github.com/ixofoundation/ixo-blockchain/x/oracles"
	"io"
	"os"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/supply"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	dbm "github.com/tendermint/tm-db"

	"github.com/ixofoundation/ixo-blockchain/x/bonds"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"github.com/ixofoundation/ixo-blockchain/x/payments"
	"github.com/ixofoundation/ixo-blockchain/x/project"
	"github.com/ixofoundation/ixo-blockchain/x/treasury"
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
	DefaultCLIHome = os.ExpandEnv("$HOME/.ixocli")

	// default home directories for ixod
	DefaultNodeHome = os.ExpandEnv("$HOME/.ixod")

	// The module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		// Standard Cosmos modules
		auth.AppModuleBasic{},
		supply.AppModuleBasic{},
		genutil.AppModuleBasic{},
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(
			paramsclient.ProposalHandler, distr.ProposalHandler, upgradeclient.ProposalHandler,
		),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},

		// Custom ixo modules
		did.AppModuleBasic{},
		payments.AppModuleBasic{},
		project.AppModuleBasic{},
		bonds.AppModuleBasic{},
		treasury.AppModuleBasic{},
		oracles.AppModuleBasic{},
	)

	// Module account permissions
	maccPerms = map[string][]string{
		// Standard Cosmos module accounts
		auth.FeeCollectorName:     nil,
		distr.ModuleName:          nil,
		mint.ModuleName:           {supply.Minter},
		staking.BondedPoolName:    {supply.Burner, supply.Staking},
		staking.NotBondedPoolName: {supply.Burner, supply.Staking},
		gov.ModuleName:            {supply.Burner},

		// Custom ixo module accounts
		bonds.BondsMintBurnAccount:       {supply.Minter, supply.Burner},
		bonds.BatchesIntermediaryAccount: nil,
		bonds.BondsReserveAccount:        nil,
		treasury.ModuleName:              {supply.Minter, supply.Burner},
		payments.PayRemainderPool:        nil,
	}

	// module accounts that are allowed to receive tokens
	allowedReceivingModAcc = map[string]bool{
		distr.ModuleName: true,
	}

	// Reserved payments module ID prefixes
	paymentsReservedIdPrefixes = []string{
		project.ModuleName,
	}
)

// MakeCodec - custom tx codec
func MakeCodec() *codec.Codec {
	var cdc = codec.New()

	// Register standard Cosmos codecs
	ModuleBasics.RegisterCodec(cdc)
	vesting.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	//codec.RegisterEvidences(cdc)

	// Register ixo codec
	ixo.RegisterCodec(cdc)

	return cdc
}

// Verify app interface at compile time
var _ simapp.App = (*ixoApp)(nil)

// Extended ABCI application
type ixoApp struct {
	*bam.BaseApp
	cdc *codec.Codec

	invCheckPeriod uint

	// keys to access the substores
	keys  map[string]*sdk.KVStoreKey
	tkeys map[string]*sdk.TransientStoreKey

	// subspaces
	subspaces map[string]params.Subspace

	// Standard Cosmos keepers
	AccountKeeper  auth.AccountKeeper
	BankKeeper     bank.Keeper
	SupplyKeeper   supply.Keeper
	StakingKeeper  staking.Keeper
	SlashingKeeper slashing.Keeper
	MintKeeper     mint.Keeper
	DistrKeeper    distr.Keeper
	GovKeeper      gov.Keeper
	CrisisKeeper   crisis.Keeper
	UpgradeKeeper  upgrade.Keeper
	ParamsKeeper   params.Keeper
	EvidenceKeeper evidence.Keeper

	// Custom ixo keepers
	didKeeper      did.Keeper
	paymentsKeeper payments.Keeper
	projectKeeper  project.Keeper
	bondsKeeper    bonds.Keeper
	oraclesKeeper  oracles.Keeper
	treasuryKeeper treasury.Keeper

	// the module manager
	mm *module.Manager

	// simulation manager
	sm *module.SimulationManager
}

// NewIxoApp returns a reference to an initialized IxoApp.
func NewIxoApp(
	logger log.Logger, db dbm.DB, traceStore io.Writer, loadLatest bool, skipUpgradeHeights map[int64]bool,
	invCheckPeriod uint, baseAppOptions ...func(*bam.BaseApp),
) *ixoApp {

	cdc := MakeCodec()

	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetAppVersion(version.Version)

	keys := sdk.NewKVStoreKeys(
		// Standard Cosmos store keys
		bam.MainStoreKey, auth.StoreKey, staking.StoreKey,
		supply.StoreKey, mint.StoreKey, distr.StoreKey, slashing.StoreKey,
		gov.StoreKey, params.StoreKey, upgrade.StoreKey, evidence.StoreKey,

		// Custom ixo store keys
		did.StoreKey, payments.StoreKey,
		project.StoreKey, bonds.StoreKey, treasury.StoreKey, oracles.StoreKey,
	)
	tkeys := sdk.NewTransientStoreKeys(params.TStoreKey)

	app := &ixoApp{
		BaseApp:        bApp,
		cdc:            cdc,
		invCheckPeriod: invCheckPeriod,
		keys:           keys,
		tkeys:          tkeys,
		subspaces:      make(map[string]params.Subspace),
	}

	// init params keeper and subspaces (for standard Cosmos modules)
	app.ParamsKeeper = params.NewKeeper(app.cdc, keys[params.StoreKey], tkeys[params.TStoreKey])
	app.subspaces[auth.ModuleName] = app.ParamsKeeper.Subspace(auth.DefaultParamspace)
	app.subspaces[bank.ModuleName] = app.ParamsKeeper.Subspace(bank.DefaultParamspace)
	app.subspaces[staking.ModuleName] = app.ParamsKeeper.Subspace(staking.DefaultParamspace)
	app.subspaces[mint.ModuleName] = app.ParamsKeeper.Subspace(mint.DefaultParamspace)
	app.subspaces[distr.ModuleName] = app.ParamsKeeper.Subspace(distr.DefaultParamspace)
	app.subspaces[slashing.ModuleName] = app.ParamsKeeper.Subspace(slashing.DefaultParamspace)
	app.subspaces[gov.ModuleName] = app.ParamsKeeper.Subspace(gov.DefaultParamspace).WithKeyTable(gov.ParamKeyTable())
	app.subspaces[crisis.ModuleName] = app.ParamsKeeper.Subspace(crisis.DefaultParamspace)
	app.subspaces[evidence.ModuleName] = app.ParamsKeeper.Subspace(evidence.DefaultParamspace)

	// init params keeper and subspaces (for custom ixo modules)
	app.subspaces[project.ModuleName] = app.ParamsKeeper.Subspace(project.DefaultParamspace)
	app.subspaces[bonds.ModuleName] = app.ParamsKeeper.Subspace(bonds.DefaultParamspace)

	// add keepers (for standard Cosmos modules)
	app.AccountKeeper = auth.NewAccountKeeper(
		app.cdc, keys[auth.StoreKey], app.subspaces[auth.ModuleName], auth.ProtoBaseAccount,
	)
	app.BankKeeper = bank.NewBaseKeeper(
		app.AccountKeeper, app.subspaces[bank.ModuleName], app.BlacklistedAccAddrs(),
	)
	app.SupplyKeeper = supply.NewKeeper(
		app.cdc, keys[supply.StoreKey], app.AccountKeeper, app.BankKeeper, maccPerms,
	)
	stakingKeeper := staking.NewKeeper(
		app.cdc, keys[staking.StoreKey], app.SupplyKeeper, app.subspaces[staking.ModuleName],
	)
	app.MintKeeper = mint.NewKeeper(
		app.cdc, keys[mint.StoreKey], app.subspaces[mint.ModuleName], &stakingKeeper,
		app.SupplyKeeper, auth.FeeCollectorName,
	)
	app.DistrKeeper = distr.NewKeeper(
		app.cdc, keys[distr.StoreKey], app.subspaces[distr.ModuleName], &stakingKeeper,
		app.SupplyKeeper, auth.FeeCollectorName, app.ModuleAccountAddrs(),
	)
	app.SlashingKeeper = slashing.NewKeeper(
		app.cdc, keys[slashing.StoreKey], &stakingKeeper, app.subspaces[slashing.ModuleName],
	)
	app.CrisisKeeper = crisis.NewKeeper(
		app.subspaces[crisis.ModuleName], invCheckPeriod, app.SupplyKeeper, auth.FeeCollectorName,
	)
	app.UpgradeKeeper = upgrade.NewKeeper(skipUpgradeHeights, keys[upgrade.StoreKey], app.cdc)

	// create evidence keeper with router
	evidenceKeeper := evidence.NewKeeper(
		app.cdc, keys[evidence.StoreKey], app.subspaces[evidence.ModuleName], &app.StakingKeeper, app.SlashingKeeper,
	)
	evidenceRouter := evidence.NewRouter()
	// TODO: Register evidence routes.
	evidenceKeeper.SetRouter(evidenceRouter)
	app.EvidenceKeeper = *evidenceKeeper

	// register the proposal types
	govRouter := gov.NewRouter()
	govRouter.AddRoute(gov.RouterKey, gov.ProposalHandler).
		AddRoute(params.RouterKey, params.NewParamChangeProposalHandler(app.ParamsKeeper)).
		AddRoute(distr.RouterKey, distr.NewCommunityPoolSpendProposalHandler(app.DistrKeeper)).
		AddRoute(upgrade.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.UpgradeKeeper))
	app.GovKeeper = gov.NewKeeper(
		app.cdc, keys[gov.StoreKey], app.subspaces[gov.ModuleName], app.SupplyKeeper,
		&stakingKeeper, govRouter,
	)

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	app.StakingKeeper = *stakingKeeper.SetHooks(
		staking.NewMultiStakingHooks(app.DistrKeeper.Hooks(), app.SlashingKeeper.Hooks()),
	)

	// add keepers (for custom ixo modules)
	app.didKeeper = did.NewKeeper(app.cdc, keys[did.StoreKey])
	app.paymentsKeeper = payments.NewKeeper(app.cdc, keys[payments.StoreKey],
		app.BankKeeper, app.didKeeper, paymentsReservedIdPrefixes)
	app.projectKeeper = project.NewKeeper(app.cdc, keys[project.StoreKey], app.subspaces[project.ModuleName],
		app.AccountKeeper, app.didKeeper, app.paymentsKeeper)
	app.bondsKeeper = bonds.NewKeeper(app.BankKeeper, app.SupplyKeeper, app.AccountKeeper,
		app.StakingKeeper, app.didKeeper, keys[bonds.StoreKey], app.subspaces[bonds.ModuleName], app.cdc)
	app.oraclesKeeper = oracles.NewKeeper(app.cdc, keys[oracles.StoreKey])
	app.treasuryKeeper = treasury.NewKeeper(app.cdc, keys[treasury.StoreKey], app.BankKeeper,
		app.oraclesKeeper, app.SupplyKeeper, app.didKeeper)

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.
	app.mm = module.NewManager(
		// Standard Cosmos appmodules
		genutil.NewAppModule(app.AccountKeeper, app.StakingKeeper, app.BaseApp.DeliverTx),
		auth.NewAppModule(app.AccountKeeper),
		bank.NewAppModule(app.BankKeeper, app.AccountKeeper),
		crisis.NewAppModule(&app.CrisisKeeper),
		supply.NewAppModule(app.SupplyKeeper, app.AccountKeeper),
		gov.NewAppModule(app.GovKeeper, app.AccountKeeper, app.SupplyKeeper),
		mint.NewAppModule(app.MintKeeper),
		slashing.NewAppModule(app.SlashingKeeper, app.AccountKeeper, app.StakingKeeper),
		distr.NewAppModule(app.DistrKeeper, app.AccountKeeper, app.SupplyKeeper, app.StakingKeeper),
		staking.NewAppModule(app.StakingKeeper, app.AccountKeeper, app.SupplyKeeper),
		upgrade.NewAppModule(app.UpgradeKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),

		// Custom ixo AppModules
		did.NewAppModule(app.didKeeper),
		payments.NewAppModule(app.paymentsKeeper, app.BankKeeper),
		project.NewAppModule(app.projectKeeper, app.paymentsKeeper, app.BankKeeper),
		bonds.NewAppModule(app.bondsKeeper, app.AccountKeeper),
		treasury.NewAppModule(app.treasuryKeeper),
		oracles.NewAppModule(app.oraclesKeeper),
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	app.mm.SetOrderBeginBlockers(
		// Standard Cosmos modules
		upgrade.ModuleName, mint.ModuleName, distr.ModuleName, slashing.ModuleName, evidence.ModuleName,
		// Custom ixo modules
		bonds.ModuleName,
	)
	app.mm.SetOrderEndBlockers(
		// Standard Cosmos modules
		crisis.ModuleName, gov.ModuleName, staking.ModuleName,
		// Custom ixo modules
		bonds.ModuleName, payments.ModuleName,
	)

	// NOTE: The genutils moodule must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	app.mm.SetOrderInitGenesis(
		// Standard Cosmos modules
		auth.ModuleName, distr.ModuleName, staking.ModuleName, bank.ModuleName,
		slashing.ModuleName, gov.ModuleName, mint.ModuleName, supply.ModuleName,
		crisis.ModuleName, genutil.ModuleName, evidence.ModuleName,
		// Custom ixo modules
		did.ModuleName, project.ModuleName, payments.ModuleName,
		bonds.ModuleName, treasury.ModuleName, oracles.ModuleName,
	)

	app.mm.RegisterInvariants(&app.CrisisKeeper)
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter())

	// create the simulation manager and define the order of the modules for deterministic simulations
	//
	// NOTE: this is not required apps that don't use the simulator for fuzz testing
	// transactions
	app.sm = module.NewSimulationManager(
		auth.NewAppModule(app.AccountKeeper),
		bank.NewAppModule(app.BankKeeper, app.AccountKeeper),
		supply.NewAppModule(app.SupplyKeeper, app.AccountKeeper),
		gov.NewAppModule(app.GovKeeper, app.AccountKeeper, app.SupplyKeeper),
		mint.NewAppModule(app.MintKeeper),
		staking.NewAppModule(app.StakingKeeper, app.AccountKeeper, app.SupplyKeeper),
		distr.NewAppModule(app.DistrKeeper, app.AccountKeeper, app.SupplyKeeper, app.StakingKeeper),
		slashing.NewAppModule(app.SlashingKeeper, app.AccountKeeper, app.StakingKeeper),
		params.NewAppModule(), // NOTE: only used for simulation to generate randomized param change proposals
	)

	app.sm.RegisterStoreDecoders()

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)

	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetAnteHandler(NewIxoAnteHandler(app))
	app.SetEndBlocker(app.EndBlocker)

	if loadLatest {
		err := app.LoadLatestVersion(app.keys[bam.MainStoreKey])
		if err != nil {
			tmos.Exit(err.Error())
		}
	}

	return app
}

// Name returns the name of the App
func (app *ixoApp) Name() string { return app.BaseApp.Name() }

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
	app.cdc.MustUnmarshalJSON(req.AppStateBytes, &genesisState)
	return app.mm.InitGenesis(ctx, genesisState)
}

// LoadHeight loads a particular height
func (app *ixoApp) LoadHeight(height int64) error {
	return app.LoadVersion(height, app.keys[bam.MainStoreKey])
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *ixoApp) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[supply.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// BlacklistedAccAddrs returns all the app's module account addresses black listed for receiving tokens.
func (app *ixoApp) BlacklistedAccAddrs() map[string]bool {
	blacklistedAddrs := make(map[string]bool)
	for acc := range maccPerms {
		blacklistedAddrs[supply.NewModuleAddress(acc).String()] = !allowedReceivingModAcc[acc]
	}

	return blacklistedAddrs
}

// Codec returns SimApp's codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *ixoApp) Codec() *codec.Codec {
	return app.cdc
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

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *ixoApp) GetSubspace(moduleName string) params.Subspace {
	return app.subspaces[moduleName]
}

// SimulationManager implements the SimulationApp interface
func (app *ixoApp) SimulationManager() *module.SimulationManager {
	return app.sm
}

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}
	return dupMaccPerms
}

func NewIxoAnteHandler(app *ixoApp) sdk.AnteHandler {

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

	defaultPubKeyGetter := did.NewDefaultPubKeyGetter(app.didKeeper)
	didPubKeyGetter := did.NewModulePubKeyGetter(app.didKeeper)
	projectPubKeyGetter := project.NewModulePubKeyGetter(app.projectKeeper, app.didKeeper)

	// Since we have parameterised pubkey getters, we can use the same default
	// ixo AnteHandler (ixo.NewDefaultAnteHandler) for all three pubkey getters
	// instead of having to implement three unique AnteHandlers.

	defaultIxoAnteHandler := ixo.NewDefaultAnteHandler(
		app.AccountKeeper, app.SupplyKeeper, ixo.IxoSigVerificationGasConsumer, defaultPubKeyGetter)
	didAnteHandler := ixo.NewDefaultAnteHandler(
		app.AccountKeeper, app.SupplyKeeper, ixo.IxoSigVerificationGasConsumer, didPubKeyGetter)
	projectAnteHandler := ixo.NewDefaultAnteHandler(
		app.AccountKeeper, app.SupplyKeeper, ixo.IxoSigVerificationGasConsumer, projectPubKeyGetter)

	// The default Cosmos AnteHandler is still used for standard Cosmos messages
	// implemented in standard Cosmos modules (bank, gov, etc.). The only change
	// is that we use an IxoSigVerificationGasConsumer instead of the default
	// one, since the default does not allow ED25519 signatures. Thus, like this
	// we enable ED25519 (as well as Secp) signing of standard Cosmos messages.

	cosmosAnteHandler := auth.NewAnteHandler(
		app.AccountKeeper, app.SupplyKeeper, ixo.IxoSigVerificationGasConsumer)

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

	projectCreationAnteHandler := project.NewProjectCreationAnteHandler(
		app.AccountKeeper, app.SupplyKeeper, app.BankKeeper,
		app.didKeeper, projectPubKeyGetter)

	return func(ctx sdk.Context, tx sdk.Tx, simulate bool) (_ sdk.Context, err error) {
		// Route message based on ixo module router key
		// Otherwise, route to Cosmos ante handler
		msg := tx.GetMsgs()[0]
		switch msg.Route() {
		case did.RouterKey:
			return didAnteHandler(ctx, tx, simulate)
		case project.RouterKey:
			switch msg.Type() {
			case project.TypeMsgCreateProject:
				return projectCreationAnteHandler(ctx, tx, simulate)
			default:
				return projectAnteHandler(ctx, tx, simulate)
			}
		case bonds.RouterKey:
			fallthrough
		case treasury.RouterKey:
			fallthrough
		case payments.RouterKey:
			return defaultIxoAnteHandler(ctx, tx, simulate)
		default:
			return cosmosAnteHandler(ctx, tx, simulate)
		}
	}
}

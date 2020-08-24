package app

import (
	"encoding/json"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
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
		genutil.AppModuleBasic{},
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(paramsclient.ProposalHandler, distr.ProposalHandler),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		supply.AppModuleBasic{},

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

	// Reserved payments module ID prefixes
	paymentsReservedIdPrefixes = []string{}
)

func MakeCodec() *codec.Codec {
	var cdc = codec.New()

	ModuleBasics.RegisterCodec(cdc)
	ixo.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	codec.RegisterEvidences(cdc)

	return cdc
}

// Extended ABCI application
type ixoApp struct {
	*bam.BaseApp
	cdc *codec.Codec

	invCheckPeriod uint

	// keys to access the substores
	keys  map[string]*sdk.KVStoreKey
	tkeys map[string]*sdk.TransientStoreKey

	// Standard Cosmos keepers
	accountKeeper  auth.AccountKeeper
	bankKeeper     bank.Keeper
	supplyKeeper   supply.Keeper
	stakingKeeper  staking.Keeper
	slashingKeeper slashing.Keeper
	mintKeeper     mint.Keeper
	distrKeeper    distr.Keeper
	govKeeper      gov.Keeper
	crisisKeeper   crisis.Keeper
	paramsKeeper   params.Keeper

	// Custom ixo keepers
	didKeeper      did.Keeper
	paymentsKeeper payments.Keeper
	projectKeeper  project.Keeper
	bondsKeeper    bonds.Keeper
	oraclesKeeper  oracles.Keeper
	treasuryKeeper treasury.Keeper

	// the module manager
	mm *module.Manager
}

// NewIxoApp returns a reference to an initialized IxoApp.
func NewIxoApp(logger log.Logger, db dbm.DB, traceStore io.Writer, loadLatest bool,
	invCheckPeriod uint, baseAppOptions ...func(*bam.BaseApp)) *ixoApp {

	cdc := MakeCodec()

	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)

	keys := sdk.NewKVStoreKeys(
		// Standard Cosmos store keys
		bam.MainStoreKey, auth.StoreKey, staking.StoreKey,
		supply.StoreKey, mint.StoreKey, distr.StoreKey, slashing.StoreKey,
		gov.StoreKey, params.StoreKey,

		// Custom ixo store keys
		did.StoreKey, payments.StoreKey,
		project.StoreKey, bonds.StoreKey, treasury.StoreKey, oracles.StoreKey,
	)

	tkeys := sdk.NewTransientStoreKeys(staking.TStoreKey, params.TStoreKey)

	app := &ixoApp{
		BaseApp:        bApp,
		cdc:            cdc,
		invCheckPeriod: invCheckPeriod,
		keys:           keys,
		tkeys:          tkeys,
	}

	// init params keeper and subspaces (for standard Cosmos modules)
	app.paramsKeeper = params.NewKeeper(app.cdc, keys[params.StoreKey], tkeys[params.TStoreKey])
	authSubspace := app.paramsKeeper.Subspace(auth.DefaultParamspace)
	bankSubspace := app.paramsKeeper.Subspace(bank.DefaultParamspace)
	stakingSubspace := app.paramsKeeper.Subspace(staking.DefaultParamspace)
	mintSubspace := app.paramsKeeper.Subspace(mint.DefaultParamspace)
	distrSubspace := app.paramsKeeper.Subspace(distr.DefaultParamspace)
	slashingSubspace := app.paramsKeeper.Subspace(slashing.DefaultParamspace)
	govSubspace := app.paramsKeeper.Subspace(gov.DefaultParamspace).WithKeyTable(gov.ParamKeyTable())
	crisisSubspace := app.paramsKeeper.Subspace(crisis.DefaultParamspace)

	// init params keeper and subspaces (for custom ixo modules)
	paymentsSubspace := app.paramsKeeper.Subspace(payments.DefaultParamspace)
	projectSubspace := app.paramsKeeper.Subspace(project.DefaultParamspace)

	// add keepers (for standard Cosmos modules)
	app.accountKeeper = auth.NewAccountKeeper(app.cdc, keys[auth.StoreKey], authSubspace, auth.ProtoBaseAccount)
	app.bankKeeper = bank.NewBaseKeeper(app.accountKeeper, bankSubspace, app.ModuleAccountAddrs())
	app.supplyKeeper = supply.NewKeeper(app.cdc, keys[supply.StoreKey], app.accountKeeper, app.bankKeeper, maccPerms)
	stakingKeeper := staking.NewKeeper(app.cdc, keys[staking.StoreKey], app.supplyKeeper, stakingSubspace)
	app.mintKeeper = mint.NewKeeper(app.cdc, keys[mint.StoreKey], mintSubspace, &stakingKeeper, app.supplyKeeper, auth.FeeCollectorName)
	app.distrKeeper = distr.NewKeeper(app.cdc, keys[distr.StoreKey], distrSubspace, &stakingKeeper,
		app.supplyKeeper, auth.FeeCollectorName, app.ModuleAccountAddrs())
	app.slashingKeeper = slashing.NewKeeper(
		app.cdc, keys[slashing.StoreKey], &stakingKeeper, slashingSubspace,
	)
	app.crisisKeeper = crisis.NewKeeper(crisisSubspace, invCheckPeriod, app.supplyKeeper, auth.FeeCollectorName)

	// register the proposal types
	govRouter := gov.NewRouter()
	govRouter.AddRoute(gov.RouterKey, gov.ProposalHandler).
		AddRoute(params.RouterKey, params.NewParamChangeProposalHandler(app.paramsKeeper)).
		AddRoute(distr.RouterKey, distr.NewCommunityPoolSpendProposalHandler(app.distrKeeper))
	app.govKeeper = gov.NewKeeper(
		app.cdc, keys[gov.StoreKey], govSubspace,
		app.supplyKeeper, &stakingKeeper, govRouter,
	)

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	app.stakingKeeper = *stakingKeeper.SetHooks(
		staking.NewMultiStakingHooks(app.distrKeeper.Hooks(), app.slashingKeeper.Hooks()),
	)

	// add keepers (for custom ixo modules)
	app.didKeeper = did.NewKeeper(app.cdc, keys[did.StoreKey])
	app.paymentsKeeper = payments.NewKeeper(app.cdc, keys[payments.StoreKey], paymentsSubspace,
		app.bankKeeper, app.didKeeper, paymentsReservedIdPrefixes)
	app.projectKeeper = project.NewKeeper(app.cdc, keys[project.StoreKey], projectSubspace,
		app.accountKeeper, app.didKeeper, app.paymentsKeeper)
	app.bondsKeeper = bonds.NewKeeper(app.bankKeeper, app.supplyKeeper, app.accountKeeper,
		app.stakingKeeper, app.didKeeper, keys[bonds.StoreKey], app.cdc)
	app.oraclesKeeper = oracles.NewKeeper(app.cdc, keys[oracles.StoreKey])
	app.treasuryKeeper = treasury.NewKeeper(app.cdc, keys[treasury.StoreKey], app.bankKeeper,
		app.oraclesKeeper, app.supplyKeeper, app.didKeeper)

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.
	app.mm = module.NewManager(
		// Standard Cosmos appmodules
		genutil.NewAppModule(app.accountKeeper, app.stakingKeeper, app.BaseApp.DeliverTx),
		auth.NewAppModule(app.accountKeeper),
		bank.NewAppModule(app.bankKeeper, app.accountKeeper),
		crisis.NewAppModule(&app.crisisKeeper),
		supply.NewAppModule(app.supplyKeeper, app.accountKeeper),
		distr.NewAppModule(app.distrKeeper, app.accountKeeper, app.supplyKeeper, app.stakingKeeper),
		gov.NewAppModule(app.govKeeper, app.accountKeeper, app.supplyKeeper),
		mint.NewAppModule(app.mintKeeper),
		slashing.NewAppModule(app.slashingKeeper, app.accountKeeper, app.stakingKeeper),
		staking.NewAppModule(app.stakingKeeper, app.accountKeeper, app.supplyKeeper),

		// Custom ixo AppModules
		did.NewAppModule(app.didKeeper),
		payments.NewAppModule(app.paymentsKeeper, app.bankKeeper),
		project.NewAppModule(app.projectKeeper, app.paymentsKeeper, app.bankKeeper),
		bonds.NewAppModule(app.bondsKeeper, app.accountKeeper),
		treasury.NewAppModule(app.treasuryKeeper),
		oracles.NewAppModule(app.oraclesKeeper),
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	app.mm.SetOrderBeginBlockers(
		// Standard Cosmos modules
		mint.ModuleName, distr.ModuleName, slashing.ModuleName,
		// Custom ixo modules
		bonds.ModuleName,
	)

	app.mm.SetOrderEndBlockers(
		// Standard Cosmos modules
		crisis.ModuleName, gov.ModuleName, staking.ModuleName,
		// Custom ixo modules
		bonds.ModuleName, payments.ModuleName,
	)

	app.mm.SetOrderInitGenesis(
		// Standard Cosmos modules
		distr.ModuleName, staking.ModuleName,
		auth.ModuleName, bank.ModuleName, slashing.ModuleName, gov.ModuleName,
		mint.ModuleName, supply.ModuleName, crisis.ModuleName, genutil.ModuleName,
		// Custom ixo modules
		did.ModuleName, project.ModuleName, payments.ModuleName,
		bonds.ModuleName, treasury.ModuleName, oracles.ModuleName,
	)

	app.mm.RegisterInvariants(&app.crisisKeeper)
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter())

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

// application updates every begin block
func (app *ixoApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// application updates every end block
func (app *ixoApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

// application update at chain initialization
func (app *ixoApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState map[string]json.RawMessage
	app.cdc.MustUnmarshalJSON(req.AppStateBytes, &genesisState)

	return app.mm.InitGenesis(ctx, genesisState)
}

// load a particular height
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

func NewIxoAnteHandler(app *ixoApp) sdk.AnteHandler {

	didPubKeyGetter := did.GetPubKeyGetter(app.didKeeper)
	projectPubKeyGetter := project.GetPubKeyGetter(app.projectKeeper, app.didKeeper)

	defaultIxoAnteHandler := ixo.NewDefaultAnteHandler(
		app.accountKeeper, app.supplyKeeper, auth.DefaultSigVerificationGasConsumer)
	didAnteHandler := ixo.NewDefaultAnteHandler(
		app.accountKeeper, app.supplyKeeper, auth.DefaultSigVerificationGasConsumer)
	projectAnteHandler := ixo.NewDefaultAnteHandler(
		app.accountKeeper, app.supplyKeeper, auth.DefaultSigVerificationGasConsumer)
	cosmosAnteHandler := auth.NewAnteHandler(
		app.accountKeeper, app.supplyKeeper, ixo.IxoSigVerificationGasConsumer)

	addDidAnteHandler := did.NewAddDidAnteHandler(app.accountKeeper, app.supplyKeeper, didPubKeyGetter)
	projectCreationAnteHandler := project.NewProjectCreationAnteHandler(
		app.accountKeeper, app.supplyKeeper, app.bankKeeper,
		app.didKeeper, projectPubKeyGetter)

	return func(ctx sdk.Context, tx sdk.Tx, simulate bool) (_ sdk.Context, err error) {
		// Route message based on ixo module router key
		// Otherwise, route to Cosmos ante handler
		msg := tx.GetMsgs()[0]
		switch msg.Route() {
		case did.RouterKey:
			switch msg.Type() {
			case did.TypeMsgAddDid:
				return addDidAnteHandler(ctx, tx, simulate)
			default:
				return didAnteHandler(ctx, tx, simulate)
			}
		case project.RouterKey:
			switch msg.Type() {
			case project.TypeMsgCreateProject:
				return projectCreationAnteHandler(ctx, tx, simulate)
			default:
				return projectAnteHandler(ctx, tx, simulate)
			}
		case bonds.RouterKey:
			return defaultIxoAnteHandler(ctx, tx, simulate)
		case treasury.RouterKey:
			return defaultIxoAnteHandler(ctx, tx, simulate)
		case payments.RouterKey:
			return defaultIxoAnteHandler(ctx, tx, simulate)
		default:
			return cosmosAnteHandler(ctx, tx, simulate)
		}
	}
}

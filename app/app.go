package app

import (
	"encoding/json"
	"io"
	"os"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/genaccounts"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsClient "github.com/cosmos/cosmos-sdk/x/params/client"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/supply"
	abci "github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/ixofoundation/ixo-cosmos/x/bonddoc"
	"github.com/ixofoundation/ixo-cosmos/x/bonds"
	"github.com/ixofoundation/ixo-cosmos/x/did"
	"github.com/ixofoundation/ixo-cosmos/x/fees"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
	"github.com/ixofoundation/ixo-cosmos/x/project"
	"github.com/ixofoundation/ixo-cosmos/x/treasury"
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
	DefaultCLIHome  = os.ExpandEnv("$HOME/.ixocli")
	DefaultNodeHome = os.ExpandEnv("$HOME/.ixod")

	ModuleBasics = module.NewBasicManager(
		genaccounts.AppModuleBasic{},
		genutil.AppModuleBasic{},
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distribution.AppModuleBasic{},
		gov.NewAppModuleBasic(paramsClient.ProposalHandler, distribution.ProposalHandler),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		supply.AppModuleBasic{},

		did.AppModuleBasic{},
		fees.AppModuleBasic{},
		project.AppModuleBasic{},
		bonddoc.AppModuleBasic{},
		bonds.AppModuleBasic{},
		treasury.AppModuleBasic{},
	)

	maccPerms = map[string][]string{
		auth.FeeCollectorName:            nil,
		distribution.ModuleName:          nil,
		mint.ModuleName:                  {supply.Minter},
		staking.BondedPoolName:           {supply.Burner, supply.Staking},
		staking.NotBondedPoolName:        {supply.Burner, supply.Staking},
		gov.ModuleName:                   {supply.Burner},
		bonds.BondsMintBurnAccount:       {supply.Minter, supply.Burner},
		bonds.BatchesIntermediaryAccount: nil,
	}
)

func MakeCodec() *codec.Codec {
	var cdc = codec.New()
	ModuleBasics.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	return cdc
}

type ixoApp struct {
	*bam.BaseApp
	cdc            *codec.Codec
	invCheckPeriod uint

	keys  map[string]*sdk.KVStoreKey
	tKeys map[string]*sdk.TransientStoreKey

	accountKeeper      auth.AccountKeeper
	bankKeeper         bank.Keeper
	supplyKeeper       supply.Keeper
	stakingKeeper      staking.Keeper
	distributionKeeper distribution.Keeper
	slashingKeeper     slashing.Keeper
	govKeeper          gov.Keeper
	mintKeeper         mint.Keeper
	crisisKeeper       crisis.Keeper
	paramsKeeper       params.Keeper

	didKeeper     did.Keeper
	feesKeeper    fees.Keeper
	projectKeeper project.Keeper
	bonddocKeeper bonddoc.Keeper
	bondsKeeper   bonds.Keeper

	mm *module.Manager
}

func NewIxoApp(logger log.Logger, db dbm.DB, traceStore io.Writer, loadLatest bool,
	invCheckPeriod uint, baseAppOptions ...func(*bam.BaseApp)) *ixoApp {

	cdc := MakeCodec()

	bApp := bam.NewBaseApp(appName, logger, db, ixo.DefaultTxDecoder(cdc), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)

	keys := sdk.NewKVStoreKeys(bam.MainStoreKey, auth.StoreKey, staking.StoreKey,
		supply.StoreKey, mint.StoreKey, distribution.StoreKey, slashing.StoreKey,
		gov.StoreKey, params.StoreKey, did.StoreKey, fees.StoreKey,
		project.StoreKey, bonds.StoreKey, bonddoc.StoreKey, treasury.StoreKey)

	tKeys := sdk.NewTransientStoreKeys(staking.TStoreKey, params.TStoreKey)

	app := &ixoApp{
		BaseApp:        bApp,
		cdc:            cdc,
		invCheckPeriod: invCheckPeriod,
		keys:           keys,
		tKeys:          tKeys,
	}

	app.paramsKeeper = params.NewKeeper(app.cdc, keys[params.StoreKey], tKeys[params.TStoreKey], params.DefaultCodespace)
	authSubspace := app.paramsKeeper.Subspace(auth.DefaultParamspace)
	bankSubspace := app.paramsKeeper.Subspace(bank.DefaultParamspace)
	stakingSubspace := app.paramsKeeper.Subspace(staking.DefaultParamspace)
	mintSubspace := app.paramsKeeper.Subspace(mint.DefaultParamspace)
	distrSubspace := app.paramsKeeper.Subspace(distribution.DefaultParamspace)
	slashingSubspace := app.paramsKeeper.Subspace(slashing.DefaultParamspace)
	govSubspace := app.paramsKeeper.Subspace(gov.DefaultParamspace)
	crisisSubspace := app.paramsKeeper.Subspace(crisis.DefaultParamspace)
	feesSubspace := app.paramsKeeper.Subspace(fees.DefaultParamspace)
	projectSubspace := app.paramsKeeper.Subspace(project.DefaultParamspace)

	app.accountKeeper = auth.NewAccountKeeper(app.cdc, keys[auth.StoreKey], authSubspace, auth.ProtoBaseAccount)
	app.bankKeeper = bank.NewBaseKeeper(app.accountKeeper, bankSubspace, bank.DefaultCodespace, app.ModuleAccountAddrs())
	app.supplyKeeper = supply.NewKeeper(app.cdc, keys[supply.StoreKey], app.accountKeeper, app.bankKeeper, maccPerms)
	stakingKeeper := staking.NewKeeper(app.cdc, keys[staking.StoreKey], tKeys[staking.TStoreKey],
		app.supplyKeeper, stakingSubspace, staking.DefaultCodespace)
	app.mintKeeper = mint.NewKeeper(app.cdc, keys[mint.StoreKey], mintSubspace, &stakingKeeper, app.supplyKeeper, auth.FeeCollectorName)
	app.distributionKeeper = distribution.NewKeeper(app.cdc, keys[distribution.StoreKey], distrSubspace, &stakingKeeper,
		app.supplyKeeper, distribution.DefaultCodespace, auth.FeeCollectorName, app.ModuleAccountAddrs())
	app.slashingKeeper = slashing.NewKeeper(app.cdc, keys[slashing.StoreKey], &stakingKeeper,
		slashingSubspace, slashing.DefaultCodespace)
	app.crisisKeeper = crisis.NewKeeper(crisisSubspace, invCheckPeriod, app.supplyKeeper, auth.FeeCollectorName)

	govRouter := gov.NewRouter()
	govRouter.AddRoute(gov.RouterKey, gov.ProposalHandler).
		AddRoute(params.RouterKey, params.NewParamChangeProposalHandler(app.paramsKeeper)).
		AddRoute(distribution.RouterKey, distribution.NewCommunityPoolSpendProposalHandler(app.distributionKeeper))
	app.govKeeper = gov.NewKeeper(app.cdc, keys[gov.StoreKey], app.paramsKeeper, govSubspace,
		app.supplyKeeper, &stakingKeeper, gov.DefaultCodespace, govRouter)

	app.stakingKeeper = *stakingKeeper.SetHooks(staking.NewMultiStakingHooks(app.distributionKeeper.Hooks(),
		app.slashingKeeper.Hooks()))

	app.didKeeper = did.NewKeeper(app.cdc, keys[did.StoreKey])
	app.feesKeeper = fees.NewKeeper(app.cdc, feesSubspace)
	app.projectKeeper = project.NewKeeper(app.cdc, keys[project.StoreKey], projectSubspace, app.accountKeeper, app.feesKeeper)
	app.bonddocKeeper = bonddoc.NewKeeper(app.cdc, keys[bonddoc.StoreKey])
	app.bondsKeeper = bonds.NewKeeper(app.bankKeeper, app.supplyKeeper, app.accountKeeper, app.stakingKeeper, keys[bonds.StoreKey], app.cdc)

	app.mm = module.NewManager(
		genaccounts.NewAppModule(app.accountKeeper),
		genutil.NewAppModule(app.accountKeeper, app.stakingKeeper, app.BaseApp.DeliverTx),
		auth.NewAppModule(app.accountKeeper),
		bank.NewAppModule(app.bankKeeper, app.accountKeeper),
		crisis.NewAppModule(&app.crisisKeeper),
		supply.NewAppModule(app.supplyKeeper, app.accountKeeper),
		distribution.NewAppModule(app.distributionKeeper, app.supplyKeeper),
		gov.NewAppModule(app.govKeeper, app.supplyKeeper),
		mint.NewAppModule(app.mintKeeper),
		slashing.NewAppModule(app.slashingKeeper, app.stakingKeeper),
		staking.NewAppModule(app.stakingKeeper, app.distributionKeeper, app.accountKeeper, app.supplyKeeper),

		did.NewAppModule(app.didKeeper),
		fees.NewAppModule(app.feesKeeper),
		project.NewAppModule(app.projectKeeper, app.feesKeeper, app.bankKeeper),
		bonddoc.NewAppModule(app.bonddocKeeper),
		bonds.NewAppModule(app.bondsKeeper, app.accountKeeper),
		treasury.NewAppModule(app.bankKeeper),
	)

	app.mm.SetOrderBeginBlockers(mint.ModuleName, distribution.ModuleName, slashing.ModuleName, bonds.ModuleName)
	app.mm.SetOrderEndBlockers(gov.ModuleName, staking.ModuleName, bonds.ModuleName)

	app.mm.SetOrderInitGenesis(genaccounts.ModuleName, distribution.ModuleName,
		staking.ModuleName, auth.ModuleName, bank.ModuleName, slashing.ModuleName,
		gov.ModuleName, mint.ModuleName, supply.ModuleName, crisis.ModuleName,
		genutil.ModuleName, did.ModuleName, project.ModuleName, fees.ModuleName,
		bonddoc.ModuleName, bonds.ModuleName, treasury.ModuleName)

	app.mm.RegisterInvariants(&app.crisisKeeper)
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter())

	app.MountKVStores(keys)
	app.MountTransientStores(tKeys)

	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetAnteHandler(NewIxoAnteHandler(app))
	app.SetEndBlocker(app.EndBlocker)

	if loadLatest {
		err := app.LoadLatestVersion(app.keys[bam.MainStoreKey])
		if err != nil {
			cmn.Exit(err.Error())
		}
	}

	return app
}

func (app *ixoApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

func (app *ixoApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

func (app *ixoApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState map[string]json.RawMessage
	app.cdc.MustUnmarshalJSON(req.AppStateBytes, &genesisState)

	return app.mm.InitGenesis(ctx, genesisState)
}

func (app *ixoApp) LoadHeight(height int64) error {
	return app.LoadVersion(height, app.keys[bam.MainStoreKey])
}

func (app *ixoApp) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[app.supplyKeeper.GetModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

func (app *ixoApp) ExportAppStateAndValidators(forZeroHeight bool, jailWhiteList []string) (appState json.RawMessage,
	validators []tmtypes.GenesisValidator, err error) {

	ctx := app.NewContext(true, abci.Header{Height: app.LastBlockHeight()})
	genState := app.mm.ExportGenesis(ctx)
	appState, err = codec.MarshalJSONIndent(app.cdc, genState)
	if err != nil {
		return nil, nil, err
	}

	validators = staking.WriteValidators(ctx, app.stakingKeeper)

	return appState, validators, nil
}

func NewIxoAnteHandler(app *ixoApp) sdk.AnteHandler {
	cosmosAnteHandler := auth.NewAnteHandler(app.accountKeeper, app.supplyKeeper, auth.DefaultSigVerificationGasConsumer)
	didAnteHandler := did.NewAnteHandler(app.didKeeper)
	projectAnteHandler := project.NewAnteHandler(app.projectKeeper, app.didKeeper)
	bonddocAnteHandler := bonddoc.NewAnteHandler(app.bonddocKeeper)
	bondsAnteHandler := bonds.NewAnteHandler(app.bondsKeeper, app.didKeeper)
	treasuryAnteHandler := treasury.NewAnteHandler(app.didKeeper)

	return func(ctx sdk.Context, tx sdk.Tx, simulate bool) (_ sdk.Context, _ sdk.Result, abort bool) {
		msg := tx.GetMsgs()[0]
		switch msg.Type() {
		case did.ModuleName:
			return didAnteHandler(ctx, tx, false)
		case project.ModuleName:
			return projectAnteHandler(ctx, tx, false)
		case bonddoc.ModuleName:
			return bonddocAnteHandler(ctx, tx, false)
		case bonds.ModuleName:
			return bondsAnteHandler(ctx, tx, false)
		case treasury.ModuleName:
			return treasuryAnteHandler(ctx, tx, false)
		default:
			return cosmosAnteHandler(ctx, tx, true)
		}
	}
}

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
	cParams "github.com/cosmos/cosmos-sdk/x/params"
	paramsClient "github.com/cosmos/cosmos-sdk/x/params/client"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/supply"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/libs/log"
	tmTypes "github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/ixofoundation/ixo-cosmos/x/bonds"
	"github.com/ixofoundation/ixo-cosmos/x/contracts"
	"github.com/ixofoundation/ixo-cosmos/x/did"
	"github.com/ixofoundation/ixo-cosmos/x/fees"
	"github.com/ixofoundation/ixo-cosmos/x/fiat"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
	"github.com/ixofoundation/ixo-cosmos/x/node"
	"github.com/ixofoundation/ixo-cosmos/x/params"
	"github.com/ixofoundation/ixo-cosmos/x/project"
)

const (
	appName = "ixoApp"
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
		cParams.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		supply.AppModuleBasic{},

		contracts.AppModuleBasic{},
		did.AppModuleBasic{},
		fees.AppModuleBasic{},
		node.AppModuleBasic{},
		params.AppModuleBasic{},
		project.AppModuleBasic{},
		bonds.AppModuleBasic{},
		fiat.AppModuleBasic{},
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
	cParamsKeeper      cParams.Keeper

	contractKeeper contracts.Keeper
	didKeeper      did.Keeper
	feesKeeper     fees.Keeper
	nodeKeeper     node.Keeper
	paramsKeepr    params.Keeper
	projectKeeper  project.Keeper
	bondsKeeper    bonds.Keeper
	fiatKeeper     fiat.Keeper

	mm        *module.Manager
	ethClient ixo.EthClient
}

func NewIxoApp(logger log.Logger, db dbm.DB, traceStore io.Writer, loadLatest bool,
	invCheckPeriod uint, baseAppOptions ...func(*bam.BaseApp)) *ixoApp {

	cdc := MakeCodec()

	bApp := bam.NewBaseApp(appName, logger, db, ixo.DefaultTxDecoder(cdc), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)

	keys := sdk.NewKVStoreKeys(bam.MainStoreKey, auth.StoreKey, staking.StoreKey,
		supply.StoreKey, mint.StoreKey, distribution.StoreKey, slashing.StoreKey,
		gov.StoreKey, cParams.StoreKey, contracts.StoreKey, did.StoreKey, fees.StoreKey,
		node.StoreKey, params.StoreKey, project.StoreKey, bonds.StoreKey, fiat.StoreKey)

	tKeys := sdk.NewTransientStoreKeys(staking.TStoreKey, cParams.TStoreKey)

	app := &ixoApp{
		BaseApp:        bApp,
		cdc:            cdc,
		invCheckPeriod: invCheckPeriod,
		keys:           keys,
		tKeys:          tKeys,
	}

	app.cParamsKeeper = cParams.NewKeeper(app.cdc, keys[cParams.StoreKey], tKeys[cParams.TStoreKey], cParams.DefaultCodespace)
	authSubspace := app.cParamsKeeper.Subspace(auth.DefaultParamspace)
	bankSubspace := app.cParamsKeeper.Subspace(bank.DefaultParamspace)
	stakingSubspace := app.cParamsKeeper.Subspace(staking.DefaultParamspace)
	mintSubspace := app.cParamsKeeper.Subspace(mint.DefaultParamspace)
	distrSubspace := app.cParamsKeeper.Subspace(distribution.DefaultParamspace)
	slashingSubspace := app.cParamsKeeper.Subspace(slashing.DefaultParamspace)
	govSubspace := app.cParamsKeeper.Subspace(gov.DefaultParamspace)
	crisisSubspace := app.cParamsKeeper.Subspace(crisis.DefaultParamspace)

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
		AddRoute(cParams.RouterKey, cParams.NewParamChangeProposalHandler(app.cParamsKeeper)).
		AddRoute(distribution.RouterKey, distribution.NewCommunityPoolSpendProposalHandler(app.distributionKeeper))
	app.govKeeper = gov.NewKeeper(app.cdc, keys[gov.StoreKey], app.cParamsKeeper, govSubspace,
		app.supplyKeeper, &stakingKeeper, gov.DefaultCodespace, govRouter)

	app.stakingKeeper = *stakingKeeper.SetHooks(staking.NewMultiStakingHooks(app.distributionKeeper.Hooks(),
		app.slashingKeeper.Hooks()))

	app.didKeeper = did.NewKeeper(app.cdc, keys[did.StoreKey])
	app.paramsKeepr = params.NewKeeper(app.cdc, keys[params.StoreKey])
	app.feesKeeper = fees.NewKeeper(app.cdc, app.paramsKeepr)
	app.projectKeeper = project.NewKeeper(app.cdc, keys[project.StoreKey], app.accountKeeper, app.feesKeeper)
	app.nodeKeeper = node.NewKeeper(app.cdc, app.paramsKeepr)
	app.contractKeeper = contracts.NewKeeper(app.cdc, app.paramsKeepr)
	app.bondsKeeper = bonds.NewKeeper(app.bankKeeper, app.supplyKeeper, app.accountKeeper, app.stakingKeeper, keys[bonds.StoreKey], app.cdc)
	app.fiatKeeper = fiat.NewKeeper(app.cdc, keys[fiat.StoreKey], app.accountKeeper)

	newEthClient, cErr := ixo.NewEthClient(app.contractKeeper)
	if cErr != nil {
		panic(cErr)
	}

	app.ethClient = newEthClient

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

		contracts.NewAppModule(app.contractKeeper),
		did.NewAppModule(app.didKeeper),
		fees.NewAppModule(app.feesKeeper),
		node.NewAppModule(app.nodeKeeper),
		params.NewAppModule(app.paramsKeepr),
		fiat.NewAppModule(app.fiatKeeper),
		project.NewAppModule(app.projectKeeper, app.feesKeeper,
			app.contractKeeper, app.bankKeeper, app.paramsKeepr, app.ethClient),
		bonds.NewAppModule(app.bondsKeeper, app.accountKeeper),
	)

	app.mm.SetOrderBeginBlockers(mint.ModuleName, distribution.ModuleName, slashing.ModuleName, bonds.ModuleName)
	app.mm.SetOrderEndBlockers(gov.ModuleName, staking.ModuleName, bonds.ModuleName)

	app.mm.SetOrderInitGenesis(genaccounts.ModuleName, distribution.ModuleName,
		staking.ModuleName, auth.ModuleName, bank.ModuleName, slashing.ModuleName,
		gov.ModuleName, mint.ModuleName, supply.ModuleName, crisis.ModuleName,
		genutil.ModuleName, did.ModuleName, project.ModuleName, fees.ModuleName,
		contracts.ModuleName, node.ModuleName, params.ModuleName, bonds.ModuleName,
		fiat.ModuleName)

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

func (app *ixoApp) BeginBlocker(ctx sdk.Context, req abciTypes.RequestBeginBlock) abciTypes.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

func (app *ixoApp) EndBlocker(ctx sdk.Context, req abciTypes.RequestEndBlock) abciTypes.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

func (app *ixoApp) InitChainer(ctx sdk.Context, req abciTypes.RequestInitChain) abciTypes.ResponseInitChain {
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
	validators []tmTypes.GenesisValidator, err error) {

	ctx := app.NewContext(true, abciTypes.Header{Height: app.LastBlockHeight()})
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

	return func(ctx sdk.Context, tx sdk.Tx, simulate bool) (_ sdk.Context, _ sdk.Result, abort bool) {
		msg := tx.GetMsgs()[0]
		switch msg.Type() {
		case "did":
			return didAnteHandler(ctx, tx, false)
		case "project":
			return projectAnteHandler(ctx, tx, false)
		default:
			return cosmosAnteHandler(ctx, tx, true)
		}
	}
}

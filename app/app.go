package app

import (
	"bytes"
	"fmt"
	"io"
	"os"

	//	"cosmos-test/types"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/go-kit/kit/log/term"
	abci "github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/ibc"

	"github.com/ixofoundation/ixo-cosmos/types"
	"github.com/ixofoundation/ixo-cosmos/x/did"
	"github.com/ixofoundation/ixo-cosmos/x/fees"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
	"github.com/ixofoundation/ixo-cosmos/x/params"
	"github.com/ixofoundation/ixo-cosmos/x/project"

	tmtypes "github.com/tendermint/tendermint/types"
)

const (
	APP_NAME = "ixoApp"
)

// Extended ABCI application
type IxoApp struct {
	*bam.BaseApp
	cdc *wire.Codec

	// keys to access the substores
	keyMain    *sdk.KVStoreKey
	keyAccount *sdk.KVStoreKey
	keyIBC     *sdk.KVStoreKey
	keyStake   *sdk.KVStoreKey
	keyParams  *sdk.KVStoreKey
	keyDID     *sdk.KVStoreKey
	keyProject *sdk.KVStoreKey

	// Manage getting and setting accounts
	accountMapper auth.AccountMapper

	// Manage keeper
	coinKeeper          bank.Keeper
	feeCollectionKeeper auth.FeeCollectionKeeper
	ibcMapper           ibc.Mapper
	didKeeper           did.Keeper
	projectKeeper       project.Keeper
	paramsKeeper        params.Keeper
	feeKeeper           fees.Keeper
	ixoKeeper           ixo.Keeper

	// Eth client
	ethClient ixo.EthClient
}

func NewIxoApp(logger log.Logger, db dbm.DB, traceStore io.Writer, baseAppOptions ...func(*bam.BaseApp)) *IxoApp {

	colorFn := func(keyvals ...interface{}) term.FgBgColor {
		return term.FgBgColor{Fg: term.Blue}
	}
	newLogger := log.NewTMLoggerWithColorFn(os.Stdout, colorFn)
	newLogger = log.Logger.With(newLogger, "module", "ixo")

	cdc := MakeCodec()

	bApp := bam.NewBaseApp(APP_NAME, newLogger, db, nil, baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)

	ethClient, cErr := ixo.NewEthClient()
	if cErr != nil {
		logger.Error(cErr.Error())
		//		panic(cErr)
	}

	// t, _ := ethClient.GetTransactionByHash("0xfd5e66b11abfdaa0a1bee7048a9da0b14ffeaee36c5cc897e007a00f23f23b95")
	// fmt.Println(t.Result.Input)
	// ethClient.IsProjectFundingTx("", t)
	// ethClient.GetEthProjectWallet("123")

	// create your application object
	var app = &IxoApp{
		cdc:        cdc,
		BaseApp:    bApp,
		keyMain:    sdk.NewKVStoreKey("main"),
		keyAccount: sdk.NewKVStoreKey("acc"),
		keyIBC:     sdk.NewKVStoreKey("ibc"),
		keyStake:   sdk.NewKVStoreKey("stake"),
		keyParams:  sdk.NewKVStoreKey("params"),
		keyDID:     sdk.NewKVStoreKey("did"),
		keyProject: sdk.NewKVStoreKey("project"),

		ethClient: ethClient,
	}

	// define and attach the mappers and keepers
	app.accountMapper = auth.NewAccountMapper(
		cdc,
		app.keyAccount, // target store
		func() auth.Account {
			return &types.AppAccount{}
		},
	)

	// add handlers
	app.coinKeeper = bank.NewKeeper(app.accountMapper)
	app.ibcMapper = ibc.NewMapper(app.cdc, app.keyIBC, app.RegisterCodespace(ibc.DefaultCodespace))
	app.paramsKeeper = params.NewKeeper(app.cdc, app.keyParams)
	app.feeKeeper = fees.NewKeeper(app.cdc, app.paramsKeeper)
	app.didKeeper = did.NewKeeper(app.cdc, app.keyDID)
	app.ixoKeeper = ixo.NewKeeper(app.cdc, app.paramsKeeper)
	app.projectKeeper = project.NewKeeper(app.cdc, app.keyProject, app.accountMapper, app.feeKeeper)

	app.Router().
		AddRoute("bank", bank.NewHandler(app.coinKeeper)).
		//		AddRoute("pool", pool.NewHandler(app.poolKeeper)).
		AddRoute("ibc", ibc.NewHandler(app.ibcMapper, app.coinKeeper)).
		AddRoute("did", did.NewHandler(app.didKeeper)).
		AddRoute("project", project.NewHandler(app.projectKeeper, app.coinKeeper, app.ethClient))

	// initialize BaseApp
	app.SetInitChainer(app.initChainerFn(app.didKeeper, app.projectKeeper))
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)
	app.SetTxDecoder(app.txDecoder)
	app.SetAnteHandler(NewIxoAnteHandler(app, app.feeCollectionKeeper))
	app.MountStoresIAVL(app.keyMain, app.keyAccount, app.keyIBC, app.keyStake, app.keyParams, app.keyDID, app.keyProject)
	err := app.LoadLatestVersion(app.keyMain)
	if err != nil {
		cmn.Exit(err.Error())
	}

	return app
}

func MakeCodec() *wire.Codec {

	cdc := wire.NewCodec()

	wire.RegisterCrypto(cdc)
	sdk.RegisterWire(cdc)
	bank.RegisterWire(cdc)
	ibc.RegisterWire(cdc)
	auth.RegisterWire(cdc)
	did.RegisterWire(cdc)
	project.RegisterWire(cdc)

	// register custom type
	cdc.RegisterConcrete(&types.AppAccount{}, "basecoin/Account", nil)

	cdc.Seal()

	return cdc
}

// custom logic for transaction decoding
func (app *IxoApp) txDecoder(txBytes []byte) (sdk.Tx, sdk.Error) {

	if len(txBytes) == 0 {
		return nil, sdk.ErrTxDecode("txBytes are empty")
	}

	//Check if bytes start with a curly bracket
	txByteString := string(txBytes[0:1])
	if txByteString == "{" {
		var tx = ixo.IxoTx{}

		app.Logger.Info("********DECODED_TXN*********")
		app.Logger.Info(string(txBytes))
		// Lets replace the hex encoded Msg with it's unhexed json equivalent so it can be parsed correctly
		var upTx map[string]interface{}
		json.Unmarshal(txBytes, &upTx)
		payloadArray := upTx["payload"].([]interface{})
		if len(payloadArray) != 1 {
			return nil, sdk.ErrTxDecode("Multiple messages not supported")
		}

		// Parse out the signed bytes
		signByteString := getSignBytes(txBytes)
		// fmt.Println("******** SignBytes *********")
		// fmt.Println(signByteString)

		// Add them back to the message
		msgPayload := payloadArray[0].(map[string]interface{})
		msg := msgPayload["value"].(map[string]interface{})
		msg["signBytes"] = signByteString

		// Repack the message
		txBytes, _ = json.Marshal(upTx)

		// StdTx.Msg is an interface. The concrete types
		// are registered by MakeTxCodec in bank.RegisterWire.
		err := app.cdc.UnmarshalJSON(txBytes, &tx)
		if err != nil {
			return nil, sdk.ErrTxDecode("").TraceSDK(err.Error())
		}

		// fmt.Println("TXN_PAYLOAD", tx)

		return tx, nil

	} else {
		var tx = auth.StdTx{}

		// StdTx.Msg is an interface. The concrete types
		// are registered by MakeTxCodec in bank.RegisterWire.
		err := app.cdc.UnmarshalBinary(txBytes, &tx)
		if err != nil {
			return nil, sdk.ErrTxDecode("").TraceSDK(err.Error())
		}
		fmt.Println(tx)
		return tx, nil

	}
}

func getSignBytes(txBytes []byte) string {
	const strtTxt string = "\"value\":"
	const endTxt string = "}],\"signatures\":"

	strt := bytes.Index(txBytes, []byte(strtTxt)) + len(strtTxt)
	end := bytes.Index(txBytes, []byte(endTxt))

	signBytes := txBytes[strt:end]
	return string(signBytes)
}

// BeginBlocker reflects logic to run before any TXs application are processed
// by the application.
func (app *IxoApp) BeginBlocker(_ sdk.Context, _ abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return abci.ResponseBeginBlock{}
}

// EndBlocker reflects logic to run after all TXs are processed by the
// application.
func (app *IxoApp) EndBlocker(_ sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return abci.ResponseEndBlock{}
}

// initChainer implements the custom application logic that the BaseApp will
// invoke upon initialization. In this case, it will take the application's
// state provided by 'req' and attempt to deserialize said state. The state
// should contain all the genesis accounts. These accounts will be added to the
// application's account mapper.
func (app *IxoApp) initChainerFn(didKeeper did.Keeper, projectKeeper project.Keeper) sdk.InitChainer {

	return func(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
		app.Logger.Info("******Init Chain called")
		stateJSON := req.AppStateBytes

		genesisState := new(GenesisState)
		err := app.cdc.UnmarshalJSON(stateJSON, genesisState)
		if err != nil {
			panic(err)
		}

		for _, gacc := range genesisState.Accounts {
			acc, err := gacc.ToAppAccount()
			if err != nil {
				panic(err)
			}
			app.accountMapper.SetAccount(ctx, acc)
		}
		app.Logger.Info("******Init Fees")

		// load the initial fee information
		err = fees.InitGenesis(ctx, app.feeKeeper, genesisState.FeeData)
		if err != nil {
			panic(err)
		}

		return abci.ResponseInitChain{}
	}
}

// Custom logic for state export
func (app *IxoApp) ExportAppStateAndValidators() (appState json.RawMessage, validators []tmtypes.GenesisValidator, err error) {
	ctx := app.NewContext(true, abci.Header{})

	// iterate to get the accounts
	accounts := []*types.GenesisAccount{}
	appendAccount := func(acc auth.Account) (stop bool) {
		account := &types.GenesisAccount{
			Address: acc.GetAddress(),
			Coins:   acc.GetCoins(),
		}
		accounts = append(accounts, account)
		return false
	}
	app.accountMapper.IterateAccounts(ctx, appendAccount)

	genState := GenesisState{
		Accounts: accounts,
		FeeData:  fees.WriteGenesis(ctx, app.feeKeeper),
	}

	appState, err = wire.MarshalJSONIndent(app.cdc, genState)
	if err != nil {
		return nil, nil, err
	}
	return appState, validators, nil
}

// custom logic for ixo initialization
// NewIxoAnteHandler returns an AnteHandler that wrappers
// the default cosmos one for signature checking. Based on
// the message type it either checks the Sovrin signature
// or executes the defualt cosmos version
func NewIxoAnteHandler(app *IxoApp, fck auth.FeeCollectionKeeper) sdk.AnteHandler {
	cosmosAnteHandler := auth.NewAnteHandler(app.accountMapper, fck)
	didAnteHandler := did.NewAnteHandler(app.didKeeper)
	projectAnteHandler := project.NewAnteHandler(app.projectKeeper, app.didKeeper)

	return func(
		ctx sdk.Context, tx sdk.Tx,
	) (_ sdk.Context, _ sdk.Result, abort bool) {

		msg := tx.GetMsgs()[0]

		switch msg.Type() {

		case "did":
			return didAnteHandler(ctx, tx)
		case "project":
			return projectAnteHandler(ctx, tx)
		default:
			return cosmosAnteHandler(ctx, tx)
		}
	}
}

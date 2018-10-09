package app

import (
	"bytes"
	"fmt"

	//	"cosmos-test/types"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/wire"
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
	"github.com/ixofoundation/ixo-cosmos/x/ixo"

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
}

func NewIxoApp(logger log.Logger, db dbm.DB, baseAppOptions ...func(*bam.BaseApp)) *IxoApp {

	cdc := MakeCodec()
	// create your application object
	var app = &IxoApp{
		cdc:        cdc,
		BaseApp:    bam.NewBaseApp(APP_NAME, logger, db, nil, baseAppOptions...),
		keyMain:    sdk.NewKVStoreKey("main"),
		keyAccount: sdk.NewKVStoreKey("acc"),
		keyIBC:     sdk.NewKVStoreKey("ibc"),
		keyStake:   sdk.NewKVStoreKey("stake"),
		keyDID:     sdk.NewKVStoreKey("did"),
		keyProject: sdk.NewKVStoreKey("project"),
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
	app.didKeeper = did.NewKeeper(app.cdc, app.keyDID)
	app.projectKeeper = project.NewKeeper(app.cdc, app.keyProject, app.accountMapper)

	app.Router().
		AddRoute("bank", bank.NewHandler(app.coinKeeper)).
		//		AddRoute("project", project.NewHandler()).
		AddRoute("ibc", ibc.NewHandler(app.ibcMapper, app.coinKeeper)).
		AddRoute("did", did.NewHandler(app.didKeeper)).
		AddRoute("project", project.NewHandler(app.projectKeeper, app.coinKeeper))

	// initialize BaseApp
	app.SetInitChainer(app.initChainerFn(app.didKeeper, app.projectKeeper))
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)
	app.SetTxDecoder(app.txDecoder)
	app.SetAnteHandler(NewIxoAnteHandler(app, app.feeCollectionKeeper))
	app.MountStoresIAVL(app.keyMain, app.keyAccount, app.keyIBC, app.keyStake, app.keyDID, app.keyProject)
	err := app.LoadLatestVersion(app.keyMain)
	if err != nil {
		cmn.Exit(err.Error())
	}

	return app
}

// custom tx codec
// TODO: use new go-wire
func MakeCodec() *wire.Codec {
	const msgTypeSend = 0x1
	const msgTypeIssue = 0x2
	const msgTypeQuiz = 0x3
	const msgTypeSetTrend = 0x4
	const msgTypeIBCTransferMsg = 0x5
	const msgTypeIBCReceiveMsg = 0x6
	const msgTypeBondMsg = 0x7
	const msgTypeUnbondMsg = 0x8

	const msgTypeAddDidMsg = 0xA

	const msgTypeCreateProjectMsg = 0x10
	const msgTypeUpdateProjectStatusMsg = 0x19
	const msgTypeCreateAgentMsg = 0x11
	const msgTypeUpdateAgentMsg = 0x12
	const msgTypeCreateClaimMsg = 0x13
	const msgTypeCreateEvaluationMsg = 0x14

	const msgTypeFundProjectMsg = 0x15
	const msgTypeWithdrawFundsMsg = 0x16

	const msgTypeAddCredentialMsg = 0x18

	var _ = oldwire.RegisterInterface(
		struct{ sdk.Msg }{},
		oldwire.ConcreteType{bank.SendMsg{}, msgTypeSend},
		oldwire.ConcreteType{bank.IssueMsg{}, msgTypeIssue},
		oldwire.ConcreteType{ibc.IBCTransferMsg{}, msgTypeIBCTransferMsg},
		oldwire.ConcreteType{ibc.IBCReceiveMsg{}, msgTypeIBCReceiveMsg},
		oldwire.ConcreteType{simplestake.BondMsg{}, msgTypeBondMsg},
		oldwire.ConcreteType{simplestake.UnbondMsg{}, msgTypeUnbondMsg},

		oldwire.ConcreteType{did.AddDidMsg{}, msgTypeAddDidMsg},
		oldwire.ConcreteType{did.AddCredentialMsg{}, msgTypeAddCredentialMsg},

		oldwire.ConcreteType{project.CreateProjectMsg{}, msgTypeCreateProjectMsg},
		oldwire.ConcreteType{project.UpdateProjectStatusMsg{}, msgTypeUpdateProjectStatusMsg},
		oldwire.ConcreteType{project.CreateAgentMsg{}, msgTypeCreateAgentMsg},
		oldwire.ConcreteType{project.UpdateAgentMsg{}, msgTypeUpdateAgentMsg},
		oldwire.ConcreteType{project.CreateClaimMsg{}, msgTypeCreateClaimMsg},
		oldwire.ConcreteType{project.CreateEvaluationMsg{}, msgTypeCreateEvaluationMsg},

		oldwire.ConcreteType{project.FundProjectMsg{}, msgTypeFundProjectMsg},
		oldwire.ConcreteType{project.WithdrawFundsMsg{}, msgTypeWithdrawFundsMsg},
	)

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

		fmt.Println("********DECODED_TXN*********")
		fmt.Println(string(txBytes))

		// Lets replace the hex encoded Msg with it's unhexed json equivalent so it can be parsed correctly
		var upTx map[string]interface{}
		json.Unmarshal(txBytes, &upTx)
		payloadArray := upTx["payload"].([]interface{})

		// Check if it is not json
		payloadBody, isString := payloadArray[1].(string)
		if isString {
			jsonBytes := make([]byte, hex.DecodedLen(len(payloadBody)))
			jsonBytes, hexErr := hex.DecodeString(payloadBody)
			if hexErr != nil {
				fmt.Print("Error decoding hex payload: ", payloadBody)
				return nil, sdk.ErrTxDecode("").TraceCause(hexErr, "")
			}
			jsonBytes = bytes.Replace(jsonBytes, []byte("{"), []byte("{\"signBytes\":\""+payloadBody+"\","), 1)
			txBytes = bytes.Replace(txBytes, []byte("\""+payloadBody+"\""), jsonBytes, 1)
			fmt.Println("** TXN_BYTES ", string(txBytes))
		}

		// StdTx.Msg is an interface. The concrete types
		// are registered by MakeTxCodec in bank.RegisterWire.
		err := app.cdc.UnmarshalJSON(txBytes, &tx)
		if err != nil {
			return nil, sdk.ErrTxDecode("").TraceSDK(err.Error())
		}

		fmt.Println("** TXN_PAYLOAD ", tx)

		return tx, nil
	}

	var tx = sdk.StdTx{}

	// StdTx.Msg is an interface. The concrete types
	// are registered by MakeTxCodec in bank.RegisterWire.
	err := app.cdc.UnmarshalBinary(txBytes, &tx)
	if err != nil {
		return nil, sdk.ErrTxDecode("").TraceCause(err, "")
	}
	fmt.Println(tx)
	return tx, nil

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
func (app *IxoApp) EndBlocker(_ sdk.Context, _ abci.RequestEndBlock) abci.ResponseEndBlock {
	return abci.ResponseEndBlock{}
}

// initChainer implements the custom application logic that the BaseApp will
// invoke upon initialization. In this case, it will take the application's
// state provided by 'req' and attempt to deserialize said state. The state
// should contain all the genesis accounts. These accounts will be added to the
// application's account mapper.
func (app *IxoApp) initChainerFn(didKeeper did.Keeper, projectKeeper project.Keeper) sdk.InitChainer {
	return func(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
		stateJSON := req.AppStateBytes

		genesisState := new(types.GenesisState)
		err := app.cdc.UnmarshalJSON(stateJSON, genesisState)
		if err != nil {
			panic(err) // TODO https://github.com/cosmos/cosmos-sdk/issues/468
			// return sdk.ErrGenesisParse("").TraceCause(err, "")
		}

		for _, gacc := range genesisState.Accounts {
			acc, err := gacc.ToAppAccount()
			if err != nil {
				panic(err) // TODO https://github.com/cosmos/cosmos-sdk/issues/468
				//	return sdk.ErrGenesisParse("").TraceCause(err, "")
			}
			app.accountMapper.SetAccount(ctx, acc)
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

	genState := types.GenesisState{
		Accounts: accounts,
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

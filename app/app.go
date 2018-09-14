package app

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"

	//	"cosmos-test/types"
	"encoding/json"

	abci "github.com/tendermint/abci/types"
	oldwire "github.com/tendermint/go-wire"
	cmn "github.com/tendermint/tmlibs/common"
	dbm "github.com/tendermint/tmlibs/db"
	"github.com/tendermint/tmlibs/log"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/ibc"
	"github.com/cosmos/cosmos-sdk/x/simplestake"

	"github.com/ixofoundation/ixo-cosmos/types"
	"github.com/ixofoundation/ixo-cosmos/x/did"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"

	"github.com/ixofoundation/ixo-cosmos/x/project"
)

const (
	APP_NAME = "ixoApp"
)

// Extended ABCI application
type IxoApp struct {
	*bam.BaseApp
	cdc *wire.Codec

	// keys to access the substores
	capKeyMainStore    *sdk.KVStoreKey
	capKeyAccountStore *sdk.KVStoreKey
	capKeyIBCStore     *sdk.KVStoreKey
	capKeyStakingStore *sdk.KVStoreKey
	capKeyDIDStore     *sdk.KVStoreKey
	capKeyProjectStore *sdk.KVStoreKey

	// Manage getting and setting accounts
	accountMapper sdk.AccountMapper

	// Manage getting and setting dids
	didMapper     did.SealedDidMapper
	projectMapper project.SealedProjectMapper
}

func NewIxoApp(logger log.Logger, dbs map[string]dbm.DB) *IxoApp {
	// create your application object
	var app = &IxoApp{
		BaseApp:            bam.NewBaseApp(APP_NAME, logger, dbs["main"]),
		cdc:                MakeCodec(),
		capKeyMainStore:    sdk.NewKVStoreKey("main"),
		capKeyAccountStore: sdk.NewKVStoreKey("acc"),
		capKeyIBCStore:     sdk.NewKVStoreKey("ibc"),
		capKeyStakingStore: sdk.NewKVStoreKey("stake"),
		capKeyDIDStore:     sdk.NewKVStoreKey("did"),
		capKeyProjectStore: sdk.NewKVStoreKey("project"),
	}

	// define the accountMapper
	app.accountMapper = auth.NewAccountMapperSealed(
		app.capKeyMainStore, // target store
		&types.AppAccount{}, // prototype
	)

	// define the didMapper
	app.didMapper = did.NewDidMapperSealed(
		app.capKeyDIDStore, // target store
		&did.BaseDidDoc{},  // prototype
	)

	// define the projectMapper
	app.projectMapper = project.NewProjectMapperSealed(
		app.capKeyProjectStore,      // target store
		&project.StoredProjectDoc{}, // prototype
	)

	// add handlers
	coinKeeper := bank.NewCoinKeeper(app.accountMapper)
	ibcMapper := ibc.NewIBCMapper(app.cdc, app.capKeyIBCStore)
	stakeKeeper := simplestake.NewKeeper(app.capKeyStakingStore, coinKeeper)
	didKeeper := did.NewKeeper(app.didMapper)
	projectKeeper := project.NewKeeper(app.projectMapper, app.accountMapper)
	app.Router().
		AddRoute("bank", bank.NewHandler(coinKeeper)).
		//		AddRoute("project", project.NewHandler()).
		AddRoute("ibc", ibc.NewHandler(ibcMapper, coinKeeper)).
		AddRoute("simplestake", simplestake.NewHandler(stakeKeeper)).
		AddRoute("did", did.NewHandler(didKeeper)).
		AddRoute("project", project.NewHandler(projectKeeper, coinKeeper))

	// initialize BaseApp
	app.SetTxDecoder(app.txDecoder)
	app.SetInitChainer(app.initChainer)
	app.MountStoreWithDB(app.capKeyMainStore, sdk.StoreTypeIAVL, dbs["main"])
	app.MountStoreWithDB(app.capKeyAccountStore, sdk.StoreTypeIAVL, dbs["acc"])
	app.MountStoreWithDB(app.capKeyIBCStore, sdk.StoreTypeIAVL, dbs["ibc"])
	app.MountStoreWithDB(app.capKeyStakingStore, sdk.StoreTypeIAVL, dbs["staking"])
	app.MountStoreWithDB(app.capKeyDIDStore, sdk.StoreTypeIAVL, dbs["did"])
	app.MountStoreWithDB(app.capKeyProjectStore, sdk.StoreTypeIAVL, dbs["project"])
	// NOTE: Broken until #532 lands
	//app.MountStoresIAVL(app.capKeyMainStore, app.capKeyIBCStore, app.capKeyStakingStore)
	app.SetAnteHandler(NewIxoAnteHandler(app))
	err := app.LoadLatestVersion(app.capKeyMainStore)
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

	const accTypeApp = 0x1
	var _ = oldwire.RegisterInterface(
		struct{ sdk.Account }{},
		oldwire.ConcreteType{&types.AppAccount{}, accTypeApp},
	)
	cdc := wire.NewCodec()

	// cdc.RegisterInterface((*sdk.Msg)(nil), nil)
	// bank.RegisterWire(cdc)   // Register bank.[SendMsg,IssueMsg] types.
	// crypto.RegisterWire(cdc) // Register crypto.[PubKey,PrivKey,Signature] types.
	// ibc.RegisterWire(cdc) // Register ibc.[IBCTransferMsg, IBCReceiveMsg] types.
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
		strs := upTx["payload"].([]interface{})
		signedBytes := strs[1].(string)
		fmt.Println("** signedBytes: ", signedBytes)
		// Check if it is not json
		if strings.Index(signedBytes, "{") == -1 {
			jsonBytes := make([]byte, hex.DecodedLen(len(signedBytes)))
			jsonBytes, hexErr := hex.DecodeString(signedBytes)
			if hexErr != nil {
				fmt.Print("Error decoding hex payload: ", signedBytes)
				fmt.Println()
				return nil, sdk.ErrTxDecode("").TraceCause(hexErr, "")
			}
			jsonBytes = bytes.Replace(jsonBytes, []byte("{"), []byte("{\"signBytes\":\""+signedBytes+"\","), 1)
			txBytes = bytes.Replace(txBytes, []byte("\""+signedBytes+"\""), jsonBytes, 1)
		}
		// StdTx.Msg is an interface. The concrete types
		// are registered by MakeTxCodec in bank.RegisterWire.
		err := app.cdc.UnmarshalJSON(txBytes, &tx)
		if err != nil {
			return nil, sdk.ErrTxDecode("").TraceCause(err, "")
		}

		// fmt.Println("** TXN_PAYLOAD", tx)

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

// custom logic for ixo initialization
func (app *IxoApp) initChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	stateJSON := req.AppStateBytes

	genesisState := new(types.GenesisState)
	err := json.Unmarshal(stateJSON, genesisState)
	if err != nil {
		panic(err) // TODO https://github.com/cosmos/cosmos-sdk/issues/468
		// return sdk.ErrGenesisParse("").TraceCause(err, "")
	}

	projectGenesis := genesisState.ProjectGenesis
	fmt.Println("peg_key: %s", projectGenesis.PegPubKey)
	ethPegDidDoc := did.BaseDidDoc{
		Did:    "ETH_PEG",
		PubKey: projectGenesis.PegPubKey,
	}

	app.didMapper.SetDidDoc(ctx, ethPegDidDoc)

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

// NewIxoAnteHandler returns an AnteHandler that wrappers
// the default cosmos one for signature checking. Based on
// the message type it either checks the Sovrin signature
// or executes the defualt cosmos version
func NewIxoAnteHandler(app *IxoApp) sdk.AnteHandler {
	cosmosAnteHandler := auth.NewAnteHandler(app.accountMapper)
	didAnteHandler := did.NewAnteHandler(app.didMapper)
	projectAnteHandler := project.NewAnteHandler(app.projectMapper, app.didMapper)

	return func(
		ctx sdk.Context, tx sdk.Tx,
	) (_ sdk.Context, _ sdk.Result, abort bool) {

		msg := tx.GetMsg()

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

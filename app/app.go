package app

import (
	"fmt"

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

	base58 "github.com/btcsuite/btcutil/base58"
	"github.com/ixofoundation/ixo-cosmos/types"
	"github.com/ixofoundation/ixo-cosmos/x/did"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
	"github.com/ixofoundation/ixo-cosmos/x/project"
)

const (
	appName = "ixoApp"
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
		BaseApp:            bam.NewBaseApp(appName, logger, dbs["main"]),
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
		app.capKeyProjectStore,    // target store
		&project.BaseProjectDoc{}, // prototype
	)

	// add handlers
	coinKeeper := bank.NewCoinKeeper(app.accountMapper)
	ibcMapper := ibc.NewIBCMapper(app.cdc, app.capKeyIBCStore)
	stakeKeeper := simplestake.NewKeeper(app.capKeyStakingStore, coinKeeper)
	didKeeper := did.NewKeeper(app.didMapper)
	projectKeeper := project.NewKeeper(app.projectMapper)
	app.Router().
		AddRoute("bank", bank.NewHandler(coinKeeper)).
		//		AddRoute("project", project.NewHandler()).
		AddRoute("ibc", ibc.NewHandler(ibcMapper, coinKeeper)).
		AddRoute("simplestake", simplestake.NewHandler(stakeKeeper)).
		AddRoute("did", did.NewHandler(didKeeper)).
		AddRoute("project", project.NewHandler(projectKeeper))

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
	app.SetAnteHandler(NewIxoAnteHandler(didKeeper, auth.NewAnteHandler(app.accountMapper)))
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
	const msgTypeGetDidMsg = 0xA
	const msgTypeAddDidMsg = 0xB
	const msgTypeAddProjectMsg = 0xC
	var _ = oldwire.RegisterInterface(
		struct{ sdk.Msg }{},
		oldwire.ConcreteType{bank.SendMsg{}, msgTypeSend},
		oldwire.ConcreteType{bank.IssueMsg{}, msgTypeIssue},
		oldwire.ConcreteType{ibc.IBCTransferMsg{}, msgTypeIBCTransferMsg},
		oldwire.ConcreteType{ibc.IBCReceiveMsg{}, msgTypeIBCReceiveMsg},
		oldwire.ConcreteType{simplestake.BondMsg{}, msgTypeBondMsg},
		oldwire.ConcreteType{simplestake.UnbondMsg{}, msgTypeUnbondMsg},

		oldwire.ConcreteType{did.AddDidMsg{}, msgTypeAddDidMsg},
		oldwire.ConcreteType{project.AddProjectMsg{}, msgTypeAddProjectMsg},
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

func (app *IxoApp) Commit() (res abci.ResponseCommit) {
	result := app.BaseApp.Commit()
	fmt.Println("IxoAppCommit")
	fmt.Println(result)
	return result
}

// custom logic for transaction decoding
func (app *IxoApp) txDecoder(txBytes []byte) (sdk.Tx, sdk.Error) {
	var tx = ixo.IxoTx{}

	if len(txBytes) == 0 {
		return nil, sdk.ErrTxDecode("txBytes are empty")
	}

	fmt.Println("********DECODED_TXN********* \n", string(txBytes))
	// StdTx.Msg is an interface. The concrete types
	// are registered by MakeTxCodec in bank.RegisterWire.
	err := app.cdc.UnmarshalBinary(txBytes, &tx)
	if err != nil {
		return nil, sdk.ErrTxDecode("").TraceCause(err, "")
	}
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
func NewIxoAnteHandler(dk did.DidKeeper, cosmosAnteHandler sdk.AnteHandler) sdk.AnteHandler {
	return func(
		ctx sdk.Context, tx sdk.Tx,
	) (_ sdk.Context, _ sdk.Result, abort bool) {

		msg := tx.GetMsg()

		fmt.Println("********MSG_TYPE********* \n", msg.Type())

		if msg.Type() != "project" && msg.Type() != "did" {
			// Not an ixo message so execute the wrappered version
			return cosmosAnteHandler(ctx, tx)
		}
		// This always be a IxoTx
		ixoTx, ok := tx.(ixo.IxoTx)
		if !ok {
			return ctx, sdk.ErrInternal("tx must be ixo.IxoTx").Result(), true
		}

		pubKey := [32]byte{}

		if msg.Type() == "did" {
			addDidMsg := msg.(did.AddDidMsg)
			copy(pubKey[:], base58.Decode(addDidMsg.DidDoc.PubKey))

			// Assert dids are the same
			if addDidMsg.DidDoc.Did != ixoTx.Signature.Creator {
				return ctx, sdk.ErrInternal("did in payload does not match creator").Result(), true
			}
		} else if msg.Type() == "project" {
			addProjectMsg := msg.(project.AddProjectMsg)
			copy(pubKey[:], base58.Decode(addProjectMsg.ProjectDoc.PubKey))

			// Assert dids are the same
			if addProjectMsg.ProjectDoc.Did != ixoTx.Signature.Creator {
				return ctx, sdk.ErrInternal("did in payload does not match creator").Result(), true
			}
		} else {
			didDoc := dk.GetDidDoc(ctx, ixoTx.Signature.Creator)
			if didDoc != nil {
				return ctx, sdk.ErrInternal("did not found").Result(), true
			}
			copy(pubKey[:], base58.Decode(didDoc.GetPubKey()))
		}

		// Assert that there are signatures.
		var sigs = tx.GetSignatures()
		if len(sigs) != 1 {
			return ctx,
				sdk.ErrUnauthorized("no signers").Result(),
				true
		}

		res := ixo.VerifySignature(msg, pubKey, sigs[0])

		if !res {
			return ctx, sdk.ErrInternal("Signature Verification failed").Result(), true
		}
		fmt.Println("Signature Verified!")

		return ctx, sdk.Result{}, false // continue...
	}
}

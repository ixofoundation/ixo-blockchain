package app

import (
	"math/rand"
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	wasm "github.com/CosmWasm/wasmd/x/wasm"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sims "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktestutil "github.com/cosmos/cosmos-sdk/x/bank/testutil"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/require"

	"github.com/ixofoundation/ixo-blockchain/v8/lib/ixo"
	bondstypes "github.com/ixofoundation/ixo-blockchain/v8/x/bonds/types"
	entitytypes "github.com/ixofoundation/ixo-blockchain/v8/x/entity/types"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v8/x/iid/types"
)

// This file is the highest-fidelity harness: it signs REAL transactions with a
// real secp256k1 key and runs them through the actual IxoAnteHandler — the full
// decorator chain wired in app/ante.go (validate-basic → bonds guard → IID
// resolution → entity NFT block → smart-account circuit breaker → signature
// verification + fee deduction). It proves the security decorators reject
// hostile signed txs and let legitimate ones through, end to end.

// signerEnv is a fresh app + context + funded signing account + constructed
// ante handler. Each test gets its own so sequence/state never couple.
type signerEnv struct {
	app     *IxoApp
	ctx     sdk.Context
	handler sdk.AnteHandler
	priv    cryptotypes.PrivKey
	addr    sdk.AccAddress
	accNum  uint64
}

func newSignerEnv(t *testing.T) signerEnv {
	t.Helper()
	a := Setup(false)
	ctx := a.BaseApp.NewContextLegacy(false, cmtproto.Header{
		Height:  1,
		ChainID: TestChainID,
		Time:    time.Now().UTC(),
	})

	wasmConfig, err := wasm.ReadWasmConfig(sims.EmptyAppOptions{})
	require.NoError(t, err)
	handler, err := IxoAnteHandler(HandlerOptions{
		HandlerOptions: ante.HandlerOptions{
			AccountKeeper:   a.AccountKeeper,
			BankKeeper:      a.BankKeeper,
			FeegrantKeeper:  a.FeeGrantKeeper,
			SignModeHandler: a.txConfig.SignModeHandler(),
			SigGasConsumer:  ixo.IxoSigVerificationGasConsumer,
		},
		IidKeeper:          a.IidKeeper,
		EntityKeeper:       a.EntityKeeper,
		WasmConfig:         wasmConfig,
		IBCKeeper:          a.IBCKeeper,
		TxCounterStoreKey:  runtime.NewKVStoreService(a.GetKey(wasmtypes.StoreKey)),
		appCodec:           a.appCodec,
		smartAccountKeeper: a.SmartAccountKeeper,
	})
	require.NoError(t, err)

	priv := secp256k1.GenPrivKey()
	addr := sdk.AccAddress(priv.PubKey().Address())
	require.NoError(t, banktestutil.FundAccount(ctx, a.BankKeeper, addr,
		sdk.NewCoins(sdk.NewCoin(ixo.IxoNativeToken, sdkmath.NewInt(1_000_000)))))
	acc := a.AccountKeeper.GetAccount(ctx, addr)
	require.NotNil(t, acc, "funded account must exist")

	return signerEnv{app: a, ctx: ctx, handler: handler, priv: priv, addr: addr, accNum: acc.GetAccountNumber()}
}

// run signs msgs with the env's key and pushes the tx through the full ante.
func (e signerEnv) run(t *testing.T, msgs ...sdk.Msg) error {
	t.Helper()
	tx, err := sims.GenSignedMockTx(
		rand.New(rand.NewSource(1)),
		e.app.txConfig,
		msgs,
		sdk.NewCoins(sdk.NewCoin(ixo.IxoNativeToken, sdkmath.NewInt(10_000))), // fee
		2_000_000, // gas
		TestChainID,
		[]uint64{e.accNum},
		[]uint64{0}, // sequence
		e.priv,
	)
	require.NoError(t, err)
	_, err = e.handler(e.ctx, tx, false)
	return err
}

func (e signerEnv) seedDID(did string, addr sdk.AccAddress) {
	methodID := did + "#key-1"
	vm := iidtypes.NewVerificationMethod(methodID, iidtypes.DID(did), iidtypes.NewBlockchainAccountID(addr.String()))
	meta := iidtypes.NewDidMetadata(e.ctx.TxBytes(), e.ctx.BlockTime())
	e.app.IidKeeper.SetDidDocument(e.ctx, []byte(did), iidtypes.IidDocument{
		Id:                 did,
		VerificationMethod: []*iidtypes.VerificationMethod{&vm},
		Authentication:     []string{methodID},
		Metadata:           &meta,
	})
}

// --- positive control: a legitimate signed tx passes the entire ante ---

func TestFullAnte_BankSendSucceeds(t *testing.T) {
	e := newSignerEnv(t)
	recipient := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
	msg := banktypes.NewMsgSend(e.addr, recipient,
		sdk.NewCoins(sdk.NewCoin(ixo.IxoNativeToken, sdkmath.NewInt(1))))

	require.NoError(t, e.run(t, msg),
		"a correctly-signed, funded bank send must pass the full ante chain")
}

// TestFullAnte_BadSignatureRejected proves signature verification is genuinely
// active in this harness (so the positive tests pass for the right reason): a
// tx for e.addr signed by a DIFFERENT key must be rejected.
func TestFullAnte_BadSignatureRejected(t *testing.T) {
	e := newSignerEnv(t)
	recipient := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
	msg := banktypes.NewMsgSend(e.addr, recipient,
		sdk.NewCoins(sdk.NewCoin(ixo.IxoNativeToken, sdkmath.NewInt(1))))

	wrongPriv := secp256k1.GenPrivKey() // not the account's key
	tx, err := sims.GenSignedMockTx(
		rand.New(rand.NewSource(2)), e.app.txConfig, []sdk.Msg{msg},
		sdk.NewCoins(sdk.NewCoin(ixo.IxoNativeToken, sdkmath.NewInt(10_000))), 2_000_000,
		TestChainID, []uint64{e.accNum}, []uint64{0}, wrongPriv,
	)
	require.NoError(t, err)
	_, err = e.handler(e.ctx, tx, false)
	require.Error(t, err, "a tx signed by the wrong key must be rejected by signature verification")
}

// --- bonds: disabled module is rejected through the real ante ---

func bondsDrainMsg(addr sdk.AccAddress) sdk.Msg {
	return &bondstypes.MsgMakeOutcomePayment{
		BondDid:       "did:ixo:bond-x",
		SenderDid:     iidtypes.DIDFragment("did:ixo:" + addr.String() + "#v1"),
		SenderAddress: addr.String(),
		Amount:        sdkmath.NewInt(1),
	}
}

func TestFullAnte_BondsMsgRejected(t *testing.T) {
	e := newSignerEnv(t)
	err := e.run(t, bondsDrainMsg(e.addr))
	require.Error(t, err)
	require.ErrorIs(t, err, bondstypes.ErrBondsModuleDisabled)
}

func TestFullAnte_BondsMsgInsideMsgExecRejected(t *testing.T) {
	e := newSignerEnv(t)
	exec := authz.NewMsgExec(e.addr, []sdk.Msg{bondsDrainMsg(e.addr)})
	err := e.run(t, &exec)
	require.Error(t, err)
	require.ErrorIs(t, err, bondstypes.ErrBondsModuleDisabled)
}

// --- IID ante: signer→DID binding enforced through the real ante ---
//
// The IID resolution decorator gates messages implementing IidTxMsg (bonds /
// entity / claims). We use MsgUpdateEntity (GetIidController = ControllerDid).
// The ante checks only that the signer controls ControllerDid — it does not
// execute the update — so the entity record need not exist for these tests.

func updateEntityMsg(entityID, controllerDID, controllerAddr string) sdk.Msg {
	return &entitytypes.MsgUpdateEntity{
		Id:                entityID,
		ControllerDid:     iidtypes.DIDFragment(controllerDID),
		ControllerAddress: controllerAddr,
		EntityStatus:      1,
	}
}

func TestFullAnte_IidUnauthorizedSignerRejected(t *testing.T) {
	e := newSignerEnv(t)
	// ControllerDid is controlled by a victim; our signer (e.addr) is NOT on it.
	victim := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
	controllerDID := "did:ixo:victimcontroller"
	e.seedDID(controllerDID, victim)

	err := e.run(t, updateEntityMsg("did:ixo:entity:target", controllerDID, e.addr.String()))
	require.Error(t, err)
	require.Contains(t, err.Error(), "not authorized to act on behalf of the did")
}

func TestFullAnte_IidAuthorizedSignerPasses(t *testing.T) {
	e := newSignerEnv(t)
	// ControllerDid is controlled by our signer.
	controllerDID := "did:ixo:goodcontroller"
	e.seedDID(controllerDID, e.addr)

	require.NoError(t, e.run(t, updateEntityMsg("did:ixo:entity:target", controllerDID, e.addr.String())),
		"the controller DID's own authentication key must pass the full ante chain")
}

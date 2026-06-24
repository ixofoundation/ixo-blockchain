package post_test

import (
	"testing"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/ixofoundation/ixo-blockchain/v8/app/apptesting"
	smartaccountpost "github.com/ixofoundation/ixo-blockchain/v8/x/smart-account/post"
)

// PostHandlerTestSuite covers AuthenticatorPostDecorator's circuit-breaker
// short-circuit branch. The full ConfirmExecution path requires a signed
// authenticator transaction and is exercised end-to-end in interchaintest.
//
// Worth flagging: AuthenticatorPostDecorator's PostHandle implementation
// invokes the *stored* `ad.next` field on the circuit-breaker path instead
// of the `next` parameter that the cosmos-sdk `PostDecorator` interface
// supplies. The implementation works because app/posthandler.go wires a
// terminator into ad.next at construction time, but it deviates from the
// idiomatic pattern (cosmos-sdk's PostDecorator interface only guarantees
// that `next` is non-nil; `ad.next` is module-private). Replacing
// `ad.next(...)` with `next(...)` would let the decorator compose with
// ChainPostDecorators in the conventional way. Not addressed here — see
// tests.md::Bug Log entry.
type PostHandlerTestSuite struct {
	apptesting.KeeperTestHelper
}

func TestPostHandlerTestSuite(t *testing.T) { suite.Run(t, new(PostHandlerTestSuite)) }

func (s *PostHandlerTestSuite) SetupTest() { s.Setup() }

type stubFeeTx struct {
	msgs       []sdk.Msg
	feePayer   sdk.AccAddress
	feeGranter sdk.AccAddress
	fee        sdk.Coins
}

func (t *stubFeeTx) GetMsgs() []sdk.Msg                              { return t.msgs }
func (t *stubFeeTx) GetMsgsV2() ([]protoreflect.ProtoMessage, error) { return nil, nil }
func (t *stubFeeTx) ValidateBasic() error                            { return nil }
func (t *stubFeeTx) GetFee() sdk.Coins                               { return t.fee }
func (t *stubFeeTx) GetGas() uint64                                  { return 200_000 }
func (t *stubFeeTx) FeePayer() []byte                                { return t.feePayer }
func (t *stubFeeTx) FeeGranter() []byte                              { return t.feeGranter }

// trackingPostHandler returns a sdk.PostHandler closure plus a *bool flag
// the closure flips when invoked. We pass it as the *stored* `next` to the
// decorator (not as the chain `next` parameter) because that's what the
// circuit-breaker branch actually calls.
func trackingPostHandler() (sdk.PostHandler, *bool) {
	called := false
	return func(ctx sdk.Context, _ sdk.Tx, _, _ bool) (sdk.Context, error) {
		called = true
		return ctx, nil
	}, &called
}

func (s *PostHandlerTestSuite) TestPostHandle_CircuitBreakActive_SkipsToNext() {
	// Smart-account is active by default in the genesis params; a tx with no
	// extension options still triggers the IsCircuitBreakActive=true branch
	// because the selected-authenticators extension is missing.
	signer := apptesting.RandomAccountAddress()
	tx := &stubFeeTx{
		msgs: []sdk.Msg{&wasmtypes.MsgExecuteContract{
			Sender: signer.String(), Contract: signer.String(), Msg: []byte(`{}`),
		}},
		feePayer: signer,
	}

	storedNext, storedCalled := trackingPostHandler()
	dec := smartaccountpost.NewAuthenticatorPostDecorator(
		s.App.AppCodec(),
		s.App.SmartAccountKeeper,
		s.App.AccountKeeper,
		s.App.TxConfig().SignModeHandler(),
		storedNext,
	)

	chainCalled := false
	_, err := dec.PostHandle(s.Ctx, tx, false, true,
		func(ctx sdk.Context, _ sdk.Tx, _, _ bool) (sdk.Context, error) {
			chainCalled = true
			return ctx, nil
		})
	s.Require().NoError(err)
	s.Require().True(*storedCalled, "circuit-breaker branch must invoke the stored next handler")
	s.Require().False(chainCalled, "circuit-breaker branch must NOT invoke the chain next handler (current implementation)")
}

func (s *PostHandlerTestSuite) TestPostHandle_ModuleInactive_SkipsToNext() {
	// Force the module-wide active flag off — IsCircuitBreakActive returns
	// true unconditionally in this case, so the post handler should skip.
	s.App.SmartAccountKeeper.SetActiveState(s.Ctx, false)
	defer s.App.SmartAccountKeeper.SetActiveState(s.Ctx, true)

	signer := apptesting.RandomAccountAddress()
	tx := &stubFeeTx{
		msgs: []sdk.Msg{&wasmtypes.MsgExecuteContract{
			Sender: signer.String(), Contract: signer.String(), Msg: []byte(`{}`),
		}},
		feePayer: signer,
	}

	storedNext, storedCalled := trackingPostHandler()
	dec := smartaccountpost.NewAuthenticatorPostDecorator(
		s.App.AppCodec(),
		s.App.SmartAccountKeeper,
		s.App.AccountKeeper,
		s.App.TxConfig().SignModeHandler(),
		storedNext,
	)

	_, err := dec.PostHandle(s.Ctx, tx, false, true,
		func(ctx sdk.Context, _ sdk.Tx, _, _ bool) (sdk.Context, error) {
			return ctx, nil
		})
	s.Require().NoError(err)
	s.Require().True(*storedCalled)
}

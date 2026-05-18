package ante_test

import (
	"testing"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	signing "github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/ixofoundation/ixo-blockchain/v6/app/apptesting"
	smartaccountante "github.com/ixofoundation/ixo-blockchain/v6/x/smart-account/ante"
)

// PubKeyAnteTestSuite covers EmitPubKeyDecoratorEvents — the decorator that
// emits sdk.EventTypeTx events with sequence/signature attributes for legacy
// account-system compatibility.
type PubKeyAnteTestSuite struct {
	apptesting.KeeperTestHelper
}

func TestPubKeyAnteTestSuite(t *testing.T) { suite.Run(t, new(PubKeyAnteTestSuite)) }

func (s *PubKeyAnteTestSuite) SetupTest() { s.Setup() }

// stubSigVerifiableTx is a SigVerifiableTx that returns canned signers and
// signatures — enough for EmitPubKeyDecoratorEvents.
type stubSigVerifiableTx struct {
	msgs    []sdk.Msg
	signers [][]byte
	sigs    []signing.SignatureV2
}

func (t *stubSigVerifiableTx) GetMsgs() []sdk.Msg                              { return t.msgs }
func (t *stubSigVerifiableTx) GetMsgsV2() ([]protoreflect.ProtoMessage, error) { return nil, nil }
func (t *stubSigVerifiableTx) ValidateBasic() error                            { return nil }
func (t *stubSigVerifiableTx) GetSigners() ([][]byte, error)                   { return t.signers, nil }
func (t *stubSigVerifiableTx) GetPubKeys() ([]cryptotypes.PubKey, error)       { return nil, nil }
func (t *stubSigVerifiableTx) GetSignaturesV2() ([]signing.SignatureV2, error) {
	return t.sigs, nil
}

// stubNonSigTx is plain sdk.Tx — the decorator should reject it with
// "invalid tx type".
type stubNonSigTx struct{}

func (*stubNonSigTx) GetMsgs() []sdk.Msg                              { return nil }
func (*stubNonSigTx) GetMsgsV2() ([]protoreflect.ProtoMessage, error) { return nil, nil }
func (*stubNonSigTx) ValidateBasic() error                            { return nil }

func (s *PubKeyAnteTestSuite) TestEmitPubKeyDecoratorEvents_NotSigVerifiableTx() {
	dec := smartaccountante.NewEmitPubKeyDecoratorEvents(s.App.AccountKeeper)
	_, err := dec.AnteHandle(s.Ctx, &stubNonSigTx{}, false,
		func(ctx sdk.Context, _ sdk.Tx, _ bool) (sdk.Context, error) { return ctx, nil })
	s.Require().ErrorContains(err, "invalid tx type")
}

func (s *PubKeyAnteTestSuite) TestEmitPubKeyDecoratorEvents_HappyPath() {
	signer := apptesting.RandomAccountAddress()
	tx := &stubSigVerifiableTx{
		msgs: []sdk.Msg{&wasmtypes.MsgExecuteContract{
			Sender: signer.String(), Contract: signer.String(), Msg: []byte(`{}`),
		}},
		signers: [][]byte{signer.Bytes()},
		sigs: []signing.SignatureV2{
			{
				PubKey:   nil,
				Data:     &signing.SingleSignatureData{Signature: []byte("sig-bytes")},
				Sequence: 7,
			},
		},
	}

	dec := smartaccountante.NewEmitPubKeyDecoratorEvents(s.App.AccountKeeper)
	called := false
	_, err := dec.AnteHandle(s.Ctx, tx, false,
		func(ctx sdk.Context, _ sdk.Tx, _ bool) (sdk.Context, error) {
			called = true
			return ctx, nil
		})
	s.Require().NoError(err)
	s.Require().True(called, "decorator must call next on the happy path")

	// Confirm the events made it onto the event manager.
	events := s.Ctx.EventManager().Events()
	foundSeq := false
	for _, ev := range events {
		if ev.Type != sdk.EventTypeTx {
			continue
		}
		for _, a := range ev.Attributes {
			if a.Key == sdk.AttributeKeyAccountSequence {
				foundSeq = true
			}
		}
	}
	s.Require().True(foundSeq, "decorator must emit at least one tx.acc_seq attribute")
}

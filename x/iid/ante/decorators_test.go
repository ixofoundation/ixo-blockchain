package ante_test

import (
	"testing"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	signing "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/cosmos-sdk/x/authz"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/ixofoundation/ixo-blockchain/v8/app/apptesting"
	bondstypes "github.com/ixofoundation/ixo-blockchain/v8/x/bonds/types"
	iidante "github.com/ixofoundation/ixo-blockchain/v8/x/iid/ante"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v8/x/iid/types"
)

// IIDAnteTestSuite covers VerifyIidControllersAgainstSignature: every
// IidTxMsg in the tx must reference an existing DID document, and the
// declared msg signer must hold an authentication / assertion verification
// relationship on that document.
type IIDAnteTestSuite struct {
	apptesting.KeeperTestHelper
}

func TestIIDAnteTestSuite(t *testing.T) { suite.Run(t, new(IIDAnteTestSuite)) }

func (s *IIDAnteTestSuite) SetupTest() { s.Setup() }

// stubSigTx is a minimal authsigning.SigVerifiableTx — VerifyIidControllers
// only calls GetMsgs(), so the rest of the interface returns zero values.
type stubSigTx struct{ msgs []sdk.Msg }

func (t *stubSigTx) GetMsgs() []sdk.Msg                              { return t.msgs }
func (t *stubSigTx) GetMsgsV2() ([]protoreflect.ProtoMessage, error) { return nil, nil }
func (t *stubSigTx) ValidateBasic() error                            { return nil }

// SigVerifiableTx surface
func (t *stubSigTx) GetSigners() ([][]byte, error)             { return nil, nil }
func (t *stubSigTx) GetPubKeys() ([]cryptotypes.PubKey, error) { return nil, nil }
func (t *stubSigTx) GetSignaturesV2() ([]signing.SignatureV2, error) {
	return nil, nil
}

var _ authsigning.SigVerifiableTx = (*stubSigTx)(nil)

// seedIIDFor produces a DID document with `signer` registered as the
// authentication-relationship verification material so VerifyIidControllers
// will accept it as a valid actor.
func (s *IIDAnteTestSuite) seedIIDFor(signer sdk.AccAddress) string {
	did := "did:ixo:" + signer.String()[:24]
	methodID := did + "#key-1"
	vm := iidtypes.NewVerificationMethod(
		methodID,
		iidtypes.DID(did),
		iidtypes.NewBlockchainAccountID(signer.String()),
	)
	meta := iidtypes.NewDidMetadata(s.Ctx.TxBytes(), s.Ctx.BlockTime())
	doc := iidtypes.IidDocument{
		Id:                 did,
		VerificationMethod: []*iidtypes.VerificationMethod{&vm},
		Authentication:     []string{methodID},
		Metadata:           &meta,
	}
	s.App.IidKeeper.SetDidDocument(s.Ctx, []byte(did), doc)
	return did
}

func (s *IIDAnteTestSuite) cdc() codec.Codec { return s.App.AppCodec() }

func (s *IIDAnteTestSuite) TestVerifyIidControllers_HappyPath() {
	s.SetupTest()
	signer := apptesting.RandomAccountAddress()
	did := s.seedIIDFor(signer)

	msg := &bondstypes.MsgCreateBond{
		BondDid:                  "did:ixo:bond-1",
		CreatorDid:               iidtypes.DIDFragment(did),
		Token:                    "tok",
		Name:                     "test",
		ReserveTokens:            []string{"uixo"},
		FeeAddress:               signer.String(),
		ReserveWithdrawalAddress: signer.String(),
		MaxSupply:                sdk.NewCoin("tok", math.NewInt(1000)),
		CreatorAddress:           signer.String(),
	}
	tx := &stubSigTx{msgs: []sdk.Msg{msg}}

	err := iidante.VerifyIidControllersAgainstSignature(tx, s.Ctx, s.App.IidKeeper, s.cdc())
	s.Require().NoError(err)
}

func (s *IIDAnteTestSuite) TestVerifyIidControllers_DIDNotFound() {
	s.SetupTest()
	signer := apptesting.RandomAccountAddress()

	msg := &bondstypes.MsgCreateBond{
		BondDid:                  "did:ixo:bond-2",
		CreatorDid:               iidtypes.DIDFragment("did:ixo:does-not-exist"),
		FeeAddress:               signer.String(),
		ReserveWithdrawalAddress: signer.String(),
		CreatorAddress:           signer.String(),
	}
	tx := &stubSigTx{msgs: []sdk.Msg{msg}}

	err := iidante.VerifyIidControllersAgainstSignature(tx, s.Ctx, s.App.IidKeeper, s.cdc())
	s.Require().ErrorContains(err, "did document")
	s.Require().ErrorContains(err, "not found")
}

func (s *IIDAnteTestSuite) TestVerifyIidControllers_UnauthorizedSigner() {
	s.SetupTest()
	owner := apptesting.RandomAccountAddress()
	stranger := apptesting.RandomAccountAddress()
	did := s.seedIIDFor(owner)

	// Construct a Msg whose declared CreatorDid is `did` but whose tx-level
	// signer is `stranger`. CreatorDid is derived via cdc.GetMsgV1Signers(msg);
	// for MsgCreateBond the signer is FeeAddress.
	msg := &bondstypes.MsgCreateBond{
		BondDid:                  "did:ixo:bond-3",
		CreatorDid:               iidtypes.DIDFragment(did),
		FeeAddress:               stranger.String(),
		ReserveWithdrawalAddress: stranger.String(),
		CreatorAddress:           stranger.String(),
	}
	tx := &stubSigTx{msgs: []sdk.Msg{msg}}

	err := iidante.VerifyIidControllersAgainstSignature(tx, s.Ctx, s.App.IidKeeper, s.cdc())
	s.Require().ErrorContains(err, "not authorized to act on behalf of the did")
}

// TestVerifyIidControllers_UnauthorizedSignerInsideMsgExec is the core
// regression test for the 2026-06-20 defence-in-depth gap: an IidTxMsg with an
// unauthorized signer, hidden inside an authz.MsgExec, must now be rejected.
// Before the fix the ante iterated top-level messages only and this passed.
func (s *IIDAnteTestSuite) TestVerifyIidControllers_UnauthorizedSignerInsideMsgExec() {
	s.SetupTest()
	owner := apptesting.RandomAccountAddress()
	stranger := apptesting.RandomAccountAddress()
	did := s.seedIIDFor(owner)

	inner := &bondstypes.MsgCreateBond{
		BondDid:                  "did:ixo:bond-exec",
		CreatorDid:               iidtypes.DIDFragment(did),
		FeeAddress:               stranger.String(),
		ReserveWithdrawalAddress: stranger.String(),
		CreatorAddress:           stranger.String(),
	}
	exec := authz.NewMsgExec(stranger, []sdk.Msg{inner})
	tx := &stubSigTx{msgs: []sdk.Msg{&exec}}

	err := iidante.VerifyIidControllersAgainstSignature(tx, s.Ctx, s.App.IidKeeper, s.cdc())
	s.Require().ErrorContains(err, "not authorized to act on behalf of the did")
}

// TestVerifyIidControllers_AuthorizedSignerInsideMsgExec confirms the recursion
// does not over-reject: a correctly-authorized IidTxMsg nested in a MsgExec
// still passes.
func (s *IIDAnteTestSuite) TestVerifyIidControllers_AuthorizedSignerInsideMsgExec() {
	s.SetupTest()
	owner := apptesting.RandomAccountAddress()
	did := s.seedIIDFor(owner)

	inner := &bondstypes.MsgCreateBond{
		BondDid:                  "did:ixo:bond-exec-ok",
		CreatorDid:               iidtypes.DIDFragment(did),
		Token:                    "tok",
		Name:                     "test",
		ReserveTokens:            []string{"uixo"},
		FeeAddress:               owner.String(),
		ReserveWithdrawalAddress: owner.String(),
		MaxSupply:                sdk.NewCoin("tok", math.NewInt(1000)),
		CreatorAddress:           owner.String(),
	}
	exec := authz.NewMsgExec(owner, []sdk.Msg{inner})
	tx := &stubSigTx{msgs: []sdk.Msg{&exec}}

	err := iidante.VerifyIidControllersAgainstSignature(tx, s.Ctx, s.App.IidKeeper, s.cdc())
	s.Require().NoError(err)
}

// TestDecorator_RejectsUnauthorizedInsideMsgExec exercises the wired
// IidResolutionDecorator.AnteHandle end-to-end (the wrapper actually placed in
// the app ante chain), confirming an unauthorized IidTxMsg hidden in a
// MsgExec is rejected and next() is never reached.
func (s *IIDAnteTestSuite) TestDecorator_RejectsUnauthorizedInsideMsgExec() {
	s.SetupTest()
	owner := apptesting.RandomAccountAddress()
	stranger := apptesting.RandomAccountAddress()
	did := s.seedIIDFor(owner)

	inner := &bondstypes.MsgCreateBond{
		BondDid:                  "did:ixo:bond-dec",
		CreatorDid:               iidtypes.DIDFragment(did),
		FeeAddress:               stranger.String(),
		ReserveWithdrawalAddress: stranger.String(),
		CreatorAddress:           stranger.String(),
	}
	exec := authz.NewMsgExec(stranger, []sdk.Msg{inner})
	tx := &stubSigTx{msgs: []sdk.Msg{&exec}}

	dec := iidante.NewIidResolutionDecorator(s.App.IidKeeper, s.cdc())
	nextCalled := false
	next := func(ctx sdk.Context, _ sdk.Tx, _ bool) (sdk.Context, error) {
		nextCalled = true
		return ctx, nil
	}
	_, err := dec.AnteHandle(s.Ctx, tx, false, next)
	s.Require().Error(err)
	s.Require().False(nextCalled, "ante must short-circuit before next()")
}

// TestVerifyIidControllers_NonIidMsgPasses confirms that messages outside
// the IID family are skipped (the type assertion to IidTxMsg fails and
// continue runs).
func (s *IIDAnteTestSuite) TestVerifyIidControllers_NonIidMsgPasses() {
	s.SetupTest()
	// MsgCreateIidDocument doesn't implement IidTxMsg — only mutating bond
	// messages do. But we want any non-IidTxMsg here. Use a different bonds
	// msg that doesn't carry an IidController.
	tx := &stubSigTx{msgs: []sdk.Msg{}}
	err := iidante.VerifyIidControllersAgainstSignature(tx, s.Ctx, s.App.IidKeeper, s.cdc())
	s.Require().NoError(err)
}

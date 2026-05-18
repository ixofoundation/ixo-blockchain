package ante_test

import (
	"testing"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	signing "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/ixofoundation/ixo-blockchain/v6/app/apptesting"
	bondstypes "github.com/ixofoundation/ixo-blockchain/v6/x/bonds/types"
	iidante "github.com/ixofoundation/ixo-blockchain/v6/x/iid/ante"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v6/x/iid/types"
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

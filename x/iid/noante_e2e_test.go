package iid_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/ixofoundation/ixo-blockchain/v7/app/apptesting"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v7/x/iid/types"
)

// IidNoAnteTestSuite dispatches iid messages through the app's REAL
// MsgServiceRouter via RunMsg, which skips the ante handler — reproducing the
// CosmWasm / ICA-host / authz routes. It asserts the iid KEEPER itself binds
// the signer to the DID (via ExecuteOnDidWithRelationships → HasRelationship on
// the signer ADDRESS), so authorization holds without the ante.
type IidNoAnteTestSuite struct {
	apptesting.KeeperTestHelper
}

func TestIidNoAnteTestSuite(t *testing.T) { suite.Run(t, new(IidNoAnteTestSuite)) }

func (s *IidNoAnteTestSuite) SetupTest() { s.Setup() }

// seedDID registers a DID whose only authentication key is `addr`.
func (s *IidNoAnteTestSuite) seedDID(did string, addr sdk.AccAddress) {
	methodID := did + "#key-1"
	vm := iidtypes.NewVerificationMethod(methodID, iidtypes.DID(did), iidtypes.NewBlockchainAccountID(addr.String()))
	meta := iidtypes.NewDidMetadata(s.Ctx.TxBytes(), s.Ctx.BlockTime())
	doc := iidtypes.IidDocument{
		Id:                 did,
		VerificationMethod: []*iidtypes.VerificationMethod{&vm},
		Authentication:     []string{methodID},
		Metadata:           &meta,
	}
	s.App.IidKeeper.SetDidDocument(s.Ctx, []byte(did), doc)
}

func (s *IidNoAnteTestSuite) addServiceMsg(did, signer string) *iidtypes.MsgAddService {
	return &iidtypes.MsgAddService{
		Id:          did,
		ServiceData: iidtypes.NewService("https://example.com/svc", "LinkedDomains", "https://example.com"),
		Signer:      signer,
	}
}

// TestAddService_NoAnte_RejectsUnauthorizedSigner: a signer not on the DID is
// rejected by the keeper even with the ante skipped.
func (s *IidNoAnteTestSuite) TestAddService_NoAnte_RejectsUnauthorizedSigner() {
	owner := apptesting.RandomAccountAddress()
	attacker := apptesting.RandomAccountAddress()
	did := "did:ixo:noantetarget"
	s.seedDID(did, owner)

	_, err := s.RunMsg(s.addServiceMsg(did, attacker.String()))
	s.Require().Error(err)
	s.Require().ErrorIs(err, iidtypes.ErrUnauthorized)

	doc, found := s.App.IidKeeper.GetDidDocument(s.Ctx, []byte(did))
	s.Require().True(found)
	s.Require().Empty(doc.Service, "service must not have been added by an unauthorized signer")
}

// TestAddService_NoAnte_AllowsAuthorizedSigner: the DID's authentication key
// succeeds on the same no-ante route.
func (s *IidNoAnteTestSuite) TestAddService_NoAnte_AllowsAuthorizedSigner() {
	owner := apptesting.RandomAccountAddress()
	did := "did:ixo:noantetargetok"
	s.seedDID(did, owner)

	_, err := s.RunMsg(s.addServiceMsg(did, owner.String()))
	s.Require().NoError(err)

	doc, _ := s.App.IidKeeper.GetDidDocument(s.Ctx, []byte(did))
	s.Require().Len(doc.Service, 1)
}

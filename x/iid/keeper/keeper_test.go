package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/ixofoundation/ixo-blockchain/v6/app/apptesting"
	"github.com/ixofoundation/ixo-blockchain/v6/x/iid/keeper"
	"github.com/ixofoundation/ixo-blockchain/v6/x/iid/types"
)

// KeeperTestSuite covers x/iid: 20 message handlers, queries, genesis,
// DID validation, and store-key encoding. Most operations route through
// keeper.ExecuteOnDidWithRelationships, so the suite has a single helper
// (createDIDFor) that produces a DID document with the signer registered as
// a controller — that's the cheapest authorisation path through the
// HasRelationship/HasController two-step.
type KeeperTestSuite struct {
	apptesting.KeeperTestHelper

	msgServer   types.MsgServer
	queryServer types.QueryServer
	queryClient types.QueryClient
}

func TestKeeperTestSuite(t *testing.T) { suite.Run(t, new(KeeperTestSuite)) }

func (s *KeeperTestSuite) SetupTest() {
	s.Setup()
	s.msgServer = keeper.NewMsgServerImpl(s.App.IidKeeper)
	s.queryServer = s.App.IidKeeper
	s.queryClient = types.NewQueryClient(s.QueryHelper)
}

func (s *KeeperTestSuite) goCtx() context.Context { return s.Ctx }

// createDIDFor persists a DID document where `signer` can authenticate. The
// document carries a VerificationMethod whose BlockchainAccountID equals the
// signer's bech32 address, and that method is listed under the Authentication
// relationship — exactly what ExecuteOnDidWithRelationships requires.
//
// We bypass NewDidDocument here so the keeper can be seeded synchronously
// without worrying about the input-validation rules that the production
// constructor enforces. This is fine for tests but should NOT be done in
// keeper code.
func (s *KeeperTestSuite) createDIDFor(signer sdk.AccAddress) string {
	did := "did:ixo:" + sdk.AccAddress(signer).String()[:24]
	methodID := did + "#key-1"
	vm := types.NewVerificationMethod(
		methodID,
		types.DID(did),
		types.NewBlockchainAccountID(signer.String()),
	)
	meta := types.NewDidMetadata(s.Ctx.TxBytes(), s.Ctx.BlockTime())
	doc := types.IidDocument{
		Id:                 did,
		VerificationMethod: []*types.VerificationMethod{&vm},
		Authentication:     []string{methodID},
		Metadata:           &meta,
	}
	s.App.IidKeeper.SetDidDocument(s.Ctx, []byte(did), doc)
	return did
}

// freshSigner returns (signer, did) where the DID is brand-new and controlled
// by the signer.
func (s *KeeperTestSuite) freshSigner() (sdk.AccAddress, string) {
	signer := apptesting.RandomAccountAddress()
	return signer, s.createDIDFor(signer)
}

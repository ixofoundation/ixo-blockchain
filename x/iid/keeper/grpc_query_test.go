package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/ixofoundation/ixo-blockchain/v7/x/iid/types"
)

func (s *KeeperTestSuite) TestQueryIidDocument() {
	s.SetupTest()
	signer, did := s.freshSigner()
	_ = signer

	resp, err := s.queryClient.IidDocument(s.goCtx(), &types.QueryIidDocumentRequest{Id: did})
	s.Require().NoError(err)
	s.Require().Equal(did, resp.IidDocument.Id)

	_, err = s.queryClient.IidDocument(s.goCtx(), &types.QueryIidDocumentRequest{Id: "did:ixo:ghost"})
	s.Require().ErrorContains(err, "not found")
}

func (s *KeeperTestSuite) TestQueryIidDocuments() {
	s.SetupTest()
	for i := 0; i < 3; i++ {
		_, _ = s.freshSigner()
	}

	resp, err := s.queryClient.IidDocuments(s.goCtx(), &types.QueryIidDocumentsRequest{
		Pagination: &query.PageRequest{Limit: 100},
	})
	s.Require().NoError(err)
	s.Require().GreaterOrEqual(len(resp.IidDocuments), 3)
}

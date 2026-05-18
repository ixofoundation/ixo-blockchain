package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/ixofoundation/ixo-blockchain/v6/app/apptesting"
	"github.com/ixofoundation/ixo-blockchain/v6/x/names/types"
)

// ----------------------------------------------------------------------------
// Namespace
// ----------------------------------------------------------------------------

func (s *KeeperTestSuite) TestQueryNamespace() {
	s.SetupTest()
	ns := s.defaultNamespace("handles")
	s.mustCreateNamespace(ns)

	s.Run("happy path", func() {
		got, err := s.queryClient.Namespace(s.goCtx(), &types.QueryNamespaceRequest{Name: "handles"})
		s.Require().NoError(err)
		s.Require().Equal("handles", got.Namespace.Name)
	})
	s.Run("nil request (direct keeper)", func() {
		// gRPC client never delivers a nil request — the codec substitutes a
		// zero-value pointer — so we exercise the nil-guard via the keeper
		// directly to make sure that defensive branch is still wired up.
		_, err := s.queryServer.Namespace(s.goCtx(), nil)
		s.Require().ErrorContains(err, "empty request")
	})
	s.Run("empty name", func() {
		_, err := s.queryClient.Namespace(s.goCtx(), &types.QueryNamespaceRequest{Name: ""})
		s.Require().ErrorContains(err, "name is required")
	})
	s.Run("not found", func() {
		_, err := s.queryClient.Namespace(s.goCtx(), &types.QueryNamespaceRequest{Name: "missing"})
		s.Require().ErrorContains(err, "namespace not found")
	})
}

// ----------------------------------------------------------------------------
// Namespaces
// ----------------------------------------------------------------------------

func (s *KeeperTestSuite) TestQueryNamespaces() {
	s.SetupTest()
	s.mustCreateNamespace(s.defaultNamespace("handles"))
	s.mustCreateNamespace(s.defaultNamespace("groups"))

	resp, err := s.queryClient.Namespaces(s.goCtx(), &types.QueryNamespacesRequest{})
	s.Require().NoError(err)
	s.Require().Len(resp.Namespaces, 2)

	// non-existent prefix returns empty (paginated), not error
	resp, err = s.queryClient.Namespaces(s.goCtx(), &types.QueryNamespacesRequest{
		Pagination: &query.PageRequest{Limit: 1},
	})
	s.Require().NoError(err)
	s.Require().Len(resp.Namespaces, 1)
	s.Require().NotNil(resp.Pagination.NextKey, "second page key should be non-nil when more remain")
}

// ----------------------------------------------------------------------------
// ResolveName
// ----------------------------------------------------------------------------

func (s *KeeperTestSuite) TestQueryResolveName() {
	s.SetupTest()
	s.mustCreateNamespace(s.defaultNamespace("handles"))
	signer, did := s.addrAndDID()
	_, err := s.msgServer.RegisterName(s.goCtx(), &types.MsgRegisterName{
		Signer: signer.String(), Namespace: "handles", Name: "Alice", OwnerDid: did,
	})
	s.Require().NoError(err)

	s.Run("display name normalises before lookup", func() {
		resp, err := s.queryClient.ResolveName(s.goCtx(), &types.QueryResolveNameRequest{
			Namespace: "handles", Name: "ALICE",
		})
		s.Require().NoError(err)
		s.Require().Equal("alice", resp.Record.NormalizedName)
	})

	s.Run("non-active record is hidden", func() {
		// suspend
		// (registrar of self-register namespace: any signer can suspend? no, only registrar/authority)
		// Use authority to suspend.
		_, err := s.msgServer.SetNameStatus(s.goCtx(), &types.MsgSetNameStatus{
			Signer: s.authority, Namespace: "handles",
			NormalizedName: "alice", Status: types.NAME_STATUS_SUSPENDED,
		})
		s.Require().NoError(err)
		_, err = s.queryClient.ResolveName(s.goCtx(), &types.QueryResolveNameRequest{
			Namespace: "handles", Name: "alice",
		})
		s.Require().ErrorContains(err, "name not found", "ResolveName must hide non-active records")
	})

	s.Run("nil request (direct keeper)", func() {
		_, err := s.queryServer.ResolveName(s.goCtx(), nil)
		s.Require().ErrorContains(err, "empty request")
	})
	s.Run("empty namespace", func() {
		_, err := s.queryClient.ResolveName(s.goCtx(), &types.QueryResolveNameRequest{Name: "alice"})
		s.Require().ErrorContains(err, "namespace and name are required")
	})
}

// ----------------------------------------------------------------------------
// GetName
// ----------------------------------------------------------------------------

func (s *KeeperTestSuite) TestQueryGetName() {
	s.SetupTest()
	s.mustCreateNamespace(s.defaultNamespace("handles"))
	signer, did := s.addrAndDID()
	_, err := s.msgServer.RegisterName(s.goCtx(), &types.MsgRegisterName{
		Signer: signer.String(), Namespace: "handles", Name: "alice", OwnerDid: did,
	})
	s.Require().NoError(err)

	// suspend so we can confirm GetName still returns it (unlike ResolveName)
	_, err = s.msgServer.SetNameStatus(s.goCtx(), &types.MsgSetNameStatus{
		Signer: s.authority, Namespace: "handles",
		NormalizedName: "alice", Status: types.NAME_STATUS_SUSPENDED,
	})
	s.Require().NoError(err)

	resp, err := s.queryClient.GetName(s.goCtx(), &types.QueryGetNameRequest{
		Namespace: "handles", NormalizedName: "alice",
	})
	s.Require().NoError(err)
	s.Require().Equal(types.NAME_STATUS_SUSPENDED, resp.Record.Status)

	_, err = s.queryClient.GetName(s.goCtx(), &types.QueryGetNameRequest{
		Namespace: "handles", NormalizedName: "ghost",
	})
	s.Require().ErrorContains(err, "name not found")
}

// ----------------------------------------------------------------------------
// NamesByNamespace
// ----------------------------------------------------------------------------

func (s *KeeperTestSuite) TestQueryNamesByNamespace() {
	s.SetupTest()
	s.mustCreateNamespace(s.defaultNamespace("handles"))
	signer, did := s.addrAndDID()
	for _, n := range []string{"alice", "bob", "carol"} {
		_, err := s.msgServer.RegisterName(s.goCtx(), &types.MsgRegisterName{
			Signer: signer.String(), Namespace: "handles", Name: n, OwnerDid: did,
		})
		s.Require().NoError(err)
	}

	resp, err := s.queryClient.NamesByNamespace(s.goCtx(), &types.QueryNamesByNamespaceRequest{
		Namespace: "handles",
	})
	s.Require().NoError(err)
	s.Require().Len(resp.Records, 3)

	// missing namespace returns NotFound
	_, err = s.queryClient.NamesByNamespace(s.goCtx(), &types.QueryNamesByNamespaceRequest{
		Namespace: "ghost",
	})
	s.Require().ErrorContains(err, "namespace not found")
}

// ----------------------------------------------------------------------------
// NamesByOwner
// ----------------------------------------------------------------------------

func (s *KeeperTestSuite) TestQueryNamesByOwner() {
	s.SetupTest()
	s.mustCreateNamespace(s.defaultNamespace("handles"))
	s.mustCreateNamespace(s.defaultNamespace("groups"))
	signer, did := s.addrAndDID()
	_, err := s.msgServer.RegisterName(s.goCtx(), &types.MsgRegisterName{
		Signer: signer.String(), Namespace: "handles", Name: "alice", OwnerDid: did,
	})
	s.Require().NoError(err)
	_, err = s.msgServer.RegisterName(s.goCtx(), &types.MsgRegisterName{
		Signer: signer.String(), Namespace: "groups", Name: "team-a", OwnerDid: did,
	})
	s.Require().NoError(err)

	resp, err := s.queryClient.NamesByOwner(s.goCtx(), &types.QueryNamesByOwnerRequest{OwnerDid: did})
	s.Require().NoError(err)
	s.Require().Len(resp.Records, 2, "owner index must span namespaces")

	// owner with no records — empty list, no error
	other, _ := s.addrAndDID()
	_ = other
	resp, err = s.queryClient.NamesByOwner(s.goCtx(), &types.QueryNamesByOwnerRequest{
		OwnerDid: "did:ixo:" + apptesting.RandomAccountAddress().String(),
	})
	s.Require().NoError(err)
	s.Require().Len(resp.Records, 0)

	// empty owner_did is rejected
	_, err = s.queryClient.NamesByOwner(s.goCtx(), &types.QueryNamesByOwnerRequest{})
	s.Require().ErrorContains(err, "owner_did is required")
}

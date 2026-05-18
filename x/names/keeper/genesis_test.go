package keeper_test

import (
	"github.com/ixofoundation/ixo-blockchain/v6/x/names/types"
)

// TestGenesis_RoundTrip ensures InitGenesis followed by ExportGenesis returns
// the same set of namespaces and name records, regardless of insertion order.
func (s *KeeperTestSuite) TestGenesis_RoundTrip() {
	s.SetupTest()

	// Build a non-trivial genesis: two namespaces, three records spanning both.
	gs := types.GenesisState{
		Namespaces: []types.Namespace{
			{Name: "handles", Description: "user handles", AllowSelfRegister: true, MinLength: 3, MaxLength: 32},
			{Name: "verified", Description: "verified handles", AllowSelfRegister: false, AllowRegistrarOverride: true,
				RegistrarAccounts: []string{"ixo1abcdefg0000000000000000000000000000000"},
				MinLength:         3, MaxLength: 32},
		},
		Names: []types.NameRecord{
			{Namespace: "handles", NormalizedName: "alice", DisplayName: "Alice", OwnerDid: "did:ixo:a", Status: types.NAME_STATUS_ACTIVE},
			{Namespace: "handles", NormalizedName: "bob", DisplayName: "bob", OwnerDid: "did:ixo:b", Status: types.NAME_STATUS_ACTIVE},
			{Namespace: "verified", NormalizedName: "carol", DisplayName: "Carol", OwnerDid: "did:ixo:c", Status: types.NAME_STATUS_ACTIVE, Verified: true, VerifiedBy: "ixo1abcdefg0000000000000000000000000000000", Source: "linkedin"},
		},
	}

	s.App.NamesKeeper.InitGenesis(s.Ctx, gs)
	exported := s.App.NamesKeeper.ExportGenesis(s.Ctx)

	s.Require().Len(exported.Namespaces, len(gs.Namespaces))
	s.Require().Len(exported.Names, len(gs.Names))

	// Compare as sets — store iteration order is deterministic by key but we
	// shouldn't rely on the slice ordering matching the source slice.
	exportedNS := map[string]types.Namespace{}
	for _, ns := range exported.Namespaces {
		exportedNS[ns.Name] = ns
	}
	for _, ns := range gs.Namespaces {
		s.Require().Equal(ns, exportedNS[ns.Name])
	}

	type rk struct{ ns, n string }
	exportedRec := map[rk]types.NameRecord{}
	for _, r := range exported.Names {
		exportedRec[rk{r.Namespace, r.NormalizedName}] = r
	}
	for _, r := range gs.Names {
		got := exportedRec[rk{r.Namespace, r.NormalizedName}]
		s.Require().Equal(r.OwnerDid, got.OwnerDid)
		s.Require().Equal(r.Status, got.Status)
		s.Require().Equal(r.Verified, got.Verified)
	}
}

// TestGenesis_DefaultIsEmpty confirms DefaultGenesisState returns an empty,
// validatable genesis. This is what InitChain consumes when no module-level
// genesis is supplied for x/names.
func (s *KeeperTestSuite) TestGenesis_DefaultIsEmpty() {
	gs := types.DefaultGenesisState()
	s.Require().NotNil(gs)
	s.Require().Empty(gs.Namespaces)
	s.Require().Empty(gs.Names)
	s.Require().NoError(gs.Validate())
}

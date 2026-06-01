package keeper_test

import (
	"strings"

	"github.com/ixofoundation/ixo-blockchain/v7/app/apptesting"
	"github.com/ixofoundation/ixo-blockchain/v7/x/names/types"
)

// ----------------------------------------------------------------------------
// CreateNamespace
// ----------------------------------------------------------------------------

func (s *KeeperTestSuite) TestMsgCreateNamespace() {
	tcs := []struct {
		name      string
		mutate    func(m *types.MsgCreateNamespace)
		expErr    bool
		expErrSub string
	}{
		{
			name:   "happy path: gov authority creates a self-register namespace",
			mutate: func(m *types.MsgCreateNamespace) {},
			expErr: false,
		},
		{
			name: "wrong authority is rejected",
			mutate: func(m *types.MsgCreateNamespace) {
				m.Authority = apptesting.RandomAccountAddress().String()
			},
			expErr:    true,
			expErrSub: "invalid authority",
		},
		{
			name: "nil namespace is rejected",
			mutate: func(m *types.MsgCreateNamespace) {
				m.Namespace = nil
			},
			expErr:    true,
			expErrSub: "namespace is required",
		},
		{
			name: "duplicate namespace is rejected",
			mutate: func(m *types.MsgCreateNamespace) {
				// pre-populate
				s.mustCreateNamespace(*m.Namespace)
			},
			expErr:    true,
			expErrSub: "already exists",
		},
		{
			name: "invalid namespace (no self-register, no registrar) is rejected",
			mutate: func(m *types.MsgCreateNamespace) {
				m.Namespace.AllowSelfRegister = false
				m.Namespace.RegistrarAccounts = nil
			},
			expErr:    true,
			expErrSub: "must allow self-registration",
		},
		{
			name: "max_length=0 is rejected",
			mutate: func(m *types.MsgCreateNamespace) {
				m.Namespace.MaxLength = 0
			},
			expErr:    true,
			expErrSub: "max_length must be > 0",
		},
	}

	for _, tc := range tcs {
		s.Run(tc.name, func() {
			s.SetupTest()
			ns := s.defaultNamespace("handles")
			msg := &types.MsgCreateNamespace{
				Authority: s.authority,
				Namespace: &ns,
			}
			tc.mutate(msg)

			_, err := s.msgServer.CreateNamespace(s.goCtx(), msg)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(err.Error(), tc.expErrSub)
				return
			}
			s.Require().NoError(err)
			// state assertion: namespace exists post-success
			got, found := s.App.NamesKeeper.GetNamespace(s.Ctx, msg.Namespace.Name)
			s.Require().True(found)
			s.Require().Equal(msg.Namespace.Name, got.Name)
			s.AssertEventEmitted(s.Ctx, "ixo.names.v1beta1.NamespaceCreatedEvent", 1)
		})
	}
}

// ----------------------------------------------------------------------------
// UpdateNamespace
// ----------------------------------------------------------------------------

func (s *KeeperTestSuite) TestMsgUpdateNamespace() {
	tcs := []struct {
		name      string
		seed      bool
		mutate    func(m *types.MsgUpdateNamespace)
		expErr    bool
		expErrSub string
	}{
		{
			name:   "happy path: change description",
			seed:   true,
			mutate: func(m *types.MsgUpdateNamespace) { m.Namespace.Description = "updated" },
			expErr: false,
		},
		{
			name:      "missing namespace is rejected",
			seed:      false,
			mutate:    func(*types.MsgUpdateNamespace) {},
			expErr:    true,
			expErrSub: "does not exist",
		},
		{
			name:      "wrong authority is rejected",
			seed:      true,
			mutate:    func(m *types.MsgUpdateNamespace) { m.Authority = apptesting.RandomAccountAddress().String() },
			expErr:    true,
			expErrSub: "invalid authority",
		},
	}

	for _, tc := range tcs {
		s.Run(tc.name, func() {
			s.SetupTest()
			ns := s.defaultNamespace("handles")
			if tc.seed {
				s.mustCreateNamespace(ns)
			}

			toSend := ns
			toSend.Description = "test namespace v2"
			msg := &types.MsgUpdateNamespace{
				Authority: s.authority,
				Namespace: &toSend,
			}
			tc.mutate(msg)

			_, err := s.msgServer.UpdateNamespace(s.goCtx(), msg)
			if tc.expErr {
				s.Require().Error(err)
				s.Require().Contains(err.Error(), tc.expErrSub)
				return
			}
			s.Require().NoError(err)
			got, found := s.App.NamesKeeper.GetNamespace(s.Ctx, msg.Namespace.Name)
			s.Require().True(found)
			s.Require().Equal(msg.Namespace.Description, got.Description)
		})
	}
}

// ----------------------------------------------------------------------------
// RegisterName (self-registration)
// ----------------------------------------------------------------------------

func (s *KeeperTestSuite) TestMsgRegisterName() {
	s.Run("happy path: signer registers name in self-register namespace", func() {
		s.SetupTest()
		s.mustCreateNamespace(s.defaultNamespace("handles"))
		signer, did := s.addrAndDID()

		resp, err := s.msgServer.RegisterName(s.goCtx(), &types.MsgRegisterName{
			Signer:    signer.String(),
			Namespace: "handles",
			Name:      "Alice",
			OwnerDid:  did,
		})
		s.Require().NoError(err)
		s.Require().Equal("alice", resp.NormalizedName) // normalised form

		rec, found := s.App.NamesKeeper.GetNameRecord(s.Ctx, "handles", "alice")
		s.Require().True(found)
		s.Require().Equal(did, rec.OwnerDid)
		s.Require().Equal(types.NAME_STATUS_ACTIVE, rec.Status)
		s.Require().False(rec.Verified, "self-registered names start unverified")
		s.AssertEventEmitted(s.Ctx, "ixo.names.v1beta1.NameRegisteredEvent", 1)
	})

	s.Run("namespace missing is rejected", func() {
		s.SetupTest()
		signer, did := s.addrAndDID()
		_, err := s.msgServer.RegisterName(s.goCtx(), &types.MsgRegisterName{
			Signer: signer.String(), Namespace: "handles", Name: "alice", OwnerDid: did,
		})
		s.Require().ErrorContains(err, "does not exist")
	})

	s.Run("namespace forbids self-register is rejected", func() {
		s.SetupTest()
		registrar := apptesting.RandomAccountAddress()
		s.mustCreateNamespace(s.registrarNamespace("verified", registrar.String()))
		signer, did := s.addrAndDID()
		_, err := s.msgServer.RegisterName(s.goCtx(), &types.MsgRegisterName{
			Signer: signer.String(), Namespace: "verified", Name: "alice", OwnerDid: did,
		})
		s.Require().ErrorContains(err, "registrar-only")
	})

	s.Run("signer does not control DID is rejected", func() {
		s.SetupTest()
		s.mustCreateNamespace(s.defaultNamespace("handles"))
		signer := apptesting.RandomAccountAddress()
		did := s.onlyDID() // not controlled by signer
		_, err := s.msgServer.RegisterName(s.goCtx(), &types.MsgRegisterName{
			Signer: signer.String(), Namespace: "handles", Name: "alice", OwnerDid: did,
		})
		s.Require().ErrorContains(err, "does not control")
	})

	s.Run("name shorter than namespace.MinLength is rejected", func() {
		s.SetupTest()
		ns := s.defaultNamespace("handles")
		ns.MinLength = 5
		s.mustCreateNamespace(ns)
		signer, did := s.addrAndDID()
		_, err := s.msgServer.RegisterName(s.goCtx(), &types.MsgRegisterName{
			Signer: signer.String(), Namespace: "handles", Name: "abc", OwnerDid: did,
		})
		s.Require().ErrorContains(err, "shorter than namespace min_length")
	})

	s.Run("duplicate name in same namespace is rejected", func() {
		s.SetupTest()
		s.mustCreateNamespace(s.defaultNamespace("handles"))
		signer, did := s.addrAndDID()
		_, err := s.msgServer.RegisterName(s.goCtx(), &types.MsgRegisterName{
			Signer: signer.String(), Namespace: "handles", Name: "alice", OwnerDid: did,
		})
		s.Require().NoError(err)
		_, err = s.msgServer.RegisterName(s.goCtx(), &types.MsgRegisterName{
			Signer: signer.String(), Namespace: "handles", Name: "ALICE", OwnerDid: did,
		})
		s.Require().ErrorContains(err, "already taken", "case-insensitive uniqueness must hold")
	})
}

// ----------------------------------------------------------------------------
// RegisterNameByRegistrar
// ----------------------------------------------------------------------------

func (s *KeeperTestSuite) TestMsgRegisterNameByRegistrar() {
	s.Run("happy path: registrar registers verified name", func() {
		s.SetupTest()
		registrar := apptesting.RandomAccountAddress()
		s.mustCreateNamespace(s.registrarNamespace("verified", registrar.String()))
		_, did := s.addrAndDID()

		_, err := s.msgServer.RegisterNameByRegistrar(s.goCtx(), &types.MsgRegisterNameByRegistrar{
			Registrar:    registrar.String(),
			Namespace:    "verified",
			Name:         "Alice",
			OwnerDid:     did,
			Verified:     true,
			EvidenceHash: "sha256:abcd",
			Source:       "twitter-oauth",
		})
		s.Require().NoError(err)

		rec, found := s.App.NamesKeeper.GetNameRecord(s.Ctx, "verified", "alice")
		s.Require().True(found)
		s.Require().Equal(did, rec.OwnerDid)
		s.Require().True(rec.Verified)
		s.Require().Equal(registrar.String(), rec.VerifiedBy)
		s.Require().Equal("sha256:abcd", rec.EvidenceHash)
	})

	s.Run("non-registrar is rejected", func() {
		s.SetupTest()
		registrar := apptesting.RandomAccountAddress()
		s.mustCreateNamespace(s.registrarNamespace("verified", registrar.String()))
		_, did := s.addrAndDID()

		stranger := apptesting.RandomAccountAddress()
		_, err := s.msgServer.RegisterNameByRegistrar(s.goCtx(), &types.MsgRegisterNameByRegistrar{
			Registrar: stranger.String(), Namespace: "verified", Name: "alice", OwnerDid: did,
		})
		s.Require().ErrorContains(err, "is not a registrar")
	})

	s.Run("missing owner DID is rejected", func() {
		s.SetupTest()
		registrar := apptesting.RandomAccountAddress()
		s.mustCreateNamespace(s.registrarNamespace("verified", registrar.String()))

		_, err := s.msgServer.RegisterNameByRegistrar(s.goCtx(), &types.MsgRegisterNameByRegistrar{
			Registrar: registrar.String(), Namespace: "verified", Name: "alice",
			OwnerDid: "did:ixo:does-not-exist",
		})
		s.Require().ErrorContains(err, "not found")
	})

	s.Run("evidence hash exceeds cap is rejected (defence-in-depth path)", func() {
		s.SetupTest()
		registrar := apptesting.RandomAccountAddress()
		s.mustCreateNamespace(s.registrarNamespace("verified", registrar.String()))
		_, did := s.addrAndDID()

		_, err := s.msgServer.RegisterNameByRegistrar(s.goCtx(), &types.MsgRegisterNameByRegistrar{
			Registrar:    registrar.String(),
			Namespace:    "verified",
			Name:         "alice",
			OwnerDid:     did,
			EvidenceHash: strings.Repeat("a", types.MaxNameRecordEvidenceHashLength+1),
		})
		s.Require().ErrorContains(err, "evidence_hash longer than")
	})
}

// ----------------------------------------------------------------------------
// UpdateNameByRegistrar
// ----------------------------------------------------------------------------

func (s *KeeperTestSuite) TestMsgUpdateNameByRegistrar() {
	s.Run("happy path: registrar flips verified flag and stamps metadata", func() {
		s.SetupTest()
		registrar := apptesting.RandomAccountAddress()
		s.mustCreateNamespace(s.registrarNamespace("verified", registrar.String()))
		_, did := s.addrAndDID()

		// arrange: existing record
		_, err := s.msgServer.RegisterNameByRegistrar(s.goCtx(), &types.MsgRegisterNameByRegistrar{
			Registrar: registrar.String(), Namespace: "verified", Name: "alice", OwnerDid: did,
		})
		s.Require().NoError(err)

		_, err = s.msgServer.UpdateNameByRegistrar(s.goCtx(), &types.MsgUpdateNameByRegistrar{
			Registrar: registrar.String(),
			Namespace: "verified", NormalizedName: "alice",
			Verified: true, EvidenceHash: "sha256:e1", Source: "linkedin",
		})
		s.Require().NoError(err)
		rec, _ := s.App.NamesKeeper.GetNameRecord(s.Ctx, "verified", "alice")
		s.Require().True(rec.Verified)
		s.Require().Equal("sha256:e1", rec.EvidenceHash)
		s.Require().Equal("linkedin", rec.Source)
	})

	s.Run("record missing is rejected", func() {
		s.SetupTest()
		registrar := apptesting.RandomAccountAddress()
		s.mustCreateNamespace(s.registrarNamespace("verified", registrar.String()))
		_, err := s.msgServer.UpdateNameByRegistrar(s.goCtx(), &types.MsgUpdateNameByRegistrar{
			Registrar: registrar.String(),
			Namespace: "verified", NormalizedName: "ghost",
		})
		s.Require().ErrorContains(err, "name \"ghost\" in namespace \"verified\" not found")
	})

	s.Run("non-registrar is rejected", func() {
		s.SetupTest()
		registrar := apptesting.RandomAccountAddress()
		s.mustCreateNamespace(s.registrarNamespace("verified", registrar.String()))
		_, did := s.addrAndDID()
		_, err := s.msgServer.RegisterNameByRegistrar(s.goCtx(), &types.MsgRegisterNameByRegistrar{
			Registrar: registrar.String(), Namespace: "verified", Name: "alice", OwnerDid: did,
		})
		s.Require().NoError(err)

		stranger := apptesting.RandomAccountAddress()
		_, err = s.msgServer.UpdateNameByRegistrar(s.goCtx(), &types.MsgUpdateNameByRegistrar{
			Registrar: stranger.String(), Namespace: "verified", NormalizedName: "alice",
		})
		s.Require().ErrorContains(err, "is not a registrar")
	})
}

// ----------------------------------------------------------------------------
// TransferName
// ----------------------------------------------------------------------------

func (s *KeeperTestSuite) TestMsgTransferName() {
	s.Run("happy path: owner transfers to existing DID", func() {
		s.SetupTest()
		s.mustCreateNamespace(s.defaultNamespace("handles"))
		signer, did := s.addrAndDID()
		_, err := s.msgServer.RegisterName(s.goCtx(), &types.MsgRegisterName{
			Signer: signer.String(), Namespace: "handles", Name: "alice", OwnerDid: did,
		})
		s.Require().NoError(err)

		newOwner := s.onlyDID()
		_, err = s.msgServer.TransferName(s.goCtx(), &types.MsgTransferName{
			Signer:         signer.String(),
			Namespace:      "handles",
			NormalizedName: "alice",
			NewOwnerDid:    newOwner,
		})
		s.Require().NoError(err)

		rec, _ := s.App.NamesKeeper.GetNameRecord(s.Ctx, "handles", "alice")
		s.Require().Equal(newOwner, rec.OwnerDid)
		s.AssertEventEmitted(s.Ctx, "ixo.names.v1beta1.NameTransferredEvent", 1)
	})

	s.Run("registrar override succeeds when allowed", func() {
		s.SetupTest()
		registrar := apptesting.RandomAccountAddress()
		s.mustCreateNamespace(s.registrarNamespace("verified", registrar.String()))
		_, did := s.addrAndDID()
		_, err := s.msgServer.RegisterNameByRegistrar(s.goCtx(), &types.MsgRegisterNameByRegistrar{
			Registrar: registrar.String(), Namespace: "verified", Name: "alice", OwnerDid: did,
		})
		s.Require().NoError(err)

		newOwner := s.onlyDID()
		_, err = s.msgServer.TransferName(s.goCtx(), &types.MsgTransferName{
			Signer:         registrar.String(),
			Namespace:      "verified",
			NormalizedName: "alice",
			NewOwnerDid:    newOwner,
		})
		s.Require().NoError(err)
	})

	s.Run("non-owner non-registrar is rejected", func() {
		s.SetupTest()
		s.mustCreateNamespace(s.defaultNamespace("handles"))
		signer, did := s.addrAndDID()
		_, err := s.msgServer.RegisterName(s.goCtx(), &types.MsgRegisterName{
			Signer: signer.String(), Namespace: "handles", Name: "alice", OwnerDid: did,
		})
		s.Require().NoError(err)

		stranger := apptesting.RandomAccountAddress()
		newOwner := s.onlyDID()
		_, err = s.msgServer.TransferName(s.goCtx(), &types.MsgTransferName{
			Signer: stranger.String(), Namespace: "handles",
			NormalizedName: "alice", NewOwnerDid: newOwner,
		})
		s.Require().ErrorContains(err, "not the current owner")
	})

	s.Run("transfer to same owner is rejected", func() {
		s.SetupTest()
		s.mustCreateNamespace(s.defaultNamespace("handles"))
		signer, did := s.addrAndDID()
		_, err := s.msgServer.RegisterName(s.goCtx(), &types.MsgRegisterName{
			Signer: signer.String(), Namespace: "handles", Name: "alice", OwnerDid: did,
		})
		s.Require().NoError(err)
		_, err = s.msgServer.TransferName(s.goCtx(), &types.MsgTransferName{
			Signer: signer.String(), Namespace: "handles",
			NormalizedName: "alice", NewOwnerDid: did,
		})
		s.Require().ErrorContains(err, "same as current owner")
	})

	s.Run("transfer to nonexistent DID is rejected", func() {
		s.SetupTest()
		s.mustCreateNamespace(s.defaultNamespace("handles"))
		signer, did := s.addrAndDID()
		_, err := s.msgServer.RegisterName(s.goCtx(), &types.MsgRegisterName{
			Signer: signer.String(), Namespace: "handles", Name: "alice", OwnerDid: did,
		})
		s.Require().NoError(err)
		_, err = s.msgServer.TransferName(s.goCtx(), &types.MsgTransferName{
			Signer: signer.String(), Namespace: "handles",
			NormalizedName: "alice", NewOwnerDid: "did:ixo:never-existed",
		})
		s.Require().ErrorContains(err, "not found")
	})
}

// ----------------------------------------------------------------------------
// SetNameStatus
// ----------------------------------------------------------------------------

func (s *KeeperTestSuite) TestMsgSetNameStatus() {
	s.Run("happy path: registrar suspends a name", func() {
		s.SetupTest()
		registrar := apptesting.RandomAccountAddress()
		s.mustCreateNamespace(s.registrarNamespace("verified", registrar.String()))
		_, did := s.addrAndDID()
		_, err := s.msgServer.RegisterNameByRegistrar(s.goCtx(), &types.MsgRegisterNameByRegistrar{
			Registrar: registrar.String(), Namespace: "verified", Name: "alice", OwnerDid: did,
		})
		s.Require().NoError(err)
		_, err = s.msgServer.SetNameStatus(s.goCtx(), &types.MsgSetNameStatus{
			Signer: registrar.String(), Namespace: "verified",
			NormalizedName: "alice", Status: types.NAME_STATUS_SUSPENDED, Reason: "abuse",
		})
		s.Require().NoError(err)
		rec, _ := s.App.NamesKeeper.GetNameRecord(s.Ctx, "verified", "alice")
		s.Require().Equal(types.NAME_STATUS_SUSPENDED, rec.Status)
		s.AssertEventEmitted(s.Ctx, "ixo.names.v1beta1.NameStatusChangedEvent", 1)
	})

	s.Run("gov authority can change status without being a registrar", func() {
		s.SetupTest()
		registrar := apptesting.RandomAccountAddress()
		s.mustCreateNamespace(s.registrarNamespace("verified", registrar.String()))
		_, did := s.addrAndDID()
		_, err := s.msgServer.RegisterNameByRegistrar(s.goCtx(), &types.MsgRegisterNameByRegistrar{
			Registrar: registrar.String(), Namespace: "verified", Name: "alice", OwnerDid: did,
		})
		s.Require().NoError(err)
		_, err = s.msgServer.SetNameStatus(s.goCtx(), &types.MsgSetNameStatus{
			Signer: s.authority, Namespace: "verified",
			NormalizedName: "alice", Status: types.NAME_STATUS_REVOKED,
		})
		s.Require().NoError(err)
	})

	s.Run("tombstoned is terminal", func() {
		s.SetupTest()
		registrar := apptesting.RandomAccountAddress()
		s.mustCreateNamespace(s.registrarNamespace("verified", registrar.String()))
		_, did := s.addrAndDID()
		_, err := s.msgServer.RegisterNameByRegistrar(s.goCtx(), &types.MsgRegisterNameByRegistrar{
			Registrar: registrar.String(), Namespace: "verified", Name: "alice", OwnerDid: did,
		})
		s.Require().NoError(err)
		_, err = s.msgServer.SetNameStatus(s.goCtx(), &types.MsgSetNameStatus{
			Signer: registrar.String(), Namespace: "verified",
			NormalizedName: "alice", Status: types.NAME_STATUS_TOMBSTONED,
		})
		s.Require().NoError(err)
		_, err = s.msgServer.SetNameStatus(s.goCtx(), &types.MsgSetNameStatus{
			Signer: registrar.String(), Namespace: "verified",
			NormalizedName: "alice", Status: types.NAME_STATUS_ACTIVE,
		})
		s.Require().ErrorContains(err, "tombstoned records are terminal")
	})

	s.Run("no-op transition (same status) is rejected", func() {
		s.SetupTest()
		registrar := apptesting.RandomAccountAddress()
		s.mustCreateNamespace(s.registrarNamespace("verified", registrar.String()))
		_, did := s.addrAndDID()
		_, err := s.msgServer.RegisterNameByRegistrar(s.goCtx(), &types.MsgRegisterNameByRegistrar{
			Registrar: registrar.String(), Namespace: "verified", Name: "alice", OwnerDid: did,
		})
		s.Require().NoError(err)
		_, err = s.msgServer.SetNameStatus(s.goCtx(), &types.MsgSetNameStatus{
			Signer: registrar.String(), Namespace: "verified",
			NormalizedName: "alice", Status: types.NAME_STATUS_ACTIVE,
		})
		s.Require().ErrorContains(err, "already set to the requested value")
	})

	s.Run("non-registrar non-authority is rejected", func() {
		s.SetupTest()
		registrar := apptesting.RandomAccountAddress()
		s.mustCreateNamespace(s.registrarNamespace("verified", registrar.String()))
		_, did := s.addrAndDID()
		_, err := s.msgServer.RegisterNameByRegistrar(s.goCtx(), &types.MsgRegisterNameByRegistrar{
			Registrar: registrar.String(), Namespace: "verified", Name: "alice", OwnerDid: did,
		})
		s.Require().NoError(err)
		stranger := apptesting.RandomAccountAddress()
		_, err = s.msgServer.SetNameStatus(s.goCtx(), &types.MsgSetNameStatus{
			Signer: stranger.String(), Namespace: "verified",
			NormalizedName: "alice", Status: types.NAME_STATUS_REVOKED,
		})
		s.Require().ErrorContains(err, "must be a registrar or the governance authority")
	})

	s.Run("invalid status enum value is rejected", func() {
		s.SetupTest()
		registrar := apptesting.RandomAccountAddress()
		s.mustCreateNamespace(s.registrarNamespace("verified", registrar.String()))
		_, err := s.msgServer.SetNameStatus(s.goCtx(), &types.MsgSetNameStatus{
			Signer: registrar.String(), Namespace: "verified",
			NormalizedName: "alice", Status: types.NameStatus(999),
		})
		s.Require().ErrorContains(err, "unsupported target status")
	})
}

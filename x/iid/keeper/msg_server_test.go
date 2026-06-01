package keeper_test

import (
	"github.com/ixofoundation/ixo-blockchain/v7/app/apptesting"
	"github.com/ixofoundation/ixo-blockchain/v7/x/iid/types"
)

// ----------------------------------------------------------------------------
// CreateIidDocument
// ----------------------------------------------------------------------------

func (s *KeeperTestSuite) TestMsgCreateIidDocument() {
	s.SetupTest()
	signer := apptesting.RandomAccountAddress()
	did := "did:ixo:" + signer.String()

	_, err := s.msgServer.CreateIidDocument(s.goCtx(), &types.MsgCreateIidDocument{
		Id:     did,
		Signer: signer.String(),
	})
	s.Require().NoError(err)
	got, found := s.App.IidKeeper.GetDidDocument(s.Ctx, []byte(did))
	s.Require().True(found)
	s.Require().Equal(did, got.Id)

	// Duplicate is rejected.
	_, err = s.msgServer.CreateIidDocument(s.goCtx(), &types.MsgCreateIidDocument{
		Id:     did,
		Signer: signer.String(),
	})
	s.Require().ErrorContains(err, "already exists")
}

func (s *KeeperTestSuite) TestMsgCreateIidDocument_InvalidDID() {
	s.SetupTest()
	signer := apptesting.RandomAccountAddress()
	_, err := s.msgServer.CreateIidDocument(s.goCtx(), &types.MsgCreateIidDocument{
		Id:     "not-a-did",
		Signer: signer.String(),
	})
	s.Require().ErrorContains(err, "did")
}

// Reserved-namespace guard: handler must reject any DID under a module-
// reserved prefix even when ValidateBasic is skipped (direct keeper call /
// gov-proposal path). Matches ca41cb4d feat(iid): block module-reserved
// DID namespaces on MsgCreateIidDocument.
func (s *KeeperTestSuite) TestMsgCreateIidDocument_ReservedNamespace() {
	s.SetupTest()

	signer := apptesting.RandomAccountAddress()
	for _, prefix := range types.ReservedDidPrefixes {
		reserved := prefix + "squatted-id-001"
		_, err := s.msgServer.CreateIidDocument(s.goCtx(), &types.MsgCreateIidDocument{
			Id:     reserved,
			Signer: signer.String(),
		})
		s.Require().ErrorIs(err, types.ErrReservedDidNamespace,
			"handler must reject reserved prefix %q (got %v)", prefix, err)
		_, found := s.App.IidKeeper.GetDidDocument(s.Ctx, []byte(reserved))
		s.Require().False(found, "rejected create must not persist a DID under prefix %q", prefix)
	}
}

// IXO-2045: MsgCreateIidDocument must enforce the DID-form policy at the
// handler — not just in ValidateBasic — so Cosmwasm contracts (which
// bypass ValidateBasic by routing through Stargate / the MsgServiceRouter
// from their wasm-vm) cannot create DIDs they should not own. The same
// policy lives in types.ValidateMsgCreateDIDForm; this suite asserts the
// keeper actually invokes it. Each subcase pins the matching error so a
// future refactor that drops one of the checks fails loudly.
func (s *KeeperTestSuite) TestMsgCreateIidDocument_DIDFormPolicy() {
	s.SetupTest()
	signer := apptesting.RandomAccountAddress()
	other := apptesting.RandomAccountAddress()
	contract := apptesting.RandomAccountAddress()

	// --- happy paths ---
	cases := []struct {
		name string
		id   string
	}{
		// did:ixo:<signer> — the signer registers their own account DID.
		{"signer-owned account DID", "did:ixo:" + signer.String()},
		// did:ixo:wasm:<contract> — any signer may create a wasm-form DID
		// (for now) as long as the suffix is a valid bech32 address.
		{"wasm contract DID created by arbitrary signer",
			"did:ixo:wasm:" + contract.String()},
	}
	for _, tc := range cases {
		_, err := s.msgServer.CreateIidDocument(s.goCtx(), &types.MsgCreateIidDocument{
			Id:     tc.id,
			Signer: signer.String(),
		})
		s.Require().NoErrorf(err, "%s: %v", tc.name, err)
		_, found := s.App.IidKeeper.GetDidDocument(s.Ctx, []byte(tc.id))
		s.Require().True(found, "%s: did was not persisted", tc.name)
	}

	// --- rejections, each pinned to the matching sentinel error ---
	rejections := []struct {
		name    string
		id      string
		errLike error
	}{
		// did:ixo:<other-account> — caller cannot claim someone else's address.
		{"did:ixo:<other-account> rejected (account != signer)",
			"did:ixo:" + other.String(), types.ErrDIDAccountSignerMismatch},
		// did:ixo:wasm:<junk> — wasm suffix must be a valid bech32 address.
		{"did:ixo:wasm:<invalid> rejected (not bech32)",
			"did:ixo:wasm:not-a-bech32", types.ErrDIDFormNotAllowed},
		// did:ixo:wasm:<addr>:more — wasm suffix must be a single segment.
		{"did:ixo:wasm:<addr>:extra rejected (nested sub-method)",
			"did:ixo:wasm:" + contract.String() + ":extra", types.ErrDIDFormNotAllowed},
		// did:ixo:haha:haha — unknown sub-method.
		{"did:ixo:<unknown-submethod>:* rejected",
			"did:ixo:haha:haha", types.ErrDIDFormNotAllowed},
		// did:cosmos:foo — wrong DID method entirely.
		{"did:cosmos:* rejected (only did:ixo: allowed)",
			"did:cosmos:foo", types.ErrDIDFormNotAllowed},
		// did:x:ixo:foo — DidChainPrefix is not an allowed user namespace.
		{"did:x:* rejected (only did:ixo: allowed)",
			"did:x:ixo:abc123", types.ErrDIDFormNotAllowed},
	}
	for _, tc := range rejections {
		_, err := s.msgServer.CreateIidDocument(s.goCtx(), &types.MsgCreateIidDocument{
			Id:     tc.id,
			Signer: signer.String(),
		})
		s.Require().ErrorIsf(err, tc.errLike,
			"%s: expected %v, got %v", tc.name, tc.errLike, err)
		_, found := s.App.IidKeeper.GetDidDocument(s.Ctx, []byte(tc.id))
		s.Require().False(found, "%s: rejected DID must not be persisted", tc.name)
	}
}

// ----------------------------------------------------------------------------
// UpdateIidDocument
// ----------------------------------------------------------------------------

func (s *KeeperTestSuite) TestMsgUpdateIidDocument() {
	s.SetupTest()
	signer, did := s.freshSigner()

	_, err := s.msgServer.UpdateIidDocument(s.goCtx(), &types.MsgUpdateIidDocument{
		Id:          did,
		Context:     []*types.Context{{Key: "foo", Val: "bar"}},
		AlsoKnownAs: "alias",
		Signer:      signer.String(),
	})
	s.Require().NoError(err)
	got, _ := s.App.IidKeeper.GetDidDocument(s.Ctx, []byte(did))
	s.Require().Equal("alias", got.AlsoKnownAs)
}

func (s *KeeperTestSuite) TestMsgUpdateIidDocument_Unauthorized() {
	s.SetupTest()
	_, did := s.freshSigner()
	stranger := apptesting.RandomAccountAddress()
	_, err := s.msgServer.UpdateIidDocument(s.goCtx(), &types.MsgUpdateIidDocument{
		Id:     did,
		Signer: stranger.String(),
	})
	s.Require().ErrorContains(err, "not authorized")
}

// ----------------------------------------------------------------------------
// AddVerification / RevokeVerification / SetVerificationRelationships
// ----------------------------------------------------------------------------

func (s *KeeperTestSuite) TestMsgAddAndRevokeVerification() {
	s.SetupTest()
	signer, did := s.freshSigner()
	vm := types.NewVerificationMethod(
		did+"#k1", types.DID(did),
		types.NewBlockchainAccountID(signer.String()),
	)
	v := types.NewVerification(vm, []string{types.Authentication}, nil)

	_, err := s.msgServer.AddVerification(s.goCtx(), &types.MsgAddVerification{
		Id:           did,
		Verification: v,
		Signer:       signer.String(),
	})
	s.Require().NoError(err)
	got, _ := s.App.IidKeeper.GetDidDocument(s.Ctx, []byte(did))
	// fixture seeds one method (#key-1); after Add we should have two.
	s.Require().Len(got.VerificationMethod, 2)

	// Set verification relationships
	_, err = s.msgServer.SetVerificationRelationships(s.goCtx(), &types.MsgSetVerificationRelationships{
		Id:                       did,
		MethodId:                 did + "#k1",
		Relationships:            []string{types.Authentication, types.AssertionMethod},
		Signer:                   signer.String(),
	})
	s.Require().NoError(err)

	// Revoke
	_, err = s.msgServer.RevokeVerification(s.goCtx(), &types.MsgRevokeVerification{
		Id:       did,
		MethodId: did + "#k1",
		Signer:   signer.String(),
	})
	s.Require().NoError(err)
	got, _ = s.App.IidKeeper.GetDidDocument(s.Ctx, []byte(did))
	// The fixture seeds one #key-1 method; we added #k1 then revoked it, so
	// only the original key-1 remains.
	s.Require().Len(got.VerificationMethod, 1, "RevokeVerification removes only the named method")
}

// ----------------------------------------------------------------------------
// Service add/delete
// ----------------------------------------------------------------------------

func (s *KeeperTestSuite) TestMsgAddAndDeleteService() {
	s.SetupTest()
	signer, did := s.freshSigner()

	svc := types.NewService(did+"#svc1", "LinkedDomains", "https://example.com")
	_, err := s.msgServer.AddService(s.goCtx(), &types.MsgAddService{
		Id:          did,
		ServiceData: svc,
		Signer:      signer.String(),
	})
	s.Require().NoError(err)
	got, _ := s.App.IidKeeper.GetDidDocument(s.Ctx, []byte(did))
	s.Require().Len(got.Service, 1)

	// Verify the inner Service fields — this is the regression target
	// for the L3 "service: [{}]" empty-object issue. If fields persist
	// here but render as empty in the CLI, the bug is in the query/
	// CLI JSON layer, not in the keeper.
	s.Require().Equal(did+"#svc1", got.Service[0].Id, "Service.Id must persist")
	s.Require().Equal("LinkedDomains", got.Service[0].Type, "Service.Type must persist")
	s.Require().Equal("https://example.com", got.Service[0].ServiceEndpoint,
		"Service.ServiceEndpoint must persist through SetDidDocument/GetDidDocument round-trip")

	// And confirm the gRPC query handler returns the same populated
	// Service. If THIS fails, the bug lives in the query handler.
	resp, err := s.App.IidKeeper.IidDocument(s.goCtx(), &types.QueryIidDocumentRequest{Id: did})
	s.Require().NoError(err)
	s.Require().Len(resp.IidDocument.Service, 1, "query response must include the service")
	s.Require().Equal("https://example.com", resp.IidDocument.Service[0].ServiceEndpoint,
		"gRPC query handler must return Service.ServiceEndpoint populated")

	// Marshal the query response via the same codec the CLI uses and
	// confirm the JSON contains the endpoint. If it doesn't, the bug is
	// in the JSON encoder.
	bz, err := s.App.AppCodec().MarshalJSON(resp)
	s.Require().NoError(err)
	s.Require().Contains(string(bz), "https://example.com",
		"AppCodec.MarshalJSON must include nested Service.ServiceEndpoint; got %s", string(bz))

	_, err = s.msgServer.DeleteService(s.goCtx(), &types.MsgDeleteService{
		Id:        did,
		ServiceId: did + "#svc1",
		Signer:    signer.String(),
	})
	s.Require().NoError(err)
	got, _ = s.App.IidKeeper.GetDidDocument(s.Ctx, []byte(did))
	s.Require().Len(got.Service, 0)
}

// ----------------------------------------------------------------------------
// LinkedResource add/delete
// ----------------------------------------------------------------------------

func (s *KeeperTestSuite) TestMsgAddAndDeleteLinkedResource() {
	s.SetupTest()
	signer, did := s.freshSigner()

	r := types.NewLinkedResource("res1", "LinkedDataSet", "desc", "application/json", "https://r.example", "proof", "encrypted", "private")
	_, err := s.msgServer.AddLinkedResource(s.goCtx(), &types.MsgAddLinkedResource{
		Id:             did,
		LinkedResource: r,
		Signer:         signer.String(),
	})
	s.Require().NoError(err)
	got, _ := s.App.IidKeeper.GetDidDocument(s.Ctx, []byte(did))
	s.Require().Len(got.LinkedResource, 1)

	_, err = s.msgServer.DeleteLinkedResource(s.goCtx(), &types.MsgDeleteLinkedResource{
		Id:         did,
		ResourceId: "res1",
		Signer:     signer.String(),
	})
	s.Require().NoError(err)
	got, _ = s.App.IidKeeper.GetDidDocument(s.Ctx, []byte(did))
	s.Require().Len(got.LinkedResource, 0)
}

// ----------------------------------------------------------------------------
// LinkedClaim / LinkedEntity / AccordedRight / Context
// ----------------------------------------------------------------------------

func (s *KeeperTestSuite) TestMsgAddAndDeleteLinkedClaim() {
	s.SetupTest()
	signer, did := s.freshSigner()
	c := &types.LinkedClaim{
		Id:              "claim1",
		Type:            "test",
		Description:     "d",
		ServiceEndpoint: "https://c.example",
	}
	_, err := s.msgServer.AddLinkedClaim(s.goCtx(), &types.MsgAddLinkedClaim{
		Id: did, LinkedClaim: c, Signer: signer.String(),
	})
	s.Require().NoError(err)
	_, err = s.msgServer.DeleteLinkedClaim(s.goCtx(), &types.MsgDeleteLinkedClaim{
		Id: did, ClaimId: "claim1", Signer: signer.String(),
	})
	s.Require().NoError(err)
	got, _ := s.App.IidKeeper.GetDidDocument(s.Ctx, []byte(did))
	s.Require().Len(got.LinkedClaim, 0)
}

func (s *KeeperTestSuite) TestMsgAddAndDeleteLinkedEntity() {
	s.SetupTest()
	signer, did := s.freshSigner()
	e := &types.LinkedEntity{Id: "ent1", Type: "asset"}
	_, err := s.msgServer.AddLinkedEntity(s.goCtx(), &types.MsgAddLinkedEntity{
		Id: did, LinkedEntity: e, Signer: signer.String(),
	})
	s.Require().NoError(err)
	_, err = s.msgServer.DeleteLinkedEntity(s.goCtx(), &types.MsgDeleteLinkedEntity{
		Id: did, EntityId: "ent1", Signer: signer.String(),
	})
	s.Require().NoError(err)
	got, _ := s.App.IidKeeper.GetDidDocument(s.Ctx, []byte(did))
	s.Require().Len(got.LinkedEntity, 0)
}

func (s *KeeperTestSuite) TestMsgAddAndDeleteAccordedRight() {
	s.SetupTest()
	signer, did := s.freshSigner()
	r := types.NewAccordedRight("right1", "vote", "ballot", "msg", "https://x.example")
	_, err := s.msgServer.AddAccordedRight(s.goCtx(), &types.MsgAddAccordedRight{
		Id: did, AccordedRight: r, Signer: signer.String(),
	})
	s.Require().NoError(err)
	_, err = s.msgServer.DeleteAccordedRight(s.goCtx(), &types.MsgDeleteAccordedRight{
		Id: did, RightId: "right1", Signer: signer.String(),
	})
	s.Require().NoError(err)
}

func (s *KeeperTestSuite) TestMsgAddAndDeleteIidContext() {
	s.SetupTest()
	signer, did := s.freshSigner()
	c := &types.Context{Key: "ctx1", Val: "https://schema.org/"}
	_, err := s.msgServer.AddIidContext(s.goCtx(), &types.MsgAddIidContext{
		Id: did, Context: c, Signer: signer.String(),
	})
	s.Require().NoError(err)
	_, err = s.msgServer.DeleteIidContext(s.goCtx(), &types.MsgDeleteIidContext{
		Id: did, ContextKey: "ctx1", Signer: signer.String(),
	})
	s.Require().NoError(err)
}

// ----------------------------------------------------------------------------
// Controllers: AddController / DeleteController
// ----------------------------------------------------------------------------

func (s *KeeperTestSuite) TestMsgAddAndDeleteController() {
	s.SetupTest()
	signer, did := s.freshSigner()
	other := apptesting.RandomAccountAddress()
	otherDID := s.createDIDFor(other)

	_, err := s.msgServer.AddController(s.goCtx(), &types.MsgAddController{
		Id: did, ControllerDid: otherDID, Signer: signer.String(),
	})
	s.Require().NoError(err)
	got, _ := s.App.IidKeeper.GetDidDocument(s.Ctx, []byte(did))
	s.Require().Contains(got.Controller, otherDID)

	_, err = s.msgServer.DeleteController(s.goCtx(), &types.MsgDeleteController{
		Id: did, ControllerDid: otherDID, Signer: signer.String(),
	})
	s.Require().NoError(err)
}

// ----------------------------------------------------------------------------
// DeactivateIID
// ----------------------------------------------------------------------------

func (s *KeeperTestSuite) TestMsgDeactivateIID() {
	s.SetupTest()
	signer, did := s.freshSigner()
	_, err := s.msgServer.DeactivateIID(s.goCtx(), &types.MsgDeactivateIID{
		Id:     did,
		State:  true,
		Signer: signer.String(),
	})
	s.Require().NoError(err)
	got, _ := s.App.IidKeeper.GetDidDocument(s.Ctx, []byte(did))
	s.Require().True(got.Metadata.Deactivated)
}

// ----------------------------------------------------------------------------
// Document not found and unauthorized signer paths (cover the common
// ExecuteOnDidWithRelationships error branches once for the whole module).
// ----------------------------------------------------------------------------

func (s *KeeperTestSuite) TestUpdate_DocumentNotFound() {
	s.SetupTest()
	signer := apptesting.RandomAccountAddress()
	_, err := s.msgServer.AddService(s.goCtx(), &types.MsgAddService{
		Id:          "did:ixo:nope",
		ServiceData: types.NewService("svc", "LinkedDomains", "https://x.example"),
		Signer:      signer.String(),
	})
	s.Require().ErrorContains(err, "not found")
}

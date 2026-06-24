package entity_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/ixofoundation/ixo-blockchain/v8/app/apptesting"
	entitytypes "github.com/ixofoundation/ixo-blockchain/v8/x/entity/types"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v8/x/iid/types"
)

// EntityNoAnteTestSuite dispatches entity messages through the app's REAL
// MsgServiceRouter via RunMsg, which deliberately skips the ante handler. This
// reproduces the CosmWasm / ICA-host / authz routes (none of which run the IID
// ante) and asserts that the entity KEEPER itself rejects a signer that does
// not control the acting DID — i.e. authorization does not depend on the ante.
//
// This is the standing regression guard for the 2026-06 audit class: any future
// IidTxMsg handler that forgets its in-keeper signer→DID check will fail here.
type EntityNoAnteTestSuite struct {
	apptesting.KeeperTestHelper
}

func TestEntityNoAnteTestSuite(t *testing.T) { suite.Run(t, new(EntityNoAnteTestSuite)) }

func (s *EntityNoAnteTestSuite) SetupTest() { s.Setup() }

// seedControllerDID registers a DID whose only authentication key is `addr`.
func (s *EntityNoAnteTestSuite) seedControllerDID(did string, addr sdk.AccAddress) {
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

// seedEntity creates an entity DID controlled by controllerDID plus its record.
func (s *EntityNoAnteTestSuite) seedEntity(entityDID, controllerDID string) {
	now := s.Ctx.BlockTime()
	idoc := iidtypes.IidDocument{
		Id:         entityDID,
		Controller: []string{entityDID, controllerDID},
		Metadata:   &iidtypes.IidMetadata{Created: &now, VersionId: "v1"},
	}
	s.App.IidKeeper.SetDidDocument(s.Ctx, []byte(entityDID), idoc)
	entityMeta := entitytypes.NewEntityMetadata(s.Ctx.TxBytes(), now)
	s.App.EntityKeeper.SetEntity(s.Ctx, []byte(entityDID), entitytypes.Entity{
		Id:        entityDID,
		Type:      "asset",
		Status:    1,
		Metadata:  &entityMeta,
		StartDate: &now,
	})
}

// TestUpdateEntity_NoAnte_RejectsUnauthorizedSigner is the core regression: a
// caller who knows the victim entity's public controller DID but does NOT
// control it must be rejected by the keeper even with the ante skipped.
func (s *EntityNoAnteTestSuite) TestUpdateEntity_NoAnte_RejectsUnauthorizedSigner() {
	victim := apptesting.RandomAccountAddress()
	attacker := apptesting.RandomAccountAddress()
	controllerDID := "did:ixo:victimcontroller"
	entityDID := "did:ixo:entity:noante-upd"

	s.seedControllerDID(controllerDID, victim)
	s.seedEntity(entityDID, controllerDID)

	now := s.Ctx.BlockTime()
	end := now.Add(time.Hour)

	// Attacker quotes the victim's controller DID but signs as themselves.
	_, err := s.RunMsg(&entitytypes.MsgUpdateEntity{
		Id:                entityDID,
		EntityStatus:      2,
		StartDate:         &now,
		EndDate:           &end,
		ControllerDid:     iidtypes.DIDFragment(controllerDID),
		ControllerAddress: attacker.String(),
	})
	s.Require().Error(err)
	s.Require().ErrorIs(err, iidtypes.ErrUnauthorized)

	// State unchanged.
	got, found := s.App.EntityKeeper.GetEntity(s.Ctx, []byte(entityDID))
	s.Require().True(found)
	s.Require().Equal(int32(1), got.Status)
}

// TestUpdateEntity_NoAnte_AllowsAuthorizedSigner confirms the fix does not
// over-reject: the legitimate controller succeeds on the same no-ante route.
func (s *EntityNoAnteTestSuite) TestUpdateEntity_NoAnte_AllowsAuthorizedSigner() {
	victim := apptesting.RandomAccountAddress()
	controllerDID := "did:ixo:goodcontroller"
	entityDID := "did:ixo:entity:noante-upd-ok"

	s.seedControllerDID(controllerDID, victim)
	s.seedEntity(entityDID, controllerDID)

	now := s.Ctx.BlockTime()
	end := now.Add(time.Hour)

	_, err := s.RunMsg(&entitytypes.MsgUpdateEntity{
		Id:                entityDID,
		EntityStatus:      2,
		StartDate:         &now,
		EndDate:           &end,
		ControllerDid:     iidtypes.DIDFragment(controllerDID),
		ControllerAddress: victim.String(),
	})
	s.Require().NoError(err)

	got, _ := s.App.EntityKeeper.GetEntity(s.Ctx, []byte(entityDID))
	s.Require().Equal(int32(2), got.Status)
}

// TestUpdateEntityVerified_NoAnte_RejectsUnauthorizedSigner covers the verified
// flag path: quoting the entity's pinned relayer DID is not enough — the signer
// must control it.
func (s *EntityNoAnteTestSuite) TestUpdateEntityVerified_NoAnte_RejectsUnauthorizedSigner() {
	relayer := apptesting.RandomAccountAddress()
	attacker := apptesting.RandomAccountAddress()
	relayerDID := "did:ixo:relayernode"
	entityDID := "did:ixo:entity:noante-verify"

	s.seedControllerDID(relayerDID, relayer)

	now := s.Ctx.BlockTime()
	idoc := iidtypes.IidDocument{Id: entityDID, Metadata: &iidtypes.IidMetadata{Created: &now, VersionId: "v1"}}
	s.App.IidKeeper.SetDidDocument(s.Ctx, []byte(entityDID), idoc)
	entityMeta := entitytypes.NewEntityMetadata(s.Ctx.TxBytes(), now)
	s.App.EntityKeeper.SetEntity(s.Ctx, []byte(entityDID), entitytypes.Entity{
		Id: entityDID, Type: "asset", Status: 1, Metadata: &entityMeta, RelayerNode: relayerDID,
	})

	_, err := s.RunMsg(&entitytypes.MsgUpdateEntityVerified{
		Id:                 entityDID,
		EntityVerified:     true,
		RelayerNodeDid:     iidtypes.DIDFragment(relayerDID),
		RelayerNodeAddress: attacker.String(),
	})
	s.Require().Error(err)
	s.Require().ErrorIs(err, iidtypes.ErrUnauthorized)

	got, _ := s.App.EntityKeeper.GetEntity(s.Ctx, []byte(entityDID))
	s.Require().False(got.EntityVerified)
}

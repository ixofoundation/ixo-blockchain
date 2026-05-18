package keeper_test

import (
	"encoding/json"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"go.uber.org/mock/gomock"

	"github.com/ixofoundation/ixo-blockchain/v6/app/apptesting"
	entitykeeper "github.com/ixofoundation/ixo-blockchain/v6/x/entity/keeper"
	"github.com/ixofoundation/ixo-blockchain/v6/x/entity/testutil"
	"github.com/ixofoundation/ixo-blockchain/v6/x/entity/types"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v6/x/iid/types"
)

// withMockedWasm builds an entity Keeper with mocked WasmKeeper +
// WasmViewKeeper, leaving every other dependency wired to the live IxoApp.
// Returns (mockWasm, mockView, msgServer, *Keeper).
func (s *KeeperTestSuite) withMockedEntityWasm() (
	*testutil.MockWasmKeeper,
	*testutil.MockWasmViewKeeper,
	types.MsgServer,
	*entitykeeper.Keeper,
) {
	ctrl := gomock.NewController(s.T())
	s.T().Cleanup(ctrl.Finish)

	mockWasm := testutil.NewMockWasmKeeper(ctrl)
	mockView := testutil.NewMockWasmViewKeeper(ctrl)

	k := entitykeeper.NewKeeper(
		s.App.AppCodec(),
		s.App.GetKey(types.StoreKey),
		s.App.IidKeeper,
		mockWasm,
		mockView,
		s.App.GetSubspace(types.ModuleName),
		s.App.AccountKeeper,
		s.App.AuthzKeeper,
	)
	return mockWasm, mockView, entitykeeper.NewMsgServerImpl(&k), &k
}

// configureNftParams sets the params required for entity msg flows that
// touch wasm: NftContractAddress + NftContractMinter. CreateSequence is
// left at its default 0 so the generated entity id is deterministic per
// (contract, sequence) pair.
func (s *KeeperTestSuite) configureNftParams(k *entitykeeper.Keeper) (nftAddr, minterAddr sdk.AccAddress) {
	nftAddr = apptesting.RandomAccountAddress()
	minterAddr = apptesting.RandomAccountAddress()
	p := types.Params{
		NftContractAddress: nftAddr.String(),
		NftContractMinter:  minterAddr.String(),
		CreateSequence:     0,
	}
	k.SetParams(s.Ctx, &p)
	return
}

// seedDID stores a DID document with the given controller bech32 address as
// an authentication-relationship verification method. Returns the full
// DIDFragment (e.g. "did:ixo:owner#key-1") for use as msg.OwnerDid /
// msg.ControllerDid.
func (s *KeeperTestSuite) seedDID(label string, addr sdk.AccAddress) iidtypes.DIDFragment {
	did := "did:ixo:" + label
	methodID := did + "#key-1"
	vm := iidtypes.NewVerificationMethod(methodID, iidtypes.DID(did),
		iidtypes.NewBlockchainAccountID(addr.String()))
	meta := iidtypes.NewDidMetadata(s.Ctx.TxBytes(), s.Ctx.BlockTime())
	doc := iidtypes.IidDocument{
		Id:                 did,
		VerificationMethod: []*iidtypes.VerificationMethod{&vm},
		Authentication:     []string{methodID},
		Metadata:           &meta,
	}
	s.App.IidKeeper.SetDidDocument(s.Ctx, []byte(did), doc)
	return iidtypes.DIDFragment(methodID)
}

func (s *KeeperTestSuite) TestMsgCreateEntity_HappyPath() {
	s.SetupTest()
	mockWasm, _, ms, k := s.withMockedEntityWasm()
	nftAddr, minterAddr := s.configureNftParams(k)

	// Seed the relayer node DID (the only iid keeper read CreateEntity does
	// before the wasm call is the relayer-node existence check).
	relayerAddr := apptesting.RandomAccountAddress()
	relayerDID := s.seedDID("relayer", relayerAddr).Did()

	owner := apptesting.RandomAccountAddress()
	ownerDID := s.seedDID("owner", owner)

	// The cw721 mint goes through WasmKeeper.Execute. The minter (msg signer)
	// is params.NftContractMinter; the contract address is params.NftContractAddress.
	mockWasm.EXPECT().
		Execute(gomock.Any(), nftAddr, minterAddr, gomock.Any(), gomock.Any()).
		Return([]byte(nil), nil).
		Times(1)

	resp, err := ms.CreateEntity(s.Ctx, &types.MsgCreateEntity{
		EntityType:   "asset",
		EntityStatus: 1,
		RelayerNode:  relayerDID,
		OwnerDid:     ownerDID,
		OwnerAddress: owner.String(),
		Data:         json.RawMessage(`{}`),
	})
	s.Require().NoError(err)
	s.Require().NotEmpty(resp.EntityId)

	// Persisted entity matches the response id.
	got, found := s.App.EntityKeeper.GetEntity(s.Ctx, []byte(resp.EntityId))
	s.Require().True(found)
	s.Require().Equal("asset", got.Type)
	s.Require().Len(got.Accounts, 1, "admin module account must be in the Accounts list")
	s.Require().Equal(types.EntityAdminAccountName, got.Accounts[0].Name)

	// CreateSequence advanced.
	p := s.App.EntityKeeper.GetParams(s.Ctx)
	s.Require().Equal(uint64(1), p.CreateSequence)
}

// Note: the "empty NftContractAddress" branch in CreateEntity is in
// practice unreachable — the params validator rejects an empty value at
// SetParams time, so this test is intentionally omitted.

func (s *KeeperTestSuite) TestMsgCreateEntity_RelayerDIDMissing() {
	s.SetupTest()
	_, _, ms, k := s.withMockedEntityWasm()
	s.configureNftParams(k)
	owner := apptesting.RandomAccountAddress()
	ownerDID := s.seedDID("owner-no-relay", owner)

	_, err := ms.CreateEntity(s.Ctx, &types.MsgCreateEntity{
		EntityType:   "asset",
		EntityStatus: 1,
		RelayerNode:  "did:ixo:never-existed",
		OwnerDid:     ownerDID,
		OwnerAddress: owner.String(),
		Data:         json.RawMessage(`{}`),
	})
	s.Require().ErrorContains(err, "relayer node did document not found")
}

func (s *KeeperTestSuite) TestMsgCreateEntity_WasmExecuteError() {
	s.SetupTest()
	mockWasm, _, ms, k := s.withMockedEntityWasm()
	nftAddr, minterAddr := s.configureNftParams(k)

	relayerAddr := apptesting.RandomAccountAddress()
	relayerDID := s.seedDID("relayer-wasmerr", relayerAddr).Did()
	owner := apptesting.RandomAccountAddress()
	ownerDID := s.seedDID("owner-wasmerr", owner)

	mockWasm.EXPECT().
		Execute(gomock.Any(), nftAddr, minterAddr, gomock.Any(), gomock.Any()).
		Return([]byte(nil), fmt.Errorf("nft contract rejected")).
		Times(1)

	_, err := ms.CreateEntity(s.Ctx, &types.MsgCreateEntity{
		EntityType:   "asset",
		EntityStatus: 1,
		RelayerNode:  relayerDID,
		OwnerDid:     ownerDID,
		OwnerAddress: owner.String(),
		Data:         json.RawMessage(`{}`),
	})
	s.Require().ErrorContains(err, "nft contract rejected")
}

// TestMsgUpdateEntity exercises the no-wasm UpdateEntity path. It still
// needs a real DID document to exist (ExecuteOnDidWithRelationships looks
// it up) and the controller signer to authenticate against it.
func (s *KeeperTestSuite) TestMsgUpdateEntity() {
	s.SetupTest()
	_, _, ms, _ := s.withMockedEntityWasm()

	// Seed a DID for the entity, listing itself as a controller so
	// ExecuteOnDidWithRelationships's HasController fallback authorises a
	// signer of msg.ControllerDid.Did() == entityDID. (HasRelationship
	// won't match because that compares the BlockchainAccountID — a bech32
	// — against the DID string passed as signer; here the call site uses
	// msg.ControllerDid.Did(), which is a DID, so we route through
	// HasController instead.)
	entityDID := "did:ixo:entity:upd-target"
	meta := iidtypes.NewDidMetadata(s.Ctx.TxBytes(), s.Ctx.BlockTime())
	doc := iidtypes.IidDocument{
		Id:         entityDID,
		Controller: []string{entityDID},
		Metadata:   &meta,
	}
	s.App.IidKeeper.SetDidDocument(s.Ctx, []byte(entityDID), doc)

	// Pre-seed an Entity record.
	now := s.Ctx.BlockTime()
	entityMeta := types.NewEntityMetadata(s.Ctx.TxBytes(), now)
	s.App.EntityKeeper.SetEntity(s.Ctx, []byte(entityDID), types.Entity{
		Id:        entityDID,
		Type:      "asset",
		Status:    1,
		Metadata:  &entityMeta,
		StartDate: &now,
	})

	end := now.Add(time.Hour)
	_, err := ms.UpdateEntity(s.Ctx, &types.MsgUpdateEntity{
		Id:            entityDID,
		EntityStatus:  2,
		StartDate:     &now,
		EndDate:       &end,
		ControllerDid: iidtypes.DIDFragment(entityDID),
	})
	s.Require().NoError(err)

	got, _ := s.App.EntityKeeper.GetEntity(s.Ctx, []byte(entityDID))
	s.Require().Equal(int32(2), got.Status)
	s.Require().NotNil(got.EndDate)
}

// TestMsgUpdateEntityVerified exercises the wasm-free verified flag path.
// RelayerNodeDid.Did() must equal Entity.RelayerNode.
func (s *KeeperTestSuite) TestMsgUpdateEntityVerified() {
	s.SetupTest()
	_, _, ms, _ := s.withMockedEntityWasm()

	relayerAddr := apptesting.RandomAccountAddress()
	relayerDIDFragment := s.seedDID("verifier", relayerAddr)
	relayerDID := relayerDIDFragment.Did()

	entityDID := "did:ixo:entity:verify-target"
	now := s.Ctx.BlockTime()
	entityMeta := types.NewEntityMetadata(s.Ctx.TxBytes(), now)
	// We also need an iid doc for ResolveEntity → IidKeeper.ResolveDid.
	idoc := iidtypes.IidDocument{Id: entityDID, Metadata: &iidtypes.IidMetadata{Created: &now, VersionId: "v1"}}
	s.App.IidKeeper.SetDidDocument(s.Ctx, []byte(entityDID), idoc)
	s.App.EntityKeeper.SetEntity(s.Ctx, []byte(entityDID), types.Entity{
		Id:          entityDID,
		Type:        "asset",
		Status:      1,
		Metadata:    &entityMeta,
		RelayerNode: relayerDID,
	})

	_, err := ms.UpdateEntityVerified(s.Ctx, &types.MsgUpdateEntityVerified{
		Id:                 entityDID,
		EntityVerified:     true,
		RelayerNodeDid:     relayerDIDFragment,
		RelayerNodeAddress: relayerAddr.String(),
	})
	s.Require().NoError(err)
	got, _ := s.App.EntityKeeper.GetEntity(s.Ctx, []byte(entityDID))
	s.Require().True(got.EntityVerified)
}

func (s *KeeperTestSuite) TestMsgUpdateEntityVerified_WrongRelayer() {
	s.SetupTest()
	_, _, ms, _ := s.withMockedEntityWasm()

	relayerAddr := apptesting.RandomAccountAddress()
	relayerDIDFragment := s.seedDID("auth-relayer", relayerAddr)

	entityDID := "did:ixo:entity:wrong-relayer"
	now := s.Ctx.BlockTime()
	entityMeta := types.NewEntityMetadata(s.Ctx.TxBytes(), now)
	idoc := iidtypes.IidDocument{Id: entityDID, Metadata: &iidtypes.IidMetadata{Created: &now, VersionId: "v1"}}
	s.App.IidKeeper.SetDidDocument(s.Ctx, []byte(entityDID), idoc)
	s.App.EntityKeeper.SetEntity(s.Ctx, []byte(entityDID), types.Entity{
		Id:          entityDID,
		Status:      1,
		Metadata:    &entityMeta,
		RelayerNode: "did:ixo:other-relayer",
	})

	_, err := ms.UpdateEntityVerified(s.Ctx, &types.MsgUpdateEntityVerified{
		Id:                 entityDID,
		EntityVerified:     true,
		RelayerNodeDid:     relayerDIDFragment,
		RelayerNodeAddress: relayerAddr.String(),
	})
	s.Require().ErrorContains(err, "invalid relayer node did")
}

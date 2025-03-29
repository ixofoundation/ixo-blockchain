package keeper

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/v4/x/entity/types"
	nft "github.com/ixofoundation/ixo-blockchain/v4/x/entity/types/contracts"
	iidkeeper "github.com/ixofoundation/ixo-blockchain/v4/x/iid/keeper"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v4/x/iid/types"
)

type msgServer struct {
	Keeper *Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the module MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(k *Keeper) types.MsgServer {
	return &msgServer{
		Keeper: k,
	}
}

// --------------------------
// CREATE ENTITY
// --------------------------
func (s msgServer) CreateEntity(goCtx context.Context, msg *types.MsgCreateEntity) (*types.MsgCreateEntityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := s.Keeper.GetParams(ctx)
	nftContractAddressParam := params.NftContractAddress

	if len(nftContractAddressParam) == 0 {
		return nil, errors.New("nftContractAddress not set")
	}

	// check that relayerNode did exists
	_, found := s.Keeper.IidKeeper.GetDidDocument(ctx, []byte(msg.RelayerNode))
	if !found {
		return nil, errorsmod.Wrapf(iidtypes.ErrDidDocumentNotFound, "relayer node did document not found for %s", msg.RelayerNode)
	}

	nftContractAddress, err := sdk.AccAddressFromBech32(nftContractAddressParam)
	if err != nil {
		return nil, err
	}

	minterAddress, err := sdk.AccAddressFromBech32(params.NftContractMinter)
	if err != nil {
		return nil, err
	}

	generatedId := md5.Sum([]byte(fmt.Sprintf("%s/%d", nftContractAddressParam, params.CreateSequence)))
	entityId := fmt.Sprintf("did:ixo:entity:%x", generatedId)

	// check that the did is not already taken
	_, found = s.Keeper.IidKeeper.GetDidDocument(ctx, []byte(entityId))
	if found {
		err := errorsmod.Wrapf(iidtypes.ErrDidDocumentFound, "a document with did %s already exists", entityId)
		return nil, err
	}

	// create and persist the iid document
	did, err := iidtypes.NewDidDocument(ctx,
		entityId,
		iidtypes.WithServices(msg.Service...),
		iidtypes.WithRights(msg.AccordedRight...),
		iidtypes.WithResources(msg.LinkedResource...),
		iidtypes.WithClaims(msg.LinkedClaim...),
		iidtypes.WithEntities(msg.LinkedEntity...),
		iidtypes.WithVerifications(msg.Verification...),
		iidtypes.WithContexts(msg.Context...),
		iidtypes.WithControllers(append(msg.Controller, entityId, msg.OwnerDid.Did())...),
		iidtypes.WithAlsoKnownAs(msg.AlsoKnownAs),
	)
	if err != nil {
		return nil, err
	}
	s.Keeper.IidKeeper.SetDidDocument(ctx, []byte(entityId), did)

	// create admin module account
	adminAddress, err := s.Keeper.CreateNewAccount(ctx, entityId, types.EntityAdminAccountName)
	if err != nil {
		return nil, err
	}

	// add admin module account to entities accounts
	var enityAccounts []*types.EntityAccount
	enityAccounts = append(enityAccounts, &types.EntityAccount{Name: types.EntityAdminAccountName, Address: adminAddress.String()})

	// create and persist the entity
	entityMeta := types.NewEntityMetadata(ctx.TxBytes(), ctx.BlockTime())
	entity := types.Entity{
		Id:          entityId,
		Type:        msg.EntityType,
		StartDate:   msg.StartDate,
		EndDate:     msg.EndDate,
		Status:      msg.EntityStatus,
		RelayerNode: msg.RelayerNode,
		Credentials: msg.Credentials,
		Metadata:    &entityMeta,
		Accounts:    enityAccounts,
	}
	s.Keeper.SetEntity(ctx, []byte(entityId), entity)

	// update and persist createSequence
	params.CreateSequence++
	s.Keeper.SetParams(ctx, &params)

	// create the nft cw721
	finalMessage, err := nft.Marshal(nft.WasmMsgMint{
		Mint: nft.Mint{
			TokenId:   did.Id,
			Owner:     msg.OwnerAddress,
			TokenUri:  did.Id,
			Extension: msg.Data,
		},
	})
	if err != nil {
		return nil, err
	}

	_, err = s.Keeper.WasmKeeper.Execute(ctx, nftContractAddress, minterAddress, finalMessage, sdk.NewCoins(sdk.NewCoin("uixo", math.ZeroInt())))
	if err != nil {
		return nil, err
	}

	// emit the events
	if err := ctx.EventManager().EmitTypedEvents(
		iidtypes.NewIidDocumentCreatedEvent(&did),
		&types.EntityCreatedEvent{
			Entity: &entity,
			Signer: msg.OwnerDid.Did(),
		},
	); err != nil {
		return nil, err
	}

	return &types.MsgCreateEntityResponse{
		EntityId:     entityId,
		EntityType:   entity.Type,
		EntityStatus: entity.Status,
	}, nil
}

// --------------------------
// UPDATE ENTITY
// --------------------------
func (s msgServer) UpdateEntity(goCtx context.Context, msg *types.MsgUpdateEntity) (*types.MsgUpdateEntityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, entity, err := s.Keeper.ResolveEntity(ctx, msg.Id)
	if err != nil {
		return nil, err
	}

	if err := iidkeeper.ExecuteOnDidWithRelationships(
		ctx,
		s.Keeper.IidKeeper,
		[]string{iidtypes.Authentication},
		msg.Id,
		msg.ControllerDid.Did(),
		func(didDoc *iidtypes.IidDocument) error {
			entity.Status = msg.EntityStatus
			entity.StartDate = msg.StartDate
			entity.EndDate = msg.EndDate
			entity.Credentials = msg.Credentials
			return nil
		}); err != nil {
		return nil, err
	}

	types.UpdateEntityMetadata(entity.Metadata, ctx.TxBytes(), ctx.BlockTime())
	s.Keeper.SetEntity(ctx, []byte(entity.Id), entity)

	// emit the events
	if err := ctx.EventManager().EmitTypedEvents(
		&types.EntityUpdatedEvent{
			Entity: &entity,
			Signer: msg.ControllerDid.Did(),
		},
	); err != nil {
		return nil, err
	}

	return &types.MsgUpdateEntityResponse{}, nil
}

// --------------------------
// TRANSFER ENTITY
// --------------------------
func (s msgServer) TransferEntity(goCtx context.Context, msg *types.MsgTransferEntity) (*types.MsgTransferEntityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := s.Keeper.GetParams(ctx)
	nftContractAddressParam := params.NftContractAddress

	if len(nftContractAddressParam) == 0 {
		return nil, errors.New("nftContractAddress not set")
	}

	controllerAddress, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		return nil, err
	}

	recipientDidDoc, found := s.Keeper.IidKeeper.GetDidDocument(ctx, []byte(msg.RecipientDid))
	if !found {
		return nil, errors.New("recipient did not found")
	}

	recipientAddress, err := recipientDidDoc.GetVerificationMethodBlockchainAddress(recipientDidDoc.Id)
	if err != nil {
		return nil, err
	}

	nftContractAddress, err := sdk.AccAddressFromBech32(nftContractAddressParam)
	if err != nil {
		return nil, err
	}

	err = iidkeeper.ExecuteOnDidWithRelationships(
		ctx,
		s.Keeper.IidKeeper,
		[]string{iidtypes.Authentication},
		msg.Id,
		msg.OwnerDid.Did(),
		func(document *iidtypes.IidDocument) error {
			// clear controller of iid Doc and set to new owner
			document.Controller = []string{
				document.Id,
				msg.RecipientDid.Did(),
			}

			// Only remove verification method with recipient did as Id if it exists
			for _, vm := range document.VerificationMethod {
				if vm.Id == msg.RecipientDid.Did() {
					err := document.RevokeVerification(vm.Id)
					if err != nil {
						return err
					}
				}
			}

			// Add recipient did as verification method, with address as verification material
			vm := iidtypes.NewBlockchainAccountID(recipientAddress.String())
			err := document.AddVerifications(iidtypes.NewVerification(
				iidtypes.NewVerificationMethod(msg.RecipientDid.Did(), iidtypes.DID(msg.RecipientDid), vm),
				[]string{iidtypes.Authentication},
				nil,
			))
			if err != nil {
				return err
			}

			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	finalMessage, err := nft.Marshal(nft.WasmMsgTransferNft{
		TransferNft: nft.TransferNft{
			TokenId:   msg.Id,
			Recipient: recipientAddress.String(),
		},
	})
	if err != nil {
		return nil, err
	}

	_, err = s.Keeper.WasmKeeper.Execute(ctx, nftContractAddress, controllerAddress, finalMessage, sdk.NewCoins(sdk.NewCoin("uixo", math.ZeroInt())))
	if err != nil {
		return nil, err
	}

	// emit the events
	if err := ctx.EventManager().EmitTypedEvents(
		&types.EntityTransferredEvent{
			Id:   msg.Id,
			To:   msg.OwnerDid.Did(),
			From: msg.RecipientDid.Did(),
		},
	); err != nil {
		return nil, err
	}

	return &types.MsgTransferEntityResponse{}, nil
}

// --------------------------
// UPDATE ENTITY VERIFIED
// --------------------------
func (s msgServer) UpdateEntityVerified(goCtx context.Context, msg *types.MsgUpdateEntityVerified) (*types.MsgUpdateEntityVerifiedResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, entity, err := s.Keeper.ResolveEntity(ctx, msg.Id)
	if err != nil {
		return nil, err
	}

	if msg.RelayerNodeDid.Did() != entity.RelayerNode {
		return nil, errorsmod.Wrapf(types.ErrUpdateVerifiedFailed, "invalid relayer node did (%s)", msg.RelayerNodeDid.String())
	}

	entity.EntityVerified = msg.EntityVerified

	types.UpdateEntityMetadata(entity.Metadata, ctx.TxBytes(), ctx.BlockTime())
	s.Keeper.SetEntity(ctx, []byte(entity.Id), entity)

	// emit the events
	if err := ctx.EventManager().EmitTypedEvents(
		&types.EntityUpdatedEvent{
			Entity: &entity,
			Signer: msg.RelayerNodeDid.Did(),
		},
		&types.EntityVerifiedUpdatedEvent{
			Id:             msg.Id,
			Signer:         msg.RelayerNodeDid.Did(),
			EntityVerified: msg.EntityVerified,
		},
	); err != nil {
		return nil, err
	}

	return &types.MsgUpdateEntityVerifiedResponse{}, nil
}

// --------------------------
// CREATE ENTITY ACCOUNT
// --------------------------
func (s msgServer) CreateEntityAccount(goCtx context.Context, msg *types.MsgCreateEntityAccount) (*types.MsgCreateEntityAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get entity doc
	_, entity, err := s.Keeper.ResolveEntity(ctx, msg.Id)
	if err != nil {
		return nil, err
	}

	// check if entity account with name already exists
	nameExists := entity.ContainsAccountName(msg.Name)
	if nameExists {
		return nil, errorsmod.Wrapf(types.ErrAccountDuplicate, "name %s", msg.Name)
	}

	// check that owner is nft owner
	err = s.Keeper.CheckIfOwner(ctx, msg.Id, msg.OwnerAddress)
	if err != nil {
		return nil, errorsmod.Wrapf(err, "unauthorized")
	}

	// create module account
	address, err := s.Keeper.CreateNewAccount(ctx, msg.Id, msg.Name)
	if err != nil {
		return nil, err
	}

	// update entity and persist
	entity.Accounts = append(entity.Accounts, &types.EntityAccount{Name: msg.Name, Address: address.String()})
	types.UpdateEntityMetadata(entity.Metadata, ctx.TxBytes(), ctx.BlockTime())
	s.Keeper.SetEntity(ctx, []byte(entity.Id), entity)

	// emit the events
	if err := ctx.EventManager().EmitTypedEvents(
		&types.EntityUpdatedEvent{
			Entity: &entity,
			Signer: msg.OwnerAddress,
		},
		&types.EntityAccountCreatedEvent{
			Id:             entity.Id,
			Signer:         msg.OwnerAddress,
			AccountName:    msg.Name,
			AccountAddress: address.String(),
		},
	); err != nil {
		return nil, err
	}

	return &types.MsgCreateEntityAccountResponse{Account: address.String()}, nil
}

// --------------------------
// GRANT ENTITY ACCOUNT AUTHZ
// --------------------------
func (s msgServer) GrantEntityAccountAuthz(goCtx context.Context, msg *types.MsgGrantEntityAccountAuthz) (*types.MsgGrantEntityAccountAuthzResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get entity doc
	_, entity, err := s.Keeper.ResolveEntity(ctx, msg.Id)
	if err != nil {
		return nil, err
	}

	// check that owner is entity owner
	err = s.Keeper.CheckIfOwner(ctx, msg.Id, msg.OwnerAddress)
	if err != nil {
		return nil, errorsmod.Wrapf(err, "unauthorized")
	}

	// get entity account with same name
	var account *types.EntityAccount
	for _, acc := range entity.Accounts {
		if account != nil {
			break
		}
		if acc.Name == msg.Name {
			account = acc
		}
	}

	// if no account with name throw
	if account == nil {
		return nil, errorsmod.Wrapf(types.ErrAccountNotFound, "name %s", msg.Name)
	}

	// get addresses
	grantee, err := sdk.AccAddressFromBech32(msg.GranteeAddress)
	if err != nil {
		return nil, err
	}
	granter, err := sdk.AccAddressFromBech32(account.Address)
	if err != nil {
		return nil, err
	}

	// Unpack interface to both have concrete type and add to cache for validation method
	err = msg.Grant.UnpackInterfaces(s.Keeper.cdc)
	if err != nil {
		return nil, err
	}

	err = msg.Grant.ValidateBasic()
	if err != nil {
		return nil, err
	}

	// get authorization
	authorization, err := msg.Grant.GetAuthorization()
	if err != nil {
		return nil, err
	}

	// persist new grant
	if err := s.Keeper.AuthzKeeper.SaveGrant(ctx, grantee, granter, authorization, msg.Grant.Expiration); err != nil {
		return nil, err
	}

	// emit the events
	if err := ctx.EventManager().EmitTypedEvents(
		&types.EntityAccountAuthzCreatedEvent{
			Id:          entity.Id,
			Signer:      msg.OwnerAddress,
			AccountName: msg.Name,
			Granter:     granter.String(),
			Grantee:     grantee.String(),
			Grant:       &msg.Grant,
		},
	); err != nil {
		return nil, err
	}

	return &types.MsgGrantEntityAccountAuthzResponse{}, nil
}

// --------------------------
// REVOKE ENTITY ACCOUNT AUTHZ
// --------------------------
func (s msgServer) RevokeEntityAccountAuthz(goCtx context.Context, msg *types.MsgRevokeEntityAccountAuthz) (*types.MsgRevokeEntityAccountAuthzResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get entity doc
	_, entity, err := s.Keeper.ResolveEntity(ctx, msg.Id)
	if err != nil {
		return nil, err
	}

	// check that owner is entity owner
	err = s.Keeper.CheckIfOwner(ctx, msg.Id, msg.OwnerAddress)
	if err != nil {
		return nil, errorsmod.Wrapf(err, "unauthorized")
	}

	// get entity account with same name
	var account *types.EntityAccount
	for _, acc := range entity.Accounts {
		if account != nil {
			break
		}
		if acc.Name == msg.Name {
			account = acc
		}
	}

	// if no account with name throw
	if account == nil {
		return nil, errorsmod.Wrapf(types.ErrAccountNotFound, "name %s", msg.Name)
	}

	// get addresses
	grantee, err := sdk.AccAddressFromBech32(msg.GranteeAddress)
	if err != nil {
		return nil, err
	}
	granter, err := sdk.AccAddressFromBech32(account.Address)
	if err != nil {
		return nil, err
	}

	// remove grant
	if err := s.Keeper.AuthzKeeper.DeleteGrant(ctx, grantee, granter, msg.MsgTypeUrl); err != nil {
		return nil, err
	}

	// emit the events
	if err := ctx.EventManager().EmitTypedEvents(
		&types.EntityAccountAuthzRevokedEvent{
			Id:          entity.Id,
			Signer:      msg.OwnerAddress,
			AccountName: msg.Name,
			Granter:     granter.String(),
			Grantee:     grantee.String(),
			MsgTypeUrl:  msg.MsgTypeUrl,
		},
	); err != nil {
		return nil, err
	}

	return &types.MsgRevokeEntityAccountAuthzResponse{}, nil
}

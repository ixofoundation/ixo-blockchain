package keeper

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ixofoundation/ixo-blockchain/x/entity/types"
	entitycontracts "github.com/ixofoundation/ixo-blockchain/x/entity/types/contracts"
	iidkeeper "github.com/ixofoundation/ixo-blockchain/x/iid/keeper"
	iidtypes "github.com/ixofoundation/ixo-blockchain/x/iid/types"
)

type msgServer struct {
	Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the module MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(k Keeper) types.MsgServer {
	return &msgServer{
		Keeper: k,
	}
}

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
		return nil, sdkerrors.Wrapf(iidtypes.ErrDidDocumentNotFound, "relayer node did document not found for %s", msg.RelayerNode)
	}

	nftContractAddress, err := sdk.AccAddressFromBech32(nftContractAddressParam)
	if err != nil {
		return nil, err
	}

	address, err := sdk.AccAddressFromBech32(params.NftContractMinter)
	if err != nil {
		return nil, err
	}

	generatedId := md5.Sum([]byte(fmt.Sprintf("%s/%d", nftContractAddressParam, params.CreateSequence)))
	entityId := fmt.Sprintf("did:ixo:entity:%x", generatedId)

	// check that the did is not already taken
	_, found = s.Keeper.IidKeeper.GetDidDocument(ctx, []byte(entityId))
	if found {
		err := sdkerrors.Wrapf(iidtypes.ErrDidDocumentFound, "a document with did %s already exists", entityId)
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
	}
	s.Keeper.SetEntity(ctx, []byte(entityId), entity)

	// update and persist createSequence
	params.CreateSequence++
	s.Keeper.SetParams(ctx, &params)

	// create the nft cw721
	nftMint := entitycontracts.WasmMsgMint{
		Mint: entitycontracts.Mint{
			TokenId:   did.Id,
			Owner:     msg.OwnerAddress,
			TokenUri:  did.Id,
			Extension: msg.Data,
		},
	}

	finalMessage, err := nftMint.Marshal()
	if err != nil {
		return nil, err
	}

	_, err = s.Keeper.WasmKeeper.Execute(ctx, nftContractAddress, address, finalMessage, sdk.NewCoins(sdk.NewCoin("uixo", sdk.ZeroInt())))
	if err != nil {
		return nil, err
	}

	// emit the events
	if err := ctx.EventManager().EmitTypedEvents(
		iidtypes.NewIidDocumentCreatedEvent(&did),
		types.NewEntityCreatedEvent(&entity, msg.OwnerDid.Did()),
	); err != nil {
		return nil, err
	}

	return &types.MsgCreateEntityResponse{
		EntityId:     entityId,
		EntityType:   entity.Type,
		EntityStatus: entity.Status,
	}, nil
}

func (s msgServer) UpdateEntity(goCtx context.Context, msg *types.MsgUpdateEntity) (*types.MsgUpdateEntityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, entity, err := s.ResolveEntity(ctx, msg.Id)
	if err != nil {
		return nil, err
	}

	if err := iidkeeper.ExecuteOnDidWithRelationships(
		ctx,
		&s.Keeper.IidKeeper,
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
		types.NewEntityUpdatedEvent(&entity, msg.ControllerDid.String()),
	); err != nil {
		return nil, err
	}

	return &types.MsgUpdateEntityResponse{}, nil
}

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
		&s.Keeper.IidKeeper,
		[]string{iidtypes.Authentication},
		msg.Id,
		msg.OwnerDid.Did(),
		func(document *iidtypes.IidDocument) error {
			// clear controller of iid Doc and set to new owner
			document.Controller = []string{
				document.Id,
				msg.RecipientDid.Did(),
			}

			// remove old verification methods
			for _, vm := range document.VerificationMethod {
				document.RevokeVerification(vm.Id)
			}

			// Add recipient did as verification method
			vm := iidtypes.NewBlockchainAccountID(recipientAddress.String())
			document.AddVerifications(iidtypes.NewVerification(
				iidtypes.NewVerificationMethod(msg.RecipientDid.Did(), iidtypes.DID(msg.RecipientDid), vm),
				[]string{iidtypes.Authentication},
				nil,
			))

			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	nftTranferMsg := entitycontracts.WasmMsgTransferNft{
		TransferNft: entitycontracts.TransferNft{
			TokenId:   msg.Id,
			Recipient: recipientAddress.String(),
		},
	}

	finalMessage, err := nftTranferMsg.Marshal()
	if err != nil {
		return nil, err
	}

	_, err = s.Keeper.WasmKeeper.Execute(ctx, nftContractAddress, controllerAddress, finalMessage, sdk.NewCoins(sdk.NewCoin("uixo", sdk.ZeroInt())))
	if err != nil {
		return nil, err
	}

	// emit the events
	if err := ctx.EventManager().EmitTypedEvents(
		types.NewEntityTransferredEvent(msg.Id, msg.OwnerDid.Did(), msg.RecipientDid.Did()),
	); err != nil {
		return nil, err
	}

	return &types.MsgTransferEntityResponse{}, nil
}

func (s msgServer) UpdateEntityVerified(goCtx context.Context, msg *types.MsgUpdateEntityVerified) (*types.MsgUpdateEntityVerifiedResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, entity, err := s.ResolveEntity(ctx, msg.Id)
	if err != nil {
		return nil, err
	}

	if msg.RelayerNodeDid.String() != entity.RelayerNode {
		return nil, sdkerrors.Wrapf(types.ErrUpdateVerifiedFailed, "invalid relayer node did (%s)", msg.RelayerNodeDid.String())
	}

	entity.EntityVerified = msg.EntityVerified

	// if err := iidkeeper.ExecuteOnDidWithRelationships(
	// 	ctx,
	// 	&s.Keeper.IidKeeper,
	// 	[]string{iidtypes.Authentication},
	// 	msg.Id,
	// 	msg.RelayerNodeDid.Did(),
	// 	func(didDoc *iidtypes.IidDocument) error {
	// 		entity.EntityVerified = msg.EntityVerified
	// 		return nil
	// 	}); err != nil {
	// 	return nil, err
	// }

	types.UpdateEntityMetadata(entity.Metadata, ctx.TxBytes(), ctx.BlockTime())
	s.Keeper.SetEntity(ctx, []byte(entity.Id), entity)

	// emit the events
	if err := ctx.EventManager().EmitTypedEvents(
		types.NewEntityUpdatedEvent(&entity, msg.RelayerNodeDid.String()),
		types.NewEntityVerifiedUpdatedEvent(msg.Id, msg.RelayerNodeDid.String(), msg.EntityVerified),
	); err != nil {
		return nil, err
	}

	return &types.MsgUpdateEntityVerifiedResponse{}, nil
}

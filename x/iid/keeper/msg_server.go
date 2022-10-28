package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ixofoundation/ixo-blockchain/x/iid/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the identity MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// CreateDidDocument creates a new DID document
func (k msgServer) CreateIidDocument(
	goCtx context.Context,
	msg *types.MsgCreateIidDocument,
) (*types.MsgCreateIidDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.Logger(ctx).Info("request to create a did document", "target did", msg.Id)
	// setup a new did document (performs input validation)
	did, err := types.NewDidDocument(msg.Id,
		types.WithServices(msg.Services...),
		types.WithRights(msg.AccordedRight...),
		types.WithResources(msg.LinkedResource...),
		types.WithEntities(msg.LinkedEntity...),
		types.WithVerifications(msg.Verifications...),
		types.WithControllers(msg.Controllers...),
	)
	if err != nil {
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}

	// check that the did is not already taken
	_, found := k.Keeper.GetDidDocument(ctx, []byte(msg.Id))
	if found {
		err := sdkerrors.Wrapf(types.ErrDidDocumentFound, "a document with did %s already exists", msg.Id)
		k.Logger(ctx).Error(err.Error())
		return nil, err
	}

	// persist the did document
	k.Keeper.SetDidDocument(ctx, []byte(msg.Id), did)

	// now create and persist the metadata
	didM := types.NewDidMetadata(ctx.TxBytes(), ctx.BlockTime())
	k.Keeper.SetDidMetadata(ctx, []byte(msg.Id), didM)

	k.Logger(ctx).Info("created did document", "did", msg.Id, "controller", msg.Signer)

	// emit the event
	if err := ctx.EventManager().EmitTypedEvents(types.NewIidDocumentCreatedEvent(msg.Id, msg.Signer)); err != nil {
		k.Logger(ctx).Error("failed to emit DidDocumentCreatedEvent", "did", msg.Id, "signer", msg.Signer, "err", err)
	}

	return &types.MsgCreateIidDocumentResponse{}, nil
}

// UpdateDidDocument update an existing DID document
func (k msgServer) UpdateIidDocument(
	goCtx context.Context,
	msg *types.MsgUpdateIidDocument,
) (*types.MsgUpdateIidDocumentResponse, error) {

	if err := ExecuteOnDidWithRelationships(
		goCtx, &k.Keeper,
		newConstraints(types.Authentication),
		msg.Doc.Id, msg.Signer,
		//XXX: check this assignment during audit
		//nolint
		func(didDoc *types.IidDocument) error {
			if !types.IsValidDIDDocument(msg.Doc) {
				return sdkerrors.Wrapf(types.ErrInvalidDIDFormat, "invalid did document")
			}
			didDoc = msg.Doc
			return nil
		}); err != nil {
		return nil, err
	}
	return &types.MsgUpdateIidDocumentResponse{}, nil
}

// AddVerification adds a verification method and it's relationships to a DID Document
func (k msgServer) AddVerification(
	goCtx context.Context,
	msg *types.MsgAddVerification,
) (*types.MsgAddVerificationResponse, error) {

	if err := ExecuteOnDidWithRelationships(
		goCtx, &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *types.IidDocument) error {
			return didDoc.AddVerifications(msg.Verification)
		}); err != nil {
		return nil, err
	}
	return &types.MsgAddVerificationResponse{}, nil
}

// AddService adds a service to an existing DID document
func (k msgServer) AddService(
	goCtx context.Context,
	msg *types.MsgAddService,
) (*types.MsgAddServiceResponse, error) {

	if err := ExecuteOnDidWithRelationships(
		goCtx, &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *types.IidDocument) error {
			return didDoc.AddServices(msg.ServiceData)
		}); err != nil {
		return nil, err
	}

	return &types.MsgAddServiceResponse{}, nil
}

func (k msgServer) AddLinkedResource(
	goCtx context.Context,
	msg *types.MsgAddLinkedResource,
) (*types.MsgAddLinkedResourceResponse, error) {

	if err := ExecuteOnDidWithRelationships(
		goCtx, &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *types.IidDocument) error {
			return didDoc.AddLinkedResource(msg.LinkedResource)
		}); err != nil {
		return nil, err
	}

	return &types.MsgAddLinkedResourceResponse{}, nil
}

func (k msgServer) DeleteLinkedResource(
	goCtx context.Context,
	msg *types.MsgDeleteLinkedResource,
) (*types.MsgDeleteLinkedResourceResponse, error) {

	if err := ExecuteOnDidWithRelationships(
		goCtx, &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *types.IidDocument) error {
			// Only try to remove service if there are services
			if len(didDoc.LinkedResource) == 0 {
				return sdkerrors.Wrapf(types.ErrInvalidState, "the did document doesn't have resources associated")
			}
			// delete service
			didDoc.DeleteLinkedResource(msg.ResourceId)
			return nil
		}); err != nil {
		return nil, err
	}

	return &types.MsgDeleteLinkedResourceResponse{}, nil
}

func (k msgServer) AddLinkedEntity(
	goCtx context.Context,
	msg *types.MsgAddLinkedEntity,
) (*types.MsgAddLinkedEntityResponse, error) {

	if err := ExecuteOnDidWithRelationships(
		goCtx, &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *types.IidDocument) error {
			return didDoc.AddLinkedEntity(msg.LinkedEntity)
		}); err != nil {
		return nil, err
	}

	return &types.MsgAddLinkedEntityResponse{}, nil
}

func (k msgServer) DeleteLinkedEntity(
	goCtx context.Context,
	msg *types.MsgDeleteLinkedEntity,
) (*types.MsgDeleteLinkedEntityResponse, error) {

	if err := ExecuteOnDidWithRelationships(
		goCtx, &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *types.IidDocument) error {
			// Only try to remove service if there are services
			if len(didDoc.LinkedEntity) == 0 {
				return sdkerrors.Wrapf(types.ErrInvalidState, "the did document doesn't have entities associated")
			}
			// delete service
			didDoc.DeleteLinkedResource(msg.EntityId)
			return nil
		}); err != nil {
		return nil, err
	}

	return &types.MsgDeleteLinkedEntityResponse{}, nil
}

func (k msgServer) AddAccordedRight(
	goCtx context.Context,
	msg *types.MsgAddAccordedRight,
) (*types.MsgAddAccordedRightResponse, error) {

	if err := ExecuteOnDidWithRelationships(
		goCtx, &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *types.IidDocument) error {
			return didDoc.AddAccordedRight(msg.AccordedRight)
		}); err != nil {
		return nil, err
	}

	return &types.MsgAddAccordedRightResponse{}, nil
}

func (k msgServer) DeleteAccordedRight(
	goCtx context.Context,
	msg *types.MsgDeleteAccordedRight,
) (*types.MsgDeleteAccordedRightResponse, error) {

	if err := ExecuteOnDidWithRelationships(
		goCtx, &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *types.IidDocument) error {
			// Only try to remove right if there are rights
			if len(didDoc.AccordedRight) == 0 {
				return sdkerrors.Wrapf(types.ErrInvalidState, "the did document doesn't have rights associated")
			}
			// delete service
			didDoc.DeleteAccordedRight(msg.RightId)
			return nil
		}); err != nil {
		return nil, err
	}

	return &types.MsgDeleteAccordedRightResponse{}, nil
}

// Contexts
func (k msgServer) AddIidContext(
	goCtx context.Context,
	msg *types.MsgAddIidContext,
) (*types.MsgAddIidContextResponse, error) {

	if err := ExecuteOnDidWithRelationships(
		goCtx, &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *types.IidDocument) error {
			return didDoc.AddDidContext(msg.Context)
		}); err != nil {
		return nil, err
	}

	return &types.MsgAddIidContextResponse{}, nil
}

func (k msgServer) DeleteIidContext(
	goCtx context.Context,
	msg *types.MsgDeleteIidContext,
) (*types.MsgDeleteIidContextResponse, error) {

	if err := ExecuteOnDidWithRelationships(
		goCtx, &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *types.IidDocument) error {
			// Only try to remove right if there are rights
			if len(didDoc.Context) == 0 {
				return sdkerrors.Wrapf(types.ErrInvalidState, "the did document doesn't have contexts associated")
			}
			// delete service
			didDoc.DeleteDidContext(msg.ContextKey)
			return nil
		}); err != nil {
		return nil, err
	}

	return &types.MsgDeleteIidContextResponse{}, nil
}

// RevokeVerification removes a public key and controller from an existing DID document
func (k msgServer) RevokeVerification(
	goCtx context.Context,
	msg *types.MsgRevokeVerification,
) (*types.MsgRevokeVerificationResponse, error) {

	if err := ExecuteOnDidWithRelationships(
		goCtx, &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *types.IidDocument) error {
			return didDoc.RevokeVerification(msg.MethodId)
		}); err != nil {
		return nil, err
	}

	return &types.MsgRevokeVerificationResponse{}, nil
}

// DeleteService removes a service from an existing DID document
func (k msgServer) DeleteService(
	goCtx context.Context,
	msg *types.MsgDeleteService,
) (*types.MsgDeleteServiceResponse, error) {

	if err := ExecuteOnDidWithRelationships(
		goCtx, &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *types.IidDocument) error {
			// Only try to remove service if there are services
			if len(didDoc.Service) == 0 {
				return sdkerrors.Wrapf(types.ErrInvalidState, "the did document doesn't have services associated")
			}
			// delete service
			didDoc.DeleteService(msg.ServiceId)
			return nil
		}); err != nil {
		return nil, err
	}

	return &types.MsgDeleteServiceResponse{}, nil
}

// SetVerificationRelationships set the verification relationships for an existing DID document
func (k msgServer) SetVerificationRelationships(
	goCtx context.Context,
	msg *types.MsgSetVerificationRelationships,
) (*types.MsgSetVerificationRelationshipsResponse, error) {

	if err := ExecuteOnDidWithRelationships(
		goCtx, &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *types.IidDocument) error {
			return didDoc.SetVerificationRelationships(msg.MethodId, msg.Relationships...)
		}); err != nil {
		return nil, err
	}

	return &types.MsgSetVerificationRelationshipsResponse{}, nil
}

// AddController add a new controller to a DID
func (k msgServer) AddController(
	goCtx context.Context,
	msg *types.MsgAddController,
) (*types.MsgAddControllerResponse, error) {
	if err := ExecuteOnDidWithRelationships(
		goCtx, &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *types.IidDocument) error {
			return didDoc.AddControllers(msg.ControllerDid)
		}); err != nil {
		return nil, err
	}

	return &types.MsgAddControllerResponse{}, nil
}

// DeleteController remove an existing controller from a DID document
func (k msgServer) DeleteController(
	goCtx context.Context,
	msg *types.MsgDeleteController,
) (*types.MsgDeleteControllerResponse, error) {

	if err := ExecuteOnDidWithRelationships(
		goCtx, &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer, func(didDoc *types.IidDocument) error {
			return didDoc.DeleteControllers(msg.ControllerDid)
		}); err != nil {
		return nil, err
	}
	return &types.MsgDeleteControllerResponse{}, nil
}

// helper function to update the did metadata
func updateDidMetadata(keeper *Keeper, ctx sdk.Context, did string) (err error) {
	didMeta, found := keeper.GetDidMetadata(ctx, []byte(did))
	if found {
		types.UpdateDidMetadata(&didMeta, ctx.TxBytes(), ctx.BlockTime())
		keeper.SetDidMetadata(ctx, []byte(did), didMeta)
	} else {
		err = fmt.Errorf("(warning) did metadata not found")
	}
	return
}

func (k msgServer) UpdateMetaData(
	goCtx context.Context,
	msg *types.MsgUpdateIidMeta,
) (*types.MsgUpdateIidMetaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	didMeta, found := k.GetDidMetadata(ctx, []byte(msg.Id))
	nm := msg.Meta
	if found {
		newMeta := types.IidMetadata{
			VersionId:            "",
			Created:              didMeta.Created,
			Updated:              nil,
			Deactivated:          nm.Deactivated,
			EntityType:           nm.EntityType,
			StartDate:            nm.StartDate,
			EndDate:              nm.EndDate,
			Status:               nm.Status,
			Stage:                nm.Stage,
			RelayerNode:          nm.RelayerNode,
			VerifiableCredential: nm.VerifiableCredential,
			Credentials:          nm.Credentials,
		}
		txH := sha256.Sum256(ctx.TxBytes())
		newMeta.VersionId = hex.EncodeToString(txH[:])
		var upd time.Time
		upd = ctx.BlockTime()
		newMeta.Updated = &upd
		k.SetDidMetadata(ctx, []byte(msg.Id), newMeta)
	} else {
		//proper error response
	}

	return &types.MsgUpdateIidMetaResponse{}, nil
}

func (k msgServer) DeactivateIID(
	goCtx context.Context,
	msg *types.MsgDeactivateIID,
) (*types.MsgDeactivateIIDResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	didMeta, found := k.GetDidMetadata(ctx, []byte(msg.Id))
	if found {
		didMeta.Deactivated = msg.State
		txH := sha256.Sum256(ctx.TxBytes())
		didMeta.VersionId = hex.EncodeToString(txH[:])
		var upd time.Time
		upd = ctx.BlockTime()
		didMeta.Updated = &upd
		k.SetDidMetadata(ctx, []byte(msg.Id), didMeta)
	} else {
		//proper error response
	}

	return &types.MsgDeactivateIIDResponse{}, nil
}

// VerificationRelationships for did document manipulation
type VerificationRelationships []string

func newConstraints(relationships ...string) VerificationRelationships {
	return relationships
}

func ExecuteOnDidWithRelationships(goCtx context.Context, k *Keeper, constraints VerificationRelationships, did, signer string, update func(document *types.IidDocument) error) (err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.Logger(ctx).Info("request to update a did document", "target did", did)
	// TODO: fail if the input did is of type KEY (immutable)
	// eg: ErrInvalidState, "did document key is immutable"

	// get the did document
	didDoc, found := k.GetDidDocument(ctx, []byte(did))
	if !found {
		err = sdkerrors.Wrapf(types.ErrDidDocumentNotFound, "did document at %s not found", did)
		k.Logger(ctx).Error(err.Error())
		return
	}

	// Any verification method in the authentication relationship can update the DID document
	if !didDoc.HasRelationship(types.NewBlockchainAccountID(ctx.ChainID(), signer), constraints...) {
		// check also the controllers
		signerDID := types.NewKeyDID(signer)
		if !didDoc.HasController(signerDID) {
			// if also the controller was not set the error
			err = sdkerrors.Wrapf(
				types.ErrUnauthorized,
				"signer account %s not authorized to update the target did document at %s",
				signer, did,
			)
			k.Logger(ctx).Error(err.Error())
			return
		}
	}

	// apply the update
	err = update(&didDoc)
	if err != nil {
		k.Logger(ctx).Error(err.Error())
		return
	}

	// persist the did document
	k.SetDidDocument(ctx, []byte(did), didDoc)
	k.Logger(ctx).Info("Set verification relationship from did document for", "did", did, "controller", signer)

	// update the Metadata
	if err = updateDidMetadata(k, ctx, didDoc.Id); err != nil {
		k.Logger(ctx).Error(err.Error(), "did", didDoc.Id)
		return
	}
	// fire the event
	if err := ctx.EventManager().EmitTypedEvent(types.NewIidDocumentUpdatedEvent(did, signer)); err != nil {
		k.Logger(ctx).Error("failed to emit DidDocumentUpdatedEvent", "did", did, "signer", signer, "err", err)
	}
	k.Logger(ctx).Info("request to update did document success", "did", didDoc.Id)
	return
}

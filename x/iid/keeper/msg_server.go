package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/v6/x/iid/types"
)

type msgServer struct {
	Keeper
}

// TODO: add ControllerDid check for iid update types

// NewMsgServerImpl returns an implementation of the identity MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// CreateDidDocument creates a new DID document
func (k msgServer) CreateIidDocument(goCtx context.Context, msg *types.MsgCreateIidDocument,
) (*types.MsgCreateIidDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check that the did is not already taken
	_, found := k.Keeper.GetDidDocument(ctx, []byte(msg.Id))
	if found {
		err := errorsmod.Wrapf(types.ErrDidDocumentFound, "a document with did %s already exists", msg.Id)
		return nil, err
	}

	// setup a new did document (performs input validation)
	did, err := types.NewDidDocument(ctx, msg.Id,
		types.WithServices(msg.Services...),
		types.WithRights(msg.AccordedRight...),
		types.WithResources(msg.LinkedResource...),
		types.WithClaims(msg.LinkedClaim...),
		types.WithEntities(msg.LinkedEntity...),
		types.WithVerifications(msg.Verifications...),
		types.WithControllers(msg.Controllers...),
		types.WithContexts(msg.Context...),
		types.WithAlsoKnownAs(msg.AlsoKnownAs),
	)
	if err != nil {
		return nil, err
	}

	k.Keeper.SetDidDocument(ctx, []byte(msg.Id), did)

	// emit the event
	if err := ctx.EventManager().EmitTypedEvents(types.NewIidDocumentCreatedEvent(&did)); err != nil {
		return nil, err
	}

	return &types.MsgCreateIidDocumentResponse{}, nil
}

// UpdateDidDocument update an existing DID document
func (k msgServer) UpdateIidDocument(goCtx context.Context, msg *types.MsgUpdateIidDocument,
) (*types.MsgUpdateIidDocumentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := ExecuteOnDidWithRelationships(
		ctx, &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *types.IidDocument) error {
			did, err := types.NewDidDocument(ctx, msg.Id,
				types.WithServices(msg.Services...),
				types.WithRights(msg.AccordedRight...),
				types.WithResources(msg.LinkedResource...),
				types.WithClaims(msg.LinkedClaim...),
				types.WithEntities(msg.LinkedEntity...),
				types.WithVerifications(msg.Verifications...),
				types.WithControllers(msg.Controllers...),
				types.WithContexts(msg.Context...),
				types.WithAlsoKnownAs(msg.AlsoKnownAs),
			)
			if err != nil {
				return err
			}
			// Keep old iid doc metadata
			did.Metadata = didDoc.Metadata
			return nil
		}); err != nil {
		return nil, err
	}
	return &types.MsgUpdateIidDocumentResponse{}, nil
}

func (k msgServer) AddVerification(goCtx context.Context, msg *types.MsgAddVerification,
) (*types.MsgAddVerificationResponse, error) {
	if err := ExecuteOnDidWithRelationships(
		sdk.UnwrapSDKContext(goCtx), &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *types.IidDocument) error {
			return didDoc.AddVerifications(msg.Verification)
		}); err != nil {
		return nil, err
	}
	return &types.MsgAddVerificationResponse{}, nil
}

func (k msgServer) AddService(goCtx context.Context, msg *types.MsgAddService,
) (*types.MsgAddServiceResponse, error) {
	if err := ExecuteOnDidWithRelationships(
		sdk.UnwrapSDKContext(goCtx), &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *types.IidDocument) error {
			return didDoc.AddServices(msg.ServiceData)
		}); err != nil {
		return nil, err
	}

	return &types.MsgAddServiceResponse{}, nil
}

func (k msgServer) AddLinkedResource(goCtx context.Context, msg *types.MsgAddLinkedResource,
) (*types.MsgAddLinkedResourceResponse, error) {
	if err := ExecuteOnDidWithRelationships(
		sdk.UnwrapSDKContext(goCtx), &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *types.IidDocument) error {
			return didDoc.AddLinkedResource(msg.LinkedResource)
		}); err != nil {
		return nil, err
	}

	return &types.MsgAddLinkedResourceResponse{}, nil
}

func (k msgServer) DeleteLinkedResource(goCtx context.Context, msg *types.MsgDeleteLinkedResource,
) (*types.MsgDeleteLinkedResourceResponse, error) {
	if err := ExecuteOnDidWithRelationships(
		sdk.UnwrapSDKContext(goCtx), &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *types.IidDocument) error {
			if len(didDoc.LinkedResource) == 0 {
				return errorsmod.Wrapf(types.ErrInvalidState, "the did document doesn't have resources associated")
			}
			didDoc.DeleteLinkedResource(msg.ResourceId)
			return nil
		}); err != nil {
		return nil, err
	}

	return &types.MsgDeleteLinkedResourceResponse{}, nil
}

func (k msgServer) AddLinkedClaim(goCtx context.Context, msg *types.MsgAddLinkedClaim,
) (*types.MsgAddLinkedClaimResponse, error) {
	if err := ExecuteOnDidWithRelationships(
		sdk.UnwrapSDKContext(goCtx), &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *types.IidDocument) error {
			return didDoc.AddLinkedClaim(msg.LinkedClaim)
		}); err != nil {
		return nil, err
	}

	return &types.MsgAddLinkedClaimResponse{}, nil
}

func (k msgServer) DeleteLinkedClaim(goCtx context.Context, msg *types.MsgDeleteLinkedClaim,
) (*types.MsgDeleteLinkedClaimResponse, error) {
	if err := ExecuteOnDidWithRelationships(
		sdk.UnwrapSDKContext(goCtx), &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *types.IidDocument) error {
			if len(didDoc.LinkedClaim) == 0 {
				return errorsmod.Wrapf(types.ErrInvalidState, "the did document doesn't have Claims associated")
			}
			didDoc.DeleteLinkedClaim(msg.ClaimId)
			return nil
		}); err != nil {
		return nil, err
	}

	return &types.MsgDeleteLinkedClaimResponse{}, nil
}

func (k msgServer) AddLinkedEntity(goCtx context.Context, msg *types.MsgAddLinkedEntity,
) (*types.MsgAddLinkedEntityResponse, error) {
	if err := ExecuteOnDidWithRelationships(
		sdk.UnwrapSDKContext(goCtx), &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *types.IidDocument) error {
			return didDoc.AddLinkedEntity(msg.LinkedEntity)
		}); err != nil {
		return nil, err
	}

	return &types.MsgAddLinkedEntityResponse{}, nil
}

func (k msgServer) DeleteLinkedEntity(goCtx context.Context, msg *types.MsgDeleteLinkedEntity,
) (*types.MsgDeleteLinkedEntityResponse, error) {
	if err := ExecuteOnDidWithRelationships(
		sdk.UnwrapSDKContext(goCtx), &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *types.IidDocument) error {
			if len(didDoc.LinkedEntity) == 0 {
				return errorsmod.Wrapf(types.ErrInvalidState, "the did document doesn't have entities associated")
			}
			didDoc.DeleteLinkedEntity(msg.EntityId)
			return nil
		}); err != nil {
		return nil, err
	}

	return &types.MsgDeleteLinkedEntityResponse{}, nil
}

func (k msgServer) AddAccordedRight(goCtx context.Context, msg *types.MsgAddAccordedRight,
) (*types.MsgAddAccordedRightResponse, error) {
	if err := ExecuteOnDidWithRelationships(
		sdk.UnwrapSDKContext(goCtx), &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *types.IidDocument) error {
			return didDoc.AddAccordedRight(msg.AccordedRight)
		}); err != nil {
		return nil, err
	}

	return &types.MsgAddAccordedRightResponse{}, nil
}

func (k msgServer) DeleteAccordedRight(goCtx context.Context, msg *types.MsgDeleteAccordedRight,
) (*types.MsgDeleteAccordedRightResponse, error) {
	if err := ExecuteOnDidWithRelationships(
		sdk.UnwrapSDKContext(goCtx), &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *types.IidDocument) error {
			if len(didDoc.AccordedRight) == 0 {
				return errorsmod.Wrapf(types.ErrInvalidState, "the did document doesn't have rights associated")
			}
			didDoc.DeleteAccordedRight(msg.RightId)
			return nil
		}); err != nil {
		return nil, err
	}

	return &types.MsgDeleteAccordedRightResponse{}, nil
}

func (k msgServer) AddIidContext(goCtx context.Context, msg *types.MsgAddIidContext,
) (*types.MsgAddIidContextResponse, error) {
	if err := ExecuteOnDidWithRelationships(
		sdk.UnwrapSDKContext(goCtx), &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *types.IidDocument) error {
			return didDoc.AddDidContext(msg.Context)
		}); err != nil {
		return nil, err
	}

	return &types.MsgAddIidContextResponse{}, nil
}

func (k msgServer) DeleteIidContext(goCtx context.Context, msg *types.MsgDeleteIidContext,
) (*types.MsgDeleteIidContextResponse, error) {
	if err := ExecuteOnDidWithRelationships(
		sdk.UnwrapSDKContext(goCtx), &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *types.IidDocument) error {
			if len(didDoc.Context) == 0 {
				return errorsmod.Wrapf(types.ErrInvalidState, "the did document doesn't have contexts associated")
			}
			didDoc.DeleteDidContext(msg.ContextKey)
			return nil
		}); err != nil {
		return nil, err
	}

	return &types.MsgDeleteIidContextResponse{}, nil
}

func (k msgServer) RevokeVerification(goCtx context.Context, msg *types.MsgRevokeVerification,
) (*types.MsgRevokeVerificationResponse, error) {
	if err := ExecuteOnDidWithRelationships(
		sdk.UnwrapSDKContext(goCtx), &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *types.IidDocument) error {
			return didDoc.RevokeVerification(msg.MethodId)
		}); err != nil {
		return nil, err
	}

	return &types.MsgRevokeVerificationResponse{}, nil
}

func (k msgServer) DeleteService(goCtx context.Context, msg *types.MsgDeleteService,
) (*types.MsgDeleteServiceResponse, error) {
	if err := ExecuteOnDidWithRelationships(
		sdk.UnwrapSDKContext(goCtx), &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *types.IidDocument) error {
			if len(didDoc.Service) == 0 {
				return errorsmod.Wrapf(types.ErrInvalidState, "the did document doesn't have services associated")
			}
			didDoc.DeleteService(msg.ServiceId)
			return nil
		}); err != nil {
		return nil, err
	}

	return &types.MsgDeleteServiceResponse{}, nil
}

func (k msgServer) SetVerificationRelationships(goCtx context.Context, msg *types.MsgSetVerificationRelationships,
) (*types.MsgSetVerificationRelationshipsResponse, error) {
	if err := ExecuteOnDidWithRelationships(
		sdk.UnwrapSDKContext(goCtx), &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *types.IidDocument) error {
			return didDoc.SetVerificationRelationships(msg.MethodId, msg.Relationships...)
		}); err != nil {
		return nil, err
	}

	return &types.MsgSetVerificationRelationshipsResponse{}, nil
}

func (k msgServer) AddController(goCtx context.Context, msg *types.MsgAddController,
) (*types.MsgAddControllerResponse, error) {
	if err := ExecuteOnDidWithRelationships(
		sdk.UnwrapSDKContext(goCtx), &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *types.IidDocument) error {
			return didDoc.AddControllers(msg.ControllerDid)
		}); err != nil {
		return nil, err
	}

	return &types.MsgAddControllerResponse{}, nil
}

func (k msgServer) DeleteController(goCtx context.Context, msg *types.MsgDeleteController,
) (*types.MsgDeleteControllerResponse, error) {
	if err := ExecuteOnDidWithRelationships(
		sdk.UnwrapSDKContext(goCtx), &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer,
		func(didDoc *types.IidDocument) error {
			return didDoc.DeleteControllers(msg.ControllerDid)
		}); err != nil {
		return nil, err
	}
	return &types.MsgDeleteControllerResponse{}, nil
}

func (k msgServer) DeactivateIID(goCtx context.Context, msg *types.MsgDeactivateIID,
) (*types.MsgDeactivateIIDResponse, error) {
	if err := ExecuteOnDidWithRelationships(
		sdk.UnwrapSDKContext(goCtx), &k.Keeper,
		newConstraints(types.Authentication),
		msg.Id, msg.Signer, func(didDoc *types.IidDocument) error {
			return didDoc.Deactivate()
		}); err != nil {
		return nil, err
	}
	return &types.MsgDeactivateIIDResponse{}, nil
}

// VerificationRelationships for did document manipulation
type VerificationRelationships []string

func newConstraints(relationships ...string) VerificationRelationships {
	return relationships
}

// Check the relations/controllers for did if have capabilities to modify did doc, do modifications and emit doc_updated event
func ExecuteOnDidWithRelationships(ctx sdk.Context, k types.IidKeeper, constraints VerificationRelationships, did, signer string, update func(document *types.IidDocument) error) (err error) {
	// get the did document
	didDoc, found := k.GetDidDocument(ctx, []byte(did))
	if !found {
		err = errorsmod.Wrapf(types.ErrDidDocumentNotFound, "did document at %s not found", did)
		return err
	}

	// Any verification method in the authentication relationship can update the DID document
	if !didDoc.HasRelationship(types.NewBlockchainAccountID(signer), constraints...) {
		// check also the controllers
		// TODO: for usage of this in iid msg_server.go, the signer is always address, so this will never pass?
		if !didDoc.HasController(types.DID(signer)) {
			// if also the controller was not set the error
			err = errorsmod.Wrapf(
				types.ErrUnauthorized,
				"signer account %s not authorized to update the target did document at %s",
				signer, did,
			)
			return err
		}
	}

	// apply the update
	err = update(&didDoc)
	if err != nil {
		return err
	}

	// update the Metadata
	types.UpdateDidMetadata(didDoc.Metadata, ctx.TxBytes(), ctx.BlockTime())

	// persist the did document
	k.SetDidDocument(ctx, []byte(did), didDoc)

	// fire the event
	err = ctx.EventManager().EmitTypedEvent(types.NewIidDocumentUpdatedEvent(&didDoc))
	if err != nil {
		return err
	}
	return nil
}

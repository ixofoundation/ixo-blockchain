package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ixofoundation/ixo-blockchain/x/claims/types"
	iidtypes "github.com/ixofoundation/ixo-blockchain/x/iid/types"
)

type msgServer struct {
	Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

// --------------------------
// CREATE COLLECTION
// --------------------------
func (s msgServer) CreateCollection(goCtx context.Context, msg *types.MsgCreateCollection) (*types.MsgCreateCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := s.Keeper.GetParams(ctx)

	_, found := s.Keeper.IidKeeper.GetDidDocument(ctx, []byte(msg.Entity))
	if !found {
		return nil, sdkerrors.Wrapf(iidtypes.ErrDidDocumentNotFound, "for entity %s", msg.Entity)
	}

	_, found = s.Keeper.IidKeeper.GetDidDocument(ctx, []byte(msg.Protocol))
	if !found {
		return nil, sdkerrors.Wrapf(iidtypes.ErrDidDocumentNotFound, "for protocol %s", msg.Protocol)
	}

	// create and persist the Collection
	collectionId := fmt.Sprint(params.CollectionSequence)
	collection := types.Collection{
		Id:        collectionId,
		Entity:    msg.Entity,
		Admin:     msg.Admin,
		Protocol:  msg.Protocol,
		StartDate: msg.StartDate,
		EndDate:   msg.EndDate,
		Quota:     msg.Quota,
		Count:     0,
		Evaluated: 0,
		Approved:  0,
		Rejected:  0,
		Disputed:  0,
		State:     msg.State,
		Payments:  msg.Payments,
	}
	s.Keeper.SetCollection(ctx, collection)

	// update and persist createSequence
	params.CollectionSequence++
	s.Keeper.SetParams(ctx, &params)

	// emit the events
	if err := ctx.EventManager().EmitTypedEvents(
		&types.CollectionCreatedEvent{
			Collection: &collection,
		},
	); err != nil {
		return nil, err
	}

	return &types.MsgCreateCollectionResponse{}, nil
}

// --------------------------
// SUBMIT CLAIM
// --------------------------
func (s msgServer) SubmitClaim(goCtx context.Context, msg *types.MsgSubmitClaim) (*types.MsgSubmitClaimResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Make sure claim does not exist already
	_, err := s.Keeper.GetClaim(ctx, msg.ClaimId)
	if err == nil {
		return nil, sdkerrors.Wrapf(types.ErrClaimDuplicate, "id %s", msg.ClaimId)
	}

	// Get Collection for claim
	collection, err := s.Keeper.GetCollection(ctx, msg.CollectionId)
	if err != nil {
		return nil, err
	}

	// check that user is authorized, aka signer is admin for Collection
	if collection.Admin != msg.AdminAddress {
		return nil, sdkerrors.Wrapf(types.ErrClaimUnauthorized, "Collection admin %s, msg admin address %s", collection.Admin, msg.AdminAddress)
	}

	// check that collection is in open state
	if collection.State != types.CollectionState_open {
		return nil, sdkerrors.Wrapf(types.ErrCollectionNotOpen, "state %s", collection.State)
	}

	// check that collection has already started and has not ended yet
	now := ctx.BlockTime()
	if now.Before(*collection.StartDate) {
		return nil, types.ErrClaimCollectionNotStarted
	}
	if !collection.EndDate.IsZero() && now.After(*collection.EndDate) {
		return nil, types.ErrClaimCollectionEnded
	}

	// check if collection quota has been reached
	if collection.Quota != 0 && collection.Quota <= collection.Count {
		return nil, types.ErrClaimCollectionQuotaReached
	}

	// create and persist the Claim
	claimSubmissionDate := ctx.BlockTime()
	claim := types.Claim{
		CollectionId:   msg.CollectionId,
		AgentDid:       msg.AgentDid.Did(),
		AgentAddress:   msg.AgentAddress,
		ClaimId:        msg.ClaimId,
		SubmissionDate: &claimSubmissionDate,
	}
	s.Keeper.SetClaim(ctx, claim)

	// update count for collection and persist
	collection.Count++
	s.Keeper.SetCollection(ctx, collection)

	// emit the events
	if err := ctx.EventManager().EmitTypedEvents(
		&types.ClaimSubmittedEvent{
			Claim: &claim,
		},
		&types.CollectionUpdatedEvent{
			Collection: &collection,
		},
	); err != nil {
		return nil, err
	}

	return &types.MsgSubmitClaimResponse{}, nil
}

// --------------------------
// EVALUATE CLAIM
// --------------------------
func (s msgServer) EvaluateClaim(goCtx context.Context, msg *types.MsgEvaluateClaim) (*types.MsgEvaluateClaimResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get Claim for evaluation
	claim, err := s.Keeper.GetClaim(ctx, msg.ClaimId)
	if err != nil {
		return nil, err
	}

	// check that claim was not evaluated already
	if claim.Evaluation != nil {
		return nil, sdkerrors.Wrapf(types.ErrClaimDuplicateEvaluation, "id %s", claim.ClaimId)
	}

	// get Collection for claim
	collection, err := s.Keeper.GetCollection(ctx, claim.CollectionId)
	if err != nil {
		return nil, err
	}

	// check that user is authorized, aka signer is admin for Collection
	if collection.Admin != msg.AdminAddress {
		return nil, sdkerrors.Wrapf(types.ErrClaimUnauthorized, "Collection admin %s, msg admin address %s", collection.Admin, msg.AdminAddress)
	}

	// create and persist the Evaluation
	evaluationDate := ctx.BlockTime()
	evaluation := types.Evaluation{
		Oracle:            msg.Oracle,
		AgentDid:          msg.AgentDid.Did(),
		AgentAddress:      msg.AgentAddress,
		Status:            msg.Status,
		Reason:            msg.Reason,
		VerificationProof: msg.VerificationProof,
		EvaluationDate:    &evaluationDate,
	}

	claim.Evaluation = &evaluation
	s.Keeper.SetClaim(ctx, claim)

	// update amounts for collection and persist
	collection.Evaluated++
	if msg.Status == types.EvaluationStatus_approved {
		collection.Approved++
	} else if msg.Status == types.EvaluationStatus_rejected {
		collection.Rejected++
	} else if msg.Status == types.EvaluationStatus_disputed {
		collection.Disputed++
	}
	s.Keeper.SetCollection(ctx, collection)

	// emit the events
	if err := ctx.EventManager().EmitTypedEvents(
		&types.ClaimEvaluatedEvent{
			Evaluation: &evaluation,
		},
		&types.ClaimUpdatedEvent{
			Claim: &claim,
		},
		&types.CollectionUpdatedEvent{
			Collection: &collection,
		},
	); err != nil {
		return nil, err
	}
	return &types.MsgEvaluateClaimResponse{}, nil
}

// --------------------------
// DISPUTE CLAIM
// --------------------------
func (s msgServer) DisputeClaim(goCtx context.Context, msg *types.MsgDisputeClaim) (*types.MsgDisputeClaimResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check that claim exists
	_, err := s.Keeper.GetClaim(ctx, msg.ClaimId)
	if err != nil {
		return nil, err
	}

	// Make sure dispute with proof does not exist already
	_, err = s.Keeper.GetDispute(ctx, msg.Data.Proof)
	if err == nil {
		return nil, sdkerrors.Wrapf(types.ErrDisputeDuplicate, "proof %s", msg.Data.Proof)
	}

	// create and persist the dispute
	dispute := types.Dispute{
		AgentDid:     msg.AgentDid.Did(),
		AgentAddress: msg.AgentAddress,
		ClaimId:      msg.ClaimId,
		Type:         msg.DisputeType,
		Data:         msg.Data,
	}
	s.Keeper.SetDispute(ctx, dispute)

	// emit the events
	if err := ctx.EventManager().EmitTypedEvents(
		&types.ClaimDisputedEvent{
			Dispute: &dispute,
		},
	); err != nil {
		return nil, err
	}
	return &types.MsgDisputeClaimResponse{}, nil
}

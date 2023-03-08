package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ixofoundation/ixo-blockchain/lib/ixo"
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

	// check that entity exists
	_, entity, err := s.Keeper.EntityKeeper.ResolveEntity(ctx, msg.Entity)
	if err != nil {
		return nil, sdkerrors.Wrapf(iidtypes.ErrDidDocumentNotFound, "for entity %s", msg.Entity)
	}

	// check that protocol exists
	if _, found := s.Keeper.IidKeeper.GetDidDocument(ctx, []byte(msg.Protocol)); !found {
		return nil, sdkerrors.Wrapf(iidtypes.ErrDidDocumentNotFound, "for protocol %s", msg.Protocol)
	}

	// check that signer is nft owner
	if err = s.Keeper.EntityKeeper.CheckIfOwner(ctx, msg.Entity, msg.Signer); err != nil {
		return nil, sdkerrors.Wrapf(err, "unauthorized")
	}

	// check that Evaluation Payment does not have 1155 payment
	if msg.Payments.Evaluation.Contract_1155Payment != nil {
		return nil, types.ErrCollectionEvalError
	}

	// check that all payments accounts is part of entity module accounts
	if !msg.Payments.AccountsIsEntityAccounts(entity) {
		return nil, types.ErrCollNotEntityAcc
	}

	// get entity admin account
	admin, err := entity.GetAdminAccount()
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "for admin")
	}

	// create and persist the Collection
	var collection types.Collection
	collectionId := fmt.Sprint(params.CollectionSequence)
	collection = types.Collection{
		Id:        collectionId,
		Entity:    msg.Entity,
		Admin:     admin.Address,
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

	// Get agent address
	agent, err := sdk.AccAddressFromBech32(msg.AgentAddress)
	if err != nil {
		return nil, err
	}

	// check that user is authorized, aka signer is admin for Collection
	if collection.Admin != msg.AdminAddress {
		return nil, sdkerrors.Wrapf(types.ErrClaimUnauthorized, "collection admin %s, msg admin address %s", collection.Admin, msg.AdminAddress)
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
		PaymentsStatus: &types.ClaimPayments{
			Submission: types.PaymentStatus_no_payment,
			Approval:   types.PaymentStatus_no_payment,
			Evaluation: types.PaymentStatus_no_payment,
			Rejection:  types.PaymentStatus_no_payment,
		},
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

	// start payout process for claim submission
	if err = processPayment(ctx, s.Keeper, agent, collection.Payments.Submission, types.PaymentType_submission, msg.ClaimId); err != nil {
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

	// check that collectionId in message corresponds to claims collection
	if claim.CollectionId != msg.CollectionId {
		return nil, sdkerrors.Wrapf(types.ErrEvaluateWrongCollection, "claim collection %s vs message collection %s", claim.CollectionId, msg.CollectionId)
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
		return nil, sdkerrors.Wrapf(types.ErrClaimUnauthorized, "collection admin %s, msg admin address %s", collection.Admin, msg.AdminAddress)
	}

	// Get evaluation agent address
	evalAgent, err := sdk.AccAddressFromBech32(msg.AgentAddress)
	if err != nil {
		return nil, err
	}

	// Get claim agent address
	claimAgent, err := sdk.AccAddressFromBech32(claim.AgentAddress)
	if err != nil {
		return nil, err
	}

	// create and persist the Evaluation
	evaluationDate := ctx.BlockTime()
	evaluation := types.Evaluation{
		ClaimId:           msg.ClaimId,
		CollectionId:      msg.CollectionId,
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

	// start payout process for evaluation submission
	if err = processPayment(ctx, s.Keeper, evalAgent, collection.Payments.Evaluation, types.PaymentType_evaluation, msg.ClaimId); err != nil {
		return nil, err
	}

	// update amounts for collection, make payouts and persist
	collection.Evaluated++
	if msg.Status == types.EvaluationStatus_approved {
		collection.Approved++
		// payout process for evaluation approval to claim agent
		// if msg amount is not zero, it means agent set custo amount that was authenticated through authZ constraints to be valid.
		approvedPayment := collection.Payments.Approval
		if !msg.Amount.IsZero() {
			approvedPayment.Amount = msg.Amount
		}
		if err = processPayment(ctx, s.Keeper, claimAgent, approvedPayment, types.PaymentType_approval, msg.ClaimId); err != nil {
			return nil, err
		}
	} else if msg.Status == types.EvaluationStatus_rejected {
		// no payment for rejected
		collection.Rejected++
	} else if msg.Status == types.EvaluationStatus_disputed {
		// no payment for disputed
		collection.Disputed++
		// update payment status to disputed
		updatePaymentStatus(ctx, s.Keeper, types.PaymentType_approval, msg.ClaimId, types.PaymentStatus_disputed)
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

	// Make sure dispute with proof does not exist already
	_, err := s.Keeper.GetDispute(ctx, msg.Data.Proof)
	if err == nil {
		return nil, sdkerrors.Wrapf(types.ErrDisputeDuplicate, "proof %s", msg.Data.Proof)
	}

	// get Claim for dispute
	claim, err := s.Keeper.GetClaim(ctx, msg.SubjectId)
	if err != nil {
		return nil, err
	}

	// get Collection for claim
	collection, err := s.Keeper.GetCollection(ctx, claim.CollectionId)
	if err != nil {
		return nil, err
	}

	entity, found := s.Keeper.IidKeeper.GetDidDocument(ctx, []byte(collection.Entity))
	if !found {
		return nil, sdkerrors.Wrapf(iidtypes.ErrDidDocumentNotFound, "for entity %s", collection.Entity)
	}

	// check if user authorized to lay claim,
	// check if user is admin on Collection or if user is a controller on Collections entity
	if msg.AgentAddress != collection.Admin && !entity.HasController(iidtypes.DID(msg.AgentDid.Did())) {
		// check if user has authz cap, aka is agent
		isAuthorized := false
		grantee, err := sdk.AccAddressFromBech32(msg.AgentAddress)
		if err != nil {
			return nil, err
		}
		granter, err := sdk.AccAddressFromBech32(collection.Admin)
		if err != nil {
			return nil, err
		}

		// get users current authorization to see if user is agent for claim/collection
		authorizations := s.Keeper.AuthzKeeper.GetAuthorizations(ctx, grantee, granter)

		for _, auth := range authorizations {
			if isAuthorized {
				break
			}
			switch k := auth.(type) {
			case *types.SubmitClaimAuthorization:
				// check if there a constraint that has collectionId of disputed subjectId(claim)
				for _, con := range k.Constraints {
					if con.CollectionId == collection.Id {
						isAuthorized = true
					}
				}
			case *types.EvaluateClaimAuthorization:
				// check if there a constraint that has collectionId or claimId of disputed subjectId(claim)
				for _, con := range k.Constraints {
					if con.CollectionId == collection.Id || ixo.Contains(con.ClaimIds, claim.ClaimId) {
						isAuthorized = true
					}
				}
			}
		}

		if !isAuthorized {
			return nil, types.ErrDisputeUnauthorized
		}
	}

	// create and persist the dispute
	dispute := types.Dispute{
		SubjectId: msg.SubjectId,
		Type:      msg.DisputeType,
		Data:      msg.Data,
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

// --------------------------
// WITHDRAW PAYMENT
// --------------------------
func (s msgServer) WithdrawPayment(goCtx context.Context, msg *types.MsgWithdrawPayment) (*types.MsgWithdrawPaymentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get Claim for dispute
	claim, err := s.Keeper.GetClaim(ctx, msg.ClaimId)
	if err != nil {
		return nil, err
	}

	// get Collection for claim
	collection, err := s.Keeper.GetCollection(ctx, claim.CollectionId)
	if err != nil {
		return nil, err
	}

	// check that user is authorized, aka signer is admin for Collection
	if collection.Admin != msg.AdminAddress {
		return nil, sdkerrors.Wrapf(types.ErrClaimUnauthorized, "collection admin %s, msg admin address %s", collection.Admin, msg.AdminAddress)
	}

	// get from address
	fromAddress, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return nil, err
	}
	// get to address
	toAddress, err := sdk.AccAddressFromBech32(msg.ToAddress)
	if err != nil {
		return nil, err
	}

	// make payout
	err = payout(ctx, s.Keeper, msg.Inputs, msg.Outputs, msg.PaymentType, msg.ClaimId, msg.ReleaseDate, msg.Contract_1155Payment, fromAddress, toAddress)
	if err != nil {
		return nil, err
	}
	return &types.MsgWithdrawPaymentResponse{}, nil
}

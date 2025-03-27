package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/authz"
	"github.com/ixofoundation/ixo-blockchain/v4/lib/ixo"
	"github.com/ixofoundation/ixo-blockchain/v4/x/claims/types"
	entitytypes "github.com/ixofoundation/ixo-blockchain/v4/x/entity/types"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v4/x/iid/types"
)

type msgServer struct {
	Keeper *Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper *Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

// TODO: ADD 1155 and 721 capabilities to claims and payments also.
// TODO: add possibility to allow multiple intents per agent based of collection flag

// --------------------------
// CREATE COLLECTION
// --------------------------
func (s msgServer) CreateCollection(goCtx context.Context, msg *types.MsgCreateCollection) (*types.MsgCreateCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check that entity exists
	_, entity, err := s.Keeper.EntityKeeper.ResolveEntity(ctx, msg.Entity)
	if err != nil {
		return nil, errorsmod.Wrapf(iidtypes.ErrDidDocumentNotFound, "for entity %s", msg.Entity)
	}

	// check that protocol exists
	if _, found := s.Keeper.IidKeeper.GetDidDocument(ctx, []byte(msg.Protocol)); !found {
		return nil, errorsmod.Wrapf(iidtypes.ErrDidDocumentNotFound, "for protocol %s", msg.Protocol)
	}

	// check that signer is nft owner
	if err = s.Keeper.EntityKeeper.CheckIfOwner(ctx, msg.Entity, msg.Signer); err != nil {
		return nil, errorsmod.Wrapf(err, "unauthorized")
	}

	// check that Evaluation Payment does not have 1155 payment
	if msg.Payments.Evaluation.Contract_1155Payment != nil {
		return nil, types.ErrCollectionEvalError
	}
	// check that Evaluation Payment does not have CW20 payments
	if len(msg.Payments.Evaluation.Cw20Payment) > 1 {
		return nil, types.ErrCollectionEvalCW20Error
	}

	// check that all payments accounts is part of entity module accounts
	if !msg.Payments.AccountsIsEntityAccounts(entity) {
		return nil, types.ErrCollNotEntityAcc
	}

	// get entity admin account
	admin, err := entity.GetEntityAccountByName(entitytypes.EntityAdminAccountName)
	if err != nil {
		return nil, errorsmod.Wrapf(err, "for admin")
	}

	// get collection id from params and update params
	params := s.Keeper.GetParams(ctx)
	collectionSequence := params.CollectionSequence
	params.CollectionSequence++
	s.Keeper.SetParams(ctx, &params)

	// create and persist the Collection
	collection := types.Collection{
		Id:          fmt.Sprint(collectionSequence),
		Entity:      msg.Entity,
		Admin:       admin,
		Protocol:    msg.Protocol,
		StartDate:   msg.StartDate,
		EndDate:     msg.EndDate,
		Quota:       msg.Quota,
		Count:       0,
		Evaluated:   0,
		Approved:    0,
		Rejected:    0,
		Disputed:    0,
		Invalidated: 0,
		State:       msg.State,
		Payments:    msg.Payments,
		Intents:     msg.Intents,
	}

	// create escrow account for the collection
	escrowAccount, err := types.CreateNewCollectionEscrow(ctx, s.Keeper.AccountKeeper, collection.Id)
	if err != nil {
		return nil, errorsmod.Wrapf(err, "failed to create escrow account for collection: %s", collection.Id)
	}
	collection.EscrowAccount = escrowAccount.String()

	// persist and emit the events
	s.Keeper.SetCollection(ctx, collection)
	if err := ctx.EventManager().EmitTypedEvent(
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
		return nil, errorsmod.Wrapf(types.ErrClaimDuplicate, "id %s", msg.ClaimId)
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
		return nil, errorsmod.Wrapf(types.ErrClaimUnauthorized, "collection admin %s, msg admin address %s", collection.Admin, msg.AdminAddress)
	}

	// check that collection is in open state
	if collection.State != types.CollectionState_open {
		return nil, errorsmod.Wrapf(types.ErrCollectionNotOpen, "state %s", collection.State)
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

	// if intents are required then check if intent is used
	if collection.Intents == types.CollectionIntentOptions_required && !msg.UseIntent {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "collection %s requires intent", msg.CollectionId)
	}
	// if intents are not allowed then check if intent is used
	if collection.Intents == types.CollectionIntentOptions_deny && msg.UseIntent {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "collection %s does not allow intent", msg.CollectionId)
	}

	// get intent if used, if used agent must have an active intent for this collection
	var intent types.Intent
	var intentFound bool
	if msg.UseIntent {
		intent, intentFound = s.Keeper.GetActiveIntent(ctx, msg.AgentAddress, msg.CollectionId)
		if !intentFound {
			return nil, errorsmod.Wrapf(types.ErrIntentNotFound, "for agent %s and collection %s", msg.AgentAddress, msg.CollectionId)
		}
		// check if intent is expired
		if !intent.ExpireAt.IsZero() && intent.ExpireAt.Before(ctx.BlockTime()) {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "intent %s is expired", intent.Id)
		}
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
		Amount:      msg.Amount,
		Cw20Payment: msg.Cw20Payment,
		UseIntent:   msg.UseIntent,
	}

	// if intent then override payments, add intent id and update APPROVAL payment to GUARANTEED
	if msg.UseIntent {
		claim.Amount = intent.Amount
		claim.Cw20Payment = intent.Cw20Payment

		// if either payment is not empty then APPROVAL payment become GUARANTEED as funds is in escrow account
		// if both payments is empty or all amounts is 0 then APPROVAL payment stays NO_PAYMENT since no funds are in escrow account
		if !intent.Amount.IsZero() || !types.IsZeroCW20Payments(intent.Cw20Payment) {
			claim.PaymentsStatus.Approval = types.PaymentStatus_guaranteed
		}

		// mark intent as fulfilled
		intent.Status = types.IntentStatus_fulfilled
		intent.ClaimId = claim.ClaimId
		err = s.Keeper.RemoveIntentAndEmitEvents(ctx, intent)
		if err != nil {
			return nil, err
		}
	} else {
		// if no intent used, check if collection approval payment is oracle payment then only native coins allowed
		if collection.Payments.Approval.IsOraclePayment && !types.IsZeroCW20Payments(claim.Cw20Payment) {
			return nil, types.ErrOraclePaymentOnlyNative
		}
	}

	// update count for collection and persist
	collection.Count++
	if err := s.Keeper.CollectionPersistAndEmitEvents(ctx, collection); err != nil {
		return nil, err
	}

	// start payout process for claim submission
	if err = processPayment(ctx, *s.Keeper, agent, collection.Payments.Submission, types.PaymentType_submission, &claim, collection, false); err != nil {
		return nil, err
	}

	// persist claim and emit the events
	s.Keeper.SetClaim(ctx, claim)
	if err := ctx.EventManager().EmitTypedEvents(
		&types.ClaimSubmittedEvent{
			Claim: &claim,
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

	// check that collectionId in message corresponds to claims collection
	if claim.CollectionId != msg.CollectionId {
		return nil, errorsmod.Wrapf(types.ErrEvaluateWrongCollection, "claim collection %s vs message collection %s", claim.CollectionId, msg.CollectionId)
	}

	// check that claim was not evaluated already
	if claim.Evaluation != nil {
		return nil, errorsmod.Wrapf(types.ErrClaimDuplicateEvaluation, "id %s", claim.ClaimId)
	}

	// get Collection for claim
	collection, err := s.Keeper.GetCollection(ctx, claim.CollectionId)
	if err != nil {
		return nil, err
	}

	// check that user is authorized, aka signer is admin for Collection
	if collection.Admin != msg.AdminAddress {
		return nil, errorsmod.Wrapf(types.ErrClaimUnauthorized, "collection admin %s, msg admin address %s", collection.Admin, msg.AdminAddress)
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
		Amount:            msg.Amount,
		Cw20Payment:       msg.Cw20Payment,
	}

	// if intent on claim then override payments with claim payments as it used the intent
	if claim.UseIntent {
		evaluation.Amount = claim.Amount
		evaluation.Cw20Payment = claim.Cw20Payment
	}

	// start payout process for evaluation submission, if evaluation has status invalidated, don't run evaluation payout process
	if msg.Status != types.EvaluationStatus_invalidated {
		if err = processPayment(ctx, *s.Keeper, evalAgent, collection.Payments.Evaluation, types.PaymentType_evaluation, &claim, collection, false); err != nil {
			return nil, err
		}

		// update evaluated count for collection
		collection.Evaluated++
	}

	// if status is not approved and intent was used then transfer funds back out of escrow account to current APPROVAL payments account
	if msg.Status != types.EvaluationStatus_approved && claim.UseIntent {
		// Get account used for APPROVAL payments on collection
		approvalAddress, err := sdk.AccAddressFromBech32(collection.Payments.Approval.Account)
		if err != nil {
			return nil, err
		}
		// Get escrow address
		escrow, err := sdk.AccAddressFromBech32(collection.EscrowAccount)
		if err != nil {
			return nil, err
		}
		if err := s.Keeper.TransferIntentPayments(ctx, escrow, approvalAddress, claim.Amount, claim.Cw20Payment); err != nil {
			return nil, err
		}
		// Update payment status to no payment again as was guaranteed with intent
		err = updatePaymentStatus(types.PaymentType_approval, &claim, types.PaymentStatus_no_payment)
		if err != nil {
			return nil, err
		}
	}

	// update amounts for collection, make payouts and persist
	if msg.Status == types.EvaluationStatus_approved {
		// payout process for evaluation approval to claim agent
		collection.Approved++
		// Dereference the pointer to avoid changing collection payments as collection is saved in keeper below
		approvedPayment := collection.Payments.Approval.Clone()
		// if intent on claim then override payments with claim payments as it used the intent and payment must be intent amount
		// also override payment account to escrow account, so funds can be transferred from escrow to agent
		if claim.UseIntent {
			approvedPayment.Amount = claim.Amount
			approvedPayment.Cw20Payment = claim.Cw20Payment
			approvedPayment.Account = collection.EscrowAccount
		} else {
			// if either msg amount or cw20Payment length is not zero, it means agent set custom amount/cw20Payment that was authenticated
			// through authZ constraints to be valid since all evaluations must be done by collections module account through authz
			if len(msg.Amount) > 0 || len(msg.Cw20Payment) > 0 {
				approvedPayment.Amount = msg.Amount
				approvedPayment.Cw20Payment = msg.Cw20Payment
			}
			// if no intent used, check if collection approval payment is oracle payment then only native coins allowed
			if collection.Payments.Approval.IsOraclePayment && !types.IsZeroCW20Payments(approvedPayment.Cw20Payment) {
				return nil, types.ErrOraclePaymentOnlyNative
			}
		}
		if err = processPayment(ctx, *s.Keeper, claimAgent, approvedPayment, types.PaymentType_approval, &claim, collection, claim.UseIntent); err != nil {
			return nil, err
		}
	} else if msg.Status == types.EvaluationStatus_rejected {
		// payout process for evaluation rejected to claim agent
		collection.Rejected++
		if err = processPayment(ctx, *s.Keeper, claimAgent, collection.Payments.Rejection, types.PaymentType_rejection, &claim, collection, false); err != nil {
			return nil, err
		}
	} else if msg.Status == types.EvaluationStatus_disputed {
		// no payment for disputed
		collection.Disputed++
		// update payment status to disputed
		err := updatePaymentStatus(types.PaymentType_approval, &claim, types.PaymentStatus_disputed)
		if err != nil {
			return nil, err
		}
	} else if msg.Status == types.EvaluationStatus_invalidated {
		// no payment for invalidated
		collection.Invalidated++
	}

	if err := s.Keeper.CollectionPersistAndEmitEvents(ctx, collection); err != nil {
		return nil, err
	}

	// persist and emit the events
	claim.Evaluation = &evaluation
	s.Keeper.SetClaim(ctx, claim)
	if err := ctx.EventManager().EmitTypedEvents(
		&types.ClaimEvaluatedEvent{
			Evaluation: &evaluation,
		},
		&types.ClaimUpdatedEvent{
			Claim: &claim,
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
		return nil, errorsmod.Wrapf(types.ErrDisputeDuplicate, "proof %s", msg.Data.Proof)
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
		return nil, errorsmod.Wrapf(iidtypes.ErrDidDocumentNotFound, "for entity %s", collection.Entity)
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
		authorizations, err := s.Keeper.AuthzKeeper.GetAuthorizations(ctx, grantee, granter)
		if err != nil {
			return nil, types.ErrDisputeUnauthorized
		}

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

	claim, err := s.Keeper.GetClaim(ctx, msg.ClaimId)
	if err != nil {
		return nil, err
	}

	// get Collection for claim
	collection, err := s.Keeper.GetCollection(ctx, claim.CollectionId)
	if err != nil {
		return nil, err
	}

	// if any input address or the fromAddress is the escrow account then return error, as any escrow funds will be
	// paid immediately through intents and never through this function, this also prevents collection owners from taking
	// out funds from escrow account through this function
	if msg.FromAddress == collection.EscrowAccount {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "from address cannot be collection's escrow account")
	}
	if len(msg.Inputs) > 0 {
		for _, i := range msg.Inputs {
			if i.Address == collection.EscrowAccount {
				return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "input address cannot be collection's escrow account")
			}
		}
	}

	// check that user is authorized, aka signer is admin for Collection
	if collection.Admin != msg.AdminAddress {
		return nil, errorsmod.Wrapf(types.ErrClaimUnauthorized, "collection admin %s, msg admin address %s", collection.Admin, msg.AdminAddress)
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
	err = payout(ctx, *s.Keeper, msg.Inputs, msg.Outputs, msg.PaymentType, &claim, collection, msg.ReleaseDate, msg.Contract_1155Payment, msg.Cw20Payment, fromAddress, toAddress, []*types.CW20Output{})
	if err != nil {
		return nil, err
	}

	// persist and emit the events
	s.Keeper.SetClaim(ctx, claim)
	if err := ctx.EventManager().EmitTypedEvent(
		&types.ClaimUpdatedEvent{
			Claim: &claim,
		},
	); err != nil {
		return nil, err
	}

	return &types.MsgWithdrawPaymentResponse{}, nil
}

// --------------------------
// UPDATE COLLECTION STATE
// --------------------------
func (s msgServer) UpdateCollectionState(goCtx context.Context, msg *types.MsgUpdateCollectionState) (*types.MsgUpdateCollectionStateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get Collection
	collection, err := s.Keeper.GetCollection(ctx, msg.CollectionId)
	if err != nil {
		return nil, err
	}

	// check that signer is collection admin
	if collection.Admin != msg.AdminAddress {
		return nil, errorsmod.Wrapf(types.ErrClaimUnauthorized, "collection admin %s, msg admin address %s", collection.Admin, msg.AdminAddress)
	}

	// update state
	collection.State = msg.State

	// persist the Collection
	if err := s.Keeper.CollectionPersistAndEmitEvents(ctx, collection); err != nil {
		return nil, err
	}

	return &types.MsgUpdateCollectionStateResponse{}, nil
}

// --------------------------
// UPDATE COLLECTION DATES
// --------------------------
func (s msgServer) UpdateCollectionDates(goCtx context.Context, msg *types.MsgUpdateCollectionDates) (*types.MsgUpdateCollectionDatesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get Collection
	collection, err := s.Keeper.GetCollection(ctx, msg.CollectionId)
	if err != nil {
		return nil, err
	}

	// check that signer is collection admin
	if collection.Admin != msg.AdminAddress {
		return nil, errorsmod.Wrapf(types.ErrClaimUnauthorized, "collection admin %s, msg admin address %s", collection.Admin, msg.AdminAddress)
	}

	// update state
	collection.StartDate = msg.StartDate
	collection.EndDate = msg.EndDate

	// persist the Collection
	if err := s.Keeper.CollectionPersistAndEmitEvents(ctx, collection); err != nil {
		return nil, err
	}

	return &types.MsgUpdateCollectionDatesResponse{}, nil
}

// --------------------------
// UPDATE COLLECTION PAYMENTS
// --------------------------
func (s msgServer) UpdateCollectionPayments(goCtx context.Context, msg *types.MsgUpdateCollectionPayments) (*types.MsgUpdateCollectionPaymentsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get Collection
	collection, err := s.Keeper.GetCollection(ctx, msg.CollectionId)
	if err != nil {
		return nil, err
	}

	// check that signer is collection admin
	if collection.Admin != msg.AdminAddress {
		return nil, errorsmod.Wrapf(types.ErrClaimUnauthorized, "collection admin %s, msg admin address %s", collection.Admin, msg.AdminAddress)
	}

	// check that entity exists
	_, entity, err := s.Keeper.EntityKeeper.ResolveEntity(ctx, collection.Entity)
	if err != nil {
		return nil, errorsmod.Wrapf(iidtypes.ErrDidDocumentNotFound, "for entity %s", collection.Entity)
	}

	// check that Evaluation Payment does not have 1155 payment
	if msg.Payments.Evaluation.Contract_1155Payment != nil {
		return nil, types.ErrCollectionEvalError
	}
	// check that Evaluation Payment does not have CW20 payments
	if len(msg.Payments.Evaluation.Cw20Payment) > 1 {
		return nil, types.ErrCollectionEvalCW20Error
	}

	// check that all payments accounts is part of entity module accounts
	if !msg.Payments.AccountsIsEntityAccounts(entity) {
		return nil, types.ErrCollNotEntityAcc
	}

	// update state
	collection.Payments = msg.Payments

	// persist the Collection
	if err := s.Keeper.CollectionPersistAndEmitEvents(ctx, collection); err != nil {
		return nil, err
	}

	return &types.MsgUpdateCollectionPaymentsResponse{}, nil
}

// --------------------------
// UPDATE COLLECTION STATE
// --------------------------
func (s msgServer) UpdateCollectionIntents(goCtx context.Context, msg *types.MsgUpdateCollectionIntents) (*types.MsgUpdateCollectionIntentsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get Collection
	collection, err := s.Keeper.GetCollection(ctx, msg.CollectionId)
	if err != nil {
		return nil, err
	}

	// check that signer is collection admin
	if collection.Admin != msg.AdminAddress {
		return nil, errorsmod.Wrapf(types.ErrClaimUnauthorized, "collection admin %s, msg admin address %s", collection.Admin, msg.AdminAddress)
	}

	// update Intents
	collection.Intents = msg.Intents

	// persist the Collection
	if err := s.Keeper.CollectionPersistAndEmitEvents(ctx, collection); err != nil {
		return nil, err
	}

	return &types.MsgUpdateCollectionIntentsResponse{}, nil
}

// --------------------------
// CLAIM INTENT
// --------------------------
func (s msgServer) ClaimIntent(goCtx context.Context, msg *types.MsgClaimIntent) (*types.MsgClaimIntentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the agent already has an active intent for this collection, if has throw error
	_, found := s.Keeper.GetActiveIntent(ctx, msg.AgentAddress, msg.CollectionId)
	if found {
		return nil, errorsmod.Wrapf(types.ErrIntentExists, "agent already has an active intent for collection %s", msg.CollectionId)
	}

	// Get Collection for Intent
	collection, err := s.Keeper.GetCollection(ctx, msg.CollectionId)
	if err != nil {
		return nil, err
	}

	// check that intents is allowed for the collection
	if collection.Intents == types.CollectionIntentOptions_deny {
		return nil, errorsmod.Wrapf(types.ErrIntentUnauthorized, "intents is not allowed for collection %s", msg.CollectionId)
	}

	agentAddress, err := sdk.AccAddressFromBech32(msg.AgentAddress)
	if err != nil {
		return nil, err
	}
	adminAddress, err := sdk.AccAddressFromBech32(collection.Admin)
	if err != nil {
		return nil, err
	}
	// get SubmitClaimAuthorization for agent to use for intent verification
	authzMsgType := sdk.MsgTypeURL(&types.MsgSubmitClaim{})
	authz, _ := s.Keeper.AuthzKeeper.GetAuthorization(ctx, agentAddress, adminAddress, authzMsgType)

	// if no authz then return error
	if authz == nil {
		return nil, errorsmod.Wrapf(types.ErrIntentUnauthorized, "agent %s does not have authz from this collection %s", msg.AgentAddress, msg.CollectionId)
	}
	// get authz constraints for type match
	var constraints []*types.SubmitClaimConstraints
	switch k := authz.(type) {
	case *types.SubmitClaimAuthorization:
		constraints = k.Constraints
	default:
		return nil, fmt.Errorf("existing Authorizations for route %s is not of type SubmitClaimAuthorization", authzMsgType)
	}
	// get constraint for collection id
	var constraint *types.SubmitClaimConstraints
	for _, con := range constraints {
		if con.CollectionId == msg.CollectionId {
			constraint = con
			break
		}
	}
	// if no authz constraint for collection id then return error
	if constraint == nil {
		return nil, errorsmod.Wrapf(types.ErrIntentUnauthorized, "agent %s does not have authz from this collection %s", msg.AgentAddress, msg.CollectionId)
	}

	// check that intent amount and cw20 payments are within max constraints
	if !types.IsCoinsInMaxConstraints(msg.Amount, constraint.MaxAmount) {
		return nil, errorsmod.Wrapf(types.ErrIntentUnauthorized, "intent amount is not within authz max constraints")
	}
	if !types.IsCW20PaymentsInMaxConstraints(msg.Cw20Payment, constraint.MaxCw20Payment) {
		return nil, errorsmod.Wrapf(types.ErrIntentUnauthorized, "intent cw20 payments is not within authz max constraints")
	}

	// if both amount and cw20 payments are empty then use default payments for APPROVAL
	if len(msg.Amount) == 0 && len(msg.Cw20Payment) == 0 {
		msg.Amount = collection.Payments.Approval.Amount
		msg.Cw20Payment = collection.Payments.Approval.Cw20Payment
	}

	// Get account used for APPROVAL payments on collection
	approvalAddress, err := sdk.AccAddressFromBech32(collection.Payments.Approval.Account)
	if err != nil {
		return nil, err
	}
	// Get escrow address
	escrow, err := sdk.AccAddressFromBech32(collection.EscrowAccount)
	if err != nil {
		return nil, err
	}

	// get intent id from params and update params
	params := s.Keeper.GetParams(ctx)
	intentID := params.IntentSequence
	params.IntentSequence++
	s.Keeper.SetParams(ctx, &params)

	// create Intent
	createdDate := ctx.BlockTime()
	expireAt := createdDate.Add(constraint.IntentDurationNs)
	intent := types.Intent{
		Id:            fmt.Sprint(intentID),
		AgentDid:      msg.AgentDid.Did(),
		AgentAddress:  msg.AgentAddress,
		CollectionId:  msg.CollectionId,
		CreatedAt:     &createdDate,
		ExpireAt:      &expireAt,
		Status:        types.IntentStatus_active,
		Amount:        msg.Amount,
		Cw20Payment:   msg.Cw20Payment,
		FromAddress:   approvalAddress.String(),
		EscrowAddress: escrow.String(),
	}
	// transfer the payments to escrow
	err = s.Keeper.TransferIntentPayments(ctx, approvalAddress, escrow, intent.Amount, intent.Cw20Payment)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to transfer payments to escrow")
	}

	// Save the intent and emit the events
	s.Keeper.SetIntent(ctx, intent)
	if err := ctx.EventManager().EmitTypedEvents(
		&types.IntentSubmittedEvent{
			Intent: &intent,
		},
	); err != nil {
		return nil, err
	}

	return &types.MsgClaimIntentResponse{
		IntentId: intent.Id,
		ExpireAt: intent.ExpireAt,
	}, nil
}

// --------------------------
// CREATE CLAIM AUTHORIZATION
// --------------------------
func (s msgServer) CreateClaimAuthorization(goCtx context.Context, msg *types.MsgCreateClaimAuthorization) (*types.MsgCreateClaimAuthorizationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the collection to verify it exists and get admin details
	collection, err := s.Keeper.GetCollection(ctx, msg.CollectionId)
	if err != nil {
		return nil, err
	}

	// Verify that the admin address in the message matches the collection's admin
	if collection.Admin != msg.AdminAddress {
		return nil, errorsmod.Wrapf(
			types.ErrClaimUnauthorized,
			"collection admin %s, msg admin address %s",
			collection.Admin,
			msg.AdminAddress,
		)
	}

	// get current owner of entity to pass to MsgGrantEntityAccountAuthz
	currentOwner, err := s.Keeper.EntityKeeper.GetCurrentOwner(ctx, collection.Entity)
	if err != nil {
		return nil, err
	}

	// Get grantee address for checking existing authorizations
	grantee, err := sdk.AccAddressFromBech32(msg.GranteeAddress)
	if err != nil {
		return nil, err
	}

	// Get granter address (admin account)
	granter, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return nil, err
	}

	// create the authorization
	// fist check the type of authorization (submit, evaluate)
	// then check if an authorization for that type already exists
	// if it does, append the new constraint to the existing authorization
	// if it doesn't, create a new authorization with the new constraint
	var authorization authz.Authorization
	switch msg.AuthType {
	case types.CreateClaimAuthorizationType_SUBMIT:
		// Create the new constraint
		newConstraint := &types.SubmitClaimConstraints{
			CollectionId:     msg.CollectionId,
			AgentQuota:       msg.AgentQuota,
			MaxAmount:        msg.MaxAmount,
			MaxCw20Payment:   msg.MaxCw20Payment,
			IntentDurationNs: msg.IntentDurationNs,
		}

		// Check for existing SubmitClaimAuthorization
		authzMsgType := sdk.MsgTypeURL(&types.MsgSubmitClaim{})
		existingAuth, _ := s.Keeper.AuthzKeeper.GetAuthorization(ctx, grantee, granter, authzMsgType)

		if existingAuth != nil {
			// check if existing auth is generic authorization
			// if so throw error as grantee already has a generic authorization with no constraints
			_, ok := existingAuth.(*authz.GenericAuthorization)
			if ok {
				return nil, errorsmod.Wrapf(
					sdkerrors.ErrInvalidRequest,
					"grantee %s already has a generic authorization from granter %s for authzMsgType %s, please remove the generic authorization before creating a new authorization with specific constraints",
					grantee.String(),
					granter.String(),
					authzMsgType,
				)
			}
			submitAuth, ok := existingAuth.(*types.SubmitClaimAuthorization)
			if ok {
				// Append the new constraint to existing constraints and create new authorization
				constraints := submitAuth.Constraints
				constraints = append(constraints, newConstraint)
				authorization = types.NewSubmitClaimAuthorization(msg.AdminAddress, constraints)
			} else {
				return nil, errorsmod.Wrapf(
					sdkerrors.ErrInvalidRequest,
					"existing authorization is not of type SubmitClaimAuthorization",
				)
			}
		} else {
			// Create new authorization with single constraint
			authorization = types.NewSubmitClaimAuthorization(msg.AdminAddress, []*types.SubmitClaimConstraints{newConstraint})
		}

	case types.CreateClaimAuthorizationType_EVALUATE:
		// Create the new constraint
		newConstraint := &types.EvaluateClaimConstraints{
			CollectionId:         msg.CollectionId,
			AgentQuota:           msg.AgentQuota,
			MaxCustomAmount:      msg.MaxAmount,
			MaxCustomCw20Payment: msg.MaxCw20Payment,
			BeforeDate:           msg.BeforeDate,
		}

		// Check for existing EvaluateClaimAuthorization
		authzMsgType := sdk.MsgTypeURL(&types.MsgEvaluateClaim{})
		existingAuth, _ := s.Keeper.AuthzKeeper.GetAuthorization(ctx, grantee, granter, authzMsgType)

		if existingAuth != nil {
			// check if existing auth is generic authorization
			// if so throw error as grantee already has a generic authorization with no constraints
			_, ok := existingAuth.(*authz.GenericAuthorization)
			if ok {
				return nil, errorsmod.Wrapf(
					sdkerrors.ErrInvalidRequest,
					"grantee %s already has a generic authorization from granter %s for authzMsgType %s, please remove the generic authorization before creating a new authorization with specific constraints",
					grantee.String(),
					granter.String(),
					authzMsgType,
				)
			}
			evalAuth, ok := existingAuth.(*types.EvaluateClaimAuthorization)
			if ok {
				// Append the new constraint to existing constraints and create new authorization
				constraints := evalAuth.Constraints
				constraints = append(constraints, newConstraint)
				authorization = types.NewEvaluateClaimAuthorization(msg.AdminAddress, constraints)
			} else {
				return nil, errorsmod.Wrapf(
					sdkerrors.ErrInvalidRequest,
					"existing authorization is not of type EvaluateClaimAuthorization",
				)
			}
		} else {
			// Create new authorization with single constraint
			authorization = types.NewEvaluateClaimAuthorization(msg.AdminAddress, []*types.EvaluateClaimConstraints{newConstraint})
		}

	case types.CreateClaimAuthorizationType_ALL:
		return nil, errorsmod.Wrapf(
			sdkerrors.ErrInvalidRequest,
			"cannot create both submission and evaluation authorization in a single request, use separate requests",
		)
	default:
		return nil, errorsmod.Wrapf(
			sdkerrors.ErrInvalidRequest,
			"unknown authorization type: %s",
			msg.AuthType,
		)
	}

	// Create the Any type for the authorization
	authAny, err := cdctypes.NewAnyWithValue(authorization)
	if err != nil {
		return nil, errorsmod.Wrapf(err, "failed to pack authorization into Any type")
	}

	// Create the MsgGrantEntityAccountAuthz to route to entity keeper
	grantMsg := entitytypes.MsgGrantEntityAccountAuthz{
		Id:             collection.Entity,
		Name:           entitytypes.EntityAdminAccountName,
		GranteeAddress: msg.GranteeAddress,
		Grant: authz.Grant{
			Authorization: authAny,
			Expiration:    msg.Expiration,
		},
		OwnerAddress: currentOwner,
	}

	// Route the authorization grant through the keeper
	err = s.Keeper.RouteGrantEntityAccountAuthz(ctx, &grantMsg)
	if err != nil {
		return nil, errorsmod.Wrapf(
			err,
			"failed to route authorization grant for grantee %s",
			msg.GranteeAddress,
		)
	}

	// Emit the event
	if err := ctx.EventManager().EmitTypedEvents(
		&types.ClaimAuthorizationCreatedEvent{
			Creator:      msg.CreatorAddress,
			CreatorDid:   msg.CreatorDid.Did(),
			Grantee:      msg.GranteeAddress,
			Admin:        msg.AdminAddress,
			CollectionId: msg.CollectionId,
			AuthType:     msg.AuthType.String(),
		},
	); err != nil {
		return nil, err
	}

	return &types.MsgCreateClaimAuthorizationResponse{}, nil
}

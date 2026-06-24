package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/authz"
	"github.com/ixofoundation/ixo-blockchain/v8/x/claims/types"
	entitytypes "github.com/ixofoundation/ixo-blockchain/v8/x/entity/types"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v8/x/iid/types"
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

// TODO: ADD 721 capabilities to claims and payments also.

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

	// check that Evaluation Payment does not have CW20 payments
	if len(msg.Payments.Evaluation.Cw20Payment) > 1 {
		return nil, types.ErrCollectionEvalCW20Error
	}
	// check that Evaluation Payment does not have CW1155 payments
	if len(msg.Payments.Evaluation.Cw1155Payment) > 0 {
		return nil, types.ErrCollectionEvalCW1155Error
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
		Id:                          fmt.Sprint(collectionSequence),
		Entity:                      msg.Entity,
		Admin:                       admin,
		Protocol:                    msg.Protocol,
		StartDate:                   msg.StartDate,
		EndDate:                     msg.EndDate,
		Quota:                       msg.Quota,
		Count:                       0,
		Evaluated:                   0,
		Approved:                    0,
		Rejected:                    0,
		Disputed:                    0,
		Invalidated:                 0,
		State:                       msg.State,
		Payments:                    msg.Payments,
		Intents:                     msg.Intents,
		ServiceAgentDepositRequired: msg.ServiceAgentDepositRequired,
		EvaluatorDepositRequired:    msg.EvaluatorDepositRequired,
		DisputeDepositAmount:        msg.DisputeDepositAmount,
		Adjudicators:                msg.Adjudicators,
		PenaltyAmountPerDispute:     msg.PenaltyAmountPerDispute,
		MinDepositPeriod:            msg.MinDepositPeriod,
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

	// Dispute-gate (v7): SA can't submit new claims while they have any
	// OPEN dispute against them on this collection. Allows the SA to top
	// up their deposit but not perform new work until the dispute is
	// adjudicated.
	if s.Keeper.HasActiveDisputeAgainstAgent(ctx, msg.CollectionId, msg.AgentAddress) {
		return nil, errorsmod.Wrapf(types.ErrAgentHasActiveDispute,
			"submitter %s has open disputes on collection %s",
			msg.AgentAddress, msg.CollectionId)
	}
	// Performance-deposit gate (v7): if the collection requires a
	// service-agent deposit, the SA's running balance must already meet it.
	// Top-up happens via MsgAddPerformanceDeposit; this is a check, not a
	// debit. The deposit balance only moves on adjudicated dispute losses
	// or explicit withdrawal.
	if !collection.ServiceAgentDepositRequired.IsZero() {
		if !s.Keeper.HasAgentMetDepositRequirement(ctx, msg.CollectionId, msg.AgentAddress, collection.ServiceAgentDepositRequired) {
			return nil, errorsmod.Wrapf(types.ErrAgentDepositInsufficient,
				"submitter %s on collection %s needs balance ≥ %s",
				msg.AgentAddress, msg.CollectionId, collection.ServiceAgentDepositRequired)
		}
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
		Amount:        msg.Amount,
		Cw20Payment:   msg.Cw20Payment,
		Cw1155Payment: msg.Cw1155Payment,
		UseIntent:     msg.UseIntent,
	}

	// if intent then override payments, add intent id and update APPROVAL payment to GUARANTEED
	if msg.UseIntent {
		// validate member_address consistency between msg and intent (strict equality,
		// including both being empty). Prevents a user from spuriously attributing a
		// claim to a member when the intent has no member context, which would later
		// cause an incorrect budget restore on rejection.
		if msg.MemberAddress != intent.MemberAddress {
			return nil, errorsmod.Wrapf(types.ErrMemberAddressMismatch, "msg member_address %s does not match intent member_address %s", msg.MemberAddress, intent.MemberAddress)
		}

		claim.Amount = intent.Amount
		claim.Cw20Payment = intent.Cw20Payment
		claim.Cw1155Payment = intent.Cw1155Payment
		claim.Cw1155IntentPayment = intent.Cw1155IntentPayment
		claim.MemberAddress = intent.MemberAddress

		// if any payment is not empty then APPROVAL payment become GUARANTEED as funds is in escrow account
		// if all payments is empty or all amounts is 0 then APPROVAL payment stays NO_PAYMENT since no funds are in escrow account
		if !intent.Amount.IsZero() || !types.IsZeroCW20Payments(intent.Cw20Payment) || !types.IsZeroCW1155Payments(intent.Cw1155Payment) {
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
		// without intent there is no member context — reject any provided member_address
		if msg.MemberAddress != "" {
			return nil, errorsmod.Wrapf(types.ErrMemberAddressMismatch, "member_address provided without use_intent")
		}
		// if no intent used, check if collection approval payment is oracle payment then only native coins allowed
		if collection.Payments.Approval.IsOraclePayment && (!types.IsZeroCW20Payments(claim.Cw20Payment) || !types.IsZeroCW1155Payments(claim.Cw1155Payment)) {
			return nil, types.ErrOraclePaymentOnlyNative
		}
	}

	// update count for collection and persist
	collection.Count++
	if err := s.Keeper.CollectionPersistAndEmitEvents(ctx, collection); err != nil {
		return nil, err
	}

	// start payout process for claim submission
	if err = processPayment(ctx, *s.Keeper, agent, collection.Payments.Submission, types.PaymentType_submission, &claim, collection, false, []*types.CW1155IntentPayment{}); err != nil {
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

	// Re-evaluation gate. A claim with an existing evaluation can only be
	// re-evaluated when its current status is FLAGGED (the non-terminal escape
	// hatch). Any other prior status is terminal and locks the claim.
	//
	// The same agent that flagged a claim is allowed to finalise their own
	// flag (e.g. they got more information later). They are not allowed to
	// flag a claim more than once — re-flagging adds no new on-chain state
	// (ErrSelfReFlag). The check looks at both the current evaluation and
	// every entry in evaluation_history, so a flag-bomb across an
	// intervening flag from another agent (tester flag → bob flag → tester
	// flag again) is also blocked. Terminal priors are blocked by the
	// duplicate-evaluation check above, so every entry we inspect for the
	// self-reflag rule is itself a flag.
	priorEvaluation := claim.Evaluation
	isReEvaluation := priorEvaluation != nil
	if isReEvaluation {
		if priorEvaluation.Status != types.EvaluationStatus_flagged {
			return nil, errorsmod.Wrapf(types.ErrClaimDuplicateEvaluation, "id %s", claim.ClaimId)
		}
		if msg.Status == types.EvaluationStatus_flagged {
			if priorEvaluation.AgentAddress == msg.AgentAddress {
				return nil, errorsmod.Wrapf(types.ErrSelfReFlag, "agent %s already flagged claim %s", msg.AgentAddress, claim.ClaimId)
			}
			for _, h := range claim.EvaluationHistory {
				if h.AgentAddress == msg.AgentAddress {
					return nil, errorsmod.Wrapf(types.ErrSelfReFlag, "agent %s already flagged claim %s", msg.AgentAddress, claim.ClaimId)
				}
			}
		}
	}
	isFlagged := msg.Status == types.EvaluationStatus_flagged

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

	// Dispute-gate (v7): evaluator can't submit new evaluations while any
	// OPEN dispute targets one of their prior evaluations on this collection.
	if s.Keeper.HasActiveDisputeAgainstAgent(ctx, claim.CollectionId, msg.AgentAddress) {
		return nil, errorsmod.Wrapf(types.ErrAgentHasActiveDispute,
			"evaluator %s has open disputes on collection %s",
			msg.AgentAddress, claim.CollectionId)
	}
	// Performance-deposit gate (v7): if the collection requires an
	// evaluator deposit, balance must meet it. INVALIDATED still gates —
	// any work on the claim counts.
	if !collection.EvaluatorDepositRequired.IsZero() {
		if !s.Keeper.HasAgentMetDepositRequirement(ctx, claim.CollectionId, msg.AgentAddress, collection.EvaluatorDepositRequired) {
			return nil, errorsmod.Wrapf(types.ErrAgentDepositInsufficient,
				"evaluator %s on collection %s needs balance ≥ %s",
				msg.AgentAddress, claim.CollectionId, collection.EvaluatorDepositRequired)
		}
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
		Cw1155Payment:     msg.Cw1155Payment,
	}

	// if intent on claim then override payments with claim payments as it used the intent
	if claim.UseIntent {
		evaluation.Amount = claim.Amount
		evaluation.Cw20Payment = claim.Cw20Payment
		evaluation.Cw1155Payment = claim.Cw1155Payment
		evaluation.Cw1155IntentPayment = claim.Cw1155IntentPayment
	}

	// Evaluator payout. Skip for INVALIDATED and for FLAGGED
	// (a flag is not a terminal output — no payment, no Evaluated++ on flag).
	// On a finalising re-evaluation the finaliser receives the evaluator
	// payment (msg.AgentAddress), which is what we want.
	if msg.Status != types.EvaluationStatus_invalidated && !isFlagged {
		if err = processPayment(ctx, *s.Keeper, evalAgent, collection.Payments.Evaluation, types.PaymentType_evaluation, &claim, collection, false, []*types.CW1155IntentPayment{}); err != nil {
			return nil, err
		}

		// update evaluated count for collection
		collection.Evaluated++
	}

	// Intent-funded escrow handling. On any non-approved terminal status with
	// an intent, refund escrow → approval account, mark approval NO_PAYMENT
	// and restore member budget. FLAGGED keeps funds in escrow (the claim is
	// still in flight and a subsequent finaliser may still APPROVE).
	if msg.Status != types.EvaluationStatus_approved && !isFlagged && claim.UseIntent {
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
		_, err = s.Keeper.TransferIntentPayments(ctx, escrow, approvalAddress, claim.Amount, claim.Cw20Payment, claim.Cw1155Payment, claim.Cw1155IntentPayment)
		if err != nil {
			return nil, err
		}
		// Update payment status to no payment again as was guaranteed with intent
		err = updatePaymentStatus(types.PaymentType_approval, &claim, types.PaymentStatus_no_payment)
		if err != nil {
			return nil, err
		}

		// Restore member budget if this claim was on behalf of a team member
		if claim.MemberAddress != "" {
			if err := s.Keeper.RestoreMemberBudget(ctx, claim.CollectionId, claim.MemberAddress, claim.Amount, claim.Cw20Payment); err != nil {
				return nil, err
			}
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
			approvedPayment.Account = collection.EscrowAccount
			approvedPayment.Amount = claim.Amount
			approvedPayment.Cw20Payment = claim.Cw20Payment
			approvedPayment.Cw1155Payment = claim.Cw1155Payment
		} else {
			// if any amount length is not zero, it means agent set custom amount that was authenticated
			// through authZ constraints to be valid since all evaluations must be done by collections module account through authz
			if len(msg.Amount) > 0 || len(msg.Cw20Payment) > 0 || len(msg.Cw1155Payment) > 0 {
				approvedPayment.Amount = msg.Amount
				approvedPayment.Cw20Payment = msg.Cw20Payment
				approvedPayment.Cw1155Payment = msg.Cw1155Payment
			}
			// if no intent used, check if collection approval payment is oracle payment then only native coins allowed
			if collection.Payments.Approval.IsOraclePayment && (!types.IsZeroCW20Payments(approvedPayment.Cw20Payment) || !types.IsZeroCW1155Payments(approvedPayment.Cw1155Payment)) {
				return nil, types.ErrOraclePaymentOnlyNative
			}
		}
		if err = processPayment(ctx, *s.Keeper, claimAgent, approvedPayment, types.PaymentType_approval, &claim, collection, claim.UseIntent, claim.Cw1155IntentPayment); err != nil {
			return nil, err
		}
	} else if msg.Status == types.EvaluationStatus_rejected {
		// payout process for evaluation rejected to claim agent
		collection.Rejected++
		if err = processPayment(ctx, *s.Keeper, claimAgent, collection.Payments.Rejection, types.PaymentType_rejection, &claim, collection, false, []*types.CW1155IntentPayment{}); err != nil {
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
	} else if isFlagged {
		// Flagged is non-terminal: no payment, no terminal-state counters.
		// `flagged` is a cumulative event counter (every flag, including
		// flag-after-flag chains). `flagged_active` is the number of claims
		// currently sitting in FLAGGED state — only increments on the
		// transition into FLAGGED, not on a subsequent re-flag.
		collection.Flagged++
		if !isReEvaluation {
			collection.FlaggedActive++
		}
	}

	// Transitioning out of FLAGGED to a terminal status: decrement the
	// active-flag count. Re-flag (FLAGGED → FLAGGED) leaves it untouched.
	if isReEvaluation && !isFlagged && collection.FlaggedActive > 0 {
		collection.FlaggedActive--
	}

	if err := s.Keeper.CollectionPersistAndEmitEvents(ctx, collection); err != nil {
		return nil, err
	}

	// On re-evaluation, move the prior evaluation into history (chronological
	// order, oldest first) before replacing the current evaluation. First-time
	// evaluations leave evaluation_history empty.
	if isReEvaluation {
		claim.EvaluationHistory = append(claim.EvaluationHistory, priorEvaluation)
	}
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

	// Dispute filing is open to anyone with a valid registered IID (the IID
	// ante still validates that msg.AgentAddress is a key on msg.AgentDid).
	// The economic gate is the `dispute_deposit_amount` the disputer stakes
	// inline — that, rather than authz/controller permissions, prevents
	// frivolous filings. Prior versions required the disputer to be the
	// collection admin, an entity controller, or hold submit/evaluate authz
	// on the collection; that gate is removed in v7 now that disputers
	// have skin in the game.

	// v7 dispute lifecycle. Determine the targeted agent so we can write
	// the active-dispute index and (on AWARDED later) slash their balance.
	var targetAgent string
	switch msg.TargetRole {
	case types.DisputeTargetRole_target_submitter:
		targetAgent = claim.AgentAddress
	case types.DisputeTargetRole_target_evaluator:
		if claim.Evaluation == nil {
			return nil, errorsmod.Wrapf(types.ErrDisputeTargetEvaluatorNoEvaluation,
				"claim %s has no evaluation; cannot dispute evaluator", claim.ClaimId)
		}
		// FLAGGED is a non-terminal escape-hatch — the evaluator explicitly
		// declined to finalise. Disputing a flag would punish honest
		// uncertainty and discourage flagging, which is the opposite of the
		// intent. The path to resolve a FLAGGED claim is another evaluator
		// re-evaluating to a terminal status; that terminal evaluation can
		// then itself be disputed.
		if claim.Evaluation.Status == types.EvaluationStatus_flagged {
			return nil, errorsmod.Wrapf(types.ErrDisputeTargetEvaluatorFlagged,
				"claim %s evaluation is FLAGGED", claim.ClaimId)
		}
		targetAgent = claim.Evaluation.AgentAddress
	default:
		// ValidateBasic already rejects UNSPECIFIED; defensive only.
		return nil, types.ErrDisputeTargetRoleInvalid
	}

	// (subject_id, target_role) gating: no OPEN duplicate; AWARDED permanently
	// blocks. DISMISSED allows a new filing.
	if err := s.Keeper.CanFileNewDisputeForSubject(ctx, msg.SubjectId, msg.TargetRole); err != nil {
		return nil, err
	}

	// Lock the disputer's stake into the collection escrow (v7). The
	// dispute is recorded with the exact amount snapshot, so subsequent
	// admin updates to dispute_deposit_amount don't retroactively change
	// what's at stake on this dispute.
	//
	// Guard: refuse to lock funds when the collection has no adjudicator
	// whitelist configured — otherwise the deposit could not be released
	// through adjudication (MsgAdjudicateDispute requires the adjudicator
	// DID to be in the whitelist). Without this check an admin who never
	// set up the whitelist could end up with disputer funds permanently
	// stuck in escrow.
	disputeDeposit := collection.DisputeDepositAmount
	if !disputeDeposit.IsZero() {
		if len(collection.Adjudicators) == 0 {
			return nil, errorsmod.Wrap(types.ErrAdjudicationNotConfigured,
				"refusing to lock dispute deposit on a collection with no adjudicators")
		}
		disputerAddr, err := sdk.AccAddressFromBech32(msg.AgentAddress)
		if err != nil {
			return nil, err
		}
		escrowAddr, err := sdk.AccAddressFromBech32(collection.EscrowAccount)
		if err != nil {
			return nil, errorsmod.Wrapf(types.ErrInternalError, "invalid escrow address: %s", err)
		}
		if err := s.Keeper.BankKeeper.SendCoins(ctx, disputerAddr, escrowAddr, disputeDeposit); err != nil {
			return nil, err
		}
	}

	submittedAt := ctx.BlockTime()
	dispute := types.Dispute{
		SubjectId:       msg.SubjectId,
		Type:            msg.DisputeType,
		Data:            msg.Data,
		TargetRole:      msg.TargetRole,
		DisputerAddress: msg.AgentAddress,
		DisputerDid:     msg.AgentDid.Did(),
		DisputeDeposit:  disputeDeposit,
		SubmittedAt:     &submittedAt,
		Status:          types.DisputeStatus_dispute_open,
	}
	s.Keeper.SetDispute(ctx, dispute)
	// Subject index points the (subject_id, role) tuple at this dispute's
	// proof so the canonical-status lookup is one-hop. Overwrites a prior
	// DISMISSED pointer if present.
	s.Keeper.SetDisputeSubjectIndex(ctx, msg.SubjectId, msg.TargetRole, msg.Data.Proof)
	// Active-dispute presence index for the target agent, used by submit /
	// evaluate / withdraw gates.
	s.Keeper.SetActiveDispute(ctx, claim.CollectionId, targetAgent, msg.SubjectId)

	// Collection-level counter.
	collection.DisputesOpen++
	if err := s.Keeper.CollectionPersistAndEmitEvents(ctx, collection); err != nil {
		return nil, err
	}

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
	err = payout(ctx, *s.Keeper, msg.Inputs, msg.Outputs, msg.PaymentType, &claim, collection, msg.ReleaseDate, msg.Cw20Payment, fromAddress, toAddress, []*types.CW20Output{}, msg.Cw1155Payment, []*types.CW1155IntentPayment{})
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

	// check that Evaluation Payment does not have CW20 payments
	if len(msg.Payments.Evaluation.Cw20Payment) > 1 {
		return nil, types.ErrCollectionEvalCW20Error
	}
	// check that Evaluation Payment does not have CW1155 payments
	if len(msg.Payments.Evaluation.Cw1155Payment) > 1 {
		return nil, types.ErrCollectionEvalCW1155Error
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
// UPDATE COLLECTION INTENTS
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

// UpdateCollectionQuota changes the maximum claim count for a collection.
// Validation rule: the new quota must be either 0 (unlimited) or ≥ the
// collection's current `count`. Setting a quota below already-submitted
// claims would retroactively invalidate the gate on every prior submission,
// so we reject it at handler time rather than silently lock the collection.
func (s msgServer) UpdateCollectionQuota(goCtx context.Context, msg *types.MsgUpdateCollectionQuota) (*types.MsgUpdateCollectionQuotaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	collection, err := s.Keeper.GetCollection(ctx, msg.CollectionId)
	if err != nil {
		return nil, err
	}

	if collection.Admin != msg.AdminAddress {
		return nil, errorsmod.Wrapf(types.ErrClaimUnauthorized, "collection admin %s, msg admin address %s", collection.Admin, msg.AdminAddress)
	}

	// Reject quota < count (unless 0, which means unlimited). Cannot
	// retroactively cap below claims already in the collection.
	if msg.Quota != 0 && msg.Quota < collection.Count {
		return nil, errorsmod.Wrapf(types.ErrCollectionQuotaBelowCount,
			"new quota %d is below current count %d", msg.Quota, collection.Count)
	}

	collection.Quota = msg.Quota

	if err := s.Keeper.CollectionPersistAndEmitEvents(ctx, collection); err != nil {
		return nil, err
	}

	return &types.MsgUpdateCollectionQuotaResponse{}, nil
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

	// member_address presence must match collection setup:
	//   - team collection (has member budgets): member_address required
	//   - individual collection (no member budgets): member_address must be empty
	hasMemberBudgets := s.Keeper.HasMemberBudgets(ctx, msg.CollectionId)
	if hasMemberBudgets && msg.MemberAddress == "" {
		return nil, types.ErrMemberAddressRequired
	}
	if !hasMemberBudgets && msg.MemberAddress != "" {
		return nil, types.ErrMemberAddressNotAllowed
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
	// get constraint matching collection_id AND member_address (strict equality,
	// including both being empty for individual subscriptions)
	var constraint *types.SubmitClaimConstraints
	for _, con := range constraints {
		if con.CollectionId != msg.CollectionId {
			continue
		}
		// member_address must match exactly: both empty for individual subscriptions,
		// or both equal to the specific member for team subscriptions. This prevents
		// silently picking a member-tagged constraint when the oracle didn't attribute
		// the intent to a member, or vice versa.
		if con.MemberAddress != msg.MemberAddress {
			continue
		}
		constraint = con
		break
	}
	// if no authz constraint for collection id then return error
	if constraint == nil {
		return nil, errorsmod.Wrapf(types.ErrIntentUnauthorized, "agent %s does not have authz from this collection %s", msg.AgentAddress, msg.CollectionId)
	}

	// check that intent amounts are within max constraints
	if !types.IsCoinsInMaxConstraints(msg.Amount, constraint.MaxAmount) {
		return nil, errorsmod.Wrapf(types.ErrIntentUnauthorized, "intent amount is not within authz max constraints")
	}
	if !types.IsCW20PaymentsInMaxConstraints(msg.Cw20Payment, constraint.MaxCw20Payment) {
		return nil, errorsmod.Wrapf(types.ErrIntentUnauthorized, "intent cw20 payments is not within authz max constraints")
	}
	if !types.IsCW1155PaymentsInMaxConstraints(msg.Cw1155Payment, constraint.MaxCw1155Payment) {
		return nil, errorsmod.Wrapf(types.ErrIntentUnauthorized, "intent cw1155 payments is not within authz max constraints")
	}

	// if all payments are empty then use default payments for APPROVAL
	if len(msg.Amount) == 0 && len(msg.Cw20Payment) == 0 && len(msg.Cw1155Payment) == 0 {
		msg.Amount = collection.Payments.Approval.Amount
		msg.Cw20Payment = collection.Payments.Approval.Cw20Payment
		msg.Cw1155Payment = collection.Payments.Approval.Cw1155Payment
	}

	// Member budget check and deduction. Constraint matching above already enforced
	// that constraint.MemberAddress == msg.MemberAddress, and the early guard ensured
	// msg.MemberAddress is non-empty when collection has member budgets.
	if hasMemberBudgets {
		budget, err := s.Keeper.GetMemberBudget(ctx, msg.CollectionId, msg.MemberAddress)
		if err != nil {
			return nil, errorsmod.Wrapf(types.ErrMemberBudgetNotFound, "member %s not found on collection %s", msg.MemberAddress, msg.CollectionId)
		}

		// Lazy period reset
		s.Keeper.TryResetMemberBudgetPeriod(ctx, &budget)

		// Check native coin budget
		if len(msg.Amount) > 0 && !msg.Amount.IsZero() {
			remaining, hasNeg := budget.PeriodSpendLimit.SafeSub(budget.PeriodSpent...)
			if hasNeg {
				return nil, errorsmod.Wrapf(types.ErrMemberBudgetExceeded, "member %s has no remaining budget", msg.MemberAddress)
			}
			if !msg.Amount.IsAllLTE(remaining) {
				return nil, errorsmod.Wrapf(types.ErrMemberBudgetExceeded, "member %s intent amount exceeds remaining budget", msg.MemberAddress)
			}
			budget.PeriodSpent = budget.PeriodSpent.Add(msg.Amount...)
		}

		// Check CW20 budget
		if len(msg.Cw20Payment) > 0 {
			for _, payment := range msg.Cw20Payment {
				if payment.Amount == 0 {
					continue
				}
				var limitAmount uint64
				for _, limit := range budget.PeriodCw20SpendLimit {
					if limit.Address == payment.Address {
						limitAmount = limit.Amount
						break
					}
				}
				var spentAmount uint64
				for _, spent := range budget.PeriodCw20Spent {
					if spent.Address == payment.Address {
						spentAmount = spent.Amount
						break
					}
				}
				// Guard against uint64 underflow: if spent already meets or exceeds
				// limit, there is no remaining budget. Without this check, the
				// subtraction wraps around to a huge number and the comparison passes.
				if spentAmount >= limitAmount || payment.Amount > limitAmount-spentAmount {
					return nil, errorsmod.Wrapf(types.ErrMemberBudgetExceeded, "member %s cw20 intent amount exceeds remaining budget for %s", msg.MemberAddress, payment.Address)
				}
				// Deduct CW20 spent
				found := false
				for i, spent := range budget.PeriodCw20Spent {
					if spent.Address == payment.Address {
						budget.PeriodCw20Spent[i].Amount += payment.Amount
						found = true
						break
					}
				}
				if !found {
					budget.PeriodCw20Spent = append(budget.PeriodCw20Spent, &types.CW20Payment{
						Address: payment.Address,
						Amount:  payment.Amount,
					})
				}
			}
		}

		if err := s.Keeper.SetMemberBudgetAndEmitUpdatedEvent(ctx, budget); err != nil {
			return nil, err
		}
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

	// transfer the payments to escrow
	cw1155IntentPayments, err := s.Keeper.TransferIntentPayments(ctx, approvalAddress, escrow, msg.Amount, msg.Cw20Payment, msg.Cw1155Payment, []*types.CW1155IntentPayment{})
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to transfer payments to escrow")
	}

	// create Intent
	createdDate := ctx.BlockTime()
	expireAt := createdDate.Add(constraint.IntentDurationNs)
	intent := types.Intent{
		Id:                  fmt.Sprint(intentID),
		AgentDid:            msg.AgentDid.Did(),
		AgentAddress:        msg.AgentAddress,
		CollectionId:        msg.CollectionId,
		CreatedAt:           &createdDate,
		ExpireAt:            &expireAt,
		Status:              types.IntentStatus_active,
		Amount:              msg.Amount,
		Cw20Payment:         msg.Cw20Payment,
		FromAddress:         approvalAddress.String(),
		EscrowAddress:       escrow.String(),
		Cw1155Payment:       msg.Cw1155Payment,
		Cw1155IntentPayment: cw1155IntentPayments,
		MemberAddress:       msg.MemberAddress,
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
			MaxCw1155Payment: msg.MaxCw1155Payment,
			IntentDurationNs: msg.IntentDurationNs,
			MemberAddress:    msg.MemberAddress,
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
			CollectionId:           msg.CollectionId,
			AgentQuota:             msg.AgentQuota,
			MaxCustomAmount:        msg.MaxAmount,
			MaxCustomCw20Payment:   msg.MaxCw20Payment,
			MaxCustomCw1155Payment: msg.MaxCw1155Payment,
			BeforeDate:             msg.BeforeDate,
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

// --------------------------
// SET COLLECTION MEMBERS
// --------------------------
func (s msgServer) SetCollectionMembers(goCtx context.Context, msg *types.MsgSetCollectionMembers) (*types.MsgSetCollectionMembersResponse, error) {
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

	now := ctx.BlockTime()

	for _, member := range msg.Members {
		// Check if member budget already exists
		existingBudget, getErr := s.Keeper.GetMemberBudget(ctx, msg.CollectionId, member.MemberAddress)
		isNew := getErr != nil

		budget := types.MemberBudget{
			CollectionId:         msg.CollectionId,
			MemberAddress:        member.MemberAddress,
			Period:               member.Period,
			PeriodSpendLimit:     member.PeriodSpendLimit,
			PeriodCw20SpendLimit: member.PeriodCw20SpendLimit,
		}

		if !isNew {
			// Existing member - preserve period_spent and period_reset_at unless reset requested
			if member.ResetPeriodSpent {
				budget.PeriodSpent = sdk.Coins{}
				budget.PeriodCw20Spent = nil
				resetAt := now.Add(member.Period)
				budget.PeriodResetAt = &resetAt
			} else {
				budget.PeriodSpent = existingBudget.PeriodSpent
				budget.PeriodCw20Spent = existingBudget.PeriodCw20Spent
				budget.PeriodResetAt = existingBudget.PeriodResetAt
			}
		} else {
			// New member - start fresh
			budget.PeriodSpent = sdk.Coins{}
			budget.PeriodCw20Spent = nil
			resetAt := now.Add(member.Period)
			budget.PeriodResetAt = &resetAt
		}

		// New members get MemberBudgetCreatedEvent (one-time emission, indexer
		// uses INSERT). Existing member updates go through the helper which
		// emits MemberBudgetUpdatedEvent (indexer uses UPDATE).
		if isNew {
			s.Keeper.SetMemberBudget(ctx, budget)
			if err := ctx.EventManager().EmitTypedEvent(
				&types.MemberBudgetCreatedEvent{Budget: &budget},
			); err != nil {
				return nil, err
			}
		} else {
			if err := s.Keeper.SetMemberBudgetAndEmitUpdatedEvent(ctx, budget); err != nil {
				return nil, err
			}
		}
	}

	return &types.MsgSetCollectionMembersResponse{}, nil
}

// --------------------------
// REMOVE COLLECTION MEMBERS
// --------------------------
// Removes one or more member budgets from a collection. Does NOT revoke any
// existing claim authorizations the members granted to oracles — admin should
// do that separately if needed. New intents from the removed members will fail
// at the GetMemberBudget lookup.
//
// Edge case: if a member is removed while they have unresolved intents/claims,
// budget restoration on rejection or expiration is silently skipped (member
// budget no longer exists). If the member is later re-added with a fresh budget
// before those intents resolve, the restore would target the new budget — this
// could effectively reduce the new period's spent amount by amounts that belong
// to the old period. This is an accepted trade-off for v1; admins should avoid
// removing members with active intents in flight.
func (s msgServer) RemoveCollectionMembers(goCtx context.Context, msg *types.MsgRemoveCollectionMembers) (*types.MsgRemoveCollectionMembersResponse, error) {
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

	for _, memberAddress := range msg.MemberAddresses {
		// Verify the member budget exists before removing — also gives us the
		// final state to include on the event for the indexer.
		budget, err := s.Keeper.GetMemberBudget(ctx, msg.CollectionId, memberAddress)
		if err != nil {
			return nil, errorsmod.Wrapf(types.ErrMemberBudgetNotFound, "member %s not found on collection %s", memberAddress, msg.CollectionId)
		}
		if err := s.Keeper.RemoveMemberBudgetAndEmitEvent(ctx, budget); err != nil {
			return nil, err
		}
	}

	return &types.MsgRemoveCollectionMembersResponse{}, nil
}

// --------------------------
// UPDATE COLLECTION DISPUTE CONFIG
// --------------------------
// Replaces the full dispute/performance-deposit config on a collection.
// Does NOT affect in-flight disputes — each dispute snapshots the values
// it cares about at filing / adjudication time (DisputeDeposit on the
// Dispute record, slash destinations resolved per adjudication).
//
// Validation is delegated to types.ValidateCollectionDisputeConfig via
// MsgUpdateCollectionDisputeConfig.ValidateBasic; the handler only enforces
// authz (admin signer matches collection.admin) and persists.
func (s msgServer) UpdateCollectionDisputeConfig(goCtx context.Context, msg *types.MsgUpdateCollectionDisputeConfig) (*types.MsgUpdateCollectionDisputeConfigResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	collection, err := s.Keeper.GetCollection(ctx, msg.CollectionId)
	if err != nil {
		return nil, err
	}
	if collection.Admin != msg.AdminAddress {
		return nil, errorsmod.Wrapf(types.ErrClaimUnauthorized, "collection admin %s, msg admin address %s", collection.Admin, msg.AdminAddress)
	}

	// Guard: refuse to clear the adjudicator whitelist while there are
	// open disputes on the collection. Otherwise those disputes — which
	// may hold locked disputer deposits — can never be adjudicated, and
	// the deposits are stranded in escrow forever. Reducing the whitelist
	// (but not emptying it) is fine; new adjudicators can still resolve
	// existing open disputes.
	if len(msg.Adjudicators) == 0 && collection.DisputesOpen > 0 {
		return nil, errorsmod.Wrapf(types.ErrAdjudicationNotConfigured,
			"cannot clear adjudicators while %d open dispute(s) exist on collection %s",
			collection.DisputesOpen, collection.Id)
	}

	collection.ServiceAgentDepositRequired = msg.ServiceAgentDepositRequired
	collection.EvaluatorDepositRequired = msg.EvaluatorDepositRequired
	collection.DisputeDepositAmount = msg.DisputeDepositAmount
	collection.Adjudicators = msg.Adjudicators
	collection.PenaltyAmountPerDispute = msg.PenaltyAmountPerDispute
	collection.MinDepositPeriod = msg.MinDepositPeriod

	if err := s.Keeper.CollectionPersistAndEmitEvents(ctx, collection); err != nil {
		return nil, err
	}

	return &types.MsgUpdateCollectionDisputeConfigResponse{}, nil
}

// --------------------------
// ADD PERFORMANCE DEPOSIT
// --------------------------
// Agent-funded top-up of their performance-deposit balance for a collection.
// Funds move agent → collection.escrow_account. Permitted regardless of
// active disputes — top-up is needed to clear arrears caused by a slash.
func (s msgServer) AddPerformanceDeposit(goCtx context.Context, msg *types.MsgAddPerformanceDeposit) (*types.MsgAddPerformanceDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	agentAddr, err := sdk.AccAddressFromBech32(msg.AgentAddress)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid agent address (%s)", err)
	}

	balance, err := s.Keeper.AddPerformanceDeposit(ctx, msg.CollectionId, agentAddr, msg.Amount)
	if err != nil {
		return nil, err
	}

	return &types.MsgAddPerformanceDepositResponse{NewBalance: balance.Amount}, nil
}

// --------------------------
// WITHDRAW PERFORMANCE DEPOSIT
// --------------------------
// Pulls some or all of an agent's deposit balance back to their wallet.
// Blocked while the agent has any OPEN dispute targeting them on this
// collection (gate enforced in keeper.WithdrawPerformanceDeposit).
func (s msgServer) WithdrawPerformanceDeposit(goCtx context.Context, msg *types.MsgWithdrawPerformanceDeposit) (*types.MsgWithdrawPerformanceDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	agentAddr, err := sdk.AccAddressFromBech32(msg.AgentAddress)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid agent address (%s)", err)
	}

	withdrawn, remaining, err := s.Keeper.WithdrawPerformanceDeposit(ctx, msg.CollectionId, agentAddr, msg.Amount)
	if err != nil {
		return nil, err
	}

	return &types.MsgWithdrawPerformanceDepositResponse{
		Withdrawn:        withdrawn,
		RemainingBalance: remaining,
	}, nil
}

// --------------------------
// ADJUDICATE DISPUTE
// --------------------------
// Settles an OPEN dispute by AWARDED or DISMISSED, applying the penalty
// math (80/20 winner / adjudicator split, configurable), tracking actual
// vs intended payouts (since loser balance may be short), and clearing
// the active-dispute index entry so the targeted agent is unblocked.
//
// Authorization: adjudicator_did must be on the collection's whitelist, and
// adjudicator_address must be authorized for that DID via either path
// (entity account OR DID-registered key — see AuthorizeAdjudicator).
func (s msgServer) AdjudicateDispute(goCtx context.Context, msg *types.MsgAdjudicateDispute) (*types.MsgAdjudicateDisputeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 1. Fetch dispute & claim & collection.
	dispute, err := s.Keeper.GetDisputeBySubject(ctx, msg.SubjectId, msg.TargetRole)
	if err != nil {
		return nil, err
	}
	if dispute.Status != types.DisputeStatus_dispute_open {
		return nil, errorsmod.Wrapf(types.ErrDisputeNotOpen, "current status %s", dispute.Status)
	}
	claim, err := s.Keeper.GetClaim(ctx, msg.SubjectId)
	if err != nil {
		return nil, err
	}
	collection, err := s.Keeper.GetCollection(ctx, claim.CollectionId)
	if err != nil {
		return nil, err
	}

	// 2. Verify adjudicator authorization. DID must be on the collection
	//    whitelist, and the signer must be a registered key on that DID
	//    document (same DID-key rule used by every other DID-gated message).
	//    The IID ante has already enforced the DID-key check at tx time;
	//    AuthorizeAdjudicator re-runs it as defense-in-depth and resolves
	//    the payout routing flag.
	adjudicatorEntry, ok := LookupAdjudicator(collection, msg.AdjudicatorDid)
	if !ok {
		return nil, errorsmod.Wrapf(types.ErrAdjudicatorDidNotApproved,
			"did %s not in collection.adjudicators", msg.AdjudicatorDid)
	}
	signer, err := sdk.AccAddressFromBech32(msg.AdjudicatorAddress)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid adjudicator address (%s)", err)
	}
	auth, err := s.Keeper.AuthorizeAdjudicator(ctx, msg.AdjudicatorDid, signer)
	if err != nil {
		return nil, err
	}
	adjudicatorPayout, err := s.Keeper.AdjudicatorPayoutAddress(ctx, auth)
	if err != nil {
		return nil, err
	}

	// 3. Resolve the target agent (loser on AWARDED) from current claim
	//    state. Safe to re-derive because:
	//      - SUBMITTER: claim.AgentAddress is immutable.
	//      - EVALUATOR: MsgDisputeClaim rejects filing against a FLAGGED
	//        evaluation, so by the time a dispute exists the evaluation is
	//        terminal — and terminal evaluations cannot be re-evaluated
	//        (ErrClaimDuplicateEvaluation). So claim.Evaluation.AgentAddress
	//        is stable for the lifetime of any EVALUATOR-targeted dispute.
	var targetAgent string
	switch dispute.TargetRole {
	case types.DisputeTargetRole_target_submitter:
		targetAgent = claim.AgentAddress
	case types.DisputeTargetRole_target_evaluator:
		if claim.Evaluation == nil {
			return nil, types.ErrDisputeTargetEvaluatorNoEvaluation
		}
		targetAgent = claim.Evaluation.AgentAddress
	default:
		return nil, types.ErrDisputeTargetRoleInvalid
	}
	targetAgentAddr, err := sdk.AccAddressFromBech32(targetAgent)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid target agent address (%s)", err)
	}
	disputerAddr, err := sdk.AccAddressFromBech32(dispute.DisputerAddress)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid disputer address (%s)", err)
	}
	escrowAddr, err := sdk.AccAddressFromBech32(collection.EscrowAccount)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrInternalError, "invalid escrow address: %s", err)
	}

	// 4. Determine the intended penalty.
	//
	// The penalty is only economically applied on AWARDED — it sets the
	// upper bound on what gets slashed from the loser agent's deposit
	// balance. On DISMISSED, the pot is always the disputer's locked
	// dispute_deposit and the penalty field is ignored.
	//
	// For AWARDED:
	//   - if collection has fixed penalty_amount_per_dispute, use that
	//   - else require msg.PenaltyAmount, capped at the loser's role-
	//     deposit-required (current collection config — admin can adjust
	//     in-flight; actual slash is always min(intended, balance), so
	//     no overdraft is possible regardless of how the config moves).
	var intendedPenalty sdk.Coins
	if msg.Outcome == types.DisputeStatus_dispute_awarded {
		if !collection.PenaltyAmountPerDispute.IsZero() {
			intendedPenalty = collection.PenaltyAmountPerDispute
		} else {
			if msg.PenaltyAmount.IsZero() {
				return nil, types.ErrPenaltyAmountRequired
			}
			intendedPenalty = msg.PenaltyAmount.Sort()
			roleCap := DepositRequiredForRole(collection, dispute.TargetRole)
			if !roleCap.IsZero() && !intendedPenalty.IsAllLTE(roleCap.Sort()) {
				return nil, errorsmod.Wrapf(types.ErrPenaltyAmountExceedsCap,
					"penalty %s exceeds role deposit cap %s", intendedPenalty, roleCap)
			}
		}
	}

	// 5. Apply outcome.
	resolution := &types.DisputeResolution{
		AdjudicatorDid:           msg.AdjudicatorDid,
		AdjudicatorAddress:       msg.AdjudicatorAddress,
		AdjudicatorPayoutAddress: adjudicatorPayout.String(),
		Data:                     msg.Data,
		IntendedPenalty:          intendedPenalty,
	}
	resolvedAt := ctx.BlockTime()
	resolution.ResolvedAt = &resolvedAt

	var actualPenalty sdk.Coins

	if msg.Outcome == types.DisputeStatus_dispute_awarded {
		// AWARDED: loser is the target agent; pot is min(intended, target balance).
		actualPenalty = s.Keeper.PenaltyPotForAwarded(ctx, claim.CollectionId, targetAgent, intendedPenalty)
		// Compute the 80/20 split BEFORE moving funds so we know the exact
		// amounts to send. SplitPenalty rounds adjudicator share down so
		// winner+adjudicator == pot.
		winnerAmt, adjudicatorAmt := SplitPenalty(actualPenalty, adjudicatorEntry.RewardPercentage)
		// Slash from target's balance: send winner share to disputer, adjudicator
		// share to adjudicator payout. SlashAgentDepositBalance handles the
		// debit + bank transfer + event for each.
		if !winnerAmt.IsZero() {
			if _, err := s.Keeper.SlashAgentDepositBalance(ctx, claim.CollectionId, targetAgentAddr, winnerAmt, disputerAddr); err != nil {
				return nil, err
			}
		}
		if !adjudicatorAmt.IsZero() {
			if _, err := s.Keeper.SlashAgentDepositBalance(ctx, claim.CollectionId, targetAgentAddr, adjudicatorAmt, adjudicatorPayout); err != nil {
				return nil, err
			}
		}
		// Disputer wins: dispute deposit is refunded in full (not slashed).
		if !dispute.DisputeDeposit.IsZero() {
			if err := s.Keeper.BankKeeper.SendCoins(ctx, escrowAddr, disputerAddr, dispute.DisputeDeposit); err != nil {
				return nil, err
			}
		}
		resolution.WinnerAddress = dispute.DisputerAddress
		resolution.LoserAddress = targetAgent
		resolution.WinnerAmount = winnerAmt
		resolution.AdjudicatorAmount = adjudicatorAmt
		dispute.Status = types.DisputeStatus_dispute_awarded
		collection.DisputesAwarded++
	} else {
		// DISMISSED: loser is the disputer; pot is the dispute deposit itself.
		// SplitPenalty applies the same 80/20 ratio to the dispute deposit.
		actualPenalty = dispute.DisputeDeposit
		// If actualPenalty == zero (collection had no dispute_deposit at
		// filing time), pay-outs are zero.
		winnerAmt, adjudicatorAmt := SplitPenalty(actualPenalty, adjudicatorEntry.RewardPercentage)
		if !winnerAmt.IsZero() {
			if err := s.Keeper.BankKeeper.SendCoins(ctx, escrowAddr, targetAgentAddr, winnerAmt); err != nil {
				return nil, err
			}
		}
		if !adjudicatorAmt.IsZero() {
			if err := s.Keeper.BankKeeper.SendCoins(ctx, escrowAddr, adjudicatorPayout, adjudicatorAmt); err != nil {
				return nil, err
			}
		}
		resolution.WinnerAddress = targetAgent
		resolution.LoserAddress = dispute.DisputerAddress
		resolution.WinnerAmount = winnerAmt
		resolution.AdjudicatorAmount = adjudicatorAmt
		dispute.Status = types.DisputeStatus_dispute_dismissed
		collection.DisputesDismissed++
	}

	resolution.ActualPenaltyPaid = actualPenalty

	// 6. Persist resolution back onto dispute record and clear indices.
	dispute.Resolution = resolution
	s.Keeper.SetDispute(ctx, dispute)
	// Subject index continues to point at this dispute's proof so future
	// CanFileNewDisputeForSubject reads the new status (AWARDED/DISMISSED).
	s.Keeper.SetDisputeSubjectIndex(ctx, dispute.SubjectId, dispute.TargetRole, dispute.Data.Proof)
	s.Keeper.RemoveActiveDispute(ctx, claim.CollectionId, targetAgent, dispute.SubjectId)

	if collection.DisputesOpen > 0 {
		collection.DisputesOpen--
	}
	if err := s.Keeper.CollectionPersistAndEmitEvents(ctx, collection); err != nil {
		return nil, err
	}

	if err := ctx.EventManager().EmitTypedEvents(&types.DisputeResolvedEvent{Dispute: &dispute}); err != nil {
		return nil, err
	}

	return &types.MsgAdjudicateDisputeResponse{ActualPenaltyPaid: actualPenalty}, nil
}

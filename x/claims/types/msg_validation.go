package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ixo "github.com/ixofoundation/ixo-blockchain/v6/lib/ixo"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v6/x/iid/types"
)

// --------------------------
// CREATE COLLECTION
// --------------------------
func (msg MsgCreateCollection) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	if !iidtypes.IsValidDID(msg.Entity) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.Entity)
	}

	if !iidtypes.IsValidDID(msg.Protocol) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.Protocol)
	}

	if err = msg.Payments.Validate(); err != nil {
		return err
	}

	if !ixo.IsEnumValueValid(CollectionState_name, int32(msg.State)) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid enum for state")
	}
	if !ixo.IsEnumValueValid(CollectionIntentOptions_name, int32(msg.Intents)) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid enum for intents")
	}

	// Dispute / performance-deposit config (all optional; if any is set the
	// helper enforces the cross-field invariants and DID format).
	if err := ValidateCollectionDisputeConfig(CollectionDisputeConfig{
		ServiceAgentDepositRequired: msg.ServiceAgentDepositRequired,
		EvaluatorDepositRequired:    msg.EvaluatorDepositRequired,
		DisputeDepositAmount:        msg.DisputeDepositAmount,
		Adjudicators:                msg.Adjudicators,
		PenaltyAmountPerDispute:     msg.PenaltyAmountPerDispute,
		MinDepositPeriod:            msg.MinDepositPeriod,
	}); err != nil {
		return err
	}

	return nil
}

// --------------------------
// SUBMIT CLAIM
// --------------------------
func (msg MsgSubmitClaim) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.AgentAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}

	if !iidtypes.IsValidDID(msg.AgentDid.Did()) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.AgentDid.String())
	}

	if ixo.IsEmpty(msg.ClaimId) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "claim_id cannot be empty")
	}
	if ixo.IsEmpty(msg.CollectionId) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "collection_id cannot be empty")
	}

	if err = ValidateCoinsAllowZero(msg.Amount.Sort()); err != nil {
		return err
	}
	if err = ValidateCW20Payments(msg.Cw20Payment, true); err != nil {
		return err
	}
	if err = ValidateCW1155Payments(msg.Cw1155Payment, true); err != nil {
		return err
	}

	// member_address is optional but must be valid if provided
	if msg.MemberAddress != "" {
		_, err = sdk.AccAddressFromBech32(msg.MemberAddress)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid member address (%s)", err)
		}
	}

	return nil
}

// --------------------------
// EVALUATE CLAIM
// --------------------------
func (msg MsgEvaluateClaim) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.AgentAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid agent address (%s)", err)
	}

	if !iidtypes.IsValidDID(msg.AgentDid.Did()) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.AgentDid.String())
	}
	if !iidtypes.IsValidDID(msg.Oracle) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.Oracle)
	}

	if iidtypes.IsEmpty(msg.ClaimId) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "claim_id cannot be empty")
	}
	if iidtypes.IsEmpty(msg.CollectionId) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "collection_id cannot be empty")
	}
	if iidtypes.IsEmpty(msg.VerificationProof) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "verification_proof cannot be empty")
	}

	if msg.Status == EvaluationStatus_pending {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "evaluation status can't be pending")
	}

	if !ixo.IsEnumValueValid(EvaluationStatus_name, int32(msg.Status)) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid enum for status")
	}

	// DISPUTED is deprecated as a new-tx evaluation status in v7. The dispute
	// lifecycle lives on the Dispute record (MsgDisputeClaim / MsgAdjudicateDispute);
	// having "disputed" as an evaluation outcome conflated two concepts.
	// Existing on-chain evaluations with status=DISPUTED remain valid history.
	if msg.Status == EvaluationStatus_disputed {
		return ErrEvaluationStatusDisputedDeprecated
	}

	if err = ValidateCoinsAllowZero(msg.Amount.Sort()); err != nil {
		return err
	}
	if err = ValidateCW20Payments(msg.Cw20Payment, true); err != nil {
		return err
	}
	if err = ValidateCW1155Payments(msg.Cw1155Payment, true); err != nil {
		return err
	}

	return nil
}

// --------------------------
// DISPUTE CLAIM
// --------------------------
func (msg MsgDisputeClaim) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AgentAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid agent address (%s)", err)
	}

	if !iidtypes.IsValidDID(msg.AgentDid.Did()) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.AgentDid.String())
	}
	if iidtypes.IsEmpty(msg.SubjectId) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "subject id cannot be empty")
	}
	if msg.Data == nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "dispute data cannot be nil")
	}
	if iidtypes.IsEmpty(msg.Data.Proof) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "dispute data proof cannot be empty")
	}
	if iidtypes.IsEmpty(msg.Data.Uri) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "dispute data uri cannot be empty")
	}
	if iidtypes.IsEmpty(msg.Data.Type) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "dispute data type cannot be empty")
	}
	// target_role must be SUBMITTER or EVALUATOR. UNSPECIFIED only exists on
	// legacy disputes migrated from pre-v7 state; new txs must pick a role.
	if msg.TargetRole != DisputeTargetRole_target_submitter &&
		msg.TargetRole != DisputeTargetRole_target_evaluator {
		return ErrDisputeTargetRoleInvalid
	}

	return nil
}

// --------------------------
// WITHDRAW PAYMENT
// --------------------------
func (msg MsgWithdrawPayment) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.ToAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid to address (%s)", err)
	}

	if err = ValidateCW1155Payments(msg.Cw1155Payment, true); err != nil {
		return err
	}

	if err = ValidateCW20Payments(msg.Cw20Payment, true); err != nil {
		return err
	}

	if !ixo.IsEnumValueValid(PaymentType_name, int32(msg.PaymentType)) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid enum for payment type")
	}

	if iidtypes.IsEmpty(msg.ClaimId) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "claim id cannot be empty")
	}

	return nil
}

// --------------------------
// UPDATE COLLECTION STATE
// --------------------------
func (msg MsgUpdateCollectionState) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}
	if iidtypes.IsEmpty(msg.CollectionId) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "collection_id cannot be empty")
	}
	if !ixo.IsEnumValueValid(CollectionState_name, int32(msg.State)) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid enum for state")
	}

	return nil
}

// --------------------------
// UPDATE COLLECTION DATES
// --------------------------
func (msg MsgUpdateCollectionDates) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}
	if iidtypes.IsEmpty(msg.CollectionId) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "collection_id cannot be empty")
	}

	return nil
}

// --------------------------
// UPDATE COLLECTION PAYMENTS
// --------------------------
func (msg MsgUpdateCollectionPayments) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}
	if iidtypes.IsEmpty(msg.CollectionId) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "collection_id cannot be empty")
	}

	if err = msg.Payments.Validate(); err != nil {
		return err
	}

	return nil
}

// --------------------------
// UPDATE COLLECTION INTENTS
// --------------------------
func (msg MsgUpdateCollectionIntents) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}
	if iidtypes.IsEmpty(msg.CollectionId) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "collection_id cannot be empty")
	}
	if !ixo.IsEnumValueValid(CollectionIntentOptions_name, int32(msg.Intents)) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid enum for intents")
	}

	return nil
}

func (msg MsgUpdateCollectionQuota) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.AdminAddress); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}
	if iidtypes.IsEmpty(msg.CollectionId) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "collection_id cannot be empty")
	}
	// Quota is a uint64; ValidateBasic accepts any value. The
	// quota-vs-current-count check is in the handler, since it needs the
	// stored collection state to compare against.
	return nil
}

// --------------------------
// CLAIM INTENT
// --------------------------
func (msg *MsgClaimIntent) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AgentAddress)
	if err != nil {
		return errorsmod.Wrapf(err, "invalid agent address (%s)", err)
	}
	if !iidtypes.IsValidDID(msg.AgentDid.Did()) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.AgentDid.String())
	}
	if iidtypes.IsEmpty(msg.CollectionId) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "collection_id cannot be empty")
	}

	if err = ValidateCoinsAllowZero(msg.Amount.Sort()); err != nil {
		return err
	}
	if err = ValidateCW20Payments(msg.Cw20Payment, true); err != nil {
		return err
	}
	if err = ValidateCW1155Payments(msg.Cw1155Payment, true); err != nil {
		return err
	}

	// member_address is optional but must be valid if provided
	if msg.MemberAddress != "" {
		_, err = sdk.AccAddressFromBech32(msg.MemberAddress)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid member address (%s)", err)
		}
	}

	return nil
}

// --------------------------
// CREATE CLAIM AUTHORIZATION
// --------------------------
func (msg MsgCreateClaimAuthorization) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.CreatorAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.GranteeAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid grantee address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}

	if !iidtypes.IsValidDID(msg.CreatorDid.Did()) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.CreatorDid.String())
	}

	if iidtypes.IsEmpty(msg.CollectionId) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "collection_id cannot be empty")
	}

	if !ixo.IsEnumValueValid(CreateClaimAuthorizationType_name, int32(msg.AuthType)) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid enum for auth_type")
	}

	if msg.AgentQuota == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "agent_quota cannot be 0")
	}

	if err = ValidateCoinsAllowZero(msg.MaxAmount.Sort()); err != nil {
		return err
	}
	if err = ValidateCW20Payments(msg.MaxCw20Payment, true); err != nil {
		return err
	}
	if err = ValidateCW1155Payments(msg.MaxCw1155Payment, true); err != nil {
		return err
	}

	// member_address is optional but must be valid if provided
	if msg.MemberAddress != "" {
		_, err = sdk.AccAddressFromBech32(msg.MemberAddress)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid member address (%s)", err)
		}
	}

	return nil
}

// --------------------------
// SET COLLECTION MEMBERS
// --------------------------
func (msg MsgSetCollectionMembers) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}
	if iidtypes.IsEmpty(msg.CollectionId) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "collection_id cannot be empty")
	}
	if len(msg.Members) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "members cannot be empty")
	}
	seenMembers := make(map[string]bool, len(msg.Members))
	for _, member := range msg.Members {
		_, err := sdk.AccAddressFromBech32(member.MemberAddress)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid member address (%s)", err)
		}
		if seenMembers[member.MemberAddress] {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "duplicate member address %s", member.MemberAddress)
		}
		seenMembers[member.MemberAddress] = true
		if member.Period < MinMemberBudgetPeriod {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "period must be at least %s for member %s", MinMemberBudgetPeriod, member.MemberAddress)
		}
		if member.PeriodSpendLimit.IsZero() && len(member.PeriodCw20SpendLimit) == 0 {
			return errorsmod.Wrapf(ErrMemberBudgetZero, "for member %s", member.MemberAddress)
		}
		if err = ValidateCoinsAllowZero(member.PeriodSpendLimit.Sort()); err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid period_spend_limit for member %s: %s", member.MemberAddress, err)
		}
		if err = ValidateCW20Payments(member.PeriodCw20SpendLimit, true); err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid period_cw20_spend_limit for member %s: %s", member.MemberAddress, err)
		}
	}
	return nil
}

// --------------------------
// REMOVE COLLECTION MEMBERS
// --------------------------
func (msg MsgRemoveCollectionMembers) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}
	if iidtypes.IsEmpty(msg.CollectionId) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "collection_id cannot be empty")
	}
	if len(msg.MemberAddresses) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "member_addresses cannot be empty")
	}
	seenAddresses := make(map[string]bool, len(msg.MemberAddresses))
	for _, addr := range msg.MemberAddresses {
		_, err := sdk.AccAddressFromBech32(addr)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid member address (%s)", err)
		}
		if seenAddresses[addr] {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "duplicate member address %s", addr)
		}
		seenAddresses[addr] = true
	}
	return nil
}

// --------------------------
// UPDATE COLLECTION DISPUTE CONFIG
// --------------------------
func (msg MsgUpdateCollectionDisputeConfig) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.AdminAddress); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}
	if iidtypes.IsEmpty(msg.CollectionId) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "collection_id cannot be empty")
	}
	return ValidateCollectionDisputeConfig(CollectionDisputeConfig{
		ServiceAgentDepositRequired: msg.ServiceAgentDepositRequired,
		EvaluatorDepositRequired:    msg.EvaluatorDepositRequired,
		DisputeDepositAmount:        msg.DisputeDepositAmount,
		Adjudicators:                msg.Adjudicators,
		PenaltyAmountPerDispute:     msg.PenaltyAmountPerDispute,
		MinDepositPeriod:            msg.MinDepositPeriod,
	})
}

// --------------------------
// ADD PERFORMANCE DEPOSIT
// --------------------------
func (msg MsgAddPerformanceDeposit) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.AgentAddress); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid agent address (%s)", err)
	}
	if iidtypes.IsEmpty(msg.CollectionId) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "collection_id cannot be empty")
	}
	if err := msg.Amount.Sort().Validate(); err != nil {
		return errorsmod.Wrapf(ErrAgentDepositAmountInvalid, "%s", err)
	}
	if msg.Amount.IsZero() {
		return errorsmod.Wrap(ErrAgentDepositAmountInvalid, "amount cannot be zero")
	}
	return nil
}

// --------------------------
// WITHDRAW PERFORMANCE DEPOSIT
// --------------------------
func (msg MsgWithdrawPerformanceDeposit) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.AgentAddress); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid agent address (%s)", err)
	}
	if iidtypes.IsEmpty(msg.CollectionId) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "collection_id cannot be empty")
	}
	// Empty amount means "withdraw full balance"; that's allowed. If supplied,
	// must validate as positive Coins.
	if len(msg.Amount) > 0 {
		if err := msg.Amount.Sort().Validate(); err != nil {
			return errorsmod.Wrapf(ErrAgentDepositAmountInvalid, "%s", err)
		}
		if msg.Amount.IsZero() {
			return errorsmod.Wrap(ErrAgentDepositAmountInvalid, "amount cannot be zero (omit field to withdraw all)")
		}
	}
	return nil
}

// --------------------------
// ADJUDICATE DISPUTE
// --------------------------
func (msg MsgAdjudicateDispute) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.AdjudicatorAddress); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid adjudicator address (%s)", err)
	}
	if iidtypes.IsEmpty(msg.SubjectId) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "subject_id cannot be empty")
	}
	if !iidtypes.IsValidDID(msg.AdjudicatorDid) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.AdjudicatorDid)
	}
	if msg.TargetRole != DisputeTargetRole_target_submitter &&
		msg.TargetRole != DisputeTargetRole_target_evaluator {
		return ErrDisputeTargetRoleInvalid
	}
	if msg.Outcome != DisputeStatus_dispute_awarded &&
		msg.Outcome != DisputeStatus_dispute_dismissed {
		return ErrAdjudicationInvalidOutcome
	}
	// penalty_amount is optional at ValidateBasic time; the handler decides
	// whether the field is needed (depends on collection's fixed-penalty
	// config) and applies the cap. If supplied, must validate positively.
	if len(msg.PenaltyAmount) > 0 {
		if err := msg.PenaltyAmount.Sort().Validate(); err != nil {
			return errorsmod.Wrapf(ErrPenaltyAmountInvalid, "%s", err)
		}
		if msg.PenaltyAmount.IsZero() {
			return errorsmod.Wrap(ErrPenaltyAmountInvalid, "penalty_amount cannot be zero (omit to use collection fixed)")
		}
	}
	// data is the adjudicator's structured opinion (uri + proof + type +
	// encrypted flag), symmetric with MsgDisputeClaim.data. The field itself
	// is optional — an adjudicator can resolve without attaching evidence —
	// but if supplied, the three string fields must each be non-empty so
	// downstream indexers never see half-populated records.
	if msg.Data != nil {
		if iidtypes.IsEmpty(msg.Data.Proof) {
			return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "dispute resolution data proof cannot be empty")
		}
		if iidtypes.IsEmpty(msg.Data.Uri) {
			return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "dispute resolution data uri cannot be empty")
		}
		if iidtypes.IsEmpty(msg.Data.Type) {
			return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "dispute resolution data type cannot be empty")
		}
	}
	return nil
}

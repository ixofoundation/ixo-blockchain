package types

import (
	errorsmod "cosmossdk.io/errors"
)

// x/claims module sentinel errors
var (
	ErrCollectionNotFound            = errorsmod.Register(ModuleName, 1001, "collection not found")
	ErrCollectionNotOpen             = errorsmod.Register(ModuleName, 1002, "collection is not in open state")
	ErrCollectionEvalError           = errorsmod.Register(ModuleName, 1003, "evaluation payment is not allowed to have a Contract1155Payment")
	ErrCollNotEntityAcc              = errorsmod.Register(ModuleName, 1004, "collection payments accounts can only be entity accounts")
	ErrCollectionEvalCW20Error       = errorsmod.Register(ModuleName, 1005, "evaluation payment is not allowed to have CW20 payments")
	ErrCollectionEvalCW1155Error     = errorsmod.Register(ModuleName, 1006, "evaluation payment is not allowed to have CW1155 payments")
	ErrCollectionApprovalCW1155Error = errorsmod.Register(ModuleName, 1007, "approval payment that is an oracle payment is not allowed to have CW1155 payments")

	ErrClaimNotFound               = errorsmod.Register(ModuleName, 1100, "claim not found")
	ErrClaimUnauthorized           = errorsmod.Register(ModuleName, 1101, "unauthorized, incorrect admin")
	ErrClaimCollectionNotStarted   = errorsmod.Register(ModuleName, 1102, "collection for claim has not started yet")
	ErrClaimCollectionEnded        = errorsmod.Register(ModuleName, 1103, "collection for claim has ended")
	ErrClaimCollectionQuotaReached = errorsmod.Register(ModuleName, 1104, "collection for claim's quato has been reached")
	ErrClaimDuplicate              = errorsmod.Register(ModuleName, 1105, "claim with id already exists")
	ErrClaimDuplicateEvaluation    = errorsmod.Register(ModuleName, 1106, "claim with id already evaluated")

	ErrDisputeNotFound     = errorsmod.Register(ModuleName, 1200, "dispute not found")
	ErrDisputeDuplicate    = errorsmod.Register(ModuleName, 1201, "dispute with proof already exists")
	ErrDisputeUnauthorized = errorsmod.Register(ModuleName, 1202, "unauthorized, not part of collection/entity/authz agent")

	ErrEvaluateWrongCollection = errorsmod.Register(ModuleName, 1300, "evaluation claim and collection does not match")
	ErrSelfReFlag              = errorsmod.Register(ModuleName, 1301, "the same agent cannot re-flag a claim they have already flagged")
	ErrInvalidFlagTransition   = errorsmod.Register(ModuleName, 1303, "claim with terminal evaluation cannot be re-evaluated")

	ErrPaymentPresetPercentagesOverflow = errorsmod.Register(ModuleName, 1400, "preset fee percentages for node and network overflows 100%")
	ErrPaymentWithdrawFailed            = errorsmod.Register(ModuleName, 1401, "payment withdrawal failed")
	ErrDistributionFailed               = errorsmod.Register(ModuleName, 1402, "distribution calculations failed")
	ErrOraclePaymentOnlyNative          = errorsmod.Register(ModuleName, 1403, "oracle payments can only have Native Coin payments if no intent is used, CW20 payments are not allowed")

	ErrIntentNotFound     = errorsmod.Register(ModuleName, 1500, "intent not found")
	ErrIntentExists       = errorsmod.Register(ModuleName, 1501, "active intent found")
	ErrIntentUnauthorized = errorsmod.Register(ModuleName, 1502, "unauthorized")

	ErrInvalidResponse = errorsmod.Register(ModuleName, 1601, "invalid response")

	ErrInternalError = errorsmod.Register(ModuleName, 1700, "internal error")

	ErrMemberBudgetNotFound    = errorsmod.Register(ModuleName, 1800, "member budget not found")
	ErrMemberBudgetExceeded    = errorsmod.Register(ModuleName, 1801, "member budget exceeded")
	ErrMemberBudgetZero        = errorsmod.Register(ModuleName, 1802, "member budget spend limits cannot all be zero, use remove instead")
	ErrMemberAddressRequired   = errorsmod.Register(ModuleName, 1803, "member_address is required for collections with member budgets")
	ErrMemberAddressMismatch   = errorsmod.Register(ModuleName, 1804, "member_address does not match")
	ErrMemberAddressNotAllowed = errorsmod.Register(ModuleName, 1805, "member_address is not allowed for collections without member budgets")

	// Performance-deposit / dispute errors (v7).
	ErrAgentDepositInsufficient            = errorsmod.Register(ModuleName, 1900, "agent performance-deposit balance is below required amount")
	ErrAgentDepositNotFound                = errorsmod.Register(ModuleName, 1901, "agent performance-deposit balance not found")
	ErrAgentDepositBalanceCannotWithdraw   = errorsmod.Register(ModuleName, 1902, "performance-deposit cannot be withdrawn while agent has active disputes")
	ErrAgentDepositAmountExceedsBalance    = errorsmod.Register(ModuleName, 1903, "withdrawal amount exceeds current balance")
	ErrAgentDepositAmountInvalid           = errorsmod.Register(ModuleName, 1904, "performance-deposit amount must be positive and in valid denoms")
	ErrAgentHasActiveDispute               = errorsmod.Register(ModuleName, 1905, "agent has an active dispute on this collection; action blocked until adjudicated")
	ErrDisputeTargetRoleInvalid            = errorsmod.Register(ModuleName, 1906, "dispute target_role must be SUBMITTER or EVALUATOR")
	ErrDisputeTargetEvaluatorNoEvaluation  = errorsmod.Register(ModuleName, 1907, "cannot dispute evaluator role: claim has not been evaluated")
	ErrDisputeTargetEvaluatorFlagged       = errorsmod.Register(ModuleName, 1923, "cannot dispute a FLAGGED evaluation: FLAGGED is non-terminal; resolve by re-evaluation instead")
	ErrAgentDepositLocked                  = errorsmod.Register(ModuleName, 1924, "performance-deposit is still within the minimum deposit period; withdrawal is locked")
	ErrDisputeAlreadyOpenForSubjectRole    = errorsmod.Register(ModuleName, 1908, "an OPEN dispute already exists for this (subject_id, target_role)")
	ErrDisputeAlreadyAwardedForSubjectRole = errorsmod.Register(ModuleName, 1909, "this (subject_id, target_role) was previously AWARDED; further disputes are blocked")
	ErrDisputeNotFoundForSubjectRole       = errorsmod.Register(ModuleName, 1910, "no dispute found for this (subject_id, target_role)")
	ErrDisputeNotOpen                      = errorsmod.Register(ModuleName, 1911, "dispute is not OPEN; cannot adjudicate")
	ErrAdjudicatorDidNotApproved           = errorsmod.Register(ModuleName, 1912, "adjudicator_did is not in the collection's adjudicators whitelist")
	ErrAdjudicatorNotAuthorized            = errorsmod.Register(ModuleName, 1913, "adjudicator_address is not an entity account of the adjudicator_did and not a key registered on the adjudicator_did DID document")
	ErrAdjudicationInvalidOutcome          = errorsmod.Register(ModuleName, 1914, "adjudication outcome must be AWARDED or DISMISSED")
	ErrPenaltyAmountRequired               = errorsmod.Register(ModuleName, 1915, "penalty_amount must be supplied because collection has no fixed penalty_amount_per_dispute")
	ErrPenaltyAmountExceedsCap             = errorsmod.Register(ModuleName, 1916, "penalty_amount exceeds the loser's deposit-required cap")
	ErrPenaltyAmountInvalid                = errorsmod.Register(ModuleName, 1917, "penalty_amount must be positive coins")
	ErrDisputeDepositNotConfigured         = errorsmod.Register(ModuleName, 1918, "collection has no dispute_deposit_amount; disputes are not enabled here")
	ErrAdjudicationNotConfigured           = errorsmod.Register(ModuleName, 1919, "collection has no adjudicators configured; disputes cannot be adjudicated")
	ErrDisputeConfigInvalid                = errorsmod.Register(ModuleName, 1920, "invalid dispute / performance-deposit configuration")
	ErrEvaluationStatusDisputedDeprecated  = errorsmod.Register(ModuleName, 1921, "EvaluationStatus DISPUTED is deprecated; use MsgDisputeClaim instead")
	ErrCollectionQuotaBelowCount           = errorsmod.Register(ModuleName, 1925, "new quota is below the collection's current claim count")
)

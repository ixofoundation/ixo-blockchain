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
)

package types

import (
	errorsmod "cosmossdk.io/errors"
)

// x/claims module sentinel errors
var (
	ErrCollectionNotFound      = errorsmod.Register(ModuleName, 1001, "collection not found")
	ErrCollectionNotOpen       = errorsmod.Register(ModuleName, 1002, "collection is not in open state")
	ErrCollectionEvalError     = errorsmod.Register(ModuleName, 1003, "evaluation payment is not allowed to have a Contract1155Payment")
	ErrCollNotEntityAcc        = errorsmod.Register(ModuleName, 1004, "collection payments accounts can only be entity accounts")
	ErrCollectionEvalCW20Error = errorsmod.Register(ModuleName, 1005, "evaluation payment is not allowed to have CW20 payments")

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

	ErrPaymentPresetPercentagesOverflow = errorsmod.Register(ModuleName, 1400, "preset fee percentages for node and network overflows 100%")
	ErrPaymentWithdrawFailed            = errorsmod.Register(ModuleName, 1401, "payment withdrawal failed")
	ErrDistributionFailed               = errorsmod.Register(ModuleName, 1402, "distribution calculations failed")
	ErrOraclePaymentOnlyNative          = errorsmod.Register(ModuleName, 1403, "oracle payments can only have Native Coin payments if no intent is used, CW20 payments are not allowed")

	ErrIntentNotFound     = errorsmod.Register(ModuleName, 1500, "intent not found")
	ErrIntentExists       = errorsmod.Register(ModuleName, 1501, "active intent found")
	ErrIntentUnauthorized = errorsmod.Register(ModuleName, 1502, "unauthorized")
)

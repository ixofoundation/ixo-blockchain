package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/claims module sentinel errors
var (
	ErrCollectionNotFound = sdkerrors.Register(ModuleName, 1001, "collection not found")
	ErrCollectionNotOpen  = sdkerrors.Register(ModuleName, 1002, "collection is not in open state")

	ErrClaimNotFound               = sdkerrors.Register(ModuleName, 1100, "claim not found")
	ErrClaimUnauthorized           = sdkerrors.Register(ModuleName, 1101, "unauthorized, incorrect admin")
	ErrClaimCollectionNotStarted   = sdkerrors.Register(ModuleName, 1102, "collection for claim has not started yet")
	ErrClaimCollectionEnded        = sdkerrors.Register(ModuleName, 1103, "collection for claim has ended")
	ErrClaimCollectionQuotaReached = sdkerrors.Register(ModuleName, 1104, "collection for claim's quato has been reached")
	ErrClaimDuplicate              = sdkerrors.Register(ModuleName, 1105, "claim with id already exists")
	ErrClaimDuplicateEvaluation    = sdkerrors.Register(ModuleName, 1106, "claim with id already evaluated")

	ErrDisputeNotFound     = sdkerrors.Register(ModuleName, 1200, "dispute not found")
	ErrDisputeDuplicate    = sdkerrors.Register(ModuleName, 1201, "dispute with proof already exists")
	ErrDisputeUnauthorized = sdkerrors.Register(ModuleName, 1200, "unauthorized, not part of collection/entity/authz agent")

	ErrEvaluateWrongCollection = sdkerrors.Register(ModuleName, 1300, "evaluation claim and collection does not match")

	ErrPaymentPresetPercentagesOverflow = sdkerrors.Register(ModuleName, 1400, "preset fee percentages for node and network overflows 100%")
	ErrPaymentWithdrawFailed            = sdkerrors.Register(ModuleName, 1401, "payment withdrawal failed")
)

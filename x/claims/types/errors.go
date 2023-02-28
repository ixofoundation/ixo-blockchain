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
	ErrClaimDuplicateEvaluation    = sdkerrors.Register(ModuleName, 1105, "claim with id already evaluated")

	ErrDisputeNotFound  = sdkerrors.Register(ModuleName, 1200, "dispute not found")
	ErrDisputeDuplicate = sdkerrors.Register(ModuleName, 1201, "dispute with proof already exists")
)

package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/entity module sentinel errors
var (
	ErrEntityNotFound        = sdkerrors.Register(ModuleName, 1001, "entity not found")
	ErrUpdateVerifiedFailed  = sdkerrors.Register(ModuleName, 1002, "update entity verified failed")
)

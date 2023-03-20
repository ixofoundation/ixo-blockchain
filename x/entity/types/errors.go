package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/entity module sentinel errors
var (
	ErrEntityNotFound       = sdkerrors.Register(ModuleName, 1001, "entity not found")
	ErrUpdateVerifiedFailed = sdkerrors.Register(ModuleName, 1002, "update entity verified failed")
	ErrAccountNotFound      = sdkerrors.Register(ModuleName, 1003, "entity account with name not found")
	ErrAccountDuplicate     = sdkerrors.Register(ModuleName, 1004, "entity account with name already exists")
	ErrEntityUnauthorized   = sdkerrors.Register(ModuleName, 1005, "unauthorized, owner not same as nft owner")
)

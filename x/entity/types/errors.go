package types

import (
	errorsmod "cosmossdk.io/errors"
)

// x/entity module sentinel errors
var (
	ErrEntityNotFound       = errorsmod.Register(ModuleName, 1001, "entity not found")
	ErrUpdateVerifiedFailed = errorsmod.Register(ModuleName, 1002, "update entity verified failed")
	ErrAccountNotFound      = errorsmod.Register(ModuleName, 1003, "entity account with name not found")
	ErrAccountDuplicate     = errorsmod.Register(ModuleName, 1004, "entity account with name already exists")
	ErrEntityUnauthorized   = errorsmod.Register(ModuleName, 1005, "unauthorized, owner not same as nft owner")
)

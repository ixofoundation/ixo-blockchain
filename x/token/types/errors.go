package types

import (
	errorsmod "cosmossdk.io/errors"
)

// x/token module sentinel errors
var (
	ErrTokenNotFound           = errorsmod.Register(ModuleName, 1001, "token not found")
	ErrTokenPropertiesNotFound = errorsmod.Register(ModuleName, 1002, "token properties not found")
	ErrTokenNameDuplicate      = errorsmod.Register(ModuleName, 1003, "token name is already taken")
	ErrTokenNameIncorrect      = errorsmod.Register(ModuleName, 1004, "token name is incorrect")
	ErrTokenAmountIncorrect    = errorsmod.Register(ModuleName, 1005, "token amount is incorrect")
	ErrTokenPaused             = errorsmod.Register(ModuleName, 1006, "token is paused")
	ErrTokenStopped            = errorsmod.Register(ModuleName, 1007, "token is stopped")
)

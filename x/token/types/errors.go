package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/token module sentinel errors
var (
	ErrTokenNotFound             = sdkerrors.Register(ModuleName, 1001, "token not found")
	ErrTokenPropertiesNotFound   = sdkerrors.Register(ModuleName, 1002, "token properties not found")
	ErrTokenNameDuplicate        = sdkerrors.Register(ModuleName, 1003, "token name is already taken")
	ErrTokenNameIncorrect        = sdkerrors.Register(ModuleName, 1004, "token name is incorrect")
	ErrTokenAmountIncorrect      = sdkerrors.Register(ModuleName, 1005, "token amount is incorrect")
	ErrTokenPausedIncorrect      = sdkerrors.Register(ModuleName, 1006, "token is paused")
	ErrTokenDeactivatedIncorrect = sdkerrors.Register(ModuleName, 1007, "token is deactivated")
)

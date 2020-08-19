package types

import sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	DefaultCodespace  = ModuleName
	ErrInternal       = sdkErrors.Register(DefaultCodespace, 101, "not allowed format")
	ErrInvalidAddress = sdkErrors.Register(DefaultCodespace, 102, "invalid project did address")
)

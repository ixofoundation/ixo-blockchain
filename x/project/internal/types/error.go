package types

import sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	DefaultCodespace  = ModuleName
	ErrInternal       = sdkErrors.Register(DefaultCodespace, 2, "not allowed format")
	ErrInvalidAddress = sdkErrors.Register(DefaultCodespace, 3, "invalid project did address")
)

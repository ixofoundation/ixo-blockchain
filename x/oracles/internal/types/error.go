package types

import sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	DefaultCodespace = ModuleName
	ErrInternal      = sdkErrors.Register(DefaultCodespace, 2, "internal error")
)

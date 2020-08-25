package types

import sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	DefaultCodespace = "ixo"
	ErrInternal      = sdkErrors.Register(DefaultCodespace, 2, "not allowed format")
)

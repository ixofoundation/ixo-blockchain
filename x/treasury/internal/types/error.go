package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	DefaultCodespace = ModuleName
	ErrInternal      = sdkerrors.Register(DefaultCodespace, 2, "not allowed format")
)

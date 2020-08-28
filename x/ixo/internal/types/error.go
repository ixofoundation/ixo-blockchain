package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

const (
	DefaultCodespace = ModuleName
)

var (
	ErrInternal = sdkerrors.Register(DefaultCodespace, 2, "internal error")
)

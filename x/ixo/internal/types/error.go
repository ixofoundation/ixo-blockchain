package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	DefaultCodespace = "ixo"
	ErrInternal      = sdkerrors.Register(DefaultCodespace, 2, "internal error")
)

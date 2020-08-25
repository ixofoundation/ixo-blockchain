package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	DefaultCodeSpace        = ModuleName
	ErrorInvalidDid         = sdkerrors.Register(DefaultCodeSpace, 2, "invalid did")
	ErrorInvalidPubKey      = sdkerrors.Register(DefaultCodeSpace, 3, "invalid pubKey")
	ErrorDidPubKeyMismatch  = sdkerrors.Register(DefaultCodeSpace, 4, "invalid did")
	ErrorInvalidIssuer      = sdkerrors.Register(DefaultCodeSpace, 5, "invalid issuer")
	ErrorInvalidCredentials = sdkerrors.Register(DefaultCodeSpace, 6, "invalid credentials")
	ErrInternal             = sdkerrors.Register(DefaultCodeSpace, 7, "internal error")
	ErrUnauthorized         = sdkerrors.Register(DefaultCodeSpace, 8, "unauthorized")
)

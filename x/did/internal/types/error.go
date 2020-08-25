package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	DefaultCodespace = ModuleName

	ErrorInvalidDid         = sdkerrors.Register(DefaultCodespace, 2, "invalid did")
	ErrorInvalidPubKey      = sdkerrors.Register(DefaultCodespace, 3, "invalid pubKey")
	ErrorDidPubKeyMismatch  = sdkerrors.Register(DefaultCodespace, 4, "invalid did")
	ErrorInvalidIssuer      = sdkerrors.Register(DefaultCodespace, 5, "invalid issuer")
	ErrorInvalidCredentials = sdkerrors.Register(DefaultCodespace, 6, "invalid credentials")
	ErrInternal             = sdkerrors.Register(DefaultCodespace, 7, "internal error")
	ErrUnauthorized         = sdkerrors.Register(DefaultCodespace, 8, "unauthorized")
)

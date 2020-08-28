package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	DefaultCodespace = ModuleName
)

var (
	ErrInvalidDid         = sdkerrors.Register(DefaultCodespace, 2, "invalid did")
	ErrInvalidPubKey      = sdkerrors.Register(DefaultCodespace, 3, "invalid pubKey")
	ErrDidPubKeyMismatch  = sdkerrors.Register(DefaultCodespace, 4, "did pubKey mismatch")
	ErrInvalidIssuer      = sdkerrors.Register(DefaultCodespace, 5, "invalid issuer")
	ErrInvalidCredentials = sdkerrors.Register(DefaultCodespace, 6, "invalid credentials")
	ErrInvalidClaimId     = sdkerrors.Register(DefaultCodespace, 7, "invalid claim ID")
	ErrInternal           = sdkerrors.Register(DefaultCodespace, 8, "internal error")
)

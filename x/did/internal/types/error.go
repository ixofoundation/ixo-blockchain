package types

import (
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	DefaultCodeSpace        = ModuleName
	CodeInvalidDid          = 201
	CodeInvalidPubKey       = 202
	CodeInvalidIssuer       = 203
	CodeInvalidCredentials  = 204
	CodeInvalidData         = 205
	CodeUnauthorized        = 206
	ErrorInvalidDid         = sdkErrors.Register(DefaultCodeSpace, 201, "code invalid did")
	ErrorInvalidPubKey      = sdkErrors.Register(DefaultCodeSpace, 202, "code invalid pubKey")
	ErrorDidPubKeyMismatch  = sdkErrors.Register(DefaultCodeSpace, 201, "code invalid did")
	ErrorInvalidIssuer      = sdkErrors.Register(DefaultCodeSpace, 203, "code invalid issuer")
	ErrorInvalidCredentials = sdkErrors.Register(DefaultCodeSpace, 204, "code invalid credentials")
	ErrInternal             = sdkErrors.Register(DefaultCodeSpace, 205, "invalid data")
	ErrUnauthorized         = sdkErrors.Register(DefaultCodeSpace, 206, "unauthorized")
)

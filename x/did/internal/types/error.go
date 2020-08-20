package types

import (
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	DefaultCodeSpace        = ModuleName
	ErrorInvalidDid         = sdkErrors.Register(DefaultCodeSpace, 208, " invalid did")
	ErrorInvalidPubKey      = sdkErrors.Register(DefaultCodeSpace, 202, " invalid pubKey")
	ErrorDidPubKeyMismatch  = sdkErrors.Register(DefaultCodeSpace, 201, " invalid did")
	ErrorInvalidIssuer      = sdkErrors.Register(DefaultCodeSpace, 203, " invalid issuer")
	ErrorInvalidCredentials = sdkErrors.Register(DefaultCodeSpace, 204, " invalid credentials")
	ErrInternal             = sdkErrors.Register(DefaultCodeSpace, 205, "invalid data")
	ErrUnauthorized         = sdkErrors.Register(DefaultCodeSpace, 206, "unauthorized")
)

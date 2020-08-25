package types

import (
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	DefaultCodeSpace        = ModuleName
	ErrorInvalidDid         = sdkErrors.Register(DefaultCodeSpace, 2, "invalid did")
	ErrorInvalidPubKey      = sdkErrors.Register(DefaultCodeSpace, 3, "invalid pubKey")
	ErrorDidPubKeyMismatch  = sdkErrors.Register(DefaultCodeSpace, 4, "invalid did")
	ErrorInvalidIssuer      = sdkErrors.Register(DefaultCodeSpace, 5, "invalid issuer")
	ErrorInvalidCredentials = sdkErrors.Register(DefaultCodeSpace, 6, "invalid credentials")
	ErrInternal             = sdkErrors.Register(DefaultCodeSpace, 7, "invalid data")
	ErrUnauthorized         = sdkErrors.Register(DefaultCodeSpace, 8, "unauthorized")
)

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultCodespace       sdk.CodespaceType = ModuleName
	CodeInvalidDid                           = 201
	CodeInvalidPubKey                        = 202
	CodeInvalidIssuer                        = 203
	CodeInvalidCredentials                   = 204
)

func ErrorInvalidDid(codeSpace sdk.CodespaceType, msg string) sdk.Error {
	return sdk.NewError(codeSpace, CodeInvalidDid, msg)
}

func ErrorInvalidPubKey(codeSpace sdk.CodespaceType, msg string) sdk.Error {
	return sdk.NewError(codeSpace, CodeInvalidPubKey, msg)
}

func ErrorDidPubKeyMismatch(codeSpace sdk.CodespaceType, msg string) sdk.Error {
	return sdk.NewError(codeSpace, CodeInvalidDid, msg)
}

func ErrorInvalidIssuer(codeSpace sdk.CodespaceType, msg string) sdk.Error {
	return sdk.NewError(codeSpace, CodeInvalidIssuer, msg)
}

func ErrorInvalidCredentials(codeSpace sdk.CodespaceType, msg string) sdk.Error {
	return sdk.NewError(codeSpace, CodeInvalidCredentials, msg)
}

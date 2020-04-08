package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultCodeSpace       sdk.CodespaceType = ModuleName
	CodeInvalidDid                           = 201
	CodeInvalidPubKey                        = 202
	CodeInvalidIssuer                        = 203
	CodeInvalidCredentials                   = 204
)

func ErrorInvalidDid(codeSpace sdk.CodespaceType, msg string) sdk.Error {
	if msg != "" {
		return sdk.NewError(codeSpace, CodeInvalidDid, msg)
	}

	return sdk.NewError(codeSpace, CodeInvalidDid, "Invalid did")
}

func ErrorInvalidPubKey(codeSpace sdk.CodespaceType, msg string) sdk.Error {
	if msg != "" {
		return sdk.NewError(codeSpace, CodeInvalidPubKey, msg)
	}

	return sdk.NewError(codeSpace, CodeInvalidPubKey, "Invalid pubKey")
}

func ErrorInvalidIssuer(codeSpace sdk.CodespaceType, msg string) sdk.Error {
	if msg != "" {
		return sdk.NewError(codeSpace, CodeInvalidIssuer, msg)
	}

	return sdk.NewError(codeSpace, CodeInvalidIssuer, "Invalid issuer")
}

func ErrorInvalidCredentials(codeSpace sdk.CodespaceType, msg string) sdk.Error {
	if msg != "" {
		return sdk.NewError(codeSpace, CodeInvalidCredentials, msg)
	}

	return sdk.NewError(codeSpace, CodeInvalidCredentials, "Data already exist")
}

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type CodeType sdk.CodeType

const (
	DefaultCodeSpace         sdk.CodespaceType = ModuleName
	CodeInvalidAmount        sdk.CodeType      = 601
	CodeInvalidString        sdk.CodeType      = 602
	CodeInvalidInputsOutputs sdk.CodeType      = 603
	CodeInvalidPegHash       sdk.CodeType      = 604
	CodeNegativeAmount       sdk.CodeType      = 605
	CodePegHashHex           sdk.CodeType      = 606
	CodeInvalidQuery         sdk.CodeType      = 607
)

func ErrInvalidAmount(codespace sdk.CodespaceType, msg string) sdk.Error {
	if msg != "" {
		return sdk.NewError(codespace, CodeInvalidAmount, msg)
	}
	return sdk.NewError(codespace, CodeInvalidAmount, "invalid Amount")
}

func ErrInvalidString(codespace sdk.CodespaceType, msg string) sdk.Error {
	if msg != "" {
		return sdk.NewError(codespace, CodeInvalidString, msg)
	}
	return sdk.NewError(codespace, CodeInvalidString, "Invalid string")
}

func ErrNoInputs(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidInputsOutputs, "no inputs to send transaction")
}

func ErrInvalidPegHash(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidPegHash, "invalid peg hash")
}

func ErrNoOutputs(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidInputsOutputs, "no outputs to send transaction")
}

func ErrNegativeAmount(codeSpace sdk.CodespaceType, msg string) sdk.Error {
	if msg != "" {
		return sdk.NewError(codeSpace, CodeNegativeAmount, msg)
	}
	return sdk.NewError(codeSpace, CodeNegativeAmount, "Amount should not be zero")
}

func ErrPegHashHex(codeSpace sdk.CodespaceType, _type string) sdk.Error {
	return sdk.NewError(codeSpace, CodePegHashHex, "Error converting "+_type+"to hex")
}

func ErrQuery(codeSpace sdk.CodespaceType, query string) sdk.Error {
	return sdk.NewError(codeSpace, CodeInvalidQuery, "Error occurred while querying "+query+" data")
}

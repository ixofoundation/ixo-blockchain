package project

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// Project module reserves error 400-499 lawl
	CodeIncorrectProjectAnswer sdk.CodeType = 400
)

// ErrIncorrectProjectAnswer - Error returned upon an incorrect guess
func ErrIncorrectProjectAnswer(answer string) sdk.Error {
	return sdk.NewError(CodeIncorrectProjectAnswer, fmt.Sprintf("Incorrect cool answer: %v", answer))
}

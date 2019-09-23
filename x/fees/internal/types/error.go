package types

import (
	"github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultCodeSpace    types.CodespaceType = ModuleName
	CodeInvalidFeeQuery                     = 301
	CodeUnmarshal                           = 302
)

func ErrorInvalidFeeQuery() types.Error {
	return types.NewError(DefaultCodeSpace, CodeInvalidFeeQuery,
		"Error occurred while querying the fees data")
}

func ErrorUnmarshalFees() types.Error {
	return types.NewError(DefaultCodeSpace, CodeUnmarshal,
		"Error occurred while unmarshal the fees data")
}

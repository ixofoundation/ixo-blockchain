package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultCodeSpace     sdk.CodespaceType = ModuleName
	CodeInvalidQueryNode                   = 501
)

func ErrorInvalidQueryNode() sdk.Error {
	return sdk.NewError(DefaultCodeSpace, CodeInvalidQueryNode,
		"Error occurred while querying node data")
}

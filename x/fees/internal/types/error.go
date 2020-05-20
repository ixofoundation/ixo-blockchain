package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultCodespace        sdk.CodespaceType = ModuleName
	CodeInvalidDistribution sdk.CodeType      = 101
	CodeInvalidShare        sdk.CodeType      = 102
	CodeInvalidGenesis      sdk.CodeType      = 103
)

func ErrNegativeSharePercentage(codespace sdk.CodespaceType) sdk.Error {
	errMsg := fmt.Sprintf("fee distribution share percentage must be positive")
	return sdk.NewError(codespace, CodeInvalidShare, errMsg)
}

func ErrDistributionPercentagesNot100(codespace sdk.CodespaceType, total sdk.Dec) sdk.Error {
	errMsg := fmt.Sprintf("fee distribution percentages should add up to 100, not %s", total.String())
	return sdk.NewError(codespace, CodeInvalidDistribution, errMsg)
}

func ErrInvalidGenesis(codespace sdk.CodespaceType, errMsg string) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidGenesis, errMsg)
}

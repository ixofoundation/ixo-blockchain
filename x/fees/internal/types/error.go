package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultCodespace               sdk.CodespaceType = ModuleName
	CodeInvalidDistribution        sdk.CodeType      = 101
	CodeInvalidShare               sdk.CodeType      = 102
	CodeInvalidGenesis             sdk.CodeType      = 103
	CodeInvalidSubscriptionAction  sdk.CodeType      = 104
	CodeInvalidSubscriptionContent sdk.CodeType      = 105
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

func ErrSubscriptionHasNoNextPeriod(codespace sdk.CodespaceType) sdk.Error {
	errMsg := fmt.Sprintf("subscription has no next period")
	return sdk.NewError(codespace, CodeInvalidSubscriptionAction, errMsg)
}

func ErrInvalidSubscriptionContent(codespace sdk.CodespaceType, errMsg string) sdk.Error {
	errMsg = fmt.Sprintf("subscription content is invalid: %s", errMsg)
	return sdk.NewError(codespace, CodeInvalidSubscriptionContent, errMsg)
}

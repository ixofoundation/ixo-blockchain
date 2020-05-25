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
	CodeInvalidSubscriptionContent sdk.CodeType      = 104
	CodeInvalidFeeContractAction   sdk.CodeType      = 105
	CodeInvalidDiscount            sdk.CodeType      = 106
	CodeInvalidFee                 sdk.CodeType      = 107
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

func ErrInvalidSubscriptionContent(codespace sdk.CodespaceType, errMsg string) sdk.Error {
	errMsg = fmt.Sprintf("subscription content is invalid: %s", errMsg)
	return sdk.NewError(codespace, CodeInvalidSubscriptionContent, errMsg)
}

func ErrFeeContractCannotBeDeauthorised(codespace sdk.CodespaceType) sdk.Error {
	errMsg := fmt.Sprintf("fee contract cannot be deauthorised")
	return sdk.NewError(codespace, CodeInvalidFeeContractAction, errMsg)
}

func ErrDiscountIDsBeSequentialFrom1(codespace sdk.CodespaceType) sdk.Error {
	errMsg := fmt.Sprintf("discount IDs must be sequential starting with 1")
	return sdk.NewError(codespace, CodeInvalidDiscount, errMsg)
}

func ErrNegativeDiscountPercantage(codespace sdk.CodespaceType) sdk.Error {
	errMsg := fmt.Sprintf("discount percentage must be positive")
	return sdk.NewError(codespace, CodeInvalidDiscount, errMsg)
}

func ErrInvalidFee(codespace sdk.CodespaceType, errMsg string) sdk.Error {
	errMsg = fmt.Sprintf("fee invalid; %s", errMsg)
	return sdk.NewError(codespace, CodeInvalidFee, errMsg)
}

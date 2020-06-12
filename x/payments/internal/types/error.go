package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultCodespace                 sdk.CodespaceType = ModuleName
	CodeInvalidDistribution          sdk.CodeType      = 101
	CodeInvalidShare                 sdk.CodeType      = 102
	CodeInvalidPeriod                sdk.CodeType      = 103
	CodeInvalidPaymentContractAction sdk.CodeType      = 104
	CodeInvalidDiscount              sdk.CodeType      = 105
	CodeInvalidDiscountRequest       sdk.CodeType      = 106
	CodeInvalidPaymentTemplate       sdk.CodeType      = 107
	CodeInvalidSubscriptionAction    sdk.CodeType      = 108
	CodeInvalidId                    sdk.CodeType      = 109
	CodeInvalidArgument              sdk.CodeType      = 110
	CodeAlreadyExists                sdk.CodeType      = 111
)

func ErrNegativeSharePercentage(codespace sdk.CodespaceType) sdk.Error {
	errMsg := fmt.Sprintf("payment distribution share percentage must be positive")
	return sdk.NewError(codespace, CodeInvalidShare, errMsg)
}

func ErrDistributionPercentagesNot100(codespace sdk.CodespaceType, total sdk.Dec) sdk.Error {
	errMsg := fmt.Sprintf("payment distribution percentages should add up to 100, not %s", total.String())
	return sdk.NewError(codespace, CodeInvalidDistribution, errMsg)
}

func ErrInvalidPeriod(codespace sdk.CodespaceType, errMsg string) sdk.Error {
	errMsg = fmt.Sprintf("period is invalid: %s", errMsg)
	return sdk.NewError(codespace, CodeInvalidPeriod, errMsg)
}

func ErrPaymentContractCannotBeDeauthorised(codespace sdk.CodespaceType) sdk.Error {
	errMsg := fmt.Sprintf("payment contract cannot be deauthorised")
	return sdk.NewError(codespace, CodeInvalidPaymentContractAction, errMsg)
}

func ErrDiscountIDsBeSequentialFrom1(codespace sdk.CodespaceType) sdk.Error {
	errMsg := fmt.Sprintf("discount IDs must be sequential starting with 1")
	return sdk.NewError(codespace, CodeInvalidDiscount, errMsg)
}

func ErrNegativeDiscountPercentage(codespace sdk.CodespaceType) sdk.Error {
	errMsg := fmt.Sprintf("discount percentage must be positive")
	return sdk.NewError(codespace, CodeInvalidDiscount, errMsg)
}

func ErrDiscountPercentageGreaterThan100(codespace sdk.CodespaceType) sdk.Error {
	errMsg := fmt.Sprintf("discount percentage cannot exceed 100%%")
	return sdk.NewError(codespace, CodeInvalidDiscount, errMsg)
}

func ErrDiscountIdIsNotInTemplate(codespace sdk.CodespaceType) sdk.Error {
	errMsg := fmt.Sprintf("discount ID specified is not one of the template's discounts")
	return sdk.NewError(codespace, CodeInvalidDiscountRequest, errMsg)
}

func ErrInvalidPaymentTemplate(codespace sdk.CodespaceType, errMsg string) sdk.Error {
	errMsg = fmt.Sprintf("payment template invalid; %s", errMsg)
	return sdk.NewError(codespace, CodeInvalidPaymentTemplate, errMsg)
}

func ErrTriedToEffectSubscriptionPaymentWhenShouldnt(codespace sdk.CodespaceType) sdk.Error {
	errMsg := fmt.Sprintf("tried to effect subscription payment when shouldn't")
	return sdk.NewError(codespace, CodeInvalidSubscriptionAction, errMsg)
}

func ErrInvalidId(codespace sdk.CodespaceType, errMsg string) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidId, errMsg)
}

func ErrInvalidArgument(codespace sdk.CodespaceType, errMsg string) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidArgument, errMsg)
}

func ErrAlreadyExists(codespace sdk.CodespaceType, errMsg string) sdk.Error {
	return sdk.NewError(codespace, CodeAlreadyExists, errMsg)
}

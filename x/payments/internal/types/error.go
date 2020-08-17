package types

import (
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	DefaultCodespace                                = ModuleName
	CodeInvalidDistribution                         = 101
	CodeInvalidShare                                = 102
	CodeInvalidPeriod                               = 103
	CodeInvalidPaymentContractAction                = 104
	CodeInvalidDiscount                             = 105
	CodeInvalidDiscountRequest                      = 106
	CodeInvalidPaymentTemplate                      = 107
	CodeInvalidSubscriptionAction                   = 108
	CodeInvalidId                                   = 109
	CodeInvalidArgument                             = 110
	CodeAlreadyExists                               = 111
	ErrNegativeSharePercentage                      = sdkErrors.Register(DefaultCodespace, 102, "payment distribution share percentage must be positive")
	ErrDistributionPercentagesNot100                = sdkErrors.Register(DefaultCodespace, 101, "payment distribution percentages should add up to 100")
	ErrInvalidPeriod                                = sdkErrors.Register(DefaultCodespace, 103, "period is invalid")
	ErrPaymentContractCannotBeDeauthorised          = sdkErrors.Register(DefaultCodespace, 104, "payment contract cannot be deauthorised")
	ErrDiscountIDsBeSequentialFrom1                 = sdkErrors.Register(DefaultCodespace, 105, "discount IDs must be sequential starting with 1")
	ErrNegativeDiscountPercentage                   = sdkErrors.Register(DefaultCodespace, 105, "discount percentage must be positive")
	ErrDiscountPercentageGreaterThan100             = sdkErrors.Register(DefaultCodespace, 105, "discount percentage cannot exceed 100")
	ErrDiscountIdIsNotInTemplate                    = sdkErrors.Register(DefaultCodespace, 106, "discount ID specified is not one of the template's discounts")
	ErrInvalidPaymentTemplate                       = sdkErrors.Register(DefaultCodespace, 107, "payment template invalid")
	ErrTriedToEffectSubscriptionPaymentWhenShouldnt = sdkErrors.Register(DefaultCodespace, 108, "tried to effect subscription payment when shouldn't")
	ErrInvalidId                                    = sdkErrors.Register(DefaultCodespace, 108, "creator did is invalid")
	ErrInvalidArgument                              = sdkErrors.Register(DefaultCodespace, 110, "invalid argument")
	ErrAlreadyExists                                = sdkErrors.Register(DefaultCodespace, 111, "alreday exist")
	ErrInvalidAddress                               = sdkErrors.Register(DefaultCodespace, 112, "payer address is empty")
	ErrorInvalidDid                                 = sdkErrors.Register(DefaultCodespace, 113, "payer did is invalid")
	ErrInternal                                     = sdkErrors.Register(DefaultCodespace, 114, "not allowed format")
)

//func ErrNegativeSharePercentage(codespace sdk.CodespaceType) sdk.Error {
//	errMsg := fmt.Sprintf("payment distribution share percentage must be positive")
//	return sdk.NewError(codespace, CodeInvalidShare, errMsg)
//}

//func ErrDistributionPercentagesNot100(codespace sdk.CodespaceType, total sdk.Dec) sdk.Error {
//	errMsg := fmt.Sprintf("payment distribution percentages should add up to 100, not %s", total.String())
//	return sdk.NewError(codespace, CodeInvalidDistribution, errMsg)
//}

//func ErrInvalidPeriod(codespace sdk.CodespaceType, errMsg string) sdk.Error {
//	errMsg = fmt.Sprintf("period is invalid: %s", errMsg)
//	return sdk.NewError(codespace, CodeInvalidPeriod, errMsg)
//}

//func ErrPaymentContractCannotBeDeauthorised(codespace sdk.CodespaceType) sdk.Error {
//	errMsg := fmt.Sprintf("payment contract cannot be deauthorised")
//	return sdk.NewError(codespace, CodeInvalidPaymentContractAction, errMsg)
//}

//func ErrDiscountIDsBeSequentialFrom1(codespace sdk.CodespaceType) sdk.Error {
//	errMsg := fmt.Sprintf("discount IDs must be sequential starting with 1")
//	return sdk.NewError(codespace, CodeInvalidDiscount, errMsg)
//}

//func ErrNegativeDiscountPercentage(codespace sdk.CodespaceType) sdk.Error {
//	errMsg := fmt.Sprintf("discount percentage must be positive")
//	return sdk.NewError(codespace, CodeInvalidDiscount, errMsg)
//}

//func ErrDiscountPercentageGreaterThan100(codespace sdk.CodespaceType) sdk.Error {
//	errMsg := fmt.Sprintf("discount percentage cannot exceed 100%%")
//	return sdk.NewError(codespace, CodeInvalidDiscount, errMsg)
//}

//func ErrDiscountIdIsNotInTemplate(codespace sdk.CodespaceType) sdk.Error {
//	errMsg := fmt.Sprintf("discount ID specified is not one of the template's discounts")
//	return sdk.NewError(codespace, CodeInvalidDiscountRequest, errMsg)
//}

//func ErrInvalidPaymentTemplate(codespace sdk.CodespaceType, errMsg string) sdk.Error {
//	errMsg = fmt.Sprintf("payment template invalid; %s", errMsg)
//	return sdk.NewError(codespace, CodeInvalidPaymentTemplate, errMsg)
////}
//
//func ErrTriedToEffectSubscriptionPaymentWhenShouldnt(codespace sdk.CodespaceType) sdk.Error {
//	errMsg := fmt.Sprintf("tried to effect subscription payment when shouldn't")
//	return sdk.NewError(codespace, CodeInvalidSubscriptionAction, errMsg)
//}

//func ErrInvalidId(codespace sdk.CodespaceType, errMsg string) sdk.Error {
//	return sdk.NewError(codespace, CodeInvalidId, errMsg)
//}

//func ErrInvalidArgument(codespace sdk.CodespaceType, errMsg string) sdk.Error {
//	return sdk.NewError(codespace, CodeInvalidArgument, errMsg)
//}

//func ErrAlreadyExists(codespace sdk.CodespaceType, errMsg string) sdk.Error {
//	return sdk.NewError(codespace, CodeAlreadyExists, errMsg)
//}

package types

import (
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	DefaultCodespace                                = ModuleName
	ErrNegativeSharePercentage                      = sdkErrors.Register(DefaultCodespace, 2, "payment distribution share percentage must be positive")
	ErrDistributionPercentagesNot100                = sdkErrors.Register(DefaultCodespace, 3, "payment distribution percentages should add up to 100")
	ErrInvalidPeriod                                = sdkErrors.Register(DefaultCodespace, 4, "period is invalid")
	ErrPaymentContractCannotBeDeauthorised          = sdkErrors.Register(DefaultCodespace, 5, "payment contract cannot be deauthorised")
	ErrDiscountIDsBeSequentialFrom1                 = sdkErrors.Register(DefaultCodespace, 6, "discount IDs must be sequential starting with 1")
	ErrNegativeDiscountPercentage                   = sdkErrors.Register(DefaultCodespace, 7, "discount percentage must be positive")
	ErrDiscountPercentageGreaterThan100             = sdkErrors.Register(DefaultCodespace, 8, "discount percentage cannot exceed 100")
	ErrDiscountIdIsNotInTemplate                    = sdkErrors.Register(DefaultCodespace, 9, "discount ID specified is not one of the template's discounts")
	ErrInvalidPaymentTemplate                       = sdkErrors.Register(DefaultCodespace, 10, "payment template invalid")
	ErrTriedToEffectSubscriptionPaymentWhenShouldnt = sdkErrors.Register(DefaultCodespace, 11, "tried to effect subscription payment when shouldn't")
	ErrInvalidId                                    = sdkErrors.Register(DefaultCodespace, 12, "creator did is invalid")
	ErrInvalidArgument                              = sdkErrors.Register(DefaultCodespace, 13, "invalid argument")
	ErrAlreadyExists                                = sdkErrors.Register(DefaultCodespace, 14, "already exist")
	ErrorInvalidDid                                 = sdkErrors.Register(DefaultCodespace, 16, "payer did is invalid")
	ErrInternal                                     = sdkErrors.Register(DefaultCodespace, 17, "not allowed format")
)

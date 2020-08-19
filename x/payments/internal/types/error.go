package types

import (
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	DefaultCodespace                                = ModuleName
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

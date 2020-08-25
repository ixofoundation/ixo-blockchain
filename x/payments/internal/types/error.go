package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	DefaultCodespace                                = ModuleName
	ErrNegativeSharePercentage                      = sdkerrors.Register(DefaultCodespace, 2, "payment distribution share percentage must be positive")
	ErrDistributionPercentagesNot100                = sdkerrors.Register(DefaultCodespace, 3, "payment distribution percentages should add up to 100")
	ErrInvalidPeriod                                = sdkerrors.Register(DefaultCodespace, 4, "period is invalid")
	ErrPaymentContractCannotBeDeauthorised          = sdkerrors.Register(DefaultCodespace, 5, "payment contract cannot be deauthorised")
	ErrDiscountIDsBeSequentialFrom1                 = sdkerrors.Register(DefaultCodespace, 6, "discount IDs must be sequential starting with 1")
	ErrNegativeDiscountPercentage                   = sdkerrors.Register(DefaultCodespace, 7, "discount percentage must be positive")
	ErrDiscountPercentageGreaterThan100             = sdkerrors.Register(DefaultCodespace, 8, "discount percentage cannot exceed 100")
	ErrDiscountIdIsNotInTemplate                    = sdkerrors.Register(DefaultCodespace, 9, "discount ID specified is not one of the template's discounts")
	ErrInvalidPaymentTemplate                       = sdkerrors.Register(DefaultCodespace, 10, "payment template invalid")
	ErrTriedToEffectSubscriptionPaymentWhenShouldnt = sdkerrors.Register(DefaultCodespace, 11, "tried to effect subscription payment when shouldn't")
	ErrInvalidId                                    = sdkerrors.Register(DefaultCodespace, 12, "creator did is invalid")
	ErrInvalidArgument                              = sdkerrors.Register(DefaultCodespace, 13, "invalid argument")
	ErrAlreadyExists                                = sdkerrors.Register(DefaultCodespace, 14, "already exist")
	ErrorInvalidDid                                 = sdkerrors.Register(DefaultCodespace, 16, "payer did is invalid")
	ErrInternal                                     = sdkerrors.Register(DefaultCodespace, 17, "not allowed format")
)

package types

import "cosmossdk.io/errors"

// Sentinel errors for the liquidstake module.
var (
	ErrActiveLiquidValidatorsNotExists              = errors.Register(ModuleName, 1, "active liquid validators not exists")
	ErrInvalidBondDenom                             = errors.Register(ModuleName, 2, "invalid bond denom")
	ErrInvalidLiquidBondDenom                       = errors.Register(ModuleName, 3, "invalid liquid bond denom")
	ErrLessThanMinLiquidStakeAmount                 = errors.Register(ModuleName, 4, "staking amount should be over params.min_liquid_stake_amount")
	ErrInvalidStkIXOSupply                          = errors.Register(ModuleName, 5, "invalid liquid bond denom supply")
	ErrInvalidActiveLiquidValidators                = errors.Register(ModuleName, 6, "invalid active liquid validators")
	ErrLiquidValidatorsNotExists                    = errors.Register(ModuleName, 7, "liquid validators not exists")
	ErrInsufficientProxyAccBalance                  = errors.Register(ModuleName, 8, "insufficient liquid tokens or balance of proxy account, need to wait for new liquid validator to be added or unbonding of proxy account to be completed")
	ErrTooSmallLiquidStakeAmount                    = errors.Register(ModuleName, 9, "liquid stake amount is too small, the result becomes zero")
	ErrTooSmallLiquidUnstakingAmount                = errors.Register(ModuleName, 10, "liquid unstaking amount is too small, the result becomes zero")
	ErrLSMTokenizeFailed                            = errors.Register(ModuleName, 11, "LSM tokenization failed")
	ErrLSMRedeemFailed                              = errors.Register(ModuleName, 12, "LSM redemption failed")
	ErrWhitelistedValidatorsList                    = errors.Register(ModuleName, 13, "whitelisted validators list incorrect")
	ErrActiveLiquidValidatorsWeightQuorumNotReached = errors.Register(ModuleName, 14, "active liquid validators weight quorum not reached")
	ErrModulePaused                                 = errors.Register(ModuleName, 15, "module functions have been paused")
	ErrDelegationFailed                             = errors.Register(ModuleName, 16, "delegation failed")
	ErrInvalidResponse                              = errors.Register(ModuleName, 17, "invalid response")
	ErrUnstakeFailed                                = errors.Register(ModuleName, 18, "Unstaking failed")
	ErrRedelegateFailed                             = errors.Register(ModuleName, 19, "Redelegate failed")
	ErrRatioMoreThanOne                             = errors.Register(ModuleName, 20, "ratio should be less than or equal to 1")
)

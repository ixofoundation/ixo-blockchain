package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func CheckReserveTokenNames(resTokens []string, token string) error {
	// Check that no token is the same as the main token, no token
	// is duplicate, and that the token is a valid denomination
	uniqueReserveTokens := make(map[string]string)
	for _, r := range resTokens {
		// Check if same as main token
		if r == token {
			return ErrBondTokenCannotAlsoBeReserveToken
		}

		// Check if duplicate
		if _, ok := uniqueReserveTokens[r]; ok {
			return ErrDuplicateReserveToken
		} else {
			uniqueReserveTokens[r] = ""
		}

		// Check if can be parsed as coin
		err := CheckCoinDenom(r)
		if err != nil {
			return err
		}
	}
	return nil
}

func CheckNoOfReserveTokens(resTokens []string, fnType string) error {
	// Come up with number of expected reserve tokens
	expectedNoOfTokens, ok := NoOfReserveTokensForFunctionType[fnType]
	if !ok {
		return sdkerrors.Wrap(ErrUnrecognizedFunctionType, fnType)
	}

	// Check that number of reserve tokens is correct (if expecting a specific number of tokens)
	if expectedNoOfTokens != AnyNumberOfReserveTokens && len(resTokens) != expectedNoOfTokens {
		return sdkerrors.Wrapf(ErrIncorrectNumberOfReserveTokens, "expected: %d", expectedNoOfTokens)
	}

	return nil
}

func CheckCoinDenom(denom string) (err error) {
	coin, err := sdk.ParseCoinNormalized("0" + denom)
	if err != nil {
		return err
	} else if denom != coin.Denom {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, coin.String())
	}
	return nil
}

func GetRequiredParamsForFunctionType(fnType string) (fnParams []string, err error) {
	expectedParams, ok := RequiredParamsForFunctionType[fnType]
	if !ok {
		return nil, sdkerrors.Wrap(ErrUnrecognizedFunctionType, fnType)
	}
	return expectedParams, nil
}

func GetExceptionsForFunctionType(fnType string) (restrictions FunctionParamRestrictions, err error) {
	restrictions, ok := ExtraParameterRestrictions[fnType]
	if !ok {
		return nil, sdkerrors.Wrap(ErrUnrecognizedFunctionType, fnType)
	}
	return restrictions, nil
}

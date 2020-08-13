package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func CheckReserveTokenNames(resTokens []string, token string) sdk.Error {
	// Check that no token is the same as the main token, no token
	// is duplicate, and that the token is a valid denomination
	uniqueReserveTokens := make(map[string]string)
	for _, r := range resTokens {
		// Check if same as main token
		if r == token {
			return ErrBondTokenCannotAlsoBeReserveToken(DefaultCodespace)
		}

		// Check if duplicate
		if _, ok := uniqueReserveTokens[r]; ok {
			return ErrDuplicateReserveToken(DefaultCodespace)
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

func CheckNoOfReserveTokens(resTokens []string, fnType string) sdk.Error {
	// Come up with number of expected reserve tokens
	expectedNoOfTokens, ok := NoOfReserveTokensForFunctionType[fnType]
	if !ok {
		return ErrUnrecognizedFunctionType(DefaultCodespace)
	}

	// Check that number of reserve tokens is correct (if expecting a specific number of tokens)
	if expectedNoOfTokens != AnyNumberOfReserveTokens && len(resTokens) != expectedNoOfTokens {
		return ErrIncorrectNumberOfReserveTokens(DefaultCodespace, expectedNoOfTokens)
	}

	return nil
}

func CheckCoinDenom(denom string) (err sdk.Error) {
	coin, err2 := sdk.ParseCoin("0" + denom)
	if err2 != nil {
		return sdk.ErrInternal(err2.Error())
	} else if denom != coin.Denom {
		return ErrInvalidCoinDenomination(DefaultCodespace, denom)
	}
	return nil
}

func GetRequiredParamsForFunctionType(fnType string) (fnParams []string, err sdk.Error) {
	expectedParams, ok := RequiredParamsForFunctionType[fnType]
	if !ok {
		return nil, ErrUnrecognizedFunctionType(DefaultCodespace)
	}
	return expectedParams, nil
}

func GetExceptionsForFunctionType(fnType string) (restrictions FunctionParamRestrictions, err sdk.Error) {
	restrictions, ok := ExtraParameterRestrictions[fnType]
	if !ok {
		return nil, ErrUnrecognizedFunctionType(DefaultCodespace)
	}
	return restrictions, nil
}

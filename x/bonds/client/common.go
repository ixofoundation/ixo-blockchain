package client

import (
	"strings"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ixofoundation/ixo-blockchain/v3/x/bonds/types"
)

func splitParameters(fnParamsStr string) (paramValuePairs []string) {
	// If empty, just return empty list
	if strings.TrimSpace(fnParamsStr) != "" {
		// Split "a:1,b:2" into ["a:1","b:2"]
		paramValuePairs = strings.Split(fnParamsStr, ",")
	}
	return paramValuePairs
}

func paramsListToMap(paramValuePairs []string) (paramsFieldMap map[string]string, err error) {
	paramsFieldMap = make(map[string]string)
	for _, pv := range paramValuePairs {
		// Split each "a:1" into ["a","1"]
		pvArray := strings.SplitN(pv, ":", 2)
		if len(pvArray) != 2 {
			return nil, errorsmod.Wrap(types.ErrInvalidFunctionParameter, pv)
		}
		paramsFieldMap[pvArray[0]] = pvArray[1]
	}
	return paramsFieldMap, nil
}

func paramsMapToObj(paramsFieldMap map[string]string) (functionParams types.FunctionParams, err error) {
	for p, v := range paramsFieldMap {
		vDec, err := math.LegacyNewDecFromStr(v)
		if err != nil {
			return nil, errorsmod.Wrap(types.ErrArgumentMissingOrNonFloat, p)
		} else {
			functionParams = append(functionParams, types.NewFunctionParam(p, vDec))
		}
	}
	return functionParams, nil
}

// TODO Consider bypassing the use of a map to avoid non-determinism on the order of parameters at the client side
func ParseFunctionParams(fnParamsStr string) (fnParams types.FunctionParams, err error) {
	// Split (if not empty) and check number of parameters
	paramValuePairs := splitParameters(fnParamsStr)

	// Parse function parameters into map
	paramsFieldMap, err := paramsListToMap(paramValuePairs)
	if err != nil {
		return nil, err
	}

	// Parse parameters into floats
	functionParams, err := paramsMapToObj(paramsFieldMap)
	if err != nil {
		return nil, err
	}

	return functionParams, nil
}

func ParseTwoPartCoin(amount, denom string) (coin sdk.Coin, err error) {
	coin, err = sdk.ParseCoinNormalized(amount + denom)
	if err != nil {
		return sdk.Coin{}, err
	} else if denom != coin.Denom {
		return sdk.Coin{}, errorsmod.Wrap(sdkerrors.ErrInvalidCoins, denom)
	}
	return coin, nil
}

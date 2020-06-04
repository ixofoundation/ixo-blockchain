package client

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/internal/types"
	"github.com/ixofoundation/ixo-blockchain/x/ixo/sovrin"
	"strings"
)

func splitParameters(fnParamsStr string) (paramValuePairs []string) {
	// If empty, just return empty list
	if strings.TrimSpace(fnParamsStr) != "" {
		// Split "a:1,b:2" into ["a:1","b:2"]
		paramValuePairs = strings.Split(fnParamsStr, ",")
	}
	return paramValuePairs
}

func paramsListToMap(paramValuePairs []string) (paramsFieldMap map[string]string, err sdk.Error) {
	paramsFieldMap = make(map[string]string)
	for _, pv := range paramValuePairs {
		// Split each "a:1" into ["a","1"]
		pvArray := strings.SplitN(pv, ":", 2)
		if len(pvArray) != 2 {
			return nil, types.ErrInvalidFunctionParameter(types.DefaultCodespace, pv)
		}
		paramsFieldMap[pvArray[0]] = pvArray[1]
	}
	return paramsFieldMap, nil
}

func paramsMapToObj(paramsFieldMap map[string]string) (functionParams types.FunctionParams, err sdk.Error) {
	for p, v := range paramsFieldMap {
		vInt, ok := sdk.NewIntFromString(v)
		if !ok {
			return nil, types.ErrArgumentMissingOrNonInteger(types.DefaultCodespace, p)
		} else {
			functionParams = append(functionParams, types.NewFunctionParam(p, vInt))
		}
	}
	return functionParams, nil
}

func ParseFunctionParams(fnParamsStr string) (fnParams types.FunctionParams, err sdk.Error) {
	// Split (if not empty) and check number of parameters
	paramValuePairs := splitParameters(fnParamsStr)

	// Parse function parameters into map
	paramsFieldMap, err := paramsListToMap(paramValuePairs)
	if err != nil {
		return nil, err
	}

	// Parse parameters into integers
	functionParams, err := paramsMapToObj(paramsFieldMap)
	if err != nil {
		return nil, err
	}

	return functionParams, nil
}

func ParseTwoPartCoin(amount, denom string) (coin sdk.Coin, err error) {
	coin, err = sdk.ParseCoin(amount + denom)
	if err != nil {
		return sdk.Coin{}, err
	} else if denom != coin.Denom {
		return sdk.Coin{}, types.ErrInvalidCoinDenomination(types.DefaultCodespace, denom)
	}
	return coin, nil
}

func UnmarshalSovrinDID(sovrinJson string) sovrin.SovrinDid {
	sovrinDid := sovrin.SovrinDid{}
	sovrinErr := json.Unmarshal([]byte(sovrinJson), &sovrinDid)
	if sovrinErr != nil {
		panic(sovrinErr)
	}

	return sovrinDid
}

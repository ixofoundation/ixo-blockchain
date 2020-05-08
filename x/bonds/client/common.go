package client

import (
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/internal/types"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"github.com/ixofoundation/ixo-blockchain/x/ixo/sovrin"
	"strings"
)

func getRequiredParamsForFunctionType(fnType string) (fnParams []string, err sdk.Error) {
	expectedParams, ok := types.RequiredParamsForFunctionType[fnType]
	if !ok {
		return nil, types.ErrUnrecognizedFunctionType(types.DefaultCodespace)
	}
	return expectedParams, nil
}

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

func paramsMapToObj(paramsFieldMap map[string]string, expectedParams []string) (functionParams types.FunctionParams, err sdk.Error) {
	for _, p := range expectedParams {
		val, ok := sdk.NewIntFromString(paramsFieldMap[p])
		if !ok {
			return nil, types.ErrFunctionParameterMissingOrNonInteger(types.DefaultCodespace, p)
		} else {
			functionParams = append(functionParams, types.NewFunctionParam(p, val))
		}
	}
	return functionParams, nil
}

func ParseFunctionParams(fnParamsStr string, fnType string) (fnParams types.FunctionParams, err sdk.Error) {

	// Come up with list of expected parameters
	expectedParams, err := getRequiredParamsForFunctionType(fnType)
	if err != nil {
		return nil, err
	}

	// Split (if not empty) and check number of parameters
	paramValuePairs := splitParameters(fnParamsStr)
	if len(paramValuePairs) != len(expectedParams) {
		return nil, types.ErrIncorrectNumberOfFunctionParameters(types.DefaultCodespace, len(expectedParams))
	}

	// Parse function parameters into map
	paramsFieldMap, err := paramsListToMap(paramValuePairs)
	if err != nil {
		return nil, err
	}

	// Parse parameters into integers
	functionParams, err := paramsMapToObj(paramsFieldMap, expectedParams)
	if err != nil {
		return nil, err
	}

	return functionParams, nil
}

func checkReserveTokenNames(resTokens []string, token string) error {
	// Check that no token is the same as the main token, no token
	// is duplicate, and that the token is a valid denomination
	uniqueReserveTokens := make(map[string]string)
	for _, r := range resTokens {
		// Check if same as main token
		if r == token {
			return types.ErrBondTokenCannotAlsoBeReserveToken(types.DefaultCodespace)
		}

		// Check if duplicate
		if _, ok := uniqueReserveTokens[r]; ok {
			return types.ErrDuplicateReserveToken(types.DefaultCodespace)
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

func checkNoOfReserveTokens(resTokens []string, fnType string) sdk.Error {
	// Come up with number of expected reserve tokens
	expectedNoOfTokens, ok := types.NoOfReserveTokensForFunctionType[fnType]
	if !ok {
		return types.ErrUnrecognizedFunctionType(types.DefaultCodespace)
	}

	// Check that number of reserve tokens is correct (if expecting a specific number of tokens)
	if expectedNoOfTokens != types.AnyNumberOfReserveTokens && len(resTokens) != expectedNoOfTokens {
		return types.ErrIncorrectNumberOfReserveTokens(types.DefaultCodespace, expectedNoOfTokens)
	}

	return nil
}

func ParseReserveTokens(resTokensStr string, fnType string, token string) (resTokens []string, err error) {
	resTokens = strings.Split(resTokensStr, ",")
	if err = checkReserveTokenNames(resTokens, token); err != nil {
		return nil, err
	} else if err = checkNoOfReserveTokens(resTokens, fnType); err != nil {
		return nil, err
	}
	return resTokens, nil
}

func ParseMaxSupply(maxSupplyStr string, token string) (coin sdk.Coin, err error) {
	maxSupply, err := sdk.ParseCoin(maxSupplyStr)
	if err != nil {
		return sdk.Coin{}, err
	} else if maxSupply.Denom != token {
		return sdk.Coin{}, types.ErrMaxSupplyDenomDoesNotMatchTokenDenom(types.DefaultCodespace)
	}
	return maxSupply, nil
}

func parseNonNegativeDec(decStr string, decName string) (dec sdk.Dec, err sdk.Error) {
	dec, err = sdk.NewDecFromStr(decStr)
	if err != nil {
		return sdk.Dec{}, types.ErrArgumentMissingOrNonFloat(types.DefaultCodespace, decName)
	} else if dec.IsNegative() {
		return sdk.Dec{}, types.ErrArgumentCannotBeNegative(types.DefaultCodespace, decName)
	}
	return dec, nil
}

func ParseSanityValues(sanityRateStr string, sanityMarginPercentageStr string) (sanityRate, sanityMarginPercentage sdk.Dec, err sdk.Error) {

	// If sanity rate is provided, margin percentage has to be provided
	// If sanity rate is not provided, margin percentage is ignored

	if sanityRateStr == "" {
		sanityRate = sdk.ZeroDec()
		sanityMarginPercentage = sdk.ZeroDec()
	} else {
		// Check that both parsable and not negative
		sanityRate, err = parseNonNegativeDec(sanityRateStr, "sanity rate")
		if err != nil {
			return sdk.Dec{}, sdk.Dec{}, err
		}
		sanityMarginPercentage, err = parseNonNegativeDec(sanityMarginPercentageStr, "sanity margin percentage")
		if err != nil {
			return sdk.Dec{}, sdk.Dec{}, err
		}
	}

	return sanityRate, sanityMarginPercentage, nil
}

func ParseBatchBlocks(batchBlocksStr string) (batchBlocks sdk.Uint, err error) {

	batchBlocks, err = sdk.ParseUint(batchBlocksStr)
	if err != nil {
		return sdk.Uint{}, types.ErrArgumentMissingOrNonUInteger(types.DefaultCodespace, "max batch blocks")
	}
	return batchBlocks, nil
}

func CheckCoinDenom(denom string) (err error) {
	coin, err := sdk.ParseCoin("0" + denom)
	if err != nil {
		return err
	} else if denom != coin.Denom {
		return types.ErrInvalidCoinDenomination(types.DefaultCodespace, denom)
	}
	return nil
}

func ParseCoin(amount, denom string) (coin sdk.Coin, err error) {
	coin, err = sdk.ParseCoin(amount + denom)
	if err != nil {
		return sdk.Coin{}, err
	} else if denom != coin.Denom {
		return sdk.Coin{}, types.ErrInvalidCoinDenomination(types.DefaultCodespace, denom)
	}
	return coin, nil
}

func IxoSignAndBroadcast(cdc *codec.Codec, ctx context.CLIContext, msg sdk.Msg, sovrinDid sovrin.SovrinDid) error {
	privKey := [64]byte{}
	copy(privKey[:], base58.Decode(sovrinDid.Secret.SignKey))
	copy(privKey[32:], base58.Decode(sovrinDid.VerifyKey))

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	signature := ixo.SignIxoMessage(msgBytes, sovrinDid.Did, privKey)
	tx := ixo.NewIxoTxSingleMsg(msg, signature)

	bz, err := cdc.MarshalJSON(tx)
	if err != nil {
		panic(err)
	}

	res, err := ctx.BroadcastTx(bz)
	if err != nil {
		return err
	}

	fmt.Println(res.String())
	fmt.Printf("Committed at block %d. Hash: %s\n", res.Height, res.TxHash)
	return nil

}

func UnmarshalSovrinDID(sovrinJson string) sovrin.SovrinDid {
	sovrinDid := sovrin.SovrinDid{}
	sovrinErr := json.Unmarshal([]byte(sovrinJson), &sovrinDid)
	if sovrinErr != nil {
		panic(sovrinErr)
	}

	return sovrinDid
}

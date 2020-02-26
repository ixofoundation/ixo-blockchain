package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
	"strings"
)

// Local code type
type CodeType = sdk.CodeType

const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	// General
	CodeArgumentInvalid                CodeType = 301
	CodeArgumentMissingOrIncorrectType CodeType = 302
	CodeIncorrectNumberOfValues        CodeType = 303

	// Bonds
	CodeBondDoesNotExist        CodeType = 304
	CodeBondAlreadyExists       CodeType = 305
	CodeBondDoesNotAllowSelling CodeType = 306
	CodeDidNotEditAnything      CodeType = 307
	CodeInvalidSwapper          CodeType = 308
	CodeInvalidBond             CodeType = 309

	// Function types and function parameters
	CodeUnrecognizedFunctionType             CodeType = 310
	CodeInvalidFunctionParameter             CodeType = 311
	CodeFunctionNotAvailableForFunctionType  CodeType = 312
	CodeFunctionRequiresNonZeroCurrentSupply CodeType = 313

	// Token/coin names
	CodeReserveTokenInvalid     CodeType = 314
	CodeMaxSupplyDenomInvalid   CodeType = 315
	CodeBondTokenInvalid        CodeType = 316
	CodeReserveDenomsMismatch   CodeType = 317
	CodeInvalidCoinDenomination CodeType = 318

	// Amounts and fees
	CodeInvalidResultantSupply CodeType = 319
	CodeMaxPriceExceeded       CodeType = 320
	CodeSwapAmountInvalid      CodeType = 321
	CodeOrderLimitExceeded     CodeType = 322
	CodeSanityRateViolated     CodeType = 323
	CodeFeeTooLarge            CodeType = 324
)

func ErrArgumentCannotBeEmpty(codespace sdk.CodespaceType, argument string) sdk.Error {
	errMsg := fmt.Sprintf("%s argument cannot be empty", argument)
	return sdk.NewError(codespace, CodeArgumentInvalid, errMsg)
}

func ErrArgumentCannotBeNegative(codespace sdk.CodespaceType, arg string) sdk.Error {
	errMsg := fmt.Sprintf("%s argument cannot be negative", arg)
	return sdk.NewError(codespace, CodeArgumentInvalid, errMsg)
}

func ErrArgumentMustBePositive(codespace sdk.CodespaceType, arg string) sdk.Error {
	errMsg := fmt.Sprintf("%s argument must be a positive value", arg)
	return sdk.NewError(codespace, CodeArgumentInvalid, errMsg)
}

func ErrFunctionParameterMissingOrNonInteger(codespace sdk.CodespaceType, param string) sdk.Error {
	errMsg := fmt.Sprintf("%s parameter is missing or is not an integer", param)
	return sdk.NewError(codespace, CodeArgumentMissingOrIncorrectType, errMsg)
}

func ErrArgumentMissingOrNonFloat(codespace sdk.CodespaceType, arg string) sdk.Error {
	errMsg := fmt.Sprintf("%s argument is missing or is not a float", arg)
	return sdk.NewError(codespace, CodeArgumentMissingOrIncorrectType, errMsg)
}

func ErrArgumentMissingOrNonInteger(codespace sdk.CodespaceType, arg string) sdk.Error {
	errMsg := fmt.Sprintf("%s argument is missing or is not an integer", arg)
	return sdk.NewError(codespace, CodeArgumentMissingOrIncorrectType, errMsg)
}

func ErrArgumentMissingOrNonUInteger(codespace sdk.CodespaceType, arg string) sdk.Error {
	errMsg := fmt.Sprintf("%s argument is missing or is not an unsigned integer", arg)
	return sdk.NewError(codespace, CodeArgumentMissingOrIncorrectType, errMsg)
}

func ErrArgumentMissingOrNonBoolean(codespace sdk.CodespaceType, arg string) sdk.Error {
	errMsg := fmt.Sprintf("%s argument is missing or is not true or false", arg)
	return sdk.NewError(codespace, CodeArgumentMissingOrIncorrectType, errMsg)
}

func ErrIncorrectNumberOfReserveTokens(codespace sdk.CodespaceType, expected int) sdk.Error {
	errMsg := fmt.Sprintf("Incorrect number of reserve tokens; expected: %d", expected)
	return sdk.NewError(codespace, CodeIncorrectNumberOfValues, errMsg)
}

func ErrIncorrectNumberOfFunctionParameters(codespace sdk.CodespaceType, expected int) sdk.Error {
	errMsg := fmt.Sprintf("Incorrect number of function parameters; expected: %d", expected)
	return sdk.NewError(codespace, CodeIncorrectNumberOfValues, errMsg)
}

func ErrBondDoesNotExist(codespace sdk.CodespaceType, bondDid ixo.Did) sdk.Error {
	errMsg := fmt.Sprintf("Bond '%s' does not exist", bondDid)
	return sdk.NewError(codespace, CodeBondDoesNotExist, errMsg)
}

func ErrBondAlreadyExists(codespace sdk.CodespaceType, bondDid ixo.Did) sdk.Error {
	errMsg := fmt.Sprintf("Bond '%s' already exists", bondDid)
	return sdk.NewError(codespace, CodeBondAlreadyExists, errMsg)
}

func ErrBondDoesNotAllowSelling(codespace sdk.CodespaceType) sdk.Error {
	errMsg := "Bond does not allow selling."
	return sdk.NewError(codespace, CodeBondDoesNotAllowSelling, errMsg)
}

func ErrDidNotEditAnything(codespace sdk.CodespaceType) sdk.Error {
	errMsg := "Did not edit anything from the bond"
	return sdk.NewError(codespace, CodeDidNotEditAnything, errMsg)
}

func ErrFromAndToCannotBeTheSameToken(codespace sdk.CodespaceType) sdk.Error {
	errMsg := "From and To tokens cannot be the same token"
	return sdk.NewError(codespace, CodeInvalidSwapper, errMsg)
}

func ErrDuplicateReserveToken(codespace sdk.CodespaceType) sdk.Error {
	errMsg := "Cannot have duplicate tokens in reserve tokens"
	return sdk.NewError(codespace, CodeInvalidBond, errMsg)
}

func ErrUnrecognizedFunctionType(codespace sdk.CodespaceType) sdk.Error {
	errMsg := "Unrecognized function type"
	return sdk.NewError(codespace, CodeUnrecognizedFunctionType, errMsg)
}

func ErrInvalidFunctionParameter(codespace sdk.CodespaceType, parameter string) sdk.Error {
	errMsg := fmt.Sprintf("Invalid function parameter '%s'", parameter)
	return sdk.NewError(codespace, CodeInvalidFunctionParameter, errMsg)
}

func ErrFunctionNotAvailableForFunctionType(codespace sdk.CodespaceType) sdk.Error {
	errMsg := "Function is not available for the function type"
	return sdk.NewError(codespace, CodeFunctionNotAvailableForFunctionType, errMsg)
}

func ErrFunctionRequiresNonZeroCurrentSupply(codespace sdk.CodespaceType) sdk.Error {
	errMsg := "Function requires the current supply to be non zero"
	return sdk.NewError(codespace, CodeFunctionRequiresNonZeroCurrentSupply, errMsg)
}

func ErrTokenIsNotAValidReserveToken(codespace sdk.CodespaceType, denom string) sdk.Error {
	errMsg := fmt.Sprintf("Token '%s' is not a valid reserve token", denom)
	return sdk.NewError(codespace, CodeReserveTokenInvalid, errMsg)
}

func ErrMaxSupplyDenomDoesNotMatchTokenDenom(codespace sdk.CodespaceType) sdk.Error {
	errMsg := "Max supply denom does not match token denom"
	return sdk.NewError(codespace, CodeMaxSupplyDenomInvalid, errMsg)
}

func ErrBondTokenCannotAlsoBeReserveToken(codespace sdk.CodespaceType) sdk.Error {
	errMsg := "Token cannot also be a reserve token"
	return sdk.NewError(codespace, CodeBondTokenInvalid, errMsg)
}

func ErrBondTokenCannotBeStakingToken(codespace sdk.CodespaceType) sdk.Error {
	errMsg := "Bond token cannot be staking token"
	return sdk.NewError(codespace, CodeBondTokenInvalid, errMsg)
}

func ErrBondTokenDoesNotMatchBond(codespace sdk.CodespaceType) sdk.Error {
	errMsg := "Bond token does not match bond"
	return sdk.NewError(codespace, CodeBondTokenInvalid, errMsg)
}

func ErrReserveDenomsMismatch(codespace sdk.CodespaceType, inputDenoms string, actualDenoms []string) sdk.Error {
	errMsg := fmt.Sprintf("Denoms in %s do not match reserve denoms; expected: %s", inputDenoms, strings.Join(actualDenoms, ","))
	return sdk.NewError(codespace, CodeReserveDenomsMismatch, errMsg)
}

func ErrInvalidCoinDenomination(codespace sdk.CodespaceType, denom string) sdk.Error {
	errMsg := fmt.Sprintf("Invalid coin denomination '%s'", denom)
	return sdk.NewError(codespace, CodeInvalidCoinDenomination, errMsg)
}

func ErrCannotMintMoreThanMaxSupply(codespace sdk.CodespaceType) sdk.Error {
	errMsg := "Cannot mint more tokens than the max supply"
	return sdk.NewError(codespace, CodeInvalidResultantSupply, errMsg)
}

func ErrCannotBurnMoreThanSupply(codespace sdk.CodespaceType) sdk.Error {
	errMsg := "Cannot burn more tokens than the current supply"
	return sdk.NewError(codespace, CodeInvalidResultantSupply, errMsg)
}

func ErrMaxPriceExceeded(codespace sdk.CodespaceType, totalPrice, maxPrice sdk.Coins) sdk.Error {
	errMsg := fmt.Sprintf("Actual prices %s exceed max prices %s", totalPrice.String(), maxPrice.String())
	return sdk.NewError(codespace, CodeMaxPriceExceeded, errMsg)
}

func ErrSwapAmountTooSmallToGiveAnyReturn(codespace sdk.CodespaceType, fromToken, toToken string) sdk.Error {
	errMsg := fmt.Sprintf("%s swap amount too small to give any %s return", fromToken, toToken)
	return sdk.NewError(codespace, CodeSwapAmountInvalid, errMsg)
}

func ErrSwapAmountCausesReserveDepletion(codespace sdk.CodespaceType, fromToken, toToken string) sdk.Error {
	errMsg := fmt.Sprintf("%s swap amount too large and causes %s reserve to be depleted", fromToken, toToken)
	return sdk.NewError(codespace, CodeSwapAmountInvalid, errMsg)
}

func ErrOrderQuantityLimitExceeded(codespace sdk.CodespaceType) sdk.Error {
	errMsg := "Order quantity limits exceeded"
	return sdk.NewError(codespace, CodeOrderLimitExceeded, errMsg)
}

func ErrValuesViolateSanityRate(codespace sdk.CodespaceType) sdk.Error {
	errMsg := "Values violate sanity rate"
	return sdk.NewError(codespace, CodeSanityRateViolated, errMsg)
}

func ErrFeesCannotBeOrExceed100Percent(codespace sdk.CodespaceType) sdk.Error {
	errMsg := "Sum of fees is or exceeds 100 percent"
	return sdk.NewError(codespace, CodeFeeTooLarge, errMsg)
}

package types

import (
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"sort"
)

type (
	BondState              string
	BondStateTransitionMap map[BondState][]BondState
)

const (
	PowerFunction     = "power_function"
	SigmoidFunction   = "sigmoid_function"
	SwapperFunction   = "swapper_function"
	AugmentedFunction = "augmented_function"

	HatchState  BondState = "HATCH"
	OpenState   BondState = "OPEN"
	SettleState BondState = "SETTLE"
	FailedState BondState = "FAILED"

	DoNotModifyField = "[do-not-modify]"

	AnyNumberOfReserveTokens = -1
)

var StateTransitions = initStateTransitions()

func initStateTransitions() BondStateTransitionMap {
	return BondStateTransitionMap{
		HatchState:  {OpenState, FailedState},
		OpenState:   {SettleState, FailedState},
		SettleState: {},
		FailedState: {},
	}
}

func (next BondState) IsValidProgressionFrom(prev BondState) bool {
	validStatuses := StateTransitions[prev]
	for _, v := range validStatuses {
		if v == next {
			return true
		}
	}
	return false
}

func (s BondState) String() string {
	return string(s)
}

func BondStateFromString(s string) BondState {
	return BondState(s)
}

type FunctionParamRestrictions func(paramsMap map[string]sdk.Dec) error

var (
	RequiredParamsForFunctionType = map[string][]string{
		PowerFunction:     {"m", "n", "c"},
		SigmoidFunction:   {"a", "b", "c"},
		SwapperFunction:   nil,
		AugmentedFunction: {"d0", "p0", "theta", "kappa"},
	}

	NoOfReserveTokensForFunctionType = map[string]int{
		PowerFunction:     AnyNumberOfReserveTokens,
		SigmoidFunction:   AnyNumberOfReserveTokens,
		SwapperFunction:   2,
		AugmentedFunction: AnyNumberOfReserveTokens,
	}

	ExtraParameterRestrictions = map[string]FunctionParamRestrictions{
		PowerFunction:     nil,
		SigmoidFunction:   sigmoidParameterRestrictions,
		SwapperFunction:   nil,
		AugmentedFunction: augmentedParameterRestrictions,
	}
)

//type FunctionParam struct {
//	Param string  `json:"param" yaml:"param"`
//	Value sdk.Dec `json:"value" yaml:"value"`
//}

func NewFunctionParam(param string, value sdk.Dec) FunctionParam {
	return FunctionParam{
		Param: param,
		Value: value,
	}
}

type FunctionParams []FunctionParam

func (fps FunctionParams) ReplaceParam(param string, value sdk.Dec) {
	for i, fp := range fps {
		if fp.Param == param {
			fps[i] = NewFunctionParam(param, value)
			return
		}
	}
}

func (fps FunctionParams) AddParam(param string, value sdk.Dec) FunctionParams {
	return append(fps, NewFunctionParam(param, value))
}

func (fps FunctionParams) AddParams(newFps FunctionParams) FunctionParams {
	return append(fps, newFps...)
}

func (fps FunctionParams) Validate(functionType string) error {
	// Come up with list of expected parameters
	expectedParams, err := GetRequiredParamsForFunctionType(functionType)
	if err != nil {
		return err
	}

	// Check that number of params is as expected
	if len(fps) != len(expectedParams) {
		return sdkerrors.Wrapf(ErrIncorrectNumberOfFunctionParameters, "expected: %d", len(expectedParams))
	}

	// Check that params match and all values are non-negative
	paramsMap := fps.AsMap()
	for _, p := range expectedParams {
		val, ok := paramsMap[p]
		if !ok {
			return sdkerrors.Wrapf(ErrFunctionParameterMissingOrNonFloat, p)
		} else if val.IsNegative() {
			return sdkerrors.Wrapf(ErrArgumentCannotBeNegative, p)
		}
	}

	// Get extra function parameter restrictions
	extraRestrictions, err := GetExceptionsForFunctionType(functionType)
	if err != nil {
		return err
	}
	if extraRestrictions != nil {
		err := extraRestrictions(paramsMap)
		if err != nil {
			return err
		}
	}

	return nil
}

func (fps FunctionParams) String() (result string) {
	output, err := json.Marshal(fps)
	if err != nil {
		panic(err)
	}
	return string(output)
}

func (fps FunctionParams) AsMap() (paramsMap map[string]sdk.Dec) {
	paramsMap = make(map[string]sdk.Dec)
	for _, fp := range fps {
		paramsMap[fp.Param] = fp.Value
	}
	return paramsMap
}

func sigmoidParameterRestrictions(paramsMap map[string]sdk.Dec) error {
	// Sigmoid exception 1: c != 0, otherwise we run into divisions by zero
	val, ok := paramsMap["c"]
	if !ok {
		panic("did not find parameter c for sigmoid function")
	} else if !val.IsPositive() {
		return sdkerrors.Wrap(ErrArgumentMustBePositive, "c")
	}
	return nil
}

func augmentedParameterRestrictions(paramsMap map[string]sdk.Dec) error {
	// Augmented exception 1.1: d0 must be an integer, since it is a token amount
	// Augmented exception 1.2: d0 > 0, since a negative raise is not valid, but
	// also to avoid division by zero
	val, ok := paramsMap["d0"]
	if !ok {
		panic("did not find parameter d0 for augmented function")
	} else if !val.TruncateDec().Equal(val) {
		return sdkerrors.Wrap(ErrArgumentMustBeInteger, "d0")
	} else if !val.IsPositive() {
		return sdkerrors.Wrap(ErrArgumentMustBePositive, "d0")
	}

	// Augmented exception 2: p0 > 0, since a negative price is not valid, but
	// also to avoid division by zero
	val, ok = paramsMap["p0"]
	if !ok {
		panic("did not find parameter p0 for augmented function")
	} else if !val.IsPositive() {
		return sdkerrors.Wrap(ErrArgumentMustBePositive, "p0")
	}

	// Augmented exception 3: theta must be from 0 to 1 (excluding 1), since
	// theta is a percentage (0% to 100%) and we cannot charge 100% fee
	val, ok = paramsMap["theta"]
	if !ok {
		panic("did not find parameter theta for augmented function")
	} else if val.LT(sdk.ZeroDec()) || val.GTE(sdk.OneDec()) {
		return sdkerrors.Wrapf(ErrArgumentMustBeBetween,
			"argument theta must be between %s and %s",
			sdk.ZeroDec().String(), sdk.OneDec().String())
	}

	// Augmented exception 4: kappa > 0, since negative exponents are not
	// allowed, but also to avoid division by zero
	val, ok = paramsMap["kappa"]
	if !ok {
		panic("did not find parameter kappa for augmented function")
	} else if !val.TruncateDec().Equal(val) {
		return sdkerrors.Wrap(ErrArgumentMustBeInteger, "kappa")
	} else if !val.IsPositive() {
		return sdkerrors.Wrap(ErrArgumentMustBePositive, "kappa")
	}

	return nil
}

//type Bond struct {
//	Token                        string         `json:"token" yaml:"token"`
//	Name                         string         `json:"name" yaml:"name"`
//	Description                  string         `json:"description" yaml:"description"`
//	CreatorDid                   exported.Did   `json:"creator_did" yaml:"creator_did"`
//	ControllerDid                exported.Did   `json:"controller_did" yaml:"controller_did"`
//	FunctionType                 string         `json:"function_type" yaml:"function_type"`
//	FunctionParameters           FunctionParams `json:"function_parameters" yaml:"function_parameters"`
//	ReserveTokens                []string       `json:"reserve_tokens" yaml:"reserve_tokens"`
//	TxFeePercentage              sdk.Dec        `json:"tx_fee_percentage" yaml:"tx_fee_percentage"`
//	ExitFeePercentage            sdk.Dec        `json:"exit_fee_percentage" yaml:"exit_fee_percentage"`
//	FeeAddress                   sdk.AccAddress `json:"fee_address" yaml:"fee_address"`
//	MaxSupply                    sdk.Coin       `json:"max_supply" yaml:"max_supply"`
//	OrderQuantityLimits          sdk.Coins      `json:"order_quantity_limits" yaml:"order_quantity_limits"`
//	SanityRate                   sdk.Dec        `json:"sanity_rate" yaml:"sanity_rate"`
//	SanityMarginPercentage       sdk.Dec        `json:"sanity_margin_percentage" yaml:"sanity_margin_percentage"`
//	CurrentSupply                sdk.Coin       `json:"current_supply" yaml:"current_supply"`
//	CurrentReserve         		 sdk.Coins      `json:"current_reserve" yaml:"current_reserve"`
//	CurrentOutcomePaymentReserve sdk.Coins      `json:"current_outcome_payment_reserve" yaml:"current_outcome_payment_reserve"`
//	AllowSells                   bool           `json:"allow_sells" yaml:"allow_sells"`
//	AlphaBond                    bool           `json:"alpha_bond" yaml:"alpha_bond"`
//	BatchBlocks                  sdk.Uint       `json:"batch_blocks" yaml:"batch_blocks"`
//	OutcomePayment               sdk.Int        `json:"outcome_payment" yaml:"outcome_payment"`
//	State                        BondState      `json:"state" yaml:"state"`
//	BondDid                      exported.Did   `json:"bond_did" yaml:"bond_did"`
//}

func NewBond(token, name, description string, creatorDid, controllerDid did.Did,
	functionType string, functionParameters FunctionParams, reserveTokens []string,
	txFeePercentage, exitFeePercentage sdk.Dec, feeAddress sdk.AccAddress,
	maxSupply sdk.Coin, orderQuantityLimits sdk.Coins, sanityRate,
	sanityMarginPercentage sdk.Dec, allowSells, alphaBond bool, batchBlocks sdk.Uint,
	outcomePayment sdk.Int, state BondState, bondDid did.Did) Bond {

	// Ensure tokens and coins are sorted
	sort.Strings(reserveTokens)
	orderQuantityLimits = orderQuantityLimits.Sort()

	return Bond{
		Token:                        token,
		Name:                         name,
		Description:                  description,
		CreatorDid:                   creatorDid,
		ControllerDid:                controllerDid,
		FunctionType:                 functionType,
		FunctionParameters:           functionParameters,
		ReserveTokens:                reserveTokens,
		TxFeePercentage:              txFeePercentage,
		ExitFeePercentage:            exitFeePercentage,
		FeeAddress:                   feeAddress.String(),
		MaxSupply:                    maxSupply,
		OrderQuantityLimits:          orderQuantityLimits,
		SanityRate:                   sanityRate,
		SanityMarginPercentage:       sanityMarginPercentage,
		CurrentSupply:                sdk.NewCoin(token, sdk.ZeroInt()),
		CurrentReserve:               nil,
		CurrentOutcomePaymentReserve: nil,
		AllowSells:                   allowSells,
		AlphaBond:                    alphaBond,
		BatchBlocks:                  batchBlocks,
		OutcomePayment:               outcomePayment,
		State:                        state.String(),
		BondDid:                      bondDid,
	}
}

func (bond Bond) GetNewReserveCoins(amount sdk.Int) (coins sdk.Coins) {
	coins = sdk.Coins{}
	for _, r := range bond.ReserveTokens {
		coins = coins.Add(sdk.NewCoin(r, amount))
	}
	return coins
}

func (bond Bond) GetNewReserveDecCoins(amount sdk.Dec) (coins sdk.DecCoins) {
	coins = sdk.DecCoins{}
	for _, r := range bond.ReserveTokens {
		coins = coins.Add(sdk.NewDecCoinFromDec(r, amount))
	}
	return coins
}

func (bond Bond) GetPricesAtSupply(supply sdk.Int) (result sdk.DecCoins, err error) {
	if supply.IsNegative() {
		panic(fmt.Sprintf("negative supply for bond %s", bond.Token))
	}

	args := bond.FunctionParameters.AsMap()
	x := supply.ToDec()
	switch bond.FunctionType {
	case PowerFunction:
		m := args["m"]
		n := args["n"]
		c := args["c"]
		XtoN, err := ApproxPower(x, n)
		if err != nil {
			return nil, err
		}
		result = bond.GetNewReserveDecCoins(XtoN.Mul(m).Add(c))
	case SigmoidFunction:
		a := args["a"]
		b := args["b"]
		c := args["c"]
		temp1 := x.Sub(b)
		temp2 := temp1.Mul(temp1).Add(c)
		temp3, err := ApproxRoot(temp2, sdk.NewDec(2))
		if err != nil {
			return nil, err
		}
		result = bond.GetNewReserveDecCoins(
			a.Mul(temp1.Quo(temp3).Add(sdk.OneDec())))
	case AugmentedFunction:
		// Note: during the hatch phase, this function returns the hatch price
		// p0 even if the supply argument is greater than the initial supply S0
		switch bond.State {
		case HatchState.String():
			result = bond.GetNewReserveDecCoins(args["p0"])
		case OpenState.String():
			kappa := args["kappa"]
			res, err := Reserve(x, kappa, args["V0"])
			if err != nil {
				return nil, err
			}

			// If reserve < 1, default to zero price to avoid calculation issues
			if res.LT(sdk.OneDec()) {
				result = bond.GetNewReserveDecCoins(sdk.ZeroDec())
			} else {
				spotPriceDec, err := SpotPrice(res, kappa, args["V0"])
				if err != nil {
					return nil, err
				}
				result = bond.GetNewReserveDecCoins(spotPriceDec)
			}
		case SettleState.String():
			return nil, ErrInvalidStateForAction
		default:
			panic("unrecognized bond state")
		}
	case SwapperFunction:
		return nil, ErrFunctionNotAvailableForFunctionType
	default:
		panic("unrecognized function type")
	}

	if result.IsAnyNegative() {
		// assumes that the curve is above the x-axis and does not intersect it
		panic(fmt.Sprintf("negative price result for bond %s", bond.Token))
	}
	return result, nil
}

func (bond Bond) GetCurrentPricesPT(reserveBalances sdk.Coins) (sdk.DecCoins, error) {
	// Note: PT stands for "per token"
	switch bond.FunctionType {
	case PowerFunction:
		fallthrough
	case SigmoidFunction:
		fallthrough
	case AugmentedFunction:
		return bond.GetPricesAtSupply(bond.CurrentSupply.Amount)
	case SwapperFunction:
		return bond.GetPricesToMint(sdk.OneInt(), reserveBalances)
	default:
		panic("unrecognized function type")
	}
}

func (bond Bond) ReserveAtSupply(supply sdk.Int) (result sdk.Dec, err error) {
	if supply.IsNegative() {
		panic(fmt.Sprintf("negative supply for bond %s", bond.Token))
	}

	args := bond.FunctionParameters.AsMap()
	x := supply.ToDec()
	switch bond.FunctionType {
	case PowerFunction:
		m := args["m"]
		n := args["n"]
		c := args["c"]
		temp1, err := ApproxPower(x, n.Add(sdk.OneDec()))
		if err != nil {
			return sdk.Dec{}, err
		}
		temp2 := temp1.Mul(m).Quo(n.Add(sdk.OneDec()))
		temp3 := x.Mul(c)
		result = temp2.Add(temp3)
	case SigmoidFunction:
		a := args["a"]
		b := args["b"]
		c := args["c"]
		temp1 := x.Sub(b)
		temp2 := temp1.Mul(temp1).Add(c)
		temp3, err := ApproxRoot(temp2, sdk.NewDec(2))
		if err != nil {
			panic(err) // mathematical problem // TODO: consider returning err
		}
		temp5 := a.Mul(temp3.Add(x))
		temp6, err := ApproxRoot(b.Mul(b).Add(c), sdk.NewDec(2))
		if err != nil {
			panic(err) // mathematical problem // TODO: consider returning err
		}
		constant := a.Mul(temp6)
		result = temp5.Sub(constant)
	case AugmentedFunction:
		kappa := args["kappa"]
		V0 := args["V0"]
		result, err = Reserve(x, kappa, V0)
		if err != nil {
			return sdk.Dec{}, err
		}
	case SwapperFunction:
		panic("invalid function for function type")
	default:
		panic("unrecognized function type")
	}

	if result.IsNegative() {
		// For vanilla bonding curves, we assume that the curve does not
		// intersect the x-axis and is greater than zero throughout
		panic(fmt.Sprintf("negative reserve result for bond %s", bond.Token))
	}
	return result, nil
}

func (bond Bond) GetReserveDeltaForLiquidityDelta(mintOrBurn sdk.Int, reserveBalances sdk.Coins) sdk.DecCoins {
	if mintOrBurn.IsNegative() {
		panic(fmt.Sprintf("negative liquidity delta for bond %s", bond.Token))
	} else if reserveBalances.IsAnyNegative() {
		panic(fmt.Sprintf("negative reserve balance for bond %s", bond.Token))
	}

	switch bond.FunctionType {
	case PowerFunction:
		fallthrough
	case SigmoidFunction:
		fallthrough
	case AugmentedFunction:
		panic("invalid function for function type")
	case SwapperFunction:
		resToken1 := bond.ReserveTokens[0]
		resToken2 := bond.ReserveTokens[1]
		resBalance1 := reserveBalances.AmountOf(resToken1).ToDec()
		resBalance2 := reserveBalances.AmountOf(resToken2).ToDec()

		// Using Uniswap formulae: x' = (1+-α)x = x +- Δx, where α = Δx/x
		// Where x is any of the two reserve balances or the current supply
		// and x' is any of the updated reserve balances or the updated supply
		// By making Δx subject of the formula: Δx = αx
		alpha := mintOrBurn.ToDec().Quo(bond.CurrentSupply.Amount.ToDec())

		result := sdk.DecCoins{
			sdk.NewDecCoinFromDec(resToken1, alpha.Mul(resBalance1)),
			sdk.NewDecCoinFromDec(resToken2, alpha.Mul(resBalance2)),
		}
		if result.IsAnyNegative() {
			panic(fmt.Sprintf("negative reserve delta result for bond %s", bond.Token))
		}
		return result
	default:
		panic("unrecognized function type")
	}
}

func (bond Bond) GetPricesToMint(mint sdk.Int, reserveBalances sdk.Coins) (sdk.DecCoins, error) {
	if mint.IsNegative() {
		panic(fmt.Sprintf("negative mint amount for bond %s", bond.Token))
	} else if reserveBalances.IsAnyNegative() {
		panic(fmt.Sprintf("negative reserve balance for bond %s", bond.Token))
	}

	// If hatch phase for augmented function, use fixed p0 price
	if bond.FunctionType == AugmentedFunction && bond.State == HatchState.String() {
		args := bond.FunctionParameters.AsMap()
		if bond.State == HatchState.String(){
			price := args["p0"].Mul(mint.ToDec())
			return bond.GetNewReserveDecCoins(price), nil
		}
	}

	switch bond.FunctionType {
	case PowerFunction:
		fallthrough
	case SigmoidFunction:
		fallthrough
	case AugmentedFunction:
		reserve, err := bond.ReserveAtSupply(bond.CurrentSupply.Amount.Add(mint))
		if err != nil {
			return nil, err
		}

		var priceToMint sdk.Dec
		if reserveBalances.Empty() {
			priceToMint = reserve
		} else {
			// Reserve balances should all be equal given that we are always
			// applying the same additions/subtractions to all reserve balances.
			// Thus we can pick the first reserve balance as the global balance.
			commonReserveBalance := reserveBalances[0].Amount.ToDec()
			priceToMint = reserve.Sub(commonReserveBalance)
		}
		if priceToMint.IsNegative() {
			// Negative priceToMint means that the previous buyer overpaid
			// to the point that the price for this buyer is covered. However,
			// we still charge this buyer at least one token.
			priceToMint = sdk.OneDec()
		}
		return bond.GetNewReserveDecCoins(priceToMint), nil
	case SwapperFunction:
		if bond.CurrentSupply.Amount.IsZero() {
			return nil, ErrFunctionRequiresNonZeroCurrentSupply
		}
		return bond.GetReserveDeltaForLiquidityDelta(mint, reserveBalances), nil
	default:
		panic("unrecognized function type")
	}
	// Note: fees have to be added to these prices to get actual prices
}

func (bond Bond) GetReturnsForBurn(burn sdk.Int, reserveBalances sdk.Coins) (sdk.DecCoins, error) {
	if burn.IsNegative() {
		panic(fmt.Sprintf("negative burn amount for bond %s", bond.Token))
	} else if reserveBalances.IsAnyNegative() {
		panic(fmt.Sprintf("negative reserve balance for bond %s", bond.Token))
	}

	switch bond.FunctionType {
	case PowerFunction:
		fallthrough
	case SigmoidFunction:
		fallthrough
	case AugmentedFunction:
		reserve, err := bond.ReserveAtSupply(bond.CurrentSupply.Amount.Sub(burn))
		if err != nil {
			return nil, err
		}

		var reserveBalance sdk.Dec
		if reserveBalances.Empty() {
			reserveBalance = sdk.ZeroDec()
		} else {
			// Reserve balances should all be equal given that we are always
			// applying the same additions/subtractions to all reserve balances.
			// Thus we can pick the first reserve balance as the global balance.
			reserveBalance = reserveBalances[0].Amount.ToDec()
		}

		if reserve.GT(reserveBalance) {
			panic("not enough reserve available for burn")
		} else {
			returnForBurn := reserveBalance.Sub(reserve)
			return bond.GetNewReserveDecCoins(returnForBurn), nil
			// TODO: investigate possibility of negative returnForBurn
		}
	case SwapperFunction:
		return bond.GetReserveDeltaForLiquidityDelta(burn, reserveBalances), nil
	default:
		panic("unrecognized function type")
	}
	// Note: fees have to be deducted from these returns to get actual returns
}

func (bond Bond) GetReturnsForSwap(from sdk.Coin, toToken string, reserveBalances sdk.Coins) (returns sdk.Coins, txFee sdk.Coin, err error) {
	if from.IsNegative() {
		panic(fmt.Sprintf("negative from amount for bond %s", bond.Token))
	} else if reserveBalances.IsAnyNegative() {
		panic(fmt.Sprintf("negative reserve balance for bond %s", bond.Token))
	}

	switch bond.FunctionType {
	case PowerFunction:
		fallthrough
	case SigmoidFunction:
		fallthrough
	case AugmentedFunction:
		return nil, sdk.Coin{}, ErrFunctionNotAvailableForFunctionType
	case SwapperFunction:
		// Check that from and to are reserve tokens
		if from.Denom != bond.ReserveTokens[0] && from.Denom != bond.ReserveTokens[1] {
			return nil, sdk.Coin{}, sdkerrors.Wrap(ErrTokenIsNotAValidReserveToken, from.Denom)
		} else if toToken != bond.ReserveTokens[0] && toToken != bond.ReserveTokens[1] {
			return nil, sdk.Coin{}, sdkerrors.Wrap(ErrTokenIsNotAValidReserveToken, toToken)
		}

		inAmt := from.Amount
		inRes := reserveBalances.AmountOf(from.Denom)
		outRes := reserveBalances.AmountOf(toToken)

		// Calculate fee to get the adjusted input amount
		txFee = bond.GetTxFee(sdk.NewDecCoinFromCoin(from))
		inAmt = inAmt.Sub(txFee.Amount) // adjusted input

		// Check that at least 1 token is going in
		if inAmt.IsZero() {
			return nil, sdk.Coin{}, ErrSwapAmountTooSmallToGiveAnyReturn
		}

		// Calculate output amount using Uniswap formula: Δy = (Δx*y)/(x+Δx)
		outAmt := inAmt.Mul(outRes).Quo(inRes.Add(inAmt))

		// Check that not giving out all of the available outRes or nothing at all
		if outAmt.Equal(outRes) {
			return nil, sdk.Coin{}, ErrSwapAmountCausesReserveDepletion
		} else if outAmt.IsZero() {
			return nil, sdk.Coin{}, ErrSwapAmountCausesReserveDepletion
		} else if outAmt.IsNegative() {
			panic(fmt.Sprintf("negative return for swap result for bond %s", bond.Token))
		}

		return sdk.Coins{sdk.NewCoin(toToken, outAmt)}, txFee, nil
	default:
		panic("unrecognized function type")
	}
}

func (bond Bond) GetFee(reserveAmount sdk.DecCoin, percentage sdk.Dec) sdk.Coin {
	feeAmount := percentage.QuoInt64(100).Mul(reserveAmount.Amount)
	return RoundFee(sdk.NewDecCoinFromDec(reserveAmount.Denom, feeAmount))
}

func (bond Bond) GetTxFee(reserveAmount sdk.DecCoin) sdk.Coin {
	return bond.GetFee(reserveAmount, bond.TxFeePercentage)
}

func (bond Bond) GetExitFee(reserveAmount sdk.DecCoin) sdk.Coin {
	return bond.GetFee(reserveAmount, bond.ExitFeePercentage)
}

func (bond Bond) GetFees(reserveAmounts sdk.DecCoins, percentage sdk.Dec) (fees sdk.Coins) {
	for _, r := range reserveAmounts {
		fees = fees.Add(bond.GetFee(r, percentage))
	}
	return fees
}

//noinspection GoNilness
func (bond Bond) GetTxFees(reserveAmounts sdk.DecCoins) (fees sdk.Coins) {
	return bond.GetFees(reserveAmounts, bond.TxFeePercentage)
}

//noinspection GoNilness
func (bond Bond) GetExitFees(reserveAmounts sdk.DecCoins) (fees sdk.Coins) {
	return bond.GetFees(reserveAmounts, bond.ExitFeePercentage)
}

func (bond Bond) ReserveDenomsEqualTo(coins sdk.Coins) bool {
	if len(bond.ReserveTokens) != len(coins) {
		return false
	}

	for _, d := range bond.ReserveTokens {
		if coins.AmountOf(d).IsZero() {
			return false
		}
	}

	return true
}

func (bond Bond) AnyOrderQuantityLimitsExceeded(amounts sdk.Coins) bool {
	return amounts.IsAnyGT(bond.OrderQuantityLimits)
}

func (bond Bond) ReservesViolateSanityRate(newReserves sdk.Coins) bool {

	if bond.SanityRate.IsZero() {
		return false
	}

	// Get new rate from new balances
	resToken1 := bond.ReserveTokens[0]
	resToken2 := bond.ReserveTokens[1]
	resBalance1 := newReserves.AmountOf(resToken1).ToDec()
	resBalance2 := newReserves.AmountOf(resToken2).ToDec()
	exchangeRate := resBalance1.Quo(resBalance2)

	// Get max and min acceptable rates
	sanityMarginDecimal := bond.SanityMarginPercentage.Quo(sdk.NewDec(100))
	upperPercentage := sdk.OneDec().Add(sanityMarginDecimal)
	lowerPercentage := sdk.OneDec().Sub(sanityMarginDecimal)
	maxRate := bond.SanityRate.Mul(upperPercentage)
	minRate := bond.SanityRate.Mul(lowerPercentage)

	// If min rate is negative, change to zero
	if minRate.IsNegative() {
		minRate = sdk.ZeroDec()
	}

	return exchangeRate.LT(minRate) || exchangeRate.GT(maxRate)
}

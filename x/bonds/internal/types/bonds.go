package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"sort"
)

const (
	PowerFunction    = "power_function"
	SigmoidFunction  = "sigmoid_function"
	SwapperFunction  = "swapper_function"
	DoNotModifyField = "[do-not-modify]"

	AnyNumberOfReserveTokens = -1
)

var (
	RequiredParamsForFunctionType = map[string][]string{
		PowerFunction:   {"m", "n", "c"},
		SigmoidFunction: {"a", "b", "c"},
		SwapperFunction: nil,
	}

	NoOfReserveTokensForFunctionType = map[string]int{
		PowerFunction:   AnyNumberOfReserveTokens,
		SigmoidFunction: AnyNumberOfReserveTokens,
		SwapperFunction: 2,
	}
)

type FunctionParam struct {
	Param string  `json:"param" yaml:"param"`
	Value sdk.Int `json:"value" yaml:"value"`
}

func NewFunctionParam(param string, value sdk.Int) FunctionParam {
	return FunctionParam{
		Param: param,
		Value: value,
	}
}

type FunctionParams []FunctionParam

func (fps FunctionParams) String() (result string) {
	result = "{"
	for _, fp := range fps {
		result += fp.Param + ":" + fp.Value.String() + ","
	}
	if len(fps) > 0 {
		// Remove last comma
		result = result[:len(result)-1]
	}
	return result + "}"
}

func (fps FunctionParams) AsMap() (paramsMap map[string]sdk.Int) {
	paramsMap = make(map[string]sdk.Int)
	for _, fp := range fps {
		paramsMap[fp.Param] = fp.Value
	}
	return paramsMap
}

type Bond struct {
	Token                  string         `json:"token" yaml:"token"`
	Name                   string         `json:"name" yaml:"name"`
	Description            string         `json:"description" yaml:"description"`
	CreatorDid             ixo.Did        `json:"creator_did" yaml:"creator_did"`
	FunctionType           string         `json:"function_type" yaml:"function_type"`
	FunctionParameters     FunctionParams `json:"function_parameters" yaml:"function_parameters"`
	ReserveTokens          []string       `json:"reserve_tokens" yaml:"reserve_tokens"`
	ReserveAddress         sdk.AccAddress `json:"reserve_address" yaml:"reserve_address"`
	TxFeePercentage        sdk.Dec        `json:"tx_fee_percentage" yaml:"tx_fee_percentage"`
	ExitFeePercentage      sdk.Dec        `json:"exit_fee_percentage" yaml:"exit_fee_percentage"`
	FeeAddress             sdk.AccAddress `json:"fee_address" yaml:"fee_address"`
	MaxSupply              sdk.Coin       `json:"max_supply" yaml:"max_supply"`
	OrderQuantityLimits    sdk.Coins      `json:"order_quantity_limits" yaml:"order_quantity_limits"`
	SanityRate             sdk.Dec        `json:"sanity_rate" yaml:"sanity_rate"`
	SanityMarginPercentage sdk.Dec        `json:"sanity_margin_percentage" yaml:"sanity_margin_percentage"`
	CurrentSupply          sdk.Coin       `json:"current_supply" yaml:"current_supply"`
	AllowSells             string         `json:"allow_sells" yaml:"allow_sells"`
	BatchBlocks            sdk.Uint       `json:"batch_blocks" yaml:"batch_blocks"`
	BondDid                ixo.Did        `json:"bond_did" yaml:"bond_did"`
	PubKey                 string         `json:"pubKey" yaml:"pubKey"`
}

func NewBond(token, name, description string, creatorDid ixo.Did,
	functionType string, functionParameters FunctionParams,
	reserveTokens []string, reserveAdddress sdk.AccAddress,
	txFeePercentage, exitFeePercentage sdk.Dec, feeAddress sdk.AccAddress,
	maxSupply sdk.Coin, orderQuantityLimits sdk.Coins, sanityRate,
	sanityMarginPercentage sdk.Dec, allowSells string,
	batchBlocks sdk.Uint, bondDid ixo.Did, pubKey string) Bond {

	// Ensure tokens and coins are sorted
	sort.Strings(reserveTokens)
	orderQuantityLimits = orderQuantityLimits.Sort()

	return Bond{
		Token:                  token,
		Name:                   name,
		Description:            description,
		CreatorDid:             creatorDid,
		FunctionType:           functionType,
		FunctionParameters:     functionParameters,
		ReserveTokens:          reserveTokens,
		ReserveAddress:         reserveAdddress,
		TxFeePercentage:        txFeePercentage,
		ExitFeePercentage:      exitFeePercentage,
		FeeAddress:             feeAddress,
		MaxSupply:              maxSupply,
		OrderQuantityLimits:    orderQuantityLimits,
		SanityRate:             sanityRate,
		SanityMarginPercentage: sanityMarginPercentage,
		CurrentSupply:          sdk.NewCoin(token, sdk.ZeroInt()),
		AllowSells:             allowSells,
		BatchBlocks:            batchBlocks,
		BondDid:                bondDid,
		PubKey:                 pubKey,
	}
}

//noinspection GoNilness
func (bond Bond) GetNewReserveDecCoins(amount sdk.Dec) (coins sdk.DecCoins) {
	for _, r := range bond.ReserveTokens {
		coins = coins.Add(sdk.DecCoins{sdk.NewDecCoinFromDec(r, amount)})
	}
	return coins
}

func (bond Bond) GetPricesAtSupply(supply sdk.Int) (result sdk.DecCoins, err sdk.Error) {
	if supply.IsNegative() {
		panic(fmt.Sprintf("negative supply for bond %s", bond))
	}

	args := bond.FunctionParameters.AsMap()
	x := supply
	switch bond.FunctionType {
	case PowerFunction:
		// TODO: should try using simpler approach, especially
		//       if function params are changed to decimals
		m := args["m"]
		n64 := args["n"].Int64()
		c := args["c"]
		temp := x
		for i := n64; i > 1; i-- {
			temp = temp.Mul(x)
		}
		result = bond.GetNewReserveDecCoins(sdk.NewDecFromInt(temp.Mul(m).Add(c)))
	case SigmoidFunction:
		aDec := sdk.NewDecFromInt(args["a"])
		b := args["b"]
		c := args["c"]
		temp1 := x.Sub(b)
		temp2 := temp1.Mul(temp1).Add(c)
		temp3 := SquareRootInt(temp2)
		result = bond.GetNewReserveDecCoins(aDec.Mul(sdk.NewDecFromInt(temp1).Quo(temp3).Add(sdk.OneDec())))
	case SwapperFunction:
		return nil, ErrFunctionNotAvailableForFunctionType(DefaultCodespace)
	default:
		panic("unrecognized function type")
	}

	if result.IsAnyNegative() {
		// assumes that the curve is above the x-axis and does not intersect it
		panic(fmt.Sprintf("negative price result for bond %s", bond))
	}
	return result, nil
}

func (bond Bond) GetCurrentPricesPT(reserveBalances sdk.Coins) (sdk.DecCoins, sdk.Error) {
	// Note: PT stands for "per token"
	switch bond.FunctionType {
	case PowerFunction:
		fallthrough
	case SigmoidFunction:
		return bond.GetPricesAtSupply(bond.CurrentSupply.Amount)
	case SwapperFunction:
		return bond.GetPricesToMint(sdk.OneInt(), reserveBalances)
	default:
		panic("unrecognized function type")
	}
}

func (bond Bond) CurveIntegral(supply sdk.Int) (result sdk.Dec) {
	if supply.IsNegative() {
		panic(fmt.Sprintf("negative supply for bond %s", bond))
	}

	args := bond.FunctionParameters.AsMap()
	x, xDec := supply, sdk.NewDecFromInt(supply)
	switch bond.FunctionType {
	case PowerFunction:
		// TODO: should try using simpler approach, especially
		//       if function params are changed to decimals
		m := args["m"]
		n, n64 := args["n"], args["n"].Int64()
		c := args["c"]
		temp1 := x
		for i := n64 + 1; i > 1; i-- {
			temp1 = temp1.Mul(x)
		}
		temp2 := sdk.NewDecFromInt(temp1.Mul(m)).Quo(sdk.NewDecFromInt(n.Add(sdk.OneInt())))
		temp3 := sdk.NewDecFromInt(x.Mul(c))
		result = temp2.Add(temp3)
	case SigmoidFunction:
		aDec := sdk.NewDecFromInt(args["a"])
		b, bDec := args["b"], sdk.NewDecFromInt(args["b"])
		c, cDec := args["c"], sdk.NewDecFromInt(args["c"])
		temp1 := x.Sub(b)
		temp2 := temp1.Mul(temp1).Add(c)
		temp3 := SquareRootInt(temp2)
		temp5 := aDec.Mul(temp3.Add(xDec))
		constant := aDec.Mul(SquareRootDec(bDec.Mul(bDec).Add(cDec)))
		result = temp5.Sub(constant)
	case SwapperFunction:
		panic("invalid function for function type")
	default:
		panic("unrecognized function type")
	}

	if result.IsNegative() {
		// assumes that the curve is above the x-axis and does not intersect it
		panic(fmt.Sprintf("negative integral result for bond %s", bond))
	}
	return result
}

func (bond Bond) GetReserveDeltaForLiquidityDelta(mintOrBurn sdk.Int, reserveBalances sdk.Coins) sdk.DecCoins {
	if mintOrBurn.IsNegative() {
		panic(fmt.Sprintf("negative liquidity delta for bond %s", bond))
	} else if reserveBalances.IsAnyNegative() {
		panic(fmt.Sprintf("negative reserve balance for bond %s", bond))
	}

	switch bond.FunctionType {
	case PowerFunction:
		fallthrough
	case SigmoidFunction:
		panic("invalid function for function type")
	case SwapperFunction:
		resToken1 := bond.ReserveTokens[0]
		resToken2 := bond.ReserveTokens[1]
		resBalance1 := sdk.NewDecFromInt(reserveBalances.AmountOf(resToken1))
		resBalance2 := sdk.NewDecFromInt(reserveBalances.AmountOf(resToken2))
		mintOrBurnDec := sdk.NewDecFromInt(mintOrBurn)

		// Using Uniswap formulae: x' = (1+-α)x = x +- Δx, where α = Δx/x
		// Where x is any of the two reserve balances or the current supply
		// and x' is any of the updated reserve balances or the updated supply
		// By making Δx subject of the formula: Δx = αx
		alpha := mintOrBurnDec.Quo(sdk.NewDecFromInt(bond.CurrentSupply.Amount))

		result := sdk.DecCoins{
			sdk.NewDecCoinFromDec(resToken1, alpha.Mul(resBalance1)),
			sdk.NewDecCoinFromDec(resToken2, alpha.Mul(resBalance2)),
		}
		if result.IsAnyNegative() {
			panic(fmt.Sprintf("negative reserve delta result for bond %s", bond))
		}
		return result
	default:
		panic("unrecognized function type")
	}
}

func (bond Bond) GetPricesToMint(mint sdk.Int, reserveBalances sdk.Coins) (sdk.DecCoins, sdk.Error) {
	if mint.IsNegative() {
		panic(fmt.Sprintf("negative mint amount for bond %s", bond))
	} else if reserveBalances.IsAnyNegative() {
		panic(fmt.Sprintf("negative reserve balance for bond %s", bond))
	}

	switch bond.FunctionType {
	case PowerFunction:
		fallthrough
	case SigmoidFunction:
		var priceToMint sdk.Dec
		result := bond.CurveIntegral(bond.CurrentSupply.Amount.Add(mint))
		if reserveBalances.Empty() {
			priceToMint = result
		} else {
			// Reserve balances should all be equal given that we are always
			// applying the same additions/subtractions to all reserve balances
			commonReserveBalance := sdk.NewDecFromInt(reserveBalances[0].Amount)
			priceToMint = result.Sub(commonReserveBalance)
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
			return nil, ErrFunctionRequiresNonZeroCurrentSupply(DefaultCodespace)
		}
		return bond.GetReserveDeltaForLiquidityDelta(mint, reserveBalances), nil
	default:
		panic("unrecognized function type")
	}
	// Note: fees have to be added to these prices to get actual prices
}

func (bond Bond) GetReturnsForBurn(burn sdk.Int, reserveBalances sdk.Coins) sdk.DecCoins {
	if burn.IsNegative() {
		panic(fmt.Sprintf("negative burn amount for bond %s", bond))
	} else if reserveBalances.IsAnyNegative() {
		panic(fmt.Sprintf("negative reserve balance for bond %s", bond))
	}

	switch bond.FunctionType {
	case PowerFunction:
		fallthrough
	case SigmoidFunction:
		var returnForBurn sdk.Dec
		result := bond.CurveIntegral(bond.CurrentSupply.Amount.Sub(burn))
		if reserveBalances.Empty() {
			panic("no reserve available for burn")
		} else {
			// Reserve balances should all be equal given that we are always
			// applying the same additions/subtractions to all reserve balances
			commonReserveBalance := sdk.NewDecFromInt(reserveBalances[0].Amount)
			returnForBurn = commonReserveBalance.Sub(result)
		}
		// TODO: investigate possibility of negative returnForBurn
		return bond.GetNewReserveDecCoins(returnForBurn)
	case SwapperFunction:
		return bond.GetReserveDeltaForLiquidityDelta(burn, reserveBalances)
	default:
		panic("unrecognized function type")
	}
	// Note: fees have to be deducted from these returns to get actual returns
}

func (bond Bond) GetReturnsForSwap(from sdk.Coin, toToken string, reserveBalances sdk.Coins) (returns sdk.Coins, txFee sdk.Coin, err sdk.Error) {
	if from.IsNegative() {
		panic(fmt.Sprintf("negative from amount for bond %s", bond))
	} else if reserveBalances.IsAnyNegative() {
		panic(fmt.Sprintf("negative reserve balance for bond %s", bond))
	}

	switch bond.FunctionType {
	case PowerFunction:
		fallthrough
	case SigmoidFunction:
		return nil, sdk.Coin{}, ErrFunctionNotAvailableForFunctionType(DefaultCodespace)
	case SwapperFunction:
		// Check that from and to are reserve tokens
		if from.Denom != bond.ReserveTokens[0] && from.Denom != bond.ReserveTokens[1] {
			return nil, sdk.Coin{}, ErrTokenIsNotAValidReserveToken(DefaultCodespace, from.Denom)
		} else if toToken != bond.ReserveTokens[0] && toToken != bond.ReserveTokens[1] {
			return nil, sdk.Coin{}, ErrTokenIsNotAValidReserveToken(DefaultCodespace, toToken)
		}

		inAmt := from.Amount
		inRes := reserveBalances.AmountOf(from.Denom)
		outRes := reserveBalances.AmountOf(toToken)

		// Calculate fee to get the adjusted input amount
		txFee = bond.GetTxFee(sdk.NewDecCoinFromCoin(from))
		inAmt = inAmt.Sub(txFee.Amount) // adjusted input

		// Check that at least 1 token is going in
		if inAmt.IsZero() {
			return nil, sdk.Coin{}, ErrSwapAmountTooSmallToGiveAnyReturn(DefaultCodespace, from.Denom, toToken)
		}

		// Calculate output amount using Uniswap formula: Δy = (Δx*y)/(x+Δx)
		outAmt := inAmt.Mul(outRes).Quo(inRes.Add(inAmt))

		// Check that not giving out all of the available outRes or nothing at all
		if outAmt.Equal(outRes) {
			return nil, sdk.Coin{}, ErrSwapAmountCausesReserveDepletion(DefaultCodespace, from.Denom, toToken)
		} else if outAmt.IsZero() {
			return nil, sdk.Coin{}, ErrSwapAmountTooSmallToGiveAnyReturn(DefaultCodespace, from.Denom, toToken)
		} else if outAmt.IsNegative() {
			panic(fmt.Sprintf("negative return for swap result for bond %s", bond))
		}

		return sdk.Coins{sdk.NewCoin(toToken, outAmt)}, txFee, nil
	default:
		panic("unrecognized function type")
	}
}

func (bond Bond) GetTxFee(reserveAmount sdk.DecCoin) sdk.Coin {
	feeAmount := bond.TxFeePercentage.QuoInt64(100).Mul(reserveAmount.Amount)
	return RoundFee(sdk.NewDecCoinFromDec(reserveAmount.Denom, feeAmount))
}

func (bond Bond) GetExitFee(reserveAmount sdk.DecCoin) sdk.Coin {
	feeAmount := bond.ExitFeePercentage.QuoInt64(100).Mul(reserveAmount.Amount)
	return RoundFee(sdk.NewDecCoinFromDec(reserveAmount.Denom, feeAmount))
}

//noinspection GoNilness
func (bond Bond) GetTxFees(reserveAmounts sdk.DecCoins) (fees sdk.Coins) {
	for _, r := range reserveAmounts {
		fees = fees.Add(sdk.Coins{bond.GetTxFee(r)})
	}
	return fees
}

//noinspection GoNilness
func (bond Bond) GetExitFees(reserveAmounts sdk.DecCoins) (fees sdk.Coins) {
	for _, r := range reserveAmounts {
		fees = fees.Add(sdk.Coins{bond.GetExitFee(r)})
	}
	return fees
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
	resBalance1 := sdk.NewDecFromInt(newReserves.AmountOf(resToken1))
	resBalance2 := sdk.NewDecFromInt(newReserves.AmountOf(resToken2))
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

package types

import (
	"strings"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	TEN18DEC = math.LegacyMustNewDecFromStr("1000000000000000000") // 1e18
)

// ApproxRoot returns an approximation of a Dec's nth root
func ApproxRoot(d math.LegacyDec, root math.LegacyDec) (guess math.LegacyDec, err error) {
	return ApproxPower(d, math.LegacyOneDec().Quo(root))
}

// ApproxPower returns an approximation of raising a Dec to a positive power
func ApproxPower(d math.LegacyDec, power math.LegacyDec) (guess math.LegacyDec, err error) {
	// Convert Dec's to Uint's
	dUint := math.NewUintFromBigInt(d.BigInt())
	powerUint := math.NewUintFromBigInt(power.BigInt())

	// Handle panics
	defer func() {
		if r := recover(); r != nil {
			err = errorsmod.Wrapf(ErrNumericOverflow, "%s", r)
		}
	}()

	// Find answer using the Uint's
	ansUint := pow(dUint, powerUint)

	// Convert back to Dec
	return math.LegacyNewDecFromBigInt(ansUint.BigInt()).Quo(TEN18DEC), nil
}

func RoundReservePrice(p sdk.DecCoin) sdk.Coin {
	// ReservePrices are rounded up so that the account gets charged more
	roundedAmount := p.Amount.Ceil().TruncateInt()
	return sdk.NewCoin(p.Denom, roundedAmount)
}

func RoundReserveReturn(r sdk.DecCoin) sdk.Coin {
	// ReserveReturns are rounded down so that the account gets less in return
	roundedAmount := r.Amount.TruncateInt()
	return sdk.NewCoin(r.Denom, roundedAmount)
}

func RoundFee(f sdk.DecCoin) sdk.Coin {
	// Fees are rounded up so that the account gets charged more
	roundedAmount := f.Amount.Ceil().TruncateInt()
	return sdk.NewCoin(f.Denom, roundedAmount)
}

// noinspection GoNilness
func RoundReservePrices(ps sdk.DecCoins) (rounded sdk.Coins) {
	for _, p := range ps {
		rounded = rounded.Add(RoundReservePrice(p))
	}
	return rounded
}

// noinspection GoNilness
func RoundReserveReturns(rs sdk.DecCoins) (rounded sdk.Coins) {
	for _, r := range rs {
		rounded = rounded.Add(RoundReserveReturn(r))
	}
	return rounded
}

func MultiplyDecCoinByInt(dc sdk.DecCoin, scale math.Int) sdk.DecCoin {
	return sdk.NewDecCoinFromDec(dc.Denom, dc.Amount.MulInt(scale))
}

// noinspection GoNilness
func MultiplyDecCoinsByInt(dcs sdk.DecCoins, scale math.Int) (scaled sdk.DecCoins) {
	for _, dc := range dcs {
		scaled = scaled.Add(MultiplyDecCoinByInt(dc, scale))
	}
	return scaled
}

func MultiplyDecCoinByDec(dc sdk.DecCoin, scale math.LegacyDec) sdk.DecCoin {
	return sdk.NewDecCoinFromDec(dc.Denom, dc.Amount.Mul(scale))
}

// noinspection GoNilness
func MultiplyDecCoinsByDec(dcs sdk.DecCoins, scale math.LegacyDec) (scaled sdk.DecCoins) {
	for _, dc := range dcs {
		scaled = scaled.Add(MultiplyDecCoinByDec(dc, scale))
	}
	return scaled
}

func DivideDecCoinByDec(dc sdk.DecCoin, scale math.LegacyDec) sdk.DecCoin {
	return sdk.NewDecCoinFromDec(dc.Denom, dc.Amount.Quo(scale))
}

// noinspection GoNilness
func DivideDecCoinsByDec(dcs sdk.DecCoins, scale math.LegacyDec) (scaled sdk.DecCoins) {
	for _, dc := range dcs {
		scaled = scaled.Add(DivideDecCoinByDec(dc, scale))
	}
	return scaled
}

func AdjustFees(fees sdk.Coins, maxFees sdk.Coins) sdk.Coins {
	// List of extra fees to deduct at the end
	extraFees := sdk.Coins{}

	// If any fee is greater than the max fee, the extra fee is discounted
	for _, f := range fees {
		maxFee := maxFees.AmountOf(f.Denom)
		if f.Amount.GT(maxFee) {
			extraFee := f.Amount.Sub(maxFee)
			extraFees = extraFees.Add(sdk.NewCoin(f.Denom, extraFee))
		}
	}
	return fees.Sub(extraFees...)
}

func AccAddressesToString(addresses []sdk.AccAddress) (result string) {
	result = "["
	for _, a := range addresses {
		result += a.String() + ","
	}
	if len(addresses) > 0 {
		result = result[:len(result)-1]
	}
	return result + "]"
}

func StringsToString(strs []string) (result string) {
	return "[" + strings.Join(strs, ",") + "]"
}

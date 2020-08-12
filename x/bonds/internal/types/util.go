package types

import (
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"math/big"
	"strings"
)

// NOTE: copied off of more recent versions of Cosmos SDK
// ApproxRoot returns an approximate estimation of a Dec's positive real nth root
// using Newton's method (where n is positive). The algorithm starts with some guess and
// computes the sequence of improved guesses until an answer converges to an
// approximate answer.  It returns `|d|.ApproxRoot() * -1` if input is negative.
func ApproxRoot(d sdk.Dec, root uint64) (guess sdk.Dec, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = errors.New("out of bounds")
			}
		}
	}()

	if d.IsNegative() {
		absRoot, err := ApproxRoot(d.MulInt64(-1), root)
		return absRoot.MulInt64(-1), err
	}

	if root == 1 || d.IsZero() || d.Equal(sdk.OneDec()) {
		return d, nil
	}

	if root == 0 {
		return sdk.OneDec(), nil
	}

	temp := big.NewInt(0)
	temp.SetUint64(root)
	rootInt := sdk.NewIntFromBigInt(temp)
	guess, delta := sdk.OneDec(), sdk.OneDec()

	for delta.Abs().GT(sdk.SmallestDec()) {
		prev := Power(guess, root-1)
		if prev.IsZero() {
			prev = sdk.SmallestDec()
		}
		delta = d.Quo(prev)
		delta = delta.Sub(guess)
		delta = delta.QuoInt(rootInt)

		guess = guess.Add(delta)
	}

	return guess, nil
}

// NOTE: copied off of more recent versions of Cosmos SDK
// Power returns a the result of raising to a positive integer power
func Power(d sdk.Dec, power uint64) sdk.Dec {
	if power == 0 {
		return sdk.OneDec()
	}
	tmp := sdk.OneDec()
	for i := power; i > 1; {
		if i%2 == 0 {
			i /= 2
		} else {
			tmp = tmp.Mul(d)
			i = (i - 1) / 2
		}
		d = d.Mul(d)
	}
	return d.Mul(tmp)
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

//noinspection GoNilness
func RoundReservePrices(ps sdk.DecCoins) (rounded sdk.Coins) {
	for _, p := range ps {
		rounded = rounded.Add(sdk.Coins{RoundReservePrice(p)})
	}
	return rounded
}

//noinspection GoNilness
func RoundReserveReturns(rs sdk.DecCoins) (rounded sdk.Coins) {
	for _, r := range rs {
		rounded = rounded.Add(sdk.Coins{RoundReserveReturn(r)})
	}
	return rounded
}

func MultiplyDecCoinByInt(dc sdk.DecCoin, scale sdk.Int) sdk.DecCoin {
	return sdk.NewDecCoinFromDec(dc.Denom, dc.Amount.MulInt(scale))
}

//noinspection GoNilness
func MultiplyDecCoinsByInt(dcs sdk.DecCoins, scale sdk.Int) (scaled sdk.DecCoins) {
	for _, dc := range dcs {
		scaled = scaled.Add(sdk.DecCoins{MultiplyDecCoinByInt(dc, scale)})
	}
	return scaled
}

func MultiplyDecCoinByDec(dc sdk.DecCoin, scale sdk.Dec) sdk.DecCoin {
	return sdk.NewDecCoinFromDec(dc.Denom, dc.Amount.Mul(scale))
}

//noinspection GoNilness
func MultiplyDecCoinsByDec(dcs sdk.DecCoins, scale sdk.Dec) (scaled sdk.DecCoins) {
	for _, dc := range dcs {
		scaled = scaled.Add(sdk.DecCoins{MultiplyDecCoinByDec(dc, scale)})
	}
	return scaled
}

func DivideDecCoinByDec(dc sdk.DecCoin, scale sdk.Dec) sdk.DecCoin {
	return sdk.NewDecCoinFromDec(dc.Denom, dc.Amount.Quo(scale))
}

//noinspection GoNilness
func DivideDecCoinsByDec(dcs sdk.DecCoins, scale sdk.Dec) (scaled sdk.DecCoins) {
	for _, dc := range dcs {
		scaled = scaled.Add(sdk.DecCoins{DivideDecCoinByDec(dc, scale)})
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
			extraFees = extraFees.Add(sdk.Coins{
				sdk.NewCoin(f.Denom, extraFee)})
		}
	}
	return fees.Sub(extraFees)
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

package ixomath

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ErrTolerance is used to define a compare function, which checks if two
// ints are within a certain error tolerance of one another,
// and (optionally) that they are rounding in the correct direction.
// ErrTolerance.Compare(a, b) returns true iff:
// * RoundingMode = RoundUp, then a <= b
// * RoundingMode = RoundDown, then a >= b
// * |a - b| <= AdditiveTolerance
// * |a - b| / min(a, b) <= MultiplicativeTolerance
//
// Each check is respectively ignored if the entry is nil.
// So AdditiveTolerance = Int{} or ZeroInt()
// MultiplicativeTolerance = Dec{}
// RoundingDir = RoundUnconstrained.
// Note that if AdditiveTolerance == 0, then this is equivalent to a standard compare.
type ErrTolerance struct {
	AdditiveTolerance       Dec
	MultiplicativeTolerance Dec
	RoundingDir             RoundingDirection
}

// Compare returns if actual is within errTolerance of expected.
// returns 0 if it is
// returns 1 if not, and expected > actual.
// returns -1 if not, and expected < actual
func (e ErrTolerance) Compare(expected Int, actual Int) int {
	diff := expected.ToLegacyDec().Sub(actual.ToLegacyDec()).Abs()

	comparisonSign := 0
	if expected.GT(actual) {
		comparisonSign = 1
	} else {
		comparisonSign = -1
	}

	// Ensure that even if expected is within tolerance of actual, we don't count it as equal if its in the wrong direction.
	// so if were supposed to round down, it must be that `expected >= actual`.
	// likewise if were supposed to round up, it must be that `expected <= actual`.
	// If neither of the above, then rounding direction does not enforce a constraint.
	if e.RoundingDir == RoundDown {
		if expected.LT(actual) {
			return -1
		}
	} else if e.RoundingDir == RoundUp {
		if expected.GT(actual) {
			return 1
		}
	}

	// Check additive tolerance equations
	if !e.AdditiveTolerance.IsNil() {
		// if no error accepted, do a direct compare.
		if e.AdditiveTolerance.IsZero() {
			if expected.Equal(actual) {
				return 0
			}
		}

		if diff.GT(e.AdditiveTolerance) {
			return comparisonSign
		}
	}
	// Check multiplicative tolerance equations
	if !e.MultiplicativeTolerance.IsNil() && !e.MultiplicativeTolerance.IsZero() {
		minValue := MinInt(expected.Abs(), actual.Abs())
		if minValue.IsZero() {
			return comparisonSign
		}

		errTerm := diff.Quo(minValue.ToLegacyDec())
		if errTerm.GT(e.MultiplicativeTolerance) {
			return comparisonSign
		}
	}

	return 0
}

// CompareBigDec validates if actual is within errTolerance of expected.
// returns 0 if it is
// returns 1 if not, and expected > actual.
// returns -1 if not, and expected < actual
func (e ErrTolerance) CompareBigDec(expected BigDec, actual BigDec) int {
	// Ensure that even if expected is within tolerance of actual, we don't count it as equal if its in the wrong direction.
	// so if were supposed to round down, it must be that `expected >= actual`.
	// likewise if were supposed to round up, it must be that `expected <= actual`.
	// If neither of the above, then rounding direction does not enforce a constraint.
	if e.RoundingDir == RoundDown {
		if expected.LT(actual) {
			return -1
		}
	} else if e.RoundingDir == RoundUp {
		if expected.GT(actual) {
			return 1
		}
	}

	diff := expected.Sub(actual).Abs()

	comparisonSign := 0
	if expected.GT(actual) {
		comparisonSign = 1
	} else if expected.LT(actual) {
		comparisonSign = -1
	}

	// Check additive tolerance equations
	if !e.AdditiveTolerance.IsNil() {
		// if no error accepted, do a direct compare.
		if e.AdditiveTolerance.IsZero() {
			if expected.Equal(actual) {
				return 0
			}
		}

		if diff.GT(BigDecFromDec(e.AdditiveTolerance)) {
			return comparisonSign
		}
	}
	// Check multiplicative tolerance equations
	if !e.MultiplicativeTolerance.IsNil() && !e.MultiplicativeTolerance.IsZero() {
		minValue := MinBigDec(expected.Abs(), actual.Abs())
		if minValue.IsZero() {
			return comparisonSign
		}

		errTerm := diff.Quo(minValue)
		// fmt.Printf("err term %v\n", errTerm)
		if errTerm.GT(BigDecFromDec(e.MultiplicativeTolerance)) {
			return comparisonSign
		}
	}

	return 0
}

// CompareDec validates if actual is within errTolerance of expected.
// returns 0 if it is
// returns 1 if not, and expected > actual.
// returns -1 if not, and expected < actual
func (e ErrTolerance) CompareDec(expected Dec, actual Dec) int {
	// Ensure that even if expected is within tolerance of actual, we don't count it as equal if its in the wrong direction.
	// so if we're supposed to round down, it must be that `expected >= actual`.
	// likewise if we're supposed to round up, it must be that `expected <= actual`.
	// If neither of the above, then rounding direction does not enforce a constraint.
	if e.RoundingDir == RoundDown {
		if expected.LT(actual) {
			return -1
		}
	} else if e.RoundingDir == RoundUp {
		if expected.GT(actual) {
			return 1
		}
	}

	diff := expected.Sub(actual).Abs()

	comparisonSign := 0
	if expected.GT(actual) {
		comparisonSign = 1
	} else if expected.LT(actual) {
		comparisonSign = -1
	}

	// Check additive tolerance equations
	if !e.AdditiveTolerance.IsNil() {
		// if no error accepted, do a direct compare.
		if e.AdditiveTolerance.IsZero() {
			if expected.Equal(actual) {
				return 0
			}
		}

		if diff.GT(e.AdditiveTolerance) {
			return comparisonSign
		}
	}
	// Check multiplicative tolerance equations
	if !e.MultiplicativeTolerance.IsNil() && !e.MultiplicativeTolerance.IsZero() {
		minValue := MinDec(expected.Abs(), actual.Abs())
		if minValue.IsZero() {
			return comparisonSign
		}

		errTerm := diff.Quo(minValue)
		if errTerm.GT(e.MultiplicativeTolerance) {
			return comparisonSign
		}
	}

	return 0
}

// EqualCoins returns true iff the two coins are equal within the ErrTolerance constraints and false otherwise.
// TODO: move error tolerance functions to a separate file.
func (e ErrTolerance) EqualCoins(expectedCoins sdk.Coins, actualCoins sdk.Coins) bool {
	if len(expectedCoins) < len(actualCoins) {
		return false
	}

	for _, expectedCoin := range expectedCoins {
		curCoinEqual := e.Compare(expectedCoin.Amount, actualCoins.AmountOf(expectedCoin.Denom))
		if curCoinEqual != 0 {
			return false
		}
	}

	return true
}

// Binary search inputs between [lowerbound, upperbound] to a monotonic increasing function f.
// We stop once f(found_input) meets the ErrTolerance constraints.
// If we perform more than maxIterations (or equivalently lowerbound = upperbound), we return an error.
func BinarySearch(f func(Int) (Int, error),
	lowerbound Int,
	upperbound Int,
	targetOutput Int,
	errTolerance ErrTolerance,
	maxIterations int,
) (Int, error) {
	var (
		curEstimate, curOutput Int
		err                    error
	)

	curIteration := 0
	for ; curIteration < maxIterations; curIteration += 1 {
		curEstimate = lowerbound.Add(upperbound).QuoRaw(2)
		curOutput, err = f(curEstimate)
		if err != nil {
			return Int{}, err
		}

		compRes := errTolerance.Compare(targetOutput, curOutput)
		if compRes < 0 {
			upperbound = curEstimate
		} else if compRes > 0 {
			lowerbound = curEstimate
		} else {
			return curEstimate, nil
		}
	}

	return Int{}, errors.New("hit maximum iterations, did not converge fast enough")
}

// SdkDec
type SdkDec[D any] interface {
	Add(SdkDec[D]) SdkDec[D]
	Quo(SdkDec[D]) SdkDec[D]
	QuoRaw(int64) SdkDec[D]
}

// BinarySearchBigDec takes as input:
// * an input range [lowerbound, upperbound]
// * an increasing function f
// * a target output x
// * max number of iterations (for gas control / handling does-not-converge cases)
//
// It binary searches on the input range, until it finds an input y s.t. f(y) meets the err tolerance constraints for how close it is to x.
// If we perform more than maxIterations (or equivalently lowerbound = upperbound), we return an error.
func BinarySearchBigDec(f func(BigDec) BigDec,
	lowerbound BigDec,
	upperbound BigDec,
	targetOutput BigDec,
	errTolerance ErrTolerance,
	maxIterations int,
) (BigDec, error) {
	var (
		curEstimate, curOutput BigDec
	)

	curIteration := 0
	for ; curIteration < maxIterations; curIteration += 1 {
		// (lowerbound + upperbound) / 2
		curEstimate = lowerbound.Add(upperbound)
		curEstimateBi := curEstimate.BigIntMut()
		curEstimateBi.Rsh(curEstimateBi, 1)
		curOutput = f(curEstimate)

		// fmt.Println("binary search, input, target output, cur output", curEstimate, targetOutput, curOutput)
		compRes := errTolerance.CompareBigDec(targetOutput, curOutput)
		if compRes < 0 {
			upperbound = curEstimate
		} else if compRes > 0 {
			lowerbound = curEstimate
		} else {
			return curEstimate, nil
		}
	}

	return BigDec{}, errors.New("hit maximum iterations, did not converge fast enough")
}

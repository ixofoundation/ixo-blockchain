package ixomath

import (
	"errors"
	"math/big"
)

// nolint:unused
var smallestDec = SmallestDec()
var tenTo18 = big.NewInt(1e18)
var tenTo36 = big.NewInt(0).Mul(tenTo18, tenTo18)
var oneBigInt = big.NewInt(1)

// Returns square root of d
// returns an error if one of the following conditions is met:
// - d is negative
// - d is too small to have a representable square root.
// This function guarantees:
// the returned root r, will be such that r^2 >= d
// This function is monotonic, i.e. if d1 >= d2, then sqrt(d1) >= sqrt(d2)
func MonotonicSqrt(d Dec) (Dec, error) {
	return MonotonicSqrtMut(d.Clone())
}

func MonotonicSqrtMut(d Dec) (Dec, error) {
	if d.IsNegative() {
		return d, errors.New("cannot take square root of negative number")
	}

	// A decimal value of d, is represented as an integer of value v = 10^18 * d.
	// We have an integer square root function, and we'd like to get the square root of d.
	// recall integer square root is floor(sqrt(x)), hence its accurate up to 1 integer.
	// we want sqrt d accurate to 18 decimal places.
	// So first we multiply our current value by 10^18, then we take the integer square root.
	// since sqrt(10^18 * v) = 10^9 * sqrt(v) = 10^18 * sqrt(d), we get the answer we want.
	//
	// We can than interpret sqrt(10^18 * v) as our resulting decimal and return it.
	// monotonicity is guaranteed by correctness of integer square root.
	dBi := d.BigIntMut()
	r := big.NewInt(0).Mul(dBi, tenTo18)
	r.Sqrt(r)
	// However this square root r is s.t. r^2 <= d. We want to flip this to be r^2 >= d.
	// To do so, we check that if r^2 < d, do r += 1. Then by correctness we will be in the case we want.
	// To compare r^2 and d, we can just compare r^2 and 10^18 * v. (recall r = 10^18 * sqrt(d), v = 10^18 * d)
	check := big.NewInt(0).Mul(r, r)
	// dBi is a copy of d, so we can modify it.
	shiftedD := dBi.Mul(dBi, tenTo18)
	if check.Cmp(shiftedD) == -1 {
		r.Add(r, oneBigInt)
	}

	dBi.Set(r)
	return d, nil
}

func MonotonicSqrtBigDec(d BigDec) (BigDec, error) {
	return MonotonicSqrtBigDecMut(d.Clone())
}

func MonotonicSqrtBigDecMut(d BigDec) (BigDec, error) {
	if d.IsNegative() {
		return d, errors.New("cannot take square root of negative number")
	}

	// A decimal value of d, is represented as an integer of value v = 10^18 * d.
	// We have an integer square root function, and we'd like to get the square root of d.
	// recall integer square root is floor(sqrt(x)), hence its accurate up to 1 integer.
	// we want sqrt d accurate to 18 decimal places.
	// So first we multiply our current value by 10^18, then we take the integer square root.
	// since sqrt(10^18 * v) = 10^9 * sqrt(v) = 10^18 * sqrt(d), we get the answer we want.
	//
	// We can than interpret sqrt(10^18 * v) as our resulting decimal and return it.
	// monotonicity is guaranteed by correctness of integer square root.
	dBi := d.BigIntMut()
	r := big.NewInt(0).Mul(dBi, tenTo36)
	r.Sqrt(r)
	// However this square root r is s.t. r^2 <= d. We want to flip this to be r^2 >= d.
	// To do so, we check that if r^2 < d, do r += 1. Then by correctness we will be in the case we want.
	// To compare r^2 and d, we can just compare r^2 and 10^18 * v. (recall r = 10^18 * sqrt(d), v = 10^18 * d)
	check := big.NewInt(0).Mul(r, r)
	// dBi is a copy of d, so we can modify it.
	shiftedD := dBi.Mul(dBi, tenTo36)
	if check.Cmp(shiftedD) == -1 {
		r.Add(r, oneBigInt)
	}
	dBi.Set(r)

	return d, nil
}

// MustMonotonicSqrt returns the output of MonotonicSqrt, panicking on error.
func MustMonotonicSqrt(d Dec) Dec {
	sqrt, err := MonotonicSqrt(d)
	if err != nil {
		panic(err)
	}
	return sqrt
}

// MustMonotonicSqrt returns the output of MonotonicSqrt, panicking on error.
func MustMonotonicSqrtBigDec(d BigDec) BigDec {
	sqrt, err := MonotonicSqrtBigDec(d)
	if err != nil {
		panic(err)
	}
	return sqrt
}

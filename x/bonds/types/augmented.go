package types

import (
	"cosmossdk.io/math"
)

// Inspired by work from BlockScience:
// https://github.com/BlockScience/cadCAD-Tutorials/tree/master/00-Reference-Mechanisms

// value function for a given state (R,S)
func Invariant(R, S math.LegacyDec, kappa math.LegacyDec) (math.LegacyDec, error) {
	SPowK, err := ApproxPower(S, kappa)
	if err != nil {
		return math.LegacyDec{}, err
	}
	return SPowK.Quo(R), nil
}

// given a value function (parameterized by kappa)
// and an invariant coefficient V0
// return Supply S as a function of reserve R
// func Supply(R math.LegacyDec, kappa math.LegacyDec, V0 math.LegacyDec) (math.LegacyDec, error) {
// 	result, err := ApproxRoot(V0.Mul(R), kappa)
// 	if err != nil {
// 		return math.LegacyDec{}, nil
// 	}
// 	return result, nil
// }

// This is the reverse of Supply(...) function
func Reserve(S math.LegacyDec, kappa math.LegacyDec, V0 math.LegacyDec) (math.LegacyDec, error) {
	SPowK, err := ApproxPower(S, kappa)
	if err != nil {
		return math.LegacyDec{}, err
	}
	return SPowK.Quo(V0), nil
}

// given a value function (parameterized by kappa)
// and an invariant coefficient V0
// return a spot price P as a function of reserve R
func SpotPrice(R math.LegacyDec, kappa math.LegacyDec, V0 math.LegacyDec) (math.LegacyDec, error) {
	temp1, err := ApproxRoot(V0, kappa)
	if err != nil {
		return math.LegacyDec{}, err
	}
	temp2, err := ApproxPower(R, kappa.Sub(math.LegacyOneDec()))
	if err != nil {
		return math.LegacyDec{}, err
	}
	temp3, err := ApproxRoot(temp2, kappa)
	if err != nil {
		return math.LegacyDec{}, err
	}
	return (kappa.Mul(temp3)).Quo(temp1), nil
}

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Inspired by work from BlockScience:
// https://github.com/BlockScience/cadCAD-Tutorials/tree/master/00-Reference-Mechanisms

// value function for a given state (R,S)
func Invariant(R, S sdk.Dec, kappa sdk.Dec) (sdk.Dec, error) {
	SPowK, err := ApproxPower(S, kappa)
	if err != nil {
		return sdk.Dec{}, err
	}
	return SPowK.Quo(R), nil
}

// given a value function (parameterized by kappa)
// and an invariant coeficient V0
// return Supply S as a function of reserve R
// func Supply(R sdk.Dec, kappa sdk.Dec, V0 sdk.Dec) (sdk.Dec, error) {
// 	result, err := ApproxRoot(V0.Mul(R), kappa)
// 	if err != nil {
// 		return sdk.Dec{}, nil
// 	}
// 	return result, nil
// }

// This is the reverse of Supply(...) function
func Reserve(S sdk.Dec, kappa sdk.Dec, V0 sdk.Dec) (sdk.Dec, error) {
	SPowK, err := ApproxPower(S, kappa)
	if err != nil {
		return sdk.Dec{}, err
	}
	return SPowK.Quo(V0), nil
}

// given a value function (parameterized by kappa)
// and an invariant coeficient V0
// return a spot price P as a function of reserve R
func SpotPrice(R sdk.Dec, kappa sdk.Dec, V0 sdk.Dec) (sdk.Dec, error) {
	temp1, err := ApproxRoot(V0, kappa)
	if err != nil {
		return sdk.Dec{}, err
	}
	temp2, err := ApproxPower(R, kappa.Sub(sdk.OneDec()))
	if err != nil {
		return sdk.Dec{}, err
	}
	temp3, err := ApproxRoot(temp2, kappa)
	if err != nil {
		return sdk.Dec{}, err
	}
	return (kappa.Mul(temp3)).Quo(temp1), nil
}

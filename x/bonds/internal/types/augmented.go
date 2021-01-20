package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Inspired by work from BlockScience:
// https://github.com/BlockScience/cadCAD-Tutorials/tree/master/00-Reference-Mechanisms

// value function for a given state (R,S)
func Invariant(R, S sdk.Dec, kappa sdk.Dec) sdk.Dec {
	return ApproxPower(S, kappa).Quo(R)
}

// given a value function (parameterized by kappa)
// and an invariant coeficient V0
// return Supply S as a function of reserve R
func Supply(R sdk.Dec, kappa sdk.Dec, V0 sdk.Dec) sdk.Dec {
	result, err := ApproxRoot(V0.Mul(R), kappa)
	if err != nil {
		panic(err)
	}
	return result
}

// This is the reverse of Supply(...) function
func Reserve(S sdk.Dec, kappa sdk.Dec, V0 sdk.Dec) sdk.Dec {
	return ApproxPower(S, kappa).Quo(V0)
}

// given a value function (parameterized by kappa)
// and an invariant coeficient V0
// return a spot price P as a function of reserve R
func SpotPrice(R sdk.Dec, kappa sdk.Dec, V0 sdk.Dec) sdk.Dec {
	temp1, err := ApproxRoot(V0, kappa)
	if err != nil {
		panic(err)
	}
	temp2, err := ApproxRoot(ApproxPower(R, kappa.Sub(sdk.OneDec())), kappa)
	if err != nil {
		panic(err)
	}
	return (kappa.Mul(temp2)).Quo(temp1)
}

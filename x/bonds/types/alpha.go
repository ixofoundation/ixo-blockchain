package types

import (
	"cosmossdk.io/math"
)

var (
	StartingPublicAlpha = math.LegacyMustNewDecFromStr("0.5")
)

func SystemAlpha(publicAlpha math.LegacyDec, S0, S1, R, C math.Int) math.LegacyDec {
	// S0/S1: negative and positive attestations, measured in bond tokens
	// C: outcome payment
	// R: current reserve

	S0R := R.Mul(S1)
	S1R := R.Mul(S1)
	S0C := C.Mul(S0)

	x := math.LegacyNewDecFromInt(S1R)
	y := math.LegacyNewDecFromInt(S1R.Sub(S0R).Add(S0C))
	return publicAlpha.Mul(x.Quo(y))
	// return sdk.NewDecWithPrec(05, 1)
}

func Kappa(I math.LegacyDec, C math.Int, alpha math.LegacyDec) math.LegacyDec {
	// I: invariant
	// C: outcome payment

	x := I
	z := alpha.MulInt(C)
	y := I.Sub(z)
	return x.Quo(y)
}

func InvariantI(C math.Int, alpha math.LegacyDec, R math.Int) math.LegacyDec {
	// C: outcome payment
	// R: current reserve

	return alpha.MulInt(C).Add(math.LegacyNewDecFromInt(R))
}

func InvariantIAlt(C math.Int, alpha math.LegacyDec, kappa math.LegacyDec) math.LegacyDec {
	// C: outcome payment

	x := alpha.MulInt(C)
	y := math.LegacyOneDec().Sub(math.LegacyOneDec().Quo(kappa))
	return x.Quo(y)
}

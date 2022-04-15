package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	StartingPublicAlpha = sdk.MustNewDecFromStr("0.5")
)

func SystemAlpha(publicAlpha sdk.Dec, S0, S1, R, C sdk.Int) sdk.Dec {
	// S0/S1: negative and positive attestations, measured in bond tokens
	// C: outcome payment
	// R: current reserve

	S0R := R.Mul(S1)
	S1R := R.Mul(S1)
	S0C := C.Mul(S0)

	x := sdk.NewDecFromInt(S1R)
	y := sdk.NewDecFromInt(S1R.Sub(S0R).Add(S0C))
	return publicAlpha.Mul(x.Quo(y))
	// return sdk.NewDecWithPrec(05, 1)
}

func Kappa(I sdk.Dec, C sdk.Int, alpha sdk.Dec) sdk.Dec {
	// I: invariant
	// C: outcome payment

	x := I
	z := alpha.MulInt(C)
	y := I.Sub(z)
	//fmt.Println(x)
	//fmt.Println(z)
	//fmt.Println(y)
	return x.Quo(y)
}

func InvariantI(C sdk.Int, alpha sdk.Dec, R sdk.Int) sdk.Dec {
	// C: outcome payment
	// R: current reserve

	return alpha.MulInt(C).Add(sdk.NewDecFromInt(R))
}

func InvariantIAlt(C sdk.Int, alpha sdk.Dec, kappa sdk.Dec) sdk.Dec {
	// C: outcome payment

	x := alpha.MulInt(C)
	y := sdk.OneDec().Sub(sdk.OneDec().Quo(kappa))
	return x.Quo(y)
}

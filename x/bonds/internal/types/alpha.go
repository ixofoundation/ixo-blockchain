package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// Idea 1: change MsgEditKappa to MsgEditAlpha
// If alpha < 0.5, kappa = 2
// If alpha = 0.5, kappa = 3
// If alpha > 0.5, kappa = 4
// This way we can have alpha being used in the system and be able to show it.

// Idea 2: calculate alpha from kappa (rather than the other way round)
// This way we can display an alpha, but this will not match the actual alpha.

func Alpha(S0, S1, R, C sdk.Int) sdk.Dec {
	// S0/S1: negative and positive attestations, measured in bond tokens
	// C: outcome payment
	// R: current reserve

	S0R := R.Mul(S1)
	S1R := R.Mul(S1)
	S0C := C.Mul(S0)

	x := sdk.NewDecFromInt(S1R)
	y := sdk.NewDecFromInt(S1R.Sub(S0R).Add(S0C))
	return x.Quo(y)
}

func Kappa(I sdk.Dec, C sdk.Int, alpha sdk.Dec) sdk.Dec {
	// I: invariant
	// C: outcome payment

	x := I
	y := I.Sub(alpha.MulInt(C))
	return x.Quo(y)
}

func InvariantI(C sdk.Int, alpha sdk.Dec, kappa sdk.Dec) sdk.Dec {
	// C: outcome payment

	x := alpha.MulInt(C)
	y := sdk.OneDec().Sub(sdk.OneDec().Quo(kappa))
	return x.Quo(y)
}

func InvariantIAlt(C sdk.Int, alpha sdk.Dec, R sdk.Int) sdk.Dec {
	// C: outcome payment
	// R: current reserve

	return alpha.MulInt(C).Add(sdk.NewDecFromInt(R))
}

//func (bond *Bond) UpdateOnBurnOrMint() {
//	paramsMap := bond.FunctionParameters.AsMap()
//	d0, _ := paramsMap["d0"]
//	p0, _ := paramsMap["p0"]
//	theta, _ := paramsMap["theta"]
//	kappa, _ := paramsMap["kappa"]
//	alpha, _ := paramsMap["alpha"]
//	R0, _ := paramsMap["R0"]
//	S0, _ := paramsMap["S0"]
//	V0, _ := paramsMap["V0"]
//
//	commonReserveBalance := bond.CurrentReserve[0].Amount
//	I := InvariantIAlt(bond.OutcomePayment, alpha, commonReserveBalance)
//	kappa = Kappa(I, bond.OutcomePayment, alpha)
//
//	bond.FunctionParameters = FunctionParams{
//		NewFunctionParam("d0", d0),
//		NewFunctionParam("p0", p0),
//		NewFunctionParam("theta", theta),
//		NewFunctionParam("kappa", kappa),
//		NewFunctionParam("alpha", alpha),
//		NewFunctionParam("R0", R0),
//		NewFunctionParam("S0", S0),
//		NewFunctionParam("V0", V0),
//		NewFunctionParam("I", I),
//	}
//}
//
//func (bond *Bond) UpdateOnAttestation() {
//	paramsMap := bond.FunctionParameters.AsMap()
//	d0, _ := paramsMap["d0"]
//	p0, _ := paramsMap["p0"]
//	theta, _ := paramsMap["theta"]
//	//kappa, _ := paramsMap["kappa"]
//	alpha, _ := paramsMap["alpha"]
//	R0, _ := paramsMap["R0"]
//	S0, _ := paramsMap["S0"]
//	//V0, _ := paramsMap["V0"]
//
//	commonReserveBalance := bond.CurrentReserve[0].Amount
//	I := InvariantIAlt(bond.OutcomePayment, alpha, commonReserveBalance)
//	kappa := Kappa(I, bond.OutcomePayment, alpha)
//	RDec := sdk.NewDecFromInt(commonReserveBalance)
//	SDec := sdk.NewDecFromInt(bond.CurrentSupply.Amount)
//	V0 := Invariant(RDec, SDec, kappa)
//
//	bond.FunctionParameters = FunctionParams{
//		NewFunctionParam("d0", d0),
//		NewFunctionParam("p0", p0),
//		NewFunctionParam("theta", theta),
//		NewFunctionParam("kappa", kappa),
//		NewFunctionParam("alpha", alpha),
//		NewFunctionParam("R0", R0),
//		NewFunctionParam("S0", S0),
//		NewFunctionParam("V0", V0),
//		NewFunctionParam("I", I),
//	}
//}

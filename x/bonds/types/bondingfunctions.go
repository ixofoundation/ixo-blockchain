package types

import (
	"fmt"
	"math"

	sdk_math "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// "fmt"

// "github.com/ixofoundation/ixo-blockchain/v6/x/bonds/types"
func toPercentage(f float64) float64 { return f / 100 }

func ConvertFloat64ToDec(f float64) (sdk_math.LegacyDec, error) {
	s := fmt.Sprintf("%.18f", f)
	fmt.Println(f)
	dec, err := sdk_math.LegacyNewDecFromStr(s)
	if err != nil {
		fmt.Println(err)
		return sdk_math.LegacyDec{}, err
	}
	return dec, nil
}

type BondingAlgorithm interface {
	recalculate() error
	Init(Bond) error
	Revision() int64
	CalculatePriceForTokens(price sdk.Coin) (sdk.DecCoin, error)
	CalculateTokensForPrice(price sdk.Coin) (sdk.DecCoin, error)
	ExportToMap() map[string]float64
	ExportToBond(bond *Bond) error
}

// type AugmentedBondRevision interface {
// 	init() error
// 	recalculate() error
// 	UpdateAlpha(float64) error
// }

type AugmentedBondRevision1 struct {
	_m float64 // CurrentSupply
	_M float64 // Max Supply
	_C float64 // Outcome Payment
	_T float64 // Time to Maturity
	_R float64 // Current Reserve

	// Issuance
	_F      float64 // Targeted Funding
	_Fi     float64 // Targeted Funding
	_Mh     float64 // Hatch Supply
	_Ph     float64 // Fixed hatch price
	_Mi     float64 // Hatch Supply
	_APYmin float64 // Minimum APY
	_APYavg float64 // Minimum APY
	_APYmax float64 // Maximum APY

	// Initialization
	// _theta float64 // Theta
	_a   float64 // System Alpha
	_a0  float64 // Initial System Alpha
	_g   float64 // System Gamma
	_ap  float64 // System Alpha Public
	_ap0 float64 // Initial Alpha Public
	_t   float64 // Time

	_r float64 // Discounting rate
	// _p1 math.LegacyDec // Maximum Price

	_Pmin float64 // Minimum Price
	_Pavg float64 // Average Price
	_Pmax float64 // Maximum Price

	_kappa float64

	_B float64 // Beta (Shape of the curve)
}

func (algo *AugmentedBondRevision1) recalculate() error {
	algo._Mi = algo._M - algo._Mh
	algo._Fi = algo._F - (algo._Mi * algo._Ph)
	algo._a0 = math.Exp(-1 * (algo._APYmin - algo._r) * algo._T)

	algo._kappa = (algo._a / algo._a0) * math.Exp(algo._r*algo._t)

	//TODO: work need here

	algo._Pavg = algo._kappa * algo._Fi / algo._Mi
	algo._APYavg = -1 * (math.Log(algo._Pavg * algo._M / algo._C)) / algo._T

	// bond._Pmax = bond._M * bond._C

	algo._Pmax = algo._kappa * math.Exp(-1*algo._APYmin*algo._T) * (algo._C / algo._M)
	algo._Pmin = algo._kappa * math.Exp(-1*algo._APYmax*algo._T) * (algo._C / algo._M)

	algo._B = (algo._Pmax - algo._Pavg) / (algo._Pavg - algo._Pmin)

	return nil
}

func (algo *AugmentedBondRevision1) ExportToMap() map[string]float64 {
	return map[string]float64{
		"m":      algo._m,
		"M":      algo._M,
		"C":      algo._C,
		"T":      algo._T,
		"R":      algo._R,
		"F":      algo._F,
		"Fi":     algo._Fi,
		"Mh":     algo._Mh,
		"Ph":     algo._Ph,
		"Mi":     algo._Mi,
		"APYmin": algo._APYmin,
		"APYavg": algo._APYavg,
		"APYmax": algo._APYmax,
		"a":      algo._a,
		"a0":     algo._a0,
		"g":      algo._g,
		"ap":     algo._ap,
		"ap0":    algo._ap0,
		"t":      algo._t,
		"r":      algo._r,
		"Pmin":   algo._Pmin,
		"Pavg":   algo._Pavg,
		"Pmax":   algo._Pmax,
		"kappa":  algo._kappa,
		"B":      algo._B,
	}
}

func (algo *AugmentedBondRevision1) ExportToBond(bond *Bond) error {
	valueOrError := RightFlatMap(FromError(ConvertFloat64ToDec(algo._ap)), func(ap sdk_math.LegacyDec) Either[error, bool] {
		return RightFlatMap(FromError(ConvertFloat64ToDec(algo._a)), func(a sdk_math.LegacyDec) Either[error, bool] {
			bond.FunctionParameters.ReplaceParam("PUBLIC_ALPHA", ap)
			bond.FunctionParameters.ReplaceParam("SYSTEM_ALPHA", a)
			return Right[error](true)
		})
	})

	if !valueOrError.IsRight() {
		return valueOrError.Left()
	}

	return nil
}

func (algo *AugmentedBondRevision1) Revision() int64 { return 1 }

func (algo *AugmentedBondRevision1) Init(bond Bond) error {
	params := bond.FunctionParameters.AsMap()
	bondReserve := func() (total sdk_math.Int) {
		// for _, coin := range bond.CurrentReserve {
		// 	total.Add(coin.Amount)
		// }
		return bond.CurrentReserve[0].Amount
	}

	// Check Revision
	if rev, exists := params["REVISION"]; !exists {
		return fmt.Errorf("REVISION not found in bond function parameters")
	} else if rev.TruncateInt64() != algo.Revision() {
		return fmt.Errorf("REVISION in bond function parameters is not %d", algo.Revision())
	}

	var valueOrError = RightFlatMap(FromError(bond.CurrentSupply.Amount.ToLegacyDec().Float64()), func(_m float64) Either[error, bool] {
		return RightFlatMap(FromError(bond.MaxSupply.Amount.ToLegacyDec().Float64()), func(_M float64) Either[error, bool] {
			return RightFlatMap(FromError(bond.OutcomePayment.ToLegacyDec().Float64()), func(_C float64) Either[error, bool] {
				return RightFlatMap(FromError(bondReserve().ToLegacyDec().Float64()), func(_R float64) Either[error, bool] {
					return RightFlatMap(FromError(params["Funding_Target"].Float64()), func(_F float64) Either[error, bool] {
						return RightFlatMap(FromError(params["Hatch_Supply"].Float64()), func(_Mh float64) Either[error, bool] {
							return RightFlatMap(FromError(params["Hatch_Price"].Float64()), func(_Ph float64) Either[error, bool] {
								return RightFlatMap(FromError(params["APY_MAX"].Float64()), func(_APYmax float64) Either[error, bool] {
									return RightFlatMap(FromError(params["APY_MIN"].Float64()), func(_APYmin float64) Either[error, bool] {
										return RightFlatMap(FromError(params["DISCOUNT_RATE"].Float64()), func(_r float64) Either[error, bool] {
											return RightFlatMap(FromError(params["MATURITY"].Float64()), func(_T float64) Either[error, bool] {
												return RightFlatMap(FromError(params["GAMMA"].Float64()), func(_g float64) Either[error, bool] {
													return RightFlatMap(FromError(params["INITIAL_PUBLIC_ALPHA"].Float64()), func(_ap0 float64) Either[error, bool] {
														algo._m = _m
														algo._M = _M
														algo._C = _C
														algo._R = _R
														algo._F = _F
														algo._Mh = _Mh
														algo._Ph = _Ph
														algo._T = _T
														algo._g = _g
														algo._APYmax = toPercentage(_APYmax)
														algo._APYmin = toPercentage(_APYmin)
														algo._r = toPercentage(_r)
														// algo._theta = toPercentage2(_theta)
														algo._ap0 = toPercentage(_ap0)

														algo._t = 0

														return Right[error](true)
													})
												})
											})
										})
									})
								})
							})
						})
					})
				})
			})
		})
	})

	if !valueOrError.IsRight() {
		return valueOrError.Left()
	}

	// Initialize the initial alpha0
	algo._a0 = math.Exp(-1 * (algo._APYmin - algo._r) * algo._T)

	var err error
	algo._ap, err = func() (float64, error) {
		if _paDec, exists := params["PUBLIC_ALPHA"]; exists {
			_pa, err := _paDec.Float64()
			return _pa, err
		} else {
			return algo._ap0, nil
		}
	}()

	if err != nil {
		return err
	}

	algo._a, err = func() (float64, error) {
		if _aDec, exists := params["SYSTEM_ALPHA"]; exists {
			_a, err := _aDec.Float64()
			return _a, err
		} else {
			return algo._a0, nil
		}
	}()

	if err != nil {
		return err
	}

	return algo.recalculate()
}

func (algo *AugmentedBondRevision1) UpdateAlpha(_ap, _delta float64) error {
	if _ap < 0 {
		return fmt.Errorf("alpha is smaller than 0 and must be greater than or equal to 0 and smaller than or equal to 1")
	}
	if _ap > 1 {
		return fmt.Errorf("alpha is larger than 1 and must be greater than or equal to 0 and smaller than or equal to 1")
	}

	gamma1 := algo._g * (1 - (1 / algo._a0)) / (1 - (1 / algo._ap0))

	alpha_new := func() float64 {
		if _ap <= algo._ap0 {
			return algo._a0 * math.Pow((_ap/algo._ap0), gamma1)
		} else {
			return 1 - (1-algo._a0)*math.Pow((1-_ap)/(1-algo._ap0), algo._g)
		}
	}()
	algo._a = (_delta * alpha_new) + (1-_delta)*algo._a

	return algo.recalculate()
}

func (algo *AugmentedBondRevision1) CalculatePriceForTokens(price sdk.Coin) (sdk.DecCoin, error) {
	// if !alphaBond.AlphaBond {
	// return types.Price{}, errors.New("not an alpha bond")
	// }

	_mh := float64(0)
	// _dm, _ := price.Amount.ToLegacyDec().Float64()
	calc := (algo._Mh - _mh)
	dec, _ := ConvertFloat64ToDec(calc)
	return sdk.NewDecCoinFromDec(price.Denom, dec), nil
}

func (algo *AugmentedBondRevision1) CalculateTokensForPrice(price sdk.Coin) (sdk.DecCoin, error) {
	// if !bond.AlphaBond {
	// 	// return types.Price{}, errors.New("not an alpha bond")
	// }

	_dm, _ := price.Amount.ToLegacyDec().Float64()
	calc := (_dm * algo._Pmin) + (((algo._Mi * (algo._Pmax - algo._Pmin)) / (algo._B + 1)) * (math.Pow(((algo._m+_dm)/algo._Mi), (algo._B+1)) - math.Pow((algo._m/algo._M), (algo._B+1))))
	dec, _ := ConvertFloat64ToDec(calc)
	return sdk.NewDecCoinFromDec(price.Denom, dec), nil
}

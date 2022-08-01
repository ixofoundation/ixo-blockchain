package bondingmath

import (
	"fmt"
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/types"
)

// "fmt"

// "github.com/ixofoundation/ixo-blockchain/x/bonds/types"

type BondingAlgorithm interface {
	AugmentedBondRevision
	init() error
	recalculate() error
}

type AugmentedBondRevision interface {
	init() error
	recalculate() error
	UpdateAlpha(float64) error
}

type AugmentedBondRevision1 struct {
	AugmentedBondRevision
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
	_theta float64 // Theta
	_a     float64 // System Alpha
	_a0    float64 // Initial System Alpha
	_g     float64 // System Gamma
	_ap    float64 // System Alpha Public
	_ap0   float64 // Intial Alpha Public
	_t     float64 // Time

	_r float64 // Discounting rate
	// _p1 sdk.Dec // Maximum Price

	_Pmin float64 // Minimum Price
	_Pavg float64 // Average Price
	_Pmax float64 // Maximum Price

	_kappa float64

	_B float64 // Beta (Shape of the curve)
}

func (algo *AugmentedBondRevision1) recalculate() error {
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

// type Eiterr[L any, R any]  {

// }

type Either[L any, R any] interface {
	IsRightBiased() bool
	IsRight() bool
	Left() L
	Right() R
}

type leftImpl[L any, R any] struct {
	value L
}

func (l leftImpl[L, R]) IsRightBiased() bool { return true }
func (l leftImpl[L, R]) IsRight() bool       { return false }
func (l leftImpl[L, R]) Left() L             { return l.value }
func (l leftImpl[L, R]) Right() R            { return *new(R) }

type rightImpl[L any, R any] struct {
	value R
}

func (r rightImpl[L, R]) IsRightBiased() bool { return true }
func (r rightImpl[L, R]) IsRight() bool       { return true }
func (r rightImpl[L, R]) Left() L             { return *new(L) }
func (r rightImpl[L, R]) Right() R            { return r.value }

func FromError[R any](r R, err error) Either[error, R] {
	if err != nil {
		return Left[error, R](err)
	}
	return Right[error, R](r)
}

func Left[L any, R any](l L) Either[L, R] {
	return leftImpl[L, R]{value: l}
}

func Right[L any, R any](r R) Either[L, R] {
	return rightImpl[L, R]{value: r}
}

func RightFlatMap[L any, R any, RR any](either Either[L, R], f func(r R) Either[L, RR]) Either[L, RR] {
	if either.IsRightBiased() && either.IsRight() {
		return f(either.Right())
	}
	return Left[L, RR](either.Left())
}

func LeftFlatMap[L any, R any, LL any](either Either[L, R], f func(l L) Either[LL, R]) Either[LL, R] {
	if !either.IsRightBiased() && !either.IsRight() {
		return f(either.Left())
	}
	return Right[LL, R](either.Right())
}

// func FlatMap[L any, R any, LL func(B) A, RR any](either Either[L, R], f func(any) any) Either[LL, R] {
// 	switch {
// 		case either.IsRightBiased() && either.IsRight():
// 			return Right[LL, R](f(either.Right()))
// 	}
// 	return Either(either.Left(), either.Right())
// }

// func Map[L any, R any](either Either[L, R])(f func(R) L) {
// 	if either.IsRightBiased() && either.IsRight() {
// 		return f(either.Right())
// 	} else if either.IsRightBiased() && !either.IsRight() {
// 		return f(either.Left())
// 	}

// func Map[L any, R any](either Either[L, R])(f func(R) L) {
// 	if either.IsRightBiased() && either.IsRight() {
// 		f(either.Right())
// 	}
// }

func toPercentage2(f float64) float64 { return f / 100 }

func (algo *AugmentedBondRevision1) init(bond types.Bond) error {
	params := bond.FunctionParameters.AsMap()
	bondReserve := func() (total sdk.Int) {
		for _, coin := range bond.CurrentReserve {
			total.Add(coin.Amount)
		}
		return
	}

	var valueOrError Either[error, float64] = RightFlatMap(FromError(bond.CurrentSupply.Amount.ToDec().Float64()), func(_m float64) Either[error, float64] {
		return RightFlatMap(FromError(bond.MaxSupply.Amount.ToDec().Float64()), func(_M float64) Either[error, float64] {
			return RightFlatMap(FromError(bond.OutcomePayment.ToDec().Float64()), func(_C float64) Either[error, float64] {
				return RightFlatMap(FromError(bondReserve().ToDec().Float64()), func(_R float64) Either[error, float64] {
					return RightFlatMap(FromError(params["Funding_Target"].Float64()), func(_F float64) Either[error, float64] {
						return RightFlatMap(FromError(params["Hatch_Supply"].Float64()), func(_Mh float64) Either[error, float64] {
							return RightFlatMap(FromError(params["Hatch_Price"].Float64()), func(_Ph float64) Either[error, float64] {
								return RightFlatMap(FromError(params["APY_MAX"].Float64()), func(_APYmax float64) Either[error, float64] {
									return RightFlatMap(FromError(params["APY_MIN"].Float64()), func(_APYmin float64) Either[error, float64] {
										return RightFlatMap(FromError(params["DISCOUNT_RATE"].Float64()), func(_r float64) Either[error, float64] {
											return RightFlatMap(FromError(params["MATURITY"].Float64()), func(_T float64) Either[error, float64] {
												return RightFlatMap(FromError(params["GAMMA"].Float64()), func(_g float64) Either[error, float64] {
													return RightFlatMap(FromError(params["THETA"].Float64()), func(_theta float64) Either[error, float64] {
														return RightFlatMap(FromError(params["INITIAL_PUBLIC_ALPHA"].Float64()), func(_ap0 float64) Either[error, float64] {
															algo._m = _m
															algo._M = _M
															algo._C = _C
															algo._R = _R
															algo._F = _F
															algo._Mh = _Mh
															algo._Ph = _Ph
															algo._T = _T
															algo._g = _g
															algo._APYmax = toPercentage2(_APYmax)
															algo._APYmin = toPercentage2(_APYmin)
															algo._r = toPercentage2(_r)
															algo._theta = toPercentage2(_theta)
															algo._ap0 = toPercentage2(_ap0)

															algo._t = 0

															return Right[error, float64](1)
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
	})

	if !valueOrError.IsRight() {
		return valueOrError.Left()
	}

	var err error
	algo._ap, err = func() (float64, error) {
		if _paDec, exists := params["PUBLIC_ALPHA"]; exists {
			_pa, err := _paDec.Float64()
			return _pa, err
		} else {
			return algo._ap0, nil
		}
	}()

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

func (algo *AugmentedBondRevision1) UpdateAlpha(_ap float64) error {

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
	algo._a = (algo._theta * alpha_new) + (1-algo._theta)*algo._a

	return algo.recalculate()
}

func InitializeBondingAlgorithm[A BondingAlgorithm](bond types.Bond, algo A) (A, error) {
	if err := algo.init(); err != nil {
		return *new(A), nil
	}

	return algo, nil
}
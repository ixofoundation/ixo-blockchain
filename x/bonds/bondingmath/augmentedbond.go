package bondingmath

type BondingAlgorithm[A any] struct {
	params A
}




type AugmentedBondRevision1 struct {
	m float64
}

// type BondingAlgorithm interface {
// 	AugmentedBond

// 	recalculate() error

// 	Revision() int64
// 	Type() string
// }

// type AugmentedBond interface {
// }

// type AugmentedBondRevision1 struct {
// }

// func (ar AugmentedBondRevision1) UpdateAlpha(ap float64) {
// 	x, _ := setAndRecalculate(*ar, func(state any) (any, error) {
// 		return nil, nil
// 	})
// }

// type AugmentedBond interface {
// }

// type AugmentedBondRevision1 int

// func (AugmentedBondRevision1) Init()

// func InitializeAugmentedBond(bond types.Bond) (AugmentedBond, error) {
// 	revision := 1
// 	params := map[string]sdk.Dec{}

// 	switch revision {
// 	case 1:
// 		AugmentedBond
// 	}

// 	return nil, nil
// }

// type AugmentedBondRevision1 struct {

// 	// Economics
// 	_m float64 // CurrentSupply
// 	_M float64 // Max Supply
// 	_C float64 // Outcome Payment
// 	_T float64 // Time to Maturity
// 	_R float64 // Current Reserve

// 	// Issuance
// 	_F      float64 // Targeted Funding
// 	_Fi     float64 // Targeted Funding
// 	_Mh     float64 // Hatch Supply
// 	_Ph     float64 // Fixed hatch price
// 	_Mi     float64 // Hatch Supply
// 	_APYmin float64 // Minimum APY
// 	_APYavg float64 // Minimum APY
// 	_APYmax float64 // Maximum APY

// 	// Initialization
// 	_theta float64 // Theta
// 	_a     float64 // System Alpha
// 	_a0    float64 // Initial System Alpha
// 	_g     float64 // System Gamma
// 	_ap    float64 // System Alpha Public
// 	_ap0   float64 // Intial Alpha Public
// 	_t     float64 // Time

// 	_r float64 // Discounting rate
// 	// _p1 sdk.Dec // Maximum Price

// 	_Pmin float64 // Minimum Price
// 	_Pavg float64 // Average Price
// 	_Pmax float64 // Maximum Price

// 	_kappa float64

// 	_B float64 // Beta (Shape of the curve)

// }

// func (b *AugmentedBondRevision1) Init() {

// }
// type AugmentedBond[R AugmentedBondRevision] struct {
// 	m
// }

// type AugmentedBondRevision interface {
// 	init()
// 	updateAlpha()
// }

// type AugmentedBondRevision1 struct {

// }

// func alphabondRev1() {
// 	// Economics
// 	_m float64 // CurrentSupply
// 	_M float64 // Max Supply
// 	_C float64 // Outcome Payment
// 	_T float64 // Time to Maturity
// 	_R float64 // Current Reserve

// 	// Issuance
// 	_F      float64 // Targeted Funding
// 	_Fi     float64 // Targeted Funding
// 	_Mh     float64 // Hatch Supply
// 	_Ph     float64 // Fixed hatch price
// 	_Mi     float64 // Hatch Supply
// 	_APYmin float64 // Minimum APY
// 	_APYavg float64 // Minimum APY
// 	_APYmax float64 // Maximum APY
// }

// func (AugmentedBond[R]) Revision() int64 { return 1 }
// func (AugmentedBond[R]) Type() string    { return "AugmentedBond" }

// func test() {

// 	var x BondingAlgorithm = AugmentedBondRevision1{}
// }

// func (a BondingAlgorithm[A]) PricePerToken() {
// 	a.algorithm
// }

// type BondingAlgorithmRevision interface {
// 	AugmentedBond
// 	Revision() int64
// 	Type()
// }

// type AugmentedBond interface {
// 	Init() (func(types.Bond) BondingAlgorithm[AugmentedBond], int64)
// 	Type()
// 	Revision() int64
// 	UpdateAlpha()
// }

// type AugmentedBondRevision1 struct {
// }

// func (AugmentedBondRevision1) Init() (func(types.Bond) BondingAlgorithm[AugmentedBond], int64) {
// 	return func(bond types.Bond) BondingAlgorithm[AugmentedBond] {

// 	}
// }

// func (AugmentedBondRevision1) Type()
// func (AugmentedBondRevision1) Revision() int64 { return 1 }
// func (AugmentedBondRevision1) UpdateAlpha()

// type AugmentedBondRevision2 struct {
// }

// func (AugmentedBondRevision2) Init() (func(types.Bond) BondingAlgorithm[AugmentedBond], int64) {
// 	return func(bond types.Bond) BondingAlgorithm[AugmentedBond] {

// 	}
// }

// func (AugmentedBondRevision2) Type()
// func (AugmentedBondRevision2) Revision() int64 { return 1 }
// func (AugmentedBondRevision2) UpdateAlpha()

// func InitializeAugmentedBond(bond types.Bond) (BondingAlgorithm[AugmentedBond], error) {

// 	x := AugmentedBondRevision1{}.Init()
// 	y, rev := AugmentedBondRevision2{}.Init()

// switch {
// case y, rev := AugmentedBondRevision2{}.Init(); rev == 0:
// }

// 	// f := func(revision int64) AugmentedBond {
// 	// 	switch revision {
// 	// 	case AugmentedBondRevision1{}.Revision():
// 	// 		return AugmentedBondRevision1{}.Init()
// 	// 	default:
// 	// 		return AugmentedBondRevision1{}
// 	// 	}
// 	// }(0)

// 	return BondingAlgorithm[AugmentedBond]{}, nil
// }

// func Test() {

// 	y := AugmentedBond(AugmentedBondRevision1{})

// 	fmt.Println(y)

// 	x, _ := InitializeAugmentedBond(types.Bond{})

// }

// type BondingAlgorithm interface {
// 	Revision() int
// }

// type AugmentedBond interface {
// 	BondingAlgorithm
// }

// func recalculate[A BondingAlgorithm](a *A, f func(A) (A, error)) (A, error) {
// 	return f(*a)
// }

// type AugmentedBondRevision1 struct {
// }

// func (AugmentedBondRevision1) Revision() int { return 1 }

// func (ab *AugmentedBondRevision1) UpdateAlpha() {
// 	ab, err = recalculate(ab, func(state AugmentedBondRevision1) {

// 	})
// 	if err != nil {
// 		return nil
// 	}
// 	ab = ab
// }

// type BondingAlgorithm interface{}

// func InitializeBondingAlgorithm[B any](bond B, f func(B) BondingAlgorithm) BondingAlgorithm {
// 	return f(bond)
// }

// type AugmentedBond struct {
// 	BondingAlgorithm
// }

// func (AugmentedBond) Init() func(types.Bond) BondingAlgorithm {
// 	return func(bb types.Bond) BondingAlgorithm {
// 		return AugmentedBond{}
// 	}
// }

// func Test() {
// 	InitializeBondingAlgorithm(types.Bond{}, AugmentedBond{}.Init())
// }

// type BondingAlgorithmRevision interface {
// }

// type BondingAlgorithm[R BondingAlgorithmRevision] interface {
// 	Revisions() []BondingAlgorithmRevision
// }

// type AugmentedBondRevision interface {
// }

// type AugmentedBond[R AugmentedBondRevision] struct {

// }

// func (a AugmentedBond[R]) Revisions() []R {
//  []R{

//  }
// }

// type AugmentedBondRevision1 struct {
// }

// func (x AugmentedBondRevision1) Revision() {

// }

// type AugmentedBondRevision2 struct{}

// type AugmentedBondRevision interface {
// 	AugmentedBondRevision1 | AugmentedBondRevision2
// 	Revision()
// }
// type AugmentedBond[R AugmentedBondRevision] func(R) struct {
// 	x string
// }

// func Test[R AugmentedBondRevision](r R) {

// }

// func Test2() {
// 	Test(AugmentedBondRevision1{})
// }

// var t AugmentedBond[AugmentedBondRevision1] = func(a AugmentedBondRevision1) struct{ x string } {
// 	return struct{ x string }{
// 		x: "",
// 	}
// }

// func (ab AugmentedBond[R]) UpdateAlpha(pa float64) {
// 	rev := func() R {
// 		return R(AugmentedBondRevision1{})
// 	}()
// 	y := ab(rev).x
// 	fmt.Println(y)
// 	// ab(x)

// }

// type AugmentedBondRevisionV1 struct {
// }

// func (x AugmentedBondRevisionV1) Revision() {

// }

// type AugmentedBondRevisionV2 struct {
// }

// type X struct {
// }

// func (x X) Revision() {

// }

// type AugmentedBond interface {
// 	AugmentedBondRevisionV1 | AugmentedBondRevisionV2
// 	Revision()
// }

// func Upgrade[A AugmentedBond](a A) {

// }

// func Test() {

// 	Upgrade(X{})
// }

// type BondingAlgorithm string

// const AugmentedBond BondingAlgorithm = "augmented_bond"

// type AlgorithmRevision interface{}

// func Init(bond types.Bond) {

// }

// type X struct {
// 	AlgorithmRevision
// }

// type BondFunction[A BondingAlgorithm, R AlgorithmRevision] interface {
// 	Revision() int64
// 	GetPice() int64
// }

// type AlphaBond[A BondingAlgorithm, R AlgorithmRevision] func(A, R)

// var AugmentedBond AlphaBond[BondingAlgorithm, X] = func(a BondingAlgorithm, r X) {

// }

// func (ab AlphaBond[A, R]) Revision() int64 { return 1 }
// func (ab AlphaBond[A, R]) GetPrice() int64 { ab(nil, nil) }

// type AlgorithmRevision interface {
// 	Upgrade(func()) (AlgorithmRevision, error)
// }

// type BondingAlgorithm interface {
// 	Revision() int64
// 	Paramaters() map[string]string
// 	SetParamaters() error
// 	// UpdateToNextRevision(func(Rev[NRev]) (Rev, error)) (BondingAlgorithm[Rev, NRev], error)

// 	CalculatePriceForTokens() (sdk.Coin, error)
// }

// type AugBond[R AlgorithmRevision] struct {
// 	algorithmRevision R
// 	_m                float64 // CurrentSupply
// 	_M                float64 // Max Supply
// 	_R                float64 // Current Reserve
// }

// func (a AugBond[R]) Y(c R)

// func InitializeAugmentedBond(bond types.Bond) (AugBond[AlgorithmRevision], error) {
// 	AugBond[AlgorithmRevision]{
// 		algorithmRevision: func(revision int) {

// 		},
// 	}

// 	if bond.FunctionType != types.AugmentedFunction {
// 		return AugBond[AlgorithmRevision]{}, fmt.Errorf("Unsupported bond revision")
// 	}

// 	params := bond.FunctionParameters.AsMap()
// 	revision := params["revision"].TruncateInt64()

// 	rev, _ := func(revision int64) (AlgorithmRevision, error) {
// 		// 	switch revision {
// 		// 	case 1:
// 		// 		return AugmentedBondRevision1{}, nil
// 		// 	case 2:
// 		// 		return AugmentedBondRevision2{}, nil
// 		// 		break
// 		// 	}
// 		return nil, fmt.Errorf("Unsupported bond revision")
// 	}(revision)

// 	state := AugBond[AlgorithmRevision]{
// 		// revision:          revision,
// 		algorithmRevision: rev,
// 	}

// 	// ab.algorithmRevision = rev.initalize()
// 	return state, nil
// }

// // func InitializeBondingAlgorithm(bond *types.Bond) (BondingAlgorithm, error) {

// // }

// // func (b *AugBond[Rev, NRev]) UpdateToNextRevision(f func(Rev) (Rev, error)) (BondingAlgorithm[Rev, NRev], error) {
// // 	return nil, nil
// // }

// func UpdateAlpha()

// // type AugmentedBond[T AugmentedBondRevision] struct {
// // 	revision          int64
// // 	algorithmRevision T
// // 	_m                float64 // CurrentSupply
// // 	_M                float64 // Max Supply
// // 	_R                float64 // Current Reserve
// // }

// // type AugmentedBondRevision interface {
// // 	X()
// // }

// // type InvalidRevision string

// // func (InvalidRevision) X() {}

// // type AugmentedBondRevision1 struct {
// // 	x, y float64
// // }

// // func (x AugmentedBondRevision1) X() {}

// // type AugmentedBondRevision2 struct {
// // 	// Economics
// // 	_C float64 // Outcome Payment
// // 	_T float64 // Time to Maturity
// // 	_F float64 // Targeted Funding

// // 	// Issuance
// // 	_Fi float64 // Targeted Funding
// // 	_Mi float64 // Hatch Supply
// // 	_Mh float64 // Hatch Supply
// // 	_Ph float64 // Fixed hatch price

// // 	_APYmin float64 // Minimum APY
// // 	_APYavg float64 // Minimum APY
// // 	_APYmax float64 // Maximum APY

// // 	// Initialization
// // 	_theta float64 // Theta
// // 	_a     float64 // System Alpha
// // 	_a0    float64 // Initial System Alpha
// // 	_g     float64 // System Gamma
// // 	_ap    float64 // System Alpha Public
// // 	_ap0   float64 // Intial Alpha Public
// // 	_t     float64 // Time

// // 	_r float64 // Discounting rate
// // 	// _p1 sdk.Dec // Maximum Price

// // 	_Pmin float64 // Minimum Price
// // 	_Pavg float64 // Average Price
// // 	_Pmax float64 // Maximum Price

// // 	_kappa float64

// // 	_B float64 // Beta (Shape of the curve)
// // }

// // func (x AugmentedBondRevision2) X() {}

// // func recalculate() {

// // }

// // func InitializeAugmentedBond(bond types.Bond) (AugmentedBond[AugmentedBondRevision], error) {
// // 	if bond.FunctionType != types.AugmentedFunction {
// // 		return AugmentedBond[AugmentedBondRevision]{}, fmt.Errorf("Unsupported bond revision")
// // 	}

// // 	params := bond.FunctionParameters.AsMap()
// // 	revision := params["revision"].TruncateInt64()

// // 	rev, _ := func(revision int64) (AugmentedBondRevision, error) {
// // 		switch revision {
// // 		case 1:
// // 			return AugmentedBondRevision1{}, nil
// // 		case 2:
// // 			return AugmentedBondRevision2{}, nil
// // 			break
// // 		}
// // 		return nil, fmt.Errorf("Unsupported bond revision")
// // 	}(revision).initalize(params, )

// // 	state := AugmentedBond[AugmentedBondRevision]{
// // 		revision:          revision,
// // 		algorithmRevision: rev,
// // 	}

// // 	// ab.algorithmRevision = rev.initalize()
// // 	return state, nil
// // }

// // func (ab *AugmentedBond[AugmentedBondRevision]) Initialize(bond *types.Bond) error {
// // 	return nil
// // }

// // func (ab *AugmentedBond) UpdatePublicAlpha() error {
// // 	return nil
// // }

// // func (ab *AugmentedBond) Apply(bond *types.Bond) error {
// // 	return nil
// // }

// // func ConvertFloat64ToDec(f float64) (sdk.Dec, error) {
// // 	s := fmt.Sprintf("%.18f", f)
// // 	fmt.Println(f)
// // 	dec, err := sdk.NewDecFromStr(s)
// // 	if err != nil {
// // 		fmt.Println(err)
// // 		return sdk.Dec{}, err
// // 	}
// // 	return dec, nil
// // }

// // type AlphabondV2 struct {
// // 	// Economics
// // 	_m float64 // CurrentSupply
// // 	_M float64 // Max Supply
// // 	_C float64 // Outcome Payment
// // 	_T float64 // Time to Maturity
// // 	_R float64 // Current Reserve

// // 	// Issuance
// // 	_F      float64 // Targeted Funding
// // 	_Fi     float64 // Targeted Funding
// // 	_Mh     float64 // Hatch Supply
// // 	_Ph     float64 // Fixed hatch price
// // 	_Mi     float64 // Hatch Supply
// // 	_APYmin float64 // Minimum APY
// // 	_APYavg float64 // Minimum APY
// // 	_APYmax float64 // Maximum APY

// // 	// Initialization
// // 	_theta float64 // Theta
// // 	_a     float64 // System Alpha
// // 	_a0    float64 // Initial System Alpha
// // 	_g     float64 // System Gamma
// // 	_ap    float64 // System Alpha Public
// // 	_ap0   float64 // Intial Alpha Public
// // 	_t     float64 // Time

// // 	_r float64 // Discounting rate
// // 	// _p1 sdk.Dec // Maximum Price

// // 	_Pmin float64 // Minimum Price
// // 	_Pavg float64 // Average Price
// // 	_Pmax float64 // Maximum Price

// // 	_kappa float64

// // 	_B float64 // Beta (Shape of the curve)
// // }

// // func toPercentage(f float64, err error) (float64, error) {
// // 	if err != nil {
// // 		return f, err
// // 	}
// // 	return f / 100, err
// // }

// // func (bond *AlphabondV2) Init(alphabond types.Bond) {

// // 	bond._m, _ = alphabond.CurrentSupply.Amount.ToDec().Float64()
// // 	bond._M, _ = alphabond.MaxSupply.Amount.ToDec().Float64()
// // 	bond._C, _ = alphabond.OutcomePayment.ToDec().Float64()
// // 	//TODO: fix this to include many coins.
// // 	bond._R, _ = alphabond.CurrentReserve[0].Amount.ToDec().Float64()

// // 	params := alphabond.FunctionParameters.AsMap()
// // 	bond._F, _ = params["Funding_Target"].Float64()
// // 	bond._Mh, _ = params["Hatch_Supply"].Float64()
// // 	bond._Ph, _ = params["Hatch_Price"].Float64()
// // 	bond._APYmax, _ = toPercentage(params["APY_MAX"].Float64())
// // 	bond._APYmin, _ = toPercentage(params["APY_MIN"].Float64())
// // 	bond._r, _ = toPercentage(params["DISCOUNT_RATE"].Float64())
// // 	bond._T, _ = params["MATURITY"].Float64()
// // 	bond._g, _ = params["GAMMA"].Float64()
// // 	bond._theta, _ = toPercentage(params["THETA"].Float64())
// // 	bond._ap0, _ = toPercentage(params["INITIAL_PUBLIC_ALPHA"].Float64())

// // 	bond._Mi = bond._M - bond._Mh
// // 	bond._Fi = bond._F - (bond._Mi * bond._Ph)

// // 	bond._a0 = math.Exp(-1 * (bond._APYmin - bond._r) * bond._T)

// // 	bond._ap, _ = func() (float64, error) {
// // 		if _paDec, exists := params["PUBLIC_ALPHA"]; exists {
// // 			_pa, err := _paDec.Float64()
// // 			return _pa, err
// // 		} else {
// // 			return bond._ap0, nil
// // 		}
// // 	}()

// // 	bond._a, _ = func() (float64, error) {
// // 		if _aDec, exists := params["SYSTEM_ALPHA"]; exists {
// // 			_a, err := _aDec.Float64()
// // 			return _a, err
// // 		} else {
// // 			return bond._a0, nil
// // 		}
// // 	}()

// // 	bond._t = 0

// // 	bond._kappa = (bond._a / bond._a0) * math.Exp(bond._r*bond._t)

// // 	//TODO: work need here

// // 	bond._Pavg = bond._kappa * bond._Fi / bond._Mi
// // 	bond._APYavg = -1 * (math.Log(bond._Pavg * bond._M / bond._C)) / bond._T

// // 	// bond._Pmax = bond._M * bond._C

// // 	bond._Pmax = bond._kappa * math.Exp(-1*bond._APYmin*bond._T) * (bond._C / bond._M)
// // 	bond._Pmin = bond._kappa * math.Exp(-1*bond._APYmax*bond._T) * (bond._C / bond._M)

// // 	// bond.setCalculatedPavg()
// // 	// bond.setCalculatedAPYavg()
// // 	// bond.setCalculatedPmax()

// // 	// bond.setCalculatedPmin()
// // 	bond._B = (bond._Pmax - bond._Pavg) / (bond._Pavg - bond._Pmin)
// // 	// bond.setCalculatedBeta()
// // 	fmt.Printf("%+v", bond)
// // }

// // func calculateIssuance() {

// // }

// // func (bond *AlphabondV2) Pmax() {

// // }

// // func (bond *AlphabondV2) setCalculatedPavg() {

// // }

// // func (bond *AlphabondV2) setCalculatedPmin() {

// // }

// // // _B = (_Pmax - _Pmin) / (_Pavg - _Pmin)
// // func (bond *AlphabondV2) setCalculatedBeta() error {
// // 	bond._B = (bond._Pmax - bond._Pmin) / (bond._Pavg - bond._Pmin)
// // 	return nil
// // }

// // func (bond *AlphabondV2) UpdateAlpha(_ap float64) error {
// // 	recalulate := func() {

// // 	}()
// // 	// Constraints
// // 	if _ap < 0 {
// // 		return fmt.Errorf("alpha is smaller than 0 and must be greater than or equal to 0 and smaller than or equal to 1")
// // 	}
// // 	if _ap > 1 {
// // 		return fmt.Errorf("alpha is larger than 1 and must be greater than or equal to 0 and smaller than or equal to 1")
// // 	}

// // 	gamma1 := bond._g * (1 - (1 / bond._a0)) / (1 - (1 / bond._ap0))
// // 	fmt.Println("gamma1", gamma1)
// // 	alpha_new := func() float64 {
// // 		if _ap <= bond._ap0 {
// // 			fmt.Println("_ap <= bond._ap0", "true")
// // 			return bond._a0 * math.Pow((_ap/bond._ap0), gamma1)
// // 		} else {
// // 			fmt.Println("_ap <= bond._ap0", "false")
// // 			return 1 - (1-bond._a0)*math.Pow((1-_ap)/(1-bond._ap0), bond._g)
// // 		}
// // 	}()
// // 	fmt.Println("alpha_new", alpha_new)
// // 	bond._a = (bond._theta * alpha_new) + (1-bond._theta)*bond._a
// // 	bond._kappa = (bond._a / bond._a0) * math.Exp(bond._r*bond._t)
// // 	bond._Pavg = bond._kappa * bond._Fi / bond._Mi
// // 	bond._APYavg = -1 * (math.Log(bond._Pavg * bond._M / bond._C)) / bond._T
// // 	bond._Pmax = bond._kappa * math.Exp(-1*bond._APYmin*bond._T) * (bond._C / bond._M)
// // 	bond._Pmin = bond._kappa * math.Exp(-1*bond._APYmax*bond._T) * (bond._C / bond._M)
// // 	bond._B = (bond._Pmax - bond._Pavg) / (bond._Pavg - bond._Pmin)
// // 	bond._ap = _ap

// // 	fmt.Println("alpha", bond._a)

// // 	return nil
// // }

// // func (bond *AlphabondV2) CalculatePriceForTokens(price sdk.Coin) (sdk.DecCoin, error) {
// // 	// if !alphaBond.AlphaBond {
// // 	// return types.Price{}, errors.New("not an alpha bond")
// // 	// }

// // 	_mh := float64(0)
// // 	// _dm, _ := price.Amount.ToDec().Float64()
// // 	calc := (bond._Mh - _mh)
// // 	dec, _ := ConvertFloat64ToDec(calc)
// // 	return sdk.NewDecCoinFromDec(price.Denom, dec), nil

// // }

// // func (bond *AlphabondV2) CalculateTokensForPrice(price sdk.Coin) (sdk.DecCoin, error) {

// // 	// if !bond.AlphaBond {
// // 	// 	// return types.Price{}, errors.New("not an alpha bond")
// // 	// }

// // 	_dm, _ := price.Amount.ToDec().Float64()
// // 	calc := (_dm * bond._Pmin) + (((bond._Mi * (bond._Pmax - bond._Pmin)) / (bond._B + 1)) * (math.Pow(((bond._m+_dm)/bond._Mi), (bond._B+1)) - math.Pow((bond._m/bond._M), (bond._B+1))))
// // 	dec, _ := ConvertFloat64ToDec(calc)
// // 	return sdk.NewDecCoinFromDec(price.Denom, dec), nil
// // }

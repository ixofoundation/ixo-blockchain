// SOURCE: https://github.com/RickGriff/decimath

package types

import sdk "github.com/cosmos/cosmos-sdk/types"

var (
	TEN = []sdk.Uint{
		sdk.NewUintFromString("1"),          // 1e0
		sdk.NewUintFromString("10"),         // 1e1
		sdk.NewUintFromString("100"),        // 1e2
		sdk.NewUintFromString("1000"),       // 1e3
		sdk.NewUintFromString("10000"),      // 1e4
		sdk.NewUintFromString("100000"),     // 1e5
		sdk.NewUintFromString("1000000"),    // 1e6
		sdk.NewUintFromString("10000000"),   // 1e7
		sdk.NewUintFromString("100000000"),  // 1e8
		sdk.NewUintFromString("1000000000"), // 1e9

		sdk.NewUintFromString("10000000000"),          // 1e10
		sdk.NewUintFromString("100000000000"),         // 1e11
		sdk.NewUintFromString("1000000000000"),        // 1e12
		sdk.NewUintFromString("10000000000000"),       // 1e13
		sdk.NewUintFromString("100000000000000"),      // 1e14
		sdk.NewUintFromString("1000000000000000"),     // 1e15
		sdk.NewUintFromString("10000000000000000"),    // 1e16
		sdk.NewUintFromString("100000000000000000"),   // 1e17
		sdk.NewUintFromString("1000000000000000000"),  // 1e18
		sdk.NewUintFromString("10000000000000000000"), // 1e19

		sdk.NewUintFromString("100000000000000000000"),          // 1e20
		sdk.NewUintFromString("1000000000000000000000"),         // 1e21
		sdk.NewUintFromString("10000000000000000000000"),        // 1e22
		sdk.NewUintFromString("100000000000000000000000"),       // 1e23
		sdk.NewUintFromString("1000000000000000000000000"),      // 1e24
		sdk.NewUintFromString("10000000000000000000000000"),     // 1e25
		sdk.NewUintFromString("100000000000000000000000000"),    // 1e26
		sdk.NewUintFromString("1000000000000000000000000000"),   // 1e27
		sdk.NewUintFromString("10000000000000000000000000000"),  // 1e28
		sdk.NewUintFromString("100000000000000000000000000000"), // 1e29

		sdk.NewUintFromString("1000000000000000000000000000000"),          // 1e30
		sdk.NewUintFromString("10000000000000000000000000000000"),         // 1e31
		sdk.NewUintFromString("100000000000000000000000000000000"),        // 1e32
		sdk.NewUintFromString("1000000000000000000000000000000000"),       // 1e33
		sdk.NewUintFromString("10000000000000000000000000000000000"),      // 1e34
		sdk.NewUintFromString("100000000000000000000000000000000000"),     // 1e35
		sdk.NewUintFromString("1000000000000000000000000000000000000"),    // 1e36
		sdk.NewUintFromString("10000000000000000000000000000000000000"),   // 1e37
		sdk.NewUintFromString("100000000000000000000000000000000000000"),  // 1e38
		sdk.NewUintFromString("1000000000000000000000000000000000000000"), // 1e39
	}

	// Abbreviation: DP stands for 'Decimal Places'

	// ln(2) - used in ln(x). 30 DP.
	LN2 = sdk.NewUintFromString("693147180559945309417232121458")

	// 1 / ln(2) - used in exp(x). 30 DP.
	ONE_OVER_LN2 = sdk.NewUintFromString("1442695040888963407359924681002")

	/***** LOOKUP TABLES *****/

	// Lookup table arrays (LUTs) for log_2(x)
	table_log_2  = make([]sdk.Uint, 100)
	table2_log_2 = make([]sdk.Uint, 100)

	// Lookup table for pow2(). Table contains 39 arrays, each array contains 10 uint slots.
	// The 10 slots are initialised in the init() function.
	table_pow2 = make([][]sdk.Uint, 39)

	ZeroUint = sdk.ZeroUint()
	OneUint  = sdk.OneUint()
	TwoUint  = sdk.NewUint(2)
	FiveUint = sdk.NewUint(5)
	TenUint  = sdk.NewUint(10)

	// LUT flags
	LUT1_isSet   = false
	LUT2_isSet   = false
	LUT3_1_isSet = false
	LUT3_2_isSet = false
	LUT3_3_isSet = false
	LUT3_4_isSet = false
)

func init() {
	for i := range table_pow2 {
		table_pow2[i] = make([]sdk.Uint, 10)
	}
	setLUT1()
	setLUT2()
	setLUT3_1()
	setLUT3_2()
	setLUT3_3()
	setLUT3_4()
}

func requireThat(flag bool, msg string) {
	if !flag {
		panic(msg)
	}
}

/******  BASIC MATH OPERATORS ******/

// Integer math operators. Identical to Zeppelin's SafeMath
func mul(a sdk.Uint, b sdk.Uint) sdk.Uint {
	if a.IsZero() {
		return ZeroUint
	}
	c := a.Mul(b)
	requireThat(c.Quo(a).Equal(b), "uint overflow from multiplication")
	return c
}

func div(a sdk.Uint, b sdk.Uint) sdk.Uint {
	requireThat(b.GT(ZeroUint), "division by zero")
	c := a.Quo(b)
	return c
}

func sub(a sdk.Uint, b sdk.Uint) sdk.Uint {
	requireThat(b.LTE(a), "uint underflow from subtraction")
	c := a.Sub(b)
	return c
}

func add(a sdk.Uint, b sdk.Uint) sdk.Uint {
	c := a.Add(b)
	requireThat(c.GTE(a), "uint overflow from multiplication")
	return c
}

// Basic decimal math operators. Inputs and outputs are uint representations of fixed-point decimals.

// 18 Decimal places
func decMul18(x sdk.Uint, y sdk.Uint) (decProd sdk.Uint) {
	prod_xy := mul(x, y)
	decProd = add(prod_xy, TEN[18]).Quo(TEN[18])
	return
}

func decDiv18(x sdk.Uint, y sdk.Uint) (decQuotient sdk.Uint) {
	prod_xTEN18 := mul(x, TEN[18])
	decQuotient = add(prod_xTEN18, y.Quo(TwoUint)).Quo(y)
	return
}

// 30 Decimal places
func decMul30(x sdk.Uint, y sdk.Uint) (decProd sdk.Uint) {
	prod_xy := mul(x, y)
	decProd = add(prod_xy, TEN[30].Quo(TwoUint)).Quo(TEN[30])
	return
}

// 38 Decimal places
func decMul38(x sdk.Uint, y sdk.Uint) (decProd sdk.Uint) {
	prod_xy := mul(x, y)
	decProd = add(prod_xy, TEN[38].Quo(TwoUint)).Quo(TEN[38])
	return
}

/****** HELPER FUNCTIONS ******/

func convert38To18DP(x sdk.Uint) (y sdk.Uint) {
	digit := (x.Mod(TEN[20])).Quo(TEN[19]) // grab 20th digit from-right
	return chopAndRound(x, digit, TEN[20])
}

func convert38To30DP(x sdk.Uint) (y sdk.Uint) {
	digit := (x.Mod(TEN[8])).Quo(TEN[7]) // grab 8th digit from-right
	return chopAndRound(x, digit, TEN[8])
}

func convert30To20DP(x sdk.Uint) (y sdk.Uint) {
	digit := (x.Mod(TEN[10])).Quo(TEN[9]) // grab 10th digit from-right
	return chopAndRound(x, digit, TEN[10])
}

func convert30To18DP(x sdk.Uint) (y sdk.Uint) {
	digit := (x.Mod(TEN[12])).Quo(TEN[11]) // grab 12th digit from-right
	return chopAndRound(x, digit, TEN[12])
}

// Chop the last digits, and round the resulting number
func chopAndRound(num sdk.Uint, digit sdk.Uint, TENpositionOfChop sdk.Uint) (chopped sdk.Uint) {
	if digit.LT(FiveUint) {
		chopped = div(num, TENpositionOfChop) // round down
	} else {
		chopped = div(num, TENpositionOfChop).Add(OneUint) // round up
	}
	return chopped
}

// return the floor of a fixed-point 20DP number
func floor(x sdk.Uint) (num sdk.Uint) {
	num = x.Sub(x.Mod(TEN[20]))
	return num
}

func countDigits(num sdk.Uint) uint {
	digits := uint(0)

	for !num.IsZero() {
		num = num.Quo(TenUint) // When num < 10, yields 0 due to EVM floor division
		digits++
	}
	return digits
}

/****** MATH FUNCTIONS ******/

// b^x for integer exponent. Use highly performant 'exponentiation-by-squaring' algorithm. O(log(n)) operations.

// b^x - integer base, integer exponent
func powBySquare(x sdk.Uint, n sdk.Uint) sdk.Uint {
	if n.IsZero() {
		return OneUint
	}

	y := OneUint

	for n.GT(OneUint) {
		if n.Mod(TwoUint).IsZero() {
			x = mul(x, x)
			n = n.Quo(TwoUint)
		} else {
			y = mul(x, y)
			x = mul(x, x)
			n = n.Sub(OneUint).Quo(TwoUint)
		}
	}
	return mul(x, y)
}

// b^x - fixed-point 18 DP base, integer exponent
func powBySquare18(base sdk.Uint, n sdk.Uint) sdk.Uint {
	if n.IsZero() {
		return TEN[18]
	}

	y := TEN[18]

	for n.GT(OneUint) {
		if n.Mod(TwoUint).IsZero() {
			base = decMul18(base, base)
			n = n.Quo(TwoUint)
		} else {
			y = decMul18(base, y)
			base = decMul18(base, base)
			n = n.Sub(OneUint).Quo(TwoUint)
		}
	}
	return decMul18(base, y)
}

// b^x - fixed-point 38 DP base, integer exponent n
func powBySquare38(base sdk.Uint, n sdk.Uint) sdk.Uint {
	if n.IsZero() {
		return TEN[38]
	}

	y := TEN[38]

	for n.GT(OneUint) {
		if n.Mod(TwoUint).IsZero() {
			base = decMul38(base, base)
			n = n.Quo(TwoUint)
		} else {
			y = decMul38(base, y)
			base = decMul38(base, base)
			n = n.Sub(OneUint).Quo(TwoUint)
		}
	}
	return decMul38(base, y)
}

/* exp(x) function. Input 18 DP, output 18 DP.  Uses identities:

   A) e^x = 2^(x / ln(2))

   and

   B) 2^y = (2^r) * 2^(y - r); where r = floor(y) - 1, and (y - r) is in range [1,2[

*/
func exp(x sdk.Uint) (num sdk.Uint) {
	var intExponent sdk.Uint // 20 DP
	var decExponent sdk.Uint // 20 DP
	var coefficient sdk.Uint // 38 DP

	x = mul(x, TEN[12]) // make x 30DP
	x = decMul30(ONE_OVER_LN2, x)
	x = convert30To20DP(x)

	if x.LT(TEN[20]) && x.GTE(TwoUint) {
		// if x < 1, do: (2^-1) * 2^(1 + x)
		decExponent = add(TEN[20], x)
		coefficient = TEN[38].Quo(TwoUint)
		num = decMul38(coefficient, pow2(decExponent))
	} else {
		// Use identity B)
		intExponent = floor(x).Sub(TEN[20])
		decExponent = x.Sub(intExponent) // decimal exponent in range [1,2[
		coefficient = powBySquare(TwoUint, div(intExponent, TEN[20]))
		num = mul(coefficient, pow2(decExponent)) //  use normal mul to avoid overflow, as coeff. is an int
	}

	return convert38To18DP(num)
}

// Base-2 logarithm function, for x in range [1,2[. For use in ln(x). Input 18 DP, output 30 DP.
func log_2(x sdk.Uint, accuracy uint) sdk.Uint {
	_onlyLUT1andLUT2AreSet()
	requireThat(x.GTE(TEN[18]) && x.LT(TEN[18].Mul(TwoUint)), "input x must be within range [1,2[")
	prod := mul(x, TEN[20]) // make x 38 DP
	newProd := TEN[38]
	output := ZeroUint

	for i := uint(1); i <= accuracy; i++ {
		newProd = decMul38(table_log_2[i], prod)

		if newProd.GTE(TEN[38]) {
			prod = newProd
			output = output.Add(table2_log_2[i])
		}
	}
	return convert38To30DP(output)
}

// pow2(x) function, for use in exp(x). Uses 2D-array LUT. Valid for x in range [1,2[. Input 20DP, output 38DP
func pow2(x sdk.Uint) sdk.Uint {
	_onlyLUT3isSet()

	requireThat(x.GTE(TEN[20]) && x.LT(TEN[20].Mul(TwoUint)), "input x must be within range [1,2[")
	x_38dp := x.Mul(TEN[18])
	prod := TwoUint.Mul(TEN[38])
	fractPart := x_38dp.Mod(TEN[38])
	digitsLength := countDigits(fractPart)

	// loop backwards through mantissa digits - multiply each by the Lookup-table value
	for i := uint(0); i < digitsLength; i++ {
		digit := (fractPart.Mod(TEN[i+1])).Quo(TEN[i]) // grab the i'th digit from right

		if digit.IsZero() {
			continue // Save gas - skip the step if digit = 0 and there would be no resultant change to prod
		}

		// computer i'th term, and new product
		term := table_pow2[37-i][digit.Uint64()]
		prod = decMul38(prod, term)
	}
	return prod
}

/* Natural log func ln(x). Input 18 DP, output 18 DP. Uses identities:

   A) ln(x) = log_2(x) * ln(2)

   and

   B) log_2(x) = log_2(2^q * y)           y in range [1,2[
               = q + log_2(y)

   The algorithm finds q and y by repeated division by powers-of-two.
*/
func ln(x sdk.Uint, accuracy uint) sdk.Uint {
	requireThat(x.GTE(TEN[18]), "input must be >= 1")
	count := ZeroUint // track
	TWO := mul(TEN[18], TwoUint)

	/* Calculate q. Use branches to divide by powers-of-two, until output is in range [1,2[. Branch approach is more performant
	   than simple successive division by 2. As max input of ln(x) is ~= 2^132, starting division at 2^30 yields sufficiently few operations for large x. */
	for x.GTE(TWO) {
		if x.GTE(sdk.NewUint(1073741824).Mul(TEN[18])) { // start at 2^30
			x = decDiv18(x, sdk.NewUint(1073741824).Mul(TEN[18]))
			count = count.Add(sdk.NewUint(30))
		} else if x.GTE(sdk.NewUint(1048576).Mul(TEN[18])) {
			x = decDiv18(x, sdk.NewUint(1048576).Mul(TEN[18]))
			count = count.Add(sdk.NewUint(20))
		} else if x.GTE(sdk.NewUint(32768).Mul(TEN[18])) {
			x = decDiv18(x, sdk.NewUint(32768).Mul(TEN[18]))
			count = count.Add(sdk.NewUint(15))
		} else if x.GTE(sdk.NewUint(1024).Mul(TEN[18])) {
			x = decDiv18(x, sdk.NewUint(1024).Mul(TEN[18]))
			count = count.Add(sdk.NewUint(10))
		} else if x.GTE(sdk.NewUint(512).Mul(TEN[18])) {
			x = decDiv18(x, sdk.NewUint(512).Mul(TEN[18]))
			count = count.Add(sdk.NewUint(9))
		} else if x.GTE(sdk.NewUint(256).Mul(TEN[18])) {
			x = decDiv18(x, sdk.NewUint(256).Mul(TEN[18]))
			count = count.Add(sdk.NewUint(8))
		} else if x.GTE(sdk.NewUint(128).Mul(TEN[18])) {
			x = decDiv18(x, sdk.NewUint(128).Mul(TEN[18]))
			count = count.Add(sdk.NewUint(7))
		} else if x.GTE(sdk.NewUint(64).Mul(TEN[18])) {
			x = decDiv18(x, sdk.NewUint(64).Mul(TEN[18]))
			count = count.Add(sdk.NewUint(6))
		} else if x.GTE(sdk.NewUint(32).Mul(TEN[18])) {
			x = decDiv18(x, sdk.NewUint(32).Mul(TEN[18]))
			count = count.Add(sdk.NewUint(5))
		} else if x.GTE(sdk.NewUint(16).Mul(TEN[18])) {
			x = decDiv18(x, sdk.NewUint(16).Mul(TEN[18]))
			count = count.Add(sdk.NewUint(4))
		} else if x.GTE(sdk.NewUint(8).Mul(TEN[18])) {
			x = decDiv18(x, sdk.NewUint(8).Mul(TEN[18]))
			count = count.Add(sdk.NewUint(3))
		} else if x.GTE(sdk.NewUint(4).Mul(TEN[18])) {
			x = decDiv18(x, sdk.NewUint(4).Mul(TEN[18]))
			count = count.Add(sdk.NewUint(2))
		} else if x.GTE(sdk.NewUint(2).Mul(TEN[18])) {
			x = decDiv18(x, sdk.NewUint(2).Mul(TEN[18]))
			count = count.Add(sdk.NewUint(1))
		}
	}

	q := count.Mul(TEN[30])
	output := decMul30(LN2, add(q, log_2(x, accuracy)))

	return convert30To18DP(output)
}

/* pow(b, x) func for 18 DP base and exponent. Output 18 DP.

   Uses identity:  b^x = exp (x * ln(b)).

   For b < 1, rewrite b^x as:
   b^x = exp( x * (-ln(1/b)) ) = 1/exp(x * ln(1/b)).

    Thus, we avoid passing argument y < 1 to ln(y), and z < 0 to exp(z).   */
func pow(base sdk.Uint, x sdk.Uint) (power sdk.Uint) {
	if base.IsZero() {
		if x.IsZero() {
			panic("Zero raised to zero is undefined")
		}
		return ZeroUint // 0^p = 0 (for p > 0)
	} else if x.IsZero() {
		return TEN[18] // b^0 = 1
	} else if base.Equal(TEN[18]) {
		return TEN[18] // 1^x = 1
	} else if base.GT(TEN[18]) {
		return exp(decMul18(x, ln(base, 70)))
	} else { // base.LT(TEN[18])
		exponent := decMul18(x, ln(decDiv18(TEN[18], base), 70))
		return decDiv18(TEN[18], exp(exponent))
	}
}

// Taylor series implementation of exp(x) - lower accuracy and higher gas cost than exp(x). 18 DP input and output.
func exp_taylor(x sdk.Uint) sdk.Uint {
	tolerance := OneUint
	term := TEN[18]
	sum := TEN[18]
	i := ZeroUint

	for term.GT(tolerance) {
		i = i.Add(TEN[18])
		term = decDiv18(decMul18(term, x), i)
		sum = sum.Add(term)
	}
	return sum
}

/* Lookup Tables (LUTs). 38 DP fixed-point numbers. */

// LUT1 for log_2(x). The i'th term is 1/(2^(1/2^i))
func setLUT1() {
	table_log_2[0] = sdk.NewUintFromString("0")
	table_log_2[1] = sdk.NewUintFromString("70710678118654752440084436210484903928")
	table_log_2[2] = sdk.NewUintFromString("84089641525371454303112547623321489504")
	table_log_2[3] = sdk.NewUintFromString("91700404320467123174354159479414442804")
	table_log_2[4] = sdk.NewUintFromString("95760328069857364693630563514791544393")
	table_log_2[5] = sdk.NewUintFromString("97857206208770013450916112581343574560")
	table_log_2[6] = sdk.NewUintFromString("98922801319397548412912495906558366777")
	table_log_2[7] = sdk.NewUintFromString("99459942348363317565247768622216631446")
	table_log_2[8] = sdk.NewUintFromString("99729605608547012625765991384792260112")
	table_log_2[9] = sdk.NewUintFromString("99864711289097017358812131808592040806")
	table_log_2[10] = sdk.NewUintFromString("99932332750265075236028365984373804116")
	table_log_2[11] = sdk.NewUintFromString("99966160649624368394219686876281565561")
	table_log_2[12] = sdk.NewUintFromString("99983078893192906311748078019767389868")
	table_log_2[13] = sdk.NewUintFromString("99991539088661349753372497156418872723")
	table_log_2[14] = sdk.NewUintFromString("99995769454843113254396753730099797524")
	table_log_2[15] = sdk.NewUintFromString("99997884705049192982650067113039327478")
	table_log_2[16] = sdk.NewUintFromString("99998942346931446424221059225315431670")
	table_log_2[17] = sdk.NewUintFromString("99999471172067428300770241277030532519")
	table_log_2[18] = sdk.NewUintFromString("99999735585684139498225234636504270993")
	table_log_2[19] = sdk.NewUintFromString("99999867792754675970531776759801063698")
	table_log_2[20] = sdk.NewUintFromString("99999933896355489526178052900624509795")
	table_log_2[21] = sdk.NewUintFromString("99999966948172282646511738368820575117")
	table_log_2[22] = sdk.NewUintFromString("99999983474084775793885880947314828005")
	table_log_2[23] = sdk.NewUintFromString("99999991737042046514572235133214264694")
	table_log_2[24] = sdk.NewUintFromString("99999995868520937911689915196095249000")
	table_log_2[25] = sdk.NewUintFromString("99999997934260447619445466250978583193")
	table_log_2[26] = sdk.NewUintFromString("99999998967130218475622805194415901619")
	table_log_2[27] = sdk.NewUintFromString("99999999483565107904286413727651274869")
	table_log_2[28] = sdk.NewUintFromString("99999999741782553618761958785587923503")
	table_log_2[29] = sdk.NewUintFromString("99999999870891276726035667265628464908")
	table_log_2[30] = sdk.NewUintFromString("99999999935445638342181505587572099682")
	table_log_2[31] = sdk.NewUintFromString("99999999967722819165881670780794171827")
	table_log_2[32] = sdk.NewUintFromString("99999999983861409581638564886938948308")
	table_log_2[33] = sdk.NewUintFromString("99999999991930704790493714817578668739")
	table_log_2[34] = sdk.NewUintFromString("99999999995965352395165465502313349139")
	table_log_2[35] = sdk.NewUintFromString("99999999997982676197562384774537267778")
	table_log_2[36] = sdk.NewUintFromString("99999999998991338098776105393113730880")
	table_log_2[37] = sdk.NewUintFromString("99999999999495669049386780948018133274")
	table_log_2[38] = sdk.NewUintFromString("99999999999747834524693072536874382794")
	table_log_2[39] = sdk.NewUintFromString("99999999999873917262346456784153520336")
	table_log_2[40] = sdk.NewUintFromString("99999999999936958631173208521005842390")
	table_log_2[41] = sdk.NewUintFromString("99999999999968479315586599292735191749")
	table_log_2[42] = sdk.NewUintFromString("99999999999984239657793298404425663513")
	table_log_2[43] = sdk.NewUintFromString("99999999999992119828896648891727348666")
	table_log_2[44] = sdk.NewUintFromString("99999999999996059914448324368242303560")
	table_log_2[45] = sdk.NewUintFromString("99999999999998029957224162164715809087")
	table_log_2[46] = sdk.NewUintFromString("99999999999999014978612081077506568870")
	table_log_2[47] = sdk.NewUintFromString("99999999999999507489306040537540450517")
	table_log_2[48] = sdk.NewUintFromString("99999999999999753744653020268467016779")
	table_log_2[49] = sdk.NewUintFromString("99999999999999876872326510134157706270")
	table_log_2[50] = sdk.NewUintFromString("99999999999999938436163255067059902605")
	table_log_2[51] = sdk.NewUintFromString("99999999999999969218081627533525213670")
	table_log_2[52] = sdk.NewUintFromString("99999999999999984609040813766761422427")
	table_log_2[53] = sdk.NewUintFromString("99999999999999992304520406883380415111")
	table_log_2[54] = sdk.NewUintFromString("99999999999999996152260203441690133530")
	table_log_2[55] = sdk.NewUintFromString("99999999999999998076130101720845048259")
	table_log_2[56] = sdk.NewUintFromString("99999999999999999038065050860422519503")
	table_log_2[57] = sdk.NewUintFromString("99999999999999999519032525430211258595")
	table_log_2[58] = sdk.NewUintFromString("99999999999999999759516262715105629008")
	table_log_2[59] = sdk.NewUintFromString("99999999999999999879758131357552814432")
	table_log_2[60] = sdk.NewUintFromString("99999999999999999939879065678776407198")
	table_log_2[61] = sdk.NewUintFromString("99999999999999999969939532839388203594")
	table_log_2[62] = sdk.NewUintFromString("99999999999999999984969766419694101796")
	table_log_2[63] = sdk.NewUintFromString("99999999999999999992484883209847050898")
	table_log_2[64] = sdk.NewUintFromString("99999999999999999996242441604923525449")
	table_log_2[65] = sdk.NewUintFromString("99999999999999999998121220802461762724")
	table_log_2[66] = sdk.NewUintFromString("99999999999999999999060610401230881362")
	table_log_2[67] = sdk.NewUintFromString("99999999999999999999530305200615440681")
	table_log_2[68] = sdk.NewUintFromString("99999999999999999999765152600307720341")
	table_log_2[69] = sdk.NewUintFromString("99999999999999999999882576300153860170")
	table_log_2[70] = sdk.NewUintFromString("99999999999999999999941288150076930085")
	table_log_2[71] = sdk.NewUintFromString("99999999999999999999970644075038465043")
	table_log_2[72] = sdk.NewUintFromString("99999999999999999999985322037519232521")
	table_log_2[73] = sdk.NewUintFromString("99999999999999999999992661018759616261")
	table_log_2[74] = sdk.NewUintFromString("99999999999999999999996330509379808130")
	table_log_2[75] = sdk.NewUintFromString("99999999999999999999998165254689904065")
	table_log_2[76] = sdk.NewUintFromString("99999999999999999999999082627344952033")
	table_log_2[77] = sdk.NewUintFromString("99999999999999999999999541313672476016")
	table_log_2[78] = sdk.NewUintFromString("99999999999999999999999770656836238008")
	table_log_2[79] = sdk.NewUintFromString("99999999999999999999999885328418119004")
	table_log_2[80] = sdk.NewUintFromString("99999999999999999999999942664209059502")
	table_log_2[81] = sdk.NewUintFromString("99999999999999999999999971332104529751")
	table_log_2[82] = sdk.NewUintFromString("99999999999999999999999985666052264876")
	table_log_2[83] = sdk.NewUintFromString("99999999999999999999999992833026132438")
	table_log_2[84] = sdk.NewUintFromString("99999999999999999999999996416513066219")
	table_log_2[85] = sdk.NewUintFromString("99999999999999999999999998208256533109")
	table_log_2[86] = sdk.NewUintFromString("99999999999999999999999999104128266555")
	table_log_2[87] = sdk.NewUintFromString("99999999999999999999999999552064133277")
	table_log_2[88] = sdk.NewUintFromString("99999999999999999999999999776032066639")
	table_log_2[89] = sdk.NewUintFromString("99999999999999999999999999888016033319")
	table_log_2[90] = sdk.NewUintFromString("99999999999999999999999999944008016660")
	table_log_2[91] = sdk.NewUintFromString("99999999999999999999999999972004008330")
	table_log_2[92] = sdk.NewUintFromString("99999999999999999999999999986002004165")
	table_log_2[93] = sdk.NewUintFromString("99999999999999999999999999993001002082")
	table_log_2[94] = sdk.NewUintFromString("99999999999999999999999999996500501041")
	table_log_2[95] = sdk.NewUintFromString("99999999999999999999999999998250250521")
	table_log_2[96] = sdk.NewUintFromString("99999999999999999999999999999125125260")
	table_log_2[97] = sdk.NewUintFromString("99999999999999999999999999999562562630")
	table_log_2[98] = sdk.NewUintFromString("99999999999999999999999999999781281315")
	table_log_2[99] = sdk.NewUintFromString("99999999999999999999999999999890640658")

	LUT1_isSet = true
}

// LUT2 for log_2(x). The i'th term is 1/(2^i)
func setLUT2() {
	table2_log_2[0] = sdk.NewUintFromString("200000000000000000000000000000000000000")
	table2_log_2[1] = sdk.NewUintFromString("50000000000000000000000000000000000000")
	table2_log_2[2] = sdk.NewUintFromString("25000000000000000000000000000000000000")
	table2_log_2[3] = sdk.NewUintFromString("12500000000000000000000000000000000000")
	table2_log_2[4] = sdk.NewUintFromString("6250000000000000000000000000000000000")
	table2_log_2[5] = sdk.NewUintFromString("3125000000000000000000000000000000000")
	table2_log_2[6] = sdk.NewUintFromString("1562500000000000000000000000000000000")
	table2_log_2[7] = sdk.NewUintFromString("781250000000000000000000000000000000")
	table2_log_2[8] = sdk.NewUintFromString("390625000000000000000000000000000000")
	table2_log_2[9] = sdk.NewUintFromString("195312500000000000000000000000000000")
	table2_log_2[10] = sdk.NewUintFromString("97656250000000000000000000000000000")
	table2_log_2[11] = sdk.NewUintFromString("48828125000000000000000000000000000")
	table2_log_2[12] = sdk.NewUintFromString("24414062500000000000000000000000000")
	table2_log_2[13] = sdk.NewUintFromString("12207031250000000000000000000000000")
	table2_log_2[14] = sdk.NewUintFromString("6103515625000000000000000000000000")
	table2_log_2[15] = sdk.NewUintFromString("3051757812500000000000000000000000")
	table2_log_2[16] = sdk.NewUintFromString("1525878906250000000000000000000000")
	table2_log_2[17] = sdk.NewUintFromString("762939453125000000000000000000000")
	table2_log_2[18] = sdk.NewUintFromString("381469726562500000000000000000000")
	table2_log_2[19] = sdk.NewUintFromString("190734863281250000000000000000000")
	table2_log_2[20] = sdk.NewUintFromString("95367431640625000000000000000000")
	table2_log_2[21] = sdk.NewUintFromString("47683715820312500000000000000000")
	table2_log_2[22] = sdk.NewUintFromString("23841857910156250000000000000000")
	table2_log_2[23] = sdk.NewUintFromString("11920928955078125000000000000000")
	table2_log_2[24] = sdk.NewUintFromString("5960464477539062500000000000000")
	table2_log_2[25] = sdk.NewUintFromString("2980232238769531250000000000000")
	table2_log_2[26] = sdk.NewUintFromString("1490116119384765625000000000000")
	table2_log_2[27] = sdk.NewUintFromString("745058059692382812500000000000")
	table2_log_2[28] = sdk.NewUintFromString("372529029846191406250000000000")
	table2_log_2[29] = sdk.NewUintFromString("186264514923095703125000000000")
	table2_log_2[30] = sdk.NewUintFromString("93132257461547851562500000000")
	table2_log_2[31] = sdk.NewUintFromString("46566128730773925781250000000")
	table2_log_2[32] = sdk.NewUintFromString("23283064365386962890625000000")
	table2_log_2[33] = sdk.NewUintFromString("11641532182693481445312500000")
	table2_log_2[34] = sdk.NewUintFromString("5820766091346740722656250000")
	table2_log_2[35] = sdk.NewUintFromString("2910383045673370361328125000")
	table2_log_2[36] = sdk.NewUintFromString("1455191522836685180664062500")
	table2_log_2[37] = sdk.NewUintFromString("727595761418342590332031250")
	table2_log_2[38] = sdk.NewUintFromString("363797880709171295166015625")
	table2_log_2[39] = sdk.NewUintFromString("181898940354585647583007812")
	table2_log_2[40] = sdk.NewUintFromString("90949470177292823791503906")
	table2_log_2[41] = sdk.NewUintFromString("45474735088646411895751953")
	table2_log_2[42] = sdk.NewUintFromString("22737367544323205947875976")
	table2_log_2[43] = sdk.NewUintFromString("11368683772161602973937988")
	table2_log_2[44] = sdk.NewUintFromString("5684341886080801486968994")
	table2_log_2[45] = sdk.NewUintFromString("2842170943040400743484497")
	table2_log_2[46] = sdk.NewUintFromString("1421085471520200371742248")
	table2_log_2[47] = sdk.NewUintFromString("710542735760100185871124")
	table2_log_2[48] = sdk.NewUintFromString("355271367880050092935562")
	table2_log_2[49] = sdk.NewUintFromString("177635683940025046467781")
	table2_log_2[50] = sdk.NewUintFromString("88817841970012523233890")
	table2_log_2[51] = sdk.NewUintFromString("44408920985006261616945")
	table2_log_2[52] = sdk.NewUintFromString("22204460492503130808472")
	table2_log_2[53] = sdk.NewUintFromString("11102230246251565404236")
	table2_log_2[54] = sdk.NewUintFromString("5551115123125782702118")
	table2_log_2[55] = sdk.NewUintFromString("2775557561562891351059")
	table2_log_2[56] = sdk.NewUintFromString("1387778780781445675529")
	table2_log_2[57] = sdk.NewUintFromString("693889390390722837764")
	table2_log_2[58] = sdk.NewUintFromString("346944695195361418882")
	table2_log_2[59] = sdk.NewUintFromString("173472347597680709441")
	table2_log_2[60] = sdk.NewUintFromString("86736173798840354720")
	table2_log_2[61] = sdk.NewUintFromString("43368086899420177360")
	table2_log_2[62] = sdk.NewUintFromString("21684043449710088680")
	table2_log_2[63] = sdk.NewUintFromString("10842021724855044340")
	table2_log_2[64] = sdk.NewUintFromString("5421010862427522170")
	table2_log_2[65] = sdk.NewUintFromString("2710505431213761085")
	table2_log_2[66] = sdk.NewUintFromString("1355252715606880542")
	table2_log_2[67] = sdk.NewUintFromString("677626357803440271")
	table2_log_2[68] = sdk.NewUintFromString("338813178901720135")
	table2_log_2[69] = sdk.NewUintFromString("169406589450860067")
	table2_log_2[70] = sdk.NewUintFromString("84703294725430033")
	table2_log_2[71] = sdk.NewUintFromString("42351647362715016")
	table2_log_2[72] = sdk.NewUintFromString("21175823681357508")
	table2_log_2[73] = sdk.NewUintFromString("10587911840678754")
	table2_log_2[74] = sdk.NewUintFromString("5293955920339377")
	table2_log_2[75] = sdk.NewUintFromString("2646977960169688")
	table2_log_2[76] = sdk.NewUintFromString("1323488980084844")
	table2_log_2[77] = sdk.NewUintFromString("661744490042422")
	table2_log_2[78] = sdk.NewUintFromString("330872245021211")
	table2_log_2[79] = sdk.NewUintFromString("165436122510605")
	table2_log_2[80] = sdk.NewUintFromString("82718061255302")
	table2_log_2[81] = sdk.NewUintFromString("41359030627651")
	table2_log_2[82] = sdk.NewUintFromString("20679515313825")
	table2_log_2[83] = sdk.NewUintFromString("10339757656912")
	table2_log_2[84] = sdk.NewUintFromString("5169878828456")
	table2_log_2[85] = sdk.NewUintFromString("2584939414228")
	table2_log_2[86] = sdk.NewUintFromString("1292469707114")
	table2_log_2[87] = sdk.NewUintFromString("646234853557")
	table2_log_2[88] = sdk.NewUintFromString("323117426778")
	table2_log_2[89] = sdk.NewUintFromString("161558713389")
	table2_log_2[90] = sdk.NewUintFromString("80779356694")
	table2_log_2[91] = sdk.NewUintFromString("40389678347")
	table2_log_2[92] = sdk.NewUintFromString("20194839173")
	table2_log_2[93] = sdk.NewUintFromString("10097419586")
	table2_log_2[94] = sdk.NewUintFromString("5048709793")
	table2_log_2[95] = sdk.NewUintFromString("2524354896")
	table2_log_2[96] = sdk.NewUintFromString("1262177448")
	table2_log_2[97] = sdk.NewUintFromString("631088724")
	table2_log_2[98] = sdk.NewUintFromString("315544362")
	table2_log_2[99] = sdk.NewUintFromString("157772181")

	LUT2_isSet = true
}

/* LUT for pow2() function. Table contains 39 arrays, each array contains 10 uint slots.

   table_pow2[i][d] = (2^(1 / 10^(i + 1))) ** d.
   d ranges from 0 to 9.

   LUT-setting is split into four separate setter functions to keep gas costs under block limit.
*/
func setLUT3_1() {
	table_pow2[0][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[0][1] = sdk.NewUintFromString("107177346253629316421300632502334202291")
	table_pow2[0][2] = sdk.NewUintFromString("114869835499703500679862694677792758944")
	table_pow2[0][3] = sdk.NewUintFromString("123114441334491628449939306916774310988")
	table_pow2[0][4] = sdk.NewUintFromString("131950791077289425937400197122964013303")
	table_pow2[0][5] = sdk.NewUintFromString("141421356237309504880168872420969807857")
	table_pow2[0][6] = sdk.NewUintFromString("151571656651039808234725980130644523868")
	table_pow2[0][7] = sdk.NewUintFromString("162450479271247104521941876555056330257")
	table_pow2[0][8] = sdk.NewUintFromString("174110112659224827827254003495949219796")
	table_pow2[0][9] = sdk.NewUintFromString("186606598307361483196268653229988433405")
	table_pow2[1][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[1][1] = sdk.NewUintFromString("100695555005671880883269821411323978545")
	table_pow2[1][2] = sdk.NewUintFromString("101395947979002913869016599962823042584")
	table_pow2[1][3] = sdk.NewUintFromString("102101212570719324976409517478306437354")
	table_pow2[1][4] = sdk.NewUintFromString("102811382665606650934634495879263497655")
	table_pow2[1][5] = sdk.NewUintFromString("103526492384137750434778819421124619773")
	table_pow2[1][6] = sdk.NewUintFromString("104246576084112139095471141872690784007")
	table_pow2[1][7] = sdk.NewUintFromString("104971668362306726904934732174028851665")
	table_pow2[1][8] = sdk.NewUintFromString("105701804056138037449949421408611430989")
	table_pow2[1][9] = sdk.NewUintFromString("106437018245335988793865835140404338206")
	table_pow2[2][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[2][1] = sdk.NewUintFromString("100069338746258063253756863930385919571")
	table_pow2[2][2] = sdk.NewUintFromString("100138725571133452908322477441877746756")
	table_pow2[2][3] = sdk.NewUintFromString("100208160507963279436035132489114568295")
	table_pow2[2][4] = sdk.NewUintFromString("100277643590107768843673305907248072983")
	table_pow2[2][5] = sdk.NewUintFromString("100347174850950278700477431086959080340")
	table_pow2[2][6] = sdk.NewUintFromString("100416754323897314177285298995922943429")
	table_pow2[2][7] = sdk.NewUintFromString("100486382042378544096788794597976421668")
	table_pow2[2][8] = sdk.NewUintFromString("100556058039846816994919680064517944020")
	table_pow2[2][9] = sdk.NewUintFromString("100625782349778177193372141519657470417")
	table_pow2[3][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[3][1] = sdk.NewUintFromString("100006931712037656919243991260264256542")
	table_pow2[3][2] = sdk.NewUintFromString("100013863904561631568466376833067115945")
	table_pow2[3][3] = sdk.NewUintFromString("100020796577605229875592540103010552992")
	table_pow2[3][4] = sdk.NewUintFromString("100027729731201760077218879711834041246")
	table_pow2[3][5] = sdk.NewUintFromString("100034663365384532718772839985089028270")
	table_pow2[3][6] = sdk.NewUintFromString("100041597480186860654672952451661760537")
	table_pow2[3][7] = sdk.NewUintFromString("100048532075642059048488888456913382370")
	table_pow2[3][8] = sdk.NewUintFromString("100055467151783445373101522870206286485")
	table_pow2[3][9] = sdk.NewUintFromString("100062402708644339410863008887585747065")
	table_pow2[4][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[4][1] = sdk.NewUintFromString("100000693149582830565320908980056168150")
	table_pow2[4][2] = sdk.NewUintFromString("100001386303970224572423685307245831542")
	table_pow2[4][3] = sdk.NewUintFromString("100002079463162215324119782522433627138")
	table_pow2[4][4] = sdk.NewUintFromString("100002772627158836123451492465145260129")
	table_pow2[4][5] = sdk.NewUintFromString("100003465795960120273691946873622208121")
	table_pow2[4][6] = sdk.NewUintFromString("100004158969566101078345118984887516084")
	table_pow2[4][7] = sdk.NewUintFromString("100004852147976811841145825134822682163")
	table_pow2[4][8] = sdk.NewUintFromString("100005545331192285866059726358255634403")
	table_pow2[4][9] = sdk.NewUintFromString("100006238519212556457283329989059798485")
	table_pow2[5][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[5][1] = sdk.NewUintFromString("100000069314742078650777263622740703038")
	table_pow2[5][2] = sdk.NewUintFromString("100000138629532202636248826052225815048")
	table_pow2[5][3] = sdk.NewUintFromString("100000207944370371989717187112633071811")
	table_pow2[5][4] = sdk.NewUintFromString("100000277259256586744484869711682067979")
	table_pow2[5][5] = sdk.NewUintFromString("100000346574190846933854419840650257373")
	table_pow2[5][6] = sdk.NewUintFromString("100000415889173152591128406574388953292")
	table_pow2[5][7] = sdk.NewUintFromString("100000485204203503749609422071339328833")
	table_pow2[5][8] = sdk.NewUintFromString("100000554519281900442600081573548417222")
	table_pow2[5][9] = sdk.NewUintFromString("100000623834408342703403023406685112154")
	table_pow2[6][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[6][1] = sdk.NewUintFromString("100000006931472045825965603683996211583")
	table_pow2[6][2] = sdk.NewUintFromString("100000013862944572104978428035962521332")
	table_pow2[6][3] = sdk.NewUintFromString("100000020794417578837071775524560348874")
	table_pow2[6][4] = sdk.NewUintFromString("100000027725891066022278948620759465140")
	table_pow2[6][5] = sdk.NewUintFromString("100000034657365033660633249797837992529")
	table_pow2[6][6] = sdk.NewUintFromString("100000041588839481752167981531382405066")
	table_pow2[6][7] = sdk.NewUintFromString("100000048520314410296916446299287528561")
	table_pow2[6][8] = sdk.NewUintFromString("100000055451789819294911946581756540768")
	table_pow2[6][9] = sdk.NewUintFromString("100000062383265708746187784861300971552")
	table_pow2[7][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[7][1] = sdk.NewUintFromString("100000000693147182962210384558650120894")
	table_pow2[7][2] = sdk.NewUintFromString("100000001386294370728950941601779822006")
	table_pow2[7][3] = sdk.NewUintFromString("100000002079441563300221704431854648481")
	table_pow2[7][4] = sdk.NewUintFromString("100000002772588760676022706351340376300")
	table_pow2[7][5] = sdk.NewUintFromString("100000003465735962856353980662703012279")
	table_pow2[7][6] = sdk.NewUintFromString("100000004158883169841215560668408794069")
	table_pow2[7][7] = sdk.NewUintFromString("100000004852030381630607479670924190156")
	table_pow2[7][8] = sdk.NewUintFromString("100000005545177598224529770972715899860")
	table_pow2[7][9] = sdk.NewUintFromString("100000006238324819622982467876250853339")
	table_pow2[8][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[8][1] = sdk.NewUintFromString("100000000069314718080017181643183694247")
	table_pow2[8][2] = sdk.NewUintFromString("100000000138629436208079664711489996172")
	table_pow2[8][3] = sdk.NewUintFromString("100000000207944154384187449238221371011")
	table_pow2[8][4] = sdk.NewUintFromString("100000000277258872608340535256680284018")
	table_pow2[8][5] = sdk.NewUintFromString("100000000346573590880538922800169200474")
	table_pow2[8][6] = sdk.NewUintFromString("100000000415888309200782611901990585682")
	table_pow2[8][7] = sdk.NewUintFromString("100000000485203027569071602595446904968")
	table_pow2[8][8] = sdk.NewUintFromString("100000000554517745985405894913840623680")
	table_pow2[8][9] = sdk.NewUintFromString("100000000623832464449785488890474207190")
	table_pow2[9][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[9][1] = sdk.NewUintFromString("100000000006931471805839679601136972338")
	table_pow2[9][2] = sdk.NewUintFromString("100000000013862943612159812216225448565")
	table_pow2[9][3] = sdk.NewUintFromString("100000000020794415418960397845298731148")
	table_pow2[9][4] = sdk.NewUintFromString("100000000027725887226241436488390122551")
	table_pow2[9][5] = sdk.NewUintFromString("100000000034657359034002928145532925240")
	table_pow2[9][6] = sdk.NewUintFromString("100000000041588830842244872816760441679")
	table_pow2[9][7] = sdk.NewUintFromString("100000000048520302650967270502105974334")
	table_pow2[9][8] = sdk.NewUintFromString("100000000055451774460170121201602825670")
	table_pow2[9][9] = sdk.NewUintFromString("100000000062383246269853424915284298153")
	table_pow2[10][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[10][1] = sdk.NewUintFromString("100000000000693147180562347574486828679")
	table_pow2[10][2] = sdk.NewUintFromString("100000000001386294361129499679112872675")
	table_pow2[10][3] = sdk.NewUintFromString("100000000002079441541701456313878165290")
	table_pow2[10][4] = sdk.NewUintFromString("100000000002772588722278217478782739826")
	table_pow2[10][5] = sdk.NewUintFromString("100000000003465735902859783173826629587")
	table_pow2[10][6] = sdk.NewUintFromString("100000000004158883083446153399009867874")
	table_pow2[10][7] = sdk.NewUintFromString("100000000004852030264037328154332487990")
	table_pow2[10][8] = sdk.NewUintFromString("100000000005545177444633307439794523238")
	table_pow2[10][9] = sdk.NewUintFromString("100000000006238324625234091255396006920")

	LUT3_1_isSet = true
}

func setLUT3_2() {
	table_pow2[11][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[11][1] = sdk.NewUintFromString("100000000000069314718056018553592419128")
	table_pow2[11][2] = sdk.NewUintFromString("100000000000138629436112085152486230109")
	table_pow2[11][3] = sdk.NewUintFromString("100000000000207944154168199796681432977")
	table_pow2[11][4] = sdk.NewUintFromString("100000000000277258872224362486178027765")
	table_pow2[11][5] = sdk.NewUintFromString("100000000000346573590280573220976014506")
	table_pow2[11][6] = sdk.NewUintFromString("100000000000415888308336832001075393234")
	table_pow2[11][7] = sdk.NewUintFromString("100000000000485203026393138826476163982")
	table_pow2[11][8] = sdk.NewUintFromString("100000000000554517744449493697178326784")
	table_pow2[11][9] = sdk.NewUintFromString("100000000000623832462505896613181881671")
	table_pow2[12][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[12][1] = sdk.NewUintFromString("100000000000006931471805599693320679280")
	table_pow2[12][2] = sdk.NewUintFromString("100000000000013862943611199867094372479")
	table_pow2[12][3] = sdk.NewUintFromString("100000000000020794415416800521321079596")
	table_pow2[12][4] = sdk.NewUintFromString("100000000000027725887222401656000800631")
	table_pow2[12][5] = sdk.NewUintFromString("100000000000034657359028003271133535584")
	table_pow2[12][6] = sdk.NewUintFromString("100000000000041588830833605366719284456")
	table_pow2[12][7] = sdk.NewUintFromString("100000000000048520302639207942758047246")
	table_pow2[12][8] = sdk.NewUintFromString("100000000000055451774444810999249823955")
	table_pow2[12][9] = sdk.NewUintFromString("100000000000062383246250414536194614582")
	table_pow2[13][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[13][1] = sdk.NewUintFromString("100000000000000693147180559947711682302")
	table_pow2[13][2] = sdk.NewUintFromString("100000000000001386294361119900227894743")
	table_pow2[13][3] = sdk.NewUintFromString("100000000000002079441541679857548637323")
	table_pow2[13][4] = sdk.NewUintFromString("100000000000002772588722239819673910042")
	table_pow2[13][5] = sdk.NewUintFromString("100000000000003465735902799786603712900")
	table_pow2[13][6] = sdk.NewUintFromString("100000000000004158883083359758338045898")
	table_pow2[13][7] = sdk.NewUintFromString("100000000000004852030263919734876909035")
	table_pow2[13][8] = sdk.NewUintFromString("100000000000005545177444479716220302311")
	table_pow2[13][9] = sdk.NewUintFromString("100000000000006238324625039702368225726")
	table_pow2[14][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[14][1] = sdk.NewUintFromString("100000000000000069314718055994554964374")
	table_pow2[14][2] = sdk.NewUintFromString("100000000000000138629436111989157974049")
	table_pow2[14][3] = sdk.NewUintFromString("100000000000000207944154167983809029026")
	table_pow2[14][4] = sdk.NewUintFromString("100000000000000277258872223978508129304")
	table_pow2[14][5] = sdk.NewUintFromString("100000000000000346573590279973255274883")
	table_pow2[14][6] = sdk.NewUintFromString("100000000000000415888308335968050465764")
	table_pow2[14][7] = sdk.NewUintFromString("100000000000000485203026391962893701947")
	table_pow2[14][8] = sdk.NewUintFromString("100000000000000554517744447957784983430")
	table_pow2[14][9] = sdk.NewUintFromString("100000000000000623832462503952724310215")
	table_pow2[15][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[15][1] = sdk.NewUintFromString("100000000000000006931471805599453334399")
	table_pow2[15][2] = sdk.NewUintFromString("100000000000000013862943611198907149251")
	table_pow2[15][3] = sdk.NewUintFromString("100000000000000020794415416798361444556")
	table_pow2[15][4] = sdk.NewUintFromString("100000000000000027725887222397816220313")
	table_pow2[15][5] = sdk.NewUintFromString("100000000000000034657359027997271476524")
	table_pow2[15][6] = sdk.NewUintFromString("100000000000000041588830833596727213188")
	table_pow2[15][7] = sdk.NewUintFromString("100000000000000048520302639196183430305")
	table_pow2[15][8] = sdk.NewUintFromString("100000000000000055451774444795640127875")
	table_pow2[15][9] = sdk.NewUintFromString("100000000000000062383246250395097305898")
	table_pow2[16][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[16][1] = sdk.NewUintFromString("100000000000000000693147180559945311819")
	table_pow2[16][2] = sdk.NewUintFromString("100000000000000001386294361119890628444")
	table_pow2[16][3] = sdk.NewUintFromString("100000000000000002079441541679835949872")
	table_pow2[16][4] = sdk.NewUintFromString("100000000000000002772588722239781276105")
	table_pow2[16][5] = sdk.NewUintFromString("100000000000000003465735902799726607143")
	table_pow2[16][6] = sdk.NewUintFromString("100000000000000004158883083359671942985")
	table_pow2[16][7] = sdk.NewUintFromString("100000000000000004852030263919617283632")
	table_pow2[16][8] = sdk.NewUintFromString("100000000000000005545177444479562629083")
	table_pow2[16][9] = sdk.NewUintFromString("100000000000000006238324625039507979339")
	table_pow2[17][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[17][1] = sdk.NewUintFromString("100000000000000000069314718055994530966")
	table_pow2[17][2] = sdk.NewUintFromString("100000000000000000138629436111989061980")
	table_pow2[17][3] = sdk.NewUintFromString("100000000000000000207944154167983593041")
	table_pow2[17][4] = sdk.NewUintFromString("100000000000000000277258872223978124151")
	table_pow2[17][5] = sdk.NewUintFromString("100000000000000000346573590279972655309")
	table_pow2[17][6] = sdk.NewUintFromString("100000000000000000415888308335967186515")
	table_pow2[17][7] = sdk.NewUintFromString("100000000000000000485203026391961717769")
	table_pow2[17][8] = sdk.NewUintFromString("100000000000000000554517744447956249071")
	table_pow2[17][9] = sdk.NewUintFromString("100000000000000000623832462503950780421")
	table_pow2[18][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[18][1] = sdk.NewUintFromString("100000000000000000006931471805599453094")
	table_pow2[18][2] = sdk.NewUintFromString("100000000000000000013862943611198906189")
	table_pow2[18][3] = sdk.NewUintFromString("100000000000000000020794415416798359285")
	table_pow2[18][4] = sdk.NewUintFromString("100000000000000000027725887222397812381")
	table_pow2[18][5] = sdk.NewUintFromString("100000000000000000034657359027997265477")
	table_pow2[18][6] = sdk.NewUintFromString("100000000000000000041588830833596718574")
	table_pow2[18][7] = sdk.NewUintFromString("100000000000000000048520302639196171671")
	table_pow2[18][8] = sdk.NewUintFromString("100000000000000000055451774444795624769")
	table_pow2[18][9] = sdk.NewUintFromString("100000000000000000062383246250395077867")
	table_pow2[19][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[19][1] = sdk.NewUintFromString("100000000000000000000693147180559945309")
	table_pow2[19][2] = sdk.NewUintFromString("100000000000000000001386294361119890619")
	table_pow2[19][3] = sdk.NewUintFromString("100000000000000000002079441541679835928")
	table_pow2[19][4] = sdk.NewUintFromString("100000000000000000002772588722239781238")
	table_pow2[19][5] = sdk.NewUintFromString("100000000000000000003465735902799726547")
	table_pow2[19][6] = sdk.NewUintFromString("100000000000000000004158883083359671857")
	table_pow2[19][7] = sdk.NewUintFromString("100000000000000000004852030263919617166")
	table_pow2[19][8] = sdk.NewUintFromString("100000000000000000005545177444479562475")
	table_pow2[19][9] = sdk.NewUintFromString("100000000000000000006238324625039507785")
	table_pow2[20][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[20][1] = sdk.NewUintFromString("100000000000000000000069314718055994531")
	table_pow2[20][2] = sdk.NewUintFromString("100000000000000000000138629436111989062")
	table_pow2[20][3] = sdk.NewUintFromString("100000000000000000000207944154167983593")
	table_pow2[20][4] = sdk.NewUintFromString("100000000000000000000277258872223978124")
	table_pow2[20][5] = sdk.NewUintFromString("100000000000000000000346573590279972655")
	table_pow2[20][6] = sdk.NewUintFromString("100000000000000000000415888308335967186")
	table_pow2[20][7] = sdk.NewUintFromString("100000000000000000000485203026391961717")
	table_pow2[20][8] = sdk.NewUintFromString("100000000000000000000554517744447956248")
	table_pow2[20][9] = sdk.NewUintFromString("100000000000000000000623832462503950778")

	LUT3_2_isSet = true
}

func setLUT3_3() {
	table_pow2[21][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[21][1] = sdk.NewUintFromString("100000000000000000000006931471805599453")
	table_pow2[21][2] = sdk.NewUintFromString("100000000000000000000013862943611198906")
	table_pow2[21][3] = sdk.NewUintFromString("100000000000000000000020794415416798359")
	table_pow2[21][4] = sdk.NewUintFromString("100000000000000000000027725887222397812")
	table_pow2[21][5] = sdk.NewUintFromString("100000000000000000000034657359027997265")
	table_pow2[21][6] = sdk.NewUintFromString("100000000000000000000041588830833596719")
	table_pow2[21][7] = sdk.NewUintFromString("100000000000000000000048520302639196172")
	table_pow2[21][8] = sdk.NewUintFromString("100000000000000000000055451774444795625")
	table_pow2[21][9] = sdk.NewUintFromString("100000000000000000000062383246250395078")
	table_pow2[22][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[22][1] = sdk.NewUintFromString("100000000000000000000000693147180559945")
	table_pow2[22][2] = sdk.NewUintFromString("100000000000000000000001386294361119891")
	table_pow2[22][3] = sdk.NewUintFromString("100000000000000000000002079441541679836")
	table_pow2[22][4] = sdk.NewUintFromString("100000000000000000000002772588722239781")
	table_pow2[22][5] = sdk.NewUintFromString("100000000000000000000003465735902799727")
	table_pow2[22][6] = sdk.NewUintFromString("100000000000000000000004158883083359672")
	table_pow2[22][7] = sdk.NewUintFromString("100000000000000000000004852030263919617")
	table_pow2[22][8] = sdk.NewUintFromString("100000000000000000000005545177444479562")
	table_pow2[22][9] = sdk.NewUintFromString("100000000000000000000006238324625039508")
	table_pow2[23][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[23][1] = sdk.NewUintFromString("100000000000000000000000069314718055995")
	table_pow2[23][2] = sdk.NewUintFromString("100000000000000000000000138629436111989")
	table_pow2[23][3] = sdk.NewUintFromString("100000000000000000000000207944154167984")
	table_pow2[23][4] = sdk.NewUintFromString("100000000000000000000000277258872223978")
	table_pow2[23][5] = sdk.NewUintFromString("100000000000000000000000346573590279973")
	table_pow2[23][6] = sdk.NewUintFromString("100000000000000000000000415888308335967")
	table_pow2[23][7] = sdk.NewUintFromString("100000000000000000000000485203026391962")
	table_pow2[23][8] = sdk.NewUintFromString("100000000000000000000000554517744447956")
	table_pow2[23][9] = sdk.NewUintFromString("100000000000000000000000623832462503951")
	table_pow2[24][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[24][1] = sdk.NewUintFromString("100000000000000000000000006931471805599")
	table_pow2[24][2] = sdk.NewUintFromString("100000000000000000000000013862943611199")
	table_pow2[24][3] = sdk.NewUintFromString("100000000000000000000000020794415416798")
	table_pow2[24][4] = sdk.NewUintFromString("100000000000000000000000027725887222398")
	table_pow2[24][5] = sdk.NewUintFromString("100000000000000000000000034657359027997")
	table_pow2[24][6] = sdk.NewUintFromString("100000000000000000000000041588830833597")
	table_pow2[24][7] = sdk.NewUintFromString("100000000000000000000000048520302639196")
	table_pow2[24][8] = sdk.NewUintFromString("100000000000000000000000055451774444796")
	table_pow2[24][9] = sdk.NewUintFromString("100000000000000000000000062383246250395")
	table_pow2[25][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[25][1] = sdk.NewUintFromString("100000000000000000000000000693147180560")
	table_pow2[25][2] = sdk.NewUintFromString("100000000000000000000000001386294361120")
	table_pow2[25][3] = sdk.NewUintFromString("100000000000000000000000002079441541680")
	table_pow2[25][4] = sdk.NewUintFromString("100000000000000000000000002772588722240")
	table_pow2[25][5] = sdk.NewUintFromString("100000000000000000000000003465735902800")
	table_pow2[25][6] = sdk.NewUintFromString("100000000000000000000000004158883083360")
	table_pow2[25][7] = sdk.NewUintFromString("100000000000000000000000004852030263920")
	table_pow2[25][8] = sdk.NewUintFromString("100000000000000000000000005545177444480")
	table_pow2[25][9] = sdk.NewUintFromString("100000000000000000000000006238324625040")
	table_pow2[26][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[26][1] = sdk.NewUintFromString("100000000000000000000000000069314718056")
	table_pow2[26][2] = sdk.NewUintFromString("100000000000000000000000000138629436112")
	table_pow2[26][3] = sdk.NewUintFromString("100000000000000000000000000207944154168")
	table_pow2[26][4] = sdk.NewUintFromString("100000000000000000000000000277258872224")
	table_pow2[26][5] = sdk.NewUintFromString("100000000000000000000000000346573590280")
	table_pow2[26][6] = sdk.NewUintFromString("100000000000000000000000000415888308336")
	table_pow2[26][7] = sdk.NewUintFromString("100000000000000000000000000485203026392")
	table_pow2[26][8] = sdk.NewUintFromString("100000000000000000000000000554517744448")
	table_pow2[26][9] = sdk.NewUintFromString("100000000000000000000000000623832462504")
	table_pow2[27][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[27][1] = sdk.NewUintFromString("100000000000000000000000000006931471806")
	table_pow2[27][2] = sdk.NewUintFromString("100000000000000000000000000013862943611")
	table_pow2[27][3] = sdk.NewUintFromString("100000000000000000000000000020794415417")
	table_pow2[27][4] = sdk.NewUintFromString("100000000000000000000000000027725887222")
	table_pow2[27][5] = sdk.NewUintFromString("100000000000000000000000000034657359028")
	table_pow2[27][6] = sdk.NewUintFromString("100000000000000000000000000041588830834")
	table_pow2[27][7] = sdk.NewUintFromString("100000000000000000000000000048520302639")
	table_pow2[27][8] = sdk.NewUintFromString("100000000000000000000000000055451774445")
	table_pow2[27][9] = sdk.NewUintFromString("100000000000000000000000000062383246250")
	table_pow2[28][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[28][1] = sdk.NewUintFromString("100000000000000000000000000000693147181")
	table_pow2[28][2] = sdk.NewUintFromString("100000000000000000000000000001386294361")
	table_pow2[28][3] = sdk.NewUintFromString("100000000000000000000000000002079441542")
	table_pow2[28][4] = sdk.NewUintFromString("100000000000000000000000000002772588722")
	table_pow2[28][5] = sdk.NewUintFromString("100000000000000000000000000003465735903")
	table_pow2[28][6] = sdk.NewUintFromString("100000000000000000000000000004158883083")
	table_pow2[28][7] = sdk.NewUintFromString("100000000000000000000000000004852030264")
	table_pow2[28][8] = sdk.NewUintFromString("100000000000000000000000000005545177444")
	table_pow2[28][9] = sdk.NewUintFromString("100000000000000000000000000006238324625")
	table_pow2[29][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[29][1] = sdk.NewUintFromString("100000000000000000000000000000069314718")
	table_pow2[29][2] = sdk.NewUintFromString("100000000000000000000000000000138629436")
	table_pow2[29][3] = sdk.NewUintFromString("100000000000000000000000000000207944154")
	table_pow2[29][4] = sdk.NewUintFromString("100000000000000000000000000000277258872")
	table_pow2[29][5] = sdk.NewUintFromString("100000000000000000000000000000346573590")
	table_pow2[29][6] = sdk.NewUintFromString("100000000000000000000000000000415888308")
	table_pow2[29][7] = sdk.NewUintFromString("100000000000000000000000000000485203026")
	table_pow2[29][8] = sdk.NewUintFromString("100000000000000000000000000000554517744")
	table_pow2[29][9] = sdk.NewUintFromString("100000000000000000000000000000623832463")
	table_pow2[30][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[30][1] = sdk.NewUintFromString("100000000000000000000000000000006931472")
	table_pow2[30][2] = sdk.NewUintFromString("100000000000000000000000000000013862944")
	table_pow2[30][3] = sdk.NewUintFromString("100000000000000000000000000000020794415")
	table_pow2[30][4] = sdk.NewUintFromString("100000000000000000000000000000027725887")
	table_pow2[30][5] = sdk.NewUintFromString("100000000000000000000000000000034657359")
	table_pow2[30][6] = sdk.NewUintFromString("100000000000000000000000000000041588831")
	table_pow2[30][7] = sdk.NewUintFromString("100000000000000000000000000000048520303")
	table_pow2[30][8] = sdk.NewUintFromString("100000000000000000000000000000055451774")
	table_pow2[30][9] = sdk.NewUintFromString("100000000000000000000000000000062383246")

	LUT3_3_isSet = true
}

func setLUT3_4() {
	table_pow2[31][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[31][1] = sdk.NewUintFromString("100000000000000000000000000000000693147")
	table_pow2[31][2] = sdk.NewUintFromString("100000000000000000000000000000001386294")
	table_pow2[31][3] = sdk.NewUintFromString("100000000000000000000000000000002079442")
	table_pow2[31][4] = sdk.NewUintFromString("100000000000000000000000000000002772589")
	table_pow2[31][5] = sdk.NewUintFromString("100000000000000000000000000000003465736")
	table_pow2[31][6] = sdk.NewUintFromString("100000000000000000000000000000004158883")
	table_pow2[31][7] = sdk.NewUintFromString("100000000000000000000000000000004852030")
	table_pow2[31][8] = sdk.NewUintFromString("100000000000000000000000000000005545177")
	table_pow2[31][9] = sdk.NewUintFromString("100000000000000000000000000000006238325")
	table_pow2[32][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[32][1] = sdk.NewUintFromString("100000000000000000000000000000000069315")
	table_pow2[32][2] = sdk.NewUintFromString("100000000000000000000000000000000138629")
	table_pow2[32][3] = sdk.NewUintFromString("100000000000000000000000000000000207944")
	table_pow2[32][4] = sdk.NewUintFromString("100000000000000000000000000000000277259")
	table_pow2[32][5] = sdk.NewUintFromString("100000000000000000000000000000000346574")
	table_pow2[32][6] = sdk.NewUintFromString("100000000000000000000000000000000415888")
	table_pow2[32][7] = sdk.NewUintFromString("100000000000000000000000000000000485203")
	table_pow2[32][8] = sdk.NewUintFromString("100000000000000000000000000000000554518")
	table_pow2[32][9] = sdk.NewUintFromString("100000000000000000000000000000000623832")
	table_pow2[33][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[33][1] = sdk.NewUintFromString("100000000000000000000000000000000006931")
	table_pow2[33][2] = sdk.NewUintFromString("100000000000000000000000000000000013863")
	table_pow2[33][3] = sdk.NewUintFromString("100000000000000000000000000000000020794")
	table_pow2[33][4] = sdk.NewUintFromString("100000000000000000000000000000000027726")
	table_pow2[33][5] = sdk.NewUintFromString("100000000000000000000000000000000034657")
	table_pow2[33][6] = sdk.NewUintFromString("100000000000000000000000000000000041589")
	table_pow2[33][7] = sdk.NewUintFromString("100000000000000000000000000000000048520")
	table_pow2[33][8] = sdk.NewUintFromString("100000000000000000000000000000000055452")
	table_pow2[33][9] = sdk.NewUintFromString("100000000000000000000000000000000062383")
	table_pow2[34][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[34][1] = sdk.NewUintFromString("100000000000000000000000000000000000693")
	table_pow2[34][2] = sdk.NewUintFromString("100000000000000000000000000000000001386")
	table_pow2[34][3] = sdk.NewUintFromString("100000000000000000000000000000000002079")
	table_pow2[34][4] = sdk.NewUintFromString("100000000000000000000000000000000002773")
	table_pow2[34][5] = sdk.NewUintFromString("100000000000000000000000000000000003466")
	table_pow2[34][6] = sdk.NewUintFromString("100000000000000000000000000000000004159")
	table_pow2[34][7] = sdk.NewUintFromString("100000000000000000000000000000000004852")
	table_pow2[34][8] = sdk.NewUintFromString("100000000000000000000000000000000005545")
	table_pow2[34][9] = sdk.NewUintFromString("100000000000000000000000000000000006238")
	table_pow2[35][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[35][1] = sdk.NewUintFromString("100000000000000000000000000000000000069")
	table_pow2[35][2] = sdk.NewUintFromString("100000000000000000000000000000000000139")
	table_pow2[35][3] = sdk.NewUintFromString("100000000000000000000000000000000000208")
	table_pow2[35][4] = sdk.NewUintFromString("100000000000000000000000000000000000277")
	table_pow2[35][5] = sdk.NewUintFromString("100000000000000000000000000000000000347")
	table_pow2[35][6] = sdk.NewUintFromString("100000000000000000000000000000000000416")
	table_pow2[35][7] = sdk.NewUintFromString("100000000000000000000000000000000000485")
	table_pow2[35][8] = sdk.NewUintFromString("100000000000000000000000000000000000555")
	table_pow2[35][9] = sdk.NewUintFromString("100000000000000000000000000000000000624")
	table_pow2[36][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[36][1] = sdk.NewUintFromString("100000000000000000000000000000000000007")
	table_pow2[36][2] = sdk.NewUintFromString("100000000000000000000000000000000000014")
	table_pow2[36][3] = sdk.NewUintFromString("100000000000000000000000000000000000021")
	table_pow2[36][4] = sdk.NewUintFromString("100000000000000000000000000000000000028")
	table_pow2[36][5] = sdk.NewUintFromString("100000000000000000000000000000000000035")
	table_pow2[36][6] = sdk.NewUintFromString("100000000000000000000000000000000000042")
	table_pow2[36][7] = sdk.NewUintFromString("100000000000000000000000000000000000049")
	table_pow2[36][8] = sdk.NewUintFromString("100000000000000000000000000000000000055")
	table_pow2[36][9] = sdk.NewUintFromString("100000000000000000000000000000000000062")
	table_pow2[37][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[37][1] = sdk.NewUintFromString("100000000000000000000000000000000000001")
	table_pow2[37][2] = sdk.NewUintFromString("100000000000000000000000000000000000001")
	table_pow2[37][3] = sdk.NewUintFromString("100000000000000000000000000000000000002")
	table_pow2[37][4] = sdk.NewUintFromString("100000000000000000000000000000000000003")
	table_pow2[37][5] = sdk.NewUintFromString("100000000000000000000000000000000000003")
	table_pow2[37][6] = sdk.NewUintFromString("100000000000000000000000000000000000004")
	table_pow2[37][7] = sdk.NewUintFromString("100000000000000000000000000000000000005")
	table_pow2[37][8] = sdk.NewUintFromString("100000000000000000000000000000000000006")
	table_pow2[37][9] = sdk.NewUintFromString("100000000000000000000000000000000000006")
	table_pow2[38][0] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[38][1] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[38][2] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[38][3] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[38][4] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[38][5] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[38][6] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[38][7] = sdk.NewUintFromString("100000000000000000000000000000000000000")
	table_pow2[38][8] = sdk.NewUintFromString("100000000000000000000000000000000000001")
	table_pow2[38][9] = sdk.NewUintFromString("100000000000000000000000000000000000001")

	LUT3_4_isSet = true
}

/***** MODIFIERS *****/

func _onlyLUT1andLUT2AreSet() {
	requireThat(LUT1_isSet && LUT2_isSet,
		"Lookup tables 1 & 2 must first be set")
}

func _onlyLUT3isSet() {
	requireThat(LUT3_1_isSet && LUT3_2_isSet && LUT3_3_isSet && LUT3_4_isSet,
		"Lookup table 3 must first be set")
}

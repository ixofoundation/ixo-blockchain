package ixomath_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"

	"github.com/ixofoundation/ixo-blockchain/v6/ixomath"
)

type decimalTestSuite struct {
	suite.Suite
}

var (
	zeroAdditiveErrTolerance = ixomath.ErrTolerance{
		AdditiveTolerance: ixomath.ZeroDec(),
	}
)

func TestDecimalTestSuite(t *testing.T) {
	suite.Run(t, new(decimalTestSuite))
}

// assertMutResult given expected value after applying a math operation, a start value,
// mutative and non mutative results with start values, asserts that mutation are only applied
// to the mutative versions. Also, asserts that both results match the expected value.
func (s *decimalTestSuite) assertMutResult(expectedResult, startValue, mutativeResult, nonMutativeResult, mutativeStartValue, nonMutativeStartValue ixomath.BigDec) {
	// assert both results are as expected.
	s.Require().Equal(expectedResult, mutativeResult)
	s.Require().Equal(expectedResult, nonMutativeResult)

	// assert that mutative method mutated the receiver
	s.Require().Equal(mutativeStartValue, expectedResult)
	// assert that non-mutative method did not mutate the receiver
	s.Require().Equal(nonMutativeStartValue, startValue)
}

func (s *decimalTestSuite) TestAddMut() {
	toAdd := ixomath.MustNewBigDecFromStr("10")
	tests := map[string]struct {
		startValue        ixomath.BigDec
		expectedMutResult ixomath.BigDec
	}{
		"0":  {ixomath.NewBigDec(0), ixomath.NewBigDec(10)},
		"1":  {ixomath.NewBigDec(1), ixomath.NewBigDec(11)},
		"10": {ixomath.NewBigDec(10), ixomath.NewBigDec(20)},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			startMut := tc.startValue.Clone()
			startNonMut := tc.startValue.Clone()

			resultMut := startMut.AddMut(toAdd)
			resultNonMut := startNonMut.Add(toAdd)

			s.assertMutResult(tc.expectedMutResult, tc.startValue, resultMut, resultNonMut, startMut, startNonMut)
		})
	}
}

func (s *decimalTestSuite) TestQuoMut() {
	quoBy := ixomath.MustNewBigDecFromStr("2")
	tests := map[string]struct {
		startValue        ixomath.BigDec
		expectedMutResult ixomath.BigDec
	}{
		"0":  {ixomath.NewBigDec(0), ixomath.NewBigDec(0)},
		"1":  {ixomath.NewBigDec(1), ixomath.MustNewBigDecFromStr("0.5")},
		"10": {ixomath.NewBigDec(10), ixomath.NewBigDec(5)},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			startMut := tc.startValue.Clone()
			startNonMut := tc.startValue.Clone()

			resultMut := startMut.QuoMut(quoBy)
			resultNonMut := startNonMut.Quo(quoBy)

			s.assertMutResult(tc.expectedMutResult, tc.startValue, resultMut, resultNonMut, startMut, startNonMut)
		})
	}
}
func TestDecApproxEq(t *testing.T) {
	// d1 = 0.55, d2 = 0.6, tol = 0.1
	d1 := ixomath.NewBigDecWithPrec(55, 2)
	d2 := ixomath.NewBigDecWithPrec(6, 1)
	tol := ixomath.NewBigDecWithPrec(1, 1)

	require.True(ixomath.DecApproxEq(t, d1, d2, tol))

	// d1 = 0.55, d2 = 0.6, tol = 1E-5
	d1 = ixomath.NewBigDecWithPrec(55, 2)
	d2 = ixomath.NewBigDecWithPrec(6, 1)
	tol = ixomath.NewBigDecWithPrec(1, 5)

	require.False(ixomath.DecApproxEq(t, d1, d2, tol))

	// d1 = 0.6, d2 = 0.61, tol = 0.01
	d1 = ixomath.NewBigDecWithPrec(6, 1)
	d2 = ixomath.NewBigDecWithPrec(61, 2)
	tol = ixomath.NewBigDecWithPrec(1, 2)

	require.True(ixomath.DecApproxEq(t, d1, d2, tol))
}

// create a decimal from a decimal string (ex. "1234.5678")
func (s *decimalTestSuite) MustNewDecFromStr(str string) (d ixomath.BigDec) {
	d, err := ixomath.NewBigDecFromStr(str)
	s.Require().NoError(err)

	return d
}

func (s *decimalTestSuite) TestNewDecFromStr() {
	largeBigInt, success := new(big.Int).SetString("3144605511029693144278234343371835", 10)
	s.Require().True(success)

	tests := []struct {
		decimalStr string
		expErr     bool
		exp        ixomath.BigDec
	}{
		{"", true, ixomath.BigDec{}},
		{"0.-75", true, ixomath.BigDec{}},
		{"0", false, ixomath.NewBigDec(0)},
		{"1", false, ixomath.NewBigDec(1)},
		{"1.1", false, ixomath.NewBigDecWithPrec(11, 1)},
		{"0.75", false, ixomath.NewBigDecWithPrec(75, 2)},
		{"0.8", false, ixomath.NewBigDecWithPrec(8, 1)},
		{"0.11111", false, ixomath.NewBigDecWithPrec(11111, 5)},
		{"314460551102969.31442782343433718353144278234343371835", true, ixomath.NewBigDec(3141203149163817869)},
		{
			"314460551102969314427823434337.18357180924882313501835718092488231350",
			true, ixomath.NewBigDecFromBigIntWithPrec(largeBigInt, 4),
		},
		{
			"314460551102969314427823434337.1835",
			false, ixomath.NewBigDecFromBigIntWithPrec(largeBigInt, 4),
		},
		{".", true, ixomath.BigDec{}},
		{".0", true, ixomath.NewBigDec(0)},
		{"1.", true, ixomath.NewBigDec(1)},
		{"foobar", true, ixomath.BigDec{}},
		{"0.foobar", true, ixomath.BigDec{}},
		{"0.foobar.", true, ixomath.BigDec{}},
		{"179769313486231590772930519078902473361797697894230657273430081157732675805500963132708477322407536021120113879871393357658789768814416622492847430639474124377767893424865485276302219601246094119453082952085005768838150682342462881473913110540827237163350510684586298239947245938479716304835356329624224137216", true, ixomath.BigDec{}},
	}

	for tcIndex, tc := range tests {
		res, err := ixomath.NewBigDecFromStr(tc.decimalStr)
		if tc.expErr {
			s.Require().NotNil(err, "error expected, decimalStr %v, tc %v", tc.decimalStr, tcIndex)
		} else {
			s.Require().Nil(err, "unexpected error, decimalStr %v, tc %v", tc.decimalStr, tcIndex)
			s.Require().True(res.Equal(tc.exp), "equality was incorrect, res %v, exp %v, tc %v", res, tc.exp, tcIndex)
		}

		// negative tc
		res, err = ixomath.NewBigDecFromStr("-" + tc.decimalStr)
		if tc.expErr {
			s.Require().NotNil(err, "error expected, decimalStr %v, tc %v", tc.decimalStr, tcIndex)
		} else {
			s.Require().Nil(err, "unexpected error, decimalStr %v, tc %v", tc.decimalStr, tcIndex)
			exp := tc.exp.Mul(ixomath.NewBigDec(-1))
			s.Require().True(res.Equal(exp), "equality was incorrect, res %v, exp %v, tc %v", res, exp, tcIndex)
		}
	}
}

var interestingDecNumbers = []string{
	"123456789012345678901234567890123456789012345678901234567890.123456789012345678901234567890123456",
	"111111111111111111111111111111111111111111111111111111111111.111111111111111111111111111111111111",
	"999999999999999999999999999999999999999999999999999999999999.999999999999999999999999999999999999",
	"3141592653589793238462643383279502884197169399375105820974944.592307816406286208998628034825342117", // Approximation of Pi, extended
	"1618033988749894848204586834365638117720309179805762862135448.622705260462818902449707207204189391", // Approximation of Phi, extended
	"2718281828459045235360287471352662497757247093699959574966967.627724076630353547594571382178525166", // Approximation of e, extended
	"101010101010101010101010101010101010101010101010101010101010.101010101010101010101010101010101010",  // Binary pattern extended
	"1234567899876543210123456789987654321012345678998765432101234.567899876543210123456789987654321012", // Ascending and descending pattern extended
	"1123581321345589144233377610987159725844181676510946173113801.986211915342546982272763642843251547", // Inspired by Fibonacci sequence, creatively adjusted
	"1428571428571428571428571428571428571428571428571428571428571.428571428571428571428571428571428571", // Repeating decimal for 1/7 extended
}

var interestingDecNumbersBigDec = []ixomath.BigDec{}
var interestingDecNumbersDec = []ixomath.Dec{}

func init() {
	for _, str := range interestingDecNumbers {
		d, err := ixomath.NewBigDecFromStr(str)
		if err != nil {
			panic(fmt.Sprintf("error parsing decimal string %v: %v", str, err))
		}
		interestingDecNumbersBigDec = append(interestingDecNumbersBigDec, d)
		interestingDecNumbersDec = append(interestingDecNumbersDec, d.Dec())
	}
}

func (s *decimalTestSuite) TestNewBigDecFromDecMulDec() {
	type testcase struct {
		s1, s2 ixomath.Dec
	}
	tests := []testcase{}
	for _, d1 := range interestingDecNumbersDec {
		for _, d2 := range interestingDecNumbersDec {
			tests = append(tests, testcase{d1, d2})
		}
	}
	s.Require().True(len(tests) > 20, "no tests to run")
	for _, tc := range tests {
		s.Run(fmt.Sprintf("d1=%v, d2=%v", tc.s1, tc.s2), func() {
			s1D := ixomath.BigDecFromDec(tc.s1)
			s2D := ixomath.BigDecFromDec(tc.s2)
			expected := s1D.MulMut(s2D)
			actual := ixomath.NewBigDecFromDecMulDec(tc.s1, tc.s2)
			s.Require().True(expected.Equal(actual), "expected %v, got %v", expected, actual)
		})
	}
}

func (s *decimalTestSuite) TestQuoRoundUpNextIntMut() {
	type testcase struct {
		s1, s2 ixomath.BigDec
	}
	tests := []testcase{}
	for _, d1 := range interestingDecNumbersBigDec {
		for _, d2 := range interestingDecNumbersBigDec {
			tests = append(tests, testcase{d1, d2})
		}
	}
	s.Require().True(len(tests) > 20, "no tests to run")
	for _, tc := range tests {
		s.Run(fmt.Sprintf("d1=%v, d2=%v", tc.s1, tc.s2), func() {
			expected := tc.s1.QuoRoundUp(tc.s2).CeilMut()
			actual := tc.s1.QuoRoundUpNextIntMut(tc.s2)
			s.Require().True(expected.Equal(actual), "expected %v, got %v", expected, actual)
		})
	}
}

func (s *decimalTestSuite) TestDecString() {
	tests := []struct {
		d    ixomath.BigDec
		want string
	}{
		{ixomath.NewBigDec(0), "0.000000000000000000000000000000000000"},
		{ixomath.NewBigDec(1), "1.000000000000000000000000000000000000"},
		{ixomath.NewBigDec(10), "10.000000000000000000000000000000000000"},
		{ixomath.NewBigDec(12340), "12340.000000000000000000000000000000000000"},
		{ixomath.NewBigDecWithPrec(12340, 4), "1.234000000000000000000000000000000000"},
		{ixomath.NewBigDecWithPrec(12340, 5), "0.123400000000000000000000000000000000"},
		{ixomath.NewBigDecWithPrec(12340, 8), "0.000123400000000000000000000000000000"},
		{ixomath.NewBigDecWithPrec(1009009009009009009, 17), "10.090090090090090090000000000000000000"},
		{ixomath.MustNewBigDecFromStr("10.090090090090090090090090090090090090"), "10.090090090090090090090090090090090090"},
	}
	for tcIndex, tc := range tests {
		s.Require().Equal(tc.want, tc.d.String(), "bad String(), index: %v", tcIndex)
	}
}

func (s *decimalTestSuite) TestDecFloat64() {
	tests := []struct {
		d    ixomath.BigDec
		want float64
	}{
		{ixomath.NewBigDec(0), 0.000000000000000000},
		{ixomath.NewBigDec(1), 1.000000000000000000},
		{ixomath.NewBigDec(10), 10.000000000000000000},
		{ixomath.NewBigDec(12340), 12340.000000000000000000},
		{ixomath.NewBigDecWithPrec(12340, 4), 1.234000000000000000},
		{ixomath.NewBigDecWithPrec(12340, 5), 0.123400000000000000},
		{ixomath.NewBigDecWithPrec(12340, 8), 0.000123400000000000},
		{ixomath.NewBigDecWithPrec(1009009009009009009, 17), 10.090090090090090090},
	}
	for tcIndex, tc := range tests {
		value, err := tc.d.Float64()
		s.Require().Nil(err, "error getting Float64(), index: %v", tcIndex)
		s.Require().Equal(tc.want, value, "bad Float64(), index: %v", tcIndex)
		s.Require().Equal(tc.want, tc.d.MustFloat64(), "bad MustFloat64(), index: %v", tcIndex)
	}
}

func (s *decimalTestSuite) TestSdkDec() {
	tests := []struct {
		d        ixomath.BigDec
		want     ixomath.Dec
		expPanic bool
	}{
		{ixomath.NewBigDec(0), ixomath.MustNewDecFromStr("0.000000000000000000"), false},
		{ixomath.NewBigDec(1), ixomath.MustNewDecFromStr("1.000000000000000000"), false},
		{ixomath.NewBigDec(10), ixomath.MustNewDecFromStr("10.000000000000000000"), false},
		{ixomath.NewBigDec(12340), ixomath.MustNewDecFromStr("12340.000000000000000000"), false},
		{ixomath.NewBigDecWithPrec(12340, 4), ixomath.MustNewDecFromStr("1.234000000000000000"), false},
		{ixomath.NewBigDecWithPrec(12340, 5), ixomath.MustNewDecFromStr("0.123400000000000000"), false},
		{ixomath.NewBigDecWithPrec(12340, 8), ixomath.MustNewDecFromStr("0.000123400000000000"), false},
		{ixomath.NewBigDecWithPrec(1009009009009009009, 17), ixomath.MustNewDecFromStr("10.090090090090090090"), false},
	}
	for tcIndex, tc := range tests {
		if tc.expPanic {
			s.Require().Panics(func() { tc.d.Dec() })
		} else {
			value := tc.d.Dec()
			s.Require().Equal(tc.want, value, "bad SdkDec(), index: %v", tcIndex)
		}
	}
}

func (s *decimalTestSuite) TestSdkDecRoundUp() {
	tests := []struct {
		d        ixomath.BigDec
		want     ixomath.Dec
		expPanic bool
	}{
		{ixomath.NewBigDec(0), ixomath.MustNewDecFromStr("0.000000000000000000"), false},
		{ixomath.NewBigDec(1), ixomath.MustNewDecFromStr("1.000000000000000000"), false},
		{ixomath.NewBigDec(10), ixomath.MustNewDecFromStr("10.000000000000000000"), false},
		{ixomath.NewBigDec(12340), ixomath.MustNewDecFromStr("12340.000000000000000000"), false},
		{ixomath.NewBigDecWithPrec(12340, 4), ixomath.MustNewDecFromStr("1.234000000000000000"), false},
		{ixomath.NewBigDecWithPrec(12340, 5), ixomath.MustNewDecFromStr("0.123400000000000000"), false},
		{ixomath.NewBigDecWithPrec(12340, 8), ixomath.MustNewDecFromStr("0.000123400000000000"), false},
		{ixomath.NewBigDecWithPrec(1009009009009009009, 17), ixomath.MustNewDecFromStr("10.090090090090090090"), false},
		{ixomath.NewBigDecWithPrec(1009009009009009009, 19), ixomath.MustNewDecFromStr("0.100900900900900901"), false},
	}
	for tcIndex, tc := range tests {
		if tc.expPanic {
			s.Require().Panics(func() { tc.d.DecRoundUp() })
		} else {
			value := tc.d.DecRoundUp()
			s.Require().Equal(tc.want, value, "bad SdkDec(), index: %v", tcIndex)
		}
	}
}

func (s *decimalTestSuite) TestBigDecFromSdkDec() {
	tests := []struct {
		d        ixomath.Dec
		want     ixomath.BigDec
		expPanic bool
	}{
		{ixomath.MustNewDecFromStr("0.000000000000000000"), ixomath.NewBigDec(0), false},
		{ixomath.MustNewDecFromStr("1.000000000000000000"), ixomath.NewBigDec(1), false},
		{ixomath.MustNewDecFromStr("10.000000000000000000"), ixomath.NewBigDec(10), false},
		{ixomath.MustNewDecFromStr("12340.000000000000000000"), ixomath.NewBigDec(12340), false},
		{ixomath.MustNewDecFromStr("1.234000000000000000"), ixomath.NewBigDecWithPrec(12340, 4), false},
		{ixomath.MustNewDecFromStr("0.123400000000000000"), ixomath.NewBigDecWithPrec(12340, 5), false},
		{ixomath.MustNewDecFromStr("0.000123400000000000"), ixomath.NewBigDecWithPrec(12340, 8), false},
		{ixomath.MustNewDecFromStr("10.090090090090090090"), ixomath.NewBigDecWithPrec(1009009009009009009, 17), false},
	}
	for tcIndex, tc := range tests {
		if tc.expPanic {
			s.Require().Panics(func() { ixomath.BigDecFromDec(tc.d) })
		} else {
			value := ixomath.BigDecFromDec(tc.d)
			s.Require().Equal(tc.want, value, "bad ixomath.BigDecFromDec(), index: %v", tcIndex)
		}
	}
}

func (s *decimalTestSuite) TestBigDecFromSdkInt() {
	tests := []struct {
		i        ixomath.Int
		want     ixomath.BigDec
		expPanic bool
	}{
		{ixomath.ZeroInt(), ixomath.NewBigDec(0), false},
		{ixomath.OneInt(), ixomath.NewBigDec(1), false},
		{ixomath.NewInt(10), ixomath.NewBigDec(10), false},
		{ixomath.NewInt(10090090090090090), ixomath.NewBigDecWithPrec(10090090090090090, 0), false},
	}
	for tcIndex, tc := range tests {
		if tc.expPanic {
			s.Require().Panics(func() { ixomath.BigDecFromSDKInt(tc.i) })
		} else {
			value := ixomath.BigDecFromSDKInt(tc.i)
			s.Require().Equal(tc.want, value, "bad ixomath.BigDecFromDec(), index: %v", tcIndex)
		}
	}
}

func (s *decimalTestSuite) TestBigDecFromSdkDecSlice() {
	tests := []struct {
		d        []ixomath.Dec
		want     []ixomath.BigDec
		expPanic bool
	}{
		{[]ixomath.Dec{ixomath.MustNewDecFromStr("0.000000000000000000")}, []ixomath.BigDec{ixomath.NewBigDec(0)}, false},
		{[]ixomath.Dec{ixomath.MustNewDecFromStr("0.000000000000000000"), ixomath.MustNewDecFromStr("1.000000000000000000")}, []ixomath.BigDec{ixomath.NewBigDec(0), ixomath.NewBigDec(1)}, false},
		{[]ixomath.Dec{ixomath.MustNewDecFromStr("1.000000000000000000"), ixomath.MustNewDecFromStr("0.000000000000000000"), ixomath.MustNewDecFromStr("0.000123400000000000")}, []ixomath.BigDec{ixomath.NewBigDec(1), ixomath.NewBigDec(0), ixomath.NewBigDecWithPrec(12340, 8)}, false},
		{[]ixomath.Dec{ixomath.MustNewDecFromStr("10.000000000000000000")}, []ixomath.BigDec{ixomath.NewBigDec(10)}, false},
		{[]ixomath.Dec{ixomath.MustNewDecFromStr("12340.000000000000000000")}, []ixomath.BigDec{ixomath.NewBigDec(12340)}, false},
		{[]ixomath.Dec{ixomath.MustNewDecFromStr("1.234000000000000000"), ixomath.MustNewDecFromStr("12340.000000000000000000")}, []ixomath.BigDec{ixomath.NewBigDecWithPrec(12340, 4), ixomath.NewBigDec(12340)}, false},
		{[]ixomath.Dec{ixomath.MustNewDecFromStr("0.123400000000000000"), ixomath.MustNewDecFromStr("12340.000000000000000000")}, []ixomath.BigDec{ixomath.NewBigDecWithPrec(12340, 5), ixomath.NewBigDec(12340)}, false},
		{[]ixomath.Dec{ixomath.MustNewDecFromStr("0.000123400000000000"), ixomath.MustNewDecFromStr("10.090090090090090090")}, []ixomath.BigDec{ixomath.NewBigDecWithPrec(12340, 8), ixomath.NewBigDecWithPrec(1009009009009009009, 17)}, false},
		{[]ixomath.Dec{ixomath.MustNewDecFromStr("10.090090090090090090"), ixomath.MustNewDecFromStr("10.090090090090090090")}, []ixomath.BigDec{ixomath.NewBigDecWithPrec(1009009009009009009, 17), ixomath.NewBigDecWithPrec(1009009009009009009, 17)}, false},
	}
	for tcIndex, tc := range tests {
		if tc.expPanic {
			s.Require().Panics(func() { ixomath.BigDecFromDecSlice(tc.d) })
		} else {
			value := ixomath.BigDecFromDecSlice(tc.d)
			s.Require().Equal(tc.want, value, "bad ixomath.BigDecFromDec(), index: %v", tcIndex)
		}
	}
}

func (s *decimalTestSuite) TestEqualities() {
	tests := []struct {
		d1, d2     ixomath.BigDec
		gt, lt, eq bool
	}{
		{ixomath.NewBigDec(0), ixomath.NewBigDec(0), false, false, true},
		{ixomath.NewBigDecWithPrec(0, 2), ixomath.NewBigDecWithPrec(0, 4), false, false, true},
		{ixomath.NewBigDecWithPrec(100, 0), ixomath.NewBigDecWithPrec(100, 0), false, false, true},
		{ixomath.NewBigDecWithPrec(-100, 0), ixomath.NewBigDecWithPrec(-100, 0), false, false, true},
		{ixomath.NewBigDecWithPrec(-1, 1), ixomath.NewBigDecWithPrec(-1, 1), false, false, true},
		{ixomath.NewBigDecWithPrec(3333, 3), ixomath.NewBigDecWithPrec(3333, 3), false, false, true},

		{ixomath.NewBigDecWithPrec(0, 0), ixomath.NewBigDecWithPrec(3333, 3), false, true, false},
		{ixomath.NewBigDecWithPrec(0, 0), ixomath.NewBigDecWithPrec(100, 0), false, true, false},
		{ixomath.NewBigDecWithPrec(-1, 0), ixomath.NewBigDecWithPrec(3333, 3), false, true, false},
		{ixomath.NewBigDecWithPrec(-1, 0), ixomath.NewBigDecWithPrec(100, 0), false, true, false},
		{ixomath.NewBigDecWithPrec(1111, 3), ixomath.NewBigDecWithPrec(100, 0), false, true, false},
		{ixomath.NewBigDecWithPrec(1111, 3), ixomath.NewBigDecWithPrec(3333, 3), false, true, false},
		{ixomath.NewBigDecWithPrec(-3333, 3), ixomath.NewBigDecWithPrec(-1111, 3), false, true, false},

		{ixomath.NewBigDecWithPrec(3333, 3), ixomath.NewBigDecWithPrec(0, 0), true, false, false},
		{ixomath.NewBigDecWithPrec(100, 0), ixomath.NewBigDecWithPrec(0, 0), true, false, false},
		{ixomath.NewBigDecWithPrec(3333, 3), ixomath.NewBigDecWithPrec(-1, 0), true, false, false},
		{ixomath.NewBigDecWithPrec(100, 0), ixomath.NewBigDecWithPrec(-1, 0), true, false, false},
		{ixomath.NewBigDecWithPrec(100, 0), ixomath.NewBigDecWithPrec(1111, 3), true, false, false},
		{ixomath.NewBigDecWithPrec(3333, 3), ixomath.NewBigDecWithPrec(1111, 3), true, false, false},
		{ixomath.NewBigDecWithPrec(-1111, 3), ixomath.NewBigDecWithPrec(-3333, 3), true, false, false},
	}

	for tcIndex, tc := range tests {
		s.Require().Equal(tc.gt, tc.d1.GT(tc.d2), "GT result is incorrect, tc %d", tcIndex)
		s.Require().Equal(tc.lt, tc.d1.LT(tc.d2), "LT result is incorrect, tc %d", tcIndex)
		s.Require().Equal(tc.eq, tc.d1.Equal(tc.d2), "equality result is incorrect, tc %d", tcIndex)
	}
}

func (s *decimalTestSuite) TestDecsEqual() {
	tests := []struct {
		d1s, d2s []ixomath.BigDec
		eq       bool
	}{
		{[]ixomath.BigDec{ixomath.NewBigDec(0)}, []ixomath.BigDec{ixomath.NewBigDec(0)}, true},
		{[]ixomath.BigDec{ixomath.NewBigDec(0)}, []ixomath.BigDec{ixomath.NewBigDec(1)}, false},
		{[]ixomath.BigDec{ixomath.NewBigDec(0)}, []ixomath.BigDec{}, false},
		{[]ixomath.BigDec{ixomath.NewBigDec(0), ixomath.NewBigDec(1)}, []ixomath.BigDec{ixomath.NewBigDec(0), ixomath.NewBigDec(1)}, true},
		{[]ixomath.BigDec{ixomath.NewBigDec(1), ixomath.NewBigDec(0)}, []ixomath.BigDec{ixomath.NewBigDec(1), ixomath.NewBigDec(0)}, true},
		{[]ixomath.BigDec{ixomath.NewBigDec(1), ixomath.NewBigDec(0)}, []ixomath.BigDec{ixomath.NewBigDec(0), ixomath.NewBigDec(1)}, false},
		{[]ixomath.BigDec{ixomath.NewBigDec(1), ixomath.NewBigDec(0)}, []ixomath.BigDec{ixomath.NewBigDec(1)}, false},
		{[]ixomath.BigDec{ixomath.NewBigDec(1), ixomath.NewBigDec(2)}, []ixomath.BigDec{ixomath.NewBigDec(2), ixomath.NewBigDec(4)}, false},
		{[]ixomath.BigDec{ixomath.NewBigDec(3), ixomath.NewBigDec(18)}, []ixomath.BigDec{ixomath.NewBigDec(1), ixomath.NewBigDec(6)}, false},
	}

	for tcIndex, tc := range tests {
		s.Require().Equal(tc.eq, ixomath.DecsEqual(tc.d1s, tc.d2s), "equality of decional arrays is incorrect, tc %d", tcIndex)
		s.Require().Equal(tc.eq, ixomath.DecsEqual(tc.d2s, tc.d1s), "equality of decional arrays is incorrect (converse), tc %d", tcIndex)
	}
}

func (s *decimalTestSuite) TestArithmetic() {
	tests := []struct {
		d1, d2                                ixomath.BigDec
		expMul, expMulTruncate, expMulRoundUp ixomath.BigDec
		expQuo, expQuoRoundUp, expQuoTruncate ixomath.BigDec
		expAdd, expSub                        ixomath.BigDec
	}{
		//  d1         d2         MUL    MulTruncate   MulRoundUp    QUO    QUORoundUp QUOTrunctate  ADD         SUB
		{ixomath.NewBigDec(0), ixomath.NewBigDec(0), ixomath.NewBigDec(0), ixomath.NewBigDec(0), ixomath.NewBigDec(0), ixomath.NewBigDec(0), ixomath.NewBigDec(0), ixomath.NewBigDec(0), ixomath.NewBigDec(0), ixomath.NewBigDec(0)},
		{ixomath.NewBigDec(1), ixomath.NewBigDec(0), ixomath.NewBigDec(0), ixomath.NewBigDec(0), ixomath.NewBigDec(0), ixomath.NewBigDec(0), ixomath.NewBigDec(0), ixomath.NewBigDec(0), ixomath.NewBigDec(1), ixomath.NewBigDec(1)},
		{ixomath.NewBigDec(0), ixomath.NewBigDec(1), ixomath.NewBigDec(0), ixomath.NewBigDec(0), ixomath.NewBigDec(0), ixomath.NewBigDec(0), ixomath.NewBigDec(0), ixomath.NewBigDec(0), ixomath.NewBigDec(1), ixomath.NewBigDec(-1)},
		{ixomath.NewBigDec(0), ixomath.NewBigDec(-1), ixomath.NewBigDec(0), ixomath.NewBigDec(0), ixomath.NewBigDec(0), ixomath.NewBigDec(0), ixomath.NewBigDec(0), ixomath.NewBigDec(0), ixomath.NewBigDec(-1), ixomath.NewBigDec(1)},
		{ixomath.NewBigDec(-1), ixomath.NewBigDec(0), ixomath.NewBigDec(0), ixomath.NewBigDec(0), ixomath.NewBigDec(0), ixomath.NewBigDec(0), ixomath.NewBigDec(0), ixomath.NewBigDec(0), ixomath.NewBigDec(-1), ixomath.NewBigDec(-1)},

		{ixomath.NewBigDec(1), ixomath.NewBigDec(1), ixomath.NewBigDec(1), ixomath.NewBigDec(1), ixomath.NewBigDec(1), ixomath.NewBigDec(1), ixomath.NewBigDec(1), ixomath.NewBigDec(1), ixomath.NewBigDec(2), ixomath.NewBigDec(0)},
		{ixomath.NewBigDec(-1), ixomath.NewBigDec(-1), ixomath.NewBigDec(1), ixomath.NewBigDec(1), ixomath.NewBigDec(1), ixomath.NewBigDec(1), ixomath.NewBigDec(1), ixomath.NewBigDec(1), ixomath.NewBigDec(-2), ixomath.NewBigDec(0)},
		{ixomath.NewBigDec(1), ixomath.NewBigDec(-1), ixomath.NewBigDec(-1), ixomath.NewBigDec(-1), ixomath.NewBigDec(-1), ixomath.NewBigDec(-1), ixomath.NewBigDec(-1), ixomath.NewBigDec(-1), ixomath.NewBigDec(0), ixomath.NewBigDec(2)},
		{ixomath.NewBigDec(-1), ixomath.NewBigDec(1), ixomath.NewBigDec(-1), ixomath.NewBigDec(-1), ixomath.NewBigDec(-1), ixomath.NewBigDec(-1), ixomath.NewBigDec(-1), ixomath.NewBigDec(-1), ixomath.NewBigDec(0), ixomath.NewBigDec(-2)},

		{
			ixomath.NewBigDec(3), ixomath.NewBigDec(7), ixomath.NewBigDec(21), ixomath.NewBigDec(21), ixomath.NewBigDec(21),
			ixomath.MustNewBigDecFromStr("0.428571428571428571428571428571428571"), ixomath.MustNewBigDecFromStr("0.428571428571428571428571428571428572"), ixomath.MustNewBigDecFromStr("0.428571428571428571428571428571428571"),
			ixomath.NewBigDec(10), ixomath.NewBigDec(-4),
		},
		{
			ixomath.NewBigDec(2), ixomath.NewBigDec(4), ixomath.NewBigDec(8), ixomath.NewBigDec(8), ixomath.NewBigDec(8), ixomath.NewBigDecWithPrec(5, 1), ixomath.NewBigDecWithPrec(5, 1), ixomath.NewBigDecWithPrec(5, 1),
			ixomath.NewBigDec(6), ixomath.NewBigDec(-2),
		},

		{ixomath.NewBigDec(100), ixomath.NewBigDec(100), ixomath.NewBigDec(10000), ixomath.NewBigDec(10000), ixomath.NewBigDec(10000), ixomath.NewBigDec(1), ixomath.NewBigDec(1), ixomath.NewBigDec(1), ixomath.NewBigDec(200), ixomath.NewBigDec(0)},

		{
			ixomath.NewBigDecWithPrec(15, 1), ixomath.NewBigDecWithPrec(15, 1), ixomath.NewBigDecWithPrec(225, 2), ixomath.NewBigDecWithPrec(225, 2), ixomath.NewBigDecWithPrec(225, 2),
			ixomath.NewBigDec(1), ixomath.NewBigDec(1), ixomath.NewBigDec(1), ixomath.NewBigDec(3), ixomath.NewBigDec(0),
		},
		{
			ixomath.NewBigDecWithPrec(3333, 4), ixomath.NewBigDecWithPrec(333, 4), ixomath.NewBigDecWithPrec(1109889, 8), ixomath.NewBigDecWithPrec(1109889, 8), ixomath.NewBigDecWithPrec(1109889, 8),
			ixomath.MustNewBigDecFromStr("10.009009009009009009009009009009009009"), ixomath.MustNewBigDecFromStr("10.009009009009009009009009009009009010"), ixomath.MustNewBigDecFromStr("10.009009009009009009009009009009009009"),
			ixomath.NewBigDecWithPrec(3666, 4), ixomath.NewBigDecWithPrec(3, 1),
		},
	}

	for tcIndex, tc := range tests {
		tc := tc
		resAdd := tc.d1.Add(tc.d2)
		resSub := tc.d1.Sub(tc.d2)
		resMul := tc.d1.Mul(tc.d2)
		resMulTruncate := tc.d1.MulTruncate(tc.d2)
		resMulRoundUp := tc.d1.MulRoundUp(tc.d2)
		s.Require().True(tc.expAdd.Equal(resAdd), "exp %v, res %v, tc %d", tc.expAdd, resAdd, tcIndex)
		s.Require().True(tc.expSub.Equal(resSub), "exp %v, res %v, tc %d", tc.expSub, resSub, tcIndex)
		s.Require().True(tc.expMul.Equal(resMul), "exp %v, res %v, tc %d", tc.expMul, resMul, tcIndex)
		s.Require().True(tc.expMulTruncate.Equal(resMulTruncate), "exp %v, res %v, tc %d", tc.expMulTruncate, resMulTruncate, tcIndex)
		s.Require().True(tc.expMulRoundUp.Equal(resMulRoundUp), "exp %v, res %v, tc %d", tc.expMulRoundUp, resMulRoundUp, tcIndex)

		if tc.d2.IsZero() { // panic for divide by zero
			s.Require().Panics(func() { tc.d1.Quo(tc.d2) })
		} else {
			resQuo := tc.d1.Quo(tc.d2)
			s.Require().True(tc.expQuo.Equal(resQuo), "exp %v, res %v, tc %d", tc.expQuo.String(), resQuo.String(), tcIndex)

			resQuoRoundUp := tc.d1.QuoRoundUp(tc.d2)
			s.Require().True(tc.expQuoRoundUp.Equal(resQuoRoundUp), "exp %v, res %v, tc %d",
				tc.expQuoRoundUp.String(), resQuoRoundUp.String(), tcIndex)

			resQuoRoundUpDec := tc.d1.QuoByDecRoundUp(tc.d2.Dec())
			expResQuoRoundUpDec := tc.d1.QuoRoundUp(ixomath.BigDecFromDec(tc.d2.Dec()))
			s.Require().True(expResQuoRoundUpDec.Equal(resQuoRoundUpDec), "exp %v, res %v, tc %d",
				expResQuoRoundUpDec.String(), resQuoRoundUpDec.String(), tcIndex)

			resQuoTruncate := tc.d1.QuoTruncate(tc.d2)
			s.Require().True(tc.expQuoTruncate.Equal(resQuoTruncate), "exp %v, res %v, tc %d",
				tc.expQuoTruncate.String(), resQuoTruncate.String(), tcIndex)
		}
	}
}

func (s *decimalTestSuite) TestMulRoundUp_RoundingAtPrecisionEnd() {
	var (
		a                = ixomath.MustNewBigDecFromStr("0.000000000000000000000000000000000009")
		b                = ixomath.MustNewBigDecFromStr("0.000000000000000000000000000000000009")
		expectedRoundUp  = ixomath.MustNewBigDecFromStr("0.000000000000000000000000000000000001")
		expectedTruncate = ixomath.MustNewBigDecFromStr("0.000000000000000000000000000000000000")
	)

	actualRoundUp := a.MulRoundUp(b)
	s.Require().Equal(expectedRoundUp.String(), actualRoundUp.String(), "exp %v, res %v", expectedRoundUp, actualRoundUp)

	actualTruncate := a.MulTruncate(b)
	s.Require().Equal(expectedTruncate.String(), actualTruncate.String(), "exp %v, res %v", expectedTruncate, actualTruncate)
}

func (s *decimalTestSuite) TestBankerRoundChop() {
	tests := []struct {
		d1  ixomath.BigDec
		exp int64
	}{
		{s.MustNewDecFromStr("0.25"), 0},
		{s.MustNewDecFromStr("0"), 0},
		{s.MustNewDecFromStr("1"), 1},
		{s.MustNewDecFromStr("0.75"), 1},
		{s.MustNewDecFromStr("0.5"), 0},
		{s.MustNewDecFromStr("7.5"), 8},
		{s.MustNewDecFromStr("1.5"), 2},
		{s.MustNewDecFromStr("2.5"), 2},
		{s.MustNewDecFromStr("0.545"), 1}, // 0.545-> 1 even though 5 is first decimal and 1 not even
		{s.MustNewDecFromStr("1.545"), 2},
	}

	for tcIndex, tc := range tests {
		resNeg := tc.d1.Neg().RoundInt64()
		s.Require().Equal(-1*tc.exp, resNeg, "negative tc %d", tcIndex)

		resPos := tc.d1.RoundInt64()
		s.Require().Equal(tc.exp, resPos, "positive tc %d", tcIndex)
	}
}

func (s *decimalTestSuite) TestTruncate() {
	tests := []struct {
		d1  ixomath.BigDec
		exp int64
	}{
		{s.MustNewDecFromStr("0"), 0},
		{s.MustNewDecFromStr("0.25"), 0},
		{s.MustNewDecFromStr("0.75"), 0},
		{s.MustNewDecFromStr("1"), 1},
		{s.MustNewDecFromStr("1.5"), 1},
		{s.MustNewDecFromStr("7.5"), 7},
		{s.MustNewDecFromStr("7.6"), 7},
		{s.MustNewDecFromStr("7.4"), 7},
		{s.MustNewDecFromStr("100.1"), 100},
		{s.MustNewDecFromStr("1000.1"), 1000},
	}

	for tcIndex, tc := range tests {
		resNeg := tc.d1.Neg().TruncateInt64()
		s.Require().Equal(-1*tc.exp, resNeg, "negative tc %d", tcIndex)

		resPos := tc.d1.TruncateInt64()
		s.Require().Equal(tc.exp, resPos, "positive tc %d", tcIndex)
	}
}

func (s *decimalTestSuite) TestStringOverflow() {
	// two random 64 bit primes
	dec1, err := ixomath.NewBigDecFromStr("51643150036226787134389711697696177267")
	s.Require().NoError(err)
	dec2, err := ixomath.NewBigDecFromStr("-31798496660535729618459429845579852627")
	s.Require().NoError(err)
	dec3 := dec1.Add(dec2)
	s.Require().Equal(
		"19844653375691057515930281852116324640.000000000000000000000000000000000000",
		dec3.String(),
	)
}

func (s *decimalTestSuite) TestDecMulInt() {
	tests := []struct {
		sdkDec ixomath.BigDec
		sdkInt ixomath.BigInt
		want   ixomath.BigDec
	}{
		{ixomath.NewBigDec(10), ixomath.NewBigInt(2), ixomath.NewBigDec(20)},
		{ixomath.NewBigDec(1000000), ixomath.NewBigInt(100), ixomath.NewBigDec(100000000)},
		{ixomath.NewBigDecWithPrec(1, 1), ixomath.NewBigInt(10), ixomath.NewBigDec(1)},
		{ixomath.NewBigDecWithPrec(1, 5), ixomath.NewBigInt(20), ixomath.NewBigDecWithPrec(2, 4)},
	}
	for i, tc := range tests {
		got := tc.sdkDec.MulInt(tc.sdkInt)
		s.Require().Equal(tc.want, got, "Incorrect result on test case %d", i)
	}
}

func (s *decimalTestSuite) TestDecCeil() {
	testCases := []struct {
		input    ixomath.BigDec
		expected ixomath.BigDec
	}{
		{ixomath.MustNewBigDecFromStr("0.001"), ixomath.NewBigDec(1)},   // 0.001 => 1.0
		{ixomath.MustNewBigDecFromStr("-0.001"), ixomath.ZeroBigDec()},  // -0.001 => 0.0
		{ixomath.ZeroBigDec(), ixomath.ZeroBigDec()},                    // 0.0 => 0.0
		{ixomath.MustNewBigDecFromStr("0.9"), ixomath.NewBigDec(1)},     // 0.9 => 1.0
		{ixomath.MustNewBigDecFromStr("4.001"), ixomath.NewBigDec(5)},   // 4.001 => 5.0
		{ixomath.MustNewBigDecFromStr("-4.001"), ixomath.NewBigDec(-4)}, // -4.001 => -4.0
		{ixomath.MustNewBigDecFromStr("4.7"), ixomath.NewBigDec(5)},     // 4.7 => 5.0
		{ixomath.MustNewBigDecFromStr("-4.7"), ixomath.NewBigDec(-4)},   // -4.7 => -4.0
	}

	for i, tc := range testCases {
		tc_input_copy := tc.input.Clone()
		res := tc.input.Ceil()
		s.Require().Equal(tc.input, tc_input_copy, "unexpected mutation of input in test case %d, input: %v", i, tc.input)
		s.Require().True(tc.expected.Equal(res), "unexpected result for test case %d, input: %v, got %v, expected %v:", i, tc.input, res, tc.expected)
	}
}

func (s *decimalTestSuite) TestApproxRoot() {
	testCases := []struct {
		input    ixomath.BigDec
		root     uint64
		expected ixomath.BigDec
	}{
		{ixomath.OneBigDec(), 10, ixomath.OneBigDec()},                                                                              // 1.0 ^ (0.1) => 1.0
		{ixomath.NewBigDecWithPrec(25, 2), 2, ixomath.NewBigDecWithPrec(5, 1)},                                                      // 0.25 ^ (0.5) => 0.5
		{ixomath.NewBigDecWithPrec(4, 2), 2, ixomath.NewBigDecWithPrec(2, 1)},                                                       // 0.04 ^ (0.5) => 0.2
		{ixomath.NewBigDecFromInt(ixomath.NewBigInt(27)), 3, ixomath.NewBigDecFromInt(ixomath.NewBigInt(3))},                        // 27 ^ (1/3) => 3
		{ixomath.NewBigDecFromInt(ixomath.NewBigInt(-81)), 4, ixomath.NewBigDecFromInt(ixomath.NewBigInt(-3))},                      // -81 ^ (0.25) => -3
		{ixomath.NewBigDecFromInt(ixomath.NewBigInt(2)), 2, ixomath.MustNewBigDecFromStr("1.414213562373095048801688724209698079")}, // 2 ^ (0.5) => 1.414213562373095048801688724209698079
		{ixomath.NewBigDecWithPrec(1005, 3), 31536000, ixomath.MustNewBigDecFromStr("1.000000000158153903837946258002096839")},      // 1.005 ^ (1/31536000) ≈ 1.000000000158153903837946258002096839
		{ixomath.SmallestBigDec(), 2, ixomath.NewBigDecWithPrec(1, 18)},                                                             // 1e-36 ^ (0.5) => 1e-18
		{ixomath.SmallestBigDec(), 3, ixomath.MustNewBigDecFromStr("0.000000000001000000000000000002431786")},                       // 1e-36 ^ (1/3) => 1e-12
		{ixomath.NewBigDecWithPrec(1, 8), 3, ixomath.MustNewBigDecFromStr("0.002154434690031883721759293566519280")},                // 1e-8 ^ (1/3) ≈ 0.002154434690031883721759293566519
	}

	// In the case of 1e-8 ^ (1/3), the result repeats every 5 iterations starting from iteration 24
	// (i.e. 24, 29, 34, ... give the same result) and never converges enough. The maximum number of
	// iterations (100) causes the result at iteration 100 to be returned, regardless of convergence.

	for i, tc := range testCases {
		res, err := tc.input.ApproxRoot(tc.root)
		s.Require().NoError(err)
		s.Require().True(tc.expected.Sub(res).AbsMut().LTE(ixomath.SmallestBigDec()), "unexpected result for test case %d, input: %v", i, tc.input)
	}
}

func (s *decimalTestSuite) TestApproxSqrt() {
	testCases := []struct {
		input    ixomath.BigDec
		expected ixomath.BigDec
	}{
		{ixomath.OneBigDec(), ixomath.OneBigDec()},                                                                               // 1.0 => 1.0
		{ixomath.NewBigDecWithPrec(25, 2), ixomath.NewBigDecWithPrec(5, 1)},                                                      // 0.25 => 0.5
		{ixomath.NewBigDecWithPrec(4, 2), ixomath.NewBigDecWithPrec(2, 1)},                                                       // 0.09 => 0.3
		{ixomath.NewBigDecFromInt(ixomath.NewBigInt(9)), ixomath.NewBigDecFromInt(ixomath.NewBigInt(3))},                         // 9 => 3
		{ixomath.NewBigDecFromInt(ixomath.NewBigInt(-9)), ixomath.NewBigDecFromInt(ixomath.NewBigInt(-3))},                       // -9 => -3
		{ixomath.NewBigDecFromInt(ixomath.NewBigInt(2)), ixomath.MustNewBigDecFromStr("1.414213562373095048801688724209698079")}, // 2 => 1.414213562373095048801688724209698079
	}

	for i, tc := range testCases {
		res, err := tc.input.ApproxSqrt()
		s.Require().NoError(err)
		s.Require().Equal(tc.expected, res, "unexpected result for test case %d, input: %v", i, tc.input)
	}
}

func (s *decimalTestSuite) TestDecEncoding() {
	testCases := []struct {
		input   ixomath.BigDec
		rawBz   string
		jsonStr string
		yamlStr string
	}{
		{
			ixomath.NewBigDec(0), "30",
			"\"0.000000000000000000000000000000000000\"",
			"\"0.000000000000000000000000000000000000\"\n",
		},
		{
			ixomath.NewBigDecWithPrec(4, 2),
			"3430303030303030303030303030303030303030303030303030303030303030303030",
			"\"0.040000000000000000000000000000000000\"",
			"\"0.040000000000000000000000000000000000\"\n",
		},
		{
			ixomath.NewBigDecWithPrec(-4, 2),
			"2D3430303030303030303030303030303030303030303030303030303030303030303030",
			"\"-0.040000000000000000000000000000000000\"",
			"\"-0.040000000000000000000000000000000000\"\n",
		},
		{
			ixomath.MustNewBigDecFromStr("1.414213562373095048801688724209698079"),
			"31343134323133353632333733303935303438383031363838373234323039363938303739",
			"\"1.414213562373095048801688724209698079\"",
			"\"1.414213562373095048801688724209698079\"\n",
		},
		{
			ixomath.MustNewBigDecFromStr("-1.414213562373095048801688724209698079"),
			"2D31343134323133353632333733303935303438383031363838373234323039363938303739",
			"\"-1.414213562373095048801688724209698079\"",
			"\"-1.414213562373095048801688724209698079\"\n",
		},
	}

	for _, tc := range testCases {
		bz, err := tc.input.Marshal()
		s.Require().NoError(err)
		s.Require().Equal(tc.rawBz, fmt.Sprintf("%X", bz))

		var other ixomath.BigDec
		s.Require().NoError((&other).Unmarshal(bz))
		s.Require().True(tc.input.Equal(other))

		bz, err = json.Marshal(tc.input)
		s.Require().NoError(err)
		s.Require().Equal(tc.jsonStr, string(bz))
		s.Require().NoError(json.Unmarshal(bz, &other))
		s.Require().True(tc.input.Equal(other))

		bz, err = yaml.Marshal(tc.input)
		s.Require().NoError(err)
		s.Require().Equal(tc.yamlStr, string(bz))
	}
}

// Showcase that different orders of operations causes different results.
func (s *decimalTestSuite) TestOperationOrders() {
	n1 := ixomath.NewBigDec(10)
	n2 := ixomath.NewBigDec(1000000010)
	s.Require().Equal(n1.Mul(n2).Quo(n2), ixomath.NewBigDec(10))
	s.Require().NotEqual(n1.Mul(n2).Quo(n2), n1.Quo(n2).Mul(n2))
}

func BenchmarkMarshalTo(b *testing.B) {
	b.ReportAllocs()
	bis := []struct {
		in   ixomath.BigDec
		want []byte
	}{
		{
			ixomath.NewBigDec(1e8), []byte{
				0x31, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
				0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
				0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
			},
		},
		{ixomath.NewBigDec(0), []byte{0x30}},
	}
	data := make([]byte, 100)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, bi := range bis {
			if n, err := bi.in.MarshalTo(data); err != nil {
				b.Fatal(err)
			} else {
				if !bytes.Equal(data[:n], bi.want) {
					b.Fatalf("Mismatch\nGot:  % x\nWant: % x\n", data[:n], bi.want)
				}
			}
		}
	}
}

func (s *decimalTestSuite) TestLog2() {
	var expectedErrTolerance = ixomath.MustNewBigDecFromStr("0.000000000000000000000000000000000100")

	tests := map[string]struct {
		initialValue ixomath.BigDec
		expected     ixomath.BigDec

		expectedPanic bool
	}{
		"log_2{-1}; invalid; panic": {
			initialValue:  ixomath.OneBigDec().Neg(),
			expectedPanic: true,
		},
		"log_2{0}; invalid; panic": {
			initialValue:  ixomath.ZeroBigDec(),
			expectedPanic: true,
		},
		"log_2{0.001} = -9.965784284662087043610958288468170528": {
			initialValue: ixomath.MustNewBigDecFromStr("0.001"),
			// From: https://www.wolframalpha.com/input?i=log+base+2+of+0.999912345+with+33+digits
			expected: ixomath.MustNewBigDecFromStr("-9.965784284662087043610958288468170528"),
		},
		"log_2{0.56171821941421412902170941} = -0.832081497183140708984033250637831402": {
			initialValue: ixomath.MustNewBigDecFromStr("0.56171821941421412902170941"),
			// From: https://www.wolframalpha.com/input?i=log+base+2+of+0.56171821941421412902170941+with+36+digits
			expected: ixomath.MustNewBigDecFromStr("-0.832081497183140708984033250637831402"),
		},
		"log_2{0.999912345} = -0.000126464976533858080645902722235833": {
			initialValue: ixomath.MustNewBigDecFromStr("0.999912345"),
			// From: https://www.wolframalpha.com/input?i=log+base+2+of+0.999912345+with+37+digits
			expected: ixomath.MustNewBigDecFromStr("-0.000126464976533858080645902722235833"),
		},
		"log_2{1} = 0": {
			initialValue: ixomath.NewBigDec(1),
			expected:     ixomath.NewBigDec(0),
		},
		"log_2{2} = 1": {
			initialValue: ixomath.NewBigDec(2),
			expected:     ixomath.NewBigDec(1),
		},
		"log_2{7} = 2.807354922057604107441969317231830809": {
			initialValue: ixomath.NewBigDec(7),
			// From: https://www.wolframalpha.com/input?i=log+base+2+of+7+37+digits
			expected: ixomath.MustNewBigDecFromStr("2.807354922057604107441969317231830809"),
		},
		"log_2{512} = 9": {
			initialValue: ixomath.NewBigDec(512),
			expected:     ixomath.NewBigDec(9),
		},
		"log_2{580} = 9.179909090014934468590092754117374938": {
			initialValue: ixomath.NewBigDec(580),
			// From: https://www.wolframalpha.com/input?i=log+base+2+of+600+37+digits
			expected: ixomath.MustNewBigDecFromStr("9.179909090014934468590092754117374938"),
		},
		"log_2{1024} = 10": {
			initialValue: ixomath.NewBigDec(1024),
			expected:     ixomath.NewBigDec(10),
		},
		"log_2{1024.987654321} = 10.001390817654141324352719749259888355": {
			initialValue: ixomath.NewBigDecWithPrec(1024987654321, 9),
			// From: https://www.wolframalpha.com/input?i=log+base+2+of+1024.987654321+38+digits
			expected: ixomath.MustNewBigDecFromStr("10.001390817654141324352719749259888355"),
		},
		"log_2{912648174127941279170121098210.92821920190204131121} = 99.525973560175362367047484597337715868": {
			initialValue: ixomath.MustNewBigDecFromStr("912648174127941279170121098210.92821920190204131121"),
			// From: https://www.wolframalpha.com/input?i=log+base+2+of+912648174127941279170121098210.92821920190204131121+38+digits
			expected: ixomath.MustNewBigDecFromStr("99.525973560175362367047484597337715868"),
		},
		"log_2{Max Spot Price} = 128": {
			initialValue: ixomath.BigDecFromDec(ixomath.MaxSpotPrice), // 2^128 - 1
			// From: https://www.wolframalpha.com/input?i=log+base+2+of+%28%282%5E128%29+-+1%29+38+digits
			expected: ixomath.MustNewBigDecFromStr("128"),
		},
		// The value tested below is: gammtypes.MaxSpotPrice * 0.99 = (2^128 - 1) * 0.99
		"log_2{336879543251729078828740861357450529340.45} = 127.98550043030488492336620207564264562": {
			initialValue: ixomath.MustNewBigDecFromStr("336879543251729078828740861357450529340.45"),
			// From: https://www.wolframalpha.com/input?i=log+base+2+of+%28%28%282%5E128%29+-+1%29*0.99%29++38+digits
			expected: ixomath.MustNewBigDecFromStr("127.98550043030488492336620207564264562"),
		},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			ixomath.ConditionalPanic(s.T(), tc.expectedPanic, func() {
				// Create a copy to test that the original was not modified.
				// That is, that LogbBase2() is non-mutative.
				initialCopy := tc.initialValue.Clone()

				res := tc.initialValue.LogBase2()
				require.True(ixomath.DecApproxEq(s.T(), tc.expected, res, expectedErrTolerance))
				require.Equal(s.T(), initialCopy, tc.initialValue)
			})
		})
	}
}

func (s *decimalTestSuite) TestLn() {
	var expectedErrTolerance = ixomath.MustNewBigDecFromStr("0.000000000000000000000000000000000100")

	tests := map[string]struct {
		initialValue ixomath.BigDec
		expected     ixomath.BigDec

		expectedPanic bool
	}{
		"log_e{-1}; invalid; panic": {
			initialValue:  ixomath.OneBigDec().Neg(),
			expectedPanic: true,
		},
		"log_e{0}; invalid; panic": {
			initialValue:  ixomath.ZeroBigDec(),
			expectedPanic: true,
		},
		"log_e{0.001} = -6.90775527898213705205397436405309262": {
			initialValue: ixomath.MustNewBigDecFromStr("0.001"),
			// From: https://www.wolframalpha.com/input?i=log0.001+to+36+digits+with+36+decimals
			expected: ixomath.MustNewBigDecFromStr("-6.90775527898213705205397436405309262"),
		},
		"log_e{0.56171821941421412902170941} = -0.576754943768592057376050794884207180": {
			initialValue: ixomath.MustNewBigDecFromStr("0.56171821941421412902170941"),
			// From: https://www.wolframalpha.com/input?i=log0.56171821941421412902170941+to+36+digits
			expected: ixomath.MustNewBigDecFromStr("-0.576754943768592057376050794884207180"),
		},
		"log_e{0.999912345} = -0.000087658841924023373535614212850888": {
			initialValue: ixomath.MustNewBigDecFromStr("0.999912345"),
			// From: https://www.wolframalpha.com/input?i=log0.999912345+to+32+digits
			expected: ixomath.MustNewBigDecFromStr("-0.000087658841924023373535614212850888"),
		},
		"log_e{1} = 0": {
			initialValue: ixomath.NewBigDec(1),
			expected:     ixomath.NewBigDec(0),
		},
		"log_e{e} = 1": {
			initialValue: ixomath.MustNewBigDecFromStr("2.718281828459045235360287471352662498"),
			// From: https://www.wolframalpha.com/input?i=e+with+36+decimals
			expected: ixomath.NewBigDec(1),
		},
		"log_e{7} = 1.945910149055313305105352743443179730": {
			initialValue: ixomath.NewBigDec(7),
			// From: https://www.wolframalpha.com/input?i=log7+up+to+36+decimals
			expected: ixomath.MustNewBigDecFromStr("1.945910149055313305105352743443179730"),
		},
		"log_e{512} = 6.238324625039507784755089093123589113": {
			initialValue: ixomath.NewBigDec(512),
			// From: https://www.wolframalpha.com/input?i=log512+up+to+36+decimals
			expected: ixomath.MustNewBigDecFromStr("6.238324625039507784755089093123589113"),
		},
		"log_e{580} = 6.36302810354046502061849560850445238": {
			initialValue: ixomath.NewBigDec(580),
			// From: https://www.wolframalpha.com/input?i=log580+up+to+36+decimals
			expected: ixomath.MustNewBigDecFromStr("6.36302810354046502061849560850445238"),
		},
		"log_e{1024.987654321} = 6.93243584693509415029056534690631614": {
			initialValue: ixomath.NewBigDecWithPrec(1024987654321, 9),
			// From: https://www.wolframalpha.com/input?i=log1024.987654321+to+36+digits
			expected: ixomath.MustNewBigDecFromStr("6.93243584693509415029056534690631614"),
		},
		"log_e{912648174127941279170121098210.92821920190204131121} = 68.986147965719214790400745338243805015": {
			initialValue: ixomath.MustNewBigDecFromStr("912648174127941279170121098210.92821920190204131121"),
			// From: https://www.wolframalpha.com/input?i=log912648174127941279170121098210.92821920190204131121+to+38+digits
			expected: ixomath.MustNewBigDecFromStr("68.986147965719214790400745338243805015"),
		},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			ixomath.ConditionalPanic(s.T(), tc.expectedPanic, func() {
				// Create a copy to test that the original was not modified.
				// That is, that Ln() is non-mutative.
				initialCopy := tc.initialValue.Clone()

				res := tc.initialValue.Ln()
				require.True(ixomath.DecApproxEq(s.T(), tc.expected, res, expectedErrTolerance))
				require.Equal(s.T(), initialCopy, tc.initialValue)
			})
		})
	}
}

func (s *decimalTestSuite) TestTickLog() {
	tests := map[string]struct {
		initialValue ixomath.BigDec
		expected     ixomath.BigDec

		expectedErrTolerance ixomath.BigDec
		expectedPanic        bool
	}{
		"log_1.0001{-1}; invalid; panic": {
			initialValue:  ixomath.OneBigDec().Neg(),
			expectedPanic: true,
		},
		"log_1.0001{0}; invalid; panic": {
			initialValue:  ixomath.ZeroBigDec(),
			expectedPanic: true,
		},
		"log_1.0001{0.001} = -69081.006609899112313305835611219486392199": {
			initialValue: ixomath.MustNewBigDecFromStr("0.001"),
			// From: https://www.wolframalpha.com/input?i=log_1.0001%280.001%29+to+41+digits
			expectedErrTolerance: ixomath.MustNewBigDecFromStr("0.000000000000000000000000000143031879"),
			expected:             ixomath.MustNewBigDecFromStr("-69081.006609899112313305835611219486392199"),
		},
		"log_1.0001{0.999912345} = -0.876632247930741919880461740717176538": {
			initialValue: ixomath.MustNewBigDecFromStr("0.999912345"),
			// From: https://www.wolframalpha.com/input?i=log_1.0001%280.999912345%29+to+36+digits
			expectedErrTolerance: ixomath.MustNewBigDecFromStr("0.000000000000000000000000000000138702"),
			expected:             ixomath.MustNewBigDecFromStr("-0.876632247930741919880461740717176538"),
		},
		"log_1.0001{1} = 0": {
			initialValue: ixomath.NewBigDec(1),

			expectedErrTolerance: ixomath.ZeroBigDec(),
			expected:             ixomath.NewBigDec(0),
		},
		"log_1.0001{1.0001} = 1": {
			initialValue: ixomath.MustNewBigDecFromStr("1.0001"),

			expectedErrTolerance: ixomath.MustNewBigDecFromStr("0.000000000000000000000000000000152500"),
			expected:             ixomath.OneBigDec(),
		},
		"log_1.0001{512} = 62386.365360724158196763710649998441051753": {
			initialValue: ixomath.NewBigDec(512),
			// From: https://www.wolframalpha.com/input?i=log_1.0001%28512%29+to+41+digits
			expectedErrTolerance: ixomath.MustNewBigDecFromStr("0.000000000000000000000000000129292137"),
			expected:             ixomath.MustNewBigDecFromStr("62386.365360724158196763710649998441051753"),
		},
		"log_1.0001{1024.987654321} = 69327.824629506998657531621822514042777198": {
			initialValue: ixomath.NewBigDecWithPrec(1024987654321, 9),
			// From: https://www.wolframalpha.com/input?i=log_1.0001%281024.987654321%29+to+41+digits
			expectedErrTolerance: ixomath.MustNewBigDecFromStr("0.000000000000000000000000000143836264"),
			expected:             ixomath.MustNewBigDecFromStr("69327.824629506998657531621822514042777198"),
		},
		"log_1.0001{912648174127941279170121098210.92821920190204131121} = 689895.972156319183538389792485913311778672": {
			initialValue: ixomath.MustNewBigDecFromStr("912648174127941279170121098210.92821920190204131121"),
			// From: https://www.wolframalpha.com/input?i=log_1.0001%28912648174127941279170121098210.92821920190204131121%29+to+42+digits
			expectedErrTolerance: ixomath.MustNewBigDecFromStr("0.000000000000000000000000001429936067"),
			expected:             ixomath.MustNewBigDecFromStr("689895.972156319183538389792485913311778672"),
		},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			ixomath.ConditionalPanic(s.T(), tc.expectedPanic, func() {
				// Create a copy to test that the original was not modified.
				// That is, that Ln() is non-mutative.
				initialCopy := tc.initialValue.Clone()

				res := tc.initialValue.TickLog()
				fmt.Println(name, res.Sub(tc.expected).Abs())
				require.True(ixomath.DecApproxEq(s.T(), tc.expected, res, tc.expectedErrTolerance))
				require.Equal(s.T(), initialCopy, tc.initialValue)
			})
		})
	}
}

func (s *decimalTestSuite) TestCustomBaseLog() {
	tests := map[string]struct {
		initialValue ixomath.BigDec
		base         ixomath.BigDec

		expected             ixomath.BigDec
		expectedErrTolerance ixomath.BigDec

		expectedPanic bool
	}{
		"log_2{-1}: normal base, invalid argument - panics": {
			initialValue:  ixomath.NewBigDec(-1),
			base:          ixomath.NewBigDec(2),
			expectedPanic: true,
		},
		"log_2{0}: normal base, invalid argument - panics": {
			initialValue:  ixomath.NewBigDec(0),
			base:          ixomath.NewBigDec(2),
			expectedPanic: true,
		},
		"log_(-1)(2): invalid base, normal argument - panics": {
			initialValue:  ixomath.NewBigDec(2),
			base:          ixomath.NewBigDec(-1),
			expectedPanic: true,
		},
		"log_1(2): base cannot equal to 1 - panics": {
			initialValue:  ixomath.NewBigDec(2),
			base:          ixomath.NewBigDec(1),
			expectedPanic: true,
		},
		"log_30(100) = 1.353984985057691049642502891262784015": {
			initialValue: ixomath.NewBigDec(100),
			base:         ixomath.NewBigDec(30),
			// From: https://www.wolframalpha.com/input?i=log_30%28100%29+to+37+digits
			expectedErrTolerance: ixomath.ZeroBigDec(),
			expected:             ixomath.MustNewBigDecFromStr("1.353984985057691049642502891262784015"),
		},
		"log_0.2(0.99) = 0.006244624769837438271878639001855450": {
			initialValue: ixomath.MustNewBigDecFromStr("0.99"),
			base:         ixomath.MustNewBigDecFromStr("0.2"),
			// From: https://www.wolframalpha.com/input?i=log_0.2%280.99%29+to+34+digits
			expectedErrTolerance: ixomath.MustNewBigDecFromStr("0.000000000000000000000000000000000013"),
			expected:             ixomath.MustNewBigDecFromStr("0.006244624769837438271878639001855450"),
		},

		"log_0.0001(500000) = -1.424742501084004701196565276318876743": {
			initialValue: ixomath.NewBigDec(500000),
			base:         ixomath.NewBigDecWithPrec(1, 4),
			// From: https://www.wolframalpha.com/input?i=log_0.0001%28500000%29+to+37+digits
			expectedErrTolerance: ixomath.MustNewBigDecFromStr("0.000000000000000000000000000000000003"),
			expected:             ixomath.MustNewBigDecFromStr("-1.424742501084004701196565276318876743"),
		},

		"log_500000(0.0001) = -0.701881216598197542030218906945601429": {
			initialValue: ixomath.NewBigDecWithPrec(1, 4),
			base:         ixomath.NewBigDec(500000),
			// From: https://www.wolframalpha.com/input?i=log_500000%280.0001%29+to+36+digits
			expectedErrTolerance: ixomath.MustNewBigDecFromStr("0.000000000000000000000000000000000001"),
			expected:             ixomath.MustNewBigDecFromStr("-0.701881216598197542030218906945601429"),
		},

		"log_10000(5000000) = 1.674742501084004701196565276318876743": {
			initialValue: ixomath.NewBigDec(5000000),
			base:         ixomath.NewBigDec(10000),
			// From: https://www.wolframalpha.com/input?i=log_10000%285000000%29+to+37+digits
			expectedErrTolerance: ixomath.MustNewBigDecFromStr("0.000000000000000000000000000000000002"),
			expected:             ixomath.MustNewBigDecFromStr("1.674742501084004701196565276318876743"),
		},
		"log_0.123456789(1) = 0": {
			initialValue: ixomath.OneBigDec(),
			base:         ixomath.MustNewBigDecFromStr("0.123456789"),

			expectedErrTolerance: ixomath.ZeroBigDec(),
			expected:             ixomath.ZeroBigDec(),
		},
		"log_1111(1111) = 1": {
			initialValue: ixomath.NewBigDec(1111),
			base:         ixomath.NewBigDec(1111),

			expectedErrTolerance: ixomath.ZeroBigDec(),
			expected:             ixomath.OneBigDec(),
		},

		"log_1.123{1024.987654321} = 59.760484327223888489694630378785099461": {
			initialValue: ixomath.NewBigDecWithPrec(1024987654321, 9),
			base:         ixomath.NewBigDecWithPrec(1123, 3),
			// From: https://www.wolframalpha.com/input?i=log_1.123%281024.987654321%29+to+38+digits
			expectedErrTolerance: ixomath.MustNewBigDecFromStr("0.000000000000000000000000000000007686"),
			expected:             ixomath.MustNewBigDecFromStr("59.760484327223888489694630378785099461"),
		},

		"log_1.123{912648174127941279170121098210.92821920190204131121} = 594.689327867863079177915648832621538986": {
			initialValue: ixomath.MustNewBigDecFromStr("912648174127941279170121098210.92821920190204131121"),
			base:         ixomath.NewBigDecWithPrec(1123, 3),
			// From: https://www.wolframalpha.com/input?i=log_1.123%28912648174127941279170121098210.92821920190204131121%29+to+39+digits
			expectedErrTolerance: ixomath.MustNewBigDecFromStr("0.000000000000000000000000000000077705"),
			expected:             ixomath.MustNewBigDecFromStr("594.689327867863079177915648832621538986"),
		},
	}
	for name, tc := range tests {
		s.Run(name, func() {
			ixomath.ConditionalPanic(s.T(), tc.expectedPanic, func() {
				// Create a copy to test that the original was not modified.
				// That is, that Ln() is non-mutative.
				initialCopy := tc.initialValue.Clone()
				res := tc.initialValue.CustomBaseLog(tc.base)
				require.True(ixomath.DecApproxEq(s.T(), tc.expected, res, tc.expectedErrTolerance))
				require.Equal(s.T(), initialCopy, tc.initialValue)
			})
		})
	}
}

func (s *decimalTestSuite) TestPowerInteger() {
	var expectedErrTolerance = ixomath.MustNewBigDecFromStr("0.000000000000000000000000000000100000")

	tests := map[string]struct {
		base           ixomath.BigDec
		exponent       uint64
		expectedResult ixomath.BigDec

		expectedToleranceOverwrite ixomath.BigDec
	}{
		"0^2": {
			base:     ixomath.ZeroBigDec(),
			exponent: 2,

			expectedResult: ixomath.ZeroBigDec(),
		},
		"1^2": {
			base:     ixomath.OneBigDec(),
			exponent: 2,

			expectedResult: ixomath.OneBigDec(),
		},
		"4^4": {
			base:     ixomath.MustNewBigDecFromStr("4"),
			exponent: 4,

			expectedResult: ixomath.MustNewBigDecFromStr("256"),
		},
		"5^3": {
			base:     ixomath.MustNewBigDecFromStr("5"),
			exponent: 4,

			expectedResult: ixomath.MustNewBigDecFromStr("625"),
		},
		"e^10": {
			base:     ixomath.EulersNumber,
			exponent: 10,

			// https://www.wolframalpha.com/input?i=e%5E10+41+digits
			expectedResult: ixomath.MustNewBigDecFromStr("22026.465794806716516957900645284244366354"),
		},
		"geom twap overflow: 2^log_2{max spot price + 1}": {
			base: ixomath.TwoBigDec,
			// add 1 for simplicity of calculation to isolate overflow.
			exponent: uint64(ixomath.BigDecFromDec(ixomath.MaxSpotPrice).Add(ixomath.OneBigDec()).LogBase2().TruncateInt().Uint64()),

			// https://www.wolframalpha.com/input?i=2%5E%28floor%28+log+base+2+%282%5E128%29%29%29+++39+digits
			expectedResult: ixomath.MustNewBigDecFromStr("340282366920938463463374607431768211456"),
		},
		"geom twap overflow: 2^log_2{max spot price}": {
			base:     ixomath.TwoBigDec,
			exponent: uint64(ixomath.BigDecFromDec(ixomath.MaxSpotPrice).LogBase2().TruncateInt().Uint64()),

			// https://www.wolframalpha.com/input?i=2%5E%28floor%28+log+base+2+%282%5E128+-+1%29%29%29+++39+digits
			expectedResult: ixomath.MustNewBigDecFromStr("170141183460469231731687303715884105728"),
		},
		"geom twap overflow: 2^log_2{max spot price / 2 - 2017}": { // 2017 is prime.
			base:     ixomath.TwoBigDec,
			exponent: uint64(ixomath.BigDecFromDec(ixomath.MaxSpotPrice.Quo(ixomath.NewDec(2)).Sub(ixomath.NewDec(2017))).LogBase2().TruncateInt().Uint64()),

			// https://www.wolframalpha.com/input?i=e%5E10+41+digits
			expectedResult: ixomath.MustNewBigDecFromStr("85070591730234615865843651857942052864"),
		},

		// ixomath.Dec test vectors copied from osmosis-labs/cosmos-sdk:

		"1.0 ^ (10) => 1.0": {
			base:     ixomath.OneBigDec(),
			exponent: 10,

			expectedResult: ixomath.OneBigDec(),
		},
		"0.5 ^ 2 => 0.25": {
			base:     ixomath.NewBigDecWithPrec(5, 1),
			exponent: 2,

			expectedResult: ixomath.NewBigDecWithPrec(25, 2),
		},
		"0.2 ^ 2 => 0.04": {
			base:     ixomath.NewBigDecWithPrec(2, 1),
			exponent: 2,

			expectedResult: ixomath.NewBigDecWithPrec(4, 2),
		},
		"3 ^ 3 => 27": {
			base:     ixomath.NewBigDec(3),
			exponent: 3,

			expectedResult: ixomath.NewBigDec(27),
		},
		"-3 ^ 4 = 81": {
			base:     ixomath.NewBigDec(-3),
			exponent: 4,

			expectedResult: ixomath.NewBigDec(81),
		},
		"-3 ^ 50 = 717897987691852588770249": {
			base:     ixomath.NewBigDec(-3),
			exponent: 50,

			expectedResult: ixomath.MustNewBigDecFromStr("717897987691852588770249"),
		},
		"-3 ^ 51 = -2153693963075557766310747": {
			base:     ixomath.NewBigDec(-3),
			exponent: 51,

			expectedResult: ixomath.MustNewBigDecFromStr("-2153693963075557766310747"),
		},
		"1.414213562373095049 ^ 2 = 2": {
			base:     ixomath.NewBigDecWithPrec(1414213562373095049, 18),
			exponent: 2,

			expectedResult:             ixomath.NewBigDec(2),
			expectedToleranceOverwrite: ixomath.MustNewBigDecFromStr("0.0000000000000000006"),
		},
	}

	for name, tc := range tests {
		tc := tc
		s.Run(name, func() {

			tolerance := expectedErrTolerance
			if !tc.expectedToleranceOverwrite.IsNil() {
				tolerance = tc.expectedToleranceOverwrite
			}

			// Main system under test
			actualResult := tc.base.PowerInteger(tc.exponent)
			require.True(ixomath.DecApproxEq(s.T(), tc.expectedResult, actualResult, tolerance))

			// Secondary system under test.
			// To reduce boilerplate from the same test cases when exponent is a
			// positive integer, we also test Power().
			// Negative exponent and base are not supported for Power()
			if tc.exponent >= 0 && !tc.base.IsNegative() {
				actualResult2 := tc.base.Power(ixomath.NewBigDecFromInt(ixomath.NewBigIntFromUint64(tc.exponent)))
				require.True(ixomath.DecApproxEq(s.T(), tc.expectedResult, actualResult2, tolerance))
			}
		})
	}
}

func (s *decimalTestSuite) TestClone() {
	tests := map[string]struct {
		startValue ixomath.BigDec
	}{
		"1.1": {
			startValue: ixomath.MustNewBigDecFromStr("1.1"),
		},
		"-3": {
			startValue: ixomath.MustNewBigDecFromStr("-3"),
		},
		"0": {
			startValue: ixomath.MustNewBigDecFromStr("-3"),
		},
	}

	for name, tc := range tests {
		tc := tc
		s.Run(name, func() {

			copy := tc.startValue.Clone()

			s.Require().Equal(tc.startValue, copy)

			copy.MulMut(ixomath.NewBigDec(2))
			// copy and startValue do not share internals.
			s.Require().NotEqual(tc.startValue, copy)
		})
	}
}

// TestMul_Mutation tests that MulMut mutates the receiver
// while Mut is not.
func (s *decimalTestSuite) TestMul_Mutation() {

	mulBy := ixomath.MustNewBigDecFromStr("2")

	tests := map[string]struct {
		startValue        ixomath.BigDec
		expectedMulResult ixomath.BigDec
	}{
		"1.1": {
			startValue:        ixomath.MustNewBigDecFromStr("1.1"),
			expectedMulResult: ixomath.MustNewBigDecFromStr("2.2"),
		},
		"-3": {
			startValue:        ixomath.MustNewBigDecFromStr("-3"),
			expectedMulResult: ixomath.MustNewBigDecFromStr("-6"),
		},
		"0": {
			startValue:        ixomath.ZeroBigDec(),
			expectedMulResult: ixomath.ZeroBigDec(),
		},
	}

	for name, tc := range tests {
		tc := tc
		s.Run(name, func() {
			startMut := tc.startValue.Clone()
			startNonMut := tc.startValue.Clone()

			resultMut := startMut.MulMut(mulBy)
			resultNonMut := startNonMut.Mul(mulBy)

			s.assertMutResult(tc.expectedMulResult, tc.startValue, resultMut, resultNonMut, startMut, startNonMut)
		})
	}
}

// TestPowerInteger_Mutation tests that PowerIntegerMut mutates the receiver
// while PowerInteger is not.
func (s *decimalTestSuite) TestPowerInteger_Mutation() {

	exponent := uint64(2)

	tests := map[string]struct {
		startValue     ixomath.BigDec
		expectedResult ixomath.BigDec
	}{
		"1": {
			startValue:     ixomath.OneBigDec(),
			expectedResult: ixomath.OneBigDec(),
		},
		"-3": {
			startValue:     ixomath.MustNewBigDecFromStr("-3"),
			expectedResult: ixomath.MustNewBigDecFromStr("9"),
		},
		"0": {
			startValue:     ixomath.ZeroBigDec(),
			expectedResult: ixomath.ZeroBigDec(),
		},
		"4": {
			startValue:     ixomath.MustNewBigDecFromStr("4.5"),
			expectedResult: ixomath.MustNewBigDecFromStr("20.25"),
		},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			startMut := tc.startValue.Clone()
			startNonMut := tc.startValue.Clone()

			resultMut := startMut.PowerIntegerMut(exponent)
			resultNonMut := startNonMut.PowerInteger(exponent)

			s.assertMutResult(tc.expectedResult, tc.startValue, resultMut, resultNonMut, startMut, startNonMut)
		})
	}
}

func (s *decimalTestSuite) TestPower() {
	tests := map[string]struct {
		base           ixomath.BigDec
		exponent       ixomath.BigDec
		expectedResult ixomath.BigDec
		expectPanic    bool
		errTolerance   ixomath.ErrTolerance
	}{
		// N.B.: integer exponents are tested under TestPowerInteger.

		"3 ^ 2 = 9 (integer base and integer exponent)": {
			base:     ixomath.NewBigDec(3),
			exponent: ixomath.NewBigDec(2),

			expectedResult: ixomath.NewBigDec(9),

			errTolerance: zeroAdditiveErrTolerance,
		},
		"2^0.5 (base of 2 and non-integer exponent)": {
			base:     ixomath.MustNewBigDecFromStr("2"),
			exponent: ixomath.MustNewBigDecFromStr("0.5"),

			// https://www.wolframalpha.com/input?i=2%5E0.5+37+digits
			expectedResult: ixomath.MustNewBigDecFromStr("1.414213562373095048801688724209698079"),

			errTolerance: ixomath.ErrTolerance{
				AdditiveTolerance: minDecTolerance,
				RoundingDir:       ixomath.RoundDown,
			},
		},
		"3^0.33 (integer base other than 2 and non-integer exponent)": {
			base:     ixomath.MustNewBigDecFromStr("3"),
			exponent: ixomath.MustNewBigDecFromStr("0.33"),

			// https://www.wolframalpha.com/input?i=3%5E0.33+37+digits
			expectedResult: ixomath.MustNewBigDecFromStr("1.436977652184851654252692986409357265"),

			errTolerance: ixomath.ErrTolerance{
				AdditiveTolerance: minDecTolerance,
				RoundingDir:       ixomath.RoundDown,
			},
		},
		"e^0.98999 (non-integer base and non-integer exponent)": {
			base:     ixomath.EulersNumber,
			exponent: ixomath.MustNewBigDecFromStr("0.9899"),

			// https://www.wolframalpha.com/input?i=e%5E0.9899+37+digits
			expectedResult: ixomath.MustNewBigDecFromStr("2.690965362357751196751808686902156603"),

			errTolerance: ixomath.ErrTolerance{
				AdditiveTolerance: minDecTolerance,
				RoundingDir:       ixomath.RoundUnconstrained,
			},
		},
		"10^0.001 (small non-integer exponent)": {
			base:     ixomath.NewBigDec(10),
			exponent: ixomath.MustNewBigDecFromStr("0.001"),

			// https://www.wolframalpha.com/input?i=10%5E0.001+37+digits
			expectedResult: ixomath.MustNewBigDecFromStr("1.002305238077899671915404889328110554"),

			errTolerance: ixomath.ErrTolerance{
				AdditiveTolerance: minDecTolerance,
				RoundingDir:       ixomath.RoundUnconstrained,
			},
		},
		"13^100.7777 (large non-integer exponent)": {
			base:     ixomath.NewBigDec(13),
			exponent: ixomath.MustNewBigDecFromStr("100.7777"),

			// https://www.wolframalpha.com/input?i=13%5E100.7777+37+digits
			expectedResult: ixomath.MustNewBigDecFromStr("1.822422110233759706998600329118969132").Mul(ixomath.NewBigDec(10).PowerInteger(112)),

			errTolerance: ixomath.ErrTolerance{
				MultiplicativeTolerance: minDecTolerance,
				RoundingDir:             ixomath.RoundDown,
			},
		},
		"large non-integer exponent with large non-integer base - panics": {
			base:     ixomath.MustNewBigDecFromStr("169.137"),
			exponent: ixomath.MustNewBigDecFromStr("100.7777"),

			expectPanic: true,
		},
		"negative base - panic": {
			base:     ixomath.NewBigDec(-3),
			exponent: ixomath.MustNewBigDecFromStr("4"),

			expectPanic: true,
		},
		"negative exponent - panic": {
			base:     ixomath.NewBigDec(1),
			exponent: ixomath.MustNewBigDecFromStr("-4"),

			expectPanic: true,
		},
		"base < 1 - panic (see godoc)": {
			base:     ixomath.NewBigDec(1).Sub(ixomath.SmallestBigDec()),
			exponent: ixomath.OneBigDec(),

			expectPanic: true,
		},
	}

	for name, tc := range tests {
		tc := tc
		s.Run(name, func() {
			ixomath.ConditionalPanic(s.T(), tc.expectPanic, func() {
				actualResult := tc.base.Power(tc.exponent)
				ixomath.Equal(s.T(), tc.errTolerance, tc.expectedResult, actualResult)
			})
		})
	}
}

func (s *decimalTestSuite) TestDec_WithPrecision() {
	tests := []struct {
		d         ixomath.BigDec
		want      ixomath.Dec
		precision uint64
		expPanic  bool
	}{
		// test cases for basic SDKDec() conversion
		{ixomath.NewBigDec(0), ixomath.MustNewDecFromStr("0.000000000000000000"), ixomath.DecPrecision, false},
		{ixomath.NewBigDec(1), ixomath.MustNewDecFromStr("1.000000000000000000"), ixomath.DecPrecision, false},
		{ixomath.NewBigDec(10), ixomath.MustNewDecFromStr("10.000000000000000000"), ixomath.DecPrecision, false},
		{ixomath.NewBigDec(12340), ixomath.MustNewDecFromStr("12340.000000000000000000"), ixomath.DecPrecision, false},
		{ixomath.NewBigDecWithPrec(12340, 4), ixomath.MustNewDecFromStr("1.234000000000000000"), ixomath.DecPrecision, false},
		{ixomath.NewBigDecWithPrec(12340, 5), ixomath.MustNewDecFromStr("0.123400000000000000"), ixomath.DecPrecision, false},
		{ixomath.NewBigDecWithPrec(12340, 8), ixomath.MustNewDecFromStr("0.000123400000000000"), ixomath.DecPrecision, false},
		{ixomath.NewBigDecWithPrec(1009009009009009009, 17), ixomath.MustNewDecFromStr("10.090090090090090090"), ixomath.DecPrecision, false},
		// test cases with custom precision:
		{ixomath.NewBigDec(0), ixomath.MustNewDecFromStr("0.000000000000"), 12, false},
		{ixomath.NewBigDec(1), ixomath.MustNewDecFromStr("1.000000000000"), 12, false},
		// specified precision is the same as the initial precision: 12.3453123123 -> 12.3453123123
		{ixomath.NewBigDecWithPrec(123453123123, 10), ixomath.MustNewDecFromStr("12.3453123123"), 10, false},
		// cut precision to 5 decimals: 3212.4623423462346 - 3212.46234
		{ixomath.NewBigDecWithPrec(32124623423462346, 13), ixomath.MustNewDecFromStr("3212.46234"), 5, false},
		// no decimal point: 18012004 -> 18012004
		{ixomath.NewBigDecWithPrec(18012004, 0), ixomath.MustNewDecFromStr("18012004"), 13, false},
		// if we try to convert to ixomath.Dec while specifying bigger precision than sdk.Dec has, panics
		{ixomath.NewBigDecWithPrec(1009009009009009009, 17), ixomath.MustNewDecFromStr("10.090090090090090090"), ixomath.DecPrecision + 2, true},
	}

	for tcIndex, tc := range tests {
		name := "testcase_" + fmt.Sprint(tcIndex)
		s.Run(name, func() {
			ixomath.ConditionalPanic(s.T(), tc.expPanic, func() {
				var got ixomath.Dec
				if tc.precision == ixomath.DecPrecision {
					got = tc.d.Dec()
				} else {
					got = tc.d.DecWithPrecision(tc.precision)
				}
				s.Require().Equal(tc.want, got, "bad Dec conversion, index: %v", tcIndex)
			})
		})
	}
}

func (s *decimalTestSuite) TestChopPrecision_Mutative() {
	tests := []struct {
		startValue        ixomath.BigDec
		expectedMutResult ixomath.BigDec
		precision         uint64
		expPanic          bool
	}{
		{ixomath.NewBigDec(0), ixomath.MustNewBigDecFromStr("0"), 0, false},
		{ixomath.NewBigDec(1), ixomath.MustNewBigDecFromStr("1"), 0, false},
		{ixomath.NewBigDec(10), ixomath.MustNewBigDecFromStr("10"), 2, false},
		// how to read these comments: ab.cde(fgh) -> ab.cdefgh = initial BigDec; (fgh) = decimal places that will be truncated
		// 5.1()
		{ixomath.NewBigDecWithPrec(51, 1), ixomath.MustNewBigDecFromStr("5.1"), 1, false},
		// 1.(0010)
		{ixomath.NewBigDecWithPrec(10010, 4), ixomath.MustNewBigDecFromStr("1"), 0, false},
		// 1009.31254(83952)
		{ixomath.NewBigDecWithPrec(10093125483952, 10), ixomath.MustNewBigDecFromStr("1009.31254"), 5, false},
		// 0.1009312548(3952)
		{ixomath.NewBigDecWithPrec(10093125483952, 14), ixomath.MustNewBigDecFromStr("0.1009312548"), 10, false},
		// Edge case: max precision. Should remain unchanged
		{ixomath.MustNewBigDecFromStr("1.000000000000000000000000000000000001"), ixomath.MustNewBigDecFromStr("1.000000000000000000000000000000000001"), ixomath.BigDecPrecision, false},
		// Precision exceeds max precision - panic
		{ixomath.MustNewBigDecFromStr("1.000000000000000000000000000000000001"), ixomath.MustNewBigDecFromStr("1.000000000000000000000000000000000001"), ixomath.BigDecPrecision + 1, true},
	}
	for id, tc := range tests {
		name := "testcase_" + fmt.Sprint(id)
		s.Run(name, func() {
			ixomath.ConditionalPanic(s.T(), tc.expPanic, func() {
				startMut := tc.startValue.Clone()
				startNonMut := tc.startValue.Clone()

				resultMut := startMut.ChopPrecisionMut(tc.precision)
				resultNonMut := startNonMut.ChopPrecision(tc.precision)

				s.assertMutResult(tc.expectedMutResult, tc.startValue, resultMut, resultNonMut, startMut, startNonMut)
			})
		})
	}
}
func (s *decimalTestSuite) TestQuoRoundUp_MutativeAndNonMutative() {
	tests := []struct {
		d1, d2, expQuoRoundUpMut ixomath.BigDec
	}{
		{ixomath.NewBigDec(0), ixomath.NewBigDec(0), ixomath.NewBigDec(0)},
		{ixomath.NewBigDec(1), ixomath.NewBigDec(0), ixomath.NewBigDec(0)},
		{ixomath.NewBigDec(0), ixomath.NewBigDec(1), ixomath.NewBigDec(0)},
		{ixomath.NewBigDec(0), ixomath.NewBigDec(-1), ixomath.NewBigDec(0)},
		{ixomath.NewBigDec(-1), ixomath.NewBigDec(0), ixomath.NewBigDec(0)},

		{ixomath.NewBigDec(1), ixomath.NewBigDec(1), ixomath.NewBigDec(1)},
		{ixomath.NewBigDec(-1), ixomath.NewBigDec(-1), ixomath.NewBigDec(1)},
		{ixomath.NewBigDec(1), ixomath.NewBigDec(-1), ixomath.NewBigDec(-1)},
		{ixomath.NewBigDec(-1), ixomath.NewBigDec(1), ixomath.NewBigDec(-1)},

		{
			ixomath.NewBigDec(3), ixomath.NewBigDec(7), ixomath.MustNewBigDecFromStr("0.428571428571428571428571428571428572"),
		},
		{
			ixomath.NewBigDec(2), ixomath.NewBigDec(4), ixomath.NewBigDecWithPrec(5, 1),
		},

		{ixomath.NewBigDec(100), ixomath.NewBigDec(100), ixomath.NewBigDec(1)},

		{
			ixomath.NewBigDecWithPrec(15, 1), ixomath.NewBigDecWithPrec(15, 1), ixomath.NewBigDec(1),
		},
		{
			ixomath.NewBigDecWithPrec(3333, 4), ixomath.NewBigDecWithPrec(333, 4), ixomath.MustNewBigDecFromStr("10.009009009009009009009009009009009010"),
		},
	}

	for tcIndex, tc := range tests {
		tc := tc
		name := "testcase_" + fmt.Sprint(tcIndex)
		s.Run(name, func() {
			ixomath.ConditionalPanic(s.T(), tc.d2.IsZero(), func() {
				copy := tc.d1.Clone()

				nonMutResult := copy.QuoRoundUp(tc.d2)

				// Return is as expected
				s.Require().Equal(tc.expQuoRoundUpMut, nonMutResult, "exp %v, res %v, tc %d", tc.expQuoRoundUpMut.String(), tc.d1.String(), tcIndex)

				// Receiver is not mutated
				s.Require().Equal(tc.d1, copy, "exp %v, res %v, tc %d", tc.expQuoRoundUpMut.String(), tc.d1.String(), tcIndex)

				// Receiver is mutated.
				tc.d1.QuoRoundUpMut(tc.d2)

				// Make sure d1 equals to expected
				s.Require().True(tc.expQuoRoundUpMut.Equal(tc.d1), "exp %v, res %v, tc %d", tc.expQuoRoundUpMut.String(), tc.d1.String(), tcIndex)
			})
		})
	}
}

func (s *decimalTestSuite) TestQuoTruncate_MutativeAndNonMutative() {
	tests := []struct {
		d1, d2, expQuoTruncateMut ixomath.BigDec
	}{
		{ixomath.NewBigDec(0), ixomath.NewBigDec(0), ixomath.NewBigDec(0)},
		{ixomath.NewBigDec(1), ixomath.NewBigDec(0), ixomath.NewBigDec(0)},
		{ixomath.NewBigDec(0), ixomath.NewBigDec(1), ixomath.NewBigDec(0)},
		{ixomath.NewBigDec(0), ixomath.NewBigDec(-1), ixomath.NewBigDec(0)},
		{ixomath.NewBigDec(-1), ixomath.NewBigDec(0), ixomath.NewBigDec(0)},

		{ixomath.NewBigDec(1), ixomath.NewBigDec(1), ixomath.NewBigDec(1)},
		{ixomath.NewBigDec(-1), ixomath.NewBigDec(-1), ixomath.NewBigDec(1)},
		{ixomath.NewBigDec(1), ixomath.NewBigDec(-1), ixomath.NewBigDec(-1)},
		{ixomath.NewBigDec(-1), ixomath.NewBigDec(1), ixomath.NewBigDec(-1)},

		{
			ixomath.NewBigDec(3), ixomath.NewBigDec(7), ixomath.MustNewBigDecFromStr("0.428571428571428571428571428571428571"),
		},
		{
			ixomath.NewBigDec(2), ixomath.NewBigDec(4), ixomath.NewBigDecWithPrec(5, 1),
		},

		{ixomath.NewBigDec(100), ixomath.NewBigDec(100), ixomath.NewBigDec(1)},

		{
			ixomath.NewBigDecWithPrec(15, 1), ixomath.NewBigDecWithPrec(15, 1), ixomath.NewBigDec(1),
		},
		{
			ixomath.NewBigDecWithPrec(3333, 4), ixomath.NewBigDecWithPrec(333, 4), ixomath.MustNewBigDecFromStr("10.009009009009009009009009009009009009"),
		},
	}

	for tcIndex, tc := range tests {
		tc := tc

		name := "testcase_" + fmt.Sprint(tcIndex)
		s.Run(name, func() {
			ixomath.ConditionalPanic(s.T(), tc.d2.IsZero(), func() {
				copy := tc.d1.Clone()

				nonMutResult := copy.QuoTruncate(tc.d2)

				// Return is as expected
				s.Require().Equal(tc.expQuoTruncateMut, nonMutResult, "exp %v, res %v, tc %d", tc.expQuoTruncateMut.String(), tc.d1.String(), tcIndex)

				// Receiver is not mutated
				s.Require().Equal(tc.d1, copy, "exp %v, res %v, tc %d", tc.expQuoTruncateMut.String(), tc.d1.String(), tcIndex)

				// Receiver is mutated.
				tc.d1.QuoTruncateMut(tc.d2)

				// Make sure d1 equals to expected
				s.Require().True(tc.expQuoTruncateMut.Equal(tc.d1), "exp %v, res %v, tc %d", tc.expQuoTruncateMut.String(), tc.d1.String(), tcIndex)
			})
		})
	}
}

func (s *decimalTestSuite) TestBigIntMut() {
	r := big.NewInt(30)
	d := ixomath.NewBigDecFromBigInt(r)

	// Compare value of BigInt & BigIntMut
	s.Require().Equal(d.BigInt(), d.BigIntMut())

	// Modify BigIntMut() pointer and ensure i.BigIntMut() & i.BigInt() change
	p1 := d.BigIntMut()
	p1.SetInt64(40)
	s.Require().Equal(big.NewInt(40), d.BigIntMut())
	s.Require().Equal(big.NewInt(40), d.BigInt())

	// Modify big.Int() pointer and ensure i.BigIntMut() & i.BigInt() don't change
	p2 := d.BigInt()
	p2.SetInt64(50)
	s.Require().NotEqual(big.NewInt(50), d.BigIntMut())
	s.Require().NotEqual(big.NewInt(50), d.BigInt())
}

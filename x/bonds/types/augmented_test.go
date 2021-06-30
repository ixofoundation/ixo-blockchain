package types

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func arange(start, stop, step float64) []float64 {
	N := int(math.Ceil((stop - start) / step))
	rnge := make([]float64, N, N)
	i := 0
	for x := start; x < stop; x += step {
		rnge[i] = x
		i += 1
	}
	return rnge
}

func printLines(title string, values []sdk.Dec) {
	print(title + " = [")
	for i, value := range values {
		index := strings.Index(value.String(), ".") + 7
		if i == len(values)-1 {
			fmt.Print(value.String()[:index])
		} else {
			fmt.Print(value.String()[:index] + ", ")
		}
	}
	println("]")
}

func TestExample1(t *testing.T) {
	d0 := sdk.MustNewDecFromStr("500.0")  // initial raise (reserve)
	p0 := sdk.MustNewDecFromStr("0.01")   // initial price (reserve per token)
	theta := sdk.MustNewDecFromStr("0.4") // funding fee fraction

	R0 := d0.Mul(sdk.OneDec().Sub(theta)) // initial reserve (raise minus funding)
	S0 := d0.Quo(p0)                      // initial supply

	kappa := sdk.NewDec(3)              // price exponent
	V0, err := Invariant(R0, S0, kappa) // invariant
	require.NoError(t, err)

	expectedR0 := sdk.MustNewDecFromStr("300.0")
	expectedS0 := sdk.MustNewDecFromStr("50000.0")
	expectedV0 := sdk.MustNewDecFromStr("416666666666.666667491283234022")

	require.Equal(t, expectedR0, R0)
	require.Equal(t, expectedS0, S0)
	require.Equal(t, expectedV0, V0)

	reserveF64 := arange(0, 100, .01)
	reserve := make([]sdk.Dec, len(reserveF64))
	for i, r := range reserveF64 {
		reserve[i] = sdk.MustNewDecFromStr(fmt.Sprintf("%f", r))
	}

	supp := make([]sdk.Dec, len(reserve))
	for i, r := range reserve {
		supp[i], err = Supply(r, kappa, V0)
		require.NoError(t, err)
	}

	price := make([]sdk.Dec, len(reserve))
	for i, r := range reserve {
		price[i], err = SpotPrice(r, kappa, V0)
		require.NoError(t, err)
	}

	printLines("reserve", reserve)
	printLines("supp", supp)
	printLines("price", price)
}

func TestReserve(t *testing.T) {
	decimals := sdk.NewDec(1000000000) // 1e9
	testCases := []struct {
		reserve sdk.Dec
		kappa   sdk.Dec
		V0      sdk.Dec
	}{
		{sdk.MustNewDecFromStr("0.05"), sdk.NewDec(1), sdk.MustNewDecFromStr("12345678.12345678")},
		{sdk.MustNewDecFromStr("5"), sdk.NewDec(2), sdk.MustNewDecFromStr("123456.123456")},
		{sdk.MustNewDecFromStr("500.500"), sdk.NewDec(3), sdk.MustNewDecFromStr("50000.50000")},
		{sdk.MustNewDecFromStr("50000.50000"), sdk.NewDec(4), sdk.MustNewDecFromStr("500.500")},
		{sdk.MustNewDecFromStr("123456.123456"), sdk.NewDec(5), sdk.MustNewDecFromStr("5")},
		{sdk.MustNewDecFromStr("12345678.12345678"), sdk.NewDec(6), sdk.MustNewDecFromStr("0.05")},
	}
	for _, tc := range testCases {
		calculatedSupply, err := Supply(tc.reserve, tc.kappa, tc.V0)
		require.NoError(t, err)
		calculatedReserve, err := Reserve(calculatedSupply, tc.kappa, tc.V0)
		require.NoError(t, err)

		// Keep 9DP (multiply by 1e9) and round, to ignore any errors beyond 9DP
		expectedReserveInt := tc.reserve.Mul(decimals).RoundInt64()
		calculatedReserveInt := calculatedReserve.Mul(decimals).RoundInt64()

		require.Equal(t, expectedReserveInt, calculatedReserveInt)
	}
}

func TestRationalPower(t *testing.T) {
	testCases := []struct {
		x      sdk.Dec
		xFloat float64
		a      uint64
		b      uint64
	}{
		{sdk.MustNewDecFromStr("5"), 5, 3, 2},
		{sdk.MustNewDecFromStr("500"), 500, 3, 2},
		{sdk.MustNewDecFromStr("50000"), 50000, 3, 2},
		{sdk.MustNewDecFromStr("5000000"), 5000000, 3, 2},
		{sdk.MustNewDecFromStr("5"), 5, 30, 2},
		//{sdk.MustNewDecFromStr("500"), 500, 300, 2}, // Int overflow
	}
	for _, tc := range testCases {
		expectedYStr := fmt.Sprintf("%.5f",
			math.Pow(tc.xFloat, float64(tc.a)/float64(tc.b)))

		temp, err := tc.x.ApproxRoot(tc.b)
		require.Nil(t, err)
		y := temp.Power(tc.a)

		y = sdk.NewDecFromInt(
			y.MulInt64(100000).RoundInt()).QuoInt64(100000)
		yFlt, err := strconv.ParseFloat(y.String(), 64)
		require.Nil(t, err)
		yStr := fmt.Sprintf("%.5f", yFlt)

		require.Equal(t, expectedYStr, yStr)
	}
}

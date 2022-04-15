package types

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

func TestRoundReservePrice(t *testing.T) {
	token := "token"

	// In general, RoundReservePrice rounds up

	testCases := []struct {
		in  string
		out int64
	}{{"9", 9}, {"1.6", 2}, {"0.5", 1}, {"0.4", 1}, {"0", 0}}
	for _, tc := range testCases {
		inDec := sdk.NewDecCoinFromDec(token, sdk.MustNewDecFromStr(tc.in))
		outInt := sdk.NewCoin(token, sdk.NewInt(tc.out))
		require.Equal(t, outInt, RoundReservePrice(inDec))
	}
}

func TestRoundReservePricesRoundsAllValues(t *testing.T) {
	tokens := []string{"token1", "token2", "token3"}
	ins := sdk.DecCoins{
		sdk.NewDecCoinFromDec(tokens[0], sdk.MustNewDecFromStr("0.4")),
		sdk.NewDecCoinFromDec(tokens[1], sdk.MustNewDecFromStr("1.6")),
		sdk.NewDecCoinFromDec(tokens[2], sdk.MustNewDecFromStr("3")),
	}
	outs := sdk.Coins{
		sdk.NewInt64Coin(tokens[0], 1),
		sdk.NewInt64Coin(tokens[1], 2),
		sdk.NewInt64Coin(tokens[2], 3),
	}
	require.True(t, RoundReservePrices(ins).IsEqual(outs))
}

func TestRoundReserveReturn(t *testing.T) {
	token := "token"

	// In general, RoundReserveReturn rounds down

	testCases := []struct {
		in  string
		out int64
	}{{"5", 5}, {"1.4", 1}, {"1.9", 1}, {"0.5", 0}, {"0", 0}}
	for _, tc := range testCases {
		inDec := sdk.NewDecCoinFromDec(token, sdk.MustNewDecFromStr(tc.in))
		outInt := sdk.NewCoin(token, sdk.NewInt(tc.out))
		require.True(t, outInt.IsEqual(RoundReserveReturn(inDec)))
	}
}

func TestRoundReserveReturnsRoundsAllValues(t *testing.T) {
	tokens := []string{"token1", "token2", "token3"}
	ins := sdk.DecCoins{
		sdk.NewDecCoinFromDec(tokens[0], sdk.MustNewDecFromStr("0.4")),
		sdk.NewDecCoinFromDec(tokens[1], sdk.MustNewDecFromStr("1.6")),
		sdk.NewDecCoinFromDec(tokens[2], sdk.MustNewDecFromStr("3")),
	}
	outs := sdk.Coins{
		// 0token1
		sdk.NewInt64Coin(tokens[1], 1),
		sdk.NewInt64Coin(tokens[2], 3),
	}
	require.True(t, RoundReserveReturns(ins).IsEqual(outs))
}

func TestMultiplyDecCoinByDec(t *testing.T) {
	token := "token"
	testCases := []struct {
		inCoin string
		scale  string
		out    string
	}{
		{"2", "2", "4"},        // all integers
		{"0.5", "2", "1"},      // result is integer
		{"2", "0.5", "1"},      // result is integer
		{"1.5", "0.5", "0.75"}, // all numbers decimal
		{"5", "1", "5"},        // N x 1 = N
		{"5", "0", "0"},        // N x 0 = 0
	}
	for _, tc := range testCases {
		inDec := sdk.NewDecCoinFromDec(token, sdk.MustNewDecFromStr(tc.inCoin))
		scaleDec := sdk.MustNewDecFromStr(tc.scale)
		outDec := sdk.NewDecCoinFromDec(token, sdk.MustNewDecFromStr(tc.out))
		require.True(t, outDec.IsEqual(MultiplyDecCoinByDec(inDec, scaleDec)))
	}
}

func TestMultiplyDecCoinsByDecMultipliesAllValues(t *testing.T) {
	tokens := []string{"token1", "token2", "token3", "token4"}
	ins := sdk.DecCoins{
		sdk.NewDecCoinFromDec(tokens[0], sdk.MustNewDecFromStr("0")),
		sdk.NewDecCoinFromDec(tokens[1], sdk.MustNewDecFromStr("0.5")),
		sdk.NewDecCoinFromDec(tokens[2], sdk.MustNewDecFromStr("1")),
		sdk.NewDecCoinFromDec(tokens[3], sdk.MustNewDecFromStr("2")),
	}.Sort()
	scaleDec := sdk.MustNewDecFromStr("0.5")
	outs := sdk.DecCoins{
		// 0token1
		sdk.NewDecCoinFromDec(tokens[1], sdk.MustNewDecFromStr("0.25")),
		sdk.NewDecCoinFromDec(tokens[2], sdk.MustNewDecFromStr("0.5")),
		sdk.NewDecCoinFromDec(tokens[3], sdk.MustNewDecFromStr("1")),
	}.Sort()
	require.True(t, MultiplyDecCoinsByDec(ins, scaleDec).IsEqual(outs))
}

func TestMultiplyDecCoinByInt(t *testing.T) {
	token := "token"
	testCases := []struct {
		inCoin string
		scale  int64
		out    string
	}{
		{"2", 2, "4"},   // all integers
		{"0.5", 2, "1"}, // result is integer
		{"5", 1, "5"},   // N x 1 = N
		{"5", 0, "0"},   // N x 0 = 0
	}
	for _, tc := range testCases {
		inDec := sdk.NewDecCoinFromDec(token, sdk.MustNewDecFromStr(tc.inCoin))
		scaleInt := sdk.NewInt(tc.scale)
		outDec := sdk.NewDecCoinFromDec(token, sdk.MustNewDecFromStr(tc.out))
		require.True(t, outDec.IsEqual(MultiplyDecCoinByInt(inDec, scaleInt)))
	}
}

func TestMultiplyDecCoinsByIntMultipliesAllValues(t *testing.T) {
	tokens := []string{"token1", "token2", "token3", "token4"}
	ins := sdk.DecCoins{
		sdk.NewDecCoinFromDec(tokens[0], sdk.MustNewDecFromStr("0")),
		sdk.NewDecCoinFromDec(tokens[1], sdk.MustNewDecFromStr("0.25")),
		sdk.NewDecCoinFromDec(tokens[2], sdk.MustNewDecFromStr("0.5")),
		sdk.NewDecCoinFromDec(tokens[3], sdk.MustNewDecFromStr("1")),
	}.Sort()
	scaleInt := sdk.NewInt(2)
	outs := sdk.DecCoins{
		// 0token1
		sdk.NewDecCoinFromDec(tokens[1], sdk.MustNewDecFromStr("0.5")),
		sdk.NewDecCoinFromDec(tokens[2], sdk.MustNewDecFromStr("1")),
		sdk.NewDecCoinFromDec(tokens[3], sdk.MustNewDecFromStr("2")),
	}.Sort()
	require.True(t, MultiplyDecCoinsByInt(ins, scaleInt).IsEqual(outs))
}

func TestDivideDecCoinByDec(t *testing.T) {
	token := "token"
	testCases := []struct {
		inCoin string
		scale  string
		out    string
	}{
		{"4", "2", "2"},       // all integers
		{"1", "2", "0.5"},     // result is decimal
		{"1.5", "0.5", "3"},   // result is integer
		{"1.5", "2.5", "0.6"}, // all numbers decimal
		{"5", "5", "1"},       // N / N = 1
		//{"5", "0", "0"},       // N / 0 = ? (panics)
	}
	for _, tc := range testCases {
		inDec := sdk.NewDecCoinFromDec(token, sdk.MustNewDecFromStr(tc.inCoin))
		scaleDec := sdk.MustNewDecFromStr(tc.scale)
		outDec := sdk.NewDecCoinFromDec(token, sdk.MustNewDecFromStr(tc.out))
		require.True(t, outDec.IsEqual(DivideDecCoinByDec(inDec, scaleDec)))
	}
}

func TestMultiplyDecCoinsByDecDividesAllValues(t *testing.T) {
	tokens := []string{"token1", "token2", "token3", "token4"}
	ins := sdk.DecCoins{
		sdk.NewDecCoinFromDec(tokens[0], sdk.MustNewDecFromStr("0")),
		sdk.NewDecCoinFromDec(tokens[1], sdk.MustNewDecFromStr("0.5")),
		sdk.NewDecCoinFromDec(tokens[2], sdk.MustNewDecFromStr("1")),
		sdk.NewDecCoinFromDec(tokens[3], sdk.MustNewDecFromStr("2")),
	}.Sort()
	scaleDec := sdk.MustNewDecFromStr("2")
	outs := sdk.DecCoins{
		// 0token1
		sdk.NewDecCoinFromDec(tokens[1], sdk.MustNewDecFromStr("0.25")),
		sdk.NewDecCoinFromDec(tokens[2], sdk.MustNewDecFromStr("0.5")),
		sdk.NewDecCoinFromDec(tokens[3], sdk.MustNewDecFromStr("1")),
	}.Sort()
	require.True(t, DivideDecCoinsByDec(ins, scaleDec).IsEqual(outs))
}

func TestRoundFee(t *testing.T) {
	token := "token"

	// In general, RoundFee rounds up

	testCases := []struct {
		in  string
		out int64
	}{{"7", 7}, {"0.4", 1}, {"67.7", 68}, {"96.5", 97}, {"0", 0}}
	for _, tc := range testCases {
		inDec := sdk.NewDecCoinFromDec(token, sdk.MustNewDecFromStr(tc.in))
		outInt := sdk.NewCoin(token, sdk.NewInt(tc.out))
		require.True(t, outInt.IsEqual(RoundFee(inDec)))
	}
}

func TestStringsToString(t *testing.T) {
	testCases := []struct {
		in  []string
		out string
	}{
		{[]string{}, "[]"},
		{[]string{"str1"}, "[str1]"},
		{[]string{"str1", "str2", "str3"}, "[str1,str2,str3]"},
	}
	for _, tc := range testCases {
		require.Equal(t, tc.out, StringsToString(tc.in))
	}
}

func TestAccAddressesToString(t *testing.T) {

	oneIn := []sdk.AccAddress{
		sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address()),
	}
	threeIn := []sdk.AccAddress{
		sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address()),
		sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address()),
		sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address()),
	}

	testCases := []struct {
		in  []sdk.AccAddress
		out string
	}{
		{[]sdk.AccAddress{}, "[]"},
		{oneIn, fmt.Sprintf("[%s]", oneIn[0])},
		{threeIn, fmt.Sprintf("[%s,%s,%s]", threeIn[0], threeIn[1], threeIn[2])},
	}
	for _, tc := range testCases {
		require.Equal(t, tc.out, AccAddressesToString(tc.in))
	}
}

func TestApproxPower(t *testing.T) {

	testCases := []struct {
		in1            string
		in2            string
		expectedResult string
	}{
		{"0", "5", "0"},
		{"1", "5", "1"},
		{"5", "0", "1"},
		{"1", "0", "1"},
		{"2", "2", "4.000000000000000001"},
		{"2", "4", "15.999999999999999989"},
		{"2.5", "4.6", "67.689926089369764941"},
		{"4.6", "2.5", "45.383144007439590368"},
	}
	for _, tc := range testCases {
		in1 := sdk.MustNewDecFromStr(tc.in1)
		in2 := sdk.MustNewDecFromStr(tc.in2)
		expectedResult := sdk.MustNewDecFromStr(tc.expectedResult)

		actualResult, err := ApproxPower(in1, in2)
		require.NoError(t, err)
		require.Equal(t, expectedResult, actualResult)
	}
}

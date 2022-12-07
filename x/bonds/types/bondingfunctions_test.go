package types

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type simFunc func(*AugmentedBondRevision1, AugmentedBondRevision1) (AugmentedBondRevision1, error)

func buyToken(amount int64, stateF func(state AugmentedBondRevision1) AugmentedBondRevision1) func(*AugmentedBondRevision1, AugmentedBondRevision1) (AugmentedBondRevision1, error) {
	return func(algo *AugmentedBondRevision1, state AugmentedBondRevision1) (AugmentedBondRevision1, error) {
		priceCoin, _ := algo.CalculateTokensForPrice(sdk.NewInt64Coin("bond", amount))
		price, _ := priceCoin.Amount.Float64()
		// fmt.Println("Buy price:", price)
		algo._R = algo._R + price
		algo._m = algo._m + float64(amount)
		return stateF(state), nil
	}
}

func updateAlpha(publicAloha float64, stateF func(state AugmentedBondRevision1) AugmentedBondRevision1) func(*AugmentedBondRevision1, AugmentedBondRevision1) (AugmentedBondRevision1, error) {
	return func(algo *AugmentedBondRevision1, state AugmentedBondRevision1) (AugmentedBondRevision1, error) {
		algo.UpdateAlpha(publicAloha, 0.5)
		return stateF(state), nil
	}
}

// type simFunc func(*AugmentedBondRevision1, AugmentedBondRevision1) (AugmentedBondRevision1, error)

// func buyToken(amount int64, stateF func(state AugmentedBondRevision1) AugmentedBondRevision1) func(*AugmentedBondRevision1, AugmentedBondRevision1) (AugmentedBondRevision1, error) {
// 	return func(algo *AugmentedBondRevision1, state AugmentedBondRevision1) (AugmentedBondRevision1, error) {
// 		priceCoin, _ := algo.CalculateTokensForPrice(sdk.NewInt64Coin("bond", amount))
// 		price, _ := priceCoin.Amount.Float64()
// 		// fmt.Println("Buy price:", price)
// 		algo._R = algo._R + price
// 		algo._m = algo._m + float64(amount)
// 		return stateF(state), nil
// 	}
// }

// func updateAlpha(publicAloha float64, stateF func(state AugmentedBondRevision1) AugmentedBondRevision1) func(*AugmentedBondRevision1, AugmentedBondRevision1) (AugmentedBondRevision1, error) {
// 	return func(algo *AugmentedBondRevision1, state AugmentedBondRevision1) (AugmentedBondRevision1, error) {
// 		algo.UpdateAlpha(publicAloha)
// 		return stateF(state), nil
// 	}
// }

// func TestRejectBondCreation(t *testing.T) {

// }
func TestBondMustMatchRevision(t *testing.T) {

}

func TestBondWithoutMaturityOrAlphaUpdates2(t *testing.T) {

	token := "bond"
	reserveToken := "xusd"
	alphaBond := Bond{
		Token: token,
		FunctionParameters: []FunctionParam{
			NewFunctionParam("REVISION", sdk.NewDec(1)),
			NewFunctionParam("Funding_Target", sdk.NewDec(60000)),
			NewFunctionParam("Hatch_Supply", sdk.NewDec(0)),
			NewFunctionParam("Hatch_Price", sdk.NewDec(0)),
			NewFunctionParam("APY_MAX", sdk.NewDec(150)),
			NewFunctionParam("APY_MIN", sdk.NewDec(10)),
			NewFunctionParam("MATURITY", sdk.NewDec(1)),
			NewFunctionParam("DISCOUNT_RATE", sdk.NewDec(2)),
			NewFunctionParam("GAMMA", sdk.NewDec(2)),
			NewFunctionParam("INITIAL_PUBLIC_ALPHA", sdk.NewDec(50)),
		},
		CurrentSupply:            sdk.NewInt64Coin(token, 0),
		Name:                     "Test Bond",
		Description:              "Description",
		CreatorDid:               "did:test:0000000000000",
		ControllerDid:            "did:test:0000000000000",
		ReserveTokens:            []string{},
		FunctionType:             "augmented_bond_v2",
		TxFeePercentage:          sdk.ZeroDec(),
		ExitFeePercentage:        sdk.ZeroDec(),
		FeeAddress:               "address",
		ReserveWithdrawalAddress: "address",
		MaxSupply:                sdk.NewInt64Coin(token, 60000),
		OrderQuantityLimits:      sdk.Coins{},
		SanityRate:               sdk.ZeroDec(),
		SanityMarginPercentage:   sdk.ZeroDec(),
		CurrentReserve: sdk.Coins{
			sdk.NewInt64Coin(reserveToken, 0),
		},
		AvailableReserve: sdk.Coins{
			sdk.NewInt64Coin(reserveToken, 0),
		},
		CurrentOutcomePaymentReserve: sdk.Coins{
			sdk.NewInt64Coin(reserveToken, 0),
		},
		AllowSells:              false,
		AlphaBond:               true,
		AllowReserveWithdrawals: true,
		OutcomePayment:          sdk.NewInt(68100),
		State:                   "OPEN",
		BondDid:                 "did:test:0000000000000",
	}

	// ba, _ := InitializeBondingAlgorithm[AugmentedBond](bond)

	algo := AugmentedBondRevision1{}
	err := algo.Init(alphaBond)
	if err != nil {

	}
	// fmt.Println(alphaBond)
	// algo.UpdatePublicAlpha(0.5)
	// algo.UpdateRivsion()
	// algo.Export(alphaBond)

	// algo := &AugmentedBondRevision1{}
	// algo.Init(alphaBond)

	// pbond, _ := algo.CalculatePriceForTokens(sdk.NewInt64Coin(reserveToken, 20000))
	// fmt.Println("price for 20000 bond tokens:", pbond.Amount.String())

	// pxusd, _ := algo.CalculateTokensForPrice(sdk.NewInt64Coin(reserveToken, 4500))
	// fmt.Println("price for 4500 xusd:", pxusd.Amount.String())

	sim := []simFunc{
		buyToken(10000, func(state AugmentedBondRevision1) AugmentedBondRevision1 {
			state._m = 10000
			// state._R = state._R + 4199
			// state.price = 100000
			return state
		}),
		buyToken(10000, func(state AugmentedBondRevision1) AugmentedBondRevision1 {
			state._m = state._m + 10000
			// state.price = 100000
			return state
		}),
		buyToken(10000, func(state AugmentedBondRevision1) AugmentedBondRevision1 {
			state._m = state._m + 10000
			// state.price = 100000
			return state
		}),
		buyToken(10000, func(state AugmentedBondRevision1) AugmentedBondRevision1 {
			state._m = state._m + 10000
			// state.price = 100000
			return state
		}),
		buyToken(10000, func(state AugmentedBondRevision1) AugmentedBondRevision1 {
			state._m = state._m + 10000
			// state.price = 100000
			return state
		}),
		buyToken(10000, func(state AugmentedBondRevision1) AugmentedBondRevision1 {
			state._m = state._m + 10000
			// state.price = 100000
			return state
		}),
		buyToken(10000, func(state AugmentedBondRevision1) AugmentedBondRevision1 {
			state._m = state._m + 10000
			// state.price = 100000
			return state
		}),
		updateAlpha(0.6, func(state AugmentedBondRevision1) AugmentedBondRevision1 {
			// state._m = state._m + 10000
			// state.price = 100000
			return state
		}),
		buyToken(10000, func(state AugmentedBondRevision1) AugmentedBondRevision1 {
			state._m = state._m + 10000
			// state.price = 100000
			return state
		}),
		buyToken(10000, func(state AugmentedBondRevision1) AugmentedBondRevision1 {
			state._m = state._m + 10000
			// state.price = 100000
			return state
		}),
		buyToken(10000, func(state AugmentedBondRevision1) AugmentedBondRevision1 {
			state._m = state._m + 10000
			// state.price = 100000
			return state
		}),
		buyToken(10000, func(state AugmentedBondRevision1) AugmentedBondRevision1 {
			state._m = state._m + 10000
			// state.price = 100000
			return state
		}),
		buyToken(10000, func(state AugmentedBondRevision1) AugmentedBondRevision1 {
			state._m = state._m + 10000
			// state.price = 100000
			return state
		}),
		updateAlpha(0.7, func(state AugmentedBondRevision1) AugmentedBondRevision1 {
			// state._m = state._m + 10000
			// state.price = 100000
			return state
		}),
		buyToken(1000, func(state AugmentedBondRevision1) AugmentedBondRevision1 {
			state._m = state._m + 10000
			// state.price = 100000
			return state
		}),
		buyToken(1000, func(state AugmentedBondRevision1) AugmentedBondRevision1 {
			state._m = state._m + 10000
			// state.price = 100000
			return state
		}),
		buyToken(1000, func(state AugmentedBondRevision1) AugmentedBondRevision1 {
			state._m = state._m + 10000
			// state.price = 100000
			return state
		}),
		buyToken(1000, func(state AugmentedBondRevision1) AugmentedBondRevision1 {
			state._m = state._m + 10000
			// state.price = 100000
			return state
		}),
		updateAlpha(0.8, func(state AugmentedBondRevision1) AugmentedBondRevision1 {
			// state._m = state._m + 10000
			// state.price = 100000
			return state
		}),
		buyToken(1000, func(state AugmentedBondRevision1) AugmentedBondRevision1 {
			state._m = state._m + 10000
			// state.price = 100000
			return state
		}),
		buyToken(1000, func(state AugmentedBondRevision1) AugmentedBondRevision1 {
			state._m = state._m + 10000
			// state.price = 100000
			return state
		}),
		buyToken(400, func(state AugmentedBondRevision1) AugmentedBondRevision1 {
			state._m = state._m + 10000
			// state.price = 100000
			return state
		}),
		buyToken(1000, func(state AugmentedBondRevision1) AugmentedBondRevision1 {
			state._m = state._m + 10000
			// state.price = 100000
			return state
		}),
		buyToken(1000, func(state AugmentedBondRevision1) AugmentedBondRevision1 {
			state._m = state._m + 10000
			// state.price = 100000
			return state
		}),
		// updateAlpha(0.6, func(state AugmentedBondRevision1) {
		// 	state.alpha = 0.6
		// 	return state
		// }),
	}

	handleSim := func(t *testing.T, algo *AugmentedBondRevision1, sim []simFunc) {
		var nextState AugmentedBondRevision1 = *algo
		var err error

		for _, f := range sim {
			nextState, err = f(algo, nextState)
			if err != nil {
				t.Error(err)
			}
			fmt.Println()
			// fmt.Printf("\nSIM: %+v\n", nextState)
			// fmt.Println("SIM:", nextState == *algo)
			fmt.Printf("ALGO: %+v\n", algo)
			fmt.Printf("APYavg: %.2f\n", algo._APYavg)
		}
	}

	handleSim(t, &algo, sim)
}

// func TestBondWithAlphaUpdatesButWithoutMaturity(t *testing.T) {

// }

// func TestBondWithMaturityButWithoutAlphaUpdates(t *testing.T) {

// }

// func TestBondWithAlphaUpdatesAndMaturity(t *testing.T) {

// }

// 2.4
// func TestFiGeatherthan0() {}
// func TestMiGeatherthan0() {}

// func TestPmaxSmallerThanCoverM() {}

// func Test_r_GreaterThanOrEqual0andSmallerThanAPYmin() {}
// func Test_IncomeIsAvaialbleDueToTime()                {}

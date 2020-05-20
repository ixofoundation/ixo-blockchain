package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
)

type FeeType string

const (
	FeeClaimTransaction      FeeType = "ClaimTransactionFee"
	FeeEvaluationTransaction FeeType = "FeeEvaluationTransaction"
)

type Fee struct {
	Id                 string
	ChargeAmount       sdk.Coins
	ChargeMinimum      sdk.Coins
	ChargeMaximum      sdk.Coins
	DiscountId         string
	DiscountPercent    sdk.Dec
	WalletDistribution Distribution
}

type FeeContract struct {
	Id               string
	FeeId            string
	Payer            ixo.Did
	CumulativeCharge sdk.Coins
	CanDeauthorise   sdk.Coins
	Authorised       bool
}

type Distribution []DistributionShare

//IsValid Checks that shares total up to 100 percent
func (d Distribution) IsValid() bool {
	total := sdk.ZeroDec()
	for _, share := range d {
		total = total.Add(share.Percentage)
	}
	return total.Equal(sdk.NewDec(100))
}

type DistributionShare struct {
	Identifier string
	Percentage sdk.Dec
}

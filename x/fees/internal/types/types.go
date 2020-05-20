package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type FeeType string

const (
	FeeClaimTransaction      FeeType = "ClaimTransactionFee"
	FeeEvaluationTransaction FeeType = "FeeEvaluationTransaction"
)

type Fee struct {
	Id                 string       `json:"id" yaml:"id"`
	ChargeAmount       sdk.Coins    `json:"charge_amount" yaml:"charge_amount"`
	ChargeMinimum      sdk.Coins    `json:"charge_minimum" yaml:"charge_minimum"`
	ChargeMaximum      sdk.Coins    `json:"charge_maximum" yaml:"charge_maximum"`
	DiscountId         string       `json:"discount_id" yaml:"discount_id"`
	DiscountPercent    sdk.Dec      `json:"discount_percent" yaml:"discount_percent"`
	WalletDistribution Distribution `json:"wallet_distribution" yaml:"wallet_distribution"`
}

func NewFee(id string, chargeAmount, chargeMinimum, chargeMaximum sdk.Coins,
	discountId string, discountPercent sdk.Dec, walletDistribution Distribution) Fee {
	return Fee{
		Id:                 id,
		ChargeAmount:       chargeAmount,
		ChargeMinimum:      chargeMinimum,
		ChargeMaximum:      chargeMaximum,
		DiscountId:         discountId,
		DiscountPercent:    discountPercent,
		WalletDistribution: walletDistribution,
	}
}

func (f Fee) Validate() sdk.Error {
	// TODO: Validation of remaining fields

	return f.WalletDistribution.Validate()
}

type FeeContract struct {
	Id               string         `json:"id" yaml:""`
	FeeId            string         `json:"fee_id" yaml:""`
	Payer            sdk.AccAddress `json:"payer" yaml:""`
	CumulativeCharge sdk.Coins      `json:"cumulative_charge" yaml:"cumulative_charge"`
	CanDeauthorise   bool           `json:"can_deauthorise" yaml:"can_deauthorise"`
	Authorised       bool           `json:"authorised" yaml:"authorised"`
}

func NewFeeContract(id, feeId string, payer sdk.AccAddress,
	cumulativeCharge sdk.Coins, canDeauthorise, authorised bool) FeeContract {
	return FeeContract{
		Id:               id,
		FeeId:            feeId,
		Payer:            payer,
		CumulativeCharge: cumulativeCharge,
		CanDeauthorise:   canDeauthorise,
		Authorised:       authorised,
	}
}

type Distribution []DistributionShare

func NewDistribution(shares ...DistributionShare) Distribution {
	return Distribution(shares)
}

//IsValid Checks that shares total up to 100 percent
func (d Distribution) Validate() sdk.Error {
	// Validate shares and calculate total
	total := sdk.ZeroDec()
	for _, share := range d {
		total = total.Add(share.Percentage)
		if err := share.Validate(); err != nil {
			return err
		}
	}

	// Shares must add up to 100%
	if !total.Equal(sdk.NewDec(100)) {
		return ErrDistributionPercentagesNot100(DefaultCodespace, total)
	}

	return nil
}

type DistributionShare struct {
	Identifier string  `json:"identifier" yaml:"identifier"`
	Percentage sdk.Dec `json:"percentage" yaml:"percentage"`
}

func NewDistributionShare(identifier string, percentage sdk.Dec) DistributionShare {
	return DistributionShare{
		Identifier: identifier,
		Percentage: percentage,
	}
}

func (d DistributionShare) Validate() sdk.Error {
	// TODO: Identifier distribution

	if !d.Percentage.IsPositive() {
		return ErrNegativeSharePercentage(DefaultCodespace)
	}

	return nil
}

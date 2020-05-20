package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type FeeType string

const (
	FeeClaimTransaction      FeeType = "ClaimTransactionFee"
	FeeEvaluationTransaction FeeType = "FeeEvaluationTransaction"
)

type FeeContent struct {
	ChargeAmount       sdk.Coins    `json:"charge_amount" yaml:"charge_amount"`
	ChargeMinimum      sdk.Coins    `json:"charge_minimum" yaml:"charge_minimum"`
	ChargeMaximum      sdk.Coins    `json:"charge_maximum" yaml:"charge_maximum"`
	Discounts          Discounts    `json:"discounts" yaml:"discounts"`
	WalletDistribution Distribution `json:"wallet_distribution" yaml:"wallet_distribution"`
}

func NewFeeContent(chargeAmount, chargeMinimum, chargeMaximum sdk.Coins,
	discounts Discounts, walletDistribution Distribution) FeeContent {
	return FeeContent{
		ChargeAmount:       chargeAmount,
		ChargeMinimum:      chargeMinimum,
		ChargeMaximum:      chargeMaximum,
		Discounts:          discounts,
		WalletDistribution: walletDistribution,
	}
}

type Fee struct {
	Id      uint64     `json:"id" yaml:"id"`
	Content FeeContent `json:"content" yaml:"content"`
}

func NewFee(id uint64, content FeeContent) Fee {
	return Fee{
		Id:      id,
		Content: content,
	}
}

func (f Fee) Validate() sdk.Error {
	// TODO: Validation of remaining fields

	return f.Content.WalletDistribution.Validate()
}

type FeeContractContent struct {
	FeeId            uint64         `json:"fee_id" yaml:"fee_id"`
	Payer            sdk.AccAddress `json:"payer" yaml:"payer"`
	CumulativeCharge sdk.Coins      `json:"cumulative_charge" yaml:"cumulative_charge"`
	CanDeauthorise   bool           `json:"can_deauthorise" yaml:"can_deauthorise"`
	Authorised       bool           `json:"authorised" yaml:"authorised"`
}

func NewFeeContractContent(feeId uint64, payer sdk.AccAddress,
	cumulativeCharge sdk.Coins, canDeauthorise, authorised bool) FeeContractContent {
	return FeeContractContent{
		FeeId:            feeId,
		Payer:            payer,
		CumulativeCharge: cumulativeCharge,
		CanDeauthorise:   canDeauthorise,
		Authorised:       authorised,
	}
}

type FeeContract struct {
	Id      uint64             `json:"id" yaml:"id"`
	Content FeeContractContent `json:"content" yaml:"content"`
}

func NewFeeContract(id uint64, content FeeContractContent) FeeContract {
	return FeeContract{
		Id:      id,
		Content: content,
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

type Discounts []Discount

func NewDiscounts(discounts ...Discount) Discounts {
	return Discounts(discounts)
}

type Discount struct {
	Id      string  `json:"id" yaml:"id"`
	Percent sdk.Dec `json:"percent" yaml:"percent"`
}

func NewDiscount(id string, percent sdk.Dec) Discount {
	return Discount{
		Id:      id,
		Percent: percent,
	}
}

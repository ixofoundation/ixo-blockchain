package types

import sdk "github.com/cosmos/cosmos-sdk/types"

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
	Creator          sdk.AccAddress `json:"creator" yaml:"creator"`
	Payer            sdk.AccAddress `json:"payer" yaml:"payer"`
	CumulativeCharge sdk.Coins      `json:"cumulative_charge" yaml:"cumulative_charge"`
	CanDeauthorise   bool           `json:"can_deauthorise" yaml:"can_deauthorise"`
	Authorised       bool           `json:"authorised" yaml:"authorised"`
}

func NewFeeContractContent(feeId uint64, creator, payer sdk.AccAddress,
	cumulativeCharge sdk.Coins, canDeauthorise, authorised bool) FeeContractContent {
	return FeeContractContent{
		FeeId:            feeId,
		Creator:          creator,
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

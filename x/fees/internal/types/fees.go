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

func (fc FeeContent) Validate() sdk.Error {
	// Validate charge amount, minimum, maximum
	amt := &fc.ChargeAmount
	min := &fc.ChargeMinimum
	max := &fc.ChargeMaximum
	if min.IsValid() && max.IsValid() && amt.IsValid() {
		return ErrInvalidFee(DefaultCodespace, "min, max, or amt coins invalid")
	} else if min.IsAnyGT(*max) {
		return ErrInvalidFee(DefaultCodespace, "min charge includes value greater than max")
	} else if !min.DenomsSubsetOf(*amt) {
		return ErrInvalidFee(DefaultCodespace, "min charge includes denom not in fee amount")
	} else if !max.DenomsSubsetOf(*amt) {
		return ErrInvalidFee(DefaultCodespace, "max charge includes denom not in fee amount")
	}

	// Validate discounts
	if err := fc.Discounts.Validate(); err != nil {
		return err
	}

	// Validate wallet distribution
	if err := fc.WalletDistribution.Validate(); err != nil {
		return err
	}

	return nil
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
	return f.Content.Validate()
}

type FeeContractContent struct {
	FeeId            uint64         `json:"fee_id" yaml:"fee_id"`
	Creator          sdk.AccAddress `json:"creator" yaml:"creator"`
	Payer            sdk.AccAddress `json:"payer" yaml:"payer"`
	CumulativeCharge sdk.Coins      `json:"cumulative_charge" yaml:"cumulative_charge"`
	CurrentRemainder sdk.Coins      `json:"current_charge" yaml:"current_charge"`
	CanDeauthorise   bool           `json:"can_deauthorise" yaml:"can_deauthorise"`
	Authorised       bool           `json:"authorised" yaml:"authorised"`
}

func NewFeeContractContent(feeId uint64, creator, payer sdk.AccAddress,
	canDeauthorise, authorised bool) FeeContractContent {
	return FeeContractContent{
		FeeId:            feeId,
		Creator:          creator,
		Payer:            payer,
		CumulativeCharge: sdk.NewCoins(),
		CurrentRemainder: sdk.NewCoins(),
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

func (fc FeeContract) IsFirstCharge() bool {
	return fc.Content.CumulativeCharge.IsZero()
}

// CanCharge False if not authorised or max has been reached
func (fc FeeContract) CanCharge(fee Fee) bool {
	if fee.Id != fc.Content.FeeId {
		panic("fee ID mismatch in CanCharge")
	}
	return !fc.Content.Authorised || fc.Content.CumulativeCharge.IsAllGTE(fee.Content.ChargeMaximum)
}

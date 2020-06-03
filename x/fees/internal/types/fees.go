package types

import sdk "github.com/cosmos/cosmos-sdk/types"

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
	Discounts          Discounts    `json:"discounts" yaml:"discounts"`
	WalletDistribution Distribution `json:"wallet_distribution" yaml:"wallet_distribution"`
}

func NewFee(id string, chargeAmount, chargeMinimum, chargeMaximum sdk.Coins,
	discounts Discounts, walletDistribution Distribution) Fee {
	return Fee{
		Id:                 id,
		ChargeAmount:       chargeAmount,
		ChargeMinimum:      chargeMinimum,
		ChargeMaximum:      chargeMaximum,
		Discounts:          discounts,
		WalletDistribution: walletDistribution,
	}
}

func (f Fee) GetDiscountPercent(discountId sdk.Uint) (sdk.Dec, sdk.Error) {
	for _, discount := range f.Discounts {
		if discount.Id.Equal(discountId) {
			return discount.Percent, nil
		}
	}
	return sdk.Dec{}, ErrDiscountIdIsNotInFee(DefaultCodespace)
}

func (f Fee) Validate() sdk.Error {
	// Validate ID
	if !IsValidFeeId(f.Id) {
		return ErrInvalidId(DefaultCodespace, "fee id invalid")
	}

	// Validate charge amount, minimum, maximum
	amt := &f.ChargeAmount
	min := &f.ChargeMinimum
	max := &f.ChargeMaximum
	if !amt.IsValid() {
		return ErrInvalidFee(DefaultCodespace, "ChargeAmount coins invalid")
	} else if !min.IsValid() {
		return ErrInvalidFee(DefaultCodespace, "ChargeMinimum coins invalid")
	} else if !max.IsValid() {
		return ErrInvalidFee(DefaultCodespace, "ChargeMaximum coins invalid")
	} else if min.IsAnyGT(*max) {
		return ErrInvalidFee(DefaultCodespace, "min charge includes value greater than max")
	} else if !min.DenomsSubsetOf(*amt) {
		return ErrInvalidFee(DefaultCodespace, "min charge includes denom not in fee amount")
	} else if !max.DenomsSubsetOf(*amt) {
		return ErrInvalidFee(DefaultCodespace, "max charge includes denom not in fee amount")
	}

	// Validate discounts
	if err := f.Discounts.Validate(); err != nil {
		return err
	}

	// Validate wallet distribution
	if err := f.WalletDistribution.Validate(); err != nil {
		return err
	}

	return nil
}

type FeeContract struct {
	Id               string         `json:"id" yaml:"id"`
	FeeId            string         `json:"fee_id" yaml:"fee_id"`
	Creator          sdk.AccAddress `json:"creator" yaml:"creator"`
	Payer            sdk.AccAddress `json:"payer" yaml:"payer"`
	CumulativeCharge sdk.Coins      `json:"cumulative_charge" yaml:"cumulative_charge"`
	CurrentRemainder sdk.Coins      `json:"current_charge" yaml:"current_charge"`
	CanDeauthorise   bool           `json:"can_deauthorise" yaml:"can_deauthorise"`
	Authorised       bool           `json:"authorised" yaml:"authorised"`
	DiscountId       sdk.Uint       `json:"discount_id" yaml:"discount_id"`
}

func NewFeeContract(id, feeId string, creator, payer sdk.AccAddress,
	canDeauthorise, authorised bool, discountId sdk.Uint) FeeContract {
	return FeeContract{
		Id:               id,
		FeeId:            feeId,
		Creator:          creator,
		Payer:            payer,
		CumulativeCharge: sdk.NewCoins(),
		CurrentRemainder: sdk.NewCoins(),
		CanDeauthorise:   canDeauthorise,
		Authorised:       authorised,
		DiscountId:       discountId,
	}
}

func NewFeeContractNoDiscount(id, feeId string, creator, payer sdk.AccAddress,
	canDeauthorise, authorised bool) FeeContract {
	return NewFeeContract(
		id, feeId, creator, payer, canDeauthorise, authorised, sdk.ZeroUint())
}

func (fc FeeContract) Validate() sdk.Error {
	// Validate ID
	if !IsValidFeeContractId(fc.Id) {
		return ErrInvalidId(DefaultCodespace, "fee contract id invalid")
	}

	// Validate coins
	if !fc.CumulativeCharge.IsValid() {
		return ErrInvalidFee(DefaultCodespace, "CumulativeCharge coins invalid")
	} else if !fc.CurrentRemainder.IsValid() {
		return ErrInvalidFee(DefaultCodespace, "CurrentRemainder coins invalid")
	}

	// Validate addresses
	if fc.Creator.Empty() {
		return sdk.ErrInvalidAddress("empty creator address")
	} else if fc.Payer.Empty() {
		return sdk.ErrInvalidAddress("empty payer address")
	}

	// Validate IDs
	if !IsValidFeeId(fc.FeeId) {
		return ErrInvalidId(DefaultCodespace, "fee id invalid")
	}

	return nil
}

func (fc FeeContract) IsFirstCharge() bool {
	return fc.CumulativeCharge.IsZero()
}

// CanCharge False if not authorised or the (non-zero!) max has been reached
func (fc FeeContract) CanCharge(fee Fee) bool {
	if fee.Id != fc.FeeId {
		panic("fee ID mismatch in CanCharge")
	}
	max := fee.ChargeMaximum
	return fc.Authorised && (max.IsZero() || max.IsAllGT(fc.CumulativeCharge))
}

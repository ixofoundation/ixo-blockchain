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
	if err := fc.Discounts.Validate(); err != nil {
		return err
	}

	// Validate wallet distribution
	if err := fc.WalletDistribution.Validate(); err != nil {
		return err
	}

	return nil
}

func (fc FeeContent) GetDiscountPercent(discountId uint64) (sdk.Dec, sdk.Error) {
	for _, discount := range fc.Discounts {
		if discount.Id == discountId {
			return discount.Percent, nil
		}
	}
	return sdk.Dec{}, ErrDiscountIdIsNotInFee(DefaultCodespace)
}

type Fee struct {
	Id      string     `json:"id" yaml:"id"`
	Content FeeContent `json:"content" yaml:"content"`
}

func NewFee(id string, content FeeContent) Fee {
	return Fee{
		Id:      id,
		Content: content,
	}
}

func (f Fee) Validate() sdk.Error {
	if !IsValidFeeId(f.Id) {
		return ErrInvalidId(DefaultCodespace, "fee id invalid")
	}
	return f.Content.Validate()
}

type FeeContractContent struct {
	FeeId            string         `json:"fee_id" yaml:"fee_id"`
	Creator          sdk.AccAddress `json:"creator" yaml:"creator"`
	Payer            sdk.AccAddress `json:"payer" yaml:"payer"`
	CumulativeCharge sdk.Coins      `json:"cumulative_charge" yaml:"cumulative_charge"`
	CurrentRemainder sdk.Coins      `json:"current_charge" yaml:"current_charge"`
	CanDeauthorise   bool           `json:"can_deauthorise" yaml:"can_deauthorise"`
	Authorised       bool           `json:"authorised" yaml:"authorised"`
	DiscountIds      []uint64       `json:"discount_ids" yaml:"discount_ids"`
}

func NewFeeContractContent(feeId string, creator, payer sdk.AccAddress,
	canDeauthorise, authorised bool, discountIds []uint64) FeeContractContent {
	return FeeContractContent{
		FeeId:            feeId,
		Creator:          creator,
		Payer:            payer,
		CumulativeCharge: sdk.NewCoins(),
		CurrentRemainder: sdk.NewCoins(),
		CanDeauthorise:   canDeauthorise,
		Authorised:       authorised,
		DiscountIds:      discountIds,
	}
}

func (fc FeeContractContent) Validate() sdk.Error {
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

type FeeContract struct {
	Id      string             `json:"id" yaml:"id"`
	Content FeeContractContent `json:"content" yaml:"content"`
}

func NewFeeContract(id string, content FeeContractContent) FeeContract {
	return FeeContract{
		Id:      id,
		Content: content,
	}
}

func (fc FeeContract) Validate() sdk.Error {
	if !IsValidFeeContractId(fc.Id) {
		return ErrInvalidId(DefaultCodespace, "fee contract id invalid")
	}
	return fc.Content.Validate()
}

func (fc FeeContract) IsFirstCharge() bool {
	return fc.Content.CumulativeCharge.IsZero()
}

// CanCharge False if not authorised or the (non-zero!) max has been reached
func (fc FeeContract) CanCharge(fee Fee) bool {
	if fee.Id != fc.Content.FeeId {
		panic("fee ID mismatch in CanCharge")
	}
	max := fee.Content.ChargeMaximum
	return fc.Content.Authorised && (max.IsZero() || max.IsAllGT(fc.Content.CumulativeCharge))
}

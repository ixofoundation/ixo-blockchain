package types

import sdk "github.com/cosmos/cosmos-sdk/types"

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

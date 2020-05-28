package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// --------------------------------------------- Discounts

type Discounts []Discount

func NewDiscounts(discounts ...Discount) Discounts {
	return Discounts(discounts)
}

func (ds Discounts) Validate() sdk.Error {
	if len(ds) == 0 {
		return nil
	}

	// Check that discount IDs are sequential, starting with 1
	id := uint64(1)
	for _, d := range ds {
		if d.Id != id {
			return ErrDiscountIDsBeSequentialFrom1(DefaultCodespace)
		}
		id += 1
	}

	// Validate list of discounts
	for _, d := range ds {
		if err := d.Validate(); err != nil {
			return err
		}
	}

	return nil
}

type Discount struct {
	Id      uint64  `json:"id" yaml:"id"`
	Percent sdk.Dec `json:"percent" yaml:"percent"`
}

func NewDiscount(id uint64, percent sdk.Dec) Discount {
	return Discount{
		Id:      id,
		Percent: percent,
	}
}

func (d Discount) Validate() sdk.Error {
	if !d.Percent.IsPositive() {
		return ErrNegativeDiscountPercentage(DefaultCodespace)
	} else if d.Percent.GT(sdk.NewDec(100)) {
		return ErrDiscountPercentageGreaterThan100(DefaultCodespace)
	}

	return nil
}

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type DistributionShare struct {
	address    string
	percentage sdk.Dec
}

var OneHundred = sdk.NewDec(100)

type Distribution []DistributionShare

func NewDistribution(shares ...DistributionShare) Distribution {
	return Distribution(shares)
}

// func (d Distribution) Validate() error {
// 	// Shares must add up to 100% (no shares means 0%)
// 	if len(d) == 0 {
// 		return ErrDistributionPercentagesNot100
// 	}

// 	// Validate shares and calculate total
// 	total := sdk.ZeroDec()
// 	for _, share := range d {
// 		total = total.Add(share.Percentage)
// 		if err := share.Validate(); err != nil {
// 			return err
// 		}
// 	}

// 	// Shares must add up to 100%
// 	if !total.Equal(sdk.NewDec(100)) {
// 		return ErrDistributionPercentagesNot100
// 	}

// 	return nil
// }

func (d Distribution) GetDistributionsFor(amount sdk.Coins) []sdk.DecCoins {
	decAmount := sdk.NewDecCoinsFromCoins(amount...)
	distributions := make([]sdk.DecCoins, len(d))

	// Calculate distribution amount for each share of the distribution
	var distributed sdk.DecCoins
	for i, share := range d {
		distributions[i] = share.GetShareOf(decAmount)
		distributed = distributed.Add(distributions[i]...)
	}

	// Distributed amount should equal original amount
	if !distributed.IsEqual(decAmount) {
		panic("distributing more or less than original amount")
	}

	return distributions
}

func NewDistributionShare(address sdk.AccAddress, percentage sdk.Dec) DistributionShare {
	return DistributionShare{
		address:    address.String(),
		percentage: percentage,
	}
}

func NewFullDistributionShare(address sdk.AccAddress) DistributionShare {
	return DistributionShare{
		address:    address.String(),
		percentage: sdk.NewDec(100),
	}
}

// func (d DistributionShare) Validate() error {
// 	if !d.Percentage.IsPositive() {
// 		return sdkerrors.Wrap(ErrNegativeSharePercentage, "")
// 	} else if strings.TrimSpace(d.Address) == "" {
// 		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "empty distribution share address")
// 	}

// 	return nil
// }

func (d DistributionShare) GetShareOf(amount sdk.DecCoins) sdk.DecCoins {
	return amount.MulDec(d.percentage.Quo(OneHundred))
}

func (d DistributionShare) GetAddress() (sdk.AccAddress, error) {
	accAddress, err := sdk.AccAddressFromBech32(d.address)
	if err != nil {
		return nil, err
	}
	return accAddress, nil
}

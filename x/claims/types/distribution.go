package types

import (
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type DistributionShare struct {
	address    string
	percentage math.LegacyDec
}

var OneHundred = math.LegacyNewDec(100)

type Distribution []DistributionShare

func NewDistribution(shares ...DistributionShare) Distribution {
	return Distribution(shares)
}

func (d Distribution) GetDistributionsFor(amount sdk.Coins) ([]sdk.DecCoins, error) {
	decAmount := sdk.NewDecCoinsFromCoins(amount...)
	distributions := make([]sdk.DecCoins, len(d))

	// Calculate distribution amount for each share of the distribution
	var distributed sdk.DecCoins
	for i, share := range d {
		distributions[i] = share.GetShareOf(decAmount)
		distributed = distributed.Add(distributions[i]...)
	}

	// Distributed amount should equal original amount
	if !distributed.Equal(decAmount) {
		return nil, errorsmod.Wrap(ErrDistributionFailed, "distributing more or less than original amount")
	}

	return distributions, nil
}

func NewDistributionShare(address sdk.AccAddress, percentage math.LegacyDec) DistributionShare {
	return DistributionShare{
		address:    address.String(),
		percentage: percentage,
	}
}

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

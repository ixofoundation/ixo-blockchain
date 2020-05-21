package types

import sdk "github.com/cosmos/cosmos-sdk/types"

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

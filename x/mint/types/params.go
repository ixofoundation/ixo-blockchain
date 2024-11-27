package types

import (
	"errors"
	"fmt"
	"strings"

	yaml "gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/ixofoundation/ixo-blockchain/v3/ixomath"
	"github.com/ixofoundation/ixo-blockchain/v3/lib/ixo"
	epochtypes "github.com/ixofoundation/ixo-blockchain/v3/x/epochs/types"
)

var _ paramtypes.ParamSet = &Params{}

// ParamTable for minting module.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams returns new mint module parameters initialized to the given values.
func NewParams(
	mintDenom string, genesisEpochProvisions ixomath.Dec, epochIdentifier string,
	ReductionFactor ixomath.Dec, reductionPeriodInEpochs int64, distrProportions DistributionProportions,
	weightedImpactRewardsReceivers []WeightedAddress, mintingRewardsDistributionStartEpoch int64,
) Params {
	return Params{
		MintDenom:                            mintDenom,
		GenesisEpochProvisions:               genesisEpochProvisions,
		EpochIdentifier:                      epochIdentifier,
		ReductionPeriodInEpochs:              reductionPeriodInEpochs,
		ReductionFactor:                      ReductionFactor,
		DistributionProportions:              distrProportions,
		WeightedImpactRewardsReceivers:       weightedImpactRewardsReceivers,
		MintingRewardsDistributionStartEpoch: mintingRewardsDistributionStartEpoch,
	}
}

// DefaultParams returns the default minting module parameters.
func DefaultParams() Params {
	return Params{
		MintDenom:              ixo.IxoNativeToken,
		GenesisEpochProvisions: ixomath.NewDec(43_500_000_000), // 43_500 IXO
		// TODO: comment out 2min and uncomment day, 2min is for local testing
		EpochIdentifier:         "2min",  // 2 min
		ReductionPeriodInEpochs: 1000000, // 14 min
		// EpochIdentifier:         "day",                                          // 1 day
		// ReductionPeriodInEpochs: 365,                                            // 1 years
		ReductionFactor: ixomath.NewDecWithPrec(666666666666666666, 18), // 0.666666666666666666
		DistributionProportions: DistributionProportions{
			Staking:       ixomath.NewDecWithPrec(1, 0), // 1.0
			ImpactRewards: ixomath.NewDecWithPrec(0, 1), // 0.0
			CommunityPool: ixomath.NewDecWithPrec(0, 1), // 0.0
		},
		WeightedImpactRewardsReceivers:       []WeightedAddress{},
		MintingRewardsDistributionStartEpoch: 1,
	}
}

// Validate validates mint module parameters. Returns nil if valid,
// error otherwise
func (p Params) Validate() error {
	if err := validateMintDenom(p.MintDenom); err != nil {
		return err
	}
	if err := validateGenesisEpochProvisions(p.GenesisEpochProvisions); err != nil {
		return err
	}
	if err := epochtypes.ValidateEpochIdentifierInterface(p.EpochIdentifier); err != nil {
		return err
	}
	if err := validateReductionPeriodInEpochs(p.ReductionPeriodInEpochs); err != nil {
		return err
	}
	if err := validateReductionFactor(p.ReductionFactor); err != nil {
		return err
	}
	if err := validateDistributionProportions(p.DistributionProportions); err != nil {
		return err
	}
	if err := validateWeightedImpactRewardsReceivers(p.WeightedImpactRewardsReceivers); err != nil {
		return err
	}
	if err := validateMintingRewardsDistributionStartEpoch(p.MintingRewardsDistributionStartEpoch); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// Implements params.ParamSet.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMintDenom, &p.MintDenom, validateMintDenom),
		paramtypes.NewParamSetPair(KeyGenesisEpochProvisions, &p.GenesisEpochProvisions, validateGenesisEpochProvisions),
		paramtypes.NewParamSetPair(KeyEpochIdentifier, &p.EpochIdentifier, epochtypes.ValidateEpochIdentifierInterface),
		paramtypes.NewParamSetPair(KeyReductionPeriodInEpochs, &p.ReductionPeriodInEpochs, validateReductionPeriodInEpochs),
		paramtypes.NewParamSetPair(KeyReductionFactor, &p.ReductionFactor, validateReductionFactor),
		paramtypes.NewParamSetPair(KeyPoolAllocationRatio, &p.DistributionProportions, validateDistributionProportions),
		paramtypes.NewParamSetPair(KeyImpactRewardsReceivers, &p.WeightedImpactRewardsReceivers, validateWeightedImpactRewardsReceivers),
		paramtypes.NewParamSetPair(KeyMintingRewardsDistributionStartEpoch, &p.MintingRewardsDistributionStartEpoch, validateMintingRewardsDistributionStartEpoch),
	}
}

func validateMintDenom(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if strings.TrimSpace(v) == "" {
		return errors.New("mint denom cannot be blank")
	}
	if err := sdk.ValidateDenom(v); err != nil {
		return err
	}

	return nil
}

func validateGenesisEpochProvisions(i interface{}) error {
	v, ok := i.(ixomath.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("genesis epoch provision must be non-negative")
	}

	return nil
}

func validateReductionPeriodInEpochs(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("max validators must be positive: %d", v)
	}

	return nil
}

func validateReductionFactor(i interface{}) error {
	v, ok := i.(ixomath.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.GT(ixomath.NewDec(1)) {
		return fmt.Errorf("reduction factor cannot be greater than 1")
	}

	if v.IsNegative() {
		return fmt.Errorf("reduction factor cannot be negative")
	}

	return nil
}

func validateDistributionProportions(i interface{}) error {
	v, ok := i.(DistributionProportions)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.Staking.IsNegative() {
		return errors.New("staking distribution ratio should not be negative")
	}

	if v.ImpactRewards.IsNegative() {
		return errors.New("impact rewards distribution ratio should not be negative")
	}

	if v.CommunityPool.IsNegative() {
		return errors.New("community pool distribution ratio should not be negative")
	}

	totalProportions := v.Staking.Add(v.ImpactRewards).Add(v.CommunityPool)

	if !totalProportions.Equal(ixomath.NewDec(1)) {
		return errors.New("total distributions ratio should be 1")
	}

	return nil
}

func validateWeightedImpactRewardsReceivers(i interface{}) error {
	v, ok := i.([]WeightedAddress)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	// fund community pool when rewards address is empty
	if len(v) == 0 {
		return nil
	}

	weightSum := ixomath.NewDec(0)
	for i, w := range v {
		// we allow address to be "" to go to community pool
		if w.Address != "" {
			_, err := sdk.AccAddressFromBech32(w.Address)
			if err != nil {
				return fmt.Errorf("invalid address at %dth", i)
			}
		}
		if !w.Weight.IsPositive() {
			return fmt.Errorf("non-positive weight at %dth", i)
		}
		if w.Weight.GT(ixomath.NewDec(1)) {
			return fmt.Errorf("more than 1 weight at %dth", i)
		}
		weightSum = weightSum.Add(w.Weight)
	}

	if !weightSum.Equal(ixomath.NewDec(1)) {
		return fmt.Errorf("invalid weight sum: %s", weightSum.String())
	}

	return nil
}

func validateMintingRewardsDistributionStartEpoch(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v < 0 {
		return fmt.Errorf("start epoch must be non-negative")
	}

	return nil
}

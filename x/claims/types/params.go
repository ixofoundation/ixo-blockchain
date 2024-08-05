package types

import (
	fmt "fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	KeyCollectionSequence   = []byte("CollectionSequence")
	KeyIxoAccount           = []byte("IxoAccount")
	KeyNetworkFeePercentage = []byte("NetworkFeePercentage")
	KeyNodeFeePercentage    = []byte("NodeFeePercentage")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(collectionSequence uint64, ixoAccount string,
	networkFeePercentage, nodeFeePercentage math.LegacyDec) Params {
	return Params{
		CollectionSequence:   collectionSequence,
		IxoAccount:           ixoAccount,
		NetworkFeePercentage: networkFeePercentage,
		NodeFeePercentage:    nodeFeePercentage,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	defaultIxoAccount := "ixo1kqmtxkggcqa9u34lnr6shy0euvclgatw4f9zz5"
	tenPercentFee := math.LegacyNewDec(10)

	return NewParams(1, defaultIxoAccount, tenPercentFee, tenPercentFee)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.ParamSetPair{Key: KeyCollectionSequence, Value: &p.CollectionSequence, ValidatorFn: validateCollectionSequence},
		paramtypes.ParamSetPair{Key: KeyIxoAccount, Value: &p.IxoAccount, ValidatorFn: validateIxoAccount},
		paramtypes.ParamSetPair{Key: KeyNetworkFeePercentage, Value: &p.NetworkFeePercentage, ValidatorFn: validateFeePercentage},
		paramtypes.ParamSetPair{Key: KeyNodeFeePercentage, Value: &p.NodeFeePercentage, ValidatorFn: validateFeePercentage},
	}
}

func validateIxoAccount(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	_, err := sdk.AccAddressFromBech32(v)
	if err != nil {
		return err
	}

	return nil
}

func validateFeePercentage(i interface{}) error {
	v, ok := i.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.LT(math.LegacyZeroDec()) {
		return fmt.Errorf("invalid parameter fee percentage; should be >= 0.0, is %s ", v.String())
	} else if v.GT(math.LegacyNewDec(100)) {
		return fmt.Errorf("invalid parameter fee percentage; should be <= 100, is %s ", v.String())
	}

	return nil
}

func validateCollectionSequence(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T expected uint64", i)
	}

	return nil
}

// Validate validates the set of params
func (p Params) Validate() error {
	err := validateCollectionSequence(p.CollectionSequence)
	if err != nil {
		return err
	}
	err = validateIxoAccount(p.IxoAccount)
	if err != nil {
		return err
	}
	err = validateFeePercentage(p.NetworkFeePercentage)
	if err != nil {
		return err
	}
	err = validateFeePercentage(p.NodeFeePercentage)
	if err != nil {
		return err
	}
	return nil
}

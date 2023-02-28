package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	didexported "github.com/ixofoundation/ixo-blockchain/lib/legacydid"
)

var (
	KeyCollectionSequence  = []byte("CollectionSequence")
	KeyIxoDid              = []byte("IxoDid")
	KeyOracleFeePercentage = []byte("OracleFeePercentage")
	KeyNodeFeePercentage   = []byte("NodeFeePercentage")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(collectionSequence uint64, ixoDid didexported.Did,
	oracleFeePercentage, nodeFeePercentage sdk.Dec) Params {
	return Params{
		CollectionSequence:  collectionSequence,
		IxoDid:              ixoDid,
		OracleFeePercentage: oracleFeePercentage,
		NodeFeePercentage:   nodeFeePercentage,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	defaultIxoDid := didexported.Did("did:ixo:U4tSpzzv91HHqWW1YmFkHJ")
	tenPercentFee := sdk.NewDec(10)

	return NewParams(0, defaultIxoDid, tenPercentFee, tenPercentFee)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.ParamSetPair{Key: KeyCollectionSequence, Value: &p.CollectionSequence, ValidatorFn: validateCollectionSequence},
		paramtypes.ParamSetPair{Key: KeyIxoDid, Value: &p.IxoDid, ValidatorFn: validateIxoDid},
		paramtypes.ParamSetPair{Key: KeyOracleFeePercentage, Value: &p.OracleFeePercentage, ValidatorFn: validateFeePercentage},
		paramtypes.ParamSetPair{Key: KeyNodeFeePercentage, Value: &p.NodeFeePercentage, ValidatorFn: validateFeePercentage},
	}
}

func validateIxoDid(i interface{}) error {
	v, ok := i.(didexported.Did)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if len(v) == 0 {
		return fmt.Errorf("ixo did cannot be empty")
	}

	return nil
}

func validateFeePercentage(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.LT(sdk.ZeroDec()) {
		return fmt.Errorf("invalid parameter fee percentage; should be >= 0.0, is %s ", v.String())
	} else if v.GT(sdk.NewDec(100)) {
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
	err = validateIxoDid(p.IxoDid)
	if err != nil {
		return err
	}
	err = validateFeePercentage(p.OracleFeePercentage)
	if err != nil {
		return err
	}
	err = validateFeePercentage(p.NodeFeePercentage)
	if err != nil {
		return err
	}
	return nil
}

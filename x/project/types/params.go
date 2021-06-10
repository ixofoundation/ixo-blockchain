package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
	didexported "github.com/ixofoundation/ixo-blockchain/x/did/exported"
)

const (
	IxoNativeToken = "uixo"
)


// Parameter store keys
var (
	KeyIxoDid                       = []byte("IxoDid")
	KeyProjectMinimumInitialFunding = []byte("ProjectMinimumInitialFunding")
	KeyOracleFeePercentage          = []byte("OracleFeePercentage")
	KeyNodeFeePercentage            = []byte("NodeFeePercentage")
)

// ParamTable for project module.
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(projectMinimumInitialFunding sdk.Coins, ixoDid didexported.Did,
	oracleFeePercentage, nodeFeePercentage sdk.Dec) Params {
	return Params{
		IxoDid:                       ixoDid,
		ProjectMinimumInitialFunding: projectMinimumInitialFunding,
		OracleFeePercentage:          oracleFeePercentage,
		NodeFeePercentage:            nodeFeePercentage,
	}

}

// default project module parameters
func DefaultParams() Params {
	defaultIxoDid := didexported.Did("did:ixo:U4tSpzzv91HHqWW1YmFkHJ")
	defaultMinInitFunding := sdk.NewCoins(sdk.NewCoin(
		IxoNativeToken, sdk.OneInt()))
	tenPercentFee := sdk.NewDec(10)

	return Params{
		IxoDid:                       defaultIxoDid,         // invalid blank
		ProjectMinimumInitialFunding: defaultMinInitFunding, // 1uixo
		OracleFeePercentage:          tenPercentFee,         // 10.0 (10%)
		NodeFeePercentage:            tenPercentFee,         // 10.0 (10%)
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

func validateProjectMinimumInitialFunding(i interface{}) error {
	v, ok := i.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsAnyNegative() {
		return fmt.Errorf("invalid project minimum initial "+
			"funding should be positive, is %s ", v.String())
	}

	return nil
}

func validateOracleFeePercentage(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.LT(sdk.ZeroDec()) {
		return fmt.Errorf("invalid parameter oracle fee percentage; should be >= 0.0, is %s ", v.String())
	} else if v.GT(sdk.NewDec(100)) {
		return fmt.Errorf("invalid parameter oracle fee percentage; should be <= 100, is %s ", v.String())
	}

	return nil
}

func validateNodeFeePercentage(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.LT(sdk.ZeroDec()) {
		return fmt.Errorf("invalid parameter node fee percentage; should be >= 0.0, is %s ", v.String())
	} else if v.GT(sdk.NewDec(100)) {
		return fmt.Errorf("invalid parameter node fee percentage; should be <= 100, is %s ", v.String())
	}

	return nil
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{KeyIxoDid, &p.IxoDid, validateIxoDid},
		{KeyProjectMinimumInitialFunding, &p.ProjectMinimumInitialFunding, validateProjectMinimumInitialFunding},
		{KeyOracleFeePercentage, &p.OracleFeePercentage, validateOracleFeePercentage},
		{KeyNodeFeePercentage, &p.NodeFeePercentage, validateNodeFeePercentage},
	}
}

package types

import (
	"encoding/json"
	"fmt"

	"cosmossdk.io/math"
)

// DefaultMinLiquidStakeAmount is the smallest acceptable liquid-stake amount,
// applied across the entire module. Set conservatively at 1 IXO (1,000,000
// uixo) to avoid decimal-loss artefacts and gas waste from dust amounts.
var DefaultMinLiquidStakeAmount = math.NewInt(1_000_000)

// DefaultModuleParams returns the module's initial global parameters.
func DefaultModuleParams() ModuleParams {
	return ModuleParams{
		MinLiquidStakeAmount: DefaultMinLiquidStakeAmount,
		ModulePaused:         false,
	}
}

// String renders ModuleParams as indented JSON. Required because the proto
// definition disables gogoproto's auto-generated stringer.
func (p ModuleParams) String() string {
	out, _ := json.MarshalIndent(p, "", "  ")
	return string(out)
}

// Validate verifies every field of ModuleParams.
func (p ModuleParams) Validate() error {
	return validateMinLiquidStakeAmount(p.MinLiquidStakeAmount)
}

func validateMinLiquidStakeAmount(v math.Int) error {
	if v.IsNil() {
		return fmt.Errorf("min liquid stake amount must not be nil")
	}
	if v.IsNegative() {
		return fmt.Errorf("min liquid stake amount must not be negative: %s", v)
	}
	return nil
}

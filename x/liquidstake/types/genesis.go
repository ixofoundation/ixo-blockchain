package types

import (
	"fmt"

	"cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewGenesisState assembles a multi-pool GenesisState. Each entry in
// poolValidators must reference a pool_id present in pools.
func NewGenesisState(moduleParams ModuleParams, pools []Pool, poolValidators []PoolLiquidValidators) *GenesisState {
	return &GenesisState{
		ModuleParams:         moduleParams,
		Pools:                pools,
		PoolLiquidValidators: poolValidators,
	}
}

// DefaultGenesisState returns an empty multi-pool genesis: default global
// params, no pools, no liquid validators. The v7 upgrade migration is what
// populates the legacy "zero" pool from pre-upgrade state.
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(DefaultModuleParams(), []Pool{}, []PoolLiquidValidators{})
}

// ValidateGenesis runs full validation over module_params, every pool, every
// liquid-validator group, and the cross-references between them.
func ValidateGenesis(data GenesisState) error {
	if err := data.ModuleParams.Validate(); err != nil {
		return fmt.Errorf("invalid module_params: %w", err)
	}

	// Validate pools, accumulating uniqueness sets for pool_id and denom.
	poolIDs := make(map[string]struct{}, len(data.Pools))
	denoms := make(map[string]struct{}, len(data.Pools))
	for i := range data.Pools {
		p := data.Pools[i]
		if err := p.Validate(); err != nil {
			return fmt.Errorf("pools[%d] (%q): %w", i, p.PoolId, err)
		}
		if _, dup := poolIDs[p.PoolId]; dup {
			return errors.Wrapf(ErrDuplicatePoolID, "pools[%d]: %s", i, p.PoolId)
		}
		poolIDs[p.PoolId] = struct{}{}
		if _, dup := denoms[p.LiquidBondDenom]; dup {
			return errors.Wrapf(ErrDuplicateLiquidBondDenom, "pools[%d]: %s", i, p.LiquidBondDenom)
		}
		denoms[p.LiquidBondDenom] = struct{}{}
	}

	// Validate every PoolLiquidValidators group references a known pool, has
	// no duplicate validator entries, and every entry parses as a validator.
	for i, group := range data.PoolLiquidValidators {
		if _, ok := poolIDs[group.PoolId]; !ok {
			return errors.Wrapf(ErrPoolNotFound,
				"pool_liquid_validators[%d]: pool_id %q has no matching pool", i, group.PoolId)
		}
		seen := map[string]struct{}{}
		for j, lv := range group.LiquidValidators {
			if err := lv.Validate(); err != nil {
				return errors.Wrapf(sdkerrors.ErrInvalidAddress,
					"pool_liquid_validators[%d].liquid_validators[%d]: %v", i, j, err)
			}
			if _, dup := seen[lv.OperatorAddress]; dup {
				return fmt.Errorf("pool_liquid_validators[%d]: duplicate liquid validator %s",
					i, lv.OperatorAddress)
			}
			seen[lv.OperatorAddress] = struct{}{}
		}
	}
	return nil
}

package types

import (
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// ParamTable for project module.
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams() Params {
	return Params{}
}

// default project module parameters
func DefaultParams() Params {
	return Params{}
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		{KeyIxoDid, &p.IxoDid, validateIxoDid},
	}
}

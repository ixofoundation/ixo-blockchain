package pool

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ValidatorPool struct {
	Validators []sdk.Validator
}

func NewValidatorPool() ValidatorPool {
	validators := make([]sdk.Validator, 0)
	return ValidatorPool{
		validators,
	}
}

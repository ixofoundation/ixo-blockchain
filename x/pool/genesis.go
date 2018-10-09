package pool

import (
	"fmt"

	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenesisState - all staking state that must be provided at genesis
type GenesisState struct {
	Validators []sdk.Validator `json:"validators"`
}

func NewGenesisState(validators []sdk.Validator) GenesisState {
	return GenesisState{
		Validators: validators,
	}
}

// get raw genesis raw message for testing
func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

// InitGenesis sets the pool and parameters for the provided keeper. For each validator in data,
// it sets that validator in the keeper
// Returns final validator set after applying all declaration and delegations
func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) (res []abci.Validator, err error) {
	fmt.Println("***", data)
	res = make([]abci.Validator, len(data.Validators))
	for i, val := range data.Validators {
		fmt.Println("***", i, val)
		res[i] = sdk.ABCIValidator(val)
	}
	return
}

package types

import (
	"github.com/ixofoundation/ixo-cosmos/types"
)

type GenesisState struct {
	FiatAccount []types.FiatAccount `json:"fiatAccount"`
}

func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

func ValidateGenesis(data GenesisState) error {
	return nil
}

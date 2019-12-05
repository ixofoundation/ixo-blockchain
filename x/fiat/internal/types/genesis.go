package types

import (
	"github.com/ixofoundation/ixo-cosmos/types"
)

type GenesisState struct {
	FiatPegWallet []types.FiatPeg `json:"fiatPegWallet"`
}

func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

func ValidateGenesis(data GenesisState) error {
	return nil
}

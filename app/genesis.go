package app

import (
	"github.com/ixofoundation/ixo-cosmos/types"
	"github.com/ixofoundation/ixo-cosmos/x/fees"
	"github.com/ixofoundation/ixo-cosmos/x/contracts"
	"github.com/ixofoundation/ixo-cosmos/x/node"
)

//___________________________________________________________________________________
// State to Unmarshal
type GenesisState struct {
	Accounts []*types.GenesisAccount `json:"accounts"`
	FeeData  fees.GenesisState       `json:"fees"`
	NodeData []*node.GenesisState    `json:"nodes"`
	Config   contracts.GenesisState        `json:"config"`
}

package node

const KeyNodeID = "nodeId"

const KeyEthWallet = "ethWallet"

var AllNodes = []string{
	KeyNodeID,
	KeyEthWallet,
}

type ETHGenesisState struct {
	Did       string `json:"did"`
	EthWallet string `json:"ethWallet"`
}

type GenesisState struct {
	ETHGenesisStates []ETHGenesisState
}

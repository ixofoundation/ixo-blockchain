package node

// KeyNodeID The node id on the network
const KeyNodeID = "nodeId"

// KeyEthWallet The ethereum wallet address for the node
const KeyEthWallet = "ethWallet"

var AllNodes = []string{
	KeyNodeID,
	KeyEthWallet,
}

// GenesisState of the node
type GenesisState struct {
	Did       string `json:"did"`
	EthWallet string `json:"ethWallet"`
}

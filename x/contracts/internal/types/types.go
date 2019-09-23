package types

const (
	KeyFoundationWallet                       = "foundationWallet"
	KeyAuthContractAddress                    = "authContractAddress"
	KeyIxoTokenContractAddress                = "ixoTokenContractAddress"
	KeyProjectRegistryContractAddress         = "projectRegistryContractAddress"
	KeyProjectWalletAuthoriserContractAddress = "projectWalletAuthoriserAddress"
)

var AllContracts = []string{
	KeyFoundationWallet,
	KeyAuthContractAddress,
	KeyIxoTokenContractAddress,
	KeyProjectRegistryContractAddress,
	KeyProjectWalletAuthoriserContractAddress,
}

type GenesisState struct {
	FoundationWallet               string `json:"foundationWallet"`
	AuthContractAddress            string `json:"authContractAddress"`
	IxoTokenContractAddress        string `json:"ixoTokenContractAddress"`
	ProjectRegistryContractAddress string `json:"projectRegistryContractAddress"`
	ProjectWalletAuthoriserAddress string `json:"projectWalletAuthoriserAddress"`
}

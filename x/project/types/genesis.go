package types

//type GenesisState struct {
//	ProjectDocs      []ProjectDoc          `json:"project_docs" yaml:"project_docs"`
//	AccountMaps      []GenesisAccountMap   `json:"account_maps" yaml:"account_maps"`
//	WithdrawalsInfos [][]WithdrawalInfoDoc `json:"withdrawal_infos" yaml:"withdrawal_infos"`
//	Claims           [][]Claim             `json:"claims" yaml:"claims"`
//	Params           Params                `json:"params" yaml:"params"`
//}

func NewGenesisState(projectDocs []ProjectDoc, accountMaps []GenesisAccountMap,
	withdrawalInfos []WithdrawalInfoDocs, claims []Claims, params Params) *GenesisState {
	return &GenesisState{
		ProjectDocs:      projectDocs,
		AccountMaps:      accountMaps,
		WithdrawalsInfos: withdrawalInfos,
		Claims:           claims,
		Params:           params,
	}
}

func ValidateGenesis(data GenesisState) error {
	return nil
}

func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		ProjectDocs:      nil,
		AccountMaps:      nil,
		WithdrawalsInfos: nil,
		Claims:           nil,
		Params:           DefaultParams(),
	}
}

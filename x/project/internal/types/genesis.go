package types

type GenesisState struct {
	ProjectDocs      []ProjectDoc        `json:"project_docs" yaml:"project_docs"`
	AccountMaps      []GenesisAccountMap `json:"account_maps" yaml:"account_maps"`
	WithdrawalsInfos [][]WithdrawalInfo  `json:"withdrawal_infos" yaml:"withdrawal_infos"`
	Params           Params              `json:"params" yaml:"params"`
}

func NewGenesisState(projectDocs []ProjectDoc, accountMaps []GenesisAccountMap,
	withdrawalInfos [][]WithdrawalInfo, params Params) GenesisState {
	return GenesisState{
		ProjectDocs:      projectDocs,
		AccountMaps:      accountMaps,
		WithdrawalsInfos: withdrawalInfos,
		Params:           params,
	}
}

//noinspection GoUnusedParameter
func ValidateGenesis(data GenesisState) error {
	err := ValidateParams(data.Params)
	if err != nil {
		return err
	}

	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		ProjectDocs:      nil,
		AccountMaps:      nil,
		WithdrawalsInfos: nil,
		Params:           DefaultParams(),
	}
}

package types

type GenesisState struct {
	ProjectDocs      []MsgCreateProject `json:"project_docs" yaml:"project_docs"`
	AccountMaps      []AccountMap       `json:"account_maps" yaml:"account_maps"`
	WithdrawalsInfos [][]WithdrawalInfo `json:"withdrawal_infos" yaml:"withdrawal_infos"`
	Params           Params             `json:"params" yaml:"params"`
}

func NewGenesisState(projectDocs []MsgCreateProject, accountMaps []AccountMap,
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

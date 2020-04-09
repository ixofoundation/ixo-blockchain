package types

type GenesisState struct {
	ProjectDocs      []MsgCreateProject `json:"project_docs" yaml:"project_docs"`
	AccountMaps      []AccountMap       `json:"account_maps" yaml:"account_maps"`
	WithdrawalsInfos [][]WithdrawalInfo `json:"withdrawal_infos" yaml:"withdrawal_infos"`
}

func NewGenesisState(projectDocs []MsgCreateProject, accountMaps []AccountMap,
	withdrawalInfos [][]WithdrawalInfo) GenesisState {
	return GenesisState{
		ProjectDocs:      projectDocs,
		AccountMaps:      accountMaps,
		WithdrawalsInfos: withdrawalInfos,
	}
}

//noinspection GoUnusedParameter
func ValidateGenesis(data GenesisState) error {
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		ProjectDocs:      nil,
		AccountMaps:      nil,
		WithdrawalsInfos: nil,
	}
}

package types

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

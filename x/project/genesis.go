package project

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	// Marshal/Unmarshal account maps into array of AccountMap
	accountMapsBz, err := json.Marshal(data.AccountMaps)
	if err != nil {
		panic(err)
	}
	var accountMaps []AccountMap
	err = json.Unmarshal(accountMapsBz, &accountMaps)
	if err != nil {
		panic(err)
	}

	// Initialise project docs, account maps, project withdrawals, params
	for i := range data.ProjectDocs {
		keeper.SetProjectDoc(ctx, &data.ProjectDocs[i])
		keeper.SetAccountMap(ctx,
			data.ProjectDocs[i].GetProjectDid(), accountMaps[i])
		keeper.SetProjectWithdrawalTransactions(ctx,
			data.ProjectDocs[i].GetProjectDid(), data.WithdrawalsInfos[i])
	}
	keeper.SetParams(ctx, data.Params)
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	// Export project docs, account maps, project withdrawals
	var projectDocs []ProjectDoc
	var accountMaps []AccountMap
	var withdrawalInfos [][]WithdrawalInfo

	iterator := k.GetProjectDocIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		projectDoc := k.MustGetProjectDocByKey(ctx, iterator.Key())
		accountMap := k.GetAccountMap(ctx, projectDoc.GetProjectDid())
		withdrawalInfo, _ := k.GetProjectWithdrawalTransactions(ctx, projectDoc.GetProjectDid())

		projectDocs = append(projectDocs, *projectDoc.(*ProjectDoc))
		accountMaps = append(accountMaps, accountMap)
		withdrawalInfos = append(withdrawalInfos, withdrawalInfo)
	}

	params := k.GetParams(ctx)

	// Marshal/Unmarshal account maps into array of GenesisAccountMap
	accountMapsBz, err := json.Marshal(accountMaps)
	if err != nil {
		panic(err)
	}
	var genesisAccountMaps []GenesisAccountMap
	err = json.Unmarshal(accountMapsBz, &genesisAccountMaps)
	if err != nil {
		panic(err)
	}

	return GenesisState{
		ProjectDocs:      projectDocs,
		AccountMaps:      genesisAccountMaps,
		WithdrawalsInfos: withdrawalInfos,
		Params:           params,
	}
}

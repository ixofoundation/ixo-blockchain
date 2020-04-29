package project

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	// Initialise project docs, account maps, project withdrawals, params
	for i := range data.ProjectDocs {
		keeper.SetProjectDoc(ctx, &data.ProjectDocs[i])
		keeper.SetAccountMap(ctx,
			data.ProjectDocs[i].GetProjectDid(), data.AccountMaps[i])
		keeper.SetProjectWithdrawalTransactions(ctx,
			data.ProjectDocs[i].GetProjectDid(), data.WithdrawalsInfos[i])
	}
	keeper.SetParams(ctx, data.Params)
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	// Export project docs, account maps, project withdrawals
	var projectDocs []MsgCreateProject
	var accountMaps []AccountMap
	var withdrawalInfos [][]WithdrawalInfo

	iterator := k.GetProjectDocIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		projectDoc := k.MustGetProjectDocByKey(ctx, iterator.Key())
		accountMap := k.GetAccountMap(ctx, projectDoc.GetProjectDid())
		withdrawalInfo, _ := k.GetProjectWithdrawalTransactions(ctx, projectDoc.GetProjectDid())

		projectDocs = append(projectDocs, *projectDoc.(*MsgCreateProject))
		accountMaps = append(accountMaps, accountMap)
		withdrawalInfos = append(withdrawalInfos, withdrawalInfo)
	}

	params := k.GetParams(ctx)

	return GenesisState{
		ProjectDocs:      projectDocs,
		AccountMaps:      accountMaps,
		WithdrawalsInfos: withdrawalInfos,
		Params:           params,
	}
}

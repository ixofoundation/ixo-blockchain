package project

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/project/internal/types"
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

	// Initialise project docs, account maps, project withdrawals, claims
	for i := range data.ProjectDocs {
		keeper.SetProjectDoc(ctx, &data.ProjectDocs[i])
		keeper.SetAccountMap(ctx,
			data.ProjectDocs[i].GetProjectDid(), accountMaps[i])
		keeper.SetProjectWithdrawalTransactions(ctx,
			data.ProjectDocs[i].GetProjectDid(), data.WithdrawalsInfos[i])
		for j := range data.Claims {
			keeper.SetClaim(ctx,
				data.ProjectDocs[i].GetProjectDid(), data.Claims[i][j])
		}
	}

	// Initialise params
	keeper.SetParams(ctx, data.Params)
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	// Export project docs, account maps, project withdrawals
	var projectDocs []ProjectDoc
	var accountMaps []AccountMap
	var withdrawalInfos [][]WithdrawalInfo
	var claims [][]types.Claim

	iterator := k.GetProjectDocIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		projectDoc := k.MustGetProjectDocByKey(ctx, iterator.Key())
		accountMap := k.GetAccountMap(ctx, projectDoc.GetProjectDid())
		withdrawalInfo, _ := k.GetProjectWithdrawalTransactions(ctx, projectDoc.GetProjectDid())

		var subClaims []types.Claim
		claimIter := k.GetClaimIterator(ctx, projectDoc.GetProjectDid())
		for ; claimIter.Valid(); claimIter.Next() {
			claim := k.MustGetClaimByKey(ctx, claimIter.Key())
			subClaims = append(subClaims, claim)
		}

		projectDocs = append(projectDocs, *projectDoc.(*ProjectDoc))
		accountMaps = append(accountMaps, accountMap)
		withdrawalInfos = append(withdrawalInfos, withdrawalInfo)
		claims = append(claims, subClaims)
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
		Claims:           claims,
		Params:           params,
	}
}

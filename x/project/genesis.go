package project

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/project/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/project/types"
)

func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data types.GenesisState) {
	// Marshal/Unmarshal account maps into array of AccountMap
	accountMapsBz, err := json.Marshal(data.AccountMaps)
	if err != nil {
		panic(err)
	}
	var accountMaps []types.AccountMap
	err = json.Unmarshal(accountMapsBz, &accountMaps)
	if err != nil {
		panic(err)
	}

	// Initialise project docs, account maps, project withdrawals, claims
	for i := range data.ProjectDocs {
		keeper.SetProjectDoc(ctx, data.ProjectDocs[i])
		keeper.SetAccountMap(ctx,
			data.ProjectDocs[i].ProjectDid, accountMaps[i])
		keeper.SetProjectWithdrawalTransactions(ctx,
			data.ProjectDocs[i].ProjectDid, data.WithdrawalsInfos[i])
		for j := range data.Claims[i].ClaimsList {
			keeper.SetClaim(ctx,
				data.ProjectDocs[i].ProjectDid, data.Claims[i].ClaimsList[j])
		}
	}

	// Initialise params
	keeper.SetParams(ctx, data.Params)
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	// Export project docs, account maps, project withdrawals
	var projectDocs []types.ProjectDoc
	var accountMaps []types.AccountMap
	var withdrawalInfos []types.WithdrawalInfoDocs
	var claims []types.Claims

	iterator := k.GetProjectDocIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		projectDoc := k.MustGetProjectDocByKey(ctx, iterator.Key())
		accountMap := k.GetAccountMap(ctx, projectDoc.ProjectDid)
		withdrawalInfo, _ := k.GetProjectWithdrawalTransactions(ctx, projectDoc.ProjectDid)

		var subClaims types.Claims
		claimIter := k.GetClaimIterator(ctx, projectDoc.ProjectDid)
		for ; claimIter.Valid(); claimIter.Next() {
			claim := k.MustGetClaimByKey(ctx, claimIter.Key())
			subClaims = types.AppendClaims(subClaims, claim) //append(subClaims, claim)
		}

		projectDocs = append(projectDocs, projectDoc)
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
	var genesisAccountMaps []types.GenesisAccountMap
	err = json.Unmarshal(accountMapsBz, &genesisAccountMaps)
	if err != nil {
		panic(err)
	}

	return types.GenesisState{
		ProjectDocs:      projectDocs,
		AccountMaps:      genesisAccountMaps,
		WithdrawalsInfos: withdrawalInfos,
		Claims:           claims,
		Params:           params,
	}
}

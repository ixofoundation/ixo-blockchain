package keeper_test

import (
	"encoding/json"
	"github.com/ixofoundation/ixo-blockchain/x/project/keeper"
	"testing"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/ixofoundation/ixo-blockchain/x/project/types"
)

func TestQueryProjectDoc(t *testing.T) {
	legacyAmino, appl, ctx := CreateTestInput()

	require.False(t, appl.ProjectKeeper.ProjectDocExists(ctx, types.ProjectDid))
	appl.ProjectKeeper.SetProjectDoc(ctx, types.ValidProjectDoc)

	query := abci.RequestQuery{
		Path: "",
		Data: []byte{},
	}

	querier := keeper.NewQuerier(appl.ProjectKeeper, legacyAmino)
	res, err := querier(ctx, []string{"queryProjectDoc", types.ProjectDid}, query)
	require.Nil(t, err)

	emptyRes, err := querier(ctx, []string{"queryProjectDoc", "InvalidProjectDid"}, query)
	require.Nil(t, emptyRes)
	require.NotNil(t, err)

	var projectDoc types.ProjectDoc
	legacyAmino.MustUnmarshalJSON(res, &projectDoc)
}

func TestQueryProjectAccounts(t *testing.T) {
	legacyAmino, appl, ctx := CreateTestInput()

	require.False(t, appl.ProjectKeeper.ProjectDocExists(ctx, types.ProjectDid))
	appl.ProjectKeeper.SetProjectDoc(ctx, types.ValidProjectDoc)

	query := abci.RequestQuery{
		Path: "",
		Data: []byte{},
	}

	querier := keeper.NewQuerier(appl.ProjectKeeper, legacyAmino)
	_, err := querier(ctx, []string{keeper.QueryProjectDoc, types.ProjectDid}, query)
	require.Nil(t, err)

	account, err := appl.ProjectKeeper.CreateNewAccount(ctx, types.ProjectDid, types.ValidAccId1)
	require.Nil(t, err)
	appl.ProjectKeeper.AddAccountToProjectAccounts(ctx, types.ProjectDid, types.ValidAccId1, account)

	res, err := querier(ctx, []string{keeper.QueryProjectAccounts, types.ProjectDid}, query)
	require.Nil(t, err)

	var data interface{}
	require.Nil(t, json.Unmarshal(res, &data))

	accountMap := data.(map[string]interface{})
	_, errRes := json.Marshal(accountMap)
	require.Nil(t, errRes)

	account, err = appl.ProjectKeeper.CreateNewAccount(ctx, types.ProjectDid, types.ValidAccId1)
	require.NotNil(t, err)
}

func TestQueryTxs(t *testing.T) {
	legacyAmino, appl, ctx := CreateTestInput()

	require.False(t, appl.ProjectKeeper.ProjectDocExists(ctx, types.ProjectDid))
	appl.ProjectKeeper.SetProjectDoc(ctx, types.ValidProjectDoc)

	appl.ProjectKeeper.AddProjectWithdrawalTransaction(ctx, types.ProjectDid, types.ValidWithdrawalInfo)
	appl.ProjectKeeper.AddProjectWithdrawalTransaction(ctx, types.ProjectDid, types.ValidWithdrawalInfo)

	query := abci.RequestQuery{
		Path: "",
		Data: []byte{},
	}

	querier := keeper.NewQuerier(appl.ProjectKeeper, legacyAmino)
	res, err := querier(ctx, []string{keeper.QueryProjectTx, types.ProjectDid}, query)
	require.Nil(t, err)

	var txs types.WithdrawalInfoDocs
	legacyAmino.MustUnmarshalJSON(res, &txs)

	_, err = querier(ctx, []string{keeper.QueryProjectTx, "InvalidDid"}, query)
	require.NotNil(t, err)

}

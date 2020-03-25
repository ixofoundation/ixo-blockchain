package keeper

import (
	"encoding/json"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"

	"github.com/ixofoundation/ixo-cosmos/x/project/internal/types"
)

func TestQueryProjectDoc(t *testing.T) {
	ctx, k, cdc, _, _, _ := CreateTestInput()
	codec.RegisterCrypto(cdc)
	cdc.RegisterInterface((*exported.Account)(nil), nil)
	cdc.RegisterConcrete(&auth.BaseAccount{}, "", nil)

	err := k.SetProjectDoc(ctx, &types.ValidCreateProjectMsg)
	require.Nil(t, err)

	query := abciTypes.RequestQuery{
		Path: "",
		Data: []byte{},
	}

	querier := NewQuerier(k)
	res, err := querier(ctx, []string{"queryProjectDoc", types.ValidCreateProjectMsg.ProjectDid}, query)
	require.Nil(t, err)

	emptyRes, err := querier(ctx, []string{"queryProjectDoc", "InvalidProjectDid"}, query)
	require.Nil(t, emptyRes)
	require.NotNil(t, err)

	var projectDoc types.MsgCreateProject
	cdc.MustUnmarshalJSON(res, &projectDoc)
}

func TestQueryProjectAccounts(t *testing.T) {
	ctx, k, cdc, _, _, _ := CreateTestInput()
	codec.RegisterCrypto(cdc)
	cdc.RegisterInterface((*exported.Account)(nil), nil)
	cdc.RegisterConcrete(&auth.BaseAccount{}, "", nil)

	err := k.SetProjectDoc(ctx, &types.ValidCreateProjectMsg)
	require.Nil(t, err)

	query := abciTypes.RequestQuery{
		Path: "",
		Data: []byte{},
	}

	querier := NewQuerier(k)
	_, err = querier(ctx, []string{QueryProjectDoc, types.ValidCreateProjectMsg.ProjectDid}, query)
	require.Nil(t, err)

	account, err := k.CreateNewAccount(ctx, types.ValidCreateProjectMsg.ProjectDid, types.ValidAddress1.String())
	require.Nil(t, err)
	k.AddAccountToProjectAccounts(ctx, types.ValidCreateProjectMsg.ProjectDid, types.ValidAddress1.String(), account)

	res, err := querier(ctx, []string{QueryProjectAccount, types.ValidCreateProjectMsg.ProjectDid}, query)
	require.Nil(t, err)

	var data interface{}
	require.Nil(t, json.Unmarshal(res, &data))

	accountMap := data.(map[string]interface{})
	_, errRes := json.Marshal(accountMap)
	require.Nil(t, errRes)

	account, err = k.CreateNewAccount(ctx, types.ValidCreateProjectMsg.ProjectDid, types.ValidAddress1.String())
	require.NotNil(t, err)
}

func TestQueryTxs(t *testing.T) {
	ctx, k, cdc, _, _, _ := CreateTestInput()
	codec.RegisterCrypto(cdc)
	cdc.RegisterInterface((*exported.Account)(nil), nil)
	cdc.RegisterConcrete(&auth.BaseAccount{}, "", nil)

	err := k.SetProjectDoc(ctx, &types.ValidCreateProjectMsg)
	require.Nil(t, err)

	k.AddProjectWithdrawalTransaction(ctx, types.ValidCreateProjectMsg.ProjectDid, types.ValidWithdrawalInfo)
	k.AddProjectWithdrawalTransaction(ctx, types.ValidCreateProjectMsg.ProjectDid, types.ValidWithdrawalInfo)

	query := abciTypes.RequestQuery{
		Path: "",
		Data: []byte{},
	}

	querier := NewQuerier(k)
	res, err := querier(ctx, []string{QueryProjectTx, types.ValidCreateProjectMsg.ProjectDid}, query)
	require.Nil(t, err)

	var txs []types.WithdrawalInfo
	cdc.MustUnmarshalJSON(res, &txs)

	_, err = querier(ctx, []string{QueryProjectTx, "InvalidDid"}, query)
	require.NotNil(t, err)

}

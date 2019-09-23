package keeper

import (
	"testing"
	
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/stretchr/testify/require"
	
	"github.com/ixofoundation/ixo-cosmos/x/project/internal/types"
)

func TestProjectDoc(t *testing.T) {
	ctx, k, _, _, _, _ := CreateTestInput()
	
	err := k.SetProjectDoc(ctx, &types.ValidCreateProjectMsg)
	require.Nil(t, err)
	
	err = k.SetProjectDoc(ctx, &types.ValidCreateProjectMsg)
	require.NotNil(t, err)
	
	doc, err := k.GetProjectDoc(ctx, types.ValidCreateProjectMsg.ProjectDid)
	require.Nil(t, err)
	require.Equal(t, &types.ValidCreateProjectMsg, doc)
	
	resUpdated, err := k.UpdateProjectDoc(ctx, &types.ValidUpdateProjectMsg)
	require.Nil(t, err)
	
	expected, err := k.GetProjectDoc(ctx, types.ValidUpdateProjectMsg.ProjectDid)
	require.Equal(t, resUpdated, expected)
	
	_, err = k.GetProjectDoc(ctx, "Invalid Did")
	require.NotNil(t, err)
}

func TestKeeperAccountMap(t *testing.T) {
	ctx, k, cdc, _, _, _ := CreateTestInput()
	codec.RegisterCrypto(cdc)
	cdc.RegisterInterface((*exported.Account)(nil), nil)
	cdc.RegisterConcrete(&auth.BaseAccount{}, "", nil)
	
	account, err := k.CreateNewAccount(ctx, types.ValidCreateProjectMsg.ProjectDid, types.ValidAddress1.String())
	require.Nil(t, err)
	
	k.AddAccountToProjectAccounts(ctx, types.ValidCreateProjectMsg.ProjectDid, types.ValidAddress1.String(), account)
	
	accountMap := k.GetAccountMap(ctx, types.ValidCreateProjectMsg.ProjectDid)
	_, found := accountMap[types.ValidAddress1.String()]
	require.True(t, found)
	
	account, err = k.CreateNewAccount(ctx, types.ValidCreateProjectMsg.ProjectDid, types.ValidAddress1.String())
	require.NotNil(t, err)
	
}

func TestKeeperWithdrawalInfo(t *testing.T) {
	ctx, k, cdc, _, _, _ := CreateTestInput()
	codec.RegisterCrypto(cdc)
	
	withdrawals, err := k.GetProjectWithdrawalTransactions(ctx, "")
	require.NotNil(t, err)
	require.Equal(t, 0, len(withdrawals))
	
	k.AddProjectWithdrawalTransaction(ctx, types.ValidCreateProjectMsg.ProjectDid, types.ValidWithdrawalInfo)
	k.AddProjectWithdrawalTransaction(ctx, types.ValidCreateProjectMsg.ProjectDid, types.ValidWithdrawalInfo)
	
	withdrawals, err = k.GetProjectWithdrawalTransactions(ctx, types.ValidCreateProjectMsg.ProjectDid)
	require.Nil(t, err)
	require.Equal(t, 2, len(withdrawals))
}

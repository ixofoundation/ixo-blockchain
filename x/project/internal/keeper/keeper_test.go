package keeper

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/stretchr/testify/require"

	"github.com/ixofoundation/ixo-blockchain/x/project/internal/types"
)

func TestProjectDoc(t *testing.T) {
	ctx, k, _, _, _ := CreateTestInput()

	require.False(t, k.ProjectDocExists(ctx, types.ProjectDid))
	k.SetProjectDoc(ctx, types.ValidProjectDoc)

	doc, err := k.GetProjectDoc(ctx, types.ProjectDid)
	require.Nil(t, err)
	require.Equal(t, types.ValidProjectDoc, doc)

	k.SetProjectDoc(ctx, types.ValidUpdatedProjectDoc)
	require.Nil(t, err)

	docUpdated, err := k.GetProjectDoc(ctx, types.ProjectDid)
	require.Nil(t, err)
	require.Equal(t, types.ValidProjectDoc, doc)

	expected, err := k.GetProjectDoc(ctx, types.ValidUpdatedProjectDoc.ProjectDid)
	require.Equal(t, docUpdated, expected)

	_, err = k.GetProjectDoc(ctx, "Invalid Did")
	require.NotNil(t, err)
}

func TestKeeperAccountMap(t *testing.T) {
	ctx, k, cdc, _, _ := CreateTestInput()
	codec.RegisterCrypto(cdc)
	cdc.RegisterInterface((*exported.Account)(nil), nil)
	cdc.RegisterConcrete(&auth.BaseAccount{}, "", nil)

	account, err := k.CreateNewAccount(ctx, types.ProjectDid, types.ValidAccId1)
	require.Nil(t, err)

	k.AddAccountToProjectAccounts(ctx, types.ProjectDid, types.ValidAccId1, account)

	accountMap := k.GetAccountMap(ctx, types.ProjectDid)
	_, found := accountMap[types.ValidAccId1]
	require.True(t, found)

	account, err = k.CreateNewAccount(ctx, types.ProjectDid, types.ValidAccId1)
	require.NotNil(t, err)

}

func TestKeeperWithdrawalInfo(t *testing.T) {
	ctx, k, cdc, _, _ := CreateTestInput()
	codec.RegisterCrypto(cdc)

	withdrawals, err := k.GetProjectWithdrawalTransactions(ctx, "")
	require.NotNil(t, err)
	require.Equal(t, 0, len(withdrawals))

	k.AddProjectWithdrawalTransaction(ctx, types.ProjectDid, types.ValidWithdrawalInfo)
	k.AddProjectWithdrawalTransaction(ctx, types.ProjectDid, types.ValidWithdrawalInfo)

	withdrawals, err = k.GetProjectWithdrawalTransactions(ctx, types.ProjectDid)
	require.Nil(t, err)
	require.Equal(t, 2, len(withdrawals))
}

package keeper

import (
	"github.com/ixofoundation/ixo-blockchain/x/payments"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ixofoundation/ixo-blockchain/x/project/internal/types"
)

func TestProjectDoc(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)

	// Project starts off not existing
	require.False(t, k.ProjectDocExists(ctx, types.ProjectDid))
	_, err := k.GetProjectDoc(ctx, types.ProjectDid)
	require.Error(t, err)

	// After setting the project, it now exists
	k.SetProjectDoc(ctx, types.ValidProjectDoc)
	require.True(t, k.ProjectDocExists(ctx, types.ProjectDid))

	// Check that project matches what we set it as
	doc, err := k.GetProjectDoc(ctx, types.ProjectDid)
	require.NoError(t, err)
	require.Equal(t, types.ValidProjectDoc, doc)

	// Update project doc
	k.SetProjectDoc(ctx, types.ValidUpdatedProjectDoc)
	require.NoError(t, err)

	// Check that updated project matches what we set it as
	docUpdated, err := k.GetProjectDoc(ctx, types.ProjectDid)
	require.NoError(t, err)
	require.Equal(t, types.ValidUpdatedProjectDoc, docUpdated)
}

func TestValidateProjectFeesMap(t *testing.T) {
	ctx, k, pk, _ := CreateTestInput(t, false)

	templateId1 := "payment:template:1"
	templateId2 := "payment:template:2"
	pk.SetPaymentTemplate(ctx, payments.NewPaymentTemplate(templateId1, nil, nil, nil, nil))
	pk.SetPaymentTemplate(ctx, payments.NewPaymentTemplate(templateId2, nil, nil, nil, nil))

	testCases := []struct {
		feesMap     types.ProjectFeesMap
		expectError bool
	}{
		{ // no fees
			feesMap: types.ProjectFeesMap{
				Context: "",
				Items:   []types.ProjectFeesMapItem{},
			},
			expectError: false,
		},
		{ // non-existent payment template used
			feesMap: types.ProjectFeesMap{
				Context: "",
				Items: []types.ProjectFeesMapItem{
					{Type: types.OracleFee, PaymentTemplateId: "non-existent-payment-template"},
				},
			},
			expectError: true,
		},
		{ // existent payment template used, reuse of template IDs, and a blank project fee type
			feesMap: types.ProjectFeesMap{
				Context: "",
				Items: []types.ProjectFeesMapItem{
					{Type: types.OracleFee, PaymentTemplateId: templateId1},
					{Type: types.FeeForService, PaymentTemplateId: templateId2},
					{Type: types.OutcomePayment, PaymentTemplateId: templateId1},
					{Type: "", PaymentTemplateId: templateId2},
				},
			},
			expectError: false,
		},
	}

	for _, tc := range testCases {
		err := k.ValidateProjectFeesMap(ctx, tc.feesMap)
		if tc.expectError {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}

func TestKeeperAccountMap(t *testing.T) {
	ctx, k, _, _ := CreateTestInput(t, false)

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
	ctx, k, _, _ := CreateTestInput(t, false)

	withdrawals, err := k.GetProjectWithdrawalTransactions(ctx, "")
	require.NotNil(t, err)
	require.Equal(t, 0, len(withdrawals))

	k.AddProjectWithdrawalTransaction(ctx, types.ProjectDid, types.ValidWithdrawalInfo)
	k.AddProjectWithdrawalTransaction(ctx, types.ProjectDid, types.ValidWithdrawalInfo)

	withdrawals, err = k.GetProjectWithdrawalTransactions(ctx, types.ProjectDid)
	require.Nil(t, err)
	require.Equal(t, 2, len(withdrawals))
}

package keeper_test

import (
	paymentstypes "github.com/ixofoundation/ixo-blockchain/x/payments/types"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ixofoundation/ixo-blockchain/x/project/types"
)

func TestProjectDoc(t *testing.T) {
	_, appl, ctx := CreateTestInput()

	// Project starts off not existing
	require.False(t, appl.ProjectKeeper.ProjectDocExists(ctx, types.ProjectDid))
	_, err := appl.ProjectKeeper.GetProjectDoc(ctx, types.ProjectDid)
	require.Error(t, err)

	// After setting the project, it now exists
	appl.ProjectKeeper.SetProjectDoc(ctx, types.ValidProjectDoc)
	require.True(t, appl.ProjectKeeper.ProjectDocExists(ctx, types.ProjectDid))

	// Check that project matches what we set it as
	doc, err := appl.ProjectKeeper.GetProjectDoc(ctx, types.ProjectDid)
	require.NoError(t, err)
	require.Equal(t, types.ValidProjectDoc, doc)

	// Update project doc
	appl.ProjectKeeper.SetProjectDoc(ctx, types.ValidUpdatedProjectDoc)
	require.NoError(t, err)

	// Check that updated project matches what we set it as
	docUpdated, err := appl.ProjectKeeper.GetProjectDoc(ctx, types.ProjectDid)
	require.NoError(t, err)
	require.Equal(t, types.ValidUpdatedProjectDoc, docUpdated)
}

func TestValidateProjectFeesMap(t *testing.T) {
	_, appl, ctx := CreateTestInput()

	templateId1 := "payment:template:1"
	templateId2 := "payment:template:2"
	appl.PaymentsKeeper.SetPaymentTemplate(ctx, paymentstypes.NewPaymentTemplate(templateId1, nil, nil, nil, nil))
	appl.PaymentsKeeper.SetPaymentTemplate(ctx, paymentstypes.NewPaymentTemplate(templateId2, nil, nil, nil, nil))

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
		err := appl.ProjectKeeper.ValidateProjectFeesMap(ctx, tc.feesMap)
		if tc.expectError {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}

func TestKeeperAccountMap(t *testing.T) {
	_, appl, ctx := CreateTestInput()

	account, err := appl.ProjectKeeper.CreateNewAccount(ctx, types.ProjectDid, types.ValidAccId1)
	require.Nil(t, err)

	appl.ProjectKeeper.AddAccountToProjectAccounts(ctx, types.ProjectDid, types.ValidAccId1, account)

	accountMap := appl.ProjectKeeper.GetAccountMap(ctx, types.ProjectDid)
	_, found := accountMap.Map[string(types.ValidAccId1)]
	require.True(t, found)

	account, err = appl.ProjectKeeper.CreateNewAccount(ctx, types.ProjectDid, types.ValidAccId1)
	require.NotNil(t, err)

}

func TestKeeperWithdrawalInfo(t *testing.T) {
	_, appl, ctx := CreateTestInput()

	withdrawals, err := appl.ProjectKeeper.GetProjectWithdrawalTransactions(ctx, "")
	require.NotNil(t, err)
	require.Equal(t, 0, len(withdrawals.DocsList))

	appl.ProjectKeeper.AddProjectWithdrawalTransaction(ctx, types.ProjectDid, types.ValidWithdrawalInfo)
	appl.ProjectKeeper.AddProjectWithdrawalTransaction(ctx, types.ProjectDid, types.ValidWithdrawalInfo)

	withdrawals, err = appl.ProjectKeeper.GetProjectWithdrawalTransactions(ctx, types.ProjectDid)
	require.Nil(t, err)
	require.Equal(t, 2, len(withdrawals.DocsList))
}

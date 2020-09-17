package project

import (
	"encoding/json"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/stretchr/testify/require"

	"github.com/ixofoundation/ixo-blockchain/x/project/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/project/internal/types"
)

func TestHandler_CreateClaim(t *testing.T) {

	ctx, k, cdc, _, _ := keeper.CreateTestInput()
	codec.RegisterCrypto(cdc)
	cdc.RegisterConcrete(types.MsgCreateProject{}, "project/CreateProject", nil)
	cdc.RegisterInterface((*exported.Account)(nil), nil)
	cdc.RegisterConcrete(&auth.BaseAccount{}, "cosmos-sdk/Account", nil)

	projectDid := "6iftm1hHdaU6LJGKayRMev"
	txHash := "txHash"
	senderDid := "senderDid"
	data := types.NewCreateClaimDoc("claim1")

	msg := types.NewMsgCreateClaim(txHash, senderDid, data, projectDid)

	res := handleMsgCreateClaim(ctx, k, msg)
	require.NotNil(t, res)
}

func TestHandler_ProjectMsg(t *testing.T) {
	ctx, k, cdc, _, _ := keeper.CreateTestInput()
	codec.RegisterCrypto(cdc)
	types.RegisterCodec(cdc)
	cdc.RegisterInterface((*exported.Account)(nil), nil)
	cdc.RegisterConcrete(&auth.BaseAccount{}, "cosmos-sdk/Account", nil)

	res := handleMsgCreateProject(ctx, k, types.ValidCreateProjectMsg)
	require.True(t, res.IsOK())

	res = handleMsgCreateProject(ctx, k, types.ValidCreateProjectMsg)
	require.False(t, res.IsOK())

}
func Test_CreateEvaluation(t *testing.T) {
	ctx, k, cdc, pk, bk := keeper.CreateTestInput()

	codec.RegisterCrypto(cdc)
	types.RegisterCodec(cdc)
	cdc.RegisterInterface((*exported.Account)(nil), nil)
	cdc.RegisterConcrete(&auth.BaseAccount{}, "cosmos-sdk/Account", nil)

	params := types.DefaultParams()
	params.OracleFeePercentage = sdk.ZeroDec()
	params.NodeFeePercentage = sdk.NewDecWithPrec(5, 1)
	k.SetParams(ctx, params)

	txHash := "txHash"
	senderDid := "senderDid"
	projectDid := "6iftm1hHdaU6LJGKayRMev"

	evaluationMsg := types.NewMsgCreateEvaluation(txHash, senderDid,
		types.NewCreateEvaluationDoc("claim1", types.PendingClaim), projectDid)

	projectData := struct {
		NodeDid         string
		RequiredClaims  string
		ServiceEndpoint string
		CreatedOn       string
		CreatedBy       string
	}{
		NodeDid:         "Tu2QWRHuDufywDALbBQ2r",
		RequiredClaims:  "requireClaims1",
		ServiceEndpoint: "https://togo.pds.ixo.network",
		CreatedOn:       "2018-05-21T15:53:18.484Z",
		CreatedBy:       "6Fu7FbbGoCJ8tX3vMMCss9",
	}

	txHash = ""
	senderDid = ""
	projectDid = "6iftm1hHdaU6LJGKayRMev"
	pubKey := "47mm6LCDAyJmqkbUbqGoZKZkBixjBgvDFRMwQRF9HWMU"

	// Create project doc with data
	projectDoc := types.NewProjectDoc(
		txHash, projectDid, senderDid, pubKey, "CREATED", json.RawMessage{})
	projectDocData, err2 := json.Marshal(projectData)
	require.Nil(t, err2)
	projectDoc.Data = projectDocData

	// Create project creation message
	msg := types.NewMsgCreateProject(
		senderDid, projectDoc.Data, projectDid, pubKey)

	var err sdk.Error
	_, err = createAccountInProjectAccounts(ctx, k, msg.ProjectDid, IxoAccountFeesId)
	require.Nil(t, err)
	_, err = createAccountInProjectAccounts(ctx, k, msg.ProjectDid, InternalAccountID(msg.ProjectDid))
	require.Nil(t, err)

	require.False(t, k.ProjectDocExists(ctx, msg.ProjectDid))
	k.SetProjectDoc(ctx, projectDoc)

	res := handleMsgCreateEvaluation(ctx, k, pk, bk, evaluationMsg)
	require.NotNil(t, res)
}

func Test_WithdrawFunds(t *testing.T) {
	ctx, k, cdc, _, bk := keeper.CreateTestInput()
	codec.RegisterCrypto(cdc)
	types.RegisterCodec(cdc)
	cdc.RegisterInterface((*exported.Account)(nil), nil)
	cdc.RegisterConcrete(&auth.BaseAccount{}, "cosmos-sdk/Account", nil)

	senderDid := "6iftm1hHdaU6LJGKayRMev"
	projectDid := senderDid
	recipientDid := senderDid
	amount := sdk.NewInt(100)
	isRefund := true

	msg := types.NewMsgWithdrawFunds(senderDid,
		types.NewWithdrawFundsDoc(projectDid, recipientDid, amount, isRefund))

	projectData := struct {
		NodeDid         string
		RequiredClaims  string
		ServiceEndpoint string
		CreatedOn       string
		CreatedBy       string
	}{
		NodeDid:         "Tu2QWRHuDufywDALbBQ2r",
		RequiredClaims:  "requireClaims1",
		ServiceEndpoint: "https://togo.pds.ixo.network",
		CreatedOn:       "2018-05-21T15:53:18.484Z",
		CreatedBy:       "6Fu7FbbGoCJ8tX3vMMCss9",
	}

	txHash := ""
	senderDid = ""
	projectDid = "6iftm1hHdaU6LJGKayRMev"
	pubKey := "47mm6LCDAyJmqkbUbqGoZKZkBixjBgvDFRMwQRF9HWMU"

	// Create project doc with data
	projectDoc := types.NewProjectDoc(
		txHash, projectDid, senderDid, pubKey, "PAIDOUT", json.RawMessage{})
	projectDocData, err2 := json.Marshal(projectData)
	require.Nil(t, err2)
	projectDoc.Data = projectDocData

	// Create project creation message
	msg2 := types.NewMsgCreateProject(
		senderDid, projectDoc.Data, projectDid, pubKey)

	var err sdk.Error
	_, err = createAccountInProjectAccounts(ctx, k, msg2.ProjectDid, IxoAccountFeesId)
	require.Nil(t, err)
	_, err = createAccountInProjectAccounts(ctx, k, msg2.ProjectDid, InternalAccountID(msg2.ProjectDid))
	require.Nil(t, err)

	require.False(t, k.ProjectDocExists(ctx, msg2.ProjectDid))
	k.SetProjectDoc(ctx, projectDoc)

	res := handleMsgWithdrawFunds(ctx, k, bk, msg)
	require.NotNil(t, res)
}

package project
//
//import (
//	"testing"
//
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	"github.com/stretchr/testify/require"
//
//	"github.com/ixofoundation/ixo-blockchain/x/project/internal/keeper"
//	"github.com/ixofoundation/ixo-blockchain/x/project/internal/types"
//)
//
//func TestHandler_CreateClaim(t *testing.T) {
//
//	ctx, k, _, _ := keeper.CreateTestInput(t, false)
//
//	projectDid := types.ValidCreateProjectMsg.ProjectDid
//	txHash := "txHash"
//	senderDid := "senderDid"
//	data := types.NewCreateClaimDoc("claim1")
//
//	res, _ := handleMsgCreateProject(ctx, k, types.ValidCreateProjectMsg)
//	require.NotNil(t, res)
//
//	// Set status to STARTED
//	project, err := k.GetProjectDoc(ctx, projectDid)
//	require.NoError(t, err)
//	project.Status = types.StartedStatus
//	k.SetProjectDoc(ctx, project)
//
//	msg := types.NewMsgCreateClaim(txHash, senderDid, data, projectDid)
//
//	res, err = handleMsgCreateClaim(ctx, k, msg)
//	require.NoError(t, err)
//	require.NotNil(t, res)
//}
//
//func TestHandler_ProjectMsg(t *testing.T) {
//	ctx, k, _, _ := keeper.CreateTestInput(t, false)
//
//	res, _ := handleMsgCreateProject(ctx, k, types.ValidCreateProjectMsg)
//	require.NotNil(t, res)
//
//	res, _ = handleMsgCreateProject(ctx, k, types.ValidCreateProjectMsg)
//	require.Nil(t, res)
//
//}
//func TestHandler_CreateEvaluation(t *testing.T) {
//	ctx, k, pk, bk := keeper.CreateTestInput(t, false)
//
//	params := types.DefaultParams()
//	params.IxoDid = "blank"
//	params.OracleFeePercentage = sdk.ZeroDec()
//	params.NodeFeePercentage = sdk.NewDecWithPrec(5, 1)
//	k.SetParams(ctx, params)
//
//	projectDid := types.ValidCreateProjectMsg.ProjectDid
//	txHash := "txHash"
//	senderDid := "senderDid"
//
//	evaluationMsg := types.NewMsgCreateEvaluation(txHash, senderDid,
//		types.NewCreateEvaluationDoc("claim1", types.PendingClaim), projectDid)
//
//	res, _ := handleMsgCreateProject(ctx, k, types.ValidCreateProjectMsg)
//	require.NotNil(t, res)
//
//	// Set status to STARTED
//	project, err := k.GetProjectDoc(ctx, projectDid)
//	require.NoError(t, err)
//	project.Status = types.StartedStatus
//	k.SetProjectDoc(ctx, project)
//
//	msg2 := types.NewMsgCreateClaim(txHash, senderDid, types.NewCreateClaimDoc("claim1"), projectDid)
//	res, err = handleMsgCreateClaim(ctx, k, msg2)
//	require.NoError(t, err)
//	require.NotNil(t, res)
//
//	res, err = handleMsgCreateEvaluation(ctx, k, pk, bk, evaluationMsg)
//	require.Nil(t, err)
//	require.NotNil(t, res)
//}

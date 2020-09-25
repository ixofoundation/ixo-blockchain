package project

import (
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

	projectDid := types.ValidCreateProjectMsg.ProjectDid
	txHash := "txHash"
	senderDid := "senderDid"
	data := types.NewCreateClaimDoc("claim1")

	res, _ := handleMsgCreateProject(ctx, k, types.ValidCreateProjectMsg)
	require.NotNil(t, res)

	// Set status to STARTED
	project, err := k.GetProjectDoc(ctx, projectDid)
	require.NoError(t, err)
	project.Status = types.StartedStatus
	k.SetProjectDoc(ctx, project)

	msg := types.NewMsgCreateClaim(txHash, senderDid, data, projectDid)

	res, err = handleMsgCreateClaim(ctx, k, msg)
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestHandler_ProjectMsg(t *testing.T) {
	ctx, k, cdc, _, _ := keeper.CreateTestInput()
	codec.RegisterCrypto(cdc)
	types.RegisterCodec(cdc)
	cdc.RegisterInterface((*exported.Account)(nil), nil)
	cdc.RegisterConcrete(&auth.BaseAccount{}, "cosmos-sdk/Account", nil)

	res, _ := handleMsgCreateProject(ctx, k, types.ValidCreateProjectMsg)
	require.NotNil(t, res)

	res, _ = handleMsgCreateProject(ctx, k, types.ValidCreateProjectMsg)
	require.Nil(t, res)

}
func TestHandler_CreateEvaluation(t *testing.T) {
	ctx, k, cdc, pk, bk := keeper.CreateTestInput()

	codec.RegisterCrypto(cdc)
	types.RegisterCodec(cdc)
	cdc.RegisterInterface((*exported.Account)(nil), nil)
	cdc.RegisterConcrete(&auth.BaseAccount{}, "cosmos-sdk/Account", nil)

	params := types.DefaultParams()
	params.IxoDid = "blank"
	params.OracleFeePercentage = sdk.ZeroDec()
	params.NodeFeePercentage = sdk.NewDecWithPrec(5, 1)
	k.SetParams(ctx, params)

	projectDid := types.ValidCreateProjectMsg.ProjectDid
	txHash := "txHash"
	senderDid := "senderDid"

	evaluationMsg := types.NewMsgCreateEvaluation(txHash, senderDid,
		types.NewCreateEvaluationDoc("claim1", types.PendingClaim), projectDid)

	res, _ := handleMsgCreateProject(ctx, k, types.ValidCreateProjectMsg)
	require.NotNil(t, res)

	// Set status to STARTED
	project, err := k.GetProjectDoc(ctx, projectDid)
	require.NoError(t, err)
	project.Status = types.StartedStatus
	k.SetProjectDoc(ctx, project)

	msg2 := types.NewMsgCreateClaim(txHash, senderDid, types.NewCreateClaimDoc("claim1"), projectDid)
	res, err = handleMsgCreateClaim(ctx, k, msg2)
	require.NoError(t, err)
	require.NotNil(t, res)

	res, err = handleMsgCreateEvaluation(ctx, k, pk, bk, evaluationMsg)
	require.Nil(t, err)
	require.NotNil(t, res)
}

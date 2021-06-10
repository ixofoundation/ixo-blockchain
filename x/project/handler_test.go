package project_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ixofoundation/ixo-blockchain/app"
	"github.com/ixofoundation/ixo-blockchain/cmd"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/ixofoundation/ixo-blockchain/x/project/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/project/types"
)

func CreateTestInput() (*codec.LegacyAmino, *app.IxoApp, sdk.Context) {
	appl := cmd.Setup(false)
	ctx := appl.BaseApp.NewContext(false, tmproto.Header{})
	ctx = ctx.WithConsensusParams(
		&abci.ConsensusParams{
			Validator: &tmproto.ValidatorParams{
				PubKeyTypes: []string{tmtypes.ABCIPubKeyTypeEd25519},
			},
		},
	)

	feeCollectorAcc := authtypes.NewEmptyModuleAccount(authtypes.FeeCollectorName)
	notBondedPool := authtypes.NewEmptyModuleAccount(stakingtypes.NotBondedPoolName, authtypes.Burner, authtypes.Staking)
	bondPool := authtypes.NewEmptyModuleAccount(stakingtypes.BondedPoolName, authtypes.Burner, authtypes.Staking)

	blockedAddrs := make(map[string]bool)
	blockedAddrs[feeCollectorAcc.GetAddress().String()] = true
	blockedAddrs[notBondedPool.GetAddress().String()] = true
	blockedAddrs[bondPool.GetAddress().String()] = true

	appl.BankKeeper = bankkeeper.NewBaseKeeper(
		appl.AppCodec(),
		appl.GetKey(banktypes.StoreKey),
		appl.AccountKeeper,
		appl.GetSubspace(banktypes.ModuleName),
		blockedAddrs,
	)

	return appl.LegacyAmino(), appl, ctx
}

func TestHandler_CreateClaim(t *testing.T) {
	_, appl, ctx := CreateTestInput()
	msgServer := keeper.NewMsgServerImpl(appl.ProjectKeeper, appl.BankKeeper, appl.PaymentsKeeper)

	projectDid := types.ValidCreateProjectMsg.ProjectDid
	txHash := "txHash"
	senderDid := "senderDid"
	data := types.NewCreateClaimDoc("claim1", "claimTemplateA")

	resProj, _ := msgServer.CreateProject(sdk.WrapSDKContext(ctx), &types.ValidCreateProjectMsg)
	require.NotNil(t, resProj)

	// Set status to STARTED
	project, err := appl.ProjectKeeper.GetProjectDoc(ctx, projectDid)
	require.NoError(t, err)
	project.Status = string(types.StartedStatus)
	appl.ProjectKeeper.SetProjectDoc(ctx, project)

	msg := types.NewMsgCreateClaim(txHash, senderDid, data, projectDid)

	resClaim, err := msgServer.CreateClaim(sdk.WrapSDKContext(ctx), msg)
	require.NoError(t, err)
	require.NotNil(t, resClaim)
}

func TestHandler_ProjectMsg(t *testing.T) {
	_, appl, ctx := CreateTestInput()
	msgServer := keeper.NewMsgServerImpl(appl.ProjectKeeper, appl.BankKeeper, appl.PaymentsKeeper)

	res, _ := msgServer.CreateProject(sdk.WrapSDKContext(ctx), &types.ValidCreateProjectMsg)
	require.NotNil(t, res)

	res, _ = msgServer.CreateProject(sdk.WrapSDKContext(ctx), &types.ValidCreateProjectMsg)
	require.Nil(t, res)

}
func TestHandler_CreateEvaluation(t *testing.T) {
	_, appl, ctx := CreateTestInput()
	msgServer := keeper.NewMsgServerImpl(appl.ProjectKeeper, appl.BankKeeper, appl.PaymentsKeeper)

	params := types.DefaultParams()
	params.IxoDid = "blank"
	params.OracleFeePercentage = sdk.ZeroDec()
	params.NodeFeePercentage = sdk.NewDecWithPrec(5, 1)
	appl.ProjectKeeper.SetParams(ctx, params)

	projectDid := types.ValidCreateProjectMsg.ProjectDid
	txHash := "txHash"
	senderDid := "senderDid"

	evaluationMsg := types.NewMsgCreateEvaluation(txHash, senderDid,
		types.NewCreateEvaluationDoc("claim1", types.PendingClaim), projectDid)

	resProj, _ := msgServer.CreateProject(sdk.WrapSDKContext(ctx), &types.ValidCreateProjectMsg)
	require.NotNil(t, resProj)

	// Set status to STARTED
	project, err := appl.ProjectKeeper.GetProjectDoc(ctx, projectDid)
	require.NoError(t, err)
	project.Status = string(types.StartedStatus)
	appl.ProjectKeeper.SetProjectDoc(ctx, project)

	msg2 := types.NewMsgCreateClaim(txHash, senderDid, types.NewCreateClaimDoc("claim1", "claimTemplateA"), projectDid)
	resClaim, err := msgServer.CreateClaim(sdk.WrapSDKContext(ctx), msg2)
	require.NoError(t, err)
	require.NotNil(t, resClaim)

	resEval, err := msgServer.CreateEvaluation(sdk.WrapSDKContext(ctx), evaluationMsg)
	require.Nil(t, err)
	require.NotNil(t, resEval)
}

package project

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/stretchr/testify/require"

	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"github.com/ixofoundation/ixo-blockchain/x/project/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/project/internal/types"
)

func TestHandler_CreateClaim(t *testing.T) {

	ctx, k, cdc, feesKeeper, bankKeeper := keeper.CreateTestInput()
	codec.RegisterCrypto(cdc)
	cdc.RegisterConcrete(types.MsgCreateProject{}, "project/CreateProject", nil)
	cdc.RegisterInterface((*exported.Account)(nil), nil)
	cdc.RegisterConcrete(&auth.BaseAccount{}, "cosmos-sdk/Account", nil)
	params := feesKeeper.GetParams(ctx)
	params.IxoFactor = sdk.OneDec()
	params.NodeFeePercentage = sdk.ZeroDec()
	params.ClaimFeeAmount = sdk.NewDec(6).Quo(sdk.NewDec(10)).Mul(ixo.IxoDecimals)
	feesKeeper.SetParams(ctx, params)
	projectMsg := types.MsgCreateClaim{
		ProjectDid: "6iftm1hHdaU6LJGKayRMev",
		TxHash:     "txHash",
		SenderDid:  "senderDid",
		Data:       types.CreateClaimDoc{ClaimID: "claim1"},
	}

	res := handleMsgCreateClaim(ctx, k, feesKeeper, bankKeeper, projectMsg)
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
	ctx, k, cdc, fk, bk := keeper.CreateTestInput()

	codec.RegisterCrypto(cdc)
	types.RegisterCodec(cdc)
	cdc.RegisterInterface((*exported.Account)(nil), nil)
	cdc.RegisterConcrete(&auth.BaseAccount{}, "cosmos-sdk/Account", nil)

	params := fk.GetParams(ctx)
	params.IxoFactor = sdk.OneDec()
	params.NodeFeePercentage = sdk.NewDec(5).Quo(sdk.NewDec(10))
	params.ClaimFeeAmount = sdk.NewDec(6).Quo(sdk.NewDec(10)).Mul(ixo.IxoDecimals)
	params.EvaluationFeeAmount = sdk.NewDec(4).Quo(sdk.NewDec(10)).Mul(ixo.IxoDecimals) // 0.4
	params.EvaluationPayFeePercentage = sdk.ZeroDec()
	params.EvaluationPayNodeFeePercentage = sdk.NewDec(5).Quo(sdk.NewDec(10))
	fk.SetParams(ctx, params)

	evaluationMsg := types.MsgCreateEvaluation{
		TxHash:     "txHash",
		SenderDid:  "senderDid",
		ProjectDid: "6iftm1hHdaU6LJGKayRMev",
		Data: types.CreateEvaluationDoc{
			ClaimID: "claim1",
			Status:  types.PendingClaim,
		},
	}

	msg := types.MsgCreateProject{
		TxHash:     "",
		SenderDid:  "",
		ProjectDid: "6iftm1hHdaU6LJGKayRMev",
		PubKey:     "47mm6LCDAyJmqkbUbqGoZKZkBixjBgvDFRMwQRF9HWMU",
		Data: types.ProjectDoc{
			NodeDid:              "Tu2QWRHuDufywDALbBQ2r",
			RequiredClaims:       "requireClaims1",
			EvaluatorPayPerClaim: "10",
			ServiceEndpoint:      "https://togo.pds.ixo.network",
			CreatedOn:            "2018-05-21T15:53:18.484Z",
			CreatedBy:            "6Fu7FbbGoCJ8tX3vMMCss9",
			Status:               "CREATED",
		},
	}

	var err sdk.Error
	_, err = createAccountInProjectAccounts(ctx, k, msg.GetProjectDid(), IxoAccountFeesId)
	require.Nil(t, err)
	_, err = createAccountInProjectAccounts(ctx, k, msg.GetProjectDid(), InternalAccountID(msg.GetProjectDid()))
	require.Nil(t, err)

	require.False(t, k.ProjectDocExists(ctx, msg.GetProjectDid()))
	k.SetProjectDoc(ctx, &msg)

	res := handleMsgCreateEvaluation(ctx, k, fk, bk, evaluationMsg)
	require.NotNil(t, res)
}

func Test_WithdrawFunds(t *testing.T) {
	ctx, k, cdc, _, bk := keeper.CreateTestInput()
	codec.RegisterCrypto(cdc)
	types.RegisterCodec(cdc)
	cdc.RegisterInterface((*exported.Account)(nil), nil)
	cdc.RegisterConcrete(&auth.BaseAccount{}, "cosmos-sdk/Account", nil)

	msg := types.MsgWithdrawFunds{
		SenderDid: "6iftm1hHdaU6LJGKayRMev",
		Data: types.WithdrawFundsDoc{
			ProjectDid:   "6iftm1hHdaU6LJGKayRMev",
			RecipientDid: "6iftm1hHdaU6LJGKayRMev",
			Amount:       sdk.NewInt(100),
			IsRefund:     true,
		},
	}

	msg1 := types.MsgCreateProject{
		TxHash:     "",
		SenderDid:  "",
		ProjectDid: "6iftm1hHdaU6LJGKayRMev",
		PubKey:     "47mm6LCDAyJmqkbUbqGoZKZkBixjBgvDFRMwQRF9HWMU",
		Data: types.ProjectDoc{
			NodeDid:              "Tu2QWRHuDufywDALbBQ2r",
			RequiredClaims:       "requireClaims1",
			EvaluatorPayPerClaim: "10",
			ServiceEndpoint:      "https://togo.pds.ixo.network",
			CreatedOn:            "2018-05-21T15:53:18.484Z",
			CreatedBy:            "6Fu7FbbGoCJ8tX3vMMCss9",
			Status:               "PAIDOUT",
		},
	}

	var err sdk.Error
	_, err = createAccountInProjectAccounts(ctx, k, msg1.GetProjectDid(), IxoAccountFeesId)
	require.Nil(t, err)
	_, err = createAccountInProjectAccounts(ctx, k, msg1.GetProjectDid(), InternalAccountID(msg1.GetProjectDid()))
	require.Nil(t, err)

	require.False(t, k.ProjectDocExists(ctx, msg1.GetProjectDid()))
	k.SetProjectDoc(ctx, &msg1)

	res := handleMsgWithdrawFunds(ctx, k, bk, msg)
	require.NotNil(t, res)
}

package project

import (
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	ixo "github.com/ixofoundation/ixo-cosmos/x/ixo"
)

func NewHandler(k Keeper, ck bank.Keeper, ethClient ixo.EthClient) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case CreateProjectMsg:
			return handleCreateProjectMsg(ctx, k, ck, msg)
		case UpdateProjectStatusMsg:
			return handleUpdateProjectStatusMsg(ctx, k, ck, ethClient, msg)
		case CreateAgentMsg:
			return handleCreateAgentMsg(ctx, k, ck, msg)
		case UpdateAgentMsg:
			return handleUpdateAgentMsg(ctx, k, ck, msg)
		case CreateClaimMsg:
			return handleCreateClaimMsg(ctx, k, ck, msg)
		case CreateEvaluationMsg:
			return handleCreateEvaluationMsg(ctx, k, ck, msg)
		case WithdrawFundsMsg:
			return handleWithdrawFundsMsg(ctx, k, ck, msg)
		default:
			return sdk.ErrUnknownRequest("No match for message type.").Result()
		}
	}
}

func handleCreateProjectMsg(ctx sdk.Context, k Keeper, ck bank.Keeper, msg CreateProjectMsg) sdk.Result {
	addAccountToAccountProjectAccounts(ctx, k, msg.GetProjectDid(), msg.GetProjectDid())

	projectDoc, err := k.AddProjectDoc(ctx, &msg)
	if err != nil {
		return err.Result()
	}
	return sdk.Result{
		Code: sdk.ABCICodeOK,
		Data: k.encodeProject(projectDoc),
	}
}

func handleUpdateProjectStatusMsg(ctx sdk.Context, k Keeper, ck bank.Keeper, ethClient ixo.EthClient, msg UpdateProjectStatusMsg) sdk.Result {
	existingProjectDoc, found := getProjectDoc(ctx, k, msg.GetProjectDid())
	if !found {
		return sdk.Result{
			Code: sdk.ABCICodeType(sdk.CodeInvalidAddress),
			Data: []byte("Could not find Project"),
		}
	}

	newStatus := msg.GetStatus()
	existingProjectDoc.SetStatus(newStatus)

	if newStatus == FundedStatus {
		if existingProjectDoc.GetStatus() != CreatedProject {
			return sdk.Result{
				Code: sdk.ABCICodeType(sdk.CodeInternal),
				Data: []byte("Could not find Project"),
			}
		} else {
			// Check that the Project wallet is funded and mint equivalent tokens on project

			ethTx, err := ethClient.GetTransactionByHash(msg.GetEthFundingTxnID())
			if err != nil {
				return sdk.Result{
					Code: sdk.ABCICodeType(sdk.CodeInternal),
					Data: []byte("ETH tx not valid"),
				}
			}
			fundingTx := ethClient.IsProjectFundingTx(existingProjectDoc.GetProjectDid(), ethTx)
			if !fundingTx {
				return sdk.Result{
					Code: sdk.ABCICodeType(sdk.CodeInternal),
					Data: []byte("ETH tx not valid"),
				}
			}
			amt := ethClient.GetFundingAmt(ethTx)
			coin := sdk.NewInt64Coin(COIN_DENOM, amt)
			return fundProject(ctx, k, ck, existingProjectDoc, coin)
		}
	}

	storedProjectDoc, updated := k.UpdateProjectDoc(ctx, existingProjectDoc)
	if !updated {
		return sdk.Result{
			Code: sdk.ABCICodeType(sdk.CodeInvalidAddress),
			Data: []byte("Could not find Project"),
		}
	}
	return sdk.Result{
		Code: sdk.ABCICodeOK,
		Data: k.encodeProject(storedProjectDoc),
	}
}

func handleCreateAgentMsg(ctx sdk.Context, k Keeper, ck bank.Keeper, msg CreateAgentMsg) sdk.Result {
	addAccountToAccountProjectAccounts(ctx, k, msg.GetProjectDid(), msg.Data.AgentDid)
	return sdk.Result{
		Code: sdk.ABCICodeOK,
		Data: []byte("Action complete"),
	}
}
func handleUpdateAgentMsg(ctx sdk.Context, k Keeper, ck bank.Keeper, msg UpdateAgentMsg) sdk.Result {
	return sdk.Result{
		Code: sdk.ABCICodeOK,
		Data: []byte("Action complete"),
	}
}
func handleCreateClaimMsg(ctx sdk.Context, k Keeper, ck bank.Keeper, msg CreateClaimMsg) sdk.Result {
	return sdk.Result{
		Code: sdk.ABCICodeOK,
		Data: []byte("Action complete"),
	}
}
func handleCreateEvaluationMsg(ctx sdk.Context, k Keeper, ck bank.Keeper, msg CreateEvaluationMsg) sdk.Result {
	projectDoc, found := getProjectDoc(ctx, k, msg.GetProjectDid())
	if !found {
		return sdk.Result{
			Code: sdk.ABCICodeType(sdk.CodeInvalidAddress),
			Data: []byte("Could not find Project"),
		}
	}
	accMap := getProjectAccountMap(ctx, k, msg.GetProjectDid())
	projectAddrInterface, found := accMap[msg.GetProjectDid()]
	if !found {
		return sdk.Result{
			Code: sdk.ABCICodeType(sdk.CodeInvalidAddress),
			Data: []byte("Could not find Project Account"),
		}
	}
	projectAddr := projectAddrInterface.(string)
	senderAccAddrInterface, found := accMap[msg.GetSenderDid()]
	var senderAccAddr string
	if !found {
		newAcc := addAccountToAccountProjectAccounts(ctx, k, msg.GetProjectDid(), msg.GetSenderDid())

		senderAccAddr = hex.EncodeToString(newAcc.GetAddress())
	} else {
		senderAccAddr = senderAccAddrInterface.(string)
	}
	_, err := ck.SendCoins(ctx, sdk.AccAddress(projectAddr), sdk.AccAddress(senderAccAddr), sdk.Coins{sdk.NewInt64Coin(COIN_DENOM, projectDoc.GetEvaluatorPay())})
	if err != nil {
		return err.Result()
	}
	return sdk.Result{
		Code: sdk.ABCICodeOK,
		Data: []byte("Action complete"),
	}
}

func fundProject(ctx sdk.Context, k Keeper, ck bank.Keeper, projectDoc StoredProjectDoc, coin sdk.Coin) sdk.Result {
	accMap := getProjectAccountMap(ctx, k, projectDoc.GetProjectDid())
	projectAddrInterface, found := accMap[projectDoc.GetProjectDid()]
	if !found {
		return sdk.Result{
			Code: sdk.ABCICodeType(sdk.CodeInvalidAddress),
			Data: []byte("Could not find Project Account"),
		}
	}
	projectAddr := projectAddrInterface.(string)

	_, _, err := ck.AddCoins(ctx, sdk.AccAddress(projectAddr), sdk.Coins{coin})
	if err != nil {
		panic(err)
	}

	return sdk.Result{
		Code: sdk.ABCICodeOK,
		Data: []byte("Project Funded"),
	}
}
func handleWithdrawFundsMsg(ctx sdk.Context, k Keeper, ck bank.Keeper, msg WithdrawFundsMsg) sdk.Result {
	return sdk.Result{
		Code: sdk.ABCICodeOK,
		Data: []byte("Action complete"),
	}
}

func getProjectDoc(ctx sdk.Context, k Keeper, projectDid ixo.Did) (StoredProjectDoc, bool) {
	ixoProjectDoc, found := k.GetProjectDoc(ctx, projectDid)
	return ixoProjectDoc.(StoredProjectDoc), found
}

func getProjectAccountMap(ctx sdk.Context, k Keeper, projectDid ixo.Did) map[string]interface{} {
	return k.GetAccountMap(ctx, projectDid)
}

func addAccountToAccountProjectAccounts(ctx sdk.Context, k Keeper, projectDid ixo.Did, accountDid ixo.Did) auth.Account {
	acc := k.CreateNewAccount(ctx, projectDid, accountDid)
	k.AddAccountToAccountProjectAccounts(ctx, projectDid, accountDid, acc)

	return acc
}

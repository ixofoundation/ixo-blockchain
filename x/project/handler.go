package project

import (
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/ixofoundation/ixo-cosmos/x/fees"
	ixo "github.com/ixofoundation/ixo-cosmos/x/ixo"
)

const IxoAccountId = "IXO Foundation"
const InitiatingNodeAccountId = "InitatingNode"

func NewHandler(k Keeper, fk fees.Keeper, ck bank.Keeper, ethClient ixo.EthClient) sdk.Handler {
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
			return handleCreateClaimMsg(ctx, k, fk, ck, msg)
		case CreateEvaluationMsg:
			return handleCreateEvaluationMsg(ctx, k, fk, ck, msg)
		case WithdrawFundsMsg:
			return handleWithdrawFundsMsg(ctx, k, ck, msg)
		default:
			return sdk.ErrUnknownRequest("No match for message type.").Result()
		}
	}
}

func handleCreateProjectMsg(ctx sdk.Context, k Keeper, ck bank.Keeper, msg CreateProjectMsg) sdk.Result {
	// Create Project Account for Project
	getAccountInProjectAccounts(ctx, k, msg.GetProjectDid(), msg.GetProjectDid())
	// Create IXO Account for Project
	getAccountInProjectAccounts(ctx, k, msg.GetProjectDid(), IxoAccountId)
	// Create Initiating Node Account for Project
	getAccountInProjectAccounts(ctx, k, msg.GetProjectDid(), InitiatingNodeAccountId)

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
		return sdk.ErrUnknownRequest("Could not find Project").Result()
	}

	newStatus := msg.GetStatus()

	switch newStatus {
	case CreatedProject:
		if existingProjectDoc.GetStatus() != NullStatus {
			return sdk.ErrUnknownRequest("Invalid Status").Result()
		}
	case FundedStatus:
		res := checkFunded(ctx, k, ck, ethClient, msg, existingProjectDoc)
		if res.Code != sdk.ABCICodeOK {
			return res
		}
	case StartedStatus:
		if existingProjectDoc.GetStatus() != FundedStatus {
			return sdk.ErrUnknownRequest("Invalid Status").Result()
		}
	case StoppedStatus:
		if existingProjectDoc.GetStatus() != StartedStatus {
			return sdk.ErrUnknownRequest("Invalid Status").Result()
		}
	case PaidoutStatus:
		if existingProjectDoc.GetStatus() != StoppedStatus {
			return sdk.ErrUnknownRequest("Invalid Status").Result()
		}
	}

	existingProjectDoc.SetStatus(newStatus)
	storedProjectDoc, _ := k.UpdateProjectDoc(ctx, existingProjectDoc)
	return sdk.Result{
		Code: sdk.ABCICodeOK,
		Data: k.encodeProject(storedProjectDoc),
	}

}

func handleCreateAgentMsg(ctx sdk.Context, k Keeper, ck bank.Keeper, msg CreateAgentMsg) sdk.Result {
	getAccountInProjectAccounts(ctx, k, msg.GetProjectDid(), msg.Data.AgentDid)
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
func handleCreateClaimMsg(ctx sdk.Context, k Keeper, fk fees.Keeper, ck bank.Keeper, msg CreateClaimMsg) sdk.Result {

	res, err := processFees(ctx, k, fk, ck, fees.FeeClaimTransaction, msg.GetProjectDid())
	if err != nil {
		return err.Result()
	} else {
		return res
	}

}

func handleCreateEvaluationMsg(ctx sdk.Context, k Keeper, fk fees.Keeper, ck bank.Keeper, msg CreateEvaluationMsg) sdk.Result {
	_, err := processFees(ctx, k, fk, ck, fees.FeeEvaluationTransaction, msg.GetProjectDid())
	// Return error if there was an error processing the fees
	if err != nil {
		return err.Result()
	}

	projectDoc, found := getProjectDoc(ctx, k, msg.GetProjectDid())
	if !found {
		return sdk.ErrUnknownRequest("Could not find Project").Result()
	}

	// If there is an EvaluatorPay configured than we make the payment and deduct and pay those fees
	if projectDoc.GetEvaluatorPay() != 0 {
		projectAddr := getAccountInProjectAccounts(ctx, k, msg.GetProjectDid(), msg.GetProjectDid())
		nodeAddr := getAccountInProjectAccounts(ctx, k, msg.GetProjectDid(), InitiatingNodeAccountId)
		ixoAddr := getAccountInProjectAccounts(ctx, k, msg.GetProjectDid(), IxoAccountId)
		evaluatorAccAddr := getAccountInProjectAccounts(ctx, k, msg.GetProjectDid(), msg.GetSenderDid())

		// Get percentage of the Evaluator pay to pay in fees
		feePercentage := fk.GetRat(ctx, fees.KeyEvaluationPayFeePercentage)
		// Get percentage of the Evaluator Pay fees that goes to the node
		nodeFeePercentage := fk.GetRat(ctx, fees.KeyEvaluationPayNodeFeePercentage)

		totalEvaluatorPayAmount := sdk.NewRat(projectDoc.GetEvaluatorPay(), 1) // This is in IXO * 10^8
		// Calculate the fee due
		evaluatorPayFeeAmount := totalEvaluatorPayAmount.Mul(feePercentage).RoundInt64()
		// Calculate what the evaluator gets less the fees
		evaluatorPayLessFees := totalEvaluatorPayAmount.RoundInt64() - evaluatorPayFeeAmount
		// Calculate the percentage of the fees that goes to the node
		nodeFees := sdk.NewRat(evaluatorPayFeeAmount, 1).Mul(nodeFeePercentage).RoundInt64()
		// Calculate the remaining  ees that goes to the ixo foundation
		ixoFees := evaluatorPayFeeAmount - nodeFees

		// Pay Evaluator
		_, err := ck.SendCoins(ctx, projectAddr, evaluatorAccAddr, sdk.Coins{sdk.NewInt64Coin(ixo.IxoNativeToken, evaluatorPayLessFees)})
		if err != nil {
			return err.Result()
		}
		// Pay Node
		_, err = ck.SendCoins(ctx, projectAddr, nodeAddr, sdk.Coins{sdk.NewInt64Coin(ixo.IxoNativeToken, nodeFees)})
		if err != nil {
			return err.Result()
		}
		// Pay ixo Foundation
		_, err = ck.SendCoins(ctx, projectAddr, ixoAddr, sdk.Coins{sdk.NewInt64Coin(ixo.IxoNativeToken, ixoFees)})
		if err != nil {
			return err.Result()
		}
	}

	return sdk.Result{
		Code: sdk.ABCICodeOK,
		Data: []byte("Action complete"),
	}
}

// TODO: This function is not completed
func handleWithdrawFundsMsg(ctx sdk.Context, k Keeper, ck bank.Keeper, msg WithdrawFundsMsg) sdk.Result {
	return sdk.Result{
		Code: sdk.ABCICodeOK,
		Data: []byte("Action complete"),
	}
}

func checkFunded(ctx sdk.Context, k Keeper, ck bank.Keeper, ethClient ixo.EthClient, msg UpdateProjectStatusMsg, existingProjectDoc StoredProjectDoc) sdk.Result {
	if existingProjectDoc.GetStatus() != CreatedProject {
		return sdk.ErrUnknownRequest("Invalid Status").Result()
	} else {
		// Check that the Project wallet is funded and mint equivalent tokens on project

		ethTx, err := ethClient.GetTransactionByHash(msg.GetEthFundingTxnID())
		if err != nil {
			return sdk.ErrUnknownRequest("ETH tx not valid").Result()
		}
		fundingTx := ethClient.IsProjectFundingTx(ctx, existingProjectDoc.GetProjectDid(), ethTx)
		if !fundingTx {
			return sdk.ErrUnknownRequest("ETH tx not valid").Result()
		}
		amt := ethClient.GetFundingAmt(ethTx)
		coin := sdk.NewInt64Coin(ixo.IxoNativeToken, amt)
		return fundProject(ctx, k, ck, existingProjectDoc, coin)
	}
}

func fundProject(ctx sdk.Context, k Keeper, ck bank.Keeper, projectDoc StoredProjectDoc, coin sdk.Coin) sdk.Result {
	projectAddr := getAccountInProjectAccounts(ctx, k, projectDoc.GetProjectDid(), projectDoc.GetProjectDid())

	_, _, err := ck.AddCoins(ctx, projectAddr, sdk.Coins{coin})
	if err != nil {
		panic(err)
	}

	return sdk.Result{
		Code: sdk.ABCICodeOK,
		Data: []byte("Project Funded"),
	}
}

func getProjectDoc(ctx sdk.Context, k Keeper, projectDid ixo.Did) (StoredProjectDoc, bool) {
	ixoProjectDoc, found := k.GetProjectDoc(ctx, projectDid)
	return ixoProjectDoc.(StoredProjectDoc), found
}

func getProjectAccountMap(ctx sdk.Context, k Keeper, projectDid ixo.Did) map[string]interface{} {
	return k.GetAccountMap(ctx, projectDid)
}

func getAccountInProjectAccounts(ctx sdk.Context, k Keeper, projectDid ixo.Did, accountID string) sdk.AccAddress {
	accMap := getProjectAccountMap(ctx, k, projectDid)
	var accountIDAccAddr string
	accountIDAddrInterface, found := accMap[accountID]
	if !found {
		newAcc := k.CreateNewAccount(ctx)
		k.AddAccountToProjectAccounts(ctx, projectDid, accountID, newAcc)
		accountIDAccAddr = hex.EncodeToString(newAcc.GetAddress())
	} else {
		accountIDAccAddr = accountIDAddrInterface.(string)
	}
	return sdk.AccAddress(accountIDAccAddr)
}

func processFees(ctx sdk.Context, k Keeper, fk fees.Keeper, ck bank.Keeper, feeType fees.FeeType, projectDid ixo.Did) (sdk.Result, sdk.Error) {
	projectAddr := getAccountInProjectAccounts(ctx, k, projectDid, projectDid)
	nodeAddr := getAccountInProjectAccounts(ctx, k, projectDid, InitiatingNodeAccountId)
	ixoAddr := getAccountInProjectAccounts(ctx, k, projectDid, IxoAccountId)

	ixoFactor := fk.GetRat(ctx, fees.KeyIxoFactor)
	nodePercentage := fk.GetRat(ctx, fees.KeyNodeFeePercentage)
	var adjustedFeeAmount sdk.Rat
	switch feeType {
	case fees.FeeClaimTransaction:
		adjustedFeeAmount = fk.GetRat(ctx, fees.KeyClaimFeeAmount).Mul(ixoFactor)
	case fees.FeeEvaluationTransaction:
		adjustedFeeAmount = fk.GetRat(ctx, fees.KeyEvaluationFeeAmount).Mul(ixoFactor)
	default:
		return sdk.Result{}, sdk.ErrUnknownRequest("Invalid Fee type.")
	}

	// Get the adjusted fee amount and round to an int64
	nodeAmount := adjustedFeeAmount.Mul(nodePercentage).RoundInt64()
	// now subtract the nodeAmount from the adjustedAmount as the foundation gets the other part of the fee
	ixoAmount := adjustedFeeAmount.RoundInt64() - nodeAmount

	_, err := ck.SendCoins(ctx, projectAddr, nodeAddr, sdk.Coins{sdk.NewInt64Coin(ixo.IxoNativeToken, nodeAmount)})
	if err != nil {
		return sdk.Result{}, err
	}

	_, err = ck.SendCoins(ctx, projectAddr, ixoAddr, sdk.Coins{sdk.NewInt64Coin(ixo.IxoNativeToken, ixoAmount)})
	if err != nil {
		return sdk.Result{}, err
	}

	return sdk.Result{
		Code: sdk.ABCICodeOK,
		Data: []byte("Action complete"),
	}, nil
}

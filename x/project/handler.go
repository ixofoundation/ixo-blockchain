package project

import (
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/ixofoundation/ixo-cosmos/x/contracts"
	"github.com/ixofoundation/ixo-cosmos/x/fees"
	ixo "github.com/ixofoundation/ixo-cosmos/x/ixo"
	"github.com/ixofoundation/ixo-cosmos/x/params"
)

type InternalAccountID = string

const (
	IxoAccountFeesId               InternalAccountID = "IxoFees"
	IxoAccountPayFeesId            InternalAccountID = "IxoPayFees"
	InitiatingNodeAccountPayFeesId InternalAccountID = "InitiatingNodePayFees"
	ValidatingNodeSetAccountFeesId InternalAccountID = "ValidatingNodeSetFees"
)

func NewHandler(k Keeper, fk fees.Keeper, ck contracts.Keeper, bk bank.Keeper, pk params.Keeper, ethClient ixo.EthClient) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case CreateProjectMsg:
			return handleCreateProjectMsg(ctx, k, bk, msg)
		case UpdateProjectStatusMsg:
			return handleUpdateProjectStatusMsg(ctx, k, ck, bk, pk, ethClient, msg)
		case CreateAgentMsg:
			return handleCreateAgentMsg(ctx, k, bk, msg)
		case UpdateAgentMsg:
			return handleUpdateAgentMsg(ctx, k, bk, msg)
		case CreateClaimMsg:
			return handleCreateClaimMsg(ctx, k, fk, bk, msg)
		case CreateEvaluationMsg:
			return handleCreateEvaluationMsg(ctx, k, fk, bk, msg)
		case WithdrawFundsMsg:
			return handleWithdrawFundsMsg(ctx, k, bk, pk, ethClient, msg)
		default:
			return sdk.ErrUnknownRequest("No match for message type.").Result()
		}
	}
}

func handleCreateProjectMsg(ctx sdk.Context, k Keeper, bk bank.Keeper, msg CreateProjectMsg) sdk.Result {

	// Create Project Account for Project
	createAccountInProjectAccounts(ctx, k, msg.GetProjectDid(), IxoAccountFeesId)
	createAccountInProjectAccounts(ctx, k, msg.GetProjectDid(), msg.GetProjectDid())

	projectDoc, err := k.AddProjectDoc(ctx, &msg)
	if err != nil {
		return err.Result()
	}
	return sdk.Result{
		Code: sdk.ABCICodeOK,
		Data: k.encodeProject(projectDoc),
	}
}

func handleUpdateProjectStatusMsg(ctx sdk.Context, k Keeper, ck contracts.Keeper, bk bank.Keeper, pk params.Keeper, ethClient ixo.EthClient, msg UpdateProjectStatusMsg) sdk.Result {
	existingProjectDoc, found := getProjectDoc(ctx, k, msg.GetProjectDid())
	if !found {
		return sdk.ErrUnknownRequest("Could not find Project").Result()
	}

	newStatus := msg.GetStatus()
	if !newStatus.IsValidProgressionFrom(existingProjectDoc.GetStatus()) {
		return sdk.ErrUnknownRequest("Invalid Status Progression requested").Result()
	}

	if newStatus == FundedStatus {
		ethFundingTxnID := msg.GetEthFundingTxnID()
		ctx.Logger().Info("Provided ethFundingTxnID: ", ethFundingTxnID)
		if ethFundingTxnID == "" {
			ctx.Logger().Error("ETH tx not valid isFundingTx")
			return sdk.ErrUnknownRequest("Invalid EthFundingTxnID provided").Result()
		}

		res := fundIfLegitimateEthereumTx(ctx, k, bk, ethClient, ethFundingTxnID, existingProjectDoc)
		if res.Code != sdk.ABCICodeOK {
			return res
		}
	}

	if newStatus == PaidoutStatus {
		res := payoutFees(ctx, k, ck, bk, pk, ethClient, existingProjectDoc.GetProjectDid())
		if res.Code != sdk.ABCICodeOK {
			return res
		}
	}

	existingProjectDoc.SetStatus(newStatus)
	storedProjectDoc, _ := k.UpdateProjectDoc(ctx, existingProjectDoc)
	return sdk.Result{
		Code: sdk.ABCICodeOK,
		Data: k.encodeProject(storedProjectDoc),
	}

}

func payoutFees(ctx sdk.Context, k Keeper, ck contracts.Keeper, bk bank.Keeper, pk params.Keeper, ethClient ixo.EthClient, projectDid ixo.Did) sdk.Result {

	// initiate auth contract based ixo ERC20 token transfer on Ethereum
	_, err := ethClient.ProjectWalletFromProjectRegistry(ctx, projectDid)
	if err != nil {
		return sdk.ErrUnknownRequest("Could not find Project Ethereum wallet").Result()
	}

	// for now all fees go to the ixoWallet
	_, err = payAllFeesToAddress(ctx, k, bk, projectDid, IxoAccountPayFeesId, IxoAccountFeesId)
	if err != nil {
		return sdk.ErrInternal("Failed to send coins").Result()
	}
	_, err = payAllFeesToAddress(ctx, k, bk, projectDid, InitiatingNodeAccountPayFeesId, IxoAccountFeesId)
	if err != nil {
		return sdk.ErrInternal("Failed to send coins").Result()
	}
	_, err = payAllFeesToAddress(ctx, k, bk, projectDid, ValidatingNodeSetAccountFeesId, IxoAccountFeesId)
	if err != nil {
		return sdk.ErrInternal("Failed to send coins").Result()
	}

	ixoEthWallet := ck.GetContract(ctx, contracts.KeyFoundationWallet)
	return payoutERC20AndRecon(ctx, k, bk, pk, ethClient, projectDid, IxoAccountFeesId, ixoEthWallet)
}

func payAllFeesToAddress(ctx sdk.Context, k Keeper, bk bank.Keeper, projectDid ixo.Did, sendingAddress InternalAccountID, receivingAddress InternalAccountID) (sdk.Tags, sdk.Error) {
	feesToPay := getIxoAmount(ctx, k, bk, projectDid, sendingAddress)

	if feesToPay < 0 {
		return nil, sdk.ErrInternal("Negative fee to pay")
	}

	if feesToPay == 0 {
		return nil, nil
	}

	receivingAccount := getAccountInProjectAccounts(ctx, k, projectDid, receivingAddress)
	sendingAccount := getAccountInProjectAccounts(ctx, k, projectDid, sendingAddress)
	return bk.SendCoins(ctx, sendingAccount, receivingAccount, sdk.Coins{sdk.NewInt64Coin(ixo.IxoNativeToken, feesToPay)})
}

func getIxoAmount(ctx sdk.Context, k Keeper, bk bank.Keeper, projectDid ixo.Did, accountID string) int64 {
	found := checkAccountInProjectAccounts(ctx, k, projectDid, accountID)
	if found {
		accAddr := getAccountInProjectAccounts(ctx, k, projectDid, accountID)
		coins := bk.GetCoins(ctx, accAddr)
		return coins.AmountOf(ixo.IxoNativeToken).Int64()
	}
	return 0
}

func handleCreateAgentMsg(ctx sdk.Context, k Keeper, bk bank.Keeper, msg CreateAgentMsg) sdk.Result {
	createAccountInProjectAccounts(ctx, k, msg.GetProjectDid(), msg.Data.AgentDid)
	return sdk.Result{
		Code: sdk.ABCICodeOK,
		Data: []byte("Action complete"),
	}
}

func handleUpdateAgentMsg(ctx sdk.Context, k Keeper, bk bank.Keeper, msg UpdateAgentMsg) sdk.Result {
	return sdk.Result{
		Code: sdk.ABCICodeOK,
		Data: []byte("Action complete"),
	}
}
func handleCreateClaimMsg(ctx sdk.Context, k Keeper, fk fees.Keeper, bk bank.Keeper, msg CreateClaimMsg) sdk.Result {
	_, err := processFees(ctx, k, fk, bk, fees.FeeClaimTransaction, msg.GetProjectDid())
	if err != nil {
		fmt.Println("Error:", err)
		return err.Result()
	} else {
		return sdk.Result{
			Code: sdk.ABCICodeOK,
			Data: []byte("Action complete"),
		}
	}

}

func handleCreateEvaluationMsg(ctx sdk.Context, k Keeper, fk fees.Keeper, bk bank.Keeper, msg CreateEvaluationMsg) sdk.Result {
	_, err := processFees(ctx, k, fk, bk, fees.FeeEvaluationTransaction, msg.GetProjectDid())
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
		projectDid := msg.GetProjectDid()
		projectAddr := getAccountInProjectAccounts(ctx, k, projectDid, msg.GetProjectDid())
		evaluatorAccAddr := getAccountInProjectAccounts(ctx, k, projectDid, msg.GetSenderDid())

		found := checkAccountInProjectAccounts(ctx, k, projectDid, InitiatingNodeAccountPayFeesId)
		var nodeAddr sdk.AccAddress
		if !found {
			nodeAddr = createAccountInProjectAccounts(ctx, k, projectDid, InitiatingNodeAccountPayFeesId)
		} else {
			nodeAddr = getAccountInProjectAccounts(ctx, k, msg.GetProjectDid(), InitiatingNodeAccountPayFeesId)
		}

		found = checkAccountInProjectAccounts(ctx, k, projectDid, InitiatingNodeAccountPayFeesId)
		var ixoAddr sdk.AccAddress
		if !found {
			ixoAddr = createAccountInProjectAccounts(ctx, k, msg.GetProjectDid(), IxoAccountPayFeesId)
		} else {
			ixoAddr = getAccountInProjectAccounts(ctx, k, msg.GetProjectDid(), IxoAccountPayFeesId)
		}

		// Get percentage of the Evaluator pay to pay in fees
		feePercentage := fk.GetRat(ctx, fees.KeyEvaluationPayFeePercentage)
		// Get percentage of the Evaluator Pay fees that goes to the node
		nodeFeePercentage := fk.GetRat(ctx, fees.KeyEvaluationPayNodeFeePercentage)

		totalEvaluatorPayAmount := sdk.NewRat(projectDoc.GetEvaluatorPay(), 1).Mul(ixo.IxoDecimals) // This is in IXO * 10^8
		// Calculate the fee due
		evaluatorPayFeeAmount := totalEvaluatorPayAmount.Mul(feePercentage).RoundInt64()
		// Calculate what the evaluator gets less the fees
		evaluatorPayLessFees := totalEvaluatorPayAmount.RoundInt64() - evaluatorPayFeeAmount
		// Calculate the percentage of the fees that goes to the node
		nodePayFees := sdk.NewRat(evaluatorPayFeeAmount, 1).Mul(nodeFeePercentage).RoundInt64()
		// Calculate the remaining  ees that goes to the ixo foundation
		ixoPayFees := evaluatorPayFeeAmount - nodePayFees

		// Pay Evaluator
		_, err := bk.SendCoins(ctx, projectAddr, evaluatorAccAddr, sdk.Coins{sdk.NewInt64Coin(ixo.IxoNativeToken, evaluatorPayLessFees)})
		if err != nil {
			return err.Result()
		}
		// Pay Node
		_, err = bk.SendCoins(ctx, projectAddr, nodeAddr, sdk.Coins{sdk.NewInt64Coin(ixo.IxoNativeToken, nodePayFees)})
		if err != nil {
			return err.Result()
		}
		// Pay ixo Foundation
		_, err = bk.SendCoins(ctx, projectAddr, ixoAddr, sdk.Coins{sdk.NewInt64Coin(ixo.IxoNativeToken, ixoPayFees)})
		if err != nil {
			return err.Result()
		}
	}

	return sdk.Result{
		Code: sdk.ABCICodeOK,
		Data: []byte("Action complete"),
	}
}

func handleWithdrawFundsMsg(ctx sdk.Context, k Keeper, bk bank.Keeper, pk params.Keeper, ethClient ixo.EthClient, msg WithdrawFundsMsg) sdk.Result {
	withdrawFundsDoc := msg.GetWithdrawFundsDoc()
	projectDoc, found := getProjectDoc(ctx, k, withdrawFundsDoc.GetProjectDid())
	if !found {
		return sdk.ErrUnknownRequest("Could not find Project").Result()
	}

	if projectDoc.GetStatus() != PaidoutStatus {
		return sdk.ErrUnknownRequest("Project not in PAIDOUT Status").Result()
	}

	ethWalletAddress := withdrawFundsDoc.GetEthWallet()
	projectDid := withdrawFundsDoc.GetProjectDid()
	var payoutResult sdk.Result
	if withdrawFundsDoc.IsRefund {
		payoutResult = payoutERC20AndRecon(ctx, k, bk, pk, ethClient, projectDid, projectDid, ethWalletAddress)
	} else {
		senderDid := msg.GetSenderDid()
		payoutResult = payoutERC20AndRecon(ctx, k, bk, pk, ethClient, projectDid, senderDid, ethWalletAddress)
	}

	return payoutResult
}

func payoutERC20AndRecon(ctx sdk.Context, k Keeper, bk bank.Keeper, pk params.Keeper, ethClient ixo.EthClient, projectDid ixo.Did, accountID string, recipientEthAddress string) sdk.Result {
	balanceToPay := getIxoAmount(ctx, k, bk, projectDid, accountID)
	if balanceToPay <= 0 {
		return sdk.ErrUnknownRequest("No balance to pay out on Project").Result()
	}

	// initiate auth contract based ixo ERC20 token transfer on Ethereum
	projectEthWallet, err := ethClient.ProjectWalletFromProjectRegistry(ctx, projectDid)
	if err != nil {
		return sdk.ErrUnknownRequest("Could not find Project Ethereum wallet").Result()
	}

	account := getAccountInProjectAccounts(ctx, k, projectDid, accountID)
	// burn coins
	_, _, err = bk.SubtractCoins(ctx, account, sdk.Coins{sdk.NewInt64Coin(ixo.IxoNativeToken, balanceToPay)})
	if err != nil {
		return sdk.ErrUnknownRequest("Could not burn tokens from " + account.String()).Result()
	}

	_, actionID := ethClient.InitiateTokenTransfer(ctx, pk, projectEthWallet, recipientEthAddress, balanceToPay)

	addProjectWithdrawalTransaction(ctx, k, projectDid, actionID, projectEthWallet, recipientEthAddress, balanceToPay)

	return sdk.Result{
		Code: sdk.ABCICodeOK,
		Data: []byte("Action complete"),
	}
}

func fundIfLegitimateEthereumTx(ctx sdk.Context, k Keeper, bk bank.Keeper, ethClient ixo.EthClient, ethFundingTxnID string, existingProjectDoc StoredProjectDoc) sdk.Result {
	// Check that the Project wallet is funded and mint equivalent tokens on project

	ethTx, err := ethClient.GetTransactionByHash(ethFundingTxnID)
	if err != nil {
		return sdk.ErrUnknownRequest("ETH tx not valid: Could not get transaction: " + ethFundingTxnID).Result()
	}
	isFundingTx := ethClient.IsProjectFundingTx(ctx, existingProjectDoc.GetProjectDid(), ethTx)
	if !isFundingTx {
		return sdk.ErrUnknownRequest("ETH tx not valid. Not a valid project funding transaction").Result()
	}

	amt := ethClient.GetFundingAmt(ethTx)
	fmt.Println("PROJECT_FUNDING", "amt: ", amt)
	coin := sdk.NewInt64Coin(ixo.IxoNativeToken, amt)
	return fundProject(ctx, k, bk, existingProjectDoc, coin)
}

func fundProject(ctx sdk.Context, k Keeper, bk bank.Keeper, projectDoc StoredProjectDoc, coin sdk.Coin) sdk.Result {
	fmt.Printf("PROJECT_FUNDING func fundProject(_, _, _, _, [coin.Amount: %d, coin.Denom: %s])", coin.Amount.Int64(), coin.Denom)
	projectAddr := getAccountInProjectAccounts(ctx, k, projectDoc.GetProjectDid(), projectDoc.GetProjectDid())

	_, _, err := bk.AddCoins(ctx, projectAddr, sdk.Coins{coin})
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

func processFees(ctx sdk.Context, k Keeper, fk fees.Keeper, bk bank.Keeper, feeType fees.FeeType, projectDid ixo.Did) (sdk.Result, sdk.Error) {
	projectAddr := getAccountInProjectAccounts(ctx, k, projectDid, projectDid)
	var validatingNodeSetAddr sdk.AccAddress
	found := checkAccountInProjectAccounts(ctx, k, projectDid, ValidatingNodeSetAccountFeesId)
	if !found {
		validatingNodeSetAddr = createAccountInProjectAccounts(ctx, k, projectDid, ValidatingNodeSetAccountFeesId)
	} else {
		validatingNodeSetAddr = getAccountInProjectAccounts(ctx, k, projectDid, ValidatingNodeSetAccountFeesId)
	}

	var ixoAddr sdk.AccAddress
	found = checkAccountInProjectAccounts(ctx, k, projectDid, IxoAccountFeesId)
	if !found {
		ixoAddr = createAccountInProjectAccounts(ctx, k, projectDid, IxoAccountFeesId)
	} else {
		ixoAddr = getAccountInProjectAccounts(ctx, k, projectDid, IxoAccountFeesId)
	}

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

	_, err := bk.SendCoins(ctx, projectAddr, validatingNodeSetAddr, sdk.Coins{sdk.NewInt64Coin(ixo.IxoNativeToken, nodeAmount)})
	if err != nil {
		return sdk.Result{}, err
	}

	_, err = bk.SendCoins(ctx, projectAddr, ixoAddr, sdk.Coins{sdk.NewInt64Coin(ixo.IxoNativeToken, ixoAmount)})
	if err != nil {
		return sdk.Result{}, err
	}

	return sdk.Result{
		Code: sdk.ABCICodeOK,
		Data: []byte("Action complete"),
	}, nil
}

func checkAccountInProjectAccounts(ctx sdk.Context, k Keeper, projectDid ixo.Did, accountId string) bool {
	accMap := k.GetAccountMap(ctx, projectDid)
	_, found := accMap[accountId]
	return found
}

func addProjectWithdrawalTransaction(ctx sdk.Context, k Keeper, projectDid ixo.Did, actionID [32]byte, projectEthWallet string, recipientEthAddress string, amount int64) {
	actionIDStr := "0x" + hex.EncodeToString(actionID[:])
	withdrawalInfo := WithdrawalInfo{
		actionIDStr,
		projectEthWallet,
		recipientEthAddress,
		amount,
	}
	k.AddProjectWithdrawalTransaction(ctx, projectDid, withdrawalInfo)
}

func createAccountInProjectAccounts(ctx sdk.Context, k Keeper, projectDid ixo.Did, accountId string) sdk.AccAddress {
	acc := k.CreateNewAccount(ctx, projectDid, accountId)
	k.AddAccountToProjectAccounts(ctx, projectDid, accountId, acc)
	return acc.GetAddress()
}

func getProjectAccountMap(ctx sdk.Context, k Keeper, projectDid ixo.Did) map[string]interface{} {
	return k.GetAccountMap(ctx, projectDid)
}

func getAccountInProjectAccounts(ctx sdk.Context, k Keeper, projectDid ixo.Did, accountId string) sdk.AccAddress {
	accMap := getProjectAccountMap(ctx, k, projectDid)
	var accountIDAccAddr string
	accountIDAddrInterface, found := accMap[accountId]
	if found {
		accountIDAccAddr = accountIDAddrInterface.(string)
		addr := sdk.AccAddress([]byte(accountIDAccAddr))
		return addr
	} else {
		return createAccountInProjectAccounts(ctx, k, projectDid, accountId)
	}

}

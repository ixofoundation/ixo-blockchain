package project

import (
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"

	"github.com/ixofoundation/ixo-cosmos/x/contracts"
	"github.com/ixofoundation/ixo-cosmos/x/fees"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
	"github.com/ixofoundation/ixo-cosmos/x/params"
)

type InternalAccountID = string

const (
	IxoAccountFeesId               InternalAccountID = "IxoFees"
	IxoAccountPayFeesId            InternalAccountID = "IxoPayFees"
	InitiatingNodeAccountPayFeesId InternalAccountID = "InitiatingNodePayFees"
	ValidatingNodeSetAccountFeesId InternalAccountID = "ValidatingNodeSetFees"
)

func NewHandler(k Keeper, fk fees.Keeper, ck contracts.Keeper, bk bank.Keeper, pk params.Keeper,
	ethClient ixo.EthClient) sdk.Handler {

	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgCreateProject:
			return handleMsgCreateProject(ctx, k, bk, msg)
		case MsgUpdateProjectStatus:
			return handleMsgUpdateProjectStatus(ctx, k, ck, bk, pk, ethClient, msg)
		case MsgCreateAgent:
			return handleMsgCreateAgent(ctx, k, bk, msg)
		case MsgUpdateAgent:
			return handleMsgUpdateAgent(ctx, k, bk, msg)
		case MsgCreateClaim:
			return handleMsgCreateClaim(ctx, k, fk, bk, msg)
		case MsgCreateEvaluation:
			return handleMsgCreateEvaluation(ctx, k, fk, bk, msg)
		case MsgWithdrawFunds:
			return handleMsgWithdrawFunds(ctx, k, bk, pk, ethClient, msg)
		default:
			return sdk.ErrUnknownRequest("No match for message type.").Result()
		}
	}
}

func handleMsgCreateProject(ctx sdk.Context, k Keeper, bk bank.Keeper, msg MsgCreateProject) sdk.Result {

	_, err := createAccountInProjectAccounts(ctx, k, msg.GetProjectDid(), IxoAccountFeesId)
	if err != nil {
		return err.Result()
	}

	_, err = createAccountInProjectAccounts(ctx, k, msg.GetProjectDid(), msg.GetProjectDid())
	if err != nil {
		err.Result()
	}

	err = k.SetProjectDoc(ctx, &msg)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Code: sdk.CodeOK,
	}
}

func handleMsgUpdateProjectStatus(ctx sdk.Context, k Keeper, ck contracts.Keeper, bk bank.Keeper, pk params.Keeper,
	ethClient ixo.EthClient, msg MsgUpdateProjectStatus) sdk.Result {

	ExistingProjectDoc, err := getProjectDoc(ctx, k, msg.GetProjectDid())
	if err != nil {
		return sdk.ErrUnknownRequest("Could not find Project").Result()
	}

	newStatus := msg.GetStatus()
	if !newStatus.IsValidProgressionFrom(ExistingProjectDoc.GetStatus()) {
		return sdk.ErrUnknownRequest("Invalid Status Progression requested").Result()
	}

	if newStatus == FundedStatus {
		ethFundingTxnID := msg.GetEthFundingTxnID()
		ctx.Logger().Info("Provided ethFundingTxnID: ", ethFundingTxnID)
		if ethFundingTxnID == "" {
			ctx.Logger().Error("ETH tx not valid isFundingTx")

			return sdk.ErrUnknownRequest("Invalid EthFundingTxnID provided").Result()
		}

		res := fundIfLegitimateEthereumTx(ctx, k, bk, ethClient, ethFundingTxnID, ExistingProjectDoc)
		if res.Code != sdk.CodeOK {
			return res
		}
	}

	if newStatus == PaidoutStatus {
		res := payoutFees(ctx, k, ck, bk, pk, ethClient, ExistingProjectDoc.GetProjectDid())
		if res.Code != sdk.CodeOK {
			return res
		}
	}

	ExistingProjectDoc.SetStatus(newStatus)
	_, _ = k.UpdateProjectDoc(ctx, ExistingProjectDoc)

	return sdk.Result{
		Code: sdk.CodeOK,
	}
}

func payoutFees(ctx sdk.Context, k Keeper, ck contracts.Keeper, bk bank.Keeper, pk params.Keeper,
	ethClient ixo.EthClient, projectDid ixo.Did) sdk.Result {

	_, err := ethClient.ProjectWalletFromProjectRegistry(ctx, projectDid)
	if err != nil {
		return sdk.ErrUnknownRequest("Could not find Project Ethereum wallet").Result()
	}

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

func payAllFeesToAddress(ctx sdk.Context, k Keeper, bk bank.Keeper, projectDid ixo.Did,
	sendingAddress InternalAccountID, receivingAddress InternalAccountID) (sdk.Events, sdk.Error) {
	feesToPay := getIxoAmount(ctx, k, bk, projectDid, sendingAddress)

	if feesToPay < 0 {
		return nil, sdk.ErrInternal("Negative fee to pay")
	}

	if feesToPay == 0 {
		return nil, nil
	}

	receivingAccount, err := getAccountInProjectAccounts(ctx, k, projectDid, receivingAddress)
	if err != nil {
		return sdk.Events{}, err
	}

	sendingAccount, _ := getAccountInProjectAccounts(ctx, k, projectDid, sendingAddress)

	return sdk.Events{}, bk.SendCoins(ctx, sendingAccount, receivingAccount,
		sdk.Coins{sdk.NewInt64Coin(ixo.IxoNativeToken, feesToPay)})
}

func getIxoAmount(ctx sdk.Context, k Keeper, bk bank.Keeper, projectDid ixo.Did, accountID string) int64 {
	found := checkAccountInProjectAccounts(ctx, k, projectDid, accountID)
	if found {
		accAddr, _ := getAccountInProjectAccounts(ctx, k, projectDid, accountID)
		coins := bk.GetCoins(ctx, accAddr)
		return coins.AmountOf(ixo.IxoNativeToken).Int64()
	}

	return 0
}

func handleMsgCreateAgent(ctx sdk.Context, k Keeper, bk bank.Keeper, msg MsgCreateAgent) sdk.Result {
	_, err := createAccountInProjectAccounts(ctx, k, msg.GetProjectDid(), msg.Data.AgentDid)
	if err != nil {
		err.Result()
	}

	return sdk.Result{
		Code: sdk.CodeOK,
	}
}

func handleMsgUpdateAgent(ctx sdk.Context, k Keeper, bk bank.Keeper, msg MsgUpdateAgent) sdk.Result {

	return sdk.Result{
		Code: sdk.CodeOK,
	}
}

func handleMsgCreateClaim(ctx sdk.Context, k Keeper, fk fees.Keeper, bk bank.Keeper, msg MsgCreateClaim) sdk.Result {

	_, err := processFees(ctx, k, fk, bk, fees.FeeClaimTransaction, msg.GetProjectDid())
	if err != nil {

		return err.Result()
	} else {

		return sdk.Result{
			Code: sdk.CodeOK,
		}
	}
}

func handleMsgCreateEvaluation(ctx sdk.Context, k Keeper, fk fees.Keeper, bk bank.Keeper, msg MsgCreateEvaluation) sdk.Result {
	_, err := processFees(ctx, k, fk, bk, fees.FeeEvaluationTransaction, msg.GetProjectDid())
	if err != nil {
		return err.Result()
	}

	projectDoc, err := getProjectDoc(ctx, k, msg.GetProjectDid())
	if err != nil {
		return sdk.ErrUnknownRequest("Could not find Project").Result()
	}

	if projectDoc.GetEvaluatorPay() != 0 {
		projectDid := msg.GetProjectDid()
		projectAddr, _ := getAccountInProjectAccounts(ctx, k, projectDid, msg.GetProjectDid())
		evaluatorAccAddr, _ := getAccountInProjectAccounts(ctx, k, projectDid, msg.GetSenderDid())

		found := checkAccountInProjectAccounts(ctx, k, projectDid, InitiatingNodeAccountPayFeesId)
		var nodeAddr sdk.AccAddress
		if !found {
			nodeAddr, _ = createAccountInProjectAccounts(ctx, k, projectDid, InitiatingNodeAccountPayFeesId)
		} else {
			nodeAddr, _ = getAccountInProjectAccounts(ctx, k, msg.GetProjectDid(), InitiatingNodeAccountPayFeesId)
		}

		found = checkAccountInProjectAccounts(ctx, k, projectDid, InitiatingNodeAccountPayFeesId)
		var ixoAddr sdk.AccAddress
		if !found {
			ixoAddr, _ = createAccountInProjectAccounts(ctx, k, msg.GetProjectDid(), IxoAccountPayFeesId)
		} else {
			ixoAddr, _ = getAccountInProjectAccounts(ctx, k, msg.GetProjectDid(), IxoAccountPayFeesId)
		}

		feePercentage := fk.GetDec(ctx, fees.KeyEvaluationPayFeePercentage)
		nodeFeePercentage := fk.GetDec(ctx, fees.KeyEvaluationPayNodeFeePercentage)

		totalEvaluatorPayAmount := sdk.NewDec(projectDoc.GetEvaluatorPay()).Mul(ixo.IxoDecimals) // This is in IXO * 10^8
		evaluatorPayFeeAmount := totalEvaluatorPayAmount.Mul(feePercentage)
		evaluatorPayLessFees := totalEvaluatorPayAmount.Sub(evaluatorPayFeeAmount)
		nodePayFees := evaluatorPayFeeAmount.Mul(nodeFeePercentage)
		ixoPayFees := evaluatorPayFeeAmount.Sub(nodePayFees)

		err := bk.SendCoins(ctx, projectAddr, evaluatorAccAddr, sdk.Coins{sdk.NewInt64Coin(ixo.IxoNativeToken, evaluatorPayLessFees.RoundInt64())})
		if err != nil {
			return err.Result()
		}

		err = bk.SendCoins(ctx, projectAddr, nodeAddr, sdk.Coins{sdk.NewInt64Coin(ixo.IxoNativeToken, nodePayFees.RoundInt64())})
		if err != nil {
			return err.Result()
		}

		err = bk.SendCoins(ctx, projectAddr, ixoAddr, sdk.Coins{sdk.NewInt64Coin(ixo.IxoNativeToken, ixoPayFees.RoundInt64())})
		if err != nil {
			return err.Result()
		}
	}

	return sdk.Result{
		Code: sdk.CodeOK,
	}
}

func handleMsgWithdrawFunds(ctx sdk.Context, k Keeper, bk bank.Keeper, pk params.Keeper,
	ethClient ixo.EthClient, msg MsgWithdrawFunds) sdk.Result {

	withdrawFundsDoc := msg.GetWithdrawFundsDoc()
	projectDoc, err := getProjectDoc(ctx, k, withdrawFundsDoc.GetProjectDid())
	if err != nil {
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

func payoutERC20AndRecon(ctx sdk.Context, k Keeper, bk bank.Keeper, pk params.Keeper, ethClient ixo.EthClient,
	projectDid ixo.Did, accountID string, recipientEthAddress string) sdk.Result {

	balanceToPay := getIxoAmount(ctx, k, bk, projectDid, accountID)
	if balanceToPay > 0 {
		projectEthWallet, err := ethClient.ProjectWalletFromProjectRegistry(ctx, projectDid)
		if err != nil {
			return sdk.ErrUnknownRequest("Could not find Project Ethereum wallet").Result()
		}

		account, errRes := getAccountInProjectAccounts(ctx, k, projectDid, accountID)
		if errRes != nil {
			return errRes.Result()
		}

		// TODO: Why is balanceToPay is added to account and removed right after??
		_, err = bk.AddCoins(ctx, account, sdk.Coins{sdk.NewInt64Coin(ixo.IxoNativeToken, balanceToPay)})
		if err != nil {
		}

		_, err = bk.SubtractCoins(ctx, account, sdk.Coins{sdk.NewInt64Coin(ixo.IxoNativeToken, balanceToPay)})
		if err != nil {
			return sdk.ErrUnknownRequest("Could not burn tokens from " + account.String()).Result()
		}

		_, actionID := ethClient.InitiateTokenTransfer(ctx, pk, projectEthWallet, recipientEthAddress, balanceToPay)

		addProjectWithdrawalTransaction(ctx, k, projectDid, actionID, projectEthWallet, recipientEthAddress, balanceToPay)
	}

	return sdk.Result{
		Code: sdk.CodeOK,
	}
}

func fundIfLegitimateEthereumTx(ctx sdk.Context, k Keeper, bk bank.Keeper, ethClient ixo.EthClient,
	ethFundingTxnID string, ExistingProjectDoc StoredProjectDoc) sdk.Result {

	ethTx, err := ethClient.GetTransactionByHash(ethFundingTxnID)
	if err != nil {
		return sdk.ErrUnknownRequest("ETH tx not valid: Could not get transaction: " + ethFundingTxnID).Result()
	}

	isFundingTx := ethClient.IsProjectFundingTx(ctx, ExistingProjectDoc.GetProjectDid(), ethTx)
	if !isFundingTx {
		return sdk.ErrUnknownRequest("ETH tx not valid. Not a valid project funding transaction").Result()
	}

	amt := ethClient.GetFundingAmt(ethTx)
	fmt.Println("PROJECT_FUNDING", "amt: ", amt)
	coin := sdk.NewInt64Coin(ixo.IxoNativeToken, amt)

	return fundProject(ctx, k, bk, ExistingProjectDoc, coin)
}

func fundProject(ctx sdk.Context, k Keeper, bk bank.Keeper, projectDoc StoredProjectDoc, coin sdk.Coin) sdk.Result {
	fmt.Printf("PROJECT_FUNDING func fundProject(_, _, _, _, [coin.Amount: %d, coin.Denom: %s])",
		coin.Amount.Int64(), coin.Denom)
	projectAddr, errRes := getAccountInProjectAccounts(ctx, k, projectDoc.GetProjectDid(), projectDoc.GetProjectDid())
	if errRes != nil {
		return errRes.Result()
	}

	_, err := bk.AddCoins(ctx, projectAddr, sdk.Coins{coin})
	if err != nil {
		panic(err)
	}

	return sdk.Result{
		Code: sdk.CodeOK,
	}
}

func getProjectDoc(ctx sdk.Context, k Keeper, projectDid ixo.Did) (StoredProjectDoc, sdk.Error) {
	ixoProjectDoc, err := k.GetProjectDoc(ctx, projectDid)

	return ixoProjectDoc.(StoredProjectDoc), err
}

func processFees(ctx sdk.Context, k Keeper, fk fees.Keeper, bk bank.Keeper, feeType fees.FeeType, projectDid ixo.Did) (sdk.Result, sdk.Error) {

	projectAddr, _ := getAccountInProjectAccounts(ctx, k, projectDid, projectDid)
	var validatingNodeSetAddr sdk.AccAddress

	found := checkAccountInProjectAccounts(ctx, k, projectDid, ValidatingNodeSetAccountFeesId)
	if !found {
		validatingNodeSetAddr, _ = createAccountInProjectAccounts(ctx, k, projectDid, ValidatingNodeSetAccountFeesId)
	} else {
		validatingNodeSetAddr, _ = getAccountInProjectAccounts(ctx, k, projectDid, ValidatingNodeSetAccountFeesId)
	}

	var ixoAddr sdk.AccAddress
	found = checkAccountInProjectAccounts(ctx, k, projectDid, IxoAccountFeesId)
	if !found {
		ixoAddr, _ = createAccountInProjectAccounts(ctx, k, projectDid, IxoAccountFeesId)
	} else {
		ixoAddr, _ = getAccountInProjectAccounts(ctx, k, projectDid, IxoAccountFeesId)
	}

	ixoFactor := fk.GetDec(ctx, fees.KeyIxoFactor)
	nodePercentage := fk.GetDec(ctx, fees.KeyNodeFeePercentage)
	var adjustedFeeAmount sdk.Dec

	switch feeType {
	case fees.FeeClaimTransaction:
		adjustedFeeAmount = fk.GetDec(ctx, fees.KeyClaimFeeAmount).Mul(ixoFactor)
	case fees.FeeEvaluationTransaction:
		adjustedFeeAmount = fk.GetDec(ctx, fees.KeyEvaluationFeeAmount).Mul(ixoFactor)
	default:
		return sdk.Result{}, sdk.ErrUnknownRequest("Invalid Fee type.")
	}

	nodeAmount := adjustedFeeAmount.Mul(nodePercentage).RoundInt64()
	ixoAmount := adjustedFeeAmount.RoundInt64() - nodeAmount

	err := bk.SendCoins(ctx, projectAddr, validatingNodeSetAddr, sdk.Coins{sdk.NewInt64Coin(ixo.IxoNativeToken, nodeAmount)})
	if err != nil {
		return sdk.Result{}, err
	}

	err = bk.SendCoins(ctx, projectAddr, ixoAddr, sdk.Coins{sdk.NewInt64Coin(ixo.IxoNativeToken, ixoAmount)})
	if err != nil {
		return sdk.Result{}, err
	}

	return sdk.Result{
		Code: sdk.CodeOK,
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

func createAccountInProjectAccounts(ctx sdk.Context, k Keeper, projectDid ixo.Did, accountId string) (sdk.AccAddress, sdk.Error) {
	acc, err := k.CreateNewAccount(ctx, projectDid, accountId)
	if err != nil {
		return nil, err
	}

	k.AddAccountToProjectAccounts(ctx, projectDid, accountId, acc)

	return acc.GetAddress(), nil
}

func getProjectAccountMap(ctx sdk.Context, k Keeper, projectDid ixo.Did) AccountMap {
	return k.GetAccountMap(ctx, projectDid)
}

func getAccountInProjectAccounts(ctx sdk.Context, k Keeper, projectDid ixo.Did, accountId string) (sdk.AccAddress, sdk.Error) {
	accMap := getProjectAccountMap(ctx, k, projectDid)
	var accountIDAccAddr string

	accountIDAddrInterface, found := accMap[accountId]
	if found {
		accountIDAccAddr = accountIDAddrInterface.(string)
		addr := sdk.AccAddress([]byte(accountIDAccAddr))
		return addr, nil
	} else {
		return createAccountInProjectAccounts(ctx, k, projectDid, accountId)
	}
}

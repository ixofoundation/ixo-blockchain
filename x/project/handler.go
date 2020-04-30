package project

import (
	"encoding/hex"
	"fmt"
	"github.com/ixofoundation/ixo-cosmos/x/did"
	"github.com/ixofoundation/ixo-cosmos/x/project/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"

	"github.com/ixofoundation/ixo-cosmos/x/fees"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

const (
	IxoAccountFeesId               InternalAccountID = "IxoFees"
	IxoAccountPayFeesId            InternalAccountID = "IxoPayFees"
	InitiatingNodeAccountPayFeesId InternalAccountID = "InitiatingNodePayFees"
	ValidatingNodeSetAccountFeesId InternalAccountID = "ValidatingNodeSetFees"
)

func NewHandler(k Keeper, fk fees.Keeper, bk bank.Keeper) sdk.Handler {

	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgCreateProject:
			return handleMsgCreateProject(ctx, k, msg)
		case MsgUpdateProjectStatus:
			return handleMsgUpdateProjectStatus(ctx, k, bk, msg)
		case MsgCreateAgent:
			return handleMsgCreateAgent(ctx, k, bk, msg)
		case MsgUpdateAgent:
			return handleMsgUpdateAgent(ctx, k, bk, msg)
		case MsgCreateClaim:
			return handleMsgCreateClaim(ctx, k, fk, bk, msg)
		case MsgCreateEvaluation:
			return handleMsgCreateEvaluation(ctx, k, fk, bk, msg)
		case MsgWithdrawFunds:
			return handleMsgWithdrawFunds(ctx, k, bk, msg)
		default:
			return sdk.ErrUnknownRequest("No match for message type.").Result()
		}
	}
}

func handleMsgCreateProject(ctx sdk.Context, k Keeper, msg MsgCreateProject) sdk.Result {

	_, err := createAccountInProjectAccounts(ctx, k, msg.GetProjectDid(), IxoAccountFeesId)
	if err != nil {
		return err.Result()
	}

	_, err = createAccountInProjectAccounts(ctx, k, msg.GetProjectDid(), InternalAccountID(msg.GetProjectDid()))
	if err != nil {
		err.Result()
	}

	if k.ProjectDocExists(ctx, msg.GetProjectDid()) {
		return did.ErrorInvalidDid(types.DefaultCodeSpace, fmt.Sprintf("Project already exists")).Result()
	}
	k.SetProjectDoc(ctx, &msg)
	k.SetProjectWithdrawalTransactions(ctx, msg.GetProjectDid(), nil)

	return sdk.Result{
		Code: sdk.CodeOK,
	}
}

func handleMsgUpdateProjectStatus(ctx sdk.Context, k Keeper, bk bank.Keeper,
	msg MsgUpdateProjectStatus) sdk.Result {

	existingProjectDoc, err := getProjectDoc(ctx, k, msg.GetProjectDid())
	if err != nil {
		return sdk.ErrUnknownRequest("Could not find Project").Result()
	}

	newStatus := msg.GetStatus()
	if !newStatus.IsValidProgressionFrom(existingProjectDoc.GetStatus()) {
		return sdk.ErrUnknownRequest("Invalid Status Progression requested").Result()
	}

	if newStatus == FundedStatus {
		projectAddr, err := getProjectAccount(ctx, k, existingProjectDoc.GetProjectDid())
		if err != nil {
			return err.Result()
		}

		projectAcc := k.AccountKeeper.GetAccount(ctx, projectAddr)
		if projectAcc == nil {
			return sdk.ErrUnknownRequest("Could not find project account").Result()
		}

		minimumFunding := k.GetParams(ctx).ProjectMinimumInitialFunding
		if projectAcc.GetCoins().AmountOf(ixo.IxoNativeToken).LT(minimumFunding) {
			return sdk.ErrInsufficientFunds(
				fmt.Sprintf("Project has not reached minimum funding %s", minimumFunding)).Result()
		}
	}

	if newStatus == PaidoutStatus {
		res := payoutFees(ctx, k, bk, existingProjectDoc.GetProjectDid())
		if res.Code != sdk.CodeOK {
			return res
		}
	}

	existingProjectDoc.SetStatus(newStatus)
	_, _ = k.UpdateProjectDoc(ctx, existingProjectDoc)

	return sdk.Result{
		Code: sdk.CodeOK,
	}
}

func payoutFees(ctx sdk.Context, k Keeper, bk bank.Keeper,
	projectDid ixo.Did) sdk.Result {

	// TODO
	//_, err := ethClient.ProjectWalletFromProjectRegistry(ctx, projectDid)
	//if err != nil {
	//	return sdk.ErrUnknownRequest("Could not find Project Ethereum wallet").Result()
	//}

	_, err := payAllFeesToAddress(ctx, k, bk, projectDid, IxoAccountPayFeesId, IxoAccountFeesId)
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

	// TODO (contracts): ixoEthWallet := ck.GetContract(ctx, contracts.KeyFoundationWallet)

	// TODO: return payoutERC20AndRecon(ctx, k, bk, pk, ethClient, projectDid, IxoAccountFeesId, ixoEthWallet)
	return sdk.Result{}
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

func getIxoAmount(ctx sdk.Context, k Keeper, bk bank.Keeper, projectDid ixo.Did, accountID InternalAccountID) int64 {
	found := checkAccountInProjectAccounts(ctx, k, projectDid, accountID)
	if found {
		accAddr, _ := getAccountInProjectAccounts(ctx, k, projectDid, accountID)
		coins := bk.GetCoins(ctx, accAddr)
		return coins.AmountOf(ixo.IxoNativeToken).Int64()
	}

	return 0
}

func handleMsgCreateAgent(ctx sdk.Context, k Keeper, bk bank.Keeper, msg MsgCreateAgent) sdk.Result {
	_, err := createAccountInProjectAccounts(ctx, k, msg.GetProjectDid(), InternalAccountID(msg.Data.AgentDid))
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

	// TODO: check if project exists before calling processFees
	// Something will still fail but it's better to give a more meaningful error

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

	// TODO: check if project exists before calling processFees
	// Something will still fail but it's better to give a more meaningful error

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
		projectAddr, _ := getAccountInProjectAccounts(ctx, k, projectDid, InternalAccountID(msg.GetProjectDid()))
		evaluatorAccAddr, _ := getAccountInProjectAccounts(ctx, k, projectDid, InternalAccountID(msg.GetSenderDid()))

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

		feePercentage := fk.GetParams(ctx).EvaluationPayFeePercentage
		nodeFeePercentage := fk.GetParams(ctx).EvaluationPayNodeFeePercentage

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

func handleMsgWithdrawFunds(ctx sdk.Context, k Keeper, bk bank.Keeper,
	msg MsgWithdrawFunds) sdk.Result {

	withdrawFundsDoc := msg.GetWithdrawFundsDoc()
	projectDoc, err := getProjectDoc(ctx, k, withdrawFundsDoc.GetProjectDid())
	if err != nil {
		return sdk.ErrUnknownRequest("Could not find Project").Result()
	}

	if projectDoc.GetStatus() != PaidoutStatus {
		return sdk.ErrUnknownRequest("Project not in PAIDOUT Status").Result()
	}

	// TODO: implement below code

	//ethWalletAddress := withdrawFundsDoc.GetEthWallet()
	//projectDid := withdrawFundsDoc.GetProjectDid()

	//var payoutResult sdk.Result
	//if withdrawFundsDoc.IsRefund {
	//	payoutResult = payoutERC20AndRecon(ctx, k, bk, pk, ethClient, projectDid, projectDid, ethWalletAddress)
	//} else {
	//	senderDid := msg.GetSenderDid()
	//	payoutResult = payoutERC20AndRecon(ctx, k, bk, pk, ethClient, projectDid, senderDid, ethWalletAddress)
	//}

	return sdk.Result{}
}

//func payoutERC20AndRecon(ctx sdk.Context, k Keeper, bk bank.Keeper, pk params.Keeper, ethClient ixo.EthClient,
//	projectDid ixo.Did, accountID string, recipientEthAddress string) sdk.Result {
//
//	balanceToPay := getIxoAmount(ctx, k, bk, projectDid, accountID)
//	if balanceToPay > 0 {
//		projectEthWallet, err := ethClient.ProjectWalletFromProjectRegistry(ctx, projectDid)
//		if err != nil {
//			return sdk.ErrUnknownRequest("Could not find Project Ethereum wallet").Result()
//		}
//
//		account, errRes := getAccountInProjectAccounts(ctx, k, projectDid, accountID)
//		if errRes != nil {
//			return errRes.Result()
//		}
//
//		// TODO: Why is balanceToPay is added to account and removed right after??
//		_, err = bk.AddCoins(ctx, account, sdk.Coins{sdk.NewInt64Coin(ixo.IxoNativeToken, balanceToPay)})
//		if err != nil {
//		}
//
//		_, err = bk.SubtractCoins(ctx, account, sdk.Coins{sdk.NewInt64Coin(ixo.IxoNativeToken, balanceToPay)})
//		if err != nil {
//			return sdk.ErrUnknownRequest("Could not burn tokens from " + account.String()).Result()
//		}
//
//		_, actionID := ethClient.InitiateTokenTransfer(ctx, pk, projectEthWallet, recipientEthAddress, balanceToPay)
//
//		addProjectWithdrawalTransaction(ctx, k, projectDid, actionID, projectEthWallet, recipientEthAddress, balanceToPay)
//	}
//
//	return sdk.Result{
//		Code: sdk.CodeOK,
//	}
//}

func getProjectDoc(ctx sdk.Context, k Keeper, projectDid ixo.Did) (StoredProjectDoc, sdk.Error) {
	ixoProjectDoc, err := k.GetProjectDoc(ctx, projectDid)

	return ixoProjectDoc.(StoredProjectDoc), err
}

func processFees(ctx sdk.Context, k Keeper, fk fees.Keeper, bk bank.Keeper, feeType fees.FeeType, projectDid ixo.Did) (sdk.Result, sdk.Error) {

	projectAddr, _ := getProjectAccount(ctx, k, projectDid)
	var validatingNodeSetAddr sdk.AccAddress

	found := checkAccountInProjectAccounts(ctx, k, projectDid, ValidatingNodeSetAccountFeesId) // not found
	if !found {
		validatingNodeSetAddr, _ = createAccountInProjectAccounts(ctx, k, projectDid, ValidatingNodeSetAccountFeesId)
	} else {
		validatingNodeSetAddr, _ = getAccountInProjectAccounts(ctx, k, projectDid, ValidatingNodeSetAccountFeesId)
	}

	var ixoAddr sdk.AccAddress
	found = checkAccountInProjectAccounts(ctx, k, projectDid, IxoAccountFeesId) // found
	if !found {
		ixoAddr, _ = createAccountInProjectAccounts(ctx, k, projectDid, IxoAccountFeesId)
	} else {
		ixoAddr, _ = getAccountInProjectAccounts(ctx, k, projectDid, IxoAccountFeesId)
	}

	ixoFactor := fk.GetParams(ctx).IxoFactor
	nodePercentage := fk.GetParams(ctx).NodeFeePercentage

	var adjustedFeeAmount sdk.Dec
	switch feeType {
	case fees.FeeClaimTransaction:
		adjustedFeeAmount = fk.GetParams(ctx).ClaimFeeAmount.Mul(ixoFactor)
	case fees.FeeEvaluationTransaction:
		adjustedFeeAmount = fk.GetParams(ctx).EvaluationFeeAmount.Mul(ixoFactor)
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

func checkAccountInProjectAccounts(ctx sdk.Context, k Keeper, projectDid ixo.Did,
	accountId InternalAccountID) bool {
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

func createAccountInProjectAccounts(ctx sdk.Context, k Keeper, projectDid ixo.Did, accountId InternalAccountID) (sdk.AccAddress, sdk.Error) {
	acc, err := k.CreateNewAccount(ctx, projectDid, accountId)
	if err != nil {
		return nil, err
	}

	k.AddAccountToProjectAccounts(ctx, projectDid, accountId, acc)

	return acc.GetAddress(), nil
}

func getAccountInProjectAccounts(ctx sdk.Context, k Keeper, projectDid ixo.Did,
	accountId InternalAccountID) (sdk.AccAddress, sdk.Error) {
	accMap := k.GetAccountMap(ctx, projectDid)

	addr, found := accMap[accountId]
	if found {
		return addr, nil
	} else {
		return createAccountInProjectAccounts(ctx, k, projectDid, accountId)
	}
}

func getProjectAccount(ctx sdk.Context, k Keeper, projectDid ixo.Did) (sdk.AccAddress, sdk.Error) {
	return getAccountInProjectAccounts(ctx, k, projectDid, InternalAccountID(projectDid))
}

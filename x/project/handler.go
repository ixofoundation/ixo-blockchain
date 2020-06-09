package project

import (
	"encoding/hex"
	"fmt"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"github.com/ixofoundation/ixo-blockchain/x/project/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"

	"github.com/ixofoundation/ixo-blockchain/x/fees"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
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

	projectDid := msg.GetProjectDid()

	var err sdk.Error
	if _, err = createAccountInProjectAccounts(ctx, k, projectDid, IxoAccountFeesId); err != nil {
		return err.Result()
	}
	if _, err = createAccountInProjectAccounts(ctx, k, projectDid, IxoAccountPayFeesId); err != nil {
		return err.Result()
	}
	if _, err = createAccountInProjectAccounts(ctx, k, projectDid, InitiatingNodeAccountPayFeesId); err != nil {
		return err.Result()
	}
	if _, err = createAccountInProjectAccounts(ctx, k, projectDid, ValidatingNodeSetAccountFeesId); err != nil {
		return err.Result()
	}
	if _, err = createAccountInProjectAccounts(ctx, k, projectDid, InternalAccountID(msg.GetProjectDid())); err != nil {
		err.Result()
	}

	if k.ProjectDocExists(ctx, msg.GetProjectDid()) {
		return did.ErrorInvalidDid(types.DefaultCodespace, fmt.Sprintf("Project already exists")).Result()
	}
	k.SetProjectDoc(ctx, &msg)
	k.SetProjectWithdrawalTransactions(ctx, msg.GetProjectDid(), nil)

	return sdk.Result{}
}

func handleMsgUpdateProjectStatus(ctx sdk.Context, k Keeper, bk bank.Keeper,
	msg MsgUpdateProjectStatus) sdk.Result {

	existingProjectDoc, err := getProjectDoc(ctx, k, msg.ProjectDid)
	if err != nil {
		return sdk.ErrUnknownRequest("Could not find Project").Result()
	}

	newStatus := msg.Data.Status
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

	return sdk.Result{}
}

func payoutFees(ctx sdk.Context, k Keeper, bk bank.Keeper, projectDid ixo.Did) sdk.Result {

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

	ixoDid := k.GetParams(ctx).IxoDid
	amount := getIxoAmount(ctx, k, bk, projectDid, IxoAccountFeesId)
	err = payoutAndRecon(ctx, k, bk, projectDid, IxoAccountFeesId, ixoDid, amount)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{}
}

func payAllFeesToAddress(ctx sdk.Context, k Keeper, bk bank.Keeper, projectDid ixo.Did,
	sendingAddress InternalAccountID, receivingAddress InternalAccountID) (sdk.Events, sdk.Error) {
	feesToPay := getIxoAmount(ctx, k, bk, projectDid, sendingAddress)

	if feesToPay.Amount.LT(sdk.ZeroInt()) {
		return nil, sdk.ErrInternal("Negative fee to pay")
	}
	if feesToPay.Amount.IsZero() {
		return nil, nil
	}

	receivingAccount, err := getAccountInProjectAccounts(ctx, k, projectDid, receivingAddress)
	if err != nil {
		return sdk.Events{}, err
	}

	sendingAccount, _ := getAccountInProjectAccounts(ctx, k, projectDid, sendingAddress)

	return sdk.Events{}, bk.SendCoins(ctx, sendingAccount, receivingAccount, sdk.Coins{feesToPay})
}

func getIxoAmount(ctx sdk.Context, k Keeper, bk bank.Keeper, projectDid ixo.Did, accountID InternalAccountID) sdk.Coin {
	found := checkAccountInProjectAccounts(ctx, k, projectDid, accountID)
	if found {
		accAddr, _ := getAccountInProjectAccounts(ctx, k, projectDid, accountID)
		coins := bk.GetCoins(ctx, accAddr)
		return sdk.NewCoin(ixo.IxoNativeToken, coins.AmountOf(ixo.IxoNativeToken))
	}
	return sdk.NewCoin(ixo.IxoNativeToken, sdk.ZeroInt())
}

func handleMsgCreateAgent(ctx sdk.Context, k Keeper, bk bank.Keeper, msg MsgCreateAgent) sdk.Result {

	// Check if project exists
	_, err := getProjectDoc(ctx, k, msg.ProjectDid)
	if err != nil {
		return sdk.ErrUnknownRequest("Could not find Project").Result()
	}

	// Create account in project accounts for the agent
	_, err = createAccountInProjectAccounts(ctx, k, msg.ProjectDid, InternalAccountID(msg.Data.AgentDid))
	if err != nil {
		err.Result()
	}

	return sdk.Result{}
}

func handleMsgUpdateAgent(ctx sdk.Context, k Keeper, bk bank.Keeper, msg MsgUpdateAgent) sdk.Result {

	// Check if project exists
	_, err := getProjectDoc(ctx, k, msg.ProjectDid)
	if err != nil {
		return sdk.ErrUnknownRequest("Could not find Project").Result()
	}

	// TODO: implement agent update (or remove functionality)

	return sdk.Result{}
}

func handleMsgCreateClaim(ctx sdk.Context, k Keeper, fk fees.Keeper, bk bank.Keeper, msg MsgCreateClaim) sdk.Result {

	// Check if project exists
	_, err := getProjectDoc(ctx, k, msg.ProjectDid)
	if err != nil {
		return sdk.ErrUnknownRequest("Could not find Project").Result()
	}

	// Process claim fees
	err = processFees(ctx, k, fk, bk, fees.FeeClaimTransaction, msg.ProjectDid)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{}
}

func handleMsgCreateEvaluation(ctx sdk.Context, k Keeper, fk fees.Keeper, bk bank.Keeper, msg MsgCreateEvaluation) sdk.Result {

	// Check if project exists
	projectDoc, err := getProjectDoc(ctx, k, msg.ProjectDid)
	if err != nil {
		return sdk.ErrUnknownRequest("Could not find Project").Result()
	}

	// Process evaluation fees
	err = processFees(ctx, k, fk, bk, fees.FeeEvaluationTransaction, msg.ProjectDid)
	if err != nil {
		return err.Result()
	}

	// Process evaluator pay
	err = processEvaluatorPay(ctx, k, fk, bk, msg.ProjectDid,
		msg.SenderDid, projectDoc.GetEvaluatorPay())
	if err != nil {
		return err.Result()
	}

	return sdk.Result{}
}

func handleMsgWithdrawFunds(ctx sdk.Context, k Keeper, bk bank.Keeper,
	msg MsgWithdrawFunds) sdk.Result {

	withdrawFundsDoc := msg.Data
	projectDoc, err := getProjectDoc(ctx, k, withdrawFundsDoc.ProjectDid)
	if err != nil {
		return sdk.ErrUnknownRequest("Could not find Project").Result()
	}

	if projectDoc.GetStatus() != PaidoutStatus {
		return sdk.ErrUnknownRequest("Project not in PAIDOUT Status").Result()
	}

	projectDid := withdrawFundsDoc.ProjectDid
	recipientDid := withdrawFundsDoc.RecipientDid
	amount := withdrawFundsDoc.Amount

	// If this is a refund, recipient has to be the project creator
	if withdrawFundsDoc.IsRefund && (recipientDid != projectDoc.GetSenderDid()) {
		return sdk.ErrUnknownRequest("Only project creator can get a refund").Result()
	}

	var fromAccountId InternalAccountID
	if withdrawFundsDoc.IsRefund {
		fromAccountId = InternalAccountID(projectDid)
	} else {
		fromAccountId = InternalAccountID(recipientDid)
	}

	amountCoin := sdk.NewCoin(ixo.IxoNativeToken, amount)
	err = payoutAndRecon(ctx, k, bk, projectDid, fromAccountId, recipientDid, amountCoin)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{}
}

func payoutAndRecon(ctx sdk.Context, k Keeper, bk bank.Keeper, projectDid ixo.Did,
	fromAccountId InternalAccountID, recipientDid ixo.Did, amount sdk.Coin) sdk.Error {

	ixoBalance := getIxoAmount(ctx, k, bk, projectDid, fromAccountId)
	if ixoBalance.IsLT(amount) {
		return sdk.ErrInternal("insufficient funds in specified account")
	}

	fromAccount, err := getAccountInProjectAccounts(ctx, k, projectDid, fromAccountId)
	if err != nil {
		return err
	}

	recipientAddr := ixo.StringToAddr(recipientDid)
	err = bk.SendCoins(ctx, fromAccount, recipientAddr, sdk.Coins{amount})
	if err != nil {
		return err
	}

	var actionId [32]byte
	dec := sdk.OneDec() // TODO: should increment with each withdrawal
	copy(actionId[:], dec.Bytes())

	addProjectWithdrawalTransaction(ctx, k, projectDid, actionId, recipientDid, amount)
	return nil
}

func getProjectDoc(ctx sdk.Context, k Keeper, projectDid ixo.Did) (StoredProjectDoc, sdk.Error) {
	ixoProjectDoc, err := k.GetProjectDoc(ctx, projectDid)
	if err != nil {
		return nil, err
	}

	return ixoProjectDoc.(StoredProjectDoc), nil
}

func processFees(ctx sdk.Context, k Keeper, fk fees.Keeper, bk bank.Keeper, feeType fees.FeeType, projectDid ixo.Did) sdk.Error {

	projectAddr, _ := getProjectAccount(ctx, k, projectDid)

	validatingNodeSetAddr, err := getAccountInProjectAccounts(ctx, k, projectDid, ValidatingNodeSetAccountFeesId)
	if err != nil {
		return err
	}

	ixoAddr, err := getAccountInProjectAccounts(ctx, k, projectDid, IxoAccountFeesId)
	if err != nil {
		return err
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
		return sdk.ErrUnknownRequest("Invalid Fee type.")
	}

	nodeAmount := adjustedFeeAmount.Mul(nodePercentage).RoundInt64()
	ixoAmount := adjustedFeeAmount.RoundInt64() - nodeAmount

	err = bk.SendCoins(ctx, projectAddr, validatingNodeSetAddr, sdk.Coins{sdk.NewInt64Coin(ixo.IxoNativeToken, nodeAmount)})
	if err != nil {
		return err
	}

	err = bk.SendCoins(ctx, projectAddr, ixoAddr, sdk.Coins{sdk.NewInt64Coin(ixo.IxoNativeToken, ixoAmount)})
	if err != nil {
		return err
	}

	return nil
}

func processEvaluatorPay(ctx sdk.Context, k Keeper, fk fees.Keeper, bk bank.Keeper, projectDid, senderDid ixo.Did, evaluatorPay int64) sdk.Error {

	if evaluatorPay == 0 {
		return nil
	}

	projectAddr, _ := getAccountInProjectAccounts(ctx, k, projectDid, InternalAccountID(projectDid))
	evaluatorAccAddr, _ := getAccountInProjectAccounts(ctx, k, projectDid, InternalAccountID(senderDid))

	nodeAddr, err := getAccountInProjectAccounts(ctx, k, projectDid, InitiatingNodeAccountPayFeesId)
	if err != nil {
		return err
	}

	ixoAddr, err := getAccountInProjectAccounts(ctx, k, projectDid, IxoAccountPayFeesId)
	if err != nil {
		return err
	}

	feePercentage := fk.GetParams(ctx).EvaluationPayFeePercentage
	nodeFeePercentage := fk.GetParams(ctx).EvaluationPayNodeFeePercentage

	totalEvaluatorPayAmount := sdk.NewDec(evaluatorPay).Mul(ixo.IxoDecimals) // This is in IXO * 10^8
	evaluatorPayFeeAmount := totalEvaluatorPayAmount.Mul(feePercentage)
	evaluatorPayLessFees := totalEvaluatorPayAmount.Sub(evaluatorPayFeeAmount)
	nodePayFees := evaluatorPayFeeAmount.Mul(nodeFeePercentage)
	ixoPayFees := evaluatorPayFeeAmount.Sub(nodePayFees)

	err = bk.SendCoins(ctx, projectAddr, evaluatorAccAddr, sdk.Coins{sdk.NewInt64Coin(ixo.IxoNativeToken, evaluatorPayLessFees.RoundInt64())})
	if err != nil {
		return err
	}

	err = bk.SendCoins(ctx, projectAddr, nodeAddr, sdk.Coins{sdk.NewInt64Coin(ixo.IxoNativeToken, nodePayFees.RoundInt64())})
	if err != nil {
		return err
	}

	err = bk.SendCoins(ctx, projectAddr, ixoAddr, sdk.Coins{sdk.NewInt64Coin(ixo.IxoNativeToken, ixoPayFees.RoundInt64())})
	if err != nil {
		return err
	}

	return nil
}

func checkAccountInProjectAccounts(ctx sdk.Context, k Keeper, projectDid ixo.Did,
	accountId InternalAccountID) bool {
	accMap := k.GetAccountMap(ctx, projectDid)
	_, found := accMap[accountId]

	return found
}

func addProjectWithdrawalTransaction(ctx sdk.Context, k Keeper, projectDid ixo.Did,
	actionID [32]byte, recipientDid ixo.Did, amount sdk.Coin) {
	actionIDStr := "0x" + hex.EncodeToString(actionID[:])

	withdrawalInfo := WithdrawalInfo{
		ActionID:     actionIDStr,
		ProjectDid:   projectDid,
		RecipientDid: recipientDid,
		Amount:       amount,
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

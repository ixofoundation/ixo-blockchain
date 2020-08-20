package project

import (
	"encoding/hex"
	"fmt"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"github.com/ixofoundation/ixo-blockchain/x/project/internal/types"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"

	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"github.com/ixofoundation/ixo-blockchain/x/payments"
)

const (
	IxoAccountFeesId               InternalAccountID = "IxoFees"
	IxoAccountPayFeesId            InternalAccountID = "IxoPayFees"
	InitiatingNodeAccountPayFeesId InternalAccountID = "InitiatingNodePayFees"
	ValidatingNodeSetAccountFeesId InternalAccountID = "ValidatingNodeSetFees"
)

func NewHandler(k Keeper, pk payments.Keeper, bk bank.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgCreateProject:
			return handleMsgCreateProject(ctx, k, msg)
		case MsgUpdateProjectStatus:
			return handleMsgUpdateProjectStatus(ctx, k, bk, msg)
		case MsgCreateAgent:
			return handleMsgCreateAgent(ctx, k, msg)
		case MsgUpdateAgent:
			return handleMsgUpdateAgent(ctx, k, bk, msg)
		case MsgCreateClaim:
			return handleMsgCreateClaim(ctx, k, pk, bk, msg)
		case MsgCreateEvaluation:
			return handleMsgCreateEvaluation(ctx, k, pk, bk, msg)
		case MsgWithdrawFunds:
			return handleMsgWithdrawFunds(ctx, k, bk, msg)
		default:
			return sdk.ErrUnknownRequest("No match for message type.").Result()
		}
	}
}

func handleMsgCreateProject(ctx sdk.Context, k Keeper, msg MsgCreateProject) sdk.Result {

	var err sdk.Error
	if _, err = createAccountInProjectAccounts(ctx, k, msg.ProjectDid, IxoAccountFeesId); err != nil {
		return err.Result()
	}
	if _, err = createAccountInProjectAccounts(ctx, k, msg.ProjectDid, IxoAccountPayFeesId); err != nil {
		return err.Result()
	}
	if _, err = createAccountInProjectAccounts(ctx, k, msg.ProjectDid, InitiatingNodeAccountPayFeesId); err != nil {
		return err.Result()
	}
	if _, err = createAccountInProjectAccounts(ctx, k, msg.ProjectDid, ValidatingNodeSetAccountFeesId); err != nil {
		return err.Result()
	}
	if _, err = createAccountInProjectAccounts(ctx, k, msg.ProjectDid, InternalAccountID(msg.ProjectDid)); err != nil {
		err.Result()
	}

	if k.ProjectDocExists(ctx, msg.ProjectDid) {
		return did.ErrorInvalidDid(types.DefaultCodespace, fmt.Sprintf("Project already exists")).Result()
	}

	projectDoc := NewProjectDoc(
		msg.TxHash, msg.ProjectDid, msg.SenderDid,
		msg.PubKey, types.NullStatus, msg.Data)

	k.SetProjectDoc(ctx, &projectDoc)
	k.SetProjectWithdrawalTransactions(ctx, msg.ProjectDid, nil)
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateProject,
			sdk.NewAttribute(types.AttributeKeyTxHash, msg.TxHash),
			sdk.NewAttribute(types.AttributeKeySenderDid, msg.SenderDid),
			sdk.NewAttribute(types.AttributeKeyProjectDid, msg.ProjectDid),
			sdk.NewAttribute(types.AttributeKeyPubKey, msg.PubKey),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleMsgUpdateProjectStatus(ctx sdk.Context, k Keeper, bk bank.Keeper,
	msg MsgUpdateProjectStatus) sdk.Result {

	projectDoc, err := k.GetProjectDoc(ctx, msg.ProjectDid)
	if err != nil {
		return sdk.ErrUnknownRequest("Could not find Project").Result()
	}

	newStatus := msg.Data.Status

	if !newStatus.IsValidProgressionFrom(projectDoc.GetStatus()) {
		return sdk.ErrUnknownRequest("Invalid Status Progression requested").Result()
	}

	if newStatus == FundedStatus {
		projectAddr, err := getProjectAccount(ctx, k, projectDoc.GetProjectDid())
		if err != nil {
			return err.Result()
		}

		projectAcc := k.AccountKeeper.GetAccount(ctx, projectAddr)
		if projectAcc == nil {
			return sdk.ErrUnknownRequest("Could not find project account").Result()
		}

		minimumFunding := k.GetParams(ctx).ProjectMinimumInitialFunding
		if minimumFunding.IsAnyGT(projectAcc.GetCoins()) {
			return sdk.ErrInsufficientFunds(
				fmt.Sprintf("Project has not reached minimum funding %s", minimumFunding)).Result()
		}
	}

	if newStatus == PaidoutStatus {
		res := payoutFees(ctx, k, bk, projectDoc.GetProjectDid())
		if res.Code != sdk.CodeOK {
			return res
		}
	}

	projectDoc.SetStatus(newStatus)
	k.SetProjectDoc(ctx, projectDoc)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUpdateProjectStatus,
			sdk.NewAttribute(types.AttributeKeyTxHash, msg.TxHash),
			sdk.NewAttribute(types.AttributeKeySenderDid, msg.SenderDid),
			sdk.NewAttribute(types.AttributeKeyProjectDid, msg.ProjectDid),
			sdk.NewAttribute(types.AttributeKeyEthFundingTxnID, msg.Data.EthFundingTxnID),
			sdk.NewAttribute(types.AttributeKeyUpdatedStatus, fmt.Sprint(msg.Data.Status)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return sdk.Result{Events: ctx.EventManager().Events()}

}

func payoutFees(ctx sdk.Context, k Keeper, bk bank.Keeper, projectDid did.Did) sdk.Result {

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

func payAllFeesToAddress(ctx sdk.Context, k Keeper, bk bank.Keeper, projectDid did.Did,
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

func getIxoAmount(ctx sdk.Context, k Keeper, bk bank.Keeper, projectDid did.Did, accountID InternalAccountID) sdk.Coin {
	found := checkAccountInProjectAccounts(ctx, k, projectDid, accountID)
	if found {
		accAddr, _ := getAccountInProjectAccounts(ctx, k, projectDid, accountID)
		coins := bk.GetCoins(ctx, accAddr)
		return sdk.NewCoin(ixo.IxoNativeToken, coins.AmountOf(ixo.IxoNativeToken))
	}
	return sdk.NewCoin(ixo.IxoNativeToken, sdk.ZeroInt())
}

func handleMsgCreateAgent(ctx sdk.Context, k Keeper, msg MsgCreateAgent) sdk.Result {

	// Check if project exists
	_, err := k.GetProjectDoc(ctx, msg.ProjectDid)
	if err != nil {
		return sdk.ErrUnknownRequest("Could not find Project").Result()
	}

	// Create account in project accounts for the agent
	_, err = createAccountInProjectAccounts(ctx, k, msg.ProjectDid, InternalAccountID(msg.Data.AgentDid))
	if err != nil {
		err.Result()
	}
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateAgent,
			sdk.NewAttribute(types.AttributeKeyTxHash, msg.TxHash),
			sdk.NewAttribute(types.AttributeKeySenderDid, msg.SenderDid),
			sdk.NewAttribute(types.AttributeKeyProjectDid, msg.ProjectDid),
			sdk.NewAttribute(types.AttributeKeyAgentDid, msg.Data.AgentDid),
			sdk.NewAttribute(types.AttributeKeyAgentRole, msg.Data.Role),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleMsgUpdateAgent(ctx sdk.Context, k Keeper, bk bank.Keeper, msg MsgUpdateAgent) sdk.Result {

	// Check if project exists
	_, err := k.GetProjectDoc(ctx, msg.ProjectDid)
	if err != nil {
		return sdk.ErrUnknownRequest("Could not find Project").Result()
	}

	// TODO: implement agent update (or remove functionality)

	return sdk.Result{}
}

func handleMsgCreateClaim(ctx sdk.Context, k Keeper, pk payments.Keeper,
	bk bank.Keeper, msg MsgCreateClaim) sdk.Result {

	// Check if project exists
	projectDoc, err := k.GetProjectDoc(ctx, msg.ProjectDid)
	if err != nil {
		return sdk.ErrUnknownRequest("Could not find Project").Result()
	}

	// Get sender address
	senderDidDoc, err := k.DidKeeper.GetDidDoc(ctx, msg.SenderDid)
	if err != nil {
		return err.Result()
	}
	senderAddr := senderDidDoc.Address()

	// Process claim fees
	err = processFees(ctx, k, pk, bk, payments.FeeClaimTransaction,
		msg.ProjectDid)
	if err != nil {
		return err.Result()
	}

	// Process claimer pay
	err = processClaimerPay(ctx, k, bk, pk, projectDoc.GetProjectDid(),
		senderAddr, projectDoc.GetClaimerPay())
	if err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateClaim,
			sdk.NewAttribute(types.AttributeKeyTxHash, msg.TxHash),
			sdk.NewAttribute(types.AttributeKeySenderDid, msg.SenderDid),
			sdk.NewAttribute(types.AttributeKeyProjectDid, msg.ProjectDid),
			sdk.NewAttribute(types.AttributeKeyClaimID, msg.Data.ClaimID),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleMsgCreateEvaluation(ctx sdk.Context, k Keeper, pk payments.Keeper,
	bk bank.Keeper, msg MsgCreateEvaluation) sdk.Result {

	// Check if project exists
	projectDoc, err := k.GetProjectDoc(ctx, msg.ProjectDid)
	if err != nil {
		return sdk.ErrUnknownRequest("Could not find Project").Result()
	}

	// Get sender address
	senderDidDoc, err := k.DidKeeper.GetDidDoc(ctx, msg.SenderDid)
	if err != nil {
		return err.Result()
	}
	senderAddr := senderDidDoc.Address()

	// Process evaluation fees
	err = processFees(ctx, k, pk, bk, payments.FeeEvaluationTransaction,
		msg.ProjectDid)
	if err != nil {
		return err.Result()
	}

	// Process evaluator pay
	err = processEvaluatorPay(ctx, k, bk, pk, msg.ProjectDid,
		senderAddr, projectDoc.GetEvaluatorPay())
	if err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateEvaluation,
			sdk.NewAttribute(types.AttributeKeyTxHash, msg.TxHash),
			sdk.NewAttribute(types.AttributeKeySenderDid, msg.SenderDid),
			sdk.NewAttribute(types.AttributeKeyProjectDid, msg.ProjectDid),
			sdk.NewAttribute(types.AttributeKeyClaimID, msg.Data.ClaimID),
			sdk.NewAttribute(types.AttributeKeyClaimStatus, fmt.Sprint(msg.Data.Status)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleMsgWithdrawFunds(ctx sdk.Context, k Keeper, bk bank.Keeper,
	msg MsgWithdrawFunds) sdk.Result {

	withdrawFundsDoc := msg.Data
	projectDoc, err := k.GetProjectDoc(ctx, withdrawFundsDoc.ProjectDid)
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

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeWithdrawFunds,
			sdk.NewAttribute(types.AttributeKeySenderDid, msg.SenderDid),
			sdk.NewAttribute(types.AttributeKeyProjectDid, msg.Data.ProjectDid),
			sdk.NewAttribute(types.AttributeKeyRecipientDid, msg.Data.RecipientDid),
			sdk.NewAttribute(types.AttributeKeyAmount, msg.Data.Amount.String()),
			sdk.NewAttribute(types.AttributeKeyIsRefund, strconv.FormatBool(msg.Data.IsRefund)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return sdk.Result{Events: ctx.EventManager().Events()}
}

func payoutAndRecon(ctx sdk.Context, k Keeper, bk bank.Keeper, projectDid did.Did,
	fromAccountId InternalAccountID, recipientDid did.Did, amount sdk.Coin) sdk.Error {

	ixoBalance := getIxoAmount(ctx, k, bk, projectDid, fromAccountId)
	if ixoBalance.IsLT(amount) {
		return sdk.ErrInternal("insufficient funds in specified account")
	}

	fromAccount, err := getAccountInProjectAccounts(ctx, k, projectDid, fromAccountId)
	if err != nil {
		return err
	}

	// Get recipient address
	recipientDidDoc, err := k.DidKeeper.GetDidDoc(ctx, recipientDid)
	if err != nil {
		return err
	}
	recipientAddr := recipientDidDoc.Address()

	err = bk.SendCoins(ctx, fromAccount, recipientAddr, sdk.Coins{amount})
	if err != nil {
		return err
	}

	var actionId [32]byte
	dec := sdk.OneDec() // TODO: should increment with each withdrawal (ref: #113)
	copy(actionId[:], dec.Bytes())

	addProjectWithdrawalTransaction(ctx, k, projectDid, actionId, recipientDid, amount)
	return nil
}

func processFees(ctx sdk.Context, k Keeper, pk payments.Keeper, bk bank.Keeper,
	feeType payments.FeeType, projectDid did.Did) sdk.Error {

	projectAddr, _ := getProjectAccount(ctx, k, projectDid)

	validatingNodeSetAddr, err := getAccountInProjectAccounts(ctx, k, projectDid, ValidatingNodeSetAccountFeesId)
	if err != nil {
		return err
	}

	ixoAddr, err := getAccountInProjectAccounts(ctx, k, projectDid, IxoAccountFeesId)
	if err != nil {
		return err
	}

	ixoFactor := pk.GetParams(ctx).IxoFactor
	nodePercentage := pk.GetParams(ctx).NodeFeePercentage

	var feeAmount sdk.Coins
	switch feeType {
	case payments.FeeClaimTransaction:
		feeAmount = pk.GetParams(ctx).ClaimFeeAmount
	case payments.FeeEvaluationTransaction:
		feeAmount = pk.GetParams(ctx).EvaluationFeeAmount
	default:
		return sdk.ErrUnknownRequest("Invalid Fee type.")
	}
	adjustedFeeAmount := sdk.NewDecCoins(feeAmount).MulDec(ixoFactor)

	nodeAmount := adjustedFeeAmount.MulDec(nodePercentage)
	ixoAmount := adjustedFeeAmount.Sub(nodeAmount)

	nodeAmountCoins, _ := nodeAmount.TruncateDecimal()
	ixoAmountCoins, _ := ixoAmount.TruncateDecimal()

	err = bk.SendCoins(ctx, projectAddr, validatingNodeSetAddr, nodeAmountCoins)
	if err != nil {
		return err
	}

	err = bk.SendCoins(ctx, projectAddr, ixoAddr, ixoAmountCoins)
	if err != nil {
		return err
	}

	return nil
}

func processEvaluatorPay(ctx sdk.Context, k Keeper, bk bank.Keeper,
	pk payments.Keeper, projectDid did.Did, senderAddr sdk.AccAddress,
	evaluatorPay sdk.Coins) sdk.Error {

	if evaluatorPay.IsZero() {
		return nil
	}

	// Get project address
	projectAddr, err := getAccountInProjectAccounts(
		ctx, k, projectDid, InternalAccountID(projectDid))
	if err != nil {
		return err
	}

	nodeAddr, err := getAccountInProjectAccounts(ctx, k, projectDid,
		InitiatingNodeAccountPayFeesId)
	if err != nil {
		return err
	}

	ixoAddr, err := getAccountInProjectAccounts(ctx, k, projectDid,
		IxoAccountPayFeesId)
	if err != nil {
		return err
	}

	feePercentage := pk.GetParams(ctx).EvaluationPayFeePercentage
	nodeFeePercentage := pk.GetParams(ctx).EvaluationPayNodeFeePercentage

	totalEvaluatorPayAmount := sdk.NewDecCoins(evaluatorPay)
	evaluatorPayFeeAmount := totalEvaluatorPayAmount.MulDec(feePercentage)
	evaluatorPayLessFees := totalEvaluatorPayAmount.Sub(evaluatorPayFeeAmount)
	nodePayFees := evaluatorPayFeeAmount.MulDec(nodeFeePercentage)
	ixoPayFees := evaluatorPayFeeAmount.Sub(nodePayFees)

	evaluatorPayLessFeesCoins, _ := evaluatorPayLessFees.TruncateDecimal()
	nodePayFeesCoins, _ := nodePayFees.TruncateDecimal()
	ixoPayFeesCoins, _ := ixoPayFees.TruncateDecimal()

	if !evaluatorPayLessFeesCoins.IsZero() {
		err = bk.SendCoins(ctx, projectAddr, senderAddr, evaluatorPayLessFeesCoins)
		if err != nil {
			return err
		}
	}

	if !nodePayFeesCoins.IsZero() {
		err = bk.SendCoins(ctx, projectAddr, nodeAddr, nodePayFeesCoins)
		if err != nil {
			return err
		}
	}

	if !ixoPayFeesCoins.IsZero() {
		err = bk.SendCoins(ctx, projectAddr, ixoAddr, ixoPayFeesCoins)
		if err != nil {
			return err
		}
	}

	return nil
}

func processClaimerPay(ctx sdk.Context, k Keeper, bk bank.Keeper,
	pk payments.Keeper, projectDid did.Did, senderAddr sdk.AccAddress,
	claimerPay sdk.Coins) sdk.Error {

	if claimerPay.IsZero() {
		return nil
	}

	// Get project address
	projectAddr, err := getAccountInProjectAccounts(
		ctx, k, projectDid, InternalAccountID(projectDid))
	if err != nil {
		return err
	}

	contractId := fmt.Sprintf("payment:contract:%s:%s:%s",
		ModuleName, projectDid, senderAddr.String())
	templateId := fmt.Sprintf("payment:template:%s:%s:%s",
		ModuleName, projectDid, senderAddr.String())

	// Create or get payment template and contract
	var template payments.PaymentTemplate
	var contract payments.PaymentContract
	if !pk.PaymentContractExists(ctx, contractId) {
		// Check template existence
		if pk.PaymentTemplateExists(ctx, templateId) {
			panic("expected payment template to not exist")
		}

		// Create template
		payAmount := claimerPay
		payMinimum := claimerPay
		payMaximum, _ := sdk.NewDecCoins(claimerPay).MulDec(
			sdk.NewDec(10)).TruncateDecimal()
		distribution := payments.NewDistribution(
			payments.NewFullDistributionShare(senderAddr))
		template = payments.NewPaymentTemplate(
			templateId, payAmount, payMinimum, payMaximum, nil, distribution)
		if err := template.Validate(); err != nil {
			return err
		}

		// Create contract
		contract = payments.NewPaymentContract(contractId, templateId,
			projectAddr, projectAddr, false, true, sdk.ZeroUint())

		pk.SetPaymentTemplate(ctx, template)
		pk.SetPaymentContract(ctx, contract)
	} else {
		var err sdk.Error
		template, err = pk.GetPaymentTemplate(ctx, templateId)
		if err != nil {
			panic("expected payment template to exist")
		}
		contract, err = pk.GetPaymentContract(ctx, contractId)
		if err != nil {
			panic("expected payment contract to exist")
		}
	}

	// Effect payment
	if contract.CanEffectPayment(template) {
		// Check that project has enough tokens to effect contract payment
		// (assume no effect from PaymentMin, PaymentMax, Discounts)
		if !bk.HasCoins(ctx, projectAddr, template.PaymentAmount) {
			return sdk.ErrInsufficientCoins("project has insufficient funds")
		}

		// Effect payment
		effected, err := pk.EffectPayment(ctx, bk, contractId)
		if err != nil {
			return err
		} else if !effected {
			panic("expected to be able to effect contract payment")
		}
	} else {
		return sdk.ErrInternal("cannot effect contract payment (max reached?)")
	}

	return nil
}

func checkAccountInProjectAccounts(ctx sdk.Context, k Keeper, projectDid did.Did,
	accountId InternalAccountID) bool {
	accMap := k.GetAccountMap(ctx, projectDid)
	_, found := accMap[accountId]

	return found
}

func addProjectWithdrawalTransaction(ctx sdk.Context, k Keeper, projectDid did.Did,
	actionID [32]byte, recipientDid did.Did, amount sdk.Coin) {
	actionIDStr := "0x" + hex.EncodeToString(actionID[:])

	withdrawalInfo := WithdrawalInfo{
		ActionID:     actionIDStr,
		ProjectDid:   projectDid,
		RecipientDid: recipientDid,
		Amount:       amount,
	}

	k.AddProjectWithdrawalTransaction(ctx, projectDid, withdrawalInfo)
}

func createAccountInProjectAccounts(ctx sdk.Context, k Keeper, projectDid did.Did, accountId InternalAccountID) (sdk.AccAddress, sdk.Error) {
	acc, err := k.CreateNewAccount(ctx, projectDid, accountId)
	if err != nil {
		return nil, err
	}

	k.AddAccountToProjectAccounts(ctx, projectDid, accountId, acc)

	return acc.GetAddress(), nil
}

func getAccountInProjectAccounts(ctx sdk.Context, k Keeper, projectDid did.Did,
	accountId InternalAccountID) (sdk.AccAddress, sdk.Error) {
	accMap := k.GetAccountMap(ctx, projectDid)

	addr, found := accMap[accountId]
	if found {
		return addr, nil
	} else {
		return createAccountInProjectAccounts(ctx, k, projectDid, accountId)
	}
}

func getProjectAccount(ctx sdk.Context, k Keeper, projectDid did.Did) (sdk.AccAddress, sdk.Error) {
	return getAccountInProjectAccounts(ctx, k, projectDid, InternalAccountID(projectDid))
}

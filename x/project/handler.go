package project

import (
	"encoding/hex"
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

func NewHandler(k Keeper, fk payments.Keeper, bk bank.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgCreateProject:
			return handleMsgCreateProject(ctx, k, msg)
		case MsgUpdateProjectStatus:
			return handleMsgUpdateProjectStatus(ctx, k, bk, msg)
		case MsgCreateAgent:
			return handleMsgCreateAgent(ctx, k, msg)
		case MsgUpdateAgent:
			return handleMsgUpdateAgent(ctx, k, msg)
		case MsgCreateClaim:
			return handleMsgCreateClaim(ctx, k, fk, bk, msg)
		case MsgCreateEvaluation:
			return handleMsgCreateEvaluation(ctx, k, fk, bk, msg)
		case MsgWithdrawFunds:
			return handleMsgWithdrawFunds(ctx, k, bk, msg)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "No match for message type.")
		}
	}
}

func handleMsgCreateProject(ctx sdk.Context, k Keeper, msg MsgCreateProject) (*sdk.Result, error) {

	var err error
	if _, err = createAccountInProjectAccounts(ctx, k, msg.ProjectDid, IxoAccountFeesId); err != nil {
		return nil, err
	}
	if _, err = createAccountInProjectAccounts(ctx, k, msg.ProjectDid, IxoAccountPayFeesId); err != nil {
		return nil, err
	}
	if _, err = createAccountInProjectAccounts(ctx, k, msg.ProjectDid, InitiatingNodeAccountPayFeesId); err != nil {
		return nil, err
	}
	if _, err = createAccountInProjectAccounts(ctx, k, msg.ProjectDid, ValidatingNodeSetAccountFeesId); err != nil {
		return nil, err
	}
	if _, err = createAccountInProjectAccounts(ctx, k, msg.ProjectDid, InternalAccountID(msg.ProjectDid)); err != nil {
		return nil, err
	}

	if k.ProjectDocExists(ctx, msg.ProjectDid) {
		return nil, sdkerrors.Wrap(did.ErrInvalidDid, "Project already exists")
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

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgUpdateProjectStatus(ctx sdk.Context, k Keeper, bk bank.Keeper,
	msg MsgUpdateProjectStatus) (*sdk.Result, error) {

	projectDoc, err := k.GetProjectDoc(ctx, msg.ProjectDid)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Could not find Project")
	}

	newStatus := msg.Data.Status

	if !newStatus.IsValidProgressionFrom(projectDoc.GetStatus()) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Invalid Status Progression requested")
	}

	if newStatus == FundedStatus {
		projectAddr, err := getProjectAccount(ctx, k, projectDoc.GetProjectDid())
		if err != nil {
			return nil, err
		}

		projectAcc := k.AccountKeeper.GetAccount(ctx, projectAddr)
		if projectAcc == nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Could not find project account")
		}

		minimumFunding := k.GetParams(ctx).ProjectMinimumInitialFunding
		if projectAcc.GetCoins().AmountOf(ixo.IxoNativeToken).LT(minimumFunding) {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "Project has not reached minimum funding %s", minimumFunding)
		}
	}

	if newStatus == PaidoutStatus {
		result, res := payoutFees(ctx, k, bk, projectDoc.GetProjectDid())
		if res != nil {
			return result, res
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

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil

}

func payoutFees(ctx sdk.Context, k Keeper, bk bank.Keeper, projectDid did.Did) (*sdk.Result, error) {

	_, err := payAllFeesToAddress(ctx, k, bk, projectDid, IxoAccountPayFeesId, IxoAccountFeesId)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInternal, "Failed to send coins")
	}

	_, err = payAllFeesToAddress(ctx, k, bk, projectDid, InitiatingNodeAccountPayFeesId, IxoAccountFeesId)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInternal, "Failed to send coins")
	}

	_, err = payAllFeesToAddress(ctx, k, bk, projectDid, ValidatingNodeSetAccountFeesId, IxoAccountFeesId)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInternal, "Failed to send coins")
	}

	ixoDid := k.GetParams(ctx).IxoDid
	amount := getIxoAmount(ctx, k, bk, projectDid, IxoAccountFeesId)
	err = payoutAndRecon(ctx, k, bk, projectDid, IxoAccountFeesId, ixoDid, amount)
	if err != nil {
		return nil, err
	}

	return &sdk.Result{}, nil
}

func payAllFeesToAddress(ctx sdk.Context, k Keeper, bk bank.Keeper, projectDid did.Did,
	sendingAddress InternalAccountID, receivingAddress InternalAccountID) (*sdk.Result, error) {
	feesToPay := getIxoAmount(ctx, k, bk, projectDid, sendingAddress)

	if feesToPay.Amount.LT(sdk.ZeroInt()) {
		return nil, sdkerrors.Wrap(types.ErrInternal, "Negative fee to pay")
	}
	if feesToPay.Amount.IsZero() {
		return nil, nil
	}

	receivingAccount, err := getAccountInProjectAccounts(ctx, k, projectDid, receivingAddress)
	if err != nil {
		return nil, err
	}

	sendingAccount, _ := getAccountInProjectAccounts(ctx, k, projectDid, sendingAddress)

	return nil, bk.SendCoins(ctx, sendingAccount, receivingAccount, sdk.Coins{feesToPay})
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

func handleMsgCreateAgent(ctx sdk.Context, k Keeper, msg MsgCreateAgent) (*sdk.Result, error) {

	// Check if project exists
	_, err := k.GetProjectDoc(ctx, msg.ProjectDid)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Could not find Project")
	}

	// Create account in project accounts for the agent
	_, err = createAccountInProjectAccounts(ctx, k, msg.ProjectDid, InternalAccountID(msg.Data.AgentDid))
	if err != nil {
		return nil, err
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

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgUpdateAgent(ctx sdk.Context, k Keeper, msg MsgUpdateAgent) (*sdk.Result, error) {

	// Check if project exists
	_, err := k.GetProjectDoc(ctx, msg.ProjectDid)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Could not find Project")
	}

	// TODO: implement agent update (or remove functionality)

	return nil, nil
}

func handleMsgCreateClaim(ctx sdk.Context, k Keeper, fk payments.Keeper,
	bk bank.Keeper, msg MsgCreateClaim) (*sdk.Result, error) {

	// Check if project exists
	_, err := k.GetProjectDoc(ctx, msg.ProjectDid)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Could not find Project")
	}

	// Process claim fees
	err = processFees(
		ctx, k, fk, bk, payments.FeeClaimTransaction, msg.ProjectDid)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUpdateAgent,
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

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgCreateEvaluation(ctx sdk.Context, k Keeper, fk payments.Keeper, bk bank.Keeper, msg MsgCreateEvaluation) (*sdk.Result, error) {

	// Check if project exists
	projectDoc, err := k.GetProjectDoc(ctx, msg.ProjectDid)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Could not find Project")
	}

	// Process evaluation fees
	err = processFees(
		ctx, k, fk, bk, payments.FeeEvaluationTransaction, msg.ProjectDid)
	if err != nil {
		return nil, err
	}

	// Process evaluator pay
	err = processEvaluatorPay(ctx, k, fk, bk, msg.ProjectDid,
		msg.SenderDid, projectDoc.GetEvaluatorPay())
	if err != nil {
		return nil, err
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

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgWithdrawFunds(ctx sdk.Context, k Keeper, bk bank.Keeper,
	msg MsgWithdrawFunds) (*sdk.Result, error) {

	withdrawFundsDoc := msg.Data
	projectDoc, err := k.GetProjectDoc(ctx, withdrawFundsDoc.ProjectDid)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Could not find Project")
	}

	if projectDoc.GetStatus() != PaidoutStatus {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Project not in PAIDOUT Status")

	}

	projectDid := withdrawFundsDoc.ProjectDid
	recipientDid := withdrawFundsDoc.RecipientDid
	amount := withdrawFundsDoc.Amount

	// If this is a refund, recipient has to be the project creator
	if withdrawFundsDoc.IsRefund && (recipientDid != projectDoc.GetSenderDid()) {

		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Only project creator can get a refund")
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
		return nil, err
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

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func payoutAndRecon(ctx sdk.Context, k Keeper, bk bank.Keeper, projectDid did.Did,
	fromAccountId InternalAccountID, recipientDid did.Did, amount sdk.Coin) error {

	ixoBalance := getIxoAmount(ctx, k, bk, projectDid, fromAccountId)
	if ixoBalance.IsLT(amount) {
		return sdkerrors.Wrap(types.ErrInternal, "insufficient funds in specified account")
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

func processFees(ctx sdk.Context, k Keeper, fk payments.Keeper, bk bank.Keeper,
	feeType payments.FeeType, projectDid did.Did) error {

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
	case payments.FeeClaimTransaction:
		adjustedFeeAmount = fk.GetParams(ctx).ClaimFeeAmount.Mul(ixoFactor)
	case payments.FeeEvaluationTransaction:
		adjustedFeeAmount = fk.GetParams(ctx).EvaluationFeeAmount.Mul(ixoFactor)
	default:
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Invalid Fee type.")
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

func processEvaluatorPay(ctx sdk.Context, k Keeper, fk payments.Keeper,
	bk bank.Keeper, projectDid, senderDid did.Did, evaluatorPay int64) error {

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

func createAccountInProjectAccounts(ctx sdk.Context, k Keeper, projectDid did.Did, accountId InternalAccountID) (sdk.AccAddress, error) {
	acc, err := k.CreateNewAccount(ctx, projectDid, accountId)
	if err != nil {
		return nil, err
	}

	k.AddAccountToProjectAccounts(ctx, projectDid, accountId, acc)

	return acc.GetAddress(), nil
}

func getAccountInProjectAccounts(ctx sdk.Context, k Keeper, projectDid did.Did,
	accountId InternalAccountID) (sdk.AccAddress, error) {
	accMap := k.GetAccountMap(ctx, projectDid)

	addr, found := accMap[accountId]
	if found {
		return addr, nil
	} else {
		return createAccountInProjectAccounts(ctx, k, projectDid, accountId)
	}
}

func getProjectAccount(ctx sdk.Context, k Keeper, projectDid did.Did) (sdk.AccAddress, error) {
	return getAccountInProjectAccounts(ctx, k, projectDid, InternalAccountID(projectDid))
}

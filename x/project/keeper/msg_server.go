package keeper

import (
	"context"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	ixotypes "github.com/ixofoundation/ixo-blockchain/lib/ixo"
	didexported "github.com/ixofoundation/ixo-blockchain/lib/legacydid"
	didtypes "github.com/ixofoundation/ixo-blockchain/lib/legacydid"
	iidtypes "github.com/ixofoundation/ixo-blockchain/x/iid/types"
	paymentskeeper "github.com/ixofoundation/ixo-blockchain/x/payments/keeper"
	paymentstypes "github.com/ixofoundation/ixo-blockchain/x/payments/types"
	"github.com/ixofoundation/ixo-blockchain/x/project/types"
)

type msgServer struct {
	Keeper         Keeper
	BankKeeper     bankkeeper.Keeper
	PaymentsKeeper paymentskeeper.Keeper
}

const (
	IxoAccountFeesId               types.InternalAccountID = "IxoFees"
	IxoAccountPayFeesId            types.InternalAccountID = "IxoPayFees"
	InitiatingNodeAccountPayFeesId types.InternalAccountID = "InitiatingNodePayFees"
)

// NewMsgServerImpl returns an implementation of the project MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(k Keeper, bk bankkeeper.Keeper, pk paymentskeeper.Keeper) types.MsgServer {
	return &msgServer{
		Keeper:         k,
		BankKeeper:     bk,
		PaymentsKeeper: pk,
	}
}

func (s msgServer) CreateProject(goCtx context.Context, msg *types.MsgCreateProject) (*types.MsgCreateProjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k := s.Keeper

	if k.ProjectDocExists(ctx, msg.ProjectDid) {
		return nil, sdkerrors.Wrap(didtypes.ErrInvalidDid, "project already exists")
	}

	// Create project doc
	projectDoc := types.NewProjectDoc(
		msg.TxHash, msg.ProjectDid, msg.SenderDid,
		msg.PubKey, types.NullStatus, msg.Data)

	// Get and validate project fees map
	err := k.ValidateProjectFeesMap(ctx, projectDoc.GetProjectFeesMap())
	if err != nil {
		return nil, err
	}

	// Create all necessary initial project accounts
	if _, err = createAccountInProjectAccounts(ctx, k, msg.ProjectDid, IxoAccountFeesId); err != nil {
		return nil, err
	}
	if _, err = createAccountInProjectAccounts(ctx, k, msg.ProjectDid, IxoAccountPayFeesId); err != nil {
		return nil, err
	}
	if _, err = createAccountInProjectAccounts(ctx, k, msg.ProjectDid, InitiatingNodeAccountPayFeesId); err != nil {
		return nil, err
	}
	if _, err = createAccountInProjectAccounts(ctx, k, msg.ProjectDid, types.InternalAccountID(msg.ProjectDid)); err != nil {
		return nil, err
	}

	iidProjectVerificationMethod, err := iidtypes.NewPublicKeyHexFromString(msg.PubKey, iidtypes.DIDVMethodTypeEd25519VerificationKey2018)
	if err != nil {
		return nil, err
	}

	//Create project backed IID
	did, err := iidtypes.NewDidDocument(
		msg.ProjectDid,
		iidtypes.WithControllers(msg.ProjectDid),
		iidtypes.WithVerifications(iidtypes.NewVerification(
			iidtypes.NewVerificationMethod(msg.ProjectDid, iidtypes.DID(msg.ProjectDid), iidProjectVerificationMethod),
			[]string{iidtypes.Authentication},
			nil,
		)),
	)
	if err != nil {
		// k.Logger(ctx).Error(err.Error())
		return nil, err
	}
	k.IidKeeper.SetDidDocument(ctx, []byte(msg.ProjectDid), did)

	// Set project doc and initialise list of withdrawal transactions
	k.SetProjectDoc(ctx, projectDoc)
	k.SetProjectWithdrawalTransactions(ctx, msg.ProjectDid, types.WithdrawalInfoDocs{})

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

	return &types.MsgCreateProjectResponse{}, nil
}

func (s msgServer) UpdateProjectStatus(goCtx context.Context, msg *types.MsgUpdateProjectStatus) (*types.MsgUpdateProjectStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k := s.Keeper
	bk := s.BankKeeper

	projectDoc, err := k.GetProjectDoc(ctx, msg.ProjectDid)
	if err != nil {
		return nil, sdkerrors.Wrap(didtypes.ErrInvalidDid, "could not find project")
	}

	newStatus := msg.Data.Status

	newStatusFromString := types.ProjectStatusFromString(newStatus)
	projectDocStatusFromString := types.ProjectStatusFromString(projectDoc.Status)

	if !newStatusFromString.IsValidProgressionFrom(projectDocStatusFromString) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest,
			"invalid Status Progression requested")
	}

	if types.ProjectStatusFromString(newStatus) == types.FundedStatus {
		projectAddr, err := getProjectAccount(ctx, k, projectDoc.ProjectDid)
		if err != nil {
			return nil, err
		}

		projectAcc := k.AccountKeeper.GetAccount(ctx, projectAddr)
		if projectAcc == nil {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest,
				"could not find project's account with address %s", projectAddr)
		}

		// Two conditions for minimum funding not reached:
		// - Either minimumFunding has some denom that is not in the projectAcc
		//   coins, indicating that the projectAcc has zero of this denom
		// - Or minimumFunding has some denom with a larger value than the projectAcc
		//   coins, indicating that the projectAcc has less than the minimum
		minimumFunding := k.GetParams(ctx).ProjectMinimumInitialFunding
		if !minimumFunding.DenomsSubsetOf(bk.GetAllBalances(ctx, projectAcc.GetAddress())) ||
			minimumFunding.IsAnyGT(bk.GetAllBalances(ctx, projectAcc.GetAddress())) {
			//if !minimumFunding.DenomsSubsetOf(projectAcc.GetCoins()) ||
			//	minimumFunding.IsAnyGT(projectAcc.GetCoins()) {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds,
				"project has not reached minimum funding %s", minimumFunding)
		}
	}

	if types.ProjectStatusFromString(newStatus) == types.PaidoutStatus {
		err := payoutFees(ctx, k, bk, projectDoc.ProjectDid)
		if err != nil {
			return nil, err
		}
	}

	projectDoc.Status = newStatus
	k.SetProjectDoc(ctx, projectDoc)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUpdateProjectStatus,
			sdk.NewAttribute(types.AttributeKeyTxHash, msg.TxHash),
			sdk.NewAttribute(types.AttributeKeySenderDid, msg.SenderDid),
			sdk.NewAttribute(types.AttributeKeyProjectDid, msg.ProjectDid),
			sdk.NewAttribute(types.AttributeKeyEthFundingTxnID, msg.Data.EthFundingTxnId),
			sdk.NewAttribute(types.AttributeKeyUpdatedStatus, fmt.Sprint(msg.Data.Status)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return &types.MsgUpdateProjectStatusResponse{}, nil
}

func payoutFees(ctx sdk.Context, k Keeper, bk bankkeeper.Keeper, projectDid didexported.Did) error {

	_, err := payAllFeesToAddress(ctx, k, bk, projectDid, IxoAccountPayFeesId, IxoAccountFeesId)
	if err != nil {
		return err
	}

	_, err = payAllFeesToAddress(ctx, k, bk, projectDid, InitiatingNodeAccountPayFeesId, IxoAccountFeesId)
	if err != nil {
		return err
	}

	ixoDid := k.GetParams(ctx).IxoDid
	amount := getIxoAmount(ctx, k, bk, projectDid, IxoAccountFeesId)
	err = payoutAndRecon(ctx, k, bk, projectDid, IxoAccountFeesId, ixoDid, amount)
	if err != nil {
		return err
	}

	return nil
}

func payAllFeesToAddress(ctx sdk.Context, k Keeper, bk bankkeeper.Keeper, projectDid didexported.Did,
	sendingAddress types.InternalAccountID, receivingAddress types.InternalAccountID) (*sdk.Result, error) {
	feesToPay := getIxoAmount(ctx, k, bk, projectDid, sendingAddress)

	if feesToPay.Amount.LT(sdk.ZeroInt()) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "negative fee to pay")
	} else if feesToPay.Amount.IsZero() {
		return nil, nil
	}

	receivingAccount, err := getAccountInProjectAccounts(ctx, k, projectDid, receivingAddress)
	if err != nil {
		return nil, err
	}

	sendingAccount, _ := getAccountInProjectAccounts(ctx, k, projectDid, sendingAddress)

	return nil, bk.SendCoins(ctx, sendingAccount, receivingAccount, sdk.Coins{feesToPay})
}

func getIxoAmount(ctx sdk.Context, k Keeper, bk bankkeeper.Keeper, projectDid didexported.Did, accountID types.InternalAccountID) sdk.Coin {
	found := checkAccountInProjectAccounts(ctx, k, projectDid, accountID)
	if found {
		accAddr, _ := getAccountInProjectAccounts(ctx, k, projectDid, accountID)
		coins := bk.GetAllBalances(ctx, accAddr)
		return sdk.NewCoin(ixotypes.IxoNativeToken, coins.AmountOf(ixotypes.IxoNativeToken))
	}
	return sdk.NewCoin(ixotypes.IxoNativeToken, sdk.ZeroInt())
}

func (s msgServer) CreateAgent(goCtx context.Context, msg *types.MsgCreateAgent) (*types.MsgCreateAgentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k := s.Keeper

	// Check if project exists
	_, err := k.GetProjectDoc(ctx, msg.ProjectDid)
	if err != nil {
		return nil, sdkerrors.Wrap(didtypes.ErrInvalidDid, "could not find project")
	}

	// Create account in project accounts for the agent
	_, err = createAccountInProjectAccounts(ctx, k, msg.ProjectDid, types.InternalAccountID(msg.Data.AgentDid))
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

	return &types.MsgCreateAgentResponse{}, nil
}

func (s msgServer) UpdateAgent(goCtx context.Context, msg *types.MsgUpdateAgent) (*types.MsgUpdateAgentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k := s.Keeper

	// Check if project exists
	_, err := k.GetProjectDoc(ctx, msg.ProjectDid)
	if err != nil {
		return nil, sdkerrors.Wrap(didtypes.ErrInvalidDid, "could not find project")
	}

	// TODO: implement agent update (or remove functionality)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUpdateAgent,
			sdk.NewAttribute(types.AttributeKeyTxHash, msg.TxHash),
			sdk.NewAttribute(types.AttributeKeySenderDid, msg.SenderDid),
			sdk.NewAttribute(types.AttributeKeyProjectDid, msg.ProjectDid),
			sdk.NewAttribute(types.AttributeKeyAgentDid, msg.Data.Did),
			sdk.NewAttribute(types.AttributeKeyAgentRole, msg.Data.Role),
			sdk.NewAttribute(types.AttributeKeyUpdatedStatus, msg.Data.Status),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return &types.MsgUpdateAgentResponse{}, nil
}

func (s msgServer) CreateClaim(goCtx context.Context, msg *types.MsgCreateClaim) (*types.MsgCreateClaimResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k := s.Keeper

	// Check if project exists
	projectDoc, err := k.GetProjectDoc(ctx, msg.ProjectDid)
	if err != nil {
		return nil, sdkerrors.Wrap(didtypes.ErrInvalidDid, "could not find project")
	}

	// Check that project status is STARTED
	if types.ProjectStatusFromString(projectDoc.Status) != types.StartedStatus {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "project not in STARTED status")
	}

	// Check if claim already exists
	if k.ClaimExists(ctx, msg.ProjectDid, msg.Data.ClaimId) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "claim already exists")
	}

	// Create and set claim
	claim := types.NewClaim(msg.Data.ClaimId, msg.Data.ClaimTemplateId, msg.SenderDid)
	k.SetClaim(ctx, msg.ProjectDid, claim)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateClaim,
			sdk.NewAttribute(types.AttributeKeyTxHash, msg.TxHash),
			sdk.NewAttribute(types.AttributeKeySenderDid, msg.SenderDid),
			sdk.NewAttribute(types.AttributeKeyProjectDid, msg.ProjectDid),
			sdk.NewAttribute(types.AttributeKeyClaimID, msg.Data.ClaimId),
			sdk.NewAttribute(types.AttributeKeyClaimTemplateID, msg.Data.ClaimTemplateId),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return &types.MsgCreateClaimResponse{}, nil
}

func (s msgServer) CreateEvaluation(goCtx context.Context, msg *types.MsgCreateEvaluation) (*types.MsgCreateEvaluationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k := s.Keeper
	bk := s.BankKeeper
	pk := s.PaymentsKeeper

	// Check if project exists
	projectDoc, err := k.GetProjectDoc(ctx, msg.ProjectDid)
	if err != nil {
		return nil, sdkerrors.Wrap(didtypes.ErrInvalidDid, "could not find project")
	}

	// Check that project status is STARTED
	if types.ProjectStatusFromString(projectDoc.Status) != types.StartedStatus {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "project not in STARTED status")
	}

	// Get claim and confirm status is pending
	claim, err := k.GetClaim(ctx, msg.ProjectDid, msg.Data.ClaimId)
	if err != nil {
		return nil, err
	} else if claim.Status != string(types.PendingClaim) {
		return nil, fmt.Errorf("claim status must be pending")
	}

	// Get project fees map
	feesMap := projectDoc.GetProjectFeesMap()

	// If oracle fee present in project fees map, proceed with oracle pay
	templateId, err := feesMap.GetPayTemplateId(types.OracleFee)
	if err == nil {
		// Get ixo address
		ixoAddr, err := getAccountInProjectAccounts(ctx, k, msg.ProjectDid,
			IxoAccountPayFeesId)
		if err != nil {
			return nil, err
		}

		// Get node (relayer) address
		nodeAddr, err := getAccountInProjectAccounts(ctx, k, msg.ProjectDid,
			InitiatingNodeAccountPayFeesId)
		if err != nil {
			return nil, err
		}

		// Get sender (oracle) address
		senderDidDoc, exists := k.IidKeeper.GetDidDocument(ctx, []byte(msg.SenderDid))
		if !exists {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "signer must be payment contract payer")
		}
		senderAddr, err := senderDidDoc.GetVerificationMethodBlockchainAddress(senderDidDoc.Id)
		if err != nil {
			return nil, sdkerrors.Wrap(err, "Address not found in iid doc")
		}
		// Calculate evaluator pay share (totals to 100) for ixo, node, and oracle
		feePercentage := k.GetParams(ctx).OracleFeePercentage
		nodeFeeShare := feePercentage.Mul(k.GetParams(ctx).NodeFeePercentage.QuoInt64(100))
		ixoFeeShare := feePercentage.Sub(nodeFeeShare)
		oracleShareLessFees := sdk.NewDec(100).Sub(feePercentage)
		oraclePayRecipients := paymentstypes.NewDistribution(
			paymentstypes.NewDistributionShare(ixoAddr, ixoFeeShare),
			paymentstypes.NewDistributionShare(nodeAddr, nodeFeeShare),
			paymentstypes.NewDistributionShare(senderAddr, oracleShareLessFees))

		// Process oracle pay
		err = processPay(ctx, k, bk, pk, msg.ProjectDid, senderAddr,
			oraclePayRecipients, types.OracleFee, templateId)
		if err != nil {
			return nil, err
		}
	}

	// If fee for service present in project fees map and if
	// claim approved, proceed with fee-for-service payment
	templateId, err = feesMap.GetPayTemplateId(types.FeeForService)
	if err == nil && msg.Data.Status == string(types.ApprovedClaim) {
		// Get claimer address
		claimerDidDoc, exists := k.IidKeeper.GetDidDocument(ctx, []byte(claim.ClaimerDid))
		if !exists {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "signer must be payment contract payer")
		}
		claimerAddr, err := claimerDidDoc.GetVerificationMethodBlockchainAddress(claimerDidDoc.Id)
		if err != nil {
			return nil, sdkerrors.Wrap(err, "Address not found in iid doc")
		}

		// Get recipients (just the claimer)
		claimApprovedPayRecipients := paymentstypes.NewDistribution(
			paymentstypes.NewFullDistributionShare(claimerAddr))

		// Process the payment
		err = processPay(ctx, k, bk, pk, projectDoc.ProjectDid, claimerAddr,
			claimApprovedPayRecipients, types.FeeForService, templateId)
		if err != nil {
			return nil, err
		}
	}

	// Update and set claim
	claim.Status = msg.Data.Status
	k.SetClaim(ctx, msg.ProjectDid, claim)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateEvaluation,
			sdk.NewAttribute(types.AttributeKeyTxHash, msg.TxHash),
			sdk.NewAttribute(types.AttributeKeySenderDid, msg.SenderDid),
			sdk.NewAttribute(types.AttributeKeyProjectDid, msg.ProjectDid),
			sdk.NewAttribute(types.AttributeKeyClaimID, msg.Data.ClaimId),
			sdk.NewAttribute(types.AttributeKeyClaimStatus, fmt.Sprint(msg.Data.Status)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return &types.MsgCreateEvaluationResponse{}, nil
}

func (s msgServer) WithdrawFunds(goCtx context.Context, msg *types.MsgWithdrawFunds) (*types.MsgWithdrawFundsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k := s.Keeper
	bk := s.BankKeeper

	withdrawFundsDoc := msg.Data
	projectDoc, err := k.GetProjectDoc(ctx, withdrawFundsDoc.ProjectDid)
	if err != nil {
		return nil, sdkerrors.Wrap(didtypes.ErrInvalidDid, "could not find project")
	}

	//if projectDoc.Status != string(types.PaidoutStatus) {
	//	return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized,
	//		"project not in PAIDOUT status")
	//}

	projectDid := withdrawFundsDoc.ProjectDid
	recipientDid := withdrawFundsDoc.RecipientDid
	amount := withdrawFundsDoc.Amount

	// If this is a refund, recipient has to be the project creator
	if withdrawFundsDoc.IsRefund && (recipientDid != projectDoc.SenderDid) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized,
			"only project creator can get a refund")
	}

	var fromAccountId types.InternalAccountID
	if withdrawFundsDoc.IsRefund {
		fromAccountId = types.InternalAccountID(projectDid)
	} else {
		fromAccountId = types.InternalAccountID(recipientDid)
	}

	amountCoin := sdk.NewCoin(ixotypes.IxoNativeToken, amount)
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

	return &types.MsgWithdrawFundsResponse{}, nil
}

func (s msgServer) UpdateProjectDoc(goCtx context.Context, msg *types.MsgUpdateProjectDoc) (*types.MsgUpdateProjectDocResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k := s.Keeper

	projectDoc, err := k.GetProjectDoc(ctx, msg.ProjectDid)
	if err != nil {
		return nil, sdkerrors.Wrap(didtypes.ErrInvalidDid, "could not find project")
	}

	// Get and validate project fees map
	err = k.ValidateProjectFeesMap(ctx, projectDoc.GetProjectFeesMap())
	if err != nil {
		return nil, err
	}

	// Editor of project doc has to be the same as project creator
	if msg.SenderDid != projectDoc.SenderDid {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized,
			"only project creator can edit project doc")
	}

	// Project doc can be updated when in states Null, Created, Pending, or Funded
	// and cannot be edited when in states Started, Stopped, and Paidout
	status := types.ProjectStatusFromString(projectDoc.Status)
	if status == types.StartedStatus || status == types.StoppedStatus || status == types.PaidoutStatus {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest,
			"project doc cannot be updated when project is in status %s", projectDoc.Status)
	}

	projectDoc.Data = msg.Data
	k.SetProjectDoc(ctx, projectDoc)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUpdateProjectDoc,
			sdk.NewAttribute(types.AttributeKeyTxHash, msg.TxHash),
			sdk.NewAttribute(types.AttributeKeySenderDid, msg.SenderDid),
			sdk.NewAttribute(types.AttributeKeyProjectDid, msg.ProjectDid),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return &types.MsgUpdateProjectDocResponse{}, nil
}

func payoutAndRecon(ctx sdk.Context, k Keeper, bk bankkeeper.Keeper, projectDid didexported.Did,
	fromAccountId types.InternalAccountID, recipientDid didexported.Did, amount sdk.Coin) error {

	if amount.IsZero() {
		return nil
	}

	ixoBalance := getIxoAmount(ctx, k, bk, projectDid, fromAccountId)
	if ixoBalance.IsLT(amount) {
		return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "insufficient funds in specified account")
	}

	fromAccount, err := getAccountInProjectAccounts(ctx, k, projectDid, fromAccountId)
	if err != nil {
		return err
	}

	// Get recipient address
	recipientDidDoc, exists := k.IidKeeper.GetDidDocument(ctx, []byte(recipientDid))
	if !exists {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "signer must be payment contract payer")
	}
	recipientAddr, err := recipientDidDoc.GetVerificationMethodBlockchainAddress(recipientDidDoc.Id)
	if err != nil {
		return sdkerrors.Wrap(err, "Address not found in iid doc")
	}

	err = bk.SendCoins(ctx, fromAccount, recipientAddr, sdk.Coins{amount})
	if err != nil {
		return err
	}

	addProjectWithdrawalTransaction(ctx, k, projectDid, recipientDid, amount)
	return nil
}

func processPay(ctx sdk.Context, k Keeper, bk bankkeeper.Keeper, pk paymentskeeper.Keeper,
	projectDid didexported.Did, senderAddr sdk.AccAddress, recipients paymentstypes.Distribution,
	feeType types.FeeType, paymentTemplateId string) error {

	// Validate recipients
	err := recipients.Validate()
	if err != nil {
		return err
	}

	// Get project address
	projectAddr, err := getProjectAccount(ctx, k, projectDid)
	if err != nil {
		return err
	}

	// Get payment template
	template := pk.MustGetPaymentTemplate(ctx, paymentTemplateId)

	// Create or get payment contract
	contractId := fmt.Sprintf("payment:contract:%s:%s:%s:%s",
		types.ModuleName, projectDid, senderAddr.String(), feeType)
	var contract paymentstypes.PaymentContract
	if !pk.PaymentContractExists(ctx, contractId) {
		contract = paymentstypes.NewPaymentContract(contractId, paymentTemplateId,
			projectAddr, projectAddr, recipients, false, true, sdk.ZeroUint())
		pk.SetPaymentContract(ctx, contract)
	} else {
		contract = pk.MustGetPaymentContract(ctx, contractId)
	}

	// Effect payment if can effect
	if contract.CanEffectPayment(template) {
		// Check that project has enough tokens to effect contract payment
		// (assume no effect from PaymentMin, PaymentMax, Discounts)
		for _, coin := range template.PaymentAmount {
			if !bk.HasBalance(ctx, projectAddr, coin) {
				return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "project has insufficient funds")
			}
		}

		// Effect payment
		effected, err := pk.EffectPayment(ctx, bk, contractId)
		if err != nil {
			return err
		} else if !effected {
			panic("expected to be able to effect contract payment")
		}
	} else {
		return fmt.Errorf("cannot effect contract payment (max reached?)")
	}

	return nil
}

func checkAccountInProjectAccounts(ctx sdk.Context, k Keeper, projectDid didexported.Did,
	accountId types.InternalAccountID) bool {
	strAccountId := string(accountId)
	accMap := k.GetAccountMap(ctx, projectDid)
	_, found := accMap.Map[strAccountId]

	return found
}

func addProjectWithdrawalTransaction(ctx sdk.Context, k Keeper,
	projectDid didexported.Did, recipientDid didexported.Did, amount sdk.Coin) {

	withdrawalInfo := types.WithdrawalInfoDoc{
		ProjectDid:   projectDid,
		RecipientDid: recipientDid,
		Amount:       amount,
	}

	k.AddProjectWithdrawalTransaction(ctx, projectDid, withdrawalInfo)
}

func createAccountInProjectAccounts(ctx sdk.Context, k Keeper, projectDid didexported.Did,
	accountId types.InternalAccountID) (sdk.AccAddress, error) {
	acc, err := k.CreateNewAccount(ctx, projectDid, accountId)
	if err != nil {
		return nil, err
	}

	k.AddAccountToProjectAccounts(ctx, projectDid, accountId, acc)

	return acc.GetAddress(), nil
}

func getAccountInProjectAccounts(ctx sdk.Context, k Keeper, projectDid didexported.Did,
	accountId types.InternalAccountID) (sdk.AccAddress, error) {
	strAccountId := string(accountId)
	accMap := k.GetAccountMap(ctx, projectDid)

	addr, found := accMap.Map[strAccountId]
	if found {
		return sdk.AccAddressFromBech32(addr)
	} else {
		return createAccountInProjectAccounts(ctx, k, projectDid, accountId)
	}
}

func getProjectAccount(ctx sdk.Context, k Keeper, projectDid didexported.Did) (sdk.AccAddress, error) {
	return getAccountInProjectAccounts(ctx, k, projectDid, types.InternalAccountID(projectDid))
}

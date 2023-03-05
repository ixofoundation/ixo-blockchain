package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/ixofoundation/ixo-blockchain/x/claims/types"
)

// --------------------------
// PAYMENT HELPERS
// --------------------------

func processPayment(ctx sdk.Context, k Keeper, receiver sdk.AccAddress, payment *types.Payment, paymentType types.PaymentType, claimId string) error {
	// check that there is outcome payment to make, otherwise skip this with no error as no payment for action
	paymentExists := false
	for _, coin := range payment.Amount {
		if !coin.IsZero() {
			paymentExists = true
			break
		}
	}
	if !paymentExists {
		return nil
	}

	// Get payer address
	payerAddress, err := sdk.AccAddressFromBech32(payment.Account)
	if err != nil {
		return err
	}

	// Not sure if needed as authz can still be created even if sender has no fees
	// check that sender has enough tokens to make the payment
	// if !types.HasBalances(ctx, k.BankKeeper, payerAddress, payment.Amount) {
	// 	return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "sender has insufficient funds")
	// }

	// get claim and collection payment is for
	claim, err := k.GetClaim(ctx, claimId)
	if err != nil {
		return err
	}
	collection, err := k.GetCollection(ctx, claim.CollectionId)
	if err != nil {
		return err
	}

	var outputs []banktypes.Output

	// if evaluation payment then do split based on params
	if paymentType == types.PaymentType_evaluation {
		params := k.GetParams(ctx)

		// Get node address
		entity, exists := k.EntityKeeper.GetEntity(ctx, []byte(collection.Entity))
		if !exists {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "entity did doesn't exist for did %s", collection.Entity)
		}
		relayerDidDoc, exists := k.IidKeeper.GetDidDocument(ctx, []byte(entity.RelayerNode))
		if !exists {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "relayer node did doesn't exist for did %s", entity.RelayerNode)
		}
		relayerAddr, err := relayerDidDoc.GetVerificationMethodBlockchainAddress(entity.RelayerNode)
		if err != nil {
			return sdkerrors.Wrapf(err, "address not found in iid doc for %s", entity.RelayerNode)
		}

		// Calculate evaluator pay share (totals to 100) for ixo, node, and oracle
		nodeFeePercentage := params.NodeFeePercentage
		ixoFeePercentage := params.NetworkFeePercentage
		// check that the 2 preset percentages dont go over 100%
		if nodeFeePercentage.Add(ixoFeePercentage).GT(types.OneHundred) {
			return types.ErrPaymentPresetPercentagesOverflow
		}
		oracleFeePercentage := types.OneHundred.Sub(nodeFeePercentage).Sub(ixoFeePercentage)

		// Get ixo netowrk address
		ixoAddress, err := sdk.AccAddressFromBech32(params.IxoAccount)
		if err != nil {
			return err
		}

		recipients := types.NewDistribution(
			types.NewDistributionShare(ixoAddress, ixoFeePercentage),
			types.NewDistributionShare(relayerAddr, nodeFeePercentage),
			types.NewDistributionShare(receiver, oracleFeePercentage))

		// Calculate list of outputs and calculate the total output to payees based
		// on the calculated wallet distributions
		distributions := recipients.GetDistributionsFor(payment.Amount)
		for i, share := range distributions {
			// Get integer output
			outputAmt, _ := share.TruncateDecimal()

			// If amount not zero, add as output
			if !outputAmt.IsZero() {
				address, err := recipients[i].GetAddress()
				if err != nil {
					return err
				}
				outputs = append(outputs, banktypes.NewOutput(address, outputAmt))
			}
		}
	} else {
		// if no split then recipient gets 100% payment
		outputs = append(outputs, banktypes.NewOutput(receiver, payment.Amount))
	}

	inputs := []banktypes.Input{banktypes.NewInput(payerAddress, payment.Amount)}

	// if no timout in payment make payout immidiately
	if payment.TimeoutNs == 0 {
		if err := payout(ctx, k, inputs, outputs, paymentType, claimId, &time.Time{}); err != nil {
			return err
		}
	} else {
		// Get admin address
		adminAddress, err := sdk.AccAddressFromBech32(collection.Admin)
		if err != nil {
			return err
		}

		// else create authz WithdrawPaymentAuthorization for receiver to execute to receive payout once timout has passed
		if err := createAuthz(ctx, k, receiver, adminAddress, inputs, outputs, paymentType, claimId, payment.TimeoutNs); err != nil {
			return err
		}
	}

	return nil
}

func payout(ctx sdk.Context, k Keeper, inputs []banktypes.Input, outputs []banktypes.Output, paymentType types.PaymentType, claimId string, releaseDate *time.Time) error {
	// distribute the payment according to the outputs
	err := k.BankKeeper.InputOutputCoins(ctx, inputs, outputs)
	if err != nil {
		// update payment status to failed
		updatePaymentStatus(ctx, k, paymentType, claimId, types.PaymentStatus_failed)
		return sdkerrors.Wrapf(types.ErrPaymentWithdrawFailed, "%s", err)
	}

	// update payment status to success
	updatePaymentStatus(ctx, k, paymentType, claimId, types.PaymentStatus_paid)

	// emit the events
	if err := ctx.EventManager().EmitTypedEvents(
		&types.PaymentWithdrawnEvent{
			Withdraw: &types.WithdrawPaymentConstraints{
				ClaimId:     claimId,
				Inputs:      inputs,
				Outputs:     outputs,
				PaymentType: paymentType,
				ReleaseDate: releaseDate,
			},
		},
	); err != nil {
		return err
	}
	return nil
}

func createAuthz(ctx sdk.Context, k Keeper, receiver, admin sdk.AccAddress, inputs []banktypes.Input, outputs []banktypes.Output, paymentType types.PaymentType, claimId string, timeoutNs time.Duration) error {
	// get users current WithdrawPaymentAuthorization authorization
	authzMsgType := sdk.MsgTypeURL(&types.MsgWithdrawPayment{})
	auth, _ := k.AuthzKeeper.GetCleanAuthorization(ctx, receiver, admin, authzMsgType)

	// making expiration date for authz grant one year from now (until indefinite time fix for cosmos version)
	expiration := ctx.BlockTime().Add(time.Hour * 24 * 365)
	releaseDate := ctx.BlockTime().Add(timeoutNs)
	var constraints []*types.WithdrawPaymentConstraints
	constraint := types.WithdrawPaymentConstraints{
		ClaimId:     claimId,
		Inputs:      inputs,
		Outputs:     outputs,
		PaymentType: paymentType,
		ReleaseDate: &releaseDate,
	}

	// if have a WithdrawPaymentAuthorization authz use current constaints to append new one to
	if auth != nil {
		switch k := auth.(type) {
		case *types.WithdrawPaymentAuthorization:
			constraints = k.Constraints
		default:
			return fmt.Errorf("existing Authorizations for route %s is not of type WithdrawPaymentAuthorization", authzMsgType)
		}
	}
	constraints = append(constraints, &constraint)

	// persist new grant
	if err := k.AuthzKeeper.SaveGrant(ctx, receiver, admin, types.NewWithdrawPaymentAuthorization(admin.String(), constraints), expiration); err != nil {
		return err
	}

	// update payment status to authorized
	updatePaymentStatus(ctx, k, paymentType, claimId, types.PaymentStatus_authorized)

	// emit the events
	if err := ctx.EventManager().EmitTypedEvents(
		&types.PaymentWithdrawCreatedEvent{
			Withdraw: &types.WithdrawPaymentConstraints{
				ClaimId:     claimId,
				Inputs:      inputs,
				Outputs:     outputs,
				PaymentType: paymentType,
				ReleaseDate: &releaseDate,
			},
		},
	); err != nil {
		return err
	}

	return nil
}

func updatePaymentStatus(ctx sdk.Context, k Keeper, paymentType types.PaymentType, claimId string, paymentStatus types.PaymentStatus) error {
	// get claim and payment is for
	claim, err := k.GetClaim(ctx, claimId)
	if err != nil {
		return err
	}

	switch paymentType {
	case types.PaymentType_approval:
		claim.PaymentsStatus.Approval = paymentStatus
	case types.PaymentType_evaluation:
		claim.PaymentsStatus.Evaluation = paymentStatus
	case types.PaymentType_submission:
		claim.PaymentsStatus.Submission = paymentStatus
	}

	// persist claim changes
	k.SetClaim(ctx, claim)

	// emit the events
	if err := ctx.EventManager().EmitTypedEvents(
		&types.ClaimUpdatedEvent{
			Claim: &claim,
		},
	); err != nil {
		return err
	}
	return nil
}

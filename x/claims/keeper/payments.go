package keeper

import (
	"fmt"
	"time"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/ixofoundation/ixo-blockchain/v3/x/claims/types"
	"github.com/ixofoundation/ixo-blockchain/v3/x/token/types/contracts/cw20"
	"github.com/ixofoundation/ixo-blockchain/v3/x/token/types/contracts/ixo1155"
)

// --------------------------
// PAYMENT HELPERS
// --------------------------

// processPayment processes payments for the different types of payments.
// inside the function, it checks that a payment exists (native coin payment), and if it exists and is not 0, it does the following:
//  1. if paymentType is evaluation, it calculates the distribution ratio according to network, node and oracle fees set in the params
//     we only calculate the distribution ratio if amount is not 0 since all evaluation payments must be native coin payments
//  2. if any other paymentType, it sets the output for payment to the amount set in the payment
//
// after the native coin payments is calculated, it does the following:
// 1. if timeout for the payment is nil, it makes the payment by calling payout() function
// 2. if timeout for the payment is not nil, it creates an authz grant to make the payment by calling createAuthz() function
func processPayment(ctx sdk.Context, k Keeper, receiver sdk.AccAddress, payment *types.Payment, paymentType types.PaymentType, claimId string) error {
	// check that there is outcome payment to make, otherwise skip this with no error as no payment for action
	paymentExists := false
	if !payment.Amount.IsZero() {
		paymentExists = true
	}
	if payment.Contract_1155Payment != nil && payment.Contract_1155Payment.Amount != 0 {
		paymentExists = true
	}
	if payment.Cw20Payment != nil && len(payment.Cw20Payment) > 0 {
		// map to check if there is any cw20 payments
		for _, cw20Payment := range payment.Cw20Payment {
			if cw20Payment.Amount != 0 {
				paymentExists = true
				break
			}
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
	// 	return errorsmod.Wrapf(sdkerrors.ErrInsufficientFunds, "sender has insufficient funds")
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

	var inputs []banktypes.Input
	var outputs []banktypes.Output

	// if there is cosmos coins in payment get input/outputs for multisend
	if !payment.Amount.IsZero() {
		// if evaluation payment then do split based on params
		if paymentType == types.PaymentType_evaluation {
			params := k.GetParams(ctx)

			// Get node address
			entity, exists := k.EntityKeeper.GetEntity(ctx, []byte(collection.Entity))
			if !exists {
				return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "entity did doesn't exist for did %s", collection.Entity)
			}
			relayerDidDoc, exists := k.IidKeeper.GetDidDocument(ctx, []byte(entity.RelayerNode))
			if !exists {
				return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "relayer node did doesn't exist for did %s", entity.RelayerNode)
			}
			relayerAddr, err := relayerDidDoc.GetVerificationMethodBlockchainAddress(entity.RelayerNode)
			if err != nil {
				return errorsmod.Wrapf(err, "did not found in iid doc verification methods for %s", entity.RelayerNode)
			}

			// Calculate evaluator pay share (totals to 100) for ixo, node, and oracle
			nodeFeePercentage := params.NodeFeePercentage
			ixoFeePercentage := params.NetworkFeePercentage
			// check that the 2 preset percentages dont go over 100%
			if nodeFeePercentage.Add(ixoFeePercentage).GT(types.OneHundred) {
				return types.ErrPaymentPresetPercentagesOverflow
			}
			oracleFeePercentage := types.OneHundred.Sub(nodeFeePercentage).Sub(ixoFeePercentage)

			// Get ixo network address
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
			distributions, err := recipients.GetDistributionsFor(payment.Amount)
			if err != nil {
				return err
			}

			var countOutputs sdk.Coins
			for i, share := range distributions {
				// Get integer output
				outputAmt, _ := share.TruncateDecimal()
				address, err := recipients[i].GetAddress()
				if err != nil {
					return err
				}

				// If receiver address(last address in the distribution), then add the remainder to the receiver
				if address.Equals(receiver) {
					outputAmt = payment.Amount.Sub(countOutputs...)
					outputs = append(outputs, banktypes.NewOutput(address, outputAmt))
				} else if !outputAmt.IsZero() {
					// If amount not zero, add as output, for network and node
					outputs = append(outputs, banktypes.NewOutput(address, outputAmt))
					countOutputs = countOutputs.Sort().Add(outputAmt.Sort()...)
				}
			}
		} else {
			// if no split then recipient gets 100% payment
			outputs = append(outputs, banktypes.NewOutput(receiver, payment.Amount))
		}

		inputs = append(inputs, banktypes.NewInput(payerAddress, payment.Amount))
	}

	// if no timeout in payment make payout immediately
	if payment.TimeoutNs == 0 {
		if err := payout(ctx, k, inputs, outputs, paymentType, claimId, &time.Time{}, payment.Contract_1155Payment, payment.Cw20Payment, payerAddress, receiver); err != nil {
			return err
		}
	} else {
		// get admin address
		adminAddress, err := sdk.AccAddressFromBech32(collection.Admin)
		if err != nil {
			return err
		}

		// else create authz WithdrawPaymentAuthorization for receiver to execute to receive payout once timeout has passed
		if err := createAuthz(ctx, k, receiver, adminAddress, inputs, outputs, paymentType, claimId, payment.TimeoutNs, payment.Contract_1155Payment, payment.Cw20Payment, payerAddress, receiver); err != nil {
			return err
		}
	}

	return nil
}

// payout makes the payment for the different paymentType.
// it does the following:
// 1. validate if the from address and input's addresses is valid entity module account
// 2. if inputs and outputs are not empty, it makes the payment by calling k.BankKeeper.InputOutputCoins()
// 3. if payment1155 is not nil, it makes the payment by calling k.WasmKeeper.Execute()
// 4. if paymentCw20s is not nil, it makes the payments by calling k.WasmKeeper.Execute()
// 5. update the claim payment status to success
// 6. emit PaymentWithdrawnEvent event
func payout(ctx sdk.Context, k Keeper, inputs []banktypes.Input, outputs []banktypes.Output, paymentType types.PaymentType, claimId string, releaseDate *time.Time, payment1155 *types.Contract1155Payment, paymentCw20s []*types.CW20Payment, fromAddress, toAddress sdk.AccAddress) error {
	// get entity payout is for to validate if from address is valid entity module account
	claim, err := k.GetClaim(ctx, claimId)
	if err != nil {
		return err
	}
	collection, err := k.GetCollection(ctx, claim.CollectionId)
	if err != nil {
		return err
	}
	_, entity, err := k.EntityKeeper.ResolveEntity(ctx, collection.Entity)
	if err != nil {
		return err
	}

	// check that fromAddress and input addresses is entity module accounts
	if !entity.ContainsAccountAddress(fromAddress.String()) {
		return types.ErrCollNotEntityAcc
	}
	for _, i := range inputs {
		if !entity.ContainsAccountAddress(i.Address) {
			return types.ErrCollNotEntityAcc
		}
	}

	// distribute the payment according to the outputs for Cosmos Coins if has inputs and outputs
	if len(inputs) != 0 && len(outputs) != 0 {
		if err := k.BankKeeper.InputOutputCoins(ctx, inputs[0], outputs); err != nil {
			return errorsmod.Wrapf(types.ErrPaymentWithdrawFailed, "%s", err)
		}
	}

	// pay 1155 contract payment if has one
	if payment1155 != nil && payment1155.Amount != 0 {
		encodedTransferMessage, err := ixo1155.Marshal(ixo1155.WasmSendFrom{
			SendFrom: ixo1155.SendFrom{
				From:     fromAddress.String(),
				To:       toAddress.String(),
				Token_id: payment1155.TokenId,
				Value:    fmt.Sprint(payment1155.Amount),
			},
		})
		if err != nil {
			return err
		}

		contractAddress, err := sdk.AccAddressFromBech32(payment1155.Address)
		if err != nil {
			return err
		}

		_, err = k.WasmKeeper.Execute(
			ctx,
			contractAddress,
			fromAddress,
			encodedTransferMessage,
			sdk.NewCoins(sdk.NewCoin("uixo", math.ZeroInt())),
		)
		if err != nil {
			return err
		}
	}

	// pay cw20 payments if has any
	if len(paymentCw20s) > 0 {
		for _, cw20Payment := range paymentCw20s {
			// make the payments if amount is not 0
			if cw20Payment.Amount != 0 {
				encodedTransferMessage, err := cw20.Marshal(cw20.WasmTransferFrom{
					TransferFrom: cw20.TransferFrom{
						Owner:     fromAddress.String(),
						Recipient: toAddress.String(),
						Amount:    fmt.Sprint(cw20Payment.Amount),
					},
				})
				if err != nil {
					return err
				}

				contractAddress, err := sdk.AccAddressFromBech32(cw20Payment.Address)
				if err != nil {
					return err
				}

				_, err = k.WasmKeeper.Execute(
					ctx,
					contractAddress,
					fromAddress,
					encodedTransferMessage,
					sdk.NewCoins(sdk.NewCoin("uixo", math.ZeroInt())),
				)
				if err != nil {
					return err
				}
			}
		}
	}

	// update payment status to success
	if err := updatePaymentStatus(ctx, k, paymentType, claimId, types.PaymentStatus_paid); err != nil {
		return err
	}

	// emit the events
	if err := ctx.EventManager().EmitTypedEvents(
		&types.PaymentWithdrawnEvent{
			Withdraw: &types.WithdrawPaymentConstraints{
				ClaimId:              claimId,
				Inputs:               inputs,
				Outputs:              outputs,
				PaymentType:          paymentType,
				ReleaseDate:          releaseDate,
				Contract_1155Payment: payment1155,
				Cw20Payment:          paymentCw20s,
				FromAddress:          fromAddress.String(),
				ToAddress:            toAddress.String(),
			},
		},
	); err != nil {
		return err
	}
	return nil
}

// createAuthz creates an authz grant for the different paymentType if timeoutNs is not nil.
// it does the following:
// 1. get user's current WithdrawPaymentAuthorization authorization
// 2. create and add the new WithdrawPaymentConstraints to the current existing constraints, and persist
// 3. update the payment status to authorized
// 4. emit PaymentWithdrawCreatedEvent event
func createAuthz(ctx sdk.Context, k Keeper, receiver, admin sdk.AccAddress, inputs []banktypes.Input, outputs []banktypes.Output, paymentType types.PaymentType, claimId string, timeoutNs time.Duration, payment1155 *types.Contract1155Payment, paymentCw20s []*types.CW20Payment, fromAddress, toAddress sdk.AccAddress) error {
	// get user's current WithdrawPaymentAuthorization authorization
	authzMsgType := sdk.MsgTypeURL(&types.MsgWithdrawPayment{})
	auth, _ := k.AuthzKeeper.GetAuthorization(ctx, receiver, admin, authzMsgType)

	releaseDate := ctx.BlockTime().Add(timeoutNs)
	var constraints []*types.WithdrawPaymentConstraints
	constraint := types.WithdrawPaymentConstraints{
		ClaimId:              claimId,
		Inputs:               inputs,
		Outputs:              outputs,
		PaymentType:          paymentType,
		ReleaseDate:          &releaseDate,
		Contract_1155Payment: payment1155,
		Cw20Payment:          paymentCw20s,
		FromAddress:          fromAddress.String(),
		ToAddress:            toAddress.String(),
	}

	// if have a WithdrawPaymentAuthorization authz use current constraints to append new one to
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
	if err := k.AuthzKeeper.SaveGrant(ctx, receiver, admin, types.NewWithdrawPaymentAuthorization(admin.String(), constraints), nil); err != nil {
		return err
	}

	// update payment status to authorized
	if err := updatePaymentStatus(ctx, k, paymentType, claimId, types.PaymentStatus_authorized); err != nil {
		return err
	}

	// emit the events
	if err := ctx.EventManager().EmitTypedEvents(
		&types.PaymentWithdrawCreatedEvent{
			Withdraw: &types.WithdrawPaymentConstraints{
				ClaimId:              claimId,
				Inputs:               inputs,
				Outputs:              outputs,
				PaymentType:          paymentType,
				ReleaseDate:          &releaseDate,
				Contract_1155Payment: payment1155,
				Cw20Payment:          paymentCw20s,
				FromAddress:          fromAddress.String(),
				ToAddress:            toAddress.String(),
			},
		},
	); err != nil {
		return err
	}

	return nil
}

// updatePaymentStatus updates the payment status for the provided different paymentType.
// it does the following:
// 1. get claim and payment is for
// 2. update the payment status to the provided paymentStatus
// 3. persist the updated claim
// 4. emit PaymentWithdrawnEvent event
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

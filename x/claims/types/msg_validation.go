package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	iidtypes "github.com/ixofoundation/ixo-blockchain/x/iid/types"
)

// --------------------------
// CREATE COLLECTION
// --------------------------
func (msg MsgCreateCollection) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}

	if !iidtypes.IsValidDID(msg.Entity) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.Entity)
	}

	if !iidtypes.IsValidDID(msg.Protocol) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.Protocol)
	}

	_, err = sdk.AccAddressFromBech32(msg.Payments.Submission.Account)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid payments submission address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.Payments.Evaluation.Account)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid payments evaluation address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.Payments.Approval.Account)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid payments approval address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.Payments.Rejection.Account)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid payments rejection address (%s)", err)
	}

	return nil
}

// --------------------------
// SUBMIT CLAIM
// --------------------------
func (msg MsgSubmitClaim) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}

	if !iidtypes.IsValidDID(msg.AgentDid.Did()) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.AgentDid.String())
	}

	if iidtypes.IsEmpty(msg.ClaimId) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "claim_id cannot be empty")
	}

	return nil
}

// --------------------------
// EVALUATE CLAIM
// --------------------------
func (msg MsgEvaluateClaim) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}

	if !iidtypes.IsValidDID(msg.AgentDid.Did()) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.AgentDid.String())
	}

	if !iidtypes.IsValidDID(msg.Oracle) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.AgentDid.String())
	}

	if iidtypes.IsEmpty(msg.ClaimId) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "claim_id cannot be empty")
	}

	if iidtypes.IsEmpty(msg.VerificationProof) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "verification_proof cannot be empty")
	}

	if msg.Status == EvaluationStatus_pending {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "evaluation status can't be pending")
	}

	return nil
}

// --------------------------
// DISPUTE CLAIM
// --------------------------
func (msg MsgDisputeClaim) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AgentAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid agent address (%s)", err)
	}

	if !iidtypes.IsValidDID(msg.AgentDid.Did()) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.AgentDid.String())
	}

	if iidtypes.IsEmpty(msg.Data.Proof) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "dispute data proof cannot be empty")
	}
	if iidtypes.IsEmpty(msg.Data.Uri) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "dispute data uri cannot be empty")
	}

	return nil
}

// --------------------------
// WITHDRAW PAYMENT
// --------------------------
func (msg MsgWithdrawPayment) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}

	if len(msg.Inputs) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "inputs cannot be empty")
	}
	if len(msg.Outputs) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "inputs cannot be empty")
	}

	if iidtypes.IsEmpty(msg.PaymentType.String()) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "payment type cannot be empty")
	}
	if iidtypes.IsEmpty(msg.ClaimId) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "claim id cannot be empty")
	}

	return nil
}

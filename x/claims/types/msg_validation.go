package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ixo "github.com/ixofoundation/ixo-blockchain/v2/lib/ixo"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v2/x/iid/types"
)

// --------------------------
// CREATE COLLECTION
// --------------------------
func (msg MsgCreateCollection) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	if !iidtypes.IsValidDID(msg.Entity) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.Entity)
	}

	if !iidtypes.IsValidDID(msg.Protocol) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.Protocol)
	}

	if err = msg.Payments.Validate(); err != nil {
		return err
	}

	if !ixo.IsEnumValueValid(CollectionState_name, int32(msg.State)) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid enum for state")
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
	_, err = sdk.AccAddressFromBech32(msg.AgentAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}

	if !iidtypes.IsValidDID(msg.AgentDid.Did()) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.AgentDid.String())
	}

	if ixo.IsEmpty(msg.ClaimId) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "claim_id cannot be empty")
	}
	if ixo.IsEmpty(msg.CollectionId) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "collection_id cannot be empty")
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
	_, err = sdk.AccAddressFromBech32(msg.AgentAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid agent address (%s)", err)
	}

	if !iidtypes.IsValidDID(msg.AgentDid.Did()) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.AgentDid.String())
	}
	if !iidtypes.IsValidDID(msg.Oracle) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.Oracle)
	}

	if iidtypes.IsEmpty(msg.ClaimId) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "claim_id cannot be empty")
	}
	if iidtypes.IsEmpty(msg.CollectionId) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "collection_id cannot be empty")
	}
	if iidtypes.IsEmpty(msg.VerificationProof) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "verification_proof cannot be empty")
	}

	if msg.Status == EvaluationStatus_pending {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "evaluation status can't be pending")
	}

	if !ixo.IsEnumValueValid(EvaluationStatus_name, int32(msg.Status)) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid enum for status")
	}

	if msg.Amount.IsAnyNegative() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount must be positive")
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
	if iidtypes.IsEmpty(msg.SubjectId) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "subject id cannot be empty")
	}
	if iidtypes.IsEmpty(msg.Data.Proof) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "dispute data proof cannot be empty")
	}
	if iidtypes.IsEmpty(msg.Data.Uri) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "dispute data uri cannot be empty")
	}
	if iidtypes.IsEmpty(msg.Data.Type) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "dispute data type cannot be empty")
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
	_, err = sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.ToAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid to address (%s)", err)
	}

	if err = msg.Contract_1155Payment.Validate(); err != nil {
		return err
	}

	if !ixo.IsEnumValueValid(PaymentType_name, int32(msg.PaymentType)) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid enum for payment type")
	}

	if iidtypes.IsEmpty(msg.ClaimId) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "claim id cannot be empty")
	}

	return nil
}

// --------------------------
// UPDATE COLLECTION STATE
// --------------------------
func (msg MsgUpdateCollectionState) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}
	if iidtypes.IsEmpty(msg.CollectionId) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "collection_id cannot be empty")
	}
	if !ixo.IsEnumValueValid(CollectionState_name, int32(msg.State)) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid enum for state")
	}

	return nil
}

// --------------------------
// UPDATE COLLECTION DATES
// --------------------------
func (msg MsgUpdateCollectionDates) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}
	if iidtypes.IsEmpty(msg.CollectionId) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "collection_id cannot be empty")
	}

	return nil
}

// --------------------------
// UPDATE COLLECTION PAYMENTS
// --------------------------
func (msg MsgUpdateCollectionPayments) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}
	if iidtypes.IsEmpty(msg.CollectionId) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "collection_id cannot be empty")
	}

	if err = msg.Payments.Validate(); err != nil {
		return err
	}

	return nil
}

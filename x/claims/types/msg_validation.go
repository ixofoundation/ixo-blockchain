package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ixo "github.com/ixofoundation/ixo-blockchain/v4/lib/ixo"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v4/x/iid/types"
)

// --------------------------
// CREATE COLLECTION
// --------------------------
func (msg MsgCreateCollection) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	if !iidtypes.IsValidDID(msg.Entity) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.Entity)
	}

	if !iidtypes.IsValidDID(msg.Protocol) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.Protocol)
	}

	if err = msg.Payments.Validate(); err != nil {
		return err
	}

	if !ixo.IsEnumValueValid(CollectionState_name, int32(msg.State)) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid enum for state")
	}
	if !ixo.IsEnumValueValid(CollectionIntentOptions_name, int32(msg.Intents)) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid enum for intents")
	}

	return nil
}

// --------------------------
// SUBMIT CLAIM
// --------------------------
func (msg MsgSubmitClaim) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.AgentAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}

	if !iidtypes.IsValidDID(msg.AgentDid.Did()) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.AgentDid.String())
	}

	if ixo.IsEmpty(msg.ClaimId) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "claim_id cannot be empty")
	}
	if ixo.IsEmpty(msg.CollectionId) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "collection_id cannot be empty")
	}

	if err = ValidateCoinsAllowZero(msg.Amount.Sort()); err != nil {
		return err
	}
	if err = ValidateCW20Payments(msg.Cw20Payment, true); err != nil {
		return err
	}

	return nil
}

// --------------------------
// EVALUATE CLAIM
// --------------------------
func (msg MsgEvaluateClaim) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.AgentAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid agent address (%s)", err)
	}

	if !iidtypes.IsValidDID(msg.AgentDid.Did()) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.AgentDid.String())
	}
	if !iidtypes.IsValidDID(msg.Oracle) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.Oracle)
	}

	if iidtypes.IsEmpty(msg.ClaimId) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "claim_id cannot be empty")
	}
	if iidtypes.IsEmpty(msg.CollectionId) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "collection_id cannot be empty")
	}
	if iidtypes.IsEmpty(msg.VerificationProof) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "verification_proof cannot be empty")
	}

	if msg.Status == EvaluationStatus_pending {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "evaluation status can't be pending")
	}

	if !ixo.IsEnumValueValid(EvaluationStatus_name, int32(msg.Status)) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid enum for status")
	}

	if err = ValidateCoinsAllowZero(msg.Amount.Sort()); err != nil {
		return err
	}
	if err = ValidateCW20Payments(msg.Cw20Payment, true); err != nil {
		return err
	}

	return nil
}

// --------------------------
// DISPUTE CLAIM
// --------------------------
func (msg MsgDisputeClaim) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AgentAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid agent address (%s)", err)
	}

	if !iidtypes.IsValidDID(msg.AgentDid.Did()) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.AgentDid.String())
	}
	if iidtypes.IsEmpty(msg.SubjectId) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "subject id cannot be empty")
	}
	if iidtypes.IsEmpty(msg.Data.Proof) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "dispute data proof cannot be empty")
	}
	if iidtypes.IsEmpty(msg.Data.Uri) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "dispute data uri cannot be empty")
	}
	if iidtypes.IsEmpty(msg.Data.Type) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "dispute data type cannot be empty")
	}

	return nil
}

// --------------------------
// WITHDRAW PAYMENT
// --------------------------
func (msg MsgWithdrawPayment) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.ToAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid to address (%s)", err)
	}

	if err = msg.Contract_1155Payment.Validate(); err != nil {
		return err
	}

	if err = ValidateCW20Payments(msg.Cw20Payment, true); err != nil {
		return err
	}

	if !ixo.IsEnumValueValid(PaymentType_name, int32(msg.PaymentType)) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid enum for payment type")
	}

	if iidtypes.IsEmpty(msg.ClaimId) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "claim id cannot be empty")
	}

	return nil
}

// --------------------------
// UPDATE COLLECTION STATE
// --------------------------
func (msg MsgUpdateCollectionState) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}
	if iidtypes.IsEmpty(msg.CollectionId) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "collection_id cannot be empty")
	}
	if !ixo.IsEnumValueValid(CollectionState_name, int32(msg.State)) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid enum for state")
	}

	return nil
}

// --------------------------
// UPDATE COLLECTION DATES
// --------------------------
func (msg MsgUpdateCollectionDates) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}
	if iidtypes.IsEmpty(msg.CollectionId) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "collection_id cannot be empty")
	}

	return nil
}

// --------------------------
// UPDATE COLLECTION PAYMENTS
// --------------------------
func (msg MsgUpdateCollectionPayments) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}
	if iidtypes.IsEmpty(msg.CollectionId) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "collection_id cannot be empty")
	}

	if err = msg.Payments.Validate(); err != nil {
		return err
	}

	return nil
}

// --------------------------
// UPDATE COLLECTION INTENTS
// --------------------------
func (msg MsgUpdateCollectionIntents) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}
	if iidtypes.IsEmpty(msg.CollectionId) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "collection_id cannot be empty")
	}
	if !ixo.IsEnumValueValid(CollectionIntentOptions_name, int32(msg.Intents)) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid enum for intents")
	}

	return nil
}

// --------------------------
// CLAIM INTENT
// --------------------------
func (msg *MsgClaimIntent) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.AgentAddress)
	if err != nil {
		return errorsmod.Wrapf(err, "invalid agent address (%s)", err)
	}
	if !iidtypes.IsValidDID(msg.AgentDid.Did()) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.AgentDid.String())
	}
	if iidtypes.IsEmpty(msg.CollectionId) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "collection_id cannot be empty")
	}

	if err = ValidateCoinsAllowZero(msg.Amount.Sort()); err != nil {
		return err
	}
	if err = ValidateCW20Payments(msg.Cw20Payment, true); err != nil {
		return err
	}

	return nil
}

// --------------------------
// CREATE CLAIM AUTHORIZATION
// --------------------------
func (msg MsgCreateClaimAuthorization) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.CreatorAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.GranteeAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid grantee address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.AdminAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}

	if !iidtypes.IsValidDID(msg.CreatorDid.Did()) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.CreatorDid.String())
	}

	if iidtypes.IsEmpty(msg.CollectionId) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "collection_id cannot be empty")
	}

	if !ixo.IsEnumValueValid(CreateClaimAuthorizationType_name, int32(msg.AuthType)) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid enum for auth_type")
	}

	if msg.AgentQuota == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "agent_quota cannot be 0")
	}

	if err = ValidateCoinsAllowZero(msg.MaxAmount.Sort()); err != nil {
		return err
	}

	if err = ValidateCW20Payments(msg.MaxCw20Payment, true); err != nil {
		return err
	}

	return nil
}

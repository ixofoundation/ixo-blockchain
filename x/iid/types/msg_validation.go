package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// --------------------------
// CREATE IDENTIFIER
// --------------------------
// ValidateBasic performs a basic check of the MsgCreateDidDocument fields.
func (msg MsgCreateIidDocument) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}

	if msg.Verifications == nil || len(msg.Verifications) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "verifications are required")
	}

	for _, v := range msg.Verifications {
		if err := ValidateVerification(v); err != nil {
			return err
		}
	}

	// services are optional
	if msg.Services != nil {
		for _, s := range msg.Services {
			if err := ValidateService(s); err != nil {
				return err
			}
		}
	}

	// if controllers,  must be compliant
	for _, c := range msg.Controllers {
		if !IsValidDID(c) {
			return sdkerrors.Wrap(ErrInvalidDIDFormat, "controller validation error")
		}
	}
	return nil
}

// --------------------------
// UPDATE IDENTIFIER
// --------------------------
// ValidateBasic performs a basic check of the MsgUpdateDidDocument fields.
func (msg MsgUpdateIidDocument) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}

	if msg.Verifications == nil || len(msg.Verifications) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "verifications are required")
	}

	for _, v := range msg.Verifications {
		if err := ValidateVerification(v); err != nil {
			return err
		}
	}

	// services are optional
	if msg.Services != nil {
		for _, s := range msg.Services {
			if err := ValidateService(s); err != nil {
				return err
			}
		}
	}

	for _, c := range msg.Controllers {
		// if controller is set must be compliant
		if !IsValidDID(c) {
			return sdkerrors.Wrap(ErrInvalidDIDFormat, "controller validation error")
		}
	}
	return nil
}

// --------------------------
// ADD VERIFICATION METHOD
// --------------------------
// ValidateBasic performs a basic check of the MsgAddVerification fields.
func (msg MsgAddVerification) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}

	return ValidateVerification(msg.Verification)
}

// --------------------------
// REVOKE VERIFICATION METHOD
// --------------------------
// ValidateBasic performs a basic check of the MsgRevokeVerification fields.
func (msg MsgRevokeVerification) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}

	if !IsValidDIDURL(msg.MethodId) {
		return sdkerrors.Wrap(ErrInvalidDIDURLFormat, "verification method id validation error")
	}
	return nil
}

// --------------------------
// SET VERIFICATION RELATIONSHIPS
// --------------------------
// ValidateBasic performs a basic check of the MsgSetVerificationRelationships fields.
func (msg MsgSetVerificationRelationships) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}

	if !IsValidDIDURL(msg.MethodId) {
		return sdkerrors.Wrap(ErrInvalidDIDURLFormat, "verification method id")
	}

	// there should be more then one relationship
	if len(msg.Relationships) == 0 {
		return sdkerrors.Wrap(ErrEmptyRelationships, "one ore more relationships is required")
	}

	return nil
}

// --------------------------
// ADD SERVICE
// --------------------------
// ValidateBasic performs a basic check of the MsgAddService fields.
func (msg MsgAddService) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}
	return ValidateService(msg.ServiceData)
}

// --------------------------
// DELETE SERVICE
// --------------------------
// ValidateBasic performs a basic check of the MsgDeleteService fields.
func (msg MsgDeleteService) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}

	if IsEmpty(msg.ServiceId) {
		return sdkerrors.Wrap(ErrInvalidInput, "service id cannot be empty;")
	}

	if !IsValidRFC3986Uri(msg.ServiceId) {
		return sdkerrors.Wrap(ErrInvalidRFC3986UriFormat, "service id validation error")
	}
	return nil
}

// --------------------------
// ADD CONTROLLERS
// --------------------------
// ValidateBasic performs a basic check of the MsgAddService fields.
func (msg MsgAddController) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}

	if !IsValidIIDKeyFormat(msg.ControllerDid) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.ControllerDid)
	}

	return nil
}

// --------------------------
// DELETE CONTROLLERS
// --------------------------
// ValidateBasic performs a basic check of the MsgDeleteService fields.
func (msg MsgDeleteController) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}

	if !IsValidDID(msg.ControllerDid) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.ControllerDid)
	}

	return nil
}

// --------------------------
// ADD LINKED RESOURCE
// --------------------------
// ValidateBasic performs a basic check of the MsgAddService fields.
func (msg MsgAddLinkedResource) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	if msg.LinkedResource == nil {
		return sdkerrors.Wrap(ErrInvalidInput, "linked resource cannot be nil")
	}

	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}
	return nil
}

// --------------------------
// DELETE LINKED RESOURCE
// --------------------------
// ValidateBasic performs a basic check of the MsgDeleteService fields.
func (msg MsgDeleteLinkedResource) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}

	if IsEmpty(msg.ResourceId) {
		return sdkerrors.Wrap(ErrInvalidInput, "resource id cannot be empty")
	}
	return nil
}

// --------------------------
// ADD LINKED CLAIM
// --------------------------
// ValidateBasic performs a basic check of the MsgAddService fields.
func (msg MsgAddLinkedClaim) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	if msg.LinkedClaim == nil {
		return sdkerrors.Wrap(ErrInvalidInput, "linked claim cannot be nil")
	}

	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}
	return nil
}

// --------------------------
// DELETE LINKED CLAIM
// --------------------------
// ValidateBasic performs a basic check of the MsgDeleteService fields.
func (msg MsgDeleteLinkedClaim) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}

	if IsEmpty(msg.ClaimId) {
		return sdkerrors.Wrap(ErrInvalidInput, "claim id cannot be empty")
	}
	return nil
}

// --------------------------
// ADD LINKED ENTITY
// --------------------------
func (msg MsgAddLinkedEntity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	if msg.LinkedEntity == nil {
		return sdkerrors.Wrap(ErrInvalidInput, "linked entity cannot be nil")
	}

	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}
	return nil
}

// --------------------------
// DELETE LINKED ENTITY
// --------------------------
func (msg MsgDeleteLinkedEntity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}

	if IsEmpty(msg.EntityId) {
		return sdkerrors.Wrap(ErrInvalidInput, "entity id cannot be empty;")
	}
	return nil
}

// --------------------------
// ADD RIGHT
// --------------------------
func (msg MsgAddAccordedRight) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	if msg.AccordedRight == nil {
		return sdkerrors.Wrap(ErrInvalidInput, "accordede right cannot be nil")
	}

	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}
	return nil
}

// --------------------------
// DELETE RIGHT
// --------------------------
func (msg MsgDeleteAccordedRight) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}

	if IsEmpty(msg.RightId) {
		return sdkerrors.Wrap(ErrInvalidInput, "right id cannot be empty;")
	}
	return nil
}

// --------------------------
// ADD CONTEXT
// --------------------------
func (msg MsgAddIidContext) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	if msg.Context == nil {
		return sdkerrors.Wrap(ErrInvalidInput, "context cannot be nil")
	}

	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}
	return nil
}

// --------------------------
// DELETE CONTEXT
// --------------------------
func (msg MsgDeleteIidContext) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}

	if IsEmpty(msg.ContextKey) {
		return sdkerrors.Wrap(ErrInvalidInput, "context id cannot be empty - try using the key of the context")
	}
	return nil
}

// --------------------------
// DEACTIVATE IID
// --------------------------
func (msg MsgDeactivateIID) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}

	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}
	return nil
}

package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// --------------------------
// CREATE IDENTIFIER
// --------------------------

// ValidateBasic performs a basic check of the MsgCreateDidDocument fields.
func (msg MsgCreateIidDocument) ValidateBasic() error {
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

	return nil

}

// --------------------------
// UPDATE IDENTIFIER
// --------------------------

// ValidateBasic performs a basic check of the MsgUpdateDidDocument fields.
func (msg MsgUpdateIidDocument) ValidateBasic() error {
	if !IsValidDID(msg.Doc.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Doc.Id)
	}

	for _, c := range msg.Doc.Controller {
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
	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}
	return ValidateService(msg.ServiceData)
}

func (msg MsgAddAccordedRight) ValidateBasic() error {
	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}
	return nil
}

// --------------------------
// DELETE SERVICE
// --------------------------

// ValidateBasic performs a basic check of the MsgDeleteService fields.
func (msg MsgDeleteService) ValidateBasic() error {
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
	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}

	if !IsValidDID(msg.ControllerDid) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.ControllerDid)
	}

	return nil
}

// --------------------------
// ADD Resource
// --------------------------

// ValidateBasic performs a basic check of the MsgAddService fields.
func (msg MsgAddLinkedResource) ValidateBasic() error {
	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}
	return nil
}

// --------------------------
// DELETE Resource
// --------------------------

// ValidateBasic performs a basic check of the MsgDeleteService fields.
func (msg MsgDeleteLinkedResource) ValidateBasic() error {
	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}

	if IsEmpty(msg.ResourceId) {
		return sdkerrors.Wrap(ErrInvalidInput, "resource id cannot be empty;")
	}

	//if !IsValidRFC3986Uri(msg.ServiceId) {
	//	return sdkerrors.Wrap(ErrInvalidRFC3986UriFormat, "service id validation error")
	//}
	return nil
}

func (msg MsgAddLinkedEntity) ValidateBasic() error {
	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}
	return nil
}

// --------------------------
// DELETE Resource
// --------------------------

// ValidateBasic performs a basic check of the MsgDeleteService fields.
func (msg MsgDeleteLinkedEntity) ValidateBasic() error {
	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}

	if IsEmpty(msg.EntityId) {
		return sdkerrors.Wrap(ErrInvalidInput, "entity id cannot be empty;")
	}

	if !IsValidRFC3986Uri(msg.EntityId) {
		return sdkerrors.Wrap(ErrInvalidRFC3986UriFormat, "service id validation error")
	}
	return nil
}

func (msg MsgDeleteAccordedRight) ValidateBasic() error {
	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}

	if IsEmpty(msg.RightId) {
		return sdkerrors.Wrap(ErrInvalidInput, "right id cannot be empty;")
	}

	//if !IsValidRFC3986Uri(msg.ServiceId) {
	//	return sdkerrors.Wrap(ErrInvalidRFC3986UriFormat, "service id validation error")
	//}
	return nil
}
func (msg MsgAddIidContext) ValidateBasic() error {
	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}
	return nil
}
func (msg MsgDeleteIidContext) ValidateBasic() error {
	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}

	if IsEmpty(msg.ContextKey) {
		return sdkerrors.Wrap(ErrInvalidInput, "context id cannot be empty - try using the key of the context")
	}

	//if !IsValidRFC3986Uri(msg.ServiceId) {
	//	return sdkerrors.Wrap(ErrInvalidRFC3986UriFormat, "service id validation error")
	//}
	return nil
}

func (msg MsgUpdateIidMeta) ValidateBasic() error {
	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}

	if msg.Id != msg.Meta.Id {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}

	return nil
}
func (msg MsgDeactivateIID) ValidateBasic() error {
	if !IsValidDID(msg.Id) {
		return sdkerrors.Wrap(ErrInvalidDIDFormat, msg.Id)
	}
	return nil
}

package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ixo "github.com/ixofoundation/ixo-blockchain/v3/lib/ixo"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v3/x/iid/types"
)

// --------------------------
// CREATE ENTITY
// --------------------------
func (msg MsgCreateEntity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}

	if !iidtypes.IsValidDID(msg.RelayerNode) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.RelayerNode)
	}

	if !iidtypes.IsValidDID(msg.OwnerDid.Did()) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.OwnerDid.String())
	}

	if msg.Verification == nil || len(msg.Verification) == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "verifications are required")
	}

	for _, v := range msg.Verification {
		if err := iidtypes.ValidateVerification(v); err != nil {
			return err
		}
	}

	// services are optional
	if msg.Service != nil {
		for _, s := range msg.Service {
			if err := iidtypes.ValidateService(s); err != nil {
				return err
			}
		}
	}

	// if controllers,  must be compliant
	for _, c := range msg.Controller {
		if !iidtypes.IsValidDID(c) {
			return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, "controller validation error")
		}
	}

	return nil
}

// --------------------------
// UPDATE ENTITY
// --------------------------
func (msg MsgUpdateEntity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.ControllerAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid controller address (%s)", err)
	}
	if !iidtypes.IsValidDID(msg.Id) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.Id)
	}
	if !iidtypes.IsValidDID(msg.ControllerDid.Did()) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.ControllerDid.String())
	}
	return nil
}

// --------------------------
// UPDATE ENTITY VERIFIED
// --------------------------
func (msg MsgUpdateEntityVerified) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.RelayerNodeAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid relayer node address %s", msg.RelayerNodeAddress)
	}
	if !iidtypes.IsValidDID(msg.Id) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.Id)
	}
	if !iidtypes.IsValidDID(msg.RelayerNodeDid.Did()) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.RelayerNodeDid.String())
	}
	return nil
}

// --------------------------
// TRANSFER ENTITY
// --------------------------
func (msg MsgTransferEntity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	if !iidtypes.IsValidDID(msg.Id) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.Id)
	}
	if !iidtypes.IsValidDID(msg.OwnerDid.Did()) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.OwnerDid.String())
	}
	if !iidtypes.IsValidDID(msg.RecipientDid.Did()) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.RecipientDid.String())
	}
	return nil
}

// --------------------------
// CREATE ENTITY ACCOUNT
// --------------------------
func (msg MsgCreateEntityAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	if !iidtypes.IsValidDID(msg.Id) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.Id)
	}
	if ixo.IsEmpty(msg.Name) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "name cannot be empty")
	}
	return nil
}

// --------------------------
// GRANT ENTITY ACCOUNT AUTHZ
// --------------------------
func (msg MsgGrantEntityAccountAuthz) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.GranteeAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid grantee address (%s)", err)
	}
	if !iidtypes.IsValidDID(msg.Id) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.Id)
	}
	if ixo.IsEmpty(msg.Name) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "name cannot be empty")
	}
	// Can't run this here as Grant Authorization (Any) needs to be cached using Grant.UnpackInterfaces()
	// return msg.Grant.ValidateBasic()
	return nil
}

// --------------------------
// REVOKE ENTITY ACCOUNT AUTHZ
// --------------------------
func (msg MsgRevokeEntityAccountAuthz) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.GranteeAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid grantee address (%s)", err)
	}
	if !iidtypes.IsValidDID(msg.Id) {
		return errorsmod.Wrap(iidtypes.ErrInvalidDIDFormat, msg.Id)
	}
	if ixo.IsEmpty(msg.Name) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "name cannot be empty")
	}
	if ixo.IsEmpty(msg.MsgTypeUrl) {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "msgTypeUrl cannot be empty")
	}
	return nil
}

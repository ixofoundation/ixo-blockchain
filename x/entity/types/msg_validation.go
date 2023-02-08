package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	iidtypes "github.com/ixofoundation/ixo-blockchain/x/iid/types"
)

// --------------------------
// CREATE ENTITY
// --------------------------
func (msg MsgCreateEntity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}

	if !iidtypes.IsValidDID(msg.RelayerNode) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.RelayerNode)
	}

	if msg.Verification == nil || len(msg.Verification) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "verifications are required")
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
			return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, "controller validation error")
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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid controller address (%s)", err)
	}
	if !iidtypes.IsValidDID(msg.Id) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.Id)
	}
	return nil
}

// --------------------------
// UPDATE ENTITY VERIFIED
// --------------------------
func (msg MsgUpdateEntityVerified) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.RelayerNodeAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if !iidtypes.IsValidDID(msg.Id) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.Id)
	}
	return nil
}

// --------------------------
// TRANSFER ENTITY
// --------------------------
func (msg MsgTransferEntity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	if !iidtypes.IsValidDID(msg.Id) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.Id)
	}
	if !iidtypes.IsValidDID(msg.OwnerDid.Did()) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.OwnerDid.Did())
	}
	if !iidtypes.IsValidDID(msg.RecipientDid.Did()) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.RecipientDid.Did())
	}
	return nil
}

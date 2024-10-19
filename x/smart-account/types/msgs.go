package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Helper functions
func validateSender(sender string) error {
	_, err := sdk.AccAddressFromBech32(sender)
	if err != nil {
		return fmt.Errorf("invalid sender address (%s)", err)
	}
	return nil
}

func getSender(sender string) []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// --------------------------
// ADD AUTHENTICATOR
// --------------------------
const TypeMsgAddAuthenticator = "add_authenticator"

var _ sdk.Msg = &MsgAddAuthenticator{}

func (msg MsgAddAuthenticator) Type() string { return TypeMsgAddAuthenticator }

func (msg MsgAddAuthenticator) Route() string { return RouterKey }

func (msg *MsgAddAuthenticator) ValidateBasic() error {
	return validateSender(msg.Sender)
}

func (msg *MsgAddAuthenticator) GetSigners() []sdk.AccAddress {
	return getSender(msg.Sender)
}

func (msg MsgAddAuthenticator) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// --------------------------
// REMOVE AUTHENTICATOR
// --------------------------
const TypeMsgRemoveAuthenticator = "remove_authenticator"

var _ sdk.Msg = &MsgRemoveAuthenticator{}

func (msg MsgRemoveAuthenticator) Type() string { return TypeMsgRemoveAuthenticator }

func (msg MsgRemoveAuthenticator) Route() string { return RouterKey }

func (msg *MsgRemoveAuthenticator) ValidateBasic() error {
	return validateSender(msg.Sender)
}

func (msg *MsgRemoveAuthenticator) GetSigners() []sdk.AccAddress {
	return getSender(msg.Sender)
}

func (msg MsgRemoveAuthenticator) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// --------------------------
// SET ACTIVE STATE
// --------------------------
const TypeMsgSetActiveState = "set_active_state"

var _ sdk.Msg = &MsgSetActiveState{}

func (msg MsgSetActiveState) Type() string { return TypeMsgSetActiveState }

func (msg MsgSetActiveState) Route() string { return RouterKey }

func (msg *MsgSetActiveState) ValidateBasic() error {
	return validateSender(msg.Sender)
}

func (msg *MsgSetActiveState) GetSigners() []sdk.AccAddress {
	return getSender(msg.Sender)
}

func (msg MsgSetActiveState) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (msg MsgCreateEntity) GetIidController() string { return "" }
func (msg MsgCreateEntity) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgCreateEntity) ValidateBasic() error {
	return nil
}

func (msg MsgCreateEntity) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// func (msg MsgCreateEntity) Type() string  { return TypeMsgCreateEntity }
func (msg MsgCreateEntity) Route() string { return RouterKey }

func (msg MsgUpdateEntityStatus) GetIidController() string { return "" }
func (msg MsgUpdateEntityStatus) GetSigners() []sdk.AccAddress {
	return nil
}

func (msg MsgUpdateEntityStatus) ValidateBasic() error {
	return nil
}

func (msg MsgUpdateEntityStatus) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// func (msg MsgCreateEntity) Type() string  { return TypeMsgCreateEntity }
func (msg MsgUpdateEntityStatus) Route() string { return RouterKey }

func (msg MsgUpdateEntityConfig) GetIidController() string { return "" }
func (msg MsgUpdateEntityConfig) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgUpdateEntityConfig) ValidateBasic() error {
	return nil
}

func (msg MsgUpdateEntityConfig) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// func (msg MsgCreateEntity) Type() string  { return TypeMsgCreateEntity }
func (msg MsgUpdateEntityConfig) Route() string { return RouterKey }

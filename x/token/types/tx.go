package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// func didToAddressSplitter(did string) (sdk.AccAddress, error) {
// 	bech32 := strings.Split(did, ":")
// 	address, err := sdk.AccAddressFromBech32(bech32[len(bech32)-1])
// 	if err != nil {
// 		return sdk.AccAddress{}, err
// 	}
// 	return address, nil
// }

func (msg MsgCreateToken) GetIidController() string { return msg.OwnerDid }
func (msg MsgCreateToken) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgCreateToken) ValidateBasic() error {
	return nil
}

func (msg MsgCreateToken) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// func (msg MsgCreateToken) Type() string  { return TypeMsgCreateToken }
func (msg MsgCreateToken) Route() string { return RouterKey }

func (msg MsgUpdateToken) GetIidController() string { return msg.ControllerDid }
func (msg MsgUpdateToken) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.ControllerAddress)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgUpdateToken) ValidateBasic() error {
	return nil
}

func (msg MsgUpdateToken) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// func (msg MsgCreateToken) Type() string  { return TypeMsgCreateToken }
func (msg MsgUpdateToken) Route() string { return RouterKey }

func (msg MsgUpdateTokenConfig) GetIidController() string { return "" }
func (msg MsgUpdateTokenConfig) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgUpdateTokenConfig) ValidateBasic() error {
	return nil
}

func (msg MsgUpdateTokenConfig) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// func (msg MsgCreateToken) Type() string  { return TypeMsgCreateToken }
func (msg MsgUpdateTokenConfig) Route() string { return RouterKey }

func (msg MsgTransferToken) GetIidController() string { return msg.ControllerDid }
func (msg MsgTransferToken) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.ControllerAddress)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgTransferToken) ValidateBasic() error {
	return nil
}

func (msg MsgTransferToken) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// func (msg MsgCreateToken) Type() string  { return TypeMsgCreateToken }
func (msg MsgTransferToken) Route() string { return RouterKey }

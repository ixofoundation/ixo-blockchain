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

func (msg MsgCreateEntity) GetIidController() string { return msg.OwnerDid }
func (msg MsgCreateEntity) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
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

func (msg MsgUpdateEntity) GetIidController() string { return msg.ControllerDid }
func (msg MsgUpdateEntity) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.ControllerAddress)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgUpdateEntity) ValidateBasic() error {
	return nil
}

func (msg MsgUpdateEntity) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// func (msg MsgCreateEntity) Type() string  { return TypeMsgCreateEntity }
func (msg MsgUpdateEntity) Route() string { return RouterKey }

// func (msg MsgCreateEntity) Type() string  { return TypeMsgCreateEntity }

func (msg MsgTransferEntity) GetIidController() string { return msg.OwnerDid }
func (msg MsgTransferEntity) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgTransferEntity) ValidateBasic() error {
	return nil
}

func (msg MsgTransferEntity) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// func (msg MsgCreateEntity) Type() string  { return TypeMsgCreateEntity }
func (msg MsgTransferEntity) Route() string { return RouterKey }

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	iidtypes "github.com/ixofoundation/ixo-blockchain/x/iid/types"
)

//	func didToAddressSplitter(did string) (sdk.AccAddress, error) {
//		bech32 := strings.Split(did, ":")
//		address, err := sdk.AccAddressFromBech32(bech32[len(bech32)-1])
//		if err != nil {
//			return sdk.AccAddress{}, err
//		}
//		return address, nil
//	}
func (msg MsgSetupMinter) GetIidController() iidtypes.DIDFragment {
	return msg.MinterDid
}
func (msg MsgSetupMinter) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.MinterAddress)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgSetupMinter) ValidateBasic() error {
	return nil
}

func (msg MsgSetupMinter) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// func (msg MsgSetupMinter) Type() string  { return TypeMsgSetupMinter }
func (msg MsgSetupMinter) Route() string { return RouterKey }

func (msg MsgMint) GetIidController() iidtypes.DIDFragment {
	return msg.MinterDid
}
func (msg MsgMint) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.MinterAddress)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgMint) ValidateBasic() error {
	return nil
}

func (msg MsgMint) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// func (msg MsgMint) Type() string  { return TypeMsgMintToken }
func (msg MsgMint) Route() string { return RouterKey }

func (msg MsgTransferToken) GetIidController() iidtypes.DIDFragment {
	return iidtypes.DIDFragment(msg.OwnerDid)
}
func (msg MsgTransferToken) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
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

// func (msg MsgMint) Type() string  { return TypeMsgMintToken }
func (msg MsgTransferToken) Route() string { return RouterKey }

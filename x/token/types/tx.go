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

// --------------------------
// SETUP MINTER
// --------------------------
const TypeMsgSetupMinter = "setup_minter"

var _ sdk.Msg = &MsgSetupMinter{}

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

func (msg MsgSetupMinter) Type() string { return TypeMsgSetupMinter }

func (msg MsgSetupMinter) Route() string { return RouterKey }

// --------------------------
// MINT TOKEN
// --------------------------
const TypeMsgMintToken = "mint_token"

var _ sdk.Msg = &MsgMintToken{}

func (msg MsgMintToken) GetIidController() iidtypes.DIDFragment {
	return msg.MinterDid
}
func (msg MsgMintToken) GetSigners() []sdk.AccAddress {
	address, err := sdk.AccAddressFromBech32(msg.MinterAddress)
	if err != nil {
		return []sdk.AccAddress{}
	}
	return []sdk.AccAddress{address}
}

func (msg MsgMintToken) ValidateBasic() error {
	return nil
}

func (msg MsgMintToken) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgMintToken) Type() string { return TypeMsgMintToken }

func (msg MsgMintToken) Route() string { return RouterKey }

// --------------------------
// TRANSFER TOKEN
// --------------------------
const TypeMsgTransferToken = "transfer_token"

var _ sdk.Msg = &MsgTransferToken{}

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

func (msg MsgTransferToken) Type() string { return TypeMsgTransferToken }

func (msg MsgTransferToken) Route() string { return RouterKey }

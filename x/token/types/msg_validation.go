package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	iidtypes "github.com/ixofoundation/ixo-blockchain/x/iid/types"
)

func (msg MsgCreateToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.MinterAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid minter address (%s)", err)
	}

	if iidtypes.IsEmpty(msg.Name) {
		return sdkerrors.Wrap(ErrTokenNameIncorrect, "token name is empty")
	}

	if !iidtypes.IsValidDID(msg.MinterDid.Did()) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.MinterDid.Did())
	}

	if !iidtypes.IsValidDID(msg.Class.Did()) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.Class.Did())
	}

	if !iidtypes.IsValidRFC3986Uri(msg.Image) {
		return sdkerrors.Wrapf(iidtypes.ErrInvalidRFC3986UriFormat, "image %s is not a valid RFC3986 uri", msg.Image)
	}

	return nil
}

func (msg MsgMintToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.MinterAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid minter address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.ContractAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid contract address (%s)", err)
	}

	if len(msg.MintBatch) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "batch cannot be empty")
	}

	if !iidtypes.IsValidDID(msg.MinterDid.Did()) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.MinterDid.Did())
	}

	if !iidtypes.IsValidDID(msg.OwnerDid.Did()) {
		return sdkerrors.Wrap(iidtypes.ErrInvalidDIDFormat, msg.OwnerDid.Did())
	}

	for _, batch := range msg.MintBatch {
		if iidtypes.IsEmpty(batch.Name) {
			return sdkerrors.Wrap(ErrTokenNameIncorrect, "token name is empty for a batch")
		}

		if batch.Amount.IsZero() {
			return sdkerrors.Wrap(ErrTokenAmountIncorrect, "token amount must be bigger than 0, cannot mint 0")
		}
	}

	return nil
}

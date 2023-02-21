package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	iidtypes "github.com/ixofoundation/ixo-blockchain/x/iid/types"
)

func (msg MsgCreateToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Minter)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid minter address (%s)", err)
	}

	if iidtypes.IsEmpty(msg.Name) {
		return sdkerrors.Wrap(ErrTokenNameIncorrect, "token name is empty")
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
	_, err := sdk.AccAddressFromBech32(msg.Minter)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid minter address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.ContractAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid contract address (%s)", err)
	}

	if len(msg.MintBatch) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "batch cannot be empty")
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

func (msg MsgTransferToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid recipient address (%s)", err)
	}

	if len(msg.Tokens) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "batch cannot be empty")
	}

	for _, batch := range msg.Tokens {
		if iidtypes.IsEmpty(batch.Id) {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "token id is empty for a batch")
		}

		if batch.Amount.IsZero() {
			return sdkerrors.Wrap(ErrTokenAmountIncorrect, "token amount must be bigger than 0, cannot transfer 0")
		}
	}

	return nil
}

func (msg MsgRetireToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}

	if len(msg.Tokens) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "batch cannot be empty")
	}

	if iidtypes.IsEmpty(msg.Jurisdiction) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "jurisdiction is empty")
	}

	for _, batch := range msg.Tokens {
		if iidtypes.IsEmpty(batch.Id) {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "token id is empty for a batch")
		}

		if batch.Amount.IsZero() {
			return sdkerrors.Wrap(ErrTokenAmountIncorrect, "token amount must be bigger than 0")
		}
	}

	return nil
}

func (msg MsgCancelToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}

	if len(msg.Tokens) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "batch cannot be empty")
	}

	for _, batch := range msg.Tokens {
		if iidtypes.IsEmpty(batch.Id) {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "token id is empty for a batch")
		}

		if batch.Amount.IsZero() {
			return sdkerrors.Wrap(ErrTokenAmountIncorrect, "token amount must be bigger than 0")
		}
	}

	return nil
}

func (msg MsgPauseToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Minter)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid minter address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.ContractAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid contract address (%s)", err)
	}

	return nil
}

func (msg MsgStopToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Minter)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid minter address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.ContractAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid contract address (%s)", err)
	}

	return nil
}

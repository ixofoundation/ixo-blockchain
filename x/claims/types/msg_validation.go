package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (msg MsgCreateClaim) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Minter)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid minter address (%s)", err)
	}

	return nil
}

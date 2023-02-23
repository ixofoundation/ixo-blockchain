package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/authz"
)

var (
	_ authz.Authorization = &EvaluateAuthorization{}
)

// NewEvaluateAuthorization creates a new EvaluateAuthorization object.
func NewEvaluateAuthorization(minter string, constraints []*MintConstraints) *EvaluateAuthorization {
	return &EvaluateAuthorization{
		Minter:      minter,
		Constraints: constraints,
	}
}

// MsgTypeURL implements Authorization.MsgTypeURL.
func (a EvaluateAuthorization) MsgTypeURL() string {
	return sdk.MsgTypeURL(&MsgCreateClaim{})
}

// Accept implements Authorization.Accept.
func (a EvaluateAuthorization) Accept(ctx sdk.Context, msg sdk.Msg) (authz.AcceptResponse, error) {

	return authz.AcceptResponse{Accept: true, Updated: &a}, nil
}

// ValidateBasic implements Authorization.ValidateBasic.
func (a EvaluateAuthorization) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(a.Minter)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid minter address (%s)", err)
	}

	if len(a.Constraints) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("mint authorization must contain atleast 1 constraint")
	}

	return nil
}

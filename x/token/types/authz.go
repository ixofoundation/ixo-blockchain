package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/authz"
	iidtypes "github.com/ixofoundation/ixo-blockchain/x/iid/types"
)

var (
	_ authz.Authorization = &MintAuthorization{}
)

// NewMintAuthorization creates a new MintAuthorization object.
func NewMintAuthorization(minterDid iidtypes.DIDFragment, cw20Limit, cw721Limit, ixo1155Limit int64) *MintAuthorization {
	return &MintAuthorization{
		MinterDid: minterDid,
		// MintLimit: &MintLimit{
		// 	Cw20:    string(cw20Limit),
		// 	Cw721:   string(cw721Limit),
		// 	Ixo1155: string(ixo1155Limit),
		// },
	}
}

// MsgTypeURL implements Authorization.MsgTypeURL.
func (a MintAuthorization) MsgTypeURL() string {
	return sdk.MsgTypeURL(&MsgMintToken{})
}

// Accept implements Authorization.Accept.
func (a MintAuthorization) Accept(ctx sdk.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
	mMint, ok := msg.(*MsgMintToken)
	if !ok {
		return authz.AcceptResponse{Accept: false}, sdkerrors.ErrInvalidType.Wrap("type mismatch")
	}

	if a.MinterDid.Did() != mMint.MinterDid.Did() {
		return authz.AcceptResponse{Accept: false}, sdkerrors.ErrInvalidType.Wrapf("authorized minter (%s) did not match the minter in the msg %s", a.MinterDid.Did(), mMint.MinterDid.Did())
	}

	// var foundContractInAuthorization bool = false
	// var allLimitsAreZero bool = true

	// for _, constraints := range a.Constraints {

	// }

	updatedConstraints := []*MintConstraints{}
	var matched bool

	for _, constraints := range a.Constraints {
		if constraints.ContractAddress != mMint.ContractAddress {
			updatedConstraints = append(updatedConstraints, constraints)
			continue
		}

		// state that a match was found becuase if this is full the result will deny, as this must match at least on contract.
		matched = true

		if constraints.Limit == 0 {
			return authz.AcceptResponse{Accept: false, Delete: true, Updated: &a}, sdkerrors.ErrInvalidType.Wrap("contract mint limit reached")
		}

		constraints.Limit = constraints.Limit - 1
		updatedConstraints = append(updatedConstraints, constraints)
	}

	if !matched {
		return authz.AcceptResponse{Accept: false, Delete: false, Updated: &a}, sdkerrors.ErrInvalidType.Wrap("no contract matched the constraints specified in the grant.")
	}
	a.Constraints = updatedConstraints

	// if a.MintLimit.Cw20 == 0 && a.MintLimit.Cw721 == 0 && a.MintLimit.Ixo1155 == 0 {
	// 	return authz.AcceptResponse{Accept: true, Delete: true}, nil
	// }

	return authz.AcceptResponse{Accept: true, Delete: false, Updated: &a}, nil
}

// ValidateBasic implements Authorization.ValidateBasic.
func (a MintAuthorization) ValidateBasic() error {

	for _, constraints := range a.Constraints {
		if constraints.Limit <= 0 {
			return sdkerrors.ErrInvalidCoins.Wrap("spend limit cannot be nil")
		}
	}

	return nil
}

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
	return sdk.MsgTypeURL(&MsgMint{})
}

// Accept implements Authorization.Accept.
func (a MintAuthorization) Accept(ctx sdk.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
	mMint, ok := msg.(*MsgMint)
	if !ok {
		return authz.AcceptResponse{}, sdkerrors.ErrInvalidType.Wrap("type mismatch")
	}

	if a.MinterDid.Did() != mMint.MinterDid.Did() {
		return authz.AcceptResponse{}, sdkerrors.ErrInvalidType.Wrapf("authorized minter (%s) did not match the minter in the msg %s", a.MinterDid.Did(), mMint.MinterDid.Did())
	}

	switch mMint.MintContract.(type) {
	case *MsgMint_Cw20:
		// a.MintLimit.Cw20 = a.MintLimit.Cw20 - 1
	case *MsgMint_Cw721:
		// a.MintLimit.Cw721 = a.MintLimit.Cw721 - 1
	case *MsgMint_Cw1155:
		// a.MintLimit.Ixo1155 = a.MintLimit.Ixo1155 - 1
	default:
		// err := sdkerrors.Wrapf(ErrInvalidInput, "verification material not set for verification method %s", v.Method.Id)
		return authz.AcceptResponse{}, sdkerrors.ErrInvalidType.Wrap("invalid contract provided")
	}

	// if a.MintLimit.Cw20 == 0 && a.MintLimit.Cw721 == 0 && a.MintLimit.Ixo1155 == 0 {
	// 	return authz.AcceptResponse{Accept: true, Delete: true}, nil
	// }

	return authz.AcceptResponse{Accept: true, Delete: false, Updated: &a}, nil
}

// ValidateBasic implements Authorization.ValidateBasic.
func (a MintAuthorization) ValidateBasic() error {

	// if a.MintLimit.Cw20 >= 0 {
	// 	return sdkerrors.ErrInvalidCoins.Wrap("spend limit cannot be nil")
	// }
	// if a.MintLimit.Cw721 >= 0 {
	// 	return sdkerrors.ErrInvalidCoins.Wrap("spend limit cannot be nil")
	// }
	// if a.MintLimit.Ixo1155 >= 0 {
	// 	return sdkerrors.ErrInvalidCoins.Wrap("spend limit cannot be nil")
	// }
	return nil
}

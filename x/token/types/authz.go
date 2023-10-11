package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/authz"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v2/x/iid/types"
)

var (
	_ authz.Authorization = &MintAuthorization{}
)

// NewMintAuthorization creates a new MintAuthorization object.
func NewMintAuthorization(minter string, constraints []*MintConstraints) *MintAuthorization {
	return &MintAuthorization{
		Minter:      minter,
		Constraints: constraints,
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
		return authz.AcceptResponse{}, sdkerrors.ErrInvalidType.Wrap("type mismatch")
	}

	if a.Minter != mMint.Minter {
		return authz.AcceptResponse{}, sdkerrors.ErrInvalidRequest.Wrapf("authorized minter (%s) did not match the minter in the msg %s", a.Minter, mMint.Minter)
	}

	// state indicating if there was a mismatch for one of the msgMint batches in relation to authz constraints
	var mismatch bool

	// for all msgMint batches check if there corresponding constraints
	for _, batch := range mMint.MintBatch {
		if mismatch {
			break
		}

		// loop variable to check if current batch has a matching constraint
		var matched bool
		unhandledConstraints := []*MintConstraints{}

		// check all constraints if the msg fields correlates to a granted constraint
		for _, constraint := range a.Constraints {
			// If the msg batch fields dont correlate to granted constraint, add constraint back into list
			if constraint.ContractAddress != mMint.ContractAddress || !constraint.Amount.Equal(batch.Amount) || constraint.Name != batch.Name || constraint.Index != batch.Index || constraint.Collection != batch.Collection || len(batch.TokenData) != len(constraint.TokenData) {
				unhandledConstraints = append(unhandledConstraints, constraint)
				continue
			}

			var tokenDataMismatch bool
			// check that all constraint tokenData matches the msg tokenData
			for _, tokenData1 := range constraint.TokenData {
				var tokenDataMatch bool
				for _, tokenData2 := range batch.TokenData {
					if !tokenData1.Equal(tokenData2) {
						continue
					} else {
						tokenDataMatch = true
						break
					}
				}
				if !tokenDataMatch {
					tokenDataMismatch = true
					break
				}
			}

			// If the msg batch tokenData dont correlate to granted constraint tokenData, add constraint back into list
			if tokenDataMismatch {
				unhandledConstraints = append(unhandledConstraints, constraint)
				continue
			}

			// if reaches here it means there is a matching constraint fot the specific batch
			matched = true
		}

		if matched {
			// set Auth constraints to the currently unhandled ones after the current batch constraint removed
			a.Constraints = unhandledConstraints
		} else {
			// There no cosntaint corresponding the the current batch
			mismatch = true
		}
	}

	if mismatch {
		return authz.AcceptResponse{}, sdkerrors.ErrInvalidRequest.Wrap("a mint batch doesnt correlate to granted constraints")
	}

	// If no more contraints means no more grants for grantee to mint, so delete authorization
	if len(a.Constraints) == 0 {
		return authz.AcceptResponse{Accept: true, Delete: true}, nil
	}

	return authz.AcceptResponse{Accept: true, Updated: &a}, nil
}

// ValidateBasic implements Authorization.ValidateBasic.
func (a MintAuthorization) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(a.Minter)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid minter address (%s)", err)
	}

	if len(a.Constraints) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("mint authorization must contain atleast 1 constraint")
	}

	for _, constraint := range a.Constraints {
		if constraint.Amount.IsZero() {
			return sdkerrors.ErrInvalidRequest.Wrap("amount must be bigger than 0")
		}
		_, err := sdk.AccAddressFromBech32(constraint.ContractAddress)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid contract address (%s)", err)
		}
		if iidtypes.IsEmpty(constraint.Name) {
			return sdkerrors.ErrInvalidRequest.Wrap("name cannot be empty")
		}
	}

	return nil
}

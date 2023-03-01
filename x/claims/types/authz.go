package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/authz"
	ixo "github.com/ixofoundation/ixo-blockchain/lib/ixo"
	iidtypes "github.com/ixofoundation/ixo-blockchain/x/iid/types"
)

var (
	_ authz.Authorization = &SubmitClaimAuthorization{}
	_ authz.Authorization = &EvaluateClaimAuthorization{}
)

// ---------------------------------------
// SUBMIT CLAIM
// ---------------------------------------

// NewSubmitClaimAuthorization creates a new SubmitClaimAuthorization object.
func NewSubmitClaimAuthorization(admin string, constraints []*SubmitClaimConstraints) *SubmitClaimAuthorization {
	return &SubmitClaimAuthorization{
		Admin:       admin,
		Constraints: constraints,
	}
}

// MsgTypeURL implements Authorization.MsgTypeURL.
func (a SubmitClaimAuthorization) MsgTypeURL() string {
	return sdk.MsgTypeURL(&MsgSubmitClaim{})
}

// ValidateBasic implements Authorization.ValidateBasic.
func (a SubmitClaimAuthorization) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(a.Admin)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}

	if len(a.Constraints) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("submit claim authorization must contain atleast 1 constraint")
	}

	for _, constraint := range a.Constraints {
		if constraint.AgentQuota == 0 {
			return sdkerrors.ErrInvalidRequest.Wrap("agent quota must be bigger than 0")
		}
		if iidtypes.IsEmpty(constraint.CollectionId) {
			return sdkerrors.ErrInvalidRequest.Wrap("collection id can't be empty")
		}
	}

	return nil
}

// Accept implements Authorization.Accept.
func (a SubmitClaimAuthorization) Accept(ctx sdk.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
	mSubmit, ok := msg.(*MsgSubmitClaim)
	if !ok {
		return authz.AcceptResponse{}, sdkerrors.ErrInvalidType.Wrap("type mismatch")
	}

	if a.Admin != mSubmit.AdminAddress {
		return authz.AcceptResponse{}, sdkerrors.ErrInvalidRequest.Wrapf("authorized admin (%s) did not match the admin in the msg %s", a.Admin, mSubmit.AdminAddress)
	}

	// state indicating if there was a auth constraint that matched msgSubmitClaim fields
	var matched bool
	unhandledConstraints := []*SubmitClaimConstraints{}

	// check all constraints if the msg fields correlates to a granted constraint
	for _, constraint := range a.Constraints {
		// If the msg fields dont correlate to granted constraint, add constraint back into list
		if constraint.CollectionId != mSubmit.CollectionId {
			unhandledConstraints = append(unhandledConstraints, constraint)
			continue
		}

		// if reaches here it means there is a matching constraint for the specific batch
		matched = true
		// subtract quota by one and if not 0 re-add to constraints, otherwise new quota is 0 so remove from constraints
		if constraint.AgentQuota > 1 {
			constraint.AgentQuota--
			unhandledConstraints = append(unhandledConstraints, constraint)
		}
	}

	if !matched {
		return authz.AcceptResponse{}, sdkerrors.ErrInvalidRequest.Wrap("no granted constraints correlates to the message")
	} else {
		// set Auth constraints to the currently unhandled ones after the current msg constraint removed
		a.Constraints = unhandledConstraints
	}

	// If no more contraints means no more grants for grantee to submit claims, so delete authorization
	if len(a.Constraints) == 0 {
		return authz.AcceptResponse{Accept: true, Delete: true}, nil
	}

	return authz.AcceptResponse{Accept: true, Updated: &a}, nil
}

// ---------------------------------------
// EVALUATE CLAIM
// ---------------------------------------

// NewEvaluateClaimAuthorization creates a new EvaluateClaimAuthorization object.
func NewEvaluateClaimAuthorization(admin string, constraints []*EvaluateClaimConstraints) *EvaluateClaimAuthorization {
	return &EvaluateClaimAuthorization{
		Admin:       admin,
		Constraints: constraints,
	}
}

// MsgTypeURL implements Authorization.MsgTypeURL.
func (a EvaluateClaimAuthorization) MsgTypeURL() string {
	return sdk.MsgTypeURL(&MsgEvaluateClaim{})
}

// ValidateBasic implements Authorization.ValidateBasic.
func (a EvaluateClaimAuthorization) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(a.Admin)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}

	if len(a.Constraints) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("evaluate claim authorization must contain atleast 1 constraint")
	}

	for _, constraint := range a.Constraints {
		if constraint.AgentQuota == 0 {
			return sdkerrors.ErrInvalidRequest.Wrap("agent quota must be bigger than 0")
		}
		if iidtypes.IsEmpty(constraint.CollectionId) && len(constraint.ClaimIds) == 0 {
			return sdkerrors.ErrInvalidRequest.Wrap("constraint must have either a collection_id or some claim ids")
		}
		if !iidtypes.IsEmpty(constraint.CollectionId) && len(constraint.ClaimIds) != 0 {
			return sdkerrors.ErrInvalidRequest.Wrap("constraint must have either a collection_id or some claim ids, not both")
		}
	}

	return nil
}

// Accept implements Authorization.Accept.
func (a EvaluateClaimAuthorization) Accept(ctx sdk.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
	mEval, ok := msg.(*MsgEvaluateClaim)
	if !ok {
		return authz.AcceptResponse{}, sdkerrors.ErrInvalidType.Wrap("type mismatch")
	}

	if a.Admin != mEval.AdminAddress {
		return authz.AcceptResponse{}, sdkerrors.ErrInvalidRequest.Wrapf("authorized admin (%s) did not match the admin in the msg %s", a.Admin, mEval.AdminAddress)
	}

	// state indicating if there was an auth constraint that matched msgEvaluateClaim fields
	var matched bool
	unhandledConstraints := []*EvaluateClaimConstraints{}

	// check all constraints if the msg fields correlates to a granted constraint
	for _, constraint := range a.Constraints {
		// if before_date is not zero(no validation) and is in the past then remove authZ constraint by not adding into unhandledConstraints,
		// same for when quota is 0, which should not get in constraints but adding extra check
		if (!constraint.BeforeDate.IsZero() && constraint.BeforeDate.Before(ctx.BlockTime())) || constraint.AgentQuota == 0 {
			continue
		}
		// If the msg fields dont correlate to granted constraint, add constraint back into list
		if constraint.CollectionId != mEval.CollectionId && !ixo.Contains(constraint.ClaimIds, mEval.ClaimId) {
			unhandledConstraints = append(unhandledConstraints, constraint)
			continue
		}

		// check if evaluator defined own custom amounts if allowed in constraints
		if len(mEval.Amount) != 0 {
			// state for below loop if one msg Amount is invalid whole msg is
			invalid := false
			// for each custom amount check if it within allowed max amount
			for _, mAmount := range mEval.Amount {
				// state if this specific coin amount is within allowed max
				valid := false
				for _, cAmount := range constraint.MaxCustomAmount {
					if mAmount.Denom == cAmount.Denom && cAmount.IsGTE(mAmount) {
						valid = true
					}
				}
				if !valid {
					invalid = true
					break
				}
			}

			// no amounts in constaints means not allowed to define custom amounts in msg
			if invalid || len(constraint.MaxCustomAmount) == 0 {
				unhandledConstraints = append(unhandledConstraints, constraint)
				continue
			}
		}

		// if reaches here it means there is a matching constraint for the specific batch
		matched = true
		// subtract quota by one and if not 0 re-add to constraints
		if constraint.AgentQuota > 1 {
			constraint.AgentQuota--
			// if constraint based of ClaimId then remove claimId once done
			if iidtypes.IsEmpty(constraint.CollectionId) {
				// if current constraint only has one ClaimId, which used now, dont re-add constraint once done
				if len(constraint.ClaimIds) == 1 {
					continue
				}
				var claimIds []string
				for _, claim := range constraint.ClaimIds {
					if claim != mEval.ClaimId {
						claimIds = append(claimIds, claim)
					}
				}
				constraint.ClaimIds = claimIds
			}
			unhandledConstraints = append(unhandledConstraints, constraint)
		}
	}

	// set Auth constraints to the currently unhandled ones after the current msg constraint removed or atleast outdated ones removed
	a.Constraints = unhandledConstraints

	if !matched {
		// still update constraints as above logic removes auths with passed end_date
		return authz.AcceptResponse{Accept: false, Updated: &a}, sdkerrors.ErrInvalidRequest.Wrap("no granted constraints correlates to the message")
	}

	// If no more contraints means no more grants for grantee to submit claims, so delete authorization
	if len(a.Constraints) == 0 {
		return authz.AcceptResponse{Accept: true, Delete: true}, nil
	}

	return authz.AcceptResponse{Accept: true, Updated: &a}, nil
}

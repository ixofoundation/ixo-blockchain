package types

import (
	context "context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/authz"
	ixo "github.com/ixofoundation/ixo-blockchain/v6/lib/ixo"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v6/x/iid/types"
)

var (
	_ authz.Authorization = &SubmitClaimAuthorization{}
	_ authz.Authorization = &EvaluateClaimAuthorization{}
	_ authz.Authorization = &WithdrawPaymentAuthorization{}
	_ authz.Authorization = &CreateClaimAuthorizationAuthorization{}
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
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}

	if len(a.Constraints) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("submit claim authorization must contain at least 1 constraint")
	}

	for _, constraint := range a.Constraints {
		if constraint.AgentQuota == 0 {
			return sdkerrors.ErrInvalidRequest.Wrap("agent quota must be bigger than 0")
		}
		if iidtypes.IsEmpty(constraint.CollectionId) {
			return sdkerrors.ErrInvalidRequest.Wrap("collection id can't be empty")
		}
		if err = ValidateCoinsAllowZero(constraint.MaxAmount.Sort()); err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "max amounts not valid: (%s)", err)
		}
		if err = ValidateCW20Payments(constraint.MaxCw20Payment, true); err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "max cw20 payments not valid: (%s)", err)
		}
		if err = ValidateCW1155Payments(constraint.MaxCw1155Payment, true); err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "max cw1155 payments not valid: (%s)", err)
		}
	}

	return nil
}

// Accept implements Authorization.Accept.
func (a SubmitClaimAuthorization) Accept(_ context.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
	mSubmit, ok := msg.(*MsgSubmitClaim)
	if !ok {
		return authz.AcceptResponse{}, sdkerrors.ErrInvalidType.Wrap("type mismatch")
	}

	if a.Admin != mSubmit.AdminAddress {
		return authz.AcceptResponse{}, sdkerrors.ErrInvalidRequest.Wrapf("authorized admin (%s) did not match the admin in the msg %s", a.Admin, mSubmit.AdminAddress)
	}

	// state indicating if there was a auth constraint that matched msgSubmitClaim fields
	var matched bool
	var unhandledConstraints []*SubmitClaimConstraints

	// check all constraints if the msg fields correlates to a granted constraint
	for _, constraint := range a.Constraints {
		// if quota is 0, which should not get in constraints but adding extra check, don't add to unhandled constraints
		if constraint.AgentQuota == 0 {
			continue
		}

		// if we already found a match, don't check further, only above check to maybe remove constraint
		// or if the collection id doesn't match, don't check further
		if matched || constraint.CollectionId != mSubmit.CollectionId {
			unhandledConstraints = append(unhandledConstraints, constraint)
			continue
		}

		// if reaches here it means it is a possible matching constraint, check amounts needed are within max constraints
		if len(mSubmit.Amount) != 0 && !IsCoinsInMaxConstraints(mSubmit.Amount, constraint.MaxAmount) {
			unhandledConstraints = append(unhandledConstraints, constraint)
			continue
		}
		if len(mSubmit.Cw20Payment) != 0 && !IsCW20PaymentsInMaxConstraints(mSubmit.Cw20Payment, constraint.MaxCw20Payment) {
			unhandledConstraints = append(unhandledConstraints, constraint)
			continue
		}
		if len(mSubmit.Cw1155Payment) != 0 && !IsCW1155PaymentsInMaxConstraints(mSubmit.Cw1155Payment, constraint.MaxCw1155Payment) {
			unhandledConstraints = append(unhandledConstraints, constraint)
			continue
		}

		// now the collection matches and amounts are within max constraints
		matched = true
		// subtract quota by one and if not 0 re-add to constraints, otherwise new quota is 0 so remove from constraints
		if constraint.AgentQuota > 1 {
			constraint.AgentQuota--
			unhandledConstraints = append(unhandledConstraints, constraint)
		}
	}

	// set new constraints as unhandled constraints
	a.Constraints = unhandledConstraints

	// If no more constraints means no more grants for grantee to submit claims, so delete authorization
	mustDeleteAuth := len(a.Constraints) == 0

	if !matched {
		return authz.AcceptResponse{Accept: false, Updated: &a, Delete: mustDeleteAuth}, sdkerrors.ErrInvalidRequest.Wrap("no granted constraints correlates to the message")
	}

	return authz.AcceptResponse{Accept: true, Updated: &a, Delete: mustDeleteAuth}, nil
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
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}

	if len(a.Constraints) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("evaluate claim authorization must contain at least 1 constraint")
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
		if err = ValidateCoinsAllowZero(constraint.MaxCustomAmount.Sort()); err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "max custom amounts not valid: (%s)", err)
		}
		if err = ValidateCW20Payments(constraint.MaxCustomCw20Payment, true); err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "max custom cw20 payments not valid: (%s)", err)
		}
		if err = ValidateCW1155Payments(constraint.MaxCustomCw1155Payment, true); err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "max custom cw1155 payments not valid: (%s)", err)
		}
	}

	return nil
}

// Accept implements Authorization.Accept.
func (a EvaluateClaimAuthorization) Accept(ctx context.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
	mEval, ok := msg.(*MsgEvaluateClaim)
	if !ok {
		return authz.AcceptResponse{}, sdkerrors.ErrInvalidType.Wrap("type mismatch")
	}

	if a.Admin != mEval.AdminAddress {
		return authz.AcceptResponse{}, sdkerrors.ErrInvalidRequest.Wrapf("authorized admin (%s) did not match the admin in the msg %s", a.Admin, mEval.AdminAddress)
	}

	// state indicating if there was an auth constraint that matched msgEvaluateClaim fields
	var matched bool
	var unhandledConstraints []*EvaluateClaimConstraints
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// check all constraints if the msg fields correlates to a granted constraint
	for _, constraint := range a.Constraints {
		// if before_date is not zero(no validation) and is in the past then remove authZ constraint by not adding into unhandledConstraints,
		// same for when quota is 0, which should not get in constraints but adding extra check
		if (constraint.BeforeDate != nil && constraint.BeforeDate.Before(sdkCtx.BlockTime())) || constraint.AgentQuota == 0 {
			continue
		}

		// if we already found a match, don't check further, only above check to maybe remove constraint
		if matched {
			unhandledConstraints = append(unhandledConstraints, constraint)
			continue
		}

		// If the msg fields dont correlate to granted constraint, add constraint back into list
		if constraint.CollectionId != mEval.CollectionId && !ixo.Contains(constraint.ClaimIds, mEval.ClaimId) {
			unhandledConstraints = append(unhandledConstraints, constraint)
			continue
		}

		// if reaches here it means it is a possible matching constraint, check amounts needed are within max constraints
		if len(mEval.Amount) != 0 && !IsCoinsInMaxConstraints(mEval.Amount, constraint.MaxCustomAmount) {
			unhandledConstraints = append(unhandledConstraints, constraint)
			continue
		}
		if len(mEval.Cw20Payment) != 0 && !IsCW20PaymentsInMaxConstraints(mEval.Cw20Payment, constraint.MaxCustomCw20Payment) {
			unhandledConstraints = append(unhandledConstraints, constraint)
			continue
		}
		if len(mEval.Cw1155Payment) != 0 && !IsCW1155PaymentsInMaxConstraints(mEval.Cw1155Payment, constraint.MaxCustomCw1155Payment) {
			unhandledConstraints = append(unhandledConstraints, constraint)
			continue
		}

		// if reaches here it means there is a matching constraint for the specific batch,
		// meaning if custom amounts defined it was within constraints, otherwise just the collection id or claim id was in constraints
		matched = true

		// subtract quota by one (if eval status is not invalidated) and if not 0 re-add to constraints
		if constraint.AgentQuota > 1 || mEval.Status == EvaluationStatus_invalidated {
			// if evaluation status is invalidated then don't subtract quota
			if mEval.Status != EvaluationStatus_invalidated {
				constraint.AgentQuota--
			}

			// if constraint based of ClaimId then remove claimId once done
			if iidtypes.IsEmpty(constraint.CollectionId) {
				// if current constraint only has one ClaimId, which used now, don't re-add constraint once done
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

			// add constraint back into list only if quota is not 0
			if constraint.AgentQuota > 0 {
				unhandledConstraints = append(unhandledConstraints, constraint)
			}
		}
	}

	// set new constraints as unhandled constraints
	a.Constraints = unhandledConstraints

	// If no more constraints means no more grants for grantee to evaluate claims, so delete authorization
	mustDeleteAuth := len(a.Constraints) == 0

	if !matched {
		// update constraints as above logic removes auths with passed end_date
		return authz.AcceptResponse{Accept: false, Updated: &a, Delete: mustDeleteAuth}, sdkerrors.ErrInvalidRequest.Wrap("no granted constraints correlates to the message")
	}

	return authz.AcceptResponse{Accept: true, Updated: &a, Delete: mustDeleteAuth}, nil
}

// ---------------------------------------
// WITHDRAW
// ---------------------------------------

// NewWithdrawPaymentAuthorization creates a new WithdrawPaymentAuthorization object.
func NewWithdrawPaymentAuthorization(admin string, constraints []*WithdrawPaymentConstraints) *WithdrawPaymentAuthorization {
	return &WithdrawPaymentAuthorization{
		Admin:       admin,
		Constraints: constraints,
	}
}

// MsgTypeURL implements Authorization.MsgTypeURL.
func (a WithdrawPaymentAuthorization) MsgTypeURL() string {
	return sdk.MsgTypeURL(&MsgWithdrawPayment{})
}

// ValidateBasic implements Authorization.ValidateBasic.
func (a WithdrawPaymentAuthorization) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(a.Admin)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}

	if len(a.Constraints) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("withdraw payment authorization must contain at least 1 constraint")
	}

	for _, constraint := range a.Constraints {
		_, err := sdk.AccAddressFromBech32(constraint.FromAddress)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address (%s)", err)
		}
		_, err = sdk.AccAddressFromBech32(constraint.ToAddress)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid to address (%s)", err)
		}
		if iidtypes.IsEmpty(constraint.PaymentType.String()) {
			return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "payment type cannot be empty")
		}
		if iidtypes.IsEmpty(constraint.ClaimId) {
			return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "claim id cannot be empty")
		}
		if err = ValidateCW20Payments(constraint.Cw20Payment, true); err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "cw20 payments not valid: (%s)", err)
		}
		if err = ValidateCW1155Payments(constraint.Cw1155Payment, true); err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "cw1155 payments not valid: (%s)", err)
		}
	}

	return nil
}

// Accept implements Authorization.Accept.
func (a WithdrawPaymentAuthorization) Accept(ctx context.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
	mWith, ok := msg.(*MsgWithdrawPayment)
	if !ok {
		return authz.AcceptResponse{}, sdkerrors.ErrInvalidType.Wrap("type mismatch")
	}

	if a.Admin != mWith.AdminAddress {
		return authz.AcceptResponse{}, sdkerrors.ErrInvalidRequest.Wrapf("authorized admin (%s) did not match the admin in the msg %s", a.Admin, mWith.AdminAddress)
	}

	// state indicating if there was an auth constraint that matched msgWithdrawPayment fields
	var matched bool
	var unhandledConstraints []*WithdrawPaymentConstraints
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// check all constraints if the msg fields correlates to a granted constraint
	for _, constraint := range a.Constraints {
		// If the msg fields doesn't correlate to a granted constraint, add constraint back into list
		if constraint.ClaimId != mWith.ClaimId || constraint.PaymentType != mWith.PaymentType {
			unhandledConstraints = append(unhandledConstraints, constraint)
			continue
		}

		// check that withdraw has reached release date yet if it exists
		if constraint.ReleaseDate != nil && constraint.ReleaseDate.After(sdkCtx.BlockTime()) {
			return authz.AcceptResponse{}, sdkerrors.ErrInvalidRequest.Wrapf("constraint release date not reached")
		}

		// check that from address is same
		if mWith.FromAddress != constraint.FromAddress {
			return authz.AcceptResponse{}, sdkerrors.ErrInvalidRequest.Wrapf("from address in msg does not match constraint")
		}

		// check that to address is same
		if mWith.ToAddress != constraint.ToAddress {
			return authz.AcceptResponse{}, sdkerrors.ErrInvalidRequest.Wrapf("to address in msg does not match constraint")
		}

		// check that withdraw input and output lengths are the same
		if len(constraint.Inputs) != len(mWith.Inputs) || len(constraint.Outputs) != len(mWith.Outputs) {
			return authz.AcceptResponse{}, sdkerrors.ErrInvalidRequest.Wrapf("lengths of Input/Output in msg does not match constraint")
		}

		// if has input/output then check that valid
		if len(mWith.Inputs) != 0 {
			// for each msg input see if there corresponding constraint input
			constraintInputs := constraint.Inputs
			for _, mInput := range mWith.Inputs {
				// state if this specific input is valid
				valid := false
				for i, cInput := range constraintInputs {
					if cInput.Address == mInput.Address && mInput.Coins.Equal(cInput.Coins) {
						constraintInputs = ixo.RemoveUnordered(constraintInputs, i)
						valid = true
						break
					}
				}
				if !valid {
					return authz.AcceptResponse{}, sdkerrors.ErrInvalidRequest.Wrapf("msg inputs does not match constraint inputs")
				}
			}

			// for each msg output see if there corresponding constraint output
			constraintOutputs := constraint.Outputs
			for _, mOutput := range mWith.Outputs {
				// state if this specific Output is valid
				valid := false
				for i, cOutput := range constraintOutputs {
					if cOutput.Address == mOutput.Address && mOutput.Coins.Equal(cOutput.Coins) {
						constraintOutputs = ixo.RemoveUnordered(constraintOutputs, i)
						valid = true
						break
					}
				}
				if !valid {
					return authz.AcceptResponse{}, sdkerrors.ErrInvalidRequest.Wrapf("msg outputs does not match constraint outputs")
				}
			}
		}

		// check that cw20 payments lengths match
		if len(mWith.Cw20Payment) != len(constraint.Cw20Payment) {
			return authz.AcceptResponse{}, sdkerrors.ErrInvalidRequest.Wrapf("cw20 payments length does not match constraint")
		}

		// if has cw20 payments then check that valid
		if len(mWith.Cw20Payment) != 0 {
			constraintCw20Payments := constraint.Cw20Payment
			// then check that for each cw20 payment there is a corresponding constraint cw20 payment
			for _, mCw20Payment := range mWith.Cw20Payment {
				// state if this specific cw20Payment is valid
				valid := false
				for i, cCw20Payment := range constraintCw20Payments {
					if cCw20Payment.Equal(mCw20Payment) {
						// if found match then remove from constraint cw20 payments and mark as valid
						constraintCw20Payments = ixo.RemoveUnordered(constraintCw20Payments, i)
						valid = true
						break
					}
				}
				if !valid {
					return authz.AcceptResponse{}, sdkerrors.ErrInvalidRequest.Wrapf("msg cw20 payments does not match constraint cw20 payments")
				}
			}
		}

		// check that cw1155 payments lengths match
		if len(mWith.Cw1155Payment) != len(constraint.Cw1155Payment) {
			return authz.AcceptResponse{}, sdkerrors.ErrInvalidRequest.Wrapf("cw1155 payments length does not match constraint")
		}

		// if has cw1155 payments then check that valid
		if len(mWith.Cw1155Payment) != 0 {
			constraintCw1155Payments := constraint.Cw1155Payment
			// then check that for each cw1155 payment there is a corresponding constraint cw1155 payment
			for _, mCw1155Payment := range mWith.Cw1155Payment {
				// state if this specific cw1155 payment is valid
				valid := false
				for i, cCw1155Payment := range constraintCw1155Payments {
					if cCw1155Payment.Equal(mCw1155Payment) {
						// if found match then remove from constraint cw1155 payments and mark as valid
						constraintCw1155Payments = ixo.RemoveUnordered(constraintCw1155Payments, i)
						valid = true
						break
					}
				}
				if !valid {
					return authz.AcceptResponse{}, sdkerrors.ErrInvalidRequest.Wrapf("msg cw1155 payments does not match constraint cw1155 payments")
				}
			}
		}

		// if reaches here it means there is a matching constraint for the specific withdraw
		matched = true
	}

	if !matched {
		return authz.AcceptResponse{}, sdkerrors.ErrInvalidRequest.Wrap("no granted constraints correlates to the message")
	}

	// set new constraints as unhandled constraints
	a.Constraints = unhandledConstraints

	// If no more constraints means no more grants for grantee to submit claims, so delete authorization
	mustDeleteAuth := len(a.Constraints) == 0

	return authz.AcceptResponse{Accept: true, Updated: &a, Delete: mustDeleteAuth}, nil
}

// ---------------------------------------
// CREATE CLAIM AUTHORIZATION
// ---------------------------------------

// NewCreateClaimAuthorizationAuthorization creates a new CreateClaimAuthorizationAuthorization object.
func NewCreateClaimAuthorizationAuthorization(admin string, constraints []*CreateClaimAuthorizationConstraints) *CreateClaimAuthorizationAuthorization {
	return &CreateClaimAuthorizationAuthorization{
		Admin:       admin,
		Constraints: constraints,
	}
}

// MsgTypeURL implements Authorization.MsgTypeURL.
func (a CreateClaimAuthorizationAuthorization) MsgTypeURL() string {
	return sdk.MsgTypeURL(&MsgCreateClaimAuthorization{})
}

// ValidateBasic implements Authorization.ValidateBasic.
func (a CreateClaimAuthorizationAuthorization) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(a.Admin)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address (%s)", err)
	}

	// check that there is at least one constraint
	if len(a.Constraints) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("create claim authorization must contain at least 1 constraint")
	}

	// check that all constraints are valid
	for _, constraint := range a.Constraints {
		if err = ValidateCoinsAllowZero(constraint.MaxAmount.Sort()); err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "max amounts not valid: (%s)", err)
		}
		if err = ValidateCW20Payments(constraint.MaxCw20Payment, true); err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "max cw20 payments not valid: (%s)", err)
		}
		if err = ValidateCW1155Payments(constraint.MaxCw1155Payment, true); err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "max cw1155 payments not valid: (%s)", err)
		}

		if !ixo.IsEnumValueValid(CreateClaimAuthorizationType_name, int32(constraint.AllowedAuthTypes)) {
			return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid enum for allowed_auth_types")
		}
	}

	return nil
}

// Accept implements Authorization.Accept.
func (a CreateClaimAuthorizationAuthorization) Accept(ctx context.Context, msg sdk.Msg) (authz.AcceptResponse, error) {
	mCreate, ok := msg.(*MsgCreateClaimAuthorization)
	if !ok {
		return authz.AcceptResponse{}, sdkerrors.ErrInvalidType.Wrap("type mismatch")
	}

	if a.Admin != mCreate.AdminAddress {
		return authz.AcceptResponse{}, sdkerrors.ErrInvalidRequest.Wrapf("authorized admin (%s) did not match the admin in the msg %s", a.Admin, mCreate.AdminAddress)
	}

	// state indicating if there was a auth constraint that matched msg fields
	var matched bool
	var unhandledConstraints []*CreateClaimAuthorizationConstraints
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// check all constraints if the msg fields correlates to a granted constraint
	for _, constraint := range a.Constraints {
		// If the constraint has an expiration, check that it before now block time
		// if before (already expired) then don't add to unhandledConstraints
		if constraint.Expiration != nil && constraint.Expiration.Before(sdkCtx.BlockTime()) {
			continue
		}

		// if we already found a match, don't check further, only above check to maybe remove constraint
		if matched {
			unhandledConstraints = append(unhandledConstraints, constraint)
			continue
		}

		// If collection_ids is not empty, check if this collection is allowed
		if len(constraint.CollectionIds) > 0 && !ixo.Contains(constraint.CollectionIds, mCreate.CollectionId) {
			unhandledConstraints = append(unhandledConstraints, constraint)
			continue
		}

		// Check if the authorization type is allowed
		if constraint.AllowedAuthTypes != CreateClaimAuthorizationType_ALL &&
			constraint.AllowedAuthTypes != mCreate.AuthType {
			unhandledConstraints = append(unhandledConstraints, constraint)
			continue
		}

		// Check if the agent quota is within the allowed maximum
		if constraint.MaxAgentQuota > 0 && mCreate.AgentQuota > constraint.MaxAgentQuota {
			unhandledConstraints = append(unhandledConstraints, constraint)
			continue
		}

		// Check if custom amount is within the max amount constraint
		if len(constraint.MaxAmount) > 0 && len(mCreate.MaxAmount) > 0 && !IsCoinsInMaxConstraints(mCreate.MaxAmount, constraint.MaxAmount) {
			unhandledConstraints = append(unhandledConstraints, constraint)
			continue
		}

		// Check if CW20 payment is within the max constraints
		if len(constraint.MaxCw20Payment) > 0 && len(mCreate.MaxCw20Payment) > 0 && !IsCW20PaymentsInMaxConstraints(mCreate.MaxCw20Payment, constraint.MaxCw20Payment) {
			unhandledConstraints = append(unhandledConstraints, constraint)
			continue
		}

		// Check if CW1155 payment is within the max constraints
		if len(constraint.MaxCw1155Payment) > 0 && len(mCreate.MaxCw1155Payment) > 0 && !IsCW1155PaymentsInMaxConstraints(mCreate.MaxCw1155Payment, constraint.MaxCw1155Payment) {
			unhandledConstraints = append(unhandledConstraints, constraint)
			continue
		}

		// Check intent duration if this is for a submit authorization
		if mCreate.AuthType == CreateClaimAuthorizationType_SUBMIT {
			if mCreate.IntentDurationNs.Nanoseconds() > constraint.MaxIntentDurationNs.Nanoseconds() {
				unhandledConstraints = append(unhandledConstraints, constraint)
				continue
			}
		}

		// Mark as matched since we've found a valid constraint
		matched = true

		// If max authorizations is 0 then don't decrement and re-add constraint (unlimited authorizations)
		// Otherwise decrement the max authorizations count and re-add constraint if more than 1
		if constraint.MaxAuthorizations == 0 {
			unhandledConstraints = append(unhandledConstraints, constraint)
		} else if constraint.MaxAuthorizations > 1 {
			constraint.MaxAuthorizations--
			unhandledConstraints = append(unhandledConstraints, constraint)
		}
	}

	// set new constraints as unhandled constraints
	a.Constraints = unhandledConstraints

	// If no more constraints means no more grants for grantee to submit claims, so delete authorization
	mustDeleteAuth := len(a.Constraints) == 0

	// if no constraints matched, return error, but still update constraints to remove outdated constraints
	if !matched {
		return authz.AcceptResponse{Accept: false, Updated: &a, Delete: mustDeleteAuth}, sdkerrors.ErrInvalidRequest.Wrap("no granted constraints correlates to the message, please check that all constraints are valid")
	}

	return authz.AcceptResponse{Accept: true, Updated: &a, Delete: mustDeleteAuth}, nil
}

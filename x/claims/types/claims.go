package types

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	entitytypes "github.com/ixofoundation/ixo-blockchain/v3/x/entity/types"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v3/x/iid/types"
)

// IsValidCollection tells if a Claim Collection is valid,
func IsValidCollection(collection *Collection) bool {
	if collection == nil {
		return false
	}
	if iidtypes.IsEmpty(collection.Id) {
		return false
	}
	_, err := sdk.AccAddressFromBech32(collection.Admin)
	if err != nil {
		return false
	}
	if !iidtypes.IsValidDID(collection.Entity) {
		return false
	}
	if !iidtypes.IsValidDID(collection.Protocol) {
		return false
	}
	return true
}

// IsValidClaim tells if a Claim is valid,
func IsValidClaim(claim *Claim) bool {
	if claim == nil {
		return false
	}
	if iidtypes.IsEmpty(claim.ClaimId) {
		return false
	}
	if !iidtypes.IsValidDID(claim.AgentDid) {
		return false
	}
	return true
}

// IsValidDispute tells if a Dispute is valid,
func IsValidDispute(dispute *Dispute) bool {
	if dispute == nil {
		return false
	}
	if iidtypes.IsEmpty(dispute.SubjectId) {
		return false
	}
	if iidtypes.IsEmpty(dispute.Data.Proof) {
		return false
	}
	if iidtypes.IsEmpty(dispute.Data.Uri) {
		return false
	}
	return true
}

func HasBalances(ctx sdk.Context, bankKeeper bankkeeper.Keeper, payerAddr sdk.AccAddress,
	requiredFunds sdk.Coins) bool {
	for _, coin := range requiredFunds {
		if !bankKeeper.HasBalance(ctx, payerAddr, coin) {
			return false
		}
	}

	return true
}

func (p *Contract1155Payment) Validate() error {
	if p != nil {
		_, err := sdk.AccAddressFromBech32(p.Address)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "err %s", err)
		}
		if iidtypes.IsEmpty(p.TokenId) {
			return fmt.Errorf("token id cannot be empty")
		}
		// if p.Contract_1155Payment.Amount == 0 {
		// 	return fmt.Errorf("token amount cannot be 0")
		// }
	}

	return nil
}

func (p Payment) Validate() error {
	_, err := sdk.AccAddressFromBech32(p.Account)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "err %s", err)
	}

	if err = p.Contract_1155Payment.Validate(); err != nil {
		return err
	}

	if p.Amount.IsAnyNegative() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "payment amount must be positive")
	}

	return nil
}

func (p Payments) AccountsIsEntityAccounts(entity entitytypes.Entity) bool {
	if !entity.ContainsAccountAddress(p.Approval.Account) || !entity.ContainsAccountAddress(p.Submission.Account) || !entity.ContainsAccountAddress(p.Rejection.Account) || !entity.ContainsAccountAddress(p.Evaluation.Account) {
		return false
	}
	return true
}

func (p Payments) Validate() error {
	if p.Evaluation.Contract_1155Payment != nil {
		return ErrCollectionEvalError
	}

	if err := p.Submission.Validate(); err != nil {
		return err
	}
	if err := p.Evaluation.Validate(); err != nil {
		return err
	}
	if err := p.Approval.Validate(); err != nil {
		return err
	}
	if err := p.Rejection.Validate(); err != nil {
		return err
	}

	return nil
}

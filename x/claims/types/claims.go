package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	iidtypes "github.com/ixofoundation/ixo-blockchain/x/iid/types"
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

func (p Payment) Validate() error {
	_, err := sdk.AccAddressFromBech32(p.Account)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "err %s", err)
	}

	if p.Contract_1155Payment != nil {
		_, err := sdk.AccAddressFromBech32(p.Contract_1155Payment.Address)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "err %s", err)
		}
		if iidtypes.IsEmpty(p.Contract_1155Payment.TokenId) {
			return fmt.Errorf("token id cannot be empty")
		}
		// if p.Contract_1155Payment.Amount == 0 {
		// 	return fmt.Errorf("token amount cannot be 0")
		// }
	}

	return nil
}

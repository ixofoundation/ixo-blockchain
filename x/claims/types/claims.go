package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
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

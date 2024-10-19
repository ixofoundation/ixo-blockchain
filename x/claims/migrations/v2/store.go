package v2claims

import (
	errorsmod "cosmossdk.io/errors"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/v4/x/claims/types"
)

// MigrateStore performs in-place store migrations from ConsensusVersion 1 to 2.
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	// Get all existing collections from the store
	iterator := storetypes.KVStorePrefixIterator(store, types.CollectionKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var oldCollection Collection // Define struct for v1 collections
		err := cdc.Unmarshal(iterator.Value(), &oldCollection)
		if err != nil {
			return errorsmod.Wrap(err, "failed to unmarshal collection")
		}

		// Convert v1 collection to v2 collection
		newCollection := types.Collection{
			Id:        oldCollection.Id,
			Entity:    oldCollection.Entity,
			Admin:     oldCollection.Admin,
			Protocol:  oldCollection.Protocol,
			StartDate: oldCollection.StartDate,
			EndDate:   oldCollection.EndDate,
			Quota:     oldCollection.Quota,
			Count:     oldCollection.Count,
			Evaluated: oldCollection.Evaluated,
			Approved:  oldCollection.Approved,
			Rejected:  oldCollection.Rejected,
			Disputed:  oldCollection.Disputed,
			State:     types.CollectionState(oldCollection.State),
			Payments:  convertPayments(*oldCollection.Payments),
			Signer:    oldCollection.Signer,
			// Add the new field "Invalidated" with a default value of 0.
			Invalidated: 0,
		}

		// Marshal the new collection and store it with the same key
		marshaled, err := cdc.Marshal(&newCollection)
		if err != nil {
			return errorsmod.Wrap(err, "failed to marshal new collection")
		}
		store.Set(iterator.Key(), marshaled)
	}

	return nil
}

// Define a function to convert old Contract1155Payment type to new Contract1155Payment type
func convertContract1155Payment(oldContract *Contract1155Payment) *types.Contract1155Payment {
	if oldContract == nil {
		return nil
	}

	return &types.Contract1155Payment{
		Address: oldContract.Address,
		TokenId: oldContract.TokenId,
		Amount:  oldContract.Amount,
	}
}

// Define a function to convert old Payment type to new Payment type
func convertPayment(oldPayment *Payment) *types.Payment {
	if oldPayment == nil {
		return nil
	}

	newPayment := &types.Payment{
		Account:   oldPayment.Account,
		Amount:    oldPayment.Amount,
		TimeoutNs: oldPayment.TimeoutNs,
	}

	// Convert Contract1155Payment if present
	newPayment.Contract_1155Payment = convertContract1155Payment(oldPayment.Contract_1155Payment)

	return newPayment
}

// Define a function to convert old Payments type to new Payments type
func convertPayments(oldPayments Payments) *types.Payments {
	return &types.Payments{
		Submission: convertPayment(oldPayments.Submission),
		Evaluation: convertPayment(oldPayments.Evaluation),
		Approval:   convertPayment(oldPayments.Approval),
		Rejection:  convertPayment(oldPayments.Rejection),
	}
}

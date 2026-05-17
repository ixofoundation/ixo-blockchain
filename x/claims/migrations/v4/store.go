package v4claims

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-blockchain/v6/x/claims/types"
)

// MigrateStore performs the in-place store migrations from claims module
// ConsensusVersion 3 to 4 (the v7 chain upgrade).
//
// The new Collection and Dispute proto fields default to zero/empty on the
// existing serialized records, so a re-read on a v7 node already produces
// a struct with sensible defaults — no need to rewrite Collections.
//
// What we DO migrate:
//
//  1. Legacy disputes (those with target_role == UNSPECIFIED, which means
//     they were filed before v7) are stamped status=DISMISSED so they
//     don't permanently block new disputes under the new (subject_id,
//     target_role) gating. They remain queryable by proof under the
//     existing primary key.
//
//  2. We do NOT backfill the subject index for legacy disputes — without
//     a target_role we can't compose the index key, and forcing one would
//     misrepresent history. New disputes after v7 will write their own
//     index entries.
//
//  3. We do NOT touch agent deposit balances — no rows exist yet.
//
// No new sub-stores need to be cleared; the new key prefixes (0x06, 0x07,
// 0x08) live alongside the existing ones and start empty.
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)

	// Rewrite every existing Dispute record with status=DISMISSED if its
	// target_role is UNSPECIFIED (i.e., a pre-v7 dispute). New disputes
	// written by v7 code have an explicit role and status=OPEN, so they
	// pass through unchanged on a re-export/re-import.
	iter := storetypes.KVStorePrefixIterator(store, types.DisputeKey)
	defer iter.Close()

	migrated := 0
	for ; iter.Valid(); iter.Next() {
		var d types.Dispute
		if err := cdc.Unmarshal(iter.Value(), &d); err != nil {
			return errorsmod.Wrapf(err, "failed to unmarshal dispute at key %x", iter.Key())
		}
		// Only touch legacy disputes. Anything written by v7 code already
		// has a role and a status set, so leave it alone.
		if d.TargetRole != types.DisputeTargetRole_target_unspecified {
			continue
		}
		// Default status was zero (OPEN) on the wire; mark these legacy
		// disputes DISMISSED so the new "no further disputes after AWARDED"
		// rule doesn't accidentally block a brand-new v7 dispute against
		// the same claim.
		d.Status = types.DisputeStatus_dispute_dismissed
		bz, err := cdc.Marshal(&d)
		if err != nil {
			return errorsmod.Wrapf(err, "failed to marshal migrated dispute %x", iter.Key())
		}
		// The iter.Key() returned by KVStorePrefixIterator is the FULL key
		// including the table prefix, so we can write it straight back.
		store.Set(iter.Key(), bz)
		migrated++
	}

	ctx.Logger().Info(fmt.Sprintf("x/claims v3->v4: stamped %d legacy disputes as DISMISSED", migrated))
	return nil
}

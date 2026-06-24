package keeper

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	iidtypes "github.com/ixofoundation/ixo-blockchain/v8/x/iid/types"
	"github.com/ixofoundation/ixo-blockchain/v8/x/names/types"
)

// verifyDidController ensures `signer` controls `did`. Control is established
// by either:
//   - having a verification method on the DID document with the
//     authentication relationship that resolves to `signer`'s address; or
//   - being listed as a controller of the DID document, expressed as the
//     signer's bech32 address (consistent with how iid x/keeper handles it).
//
// Returns ErrUnauthorized otherwise.
func (k Keeper) verifyDidController(ctx sdk.Context, did, signer string) error {
	doc, found := k.iidKeeper.GetDidDocument(ctx, []byte(did))
	if !found {
		return errorsmod.Wrapf(types.ErrInvalidDID, "DID %q not found", did)
	}
	if doc.HasRelationship(iidtypes.NewBlockchainAccountID(signer), iidtypes.Authentication) {
		return nil
	}
	if doc.HasController(iidtypes.DID(signer)) {
		return nil
	}
	return errorsmod.Wrapf(types.ErrUnauthorized, "signer %s does not control DID %s", signer, did)
}

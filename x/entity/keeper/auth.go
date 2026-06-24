package keeper

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	iidtypes "github.com/ixofoundation/ixo-blockchain/v8/x/iid/types"
)

// VerifyDidSignerAuthentication re-establishes, inside the keeper, the
// signer→DID binding that the IID ante decorator (x/iid/ante) performs at the
// transaction level: it requires that signerAddr — the message's proto signer
// field (cosmos.msg.v1.signer) — holds an `authentication` verification
// relationship on `did`.
//
// This MUST be called by every entity handler that authorises an action "as" a
// DID, because the IID ante decorator does NOT run on messages dispatched via
// CosmWasm contract execution, ICA host, or other non-top-level routes. Without
// an in-keeper re-check, those routes could act as a DID the caller does not
// control — e.g. quoting a victim entity's public controller DID to mutate its
// state (see the 2026-06 audit; the entity handlers previously only verified
// DID-controls-DID via HasController and relied on the ante for the
// signer→DID half of the binding).
func (k Keeper) VerifyDidSignerAuthentication(ctx sdk.Context, did, signerAddr string) error {
	doc, found := k.IidKeeper.GetDidDocument(ctx, []byte(did))
	if !found {
		return errorsmod.Wrapf(iidtypes.ErrDidDocumentNotFound, "did document %s not found", did)
	}
	if !doc.HasRelationship(iidtypes.NewBlockchainAccountID(signerAddr), iidtypes.Authentication) {
		return errorsmod.Wrapf(iidtypes.ErrUnauthorized,
			"signer %s is not authorized to act on behalf of did %s", signerAddr, did)
	}
	return nil
}

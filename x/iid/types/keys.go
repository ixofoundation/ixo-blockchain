package types

import (
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "iid"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// DidChainPrefix defines the did prefix for this chain
	DidChainPrefix = "did:x:"

	// IxoDidPrefix is the canonical method-and-namespace prefix for every
	// user-created IID on this chain. MsgCreateIidDocument refuses any id
	// outside this prefix — see ValidateMsgCreateDIDForm. Module-internal
	// DIDs (e.g. did:ixo:entity:...) also live under this prefix; they go
	// through IidKeeper.SetDidDocument directly and are listed in
	// ReservedDidPrefixes.
	IxoDidPrefix = "did:ixo:"

	// IxoWasmDidSubPrefix marks the second segment of the wasm-contract
	// DID form did:ixo:wasm:<contract-address>. Anyone may create one of
	// these (they identify a contract, not a user account), but the part
	// after `wasm:` must be a valid bech32 account address.
	IxoWasmDidSubPrefix = "wasm:"

	// IxoWasmDidPrefix is the full prefix for the wasm-contract DID form.
	IxoWasmDidPrefix = IxoDidPrefix + IxoWasmDidSubPrefix
)

var (
	// DidDocumentKey prefix for each key to a DidDocument
	DidDocumentKey = []byte{0x01}
)

// ReservedDidPrefixes lists DID prefixes owned by other ixo modules that
// mint their DIDs via IidKeeper.SetDidDocument directly. MsgCreateIidDocument
// must refuse any id with one of these prefixes — otherwise a user could
// front-run the owning module's deterministic sequence and squat a DID it
// will later try to mint, deadlocking the module.
//
// Each entry is a full prefix (no trailing identifier). New owning modules
// add their prefix here, not in the owning module itself, so the iid-module
// guard remains the single source of truth.
var ReservedDidPrefixes = []string{
	"did:ixo:entity:",
}

// IsReservedDid reports whether `id` falls under any module-reserved DID
// namespace. Returns the matching prefix on hit (for error context).
func IsReservedDid(id string) (string, bool) {
	for _, p := range ReservedDidPrefixes {
		if len(id) >= len(p) && id[:len(p)] == p {
			return p, true
		}
	}
	return "", false
}

// ValidateMsgCreateDIDForm enforces the policy for which DID forms a user
// (or a wasm contract acting as a user) may register via
// MsgCreateIidDocument:
//
//   - did:ixo:<bech32-account>       — the suffix MUST equal the signer's
//                                      bech32 address. Users can only
//                                      register their own account DID; they
//                                      cannot claim someone else's address.
//   - did:ixo:wasm:<bech32-contract> — anyone may register one of these for
//                                      now; the suffix after `wasm:` must
//                                      still be a valid bech32 address so
//                                      the chain cannot accumulate
//                                      junk-suffix entries that don't refer
//                                      to a real account.
//
// Reserved namespaces (currently did:ixo:entity:) are minted deterministically
// by their owning module via IidKeeper.SetDidDocument and are rejected here
// so external callers cannot front-run / squat them.
//
// Anything else — did:cosmos:..., did:x:..., did:ixo:foo:bar:baz,
// did:ixo:randomstring — is rejected with ErrDIDFormNotAllowed. Cosmos SDK
// message routing makes this the single chokepoint that catches Cosmwasm
// stargate messages and any other external caller; the keeper's
// CreateIidDocument path runs the same check as defense in depth (see
// msg_server.go).
func ValidateMsgCreateDIDForm(id, signer string) error {
	if prefix, reserved := IsReservedDid(id); reserved {
		return errorsmod.Wrapf(
			ErrReservedDidNamespace,
			"did %s is in reserved namespace %s", id, prefix,
		)
	}

	if !strings.HasPrefix(id, IxoDidPrefix) {
		return errorsmod.Wrapf(
			ErrDIDFormNotAllowed,
			"did %s is not in the did:ixo: namespace", id,
		)
	}

	suffix := strings.TrimPrefix(id, IxoDidPrefix)
	if suffix == "" {
		return errorsmod.Wrapf(
			ErrDIDFormNotAllowed,
			"did %s has empty method-specific id", id,
		)
	}

	// did:ixo:wasm:<contract-address>
	if contract, ok := strings.CutPrefix(suffix, IxoWasmDidSubPrefix); ok {
		// Reject did:ixo:wasm:something:else — wasm DIDs are a single
		// bech32-address segment after `wasm:`, not an arbitrarily-nested
		// namespace.
		if contract == "" || strings.ContainsRune(contract, ':') {
			return errorsmod.Wrapf(
				ErrDIDFormNotAllowed,
				"did %s: did:ixo:wasm: must be followed by a single bech32 address", id,
			)
		}
		if _, err := sdk.AccAddressFromBech32(contract); err != nil {
			return errorsmod.Wrapf(
				ErrDIDFormNotAllowed,
				"did:ixo:wasm: suffix %q is not a valid bech32 address: %s", contract, err,
			)
		}
		return nil
	}

	// did:ixo:<bech32-account> — must be a single segment (no sub-method)
	// and must equal the signer's address.
	if strings.ContainsRune(suffix, ':') {
		return errorsmod.Wrapf(
			ErrDIDFormNotAllowed,
			"did %s: only did:ixo:<account> and did:ixo:wasm:<contract> are allowed; "+
				"unrecognised sub-method", id,
		)
	}
	if _, err := sdk.AccAddressFromBech32(suffix); err != nil {
		return errorsmod.Wrapf(
			ErrDIDFormNotAllowed,
			"did:ixo:<account> suffix %q is not a valid bech32 address: %s", suffix, err,
		)
	}
	if suffix != signer {
		return errorsmod.Wrapf(
			ErrDIDAccountSignerMismatch,
			"did:ixo:<account> account %q must equal signer %q", suffix, signer,
		)
	}
	return nil
}

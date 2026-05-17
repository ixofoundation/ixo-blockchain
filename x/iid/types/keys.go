package types

const (
	// ModuleName defines the module name
	ModuleName = "iid"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// DidChainPrefix defines the did prefix for this chain
	DidChainPrefix = "did:x:"
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

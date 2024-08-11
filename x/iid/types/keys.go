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

	// DidKeyPrefix defines the did key prefix
	// DidKeyPrefix = "did:key:"

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability_did"
)

var (
	// DidDocumentKey prefix for each key to a DidDocument
	DidDocumentKey = []byte{0x01}
)

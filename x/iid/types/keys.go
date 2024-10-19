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

package types

const (
	// ModuleName defines the module name
	ModuleName = "claims"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_claims"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

var (
	CollectionKey = []byte{0x01}
	ClaimKey      = []byte{0x02}
	DisputeKey    = []byte{0x03}
)

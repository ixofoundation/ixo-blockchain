package types

const (
	// ModuleName is the name of this module
	ModuleName = "entity"

	// DefaultParamspace is the default param space for this module
	DefaultParamspace = ModuleName

	// StoreKey is the default store key for this module
	StoreKey = ModuleName

	// RouterKey is the message route for this module
	RouterKey = ModuleName

	// QuerierRoute is the querier route for this module's store.
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key.
	MemStoreKey = "mem_entity"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

var (
	EntityKey = []byte{0x01}
)

const (
	EntityAdminAccountName = "admin"
)

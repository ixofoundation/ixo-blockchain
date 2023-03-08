package types

const (
	ModuleName        = "entity"
	DefaultParamspace = ModuleName
	StoreKey          = ModuleName
	RouterKey         = ModuleName
	QuerierRoute      = ModuleName
	MemStoreKey       = "mem_entity"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

var (
	EntityKey = []byte{0x01}
)

var (
	EntityAdminAccountName = "admin"
)

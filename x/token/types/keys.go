package types

const (
	ModuleName        = "token"
	DefaultParamspace = ModuleName
	StoreKey          = ModuleName
	RouterKey         = ModuleName
	QuerierRoute      = ModuleName
	MemStoreKey       = "mem_token"
)

var (
	TokenKey = []byte{0x01}
)

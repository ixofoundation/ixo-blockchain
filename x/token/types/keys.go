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
	TokenKey           = []byte{0x01}
	TokenPropertiesKey = []byte{0x02}
)

const (
	TokenUriBase = "https://w3id.org/token/" // https://w3id.org/token/{stringportionofthetokendid}
)

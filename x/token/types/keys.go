package types

const (
	// ModuleName defines the module name
	ModuleName        = "token"
	DefaultParamspace = ModuleName

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName
)

var (
	TokenKey           = []byte{0x01}
	TokenPropertiesKey = []byte{0x02}
)

const (
	TokenUriBase = "https://w3id.org/token/" // https://w3id.org/token/{stringportionofthetokendid}
)

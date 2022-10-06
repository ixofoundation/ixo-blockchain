package types

const (
	ModuleName        = "entity"
	DefaultParamspace = ModuleName
	StoreKey          = ModuleName
	RouterKey         = ModuleName
	QuerierRoute      = ModuleName
	MemStoreKey       = "mem_entity"
)

var (
	EntityKey = []byte{0x01}
	// DocKey         = []byte{0x02}
	// AccountMapKey  = []byte{0x03}
	// WithdrawalsKey = []byte{0x04}
	// ClaimsKey      = []byte{0x05}
)

// func GetProjectKey(projectDid didexported.Did) []byte {
// 	return append(ProjectKey, []byte(projectDid)...)
// }

// func GetAccountMapKey(projectDid didexported.Did) []byte {
// 	return append(AccountMapKey, []byte(projectDid)...)
// }

// func GetWithdrawalsKey(projectDid didexported.Did) []byte {
// 	return append(WithdrawalsKey, []byte(projectDid)...)
// }

// func GetClaimsKey(projectDid didexported.Did) []byte {
// 	return append(ClaimsKey, []byte(projectDid)...)
// }

// func GetClaimKey(projectDid didexported.Did, claimId string) []byte {
// 	return append(GetClaimsKey(projectDid), []byte(claimId)...)
// }

type EntityConfigKey string

const (
	ConfigNftContractAddress EntityConfigKey = "nft_address"
)

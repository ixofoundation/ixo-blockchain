package types

import didexported "github.com/ixofoundation/ixo-blockchain/x/did/exported"

const (
	ModuleName        = "project"
	DefaultParamspace = ModuleName
	StoreKey          = ModuleName
	RouterKey         = ModuleName
	QuerierRoute      = ModuleName
)

var (
	ProjectKey     = []byte{0x01}
	DocKey         = []byte{0x02}
	AccountMapKey  = []byte{0x03}
	WithdrawalsKey = []byte{0x04}
	ClaimsKey      = []byte{0x05}
)

func GetProjectKey(projectDid didexported.Did) []byte {
	return append(ProjectKey, []byte(projectDid)...)
}

func GetAccountMapKey(projectDid didexported.Did) []byte {
	return append(AccountMapKey, []byte(projectDid)...)
}

func GetWithdrawalsKey(projectDid didexported.Did) []byte {
	return append(WithdrawalsKey, []byte(projectDid)...)
}

func GetClaimsKey(projectDid didexported.Did) []byte {
	return append(ClaimsKey, []byte(projectDid)...)
}

func GetClaimKey(projectDid didexported.Did, claimId string) []byte {
	return append(GetClaimsKey(projectDid), []byte(claimId)...)
}

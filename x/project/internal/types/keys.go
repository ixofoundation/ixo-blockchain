package types

import (
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

const (
	ModuleName        = "project"
	DefaultParamspace = ModuleName
	StoreKey          = ModuleName
	RouterKey         = StoreKey
	QuerierRoute      = RouterKey
)

var (
	ProjectKey    = []byte{0x01}
	AccountKey    = []byte{0x02}
	WithdrawalKey = []byte{0x03}
)

func GetProjectPrefixKey(did ixo.Did) []byte {
	return append(ProjectKey, []byte(did)...)
}

func GetAccountPrefixKey(did ixo.Did) []byte {
	return append(AccountKey, []byte(did)...)
}

func GetWithdrawalPrefixKey(did ixo.Did) []byte {
	return append(WithdrawalKey, []byte(did)...)
}

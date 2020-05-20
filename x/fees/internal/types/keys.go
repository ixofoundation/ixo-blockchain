package types

import "encoding/binary"

const (
	ModuleName        = "fees"
	DefaultParamspace = ModuleName
	StoreKey          = ModuleName
	RouterKey         = ModuleName
	QuerierRoute      = ModuleName
)

var (
	FeeKeyPrefix         = []byte{0x00}
	FeeContractKeyPrefix = []byte{0x01}

	FeeIdKey         = []byte{0x10}
	FeeContractIdKey = []byte{0x11}
)

func GetFeePrefixKey(feeId uint64) []byte {
	bz := make([]byte, 8)
	binary.LittleEndian.PutUint64(bz, feeId)
	return append(FeeKeyPrefix, bz...)
}

func GetFeeContractPrefixKey(feeContractId uint64) []byte {
	bz := make([]byte, 8)
	binary.LittleEndian.PutUint64(bz, feeContractId)
	return append(FeeContractKeyPrefix, bz...)
}

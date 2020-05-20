package types

const (
	ModuleName        = "fees"
	DefaultParamspace = ModuleName
	StoreKey          = ModuleName
	RouterKey         = ModuleName
	QuerierRoute      = ModuleName
)

var (
	FeeKey         = []byte{0x00}
	FeeContractKey = []byte{0x01}
)

func GetFeePrefixKey(feeId string) []byte {
	return append(FeeKey, []byte(feeId)...)
}

func GetFeeContractPrefixKey(feeContractId string) []byte {
	return append(FeeContractKey, []byte(feeContractId)...)
}

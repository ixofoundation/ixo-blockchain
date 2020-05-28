package types

const (
	ModuleName        = "fees"
	DefaultParamspace = ModuleName
	StoreKey          = ModuleName
	RouterKey         = ModuleName
	QuerierRoute      = ModuleName

	FeeRemainderPool = "fee_remainder_pool"

	FeePrefix          = "fee:"
	FeeContractPrefix  = FeePrefix + "contract:"
	SubscriptionPrefix = FeePrefix + "subscription:"
)

var (
	FeeKeyPrefix          = []byte{0x00}
	FeeContractKeyPrefix  = []byte{0x01}
	SubscriptionKeyPrefix = []byte{0x02}
)

func GetFeeKey(feeId string) []byte {
	return append(FeeKeyPrefix, []byte(feeId)...)
}

func GetFeeContractKey(feeContractId string) []byte {
	return append(FeeContractKeyPrefix, []byte(feeContractId)...)
}

func GetSubscriptionKey(subscriptionId string) []byte {
	return append(SubscriptionKeyPrefix, []byte(subscriptionId)...)
}

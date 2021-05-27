package types

const (
	ModuleName        = "payments"
	DefaultParamspace = ModuleName
	StoreKey          = ModuleName
	RouterKey         = ModuleName
	QuerierRoute      = ModuleName

	PayRemainderPool = "pay_remainder_pool"

	PaymentIdPrefix         = "payment:"
	PaymentTemplateIdPrefix = PaymentIdPrefix + "template:"
	PaymentContractIdPrefix = PaymentIdPrefix + "contract:"
	SubscriptionIdPrefix    = PaymentIdPrefix + "subscription:"
)

var (
	PaymentTemplateKeyPrefix = []byte{0x00}
	PaymentContractKeyPrefix = []byte{0x01}
	SubscriptionKeyPrefix    = []byte{0x02}
)

func GetPaymentTemplateKey(templateId string) []byte {
	return append(PaymentTemplateKeyPrefix, []byte(templateId)...)
}

func GetPaymentContractKey(contractId string) []byte {
	return append(PaymentContractKeyPrefix, []byte(contractId)...)
}

func GetSubscriptionKey(subscriptionId string) []byte {
	return append(SubscriptionKeyPrefix, []byte(subscriptionId)...)
}

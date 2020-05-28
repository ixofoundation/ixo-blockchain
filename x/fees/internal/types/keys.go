package types

import (
	"encoding/binary"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

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
	FeeKeyPrefix            = []byte{0x00}
	FeeContractKeyPrefix    = []byte{0x01}
	SubscriptionKeyPrefix   = []byte{0x02}
	DiscountHolderKeyPrefix = []byte{0x03}
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

func GetDiscountHolderKey(feeId string, discountId uint64, addr sdk.AccAddress) []byte {
	return append(GetDiscountHoldersKey(feeId, discountId), addr.Bytes()...)
}

func GetDiscountHoldersKey(feeId string, discountId uint64) []byte {
	bz := make([]byte, 8)
	binary.LittleEndian.PutUint64(bz, discountId)
	return append(GetDiscountsHoldersKey(feeId), bz...)
}

func GetDiscountsHoldersKey(feeId string) []byte {
	return append(DiscountHolderKeyPrefix, []byte(feeId)...)
}

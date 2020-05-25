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
)

var (
	FeeKeyPrefix            = []byte{0x00}
	FeeContractKeyPrefix    = []byte{0x01}
	SubscriptionKeyPrefix   = []byte{0x02}
	DiscountHolderKeyPrefix = []byte{0x03}

	FeeIdKey          = []byte{0x10}
	FeeContractIdKey  = []byte{0x11}
	SubscriptionIdKey = []byte{0x12}
)

func GetFeeKey(feeId uint64) []byte {
	bz := make([]byte, 8)
	binary.LittleEndian.PutUint64(bz, feeId)
	return append(FeeKeyPrefix, bz...)
}

func GetFeeContractKey(feeContractId uint64) []byte {
	bz := make([]byte, 8)
	binary.LittleEndian.PutUint64(bz, feeContractId)
	return append(FeeContractKeyPrefix, bz...)
}

func GetSubscriptionKey(subscriptionId uint64) []byte {
	bz := make([]byte, 8)
	binary.LittleEndian.PutUint64(bz, subscriptionId)
	return append(SubscriptionKeyPrefix, bz...)
}

func GetDiscountHolderKey(feeId uint64, discountId uint64, addr sdk.AccAddress) []byte {
	return append(GetDiscountHoldersKey(feeId, discountId), addr.Bytes()...)
}

func GetDiscountHoldersKey(feeId uint64, discountId uint64) []byte {
	bz := make([]byte, 8)
	binary.LittleEndian.PutUint64(bz, discountId)
	return append(GetDiscountsHoldersKey(feeId), bz...)
}

func GetDiscountsHoldersKey(feeId uint64) []byte {
	bz := make([]byte, 8)
	binary.LittleEndian.PutUint64(bz, feeId)
	return append(DiscountHolderKeyPrefix, bz...)
}

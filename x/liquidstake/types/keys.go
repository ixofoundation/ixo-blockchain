package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (
	// ModuleName defines the module name
	ModuleName = "liquidstake"

	// RouterKey is the message router key for the liquidstake module
	RouterKey = ModuleName

	// StoreKey is the default store key for the liquidstake module
	StoreKey = ModuleName

	// Epoch identifiers
	// TODO: comment out 2min it is for local testing
	AutocompoundEpoch = "2min"
	RebalanceEpoch    = "2min"
	// AutocompoundEpoch = "hour"
	// RebalanceEpoch = "day"
)

var (
	ParamsKey = []byte{0x01}
	// LiquidValidatorsKey defines prefix for each key to a liquid validator
	LiquidValidatorsKey = []byte{0x02}
)

// GetLiquidValidatorKey creates the key for the liquid validator with address
// VALUE: liquidstake/LiquidValidator
func GetLiquidValidatorKey(operatorAddr sdk.ValAddress) []byte {
	tmp := append([]byte{}, LiquidValidatorsKey...)
	return append(tmp, address.MustLengthPrefix(operatorAddr)...)
}

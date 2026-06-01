package types

import (
	"context"
	"regexp"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	ModuleName = "liquidstake"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	// Epoch identifiers used by the epochs module.
	//
	// For local liquidstake testing, swap to the "2min" identifiers below
	// (and uncomment the matching "2min" epoch entry in
	// x/epochs/types/genesis.go and the EpochIdentifier in
	// x/mint/types/params.go) so autocompound + rebalance hooks fire every
	// few minutes instead of hours/days. DO NOT enable for mainnet.
	AutocompoundEpoch = "hour"
	RebalanceEpoch    = "day"
	// AutocompoundEpoch = "2min"
	// RebalanceEpoch    = "2min"

	// MinPoolIDLength and MaxPoolIDLength bound the on-chain pool identifier.
	// 16 chars keeps storage keys compact and event/log output readable.
	MinPoolIDLength = 2
	MaxPoolIDLength = 16
)

// Storage key prefixes.
//
// 0x01 holds the (single) ModuleParams record.
// 0x10 + poolID            -> Pool
// 0x11 + lp(poolID) + valAddr -> LiquidValidator (scoped per pool)
//
// 0x02 was the pre-v7 LiquidValidator prefix; the v7 migration deletes any
// remaining 0x02 entries and re-keys them under 0x11.
var (
	ModuleParamsKey         = []byte{0x01}
	LegacyLiquidValidators_ = []byte{0x02} // historical, only referenced by the v7 migration

	PoolPrefix            = []byte{0x10}
	LiquidValidatorPrefix = []byte{0x11}
)

// poolIDRegex restricts pool identifiers to lowercase alphanumerics and dashes,
// with no leading/trailing dash. The bounds are enforced separately so the
// regex can stay anchored without character-count quantifiers.
var poolIDRegex = regexp.MustCompile(`^[a-z0-9](?:[a-z0-9-]*[a-z0-9])?$`)

// ibcDenomRegex matches the canonical IBC voucher format:
//
//	ibc/<64 uppercase hex chars>
//
// where the hex is the SHA-256 of a denomination trace path. The IBC
// transfer module is the only legitimate minter in this namespace, so we
// refuse to register a pool with a denom that looks like a voucher to
// guarantee no liquidstake/IBC denom collision can ever exist.
var ibcDenomRegex = regexp.MustCompile(`^ibc/[0-9A-F]{64}$`)

// IsIBCDenom reports whether denom matches the canonical IBC voucher format.
// Used by the liquidstake registerPool path and by the bank
// MintCoinsRestriction installed at app wiring.
func IsIBCDenom(denom string) bool {
	return ibcDenomRegex.MatchString(denom)
}

// mintAuthCtxKey is the context-value key the bank MintCoinsRestriction
// (installed in app/keepers/keepers.go) checks to confirm a mint of an LST
// denom originates from the liquidstake module's own LiquidStake handler.
//
// Cosmos SDK's MintingRestrictionFn signature does not receive the calling
// module name, so the only safe in-band way to distinguish "liquidstake
// minting its own LST" from "some other module minting an LST denom it
// shouldn't" is for liquidstake to stamp its own context immediately
// before calling bank.MintCoins. The sentinel value type is unexported
// so external code can't trivially construct or forge one.
type mintAuthCtxKeyT struct{}

var mintAuthCtxKey = mintAuthCtxKeyT{}

// AuthorizeLSTMintContext returns a new context tagged as authorised to
// mint LST denoms claimed by liquidstake pools. Call this in liquidstake's
// LiquidStake handler immediately before invoking bank.MintCoins.
func AuthorizeLSTMintContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, mintAuthCtxKey, true)
}

// IsLSTMintAuthorized reports whether the supplied context was tagged via
// AuthorizeLSTMintContext. Used by the bank MintCoinsRestriction to decide
// whether to allow a mint of an LST-claimed denom.
func IsLSTMintAuthorized(ctx context.Context) bool {
	v, _ := ctx.Value(mintAuthCtxKey).(bool)
	return v
}

// ValidatePoolID returns nil if id is a syntactically valid pool identifier.
// It does NOT check uniqueness — that is the keeper's responsibility.
func ValidatePoolID(id string) error {
	if l := len(id); l < MinPoolIDLength || l > MaxPoolIDLength {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest,
			"pool_id length must be in [%d,%d]: got %d", MinPoolIDLength, MaxPoolIDLength, l)
	}
	if !poolIDRegex.MatchString(id) {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest,
			"pool_id %q must be lowercase alphanumeric with optional internal dashes", id)
	}
	return nil
}

// GetPoolKey returns the KV store key for a Pool record. The pool_id is
// length-prefixed so that no pool's key is a byte-prefix of another's
// (defends against future code that uses GetPoolKey as an iteration prefix
// rather than a point lookup).
func GetPoolKey(poolID string) []byte {
	return append(append([]byte{}, PoolPrefix...), address.MustLengthPrefix([]byte(poolID))...)
}

// GetLiquidValidatorsByPoolPrefix returns the iteration prefix for every
// LiquidValidator belonging to the given pool.
func GetLiquidValidatorsByPoolPrefix(poolID string) []byte {
	out := append([]byte{}, LiquidValidatorPrefix...)
	return append(out, address.MustLengthPrefix([]byte(poolID))...)
}

// GetLiquidValidatorKey returns the KV store key for a single per-pool
// LiquidValidator record.
func GetLiquidValidatorKey(poolID string, operatorAddr sdk.ValAddress) []byte {
	prefix := GetLiquidValidatorsByPoolPrefix(poolID)
	return append(prefix, address.MustLengthPrefix(operatorAddr)...)
}

// PoolProxyAcc derives the 32-byte module sub-account that holds delegations
// for a given pool. Newly created pools always use this derivation; the v7
// migration overrides Pool.ProxyAccountAddress for the legacy "zero" pool to
// preserve pre-upgrade delegations under the original LiquidStakeProxyAcc.
func PoolProxyAcc(poolID string) sdk.AccAddress {
	return sdk.AccAddress(address.Module(ModuleName, []byte("-LiquidStakeProxyAcc-"+poolID)))
}

// LegacyLiquidStakeProxyAcc returns the pre-v7 single-pool proxy account
// address. Stored verbatim in the migrated "zero" pool's ProxyAccountAddress
// so existing delegations remain accessible without state migration.
func LegacyLiquidStakeProxyAcc() sdk.AccAddress {
	return sdk.AccAddress(address.Module(ModuleName, []byte("-LiquidStakeProxyAcc")))
}

// DummyFeeAccountAcc is a placeholder fee-collection address used when a pool
// is created without a fee account explicitly specified. Pools should override
// this with a real address before any rewards autocompound.
var DummyFeeAccountAcc = sdk.AccAddress(address.Module(ModuleName, []byte("-FeeAcc")))

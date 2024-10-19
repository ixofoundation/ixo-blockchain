package types

import (
	context "context"
	time "time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	entitytypes "github.com/ixofoundation/ixo-blockchain/v3/x/entity/types"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v3/x/iid/types"
)

type IidKeeper interface {
	GetDidDocument(ctx sdk.Context, key []byte) (iidtypes.IidDocument, bool)
}

type EntityKeeper interface {
	ResolveEntity(ctx sdk.Context, entityId string) (iidDocument iidtypes.IidDocument, entity entitytypes.Entity, err error)
	GetEntity(ctx sdk.Context, key []byte) (entitytypes.Entity, bool)
	CheckIfOwner(ctx sdk.Context, entityId, ownerAddress string) error
}

type AuthzKeeper interface {
	GetAuthorizations(ctx context.Context, grantee, granter sdk.AccAddress) ([]authz.Authorization, error)
	SaveGrant(ctx context.Context, grantee, granter sdk.AccAddress, authorization authz.Authorization, expiration *time.Time) error
	GetAuthorization(ctx context.Context, grantee, granter sdk.AccAddress, msgType string) (authz.Authorization, *time.Time)
}

type BankKeeper interface {
	HasBalance(ctx context.Context, addr sdk.AccAddress, amt sdk.Coin) bool
	InputOutputCoins(ctx context.Context, input banktypes.Input, outputs []banktypes.Output) error
	SendCoins(ctx context.Context, fromAddr, toAddr sdk.AccAddress, amt sdk.Coins) error
}

type WasmKeeper interface {
	Execute(ctx sdk.Context, contractAddress sdk.AccAddress, caller sdk.AccAddress, msg []byte, coins sdk.Coins) ([]byte, error)
}

type AccountKeeper interface {
	NewAccountWithAddress(context.Context, sdk.AccAddress) sdk.AccountI
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI
	SetAccount(context.Context, sdk.AccountI)
}

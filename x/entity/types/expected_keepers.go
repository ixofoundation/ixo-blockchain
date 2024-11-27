package types

import (
	context "context"
	time "time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v4/x/iid/types"
)

type AccountKeeper interface {
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI
	SetAccount(context.Context, sdk.AccountI)
	NewAccountWithAddress(context.Context, sdk.AccAddress) sdk.AccountI
}

type AuthzKeeper interface {
	SaveGrant(ctx context.Context, grantee, granter sdk.AccAddress, authorization authz.Authorization, expiration *time.Time) error
	DeleteGrant(ctx context.Context, grantee, granter sdk.AccAddress, msgType string) error
}

type IidKeeper interface {
	ResolveDid(ctx sdk.Context, did iidtypes.DID) (doc iidtypes.IidDocument, err error)
	GetDidDocument(ctx sdk.Context, key []byte) (iidtypes.IidDocument, bool)
	SetDidDocument(ctx sdk.Context, key []byte, document iidtypes.IidDocument)
}

type WasmKeeper interface {
	Instantiate(
		ctx sdk.Context,
		codeID uint64,
		creator, admin sdk.AccAddress,
		initMsg []byte,
		label string,
		deposit sdk.Coins,
	) (sdk.AccAddress, []byte, error)
	Execute(ctx sdk.Context, contractAddress sdk.AccAddress, caller sdk.AccAddress, msg []byte, coins sdk.Coins) ([]byte, error)
}

type WasmViewKeeper interface {
	QuerySmart(ctx context.Context, contractAddress sdk.AccAddress, queryMsg []byte) ([]byte, error)
}

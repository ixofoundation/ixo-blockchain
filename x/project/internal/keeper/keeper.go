package keeper

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/ixofoundation/ixo-blockchain/x/payments"

	"github.com/ixofoundation/ixo-blockchain/x/did"
	"github.com/ixofoundation/ixo-blockchain/x/project/internal/types"
)

type Keeper struct {
	cdc            *codec.Codec
	storeKey       sdk.StoreKey
	paramSpace     params.Subspace
	AccountKeeper  auth.AccountKeeper
	DidKeeper      did.Keeper
	paymentsKeeper payments.Keeper
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, paramSpace params.Subspace,
	accountKeeper auth.AccountKeeper, didKeeper did.Keeper,
	paymentsKeeper payments.Keeper) Keeper {
	return Keeper{
		cdc:            cdc,
		storeKey:       key,
		paramSpace:     paramSpace.WithKeyTable(types.ParamKeyTable()),
		AccountKeeper:  accountKeeper,
		DidKeeper:      didKeeper,
		paymentsKeeper: paymentsKeeper,
	}
}

// GetParams returns the total set of project parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the total set of project parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

func (k Keeper) GetProjectDocIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.ProjectKey)
}

func (k Keeper) MustGetProjectDocByKey(ctx sdk.Context, key []byte) types.StoredProjectDoc {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(key) {
		panic("project doc not found")
	}

	bz := store.Get(key)
	var projectDoc types.ProjectDoc
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &projectDoc)

	return &projectDoc
}

func (k Keeper) ProjectDocExists(ctx sdk.Context, projectDid did.Did) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetProjectKey(projectDid))
}

func (k Keeper) GetProjectDoc(ctx sdk.Context, projectDid did.Did) (types.StoredProjectDoc, sdk.Error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetProjectKey(projectDid)

	bz := store.Get(key)
	if bz == nil {
		return nil, did.ErrorInvalidDid(types.DefaultCodespace, "Invalid ProjectDid Address")
	}

	var projectDoc types.ProjectDoc
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &projectDoc)

	return &projectDoc, nil
}

func (k Keeper) SetProjectDoc(ctx sdk.Context, projectDoc types.StoredProjectDoc) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetProjectKey(projectDoc.GetProjectDid())
	store.Set(key, k.cdc.MustMarshalBinaryLengthPrefixed(projectDoc))
}

func (k Keeper) SetAccountMap(ctx sdk.Context, projectDid did.Did, accountMap types.AccountMap) {
	store := ctx.KVStore(k.storeKey)
	bz, err := json.Marshal(accountMap)
	if err != nil {
		panic(err)
	}
	store.Set(types.GetAccountMapKey(projectDid), bz)
}

func (k Keeper) GetAccountMap(ctx sdk.Context, projectDid did.Did) types.AccountMap {
	store := ctx.KVStore(k.storeKey)
	key := types.GetAccountMapKey(projectDid)

	bz := store.Get(key)
	if bz == nil {
		return make(types.AccountMap)
	} else {
		var accountMap types.AccountMap
		if err := json.Unmarshal(bz, &accountMap); err != nil {
			panic(err)
		}

		return accountMap
	}
}

func (k Keeper) AddAccountToProjectAccounts(ctx sdk.Context, projectDid did.Did,
	accountId types.InternalAccountID, account auth.Account) {
	accountMap := k.GetAccountMap(ctx, projectDid)
	_, found := accountMap[accountId]
	if found {
		return
	}

	store := ctx.KVStore(k.storeKey)
	key := types.GetAccountMapKey(projectDid)
	accountMap[accountId] = account.GetAddress()

	bz, err := json.Marshal(accountMap)
	if err != nil {
		panic(err)
	}

	store.Set(key, bz)
}

func (k Keeper) CreateNewAccount(ctx sdk.Context, projectDid did.Did,
	accountId types.InternalAccountID) (auth.Account, sdk.Error) {
	address := supply.NewModuleAddress(accountId.ToAddressKey(projectDid))

	if k.AccountKeeper.GetAccount(ctx, address) != nil {
		return nil, sdk.ErrInvalidAddress("Generate account already exists")
	}

	account := k.AccountKeeper.NewAccountWithAddress(ctx, address)
	k.AccountKeeper.SetAccount(ctx, account)

	return account, nil
}

func (k Keeper) SetProjectWithdrawalTransactions(ctx sdk.Context, projectDid did.Did, txs []types.WithdrawalInfo) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(txs)
	store.Set(types.GetWithdrawalsKey(projectDid), bz)
}

func (k Keeper) GetProjectWithdrawalTransactions(ctx sdk.Context, projectDid did.Did) ([]types.WithdrawalInfo, sdk.Error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetWithdrawalsKey(projectDid)

	bz := store.Get(key)
	if bz == nil {
		return []types.WithdrawalInfo{}, did.ErrorInvalidDid(types.DefaultCodespace, "ProjectDoc doesn't exist")
	} else {
		var txs []types.WithdrawalInfo
		k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &txs)

		return txs, nil
	}
}

func (k Keeper) AddProjectWithdrawalTransaction(ctx sdk.Context, projectDid did.Did, info types.WithdrawalInfo) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetWithdrawalsKey(projectDid)

	txs, _ := k.GetProjectWithdrawalTransactions(ctx, projectDid)
	txs = append(txs, info)

	store.Set(key, k.cdc.MustMarshalBinaryLengthPrefixed(txs))
}

func (k Keeper) GetClaimIterator(ctx sdk.Context, projectDid did.Did) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.GetClaimsKey(projectDid))
}

func (k Keeper) MustGetClaimByKey(ctx sdk.Context, key []byte) types.Claim {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(key) {
		panic("claim not found")
	}

	bz := store.Get(key)
	var claim types.Claim
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &claim)

	return claim
}

func (k Keeper) ClaimExists(ctx sdk.Context, projectDid did.Did, claimId string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetClaimKey(projectDid, claimId))
}

func (k Keeper) GetClaim(ctx sdk.Context, projectDid did.Did, claimId string) (types.Claim, sdk.Error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetClaimKey(projectDid, claimId)

	bz := store.Get(key)
	if bz == nil {
		return types.Claim{}, did.ErrorInvalidDid(types.DefaultCodespace, "claim not found")
	}

	var claim types.Claim
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &claim)

	return claim, nil
}

func (k Keeper) SetClaim(ctx sdk.Context, projectDid did.Did, claim types.Claim) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetClaimKey(projectDid, claim.Id)
	store.Set(key, k.cdc.MustMarshalBinaryLengthPrefixed(claim))
}

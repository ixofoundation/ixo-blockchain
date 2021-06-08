package keeper

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	didexported "github.com/ixofoundation/ixo-blockchain/x/did/exported"
	didtypes "github.com/ixofoundation/ixo-blockchain/x/did/types"

	"github.com/ixofoundation/ixo-blockchain/x/project/types"
)

type Keeper struct {
	cdc            codec.BinaryMarshaler
	storeKey       sdk.StoreKey
	paramSpace     paramstypes.Subspace
	accountKeeper  types.AccountKeeper
	didKeeper      types.DidKeeper
	paymentsKeeper types.PaymentsKeeper
	bankKeeper     types.BankKeeper
}

func NewKeeper(cdc codec.BinaryMarshaler, key sdk.StoreKey, paramSpace paramstypes.Subspace,
	accountKeeper types.AccountKeeper, didKeeper types.DidKeeper,
	paymentsKeeper types.PaymentsKeeper, bankKeeper types.BankKeeper) Keeper {
	return Keeper{
		cdc:            cdc,
		storeKey:       key,
		paramSpace:     paramSpace.WithKeyTable(types.ParamKeyTable()),
		accountKeeper:  accountKeeper,
		didKeeper:      didKeeper,
		paymentsKeeper: paymentsKeeper,
		bankKeeper:     bankKeeper,
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

func (k Keeper) MustGetProjectDocByKey(ctx sdk.Context, key []byte) types.ProjectDoc {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(key) {
		panic("project doc not found")
	}

	bz := store.Get(key)
	var projectDoc types.ProjectDoc
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &projectDoc)

	return projectDoc
}

func (k Keeper) ProjectDocExists(ctx sdk.Context, projectDid didexported.Did) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetProjectKey(projectDid))
}

func (k Keeper) GetProjectDoc(ctx sdk.Context, projectDid didexported.Did) (types.ProjectDoc, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetProjectKey(projectDid)

	bz := store.Get(key)
	if bz == nil {
		return types.ProjectDoc{}, sdkerrors.Wrap(didtypes.ErrInvalidDid, projectDid)
	}

	var projectDoc types.ProjectDoc
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &projectDoc)

	return projectDoc, nil
}

func (k Keeper) ValidateProjectFeesMap(ctx sdk.Context, projectFeesMap types.ProjectFeesMap) error {
	for _, v := range projectFeesMap.Items {
		_, err := k.paymentsKeeper.GetPaymentTemplate(ctx, v.PaymentTemplateId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (k Keeper) SetProjectDoc(ctx sdk.Context, projectDoc types.ProjectDoc) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetProjectKey(projectDoc.ProjectDid)
	store.Set(key, k.cdc.MustMarshalBinaryLengthPrefixed(&projectDoc))
}

func (k Keeper) SetAccountMap(ctx sdk.Context, projectDid didexported.Did, accountMap types.AccountMap) {
	store := ctx.KVStore(k.storeKey)
	bz, err := json.Marshal(accountMap)
	if err != nil {
		panic(err)
	}
	store.Set(types.GetAccountMapKey(projectDid), bz)
}

func (k Keeper) GetAccountMap(ctx sdk.Context, projectDid didexported.Did) types.AccountMap {
	store := ctx.KVStore(k.storeKey)
	key := types.GetAccountMapKey(projectDid)

	bz := store.Get(key)
	if bz == nil {
		return types.AccountMap{Map: make(map[string]string)} //make(types.AccountMap)
	} else {
		var accountMap types.AccountMap
		if err := json.Unmarshal(bz, &accountMap); err != nil {
			panic(err)
		}

		return accountMap
	}
}

func (k Keeper) AddAccountToProjectAccounts(ctx sdk.Context, projectDid didexported.Did,
	accountId types.InternalAccountID, account authtypes.AccountI) {
	strAccountId := string(accountId)
	accountMap := k.GetAccountMap(ctx, projectDid)
	_, found := accountMap.Map[strAccountId]
	if found {
		return
	}

	store := ctx.KVStore(k.storeKey)
	key := types.GetAccountMapKey(projectDid)
	accountMap.Map[strAccountId] = account.GetAddress().String()

	bz, err := json.Marshal(accountMap)
	if err != nil {
		panic(err)
	}

	store.Set(key, bz)
}

func (k Keeper) CreateNewAccount(ctx sdk.Context, projectDid didexported.Did,
	accountId types.InternalAccountID) (authtypes.AccountI, error) {
	address := authtypes.NewModuleAddress(accountId.ToAddressKey(projectDid))

	if k.accountKeeper.GetAccount(ctx, address) != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "account already exists")
	}

	account := k.accountKeeper.NewAccountWithAddress(ctx, address)
	k.accountKeeper.SetAccount(ctx, account)

	return account, nil
}

func (k Keeper) SetProjectWithdrawalTransactions(ctx sdk.Context, projectDid didexported.Did, txs types.WithdrawalInfoDocs) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(&txs)
	store.Set(types.GetWithdrawalsKey(projectDid), bz)
}

func (k Keeper) GetProjectWithdrawalTransactions(ctx sdk.Context, projectDid didexported.Did) (types.WithdrawalInfoDocs, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetWithdrawalsKey(projectDid)

	bz := store.Get(key)
	if bz == nil {
		return types.WithdrawalInfoDocs{}, sdkerrors.Wrap(didtypes.ErrInvalidDid, "project does not exist")
	} else {
		var txs types.WithdrawalInfoDocs
		k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &txs)

		return txs, nil
	}
}

func (k Keeper) AddProjectWithdrawalTransaction(ctx sdk.Context, projectDid didexported.Did, info types.WithdrawalInfoDoc) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetWithdrawalsKey(projectDid)

	txs, _ := k.GetProjectWithdrawalTransactions(ctx, projectDid)
	txs = types.AppendWithdrawalInfoDocs(txs, info) //append(txs, info)

	store.Set(key, k.cdc.MustMarshalBinaryLengthPrefixed(&txs))
}

func (k Keeper) GetClaimIterator(ctx sdk.Context, projectDid didexported.Did) sdk.Iterator {
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

func (k Keeper) ClaimExists(ctx sdk.Context, projectDid didexported.Did, claimId string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetClaimKey(projectDid, claimId))
}

func (k Keeper) GetClaim(ctx sdk.Context, projectDid didexported.Did, claimId string) (types.Claim, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetClaimKey(projectDid, claimId)

	bz := store.Get(key)
	if bz == nil {
		return types.Claim{}, fmt.Errorf("claim not found")
	}

	var claim types.Claim
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &claim)

	return claim, nil
}

func (k Keeper) SetClaim(ctx sdk.Context, projectDid didexported.Did, claim types.Claim) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetClaimKey(projectDid, claim.Id)
	store.Set(key, k.cdc.MustMarshalBinaryLengthPrefixed(&claim))
}

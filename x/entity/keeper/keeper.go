package keeper

import (
	"errors"
	"fmt"
	"time"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	cryptosecp256k1 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/entity/types"
	entitycontracts "github.com/ixofoundation/ixo-blockchain/x/entity/types/contracts"
	iidkeeper "github.com/ixofoundation/ixo-blockchain/x/iid/keeper"
	iidtypes "github.com/ixofoundation/ixo-blockchain/x/iid/types"
)

type Keeper struct {
	cdc           codec.BinaryCodec
	storeKey      sdk.StoreKey
	memStoreKey   sdk.StoreKey
	IidKeeper     iidkeeper.Keeper
	WasmKeeper    wasmtypes.ContractOpsKeeper
	AccountKeeper authkeeper.AccountKeeper
	ParamsKeeper  paramskeeper.Keeper
}

func NewKeeper(cdc codec.BinaryCodec, key sdk.StoreKey, memStoreKey sdk.StoreKey, iidKeeper iidkeeper.Keeper, wasmKeeper wasmkeeper.Keeper,
	accountKeeper authkeeper.AccountKeeper, paramsKeeper paramskeeper.Keeper) Keeper {
	return Keeper{
		cdc:           cdc,
		storeKey:      key,
		memStoreKey:   memStoreKey,
		IidKeeper:     iidKeeper,
		WasmKeeper:    wasmkeeper.NewDefaultPermissionKeeper(wasmKeeper),
		AccountKeeper: accountKeeper,
		ParamsKeeper:  paramsKeeper,
	}
}

func (k Keeper) UpdateEntityConfig(ctx sdk.Context, key string, value string) error {
	ctx.KVStore(k.storeKey).Set(nil, nil)
	return nil
}

func (k Keeper) GetEntityConfig(ctx sdk.Context, key types.EntityConfigKey) ([]byte, error) {
	val := ctx.KVStore(k.storeKey).Get([]byte(key))
	if val == nil {
		return nil, sdkerrors.Wrap(errors.New("not found"), "Not found")
	}
	return val, nil
}

func (k Keeper) CreateEntity(ctx sdk.Context, msg *types.MsgCreateEntity) (types.MsgCreateEntityResponse, error) {

	signer, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		return types.MsgCreateEntityResponse{}, err
	}

	fmt.Println("================================================1")
	fmt.Printf("%+v\n", msg)

	privKey := cryptosecp256k1.GenPrivKey()
	pubKey := privKey.PubKey()

	fmt.Println("================================================12")
	fmt.Printf("%+v\n", string(pubKey.Address()))
	address := sdk.AccAddress(pubKey.Address().Bytes())
	// if err != nil {
	// 	fmt.Println("================================================1e")
	// 	fmt.Printf("%s\n", err)
	// 	return err
	// }

	fmt.Println("================================================11")
	fmt.Printf("%s\n", string(address.String()))

	account := k.AccountKeeper.NewAccount(ctx, authtypes.NewBaseAccount(address, pubKey, 0, 0))
	entityId := fmt.Sprintf("did:ixo:entity:%s:%s", msg.EntityType, address.String())

	verification := iidtypes.NewAccountVerification(
		iidtypes.DID(entityId),
		ctx.ChainID(),
		account.GetAddress().String(),
		iidtypes.Authentication,
	)

	fmt.Println("================================================2")
	fmt.Printf("%s\n", entityId)

	did, err := iidtypes.NewDidDocument(entityId,
		iidtypes.WithServices(msg.Service...),
		iidtypes.WithRights(msg.AccordedRight...),
		iidtypes.WithResources(msg.LinkedResource...),
		iidtypes.WithVerifications(append(msg.Verification, verification)...),
		iidtypes.WithControllers(append(msg.Controller, entityId, msg.OwnerDid)...),
	)
	if err != nil {
		fmt.Println("================================================3")
		fmt.Printf("%s\n", err)
		// k.Logger(ctx).Error(err.Error())
		return types.MsgCreateEntityResponse{}, err
	}

	fmt.Println("================================================3")
	fmt.Printf("%s\n", entityId)
	fmt.Printf("%+v\n", did)

	// check that the did is not already taken
	_, found := k.IidKeeper.GetDidDocument(ctx, []byte(entityId))
	if found {
		err := sdkerrors.Wrapf(iidtypes.ErrDidDocumentFound, "a document with did %s already exists", entityId)
		// k.Logger(ctx).Error(err.Error())
		return types.MsgCreateEntityResponse{}, err
	}

	// persist the did document

	currentTimeUtc := time.Now().UTC()
	// now create and persist the metadata
	didM := iidtypes.NewDidMetadata(ctx.TxBytes(), ctx.BlockTime())
	didM.EntityType = msg.EntityType
	didM.Deactivated = msg.Deactivated
	didM.Created = &currentTimeUtc
	didM.Updated = &currentTimeUtc
	didM.VersionId = fmt.Sprintf("%s:%d", entityId, 0)
	didM.Stage = msg.Stage
	didM.Credentials = msg.VerifiableCredential
	didM.VerifiableCredential = msg.VerificationStatus
	didM.StartDate = msg.StartDate
	didM.EndDate = msg.EndDate
	didM.RelayerNode = msg.RelayerNode

	fmt.Println("===============================================4")
	fmt.Printf("%+v\n", did)

	k.IidKeeper.SetDidDocument(ctx, []byte(entityId), did)
	k.IidKeeper.SetDidMetadata(ctx, []byte(entityId), didM)

	//nftAddresBytes, err := k.GetEntityConfig(ctx, types.ConfigNftContractAddress)
	addr := []byte("ixo14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9sqa3vn7")
	if err != nil {
		return types.MsgCreateEntityResponse{}, err
	}
	//ixo14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9sqa3vn7

	nftMsgBytes, err := entitycontracts.Mint{
		TokenId:  did.Id,
		Owner:    did.Id,
		TokenUrl: "http://nft.com",
	}.Marshal()
	if err != nil {
		return types.MsgCreateEntityResponse{}, err
	}

	_, err = k.WasmKeeper.Execute(ctx, sdk.AccAddress(addr), signer, nftMsgBytes, sdk.NewCoins(sdk.NewCoin("uixo", sdk.ZeroInt())))
	if err != nil {
		return types.MsgCreateEntityResponse{}, err
	}

	// emit the event
	if err := ctx.EventManager().EmitTypedEvents(iidtypes.NewIidDocumentCreatedEvent(entityId, msg.OwnerDid)); err != nil {
		// k.Logger(ctx).Error("failed to emit DidDocumentCreatedEvent", "did", msg.Id, "signer", msg.Signer, "err", err)
		return types.MsgCreateEntityResponse{}, err
	}

	resp := types.MsgCreateEntityResponse{
		EntityId:     entityId,
		EntityType:   didM.EntityType,
		EntityStatus: didM.Status,
	}
	return resp, err
}

// // GetParams returns the total set of project parameters.
// func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
// 	k.paramSpace.GetParamSet(ctx, &params)
// 	return params
// }

// // SetParams sets the total set of project parameters.
// func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
// 	k.paramSpace.SetParamSet(ctx, &params)
// }

func (k Keeper) GetEntityDocIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.EntityKey)
}

// func (k Keeper) MustGetProjectDocByKey(ctx sdk.Context, key []byte) types.ProjectDoc {
// 	store := ctx.KVStore(k.storeKey)
// 	if !store.Has(key) {
// 		panic("project doc not found")
// 	}

// 	bz := store.Get(key)
// 	var projectDoc types.ProjectDoc
// 	k.cdc.MustUnmarshalLengthPrefixed(bz, &projectDoc)

// 	return projectDoc
// }

func (k Keeper) EntityExists(ctx sdk.Context, entityDid string) bool {
	// store := ctx.KVStore(k.storeKey)
	_, exists := k.IidKeeper.GetDidDocument(ctx, []byte(entityDid))
	return exists
}

// func (k Keeper) GetEntityDoc(ctx sdk.Context, projectDid didexported.Did) (types.ProjectDoc, error) {
// 	store := ctx.KVStore(k.storeKey)
// 	key := types.GetProjectKey(projectDid)

// 	bz := store.Get(key)
// 	if bz == nil {
// 		return types.ProjectDoc{}, sdkerrors.Wrap(didtypes.ErrInvalidDid, projectDid)
// 	}

// 	var projectDoc types.ProjectDoc
// 	k.cdc.MustUnmarshalLengthPrefixed(bz, &projectDoc)

// 	return projectDoc, nil
// }

// func (k Keeper) ValidateProjectFeesMap(ctx sdk.Context, projectFeesMap types.ProjectFeesMap) error {
// 	for _, v := range projectFeesMap.Items {
// 		_, err := k.paymentsKeeper.GetPaymentTemplate(ctx, v.PaymentTemplateId)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func (k Keeper) SetProjectDoc(ctx sdk.Context, projectDoc types.ProjectDoc) {
// 	store := ctx.KVStore(k.storeKey)
// 	key := types.GetProjectKey(projectDoc.ProjectDid)
// 	store.Set(key, k.cdc.MustMarshalLengthPrefixed(&projectDoc))
// }

// func (k Keeper) SetAccountMap(ctx sdk.Context, projectDid didexported.Did, accountMap types.AccountMap) {
// 	store := ctx.KVStore(k.storeKey)
// 	bz, err := json.Marshal(accountMap)
// 	if err != nil {
// 		panic(err)
// 	}
// 	store.Set(types.GetAccountMapKey(projectDid), bz)
// }

// func (k Keeper) GetAccountMap(ctx sdk.Context, projectDid didexported.Did) types.AccountMap {
// 	store := ctx.KVStore(k.storeKey)
// 	key := types.GetAccountMapKey(projectDid)

// 	bz := store.Get(key)
// 	if bz == nil {
// 		return types.AccountMap{Map: make(map[string]string)} //make(types.AccountMap)
// 	} else {
// 		var accountMap types.AccountMap
// 		if err := json.Unmarshal(bz, &accountMap); err != nil {
// 			panic(err)
// 		}

// 		return accountMap
// 	}
// }

// func (k Keeper) AddAccountToProjectAccounts(ctx sdk.Context, projectDid didexported.Did,
// 	accountId types.InternalAccountID, account authtypes.AccountI) {
// 	strAccountId := string(accountId)
// 	accountMap := k.GetAccountMap(ctx, projectDid)
// 	_, found := accountMap.Map[strAccountId]
// 	if found {
// 		return
// 	}

// 	store := ctx.KVStore(k.storeKey)
// 	key := types.GetAccountMapKey(projectDid)
// 	accountMap.Map[strAccountId] = account.GetAddress().String()

// 	bz, err := json.Marshal(accountMap)
// 	if err != nil {
// 		panic(err)
// 	}

// 	store.Set(key, bz)
// }

// func (k Keeper) CreateNewAccount(ctx sdk.Context, projectDid didexported.Did,
// 	accountId types.InternalAccountID) (authtypes.AccountI, error) {
// 	address := authtypes.NewModuleAddress(accountId.ToAddressKey(projectDid))

// 	if k.AccountKeeper.GetAccount(ctx, address) != nil {
// 		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "account already exists")
// 	}

// 	account := k.AccountKeeper.NewAccountWithAddress(ctx, address)
// 	k.AccountKeeper.SetAccount(ctx, account)

// 	return account, nil
// }

// func (k Keeper) SetProjectWithdrawalTransactions(ctx sdk.Context, projectDid didexported.Did, txs types.WithdrawalInfoDocs) {
// 	store := ctx.KVStore(k.storeKey)
// 	bz := k.cdc.MustMarshalLengthPrefixed(&txs)
// 	store.Set(types.GetWithdrawalsKey(projectDid), bz)
// }

// func (k Keeper) GetProjectWithdrawalTransactions(ctx sdk.Context, projectDid didexported.Did) (types.WithdrawalInfoDocs, error) {
// 	store := ctx.KVStore(k.storeKey)
// 	key := types.GetWithdrawalsKey(projectDid)

// 	bz := store.Get(key)
// 	if bz == nil {
// 		return types.WithdrawalInfoDocs{}, sdkerrors.Wrap(didtypes.ErrInvalidDid, "project does not exist")
// 	} else {
// 		var txs types.WithdrawalInfoDocs
// 		k.cdc.MustUnmarshalLengthPrefixed(bz, &txs)

// 		return txs, nil
// 	}
// }

// func (k Keeper) AddProjectWithdrawalTransaction(ctx sdk.Context, projectDid didexported.Did, info types.WithdrawalInfoDoc) {
// 	store := ctx.KVStore(k.storeKey)
// 	key := types.GetWithdrawalsKey(projectDid)

// 	txs, _ := k.GetProjectWithdrawalTransactions(ctx, projectDid)
// 	txs = types.AppendWithdrawalInfoDocs(txs, info) //append(txs, info)

// 	store.Set(key, k.cdc.MustMarshalLengthPrefixed(&txs))
// }

// func (k Keeper) GetClaimIterator(ctx sdk.Context, projectDid didexported.Did) sdk.Iterator {
// 	store := ctx.KVStore(k.storeKey)
// 	return sdk.KVStorePrefixIterator(store, types.GetClaimsKey(projectDid))
// }

// func (k Keeper) MustGetClaimByKey(ctx sdk.Context, key []byte) types.Claim {
// 	store := ctx.KVStore(k.storeKey)
// 	if !store.Has(key) {
// 		panic("claim not found")
// 	}

// 	bz := store.Get(key)
// 	var claim types.Claim
// 	k.cdc.MustUnmarshalLengthPrefixed(bz, &claim)

// 	return claim
// }

// func (k Keeper) ClaimExists(ctx sdk.Context, projectDid didexported.Did, claimId string) bool {
// 	store := ctx.KVStore(k.storeKey)
// 	return store.Has(types.GetClaimKey(projectDid, claimId))
// }

// func (k Keeper) GetClaim(ctx sdk.Context, projectDid didexported.Did, claimId string) (types.Claim, error) {
// 	store := ctx.KVStore(k.storeKey)
// 	key := types.GetClaimKey(projectDid, claimId)

// 	bz := store.Get(key)
// 	if bz == nil {
// 		return types.Claim{}, fmt.Errorf("claim not found")
// 	}

// 	var claim types.Claim
// 	k.cdc.MustUnmarshalLengthPrefixed(bz, &claim)

// 	return claim, nil
// }

// func (k Keeper) SetClaim(ctx sdk.Context, projectDid didexported.Did, claim types.Claim) {
// 	store := ctx.KVStore(k.storeKey)
// 	key := types.GetClaimKey(projectDid, claim.Id)
// 	store.Set(key, k.cdc.MustMarshalLengthPrefixed(&claim))
// }

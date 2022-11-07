package keeper

import (
	"errors"
	"fmt"
	"time"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/btcsuite/btcutil/base58"
	"github.com/cosmos/cosmos-sdk/codec"
	cryptosecp256k1 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
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
	ParamSpace    paramstypes.Subspace
}

func NewKeeper(cdc codec.BinaryCodec, key sdk.StoreKey, memStoreKey sdk.StoreKey, iidKeeper iidkeeper.Keeper, wasmKeeper wasmkeeper.Keeper,
	accountKeeper authkeeper.AccountKeeper, paramSpace paramstypes.Subspace) Keeper {

	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:           cdc,
		storeKey:      key,
		memStoreKey:   memStoreKey,
		IidKeeper:     iidKeeper,
		WasmKeeper:    wasmkeeper.NewDefaultPermissionKeeper(wasmKeeper),
		AccountKeeper: accountKeeper,
		ParamSpace:    paramSpace,
	}
}

func (k Keeper) UpdateEntityConfig(ctx sdk.Context, key string, value string) error {
	ctx.KVStore(k.storeKey).Set(nil, nil)
	return nil
}

func (k Keeper) CreateEntity(ctx sdk.Context, msg *types.MsgCreateEntity) (types.MsgCreateEntityResponse, error) {
	params := k.GetParams(ctx)
	nftContractAddressParam := params.NftContractAddress

	fmt.Println("==============\nnftContractAddressParam", nftContractAddressParam)
	if len(nftContractAddressParam) == 0 {
		return types.MsgCreateEntityResponse{}, errors.New("nftContractAddress not set")
	}

	nftContractAddress, err := sdk.AccAddressFromBech32(nftContractAddressParam)
	if err != nil {
		return types.MsgCreateEntityResponse{}, err
	}

	privKey := cryptosecp256k1.GenPrivKey()
	pubKey := privKey.PubKey()
	address, err := sdk.AccAddressFromBech32(params.NftContractMinter)
	if err != nil {
		return types.MsgCreateEntityResponse{}, err
	}

	account := k.AccountKeeper.NewAccount(ctx, authtypes.NewBaseAccount(address, pubKey, 0, 0))
	entityId := fmt.Sprintf("did:ixo:entity:%s:%s", msg.EntityType, base58.Encode(pubKey.Bytes()[:16]))

	verification := iidtypes.NewVerification(
		iidtypes.NewVerificationMethod(
			entityId,
			iidtypes.DID(entityId),
			iidtypes.NewBlockchainAccountID(ctx.ChainID(), account.GetAddress().String()),
		),
		[]string{iidtypes.Authentication},
		nil,
	)

	defaultVerification := iidtypes.NewVerification(
		iidtypes.NewVerificationMethod(
			iidtypes.DID(entityId).NewVerificationMethodID(account.GetAddress().String()),
			iidtypes.DID(entityId),
			iidtypes.NewBlockchainAccountID(ctx.ChainID(), account.GetAddress().String()),
		),
		[]string{iidtypes.Authentication},
		nil,
	)

	did, err := iidtypes.NewDidDocument(entityId,
		iidtypes.WithServices(msg.Service...),
		iidtypes.WithRights(msg.AccordedRight...),
		iidtypes.WithResources(msg.LinkedResource...),
		iidtypes.WithVerifications(append(msg.Verification, defaultVerification, verification)...),
		iidtypes.WithControllers(append(msg.Controller, entityId, msg.OwnerDid)...),
	)
	if err != nil {
		// k.Logger(ctx).Error(err.Error())
		return types.MsgCreateEntityResponse{}, err
	}

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
	did.Context = msg.Context

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

	k.IidKeeper.SetDidDocument(ctx, []byte(entityId), did)
	k.IidKeeper.SetDidMetadata(ctx, []byte(entityId), didM)

	//nftAddresBytes, err := k.GetEntityConfig(ctx, types.ConfigNftContractAddress)

	nftMint := entitycontracts.WasmMsgMint{
		Mint: entitycontracts.Mint{
			TokenId:   did.Id,
			Owner:     msg.OwnerAddress,
			TokenUri:  did.Id,
			Extension: msg.Data,
		},
	}

	finalMessage, err := nftMint.Marshal()

	if err != nil {
		return types.MsgCreateEntityResponse{}, err
	}

	_, err = k.WasmKeeper.Execute(ctx, nftContractAddress, address, finalMessage, sdk.NewCoins(sdk.NewCoin("uixo", sdk.ZeroInt())))
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

func (k Keeper) TransferEntity(ctx sdk.Context, msg *types.MsgTransferEntity) (*types.MsgTransferEntityResponse, error) {
	params := k.GetParams(ctx)
	nftContractAddressParam := params.NftContractAddress

	if len(nftContractAddressParam) == 0 {
		return nil, errors.New("nftContractAddress not set")
	}

	controllerAddress, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		return nil, err
	}

	recipientDidDoc, found := k.IidKeeper.GetDidDocument(ctx, []byte(msg.RecipientDid))
	if !found {
		return nil, errors.New("recipient did not found")
	}

	recipientAddress, err := recipientDidDoc.GetVerificationMethodBlockchainAddress(recipientDidDoc.Id)
	if err != nil {
		return nil, err
	}

	nftContractAddress, err := sdk.AccAddressFromBech32(nftContractAddressParam)
	if err != nil {
		return nil, err
	}

	err = iidkeeper.ExecuteOnDidWithRelationships(
		ctx,
		&k.IidKeeper,
		[]string{iidtypes.Authentication},
		msg.EntityDid,
		msg.OwnerDid,
		func(document *iidtypes.IidDocument) error {
			document.Controller = []string{
				document.Id,
				msg.RecipientDid,
			}
			return nil
		},
	)

	if err != nil {
		return nil, err
	}

	nftTranferMsg := entitycontracts.WasmMsgTransferNft{
		TransferNft: entitycontracts.TransferNft{
			TokenId:   msg.EntityDid,
			Recipient: recipientAddress.String(),
		},
	}

	finalMessage, err := nftTranferMsg.Marshal()

	if err != nil {
		return nil, err
	}

	_, err = k.WasmKeeper.Execute(ctx, nftContractAddress, controllerAddress, finalMessage, sdk.NewCoins(sdk.NewCoin("uixo", sdk.ZeroInt())))
	if err != nil {
		return nil, err
	}

	return &types.MsgTransferEntityResponse{}, err
}

// GetParams returns the total set of project parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.ParamSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the total set of project parameters.
func (k Keeper) SetParams(ctx sdk.Context, params *types.Params) {
	k.ParamSpace.SetParamSet(ctx, params)
}

func (k Keeper) GetEntityDocIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.EntityKey)
}

func (k Keeper) EntityExists(ctx sdk.Context, entityDid string) bool {
	// store := ctx.KVStore(k.storeKey)
	_, exists := k.IidKeeper.GetDidDocument(ctx, []byte(entityDid))
	return exists
}

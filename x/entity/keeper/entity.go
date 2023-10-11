package keeper

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ixofoundation/ixo-blockchain/v2/x/entity/types"
	nft "github.com/ixofoundation/ixo-blockchain/v2/x/entity/types/contracts"
	iidTypes "github.com/ixofoundation/ixo-blockchain/v2/x/iid/types"
)

func (k Keeper) SetEntity(ctx sdk.Context, key []byte, meta types.Entity) {
	k.Set(ctx, key, types.EntityKey, meta, k.Marshal)
}

func (k Keeper) GetEntity(ctx sdk.Context, key []byte) (types.Entity, bool) {
	val, found := k.Get(ctx, key, types.EntityKey, k.UnmarshalEntity)
	return val.(types.Entity), found
}

func (k Keeper) UnmarshalEntity(value []byte) (interface{}, bool) {
	data := types.Entity{}
	k.Unmarshal(value, &data)
	return data, types.IsValidEntity(&data)
}

// ResolveEntity returns the Entity and IidDocument
func (k Keeper) ResolveEntity(ctx sdk.Context, entityId string) (iidDocument iidTypes.IidDocument, entity types.Entity, err error) {
	iidDocument, err = k.IidKeeper.ResolveDid(ctx, iidTypes.DID(entityId))
	if err != nil {
		return
	}

	entity, found := k.GetEntity(ctx, []byte(entityId))
	if !found {
		err = types.ErrEntityNotFound
	}
	return
}

func (k Keeper) Marshal(value interface{}) (bytes []byte) {
	switch value := value.(type) {
	case types.Entity:
		bytes = k.cdc.MustMarshal(&value)
	}
	return
}

// Unmarshal unmarshal a byte slice to a struct, return false in case of errors
func (k Keeper) Unmarshal(data []byte, val codec.ProtoMarshaler) bool {
	if len(data) == 0 {
		return false
	}
	if err := k.cdc.Unmarshal(data, val); err != nil {
		return false
	}
	return true
}

// GetAllEntityWithCondition retrieve a list of
// entitiy docs by some arbitrary criteria.
func (k Keeper) GetAllEntityWithCondition(
	ctx sdk.Context,
	key []byte,
	entitySelector func(entity types.Entity) bool,
) (entities []types.Entity) {
	iterator := k.GetAll(ctx, key)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		entity, _ := k.UnmarshalEntity(iterator.Value())
		entityTyped := entity.(types.Entity)

		if entitySelector(entityTyped) {
			entities = append(entities, entityTyped)
		}
	}

	return entities
}

// GetAllEntity returns all the Entity docs
func (k Keeper) GetAllEntity(ctx sdk.Context) []types.Entity {
	return k.GetAllEntityWithCondition(
		ctx,
		types.EntityKey,
		func(did types.Entity) bool { return true },
	)
}

// Create a module account for entity id and name of account as fragemnt in form: did#name
func (k Keeper) CreateNewAccount(ctx sdk.Context, entityId, name string) (sdk.AccAddress, error) {
	address := types.GetModuleAccountAddress(entityId, name)

	if k.AccountKeeper.GetAccount(ctx, address) != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "account already exists")
	}

	account := k.AccountKeeper.NewAccountWithAddress(ctx, address)
	k.AccountKeeper.SetAccount(ctx, account)

	return account.GetAddress(), nil
}

// checks if the provided address is the owner on the smart contract
func (k Keeper) CheckIfOwner(ctx sdk.Context, entityId, ownerAddress string) error {
	// get cw721 contract address
	params := k.GetParams(ctx)
	nftContractAddress, err := sdk.AccAddressFromBech32(params.NftContractAddress)
	if err != nil {
		return err
	}

	// create the nft cw721 query
	queryMessage, err := nft.Marshal(nft.WasmQueryOwnerOf{
		OwnerOf: nft.OwnerOf{
			TokenId: entityId,
		},
	})
	if err != nil {
		return err
	}

	// query smart contract
	ownerOfBytes, err := k.WasmViewKeeper.QuerySmart(ctx, nftContractAddress, queryMessage)
	if err != nil {
		return err
	}
	var ownerOf nft.OwnerOfResponse
	if err := json.Unmarshal([]byte(ownerOfBytes), &ownerOf); err != nil {
		return err
	}

	// check if token owner is owner provided
	if ownerOf.Owner == ownerAddress {
		return nil
	}

	return types.ErrEntityUnauthorized
}

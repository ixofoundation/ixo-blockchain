package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-blockchain/v3/x/iid/types"
)

func (k Keeper) SetDidDocument(ctx sdk.Context, key []byte, document types.IidDocument) {
	k.Set(ctx, key, types.DidDocumentKey, document, k.Marshal)
}

func (k Keeper) GetDidDocument(ctx sdk.Context, key []byte) (types.IidDocument, bool) {
	val, found := k.Get(ctx, key, types.DidDocumentKey, k.UnmarshalDidDocument)
	iid, ok := val.(types.IidDocument)
	if !ok {
		return types.IidDocument{}, false
	}
	return iid, found
}

func (k Keeper) UnmarshalDidDocument(value []byte) (interface{}, bool) {
	data := types.IidDocument{}
	k.Unmarshal(value, &data)
	return data, types.IsValidDIDDocument(&data)
}

// ResolveDid returning the did document and throwing error if not found
func (k Keeper) ResolveDid(ctx sdk.Context, did types.DID) (doc types.IidDocument, err error) {
	doc, found := k.GetDidDocument(ctx, []byte(did.String()))
	if !found {
		err = types.ErrDidDocumentNotFound
		return
	}
	return
}

func (k Keeper) Marshal(value interface{}) (bytes []byte) {
	switch value := value.(type) {
	case types.IidDocument:
		bytes = k.cdc.MustMarshal(&value)
	}
	return
}

// nolint
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

// GetAllDidDocumentsWithCondition retrieve a list of did document by some arbitrary criteria.
func (k Keeper) GetAllDidDocumentsWithCondition(
	ctx sdk.Context,
	key []byte,
	didSelector func(did types.IidDocument) bool,
) (didDocs []types.IidDocument) {
	iterator := k.GetAll(ctx, key)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		did, _ := k.UnmarshalDidDocument(iterator.Value())
		didTyped, ok := did.(types.IidDocument)
		if !ok {
			continue
		}

		if didSelector(didTyped) {
			didDocs = append(didDocs, didTyped)
		}
	}

	return didDocs
}

// GetAllDidDocuments returns all the DidDocuments
func (k Keeper) GetAllDidDocuments(ctx sdk.Context) []types.IidDocument {
	return k.GetAllDidDocumentsWithCondition(
		ctx,
		types.DidDocumentKey,
		func(did types.IidDocument) bool { return true },
	)
}

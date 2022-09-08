package keeper

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-blockchain/x/iid/types"
)

func (k Keeper) SetDidDocument(ctx sdk.Context, key []byte, document types.IidDocument) {
	k.Set(ctx, key, types.DidDocumentKey, document, k.Marshal)
}

func (k Keeper) GetDidDocument(ctx sdk.Context, key []byte) (types.IidDocument, bool) {
	val, found := k.Get(ctx, key, types.DidDocumentKey, k.UnmarshalDidDocument)
	return val.(types.IidDocument), found
}

// UnmarshalDidDocument unmarshall a did document= and check if it is empty
// ad DID document is empty if contains no context
func (k Keeper) UnmarshalDidDocument(value []byte) (interface{}, bool) {
	data := types.IidDocument{}
	k.Unmarshal(value, &data)
	return data, types.IsValidDIDDocument(&data)
}

func (k Keeper) SetDidMetadata(ctx sdk.Context, key []byte, meta types.IidMetadata) {
	k.Set(ctx, key, types.DidMetadataKey, meta, k.Marshal)
}

func (k Keeper) GetDidMetadata(ctx sdk.Context, key []byte) (types.IidMetadata, bool) {
	val, found := k.Get(ctx, key, types.DidMetadataKey, k.UnmarshalDidMetadata)
	return val.(types.IidMetadata), found
}

func (k Keeper) UnmarshalDidMetadata(value []byte) (interface{}, bool) {
	data := types.IidMetadata{}
	k.Unmarshal(value, &data)
	return data, types.IsValidDIDMetadata(&data)
}

// ResolveDid returning the did document and associated metadata
func (k Keeper) ResolveDid(ctx sdk.Context, did types.DID) (doc types.IidDocument, meta types.IidMetadata, err error) {
	if strings.HasPrefix(did.String(), types.DidKeyPrefix) {
		doc, meta, err = types.ResolveAccountDID(did.String(), ctx.ChainID())
		return
	}
	doc, found := k.GetDidDocument(ctx, []byte(did.String()))
	if !found {
		err = types.ErrDidDocumentNotFound
		return
	}
	meta, _ = k.GetDidMetadata(ctx, []byte(did.String()))
	return
}

func (k Keeper) Marshal(value interface{}) (bytes []byte) {
	switch value := value.(type) {
	case types.IidDocument:
		bytes = k.cdc.MustMarshal(&value)
	case types.IidMetadata:
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

// GetAllDidDocumentsWithCondition retrieve a list of
// did document by some arbitrary criteria. The selector filter has access
// to both the did and its metadata
func (k Keeper) GetAllDidDocumentsWithCondition(
	ctx sdk.Context,
	key []byte,
	didSelector func(did types.IidDocument) bool,
) (didDocs []types.IidDocument) {
	iterator := k.GetAll(ctx, key)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		did, _ := k.UnmarshalDidDocument(iterator.Value())
		didTyped := did.(types.IidDocument)

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

// GetDidDocumentsByPubKey retrieve a did document using a pubkey associated to the DID
// TODO: this function is used only in the issuer module ante handler !
func (k Keeper) GetDidDocumentsByPubKey(ctx sdk.Context, pubkey cryptotypes.PubKey) (dids []types.IidDocument) {

	dids = k.GetAllDidDocumentsWithCondition(
		ctx,
		types.DidDocumentKey,
		func(did types.IidDocument) bool {
			return did.HasPublicKey(pubkey)
		},
	)
	// compute the key did

	// generate the address
	addr, err := sdk.Bech32ifyAddressBytes(sdk.GetConfig().GetBech32AccountAddrPrefix(), pubkey.Address())
	if err != nil {
		return
	}
	doc, _, err := types.ResolveAccountDID(types.NewKeyDID(addr).String(), ctx.ChainID())
	if err != nil {
		return
	}
	dids = append(dids, doc)
	return
}

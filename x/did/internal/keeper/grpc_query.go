package keeper

import (
	"context"
	"github.com/ixofoundation/ixo-blockchain/x/did/internal/types"
)

func (k Keeper) DidDoc(ctx context.Context, request *types.QueryDidDocRequest) (*types.QueryDidDocResponse, error) {
	panic("implement me")
}

func (k Keeper) AllDids(ctx context.Context, request *types.QueryAllDidsRequest) (*types.QueryAllDidsResponse, error) {
	panic("implement me")
}

func (k Keeper) AllDidDocs(ctx context.Context, request *types.QueryAllDidDocsRequest) (*types.QueryAllDidDocsResponse, error) {
	panic("implement me")
}

func (k Keeper) AddressFromDid(ctx context.Context, request *types.QueryAddressFromDidRequest) (*types.QueryAddressFromDidResponse, error) {
	panic("implement me")
}

func (k Keeper) AddressFromBase58EncodedPubkey(ctx context.Context, request *types.QueryAddressFromBase58EncodedPubkeyRequest) (*types.QueryAddressFromBase58EncodedPubkeyResponse, error) {
	panic("implement me")
}

func (k Keeper) IxoDidFromMnemonic(ctx context.Context, request *types.QueryIxoDidFromMnemonicRequest) (*types.QueryIxoDidFromMnemonicResponse, error) {
	panic("implement me")
}
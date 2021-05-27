package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"github.com/ixofoundation/ixo-blockchain/x/payments/types"
	"strings"
)

type Keeper struct {
	cdc                codec.BinaryMarshaler
	storeKey           sdk.StoreKey
	paramSpace         paramtypes.Subspace
	bankKeeper         bankkeeper.Keeper
	DidKeeper          did.Keeper
	reservedIdPrefixes []string
}

func NewKeeper(cdc codec.BinaryMarshaler, storeKey sdk.StoreKey, bankKeeper bankkeeper.Keeper,
	didKeeper did.Keeper, reservedIdPrefixes []string) Keeper {
	return Keeper{
		cdc:                cdc,
		storeKey:           storeKey,
		bankKeeper:         bankKeeper,
		DidKeeper:          didKeeper,
		reservedIdPrefixes: reservedIdPrefixes,
	}
}

// idReserved checks if the id (from a template, contract, or subscription)
// is using a reserved prefix (i.e. should not be used). The prefixPrefix
// indicates whether we are referring to a template, contract, or subscription.
func (k Keeper) idReserved(id string, prefixPrefix string) bool {
	for _, prefix := range k.reservedIdPrefixes {
		if strings.HasPrefix(id, prefixPrefix+prefix) {
			return true
		}
	}
	return false
}

func (k Keeper) PaymentTemplateIdReserved(id string) bool {
	return k.idReserved(id, types.PaymentTemplateIdPrefix)
}

func (k Keeper) PaymentContractIdReserved(id string) bool {
	return k.idReserved(id, types.PaymentContractIdPrefix)
}

func (k Keeper) SubscriptionIdReserved(id string) bool {
	return k.idReserved(id, types.SubscriptionIdPrefix)
}

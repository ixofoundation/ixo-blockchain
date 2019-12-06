package keeper

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-cosmos/codec"
	"github.com/ixofoundation/ixo-cosmos/types"
	fiatTypes "github.com/ixofoundation/ixo-cosmos/x/fiat/internal/types"
)

type Keeper struct {
	storeKey sdk.StoreKey
	cdc      *codec.Codec
}

func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
	}
}

// SetFiatAccount:
func (k Keeper) SetFiatAccount(ctx sdk.Context, fiatAccount types.FiatAccount) {
	store := ctx.KVStore(k.storeKey)

	fiatAccountKey := fiatTypes.FiatAccountStoreKey(fiatAccount.GetAddress())
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(fiatAccount)
	store.Set(fiatAccountKey, bz)
}

// returns fiat account by address
func (k Keeper) GetFiatAccount(ctx sdk.Context, address sdk.AccAddress) (fiatAccount types.FiatAccount, err sdk.Error) {
	store := ctx.KVStore(k.storeKey)

	fiatAccountKey := fiatTypes.FiatAccountStoreKey(address)
	bz := store.Get(fiatAccountKey)
	if bz == nil {
		return nil, fiatTypes.ErrInvalidPegHash(fiatTypes.DefaultCodeSpace)
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &fiatAccount)
	return fiatAccount, nil
}

func (k Keeper) IterateFiatAccounts(ctx sdk.Context, handler func(fiatAccount types.FiatAccount) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, fiatTypes.FiatAccountKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var fiatAccount types.FiatAccount
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iterator.Value(), &fiatAccount)
		if handler(fiatAccount) {
			break
		}
	}
}

// get all fiat accounts => []FiatAccounts from store
func (k Keeper) GetFiatAccounts(ctx sdk.Context) (fiatAccounts []types.FiatAccount) {
	k.IterateFiatAccounts(ctx, func(fiatAccount types.FiatAccount) (stop bool) {
		fiatAccounts = append(fiatAccounts, fiatAccount)
		return false
	},
	)
	return
}

// GetNextFiatPegHash : Returns and increments the fiatPeg counter
func (k Keeper) getNextFiatPegHash(ctx sdk.Context) int {
	var fiatNumber int
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(fiatTypes.FiatPegHashKey)
	if bz == nil {
		fiatNumber = 0
	} else {
		k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &fiatNumber)
	}

	bz = k.cdc.MustMarshalBinaryLengthPrefixed(fiatNumber + 1)
	store.Set(fiatTypes.FiatPegHashKey, bz)
	//TODO
	return 11
}

func (k Keeper) IssueFiats(ctx sdk.Context, issueFiat fiatTypes.IssueFiat) sdk.Error {
	fiatAccount, err := k.GetFiatAccount(ctx, issueFiat.ToAddress)
	if err != nil {
		fiatAccount = types.NewFiatAccountWithAddress(issueFiat.ToAddress)
	}
	pegHash, err2 := types.GetFiatPegHashHex(strconv.Itoa(k.getNextFiatPegHash(ctx)))
	if err2 != nil {
		return sdk.ErrInternal(fmt.Sprintf("%s", err.Error()))
	}
	fiatPeg := types.NewFiatPegWithPegHash(pegHash)
	fiatPeg.SetTransactionAmount(issueFiat.TransactionAmount)
	fiatPeg.SetTransactionID(issueFiat.TransactionID)
	oldFiatPegWallet := fiatAccount.GetFiatPegWallet()
	newFiatPegWallet := types.AddFiatPegToWallet(oldFiatPegWallet, []types.BaseFiatPeg{types.ToBaseFiatPeg(fiatPeg)})
	fiatAccount.SetFiatPegWallet(newFiatPegWallet)
	k.SetFiatAccount(ctx, fiatAccount)
	return nil
}

func (k Keeper) RedeemFiats(ctx sdk.Context, redeemFiat fiatTypes.RedeemFiat) sdk.Error {
	fiatAccount, err := k.GetFiatAccount(ctx, redeemFiat.RedeemerAddress)
	if err != nil {
		return err
	}
	emptiedFiatPegWallet, redeemerFiatPegWallet := types.RedeemAmountFromWallet(redeemFiat.Amount, fiatAccount.GetFiatPegWallet())
	if len(redeemerFiatPegWallet) == 0 && len(emptiedFiatPegWallet) == 0 {
		return sdk.ErrInsufficientCoins(fmt.Sprintf("Redeemed amount higher than the account balance"))
	}
	fiatAccount.SetFiatPegWallet(redeemerFiatPegWallet)
	k.SetFiatAccount(ctx, fiatAccount)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		fiatTypes.EventTypeRedeemFiat,
		sdk.NewAttribute("redeemer", redeemFiat.RedeemerAddress.String()),
	))
	return nil
}

func (k Keeper) SendFiats(ctx sdk.Context, sendFiat fiatTypes.SendFiat) sdk.Error {
	fromFiatAccount, err := k.GetFiatAccount(ctx, sendFiat.FromAddress)
	if err != nil {
		return err
	}
	sentFiatPegWallet, oldFiatPegWallet := types.SubtractAmountFromWallet(sendFiat.Amount, fromFiatAccount.GetFiatPegWallet())
	if len(sentFiatPegWallet) == 0 && len(oldFiatPegWallet) == 0 {
		return sdk.ErrInsufficientCoins(fmt.Sprintf("Insufficient funds"))
	}

	toFiatAccount, err := k.GetFiatAccount(ctx, sendFiat.ToAddress)
	if err != nil {
		toFiatAccount = types.NewFiatAccountWithAddress(sendFiat.ToAddress)
	}
	toFiatAccount.SetFiatPegWallet(sentFiatPegWallet)
	k.SetFiatAccount(ctx, toFiatAccount)

	fromFiatAccount.SetFiatPegWallet(oldFiatPegWallet)
	k.SetFiatAccount(ctx, fromFiatAccount)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		fiatTypes.EventTypeSendFiat,
		sdk.NewAttribute("recipient", sendFiat.ToAddress.String()),
		sdk.NewAttribute("sender", sendFiat.FromAddress.String()),
	))

	return nil
}

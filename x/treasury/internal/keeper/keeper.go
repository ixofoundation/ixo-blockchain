package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
	"github.com/ixofoundation/ixo-cosmos/x/oracles"
	"github.com/ixofoundation/ixo-cosmos/x/treasury/internal/types"
)

type Keeper struct {
	cdc           *codec.Codec
	storeKey      sdk.StoreKey
	bankKeeper    bank.Keeper
	oraclesKeeper oracles.Keeper
	supplyKeeper  supply.Keeper
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, bankKeeper bank.Keeper,
	oraclesKeeper oracles.Keeper, supplyKeeper supply.Keeper) Keeper {

	return Keeper{
		cdc:           cdc,
		storeKey:      key,
		bankKeeper:    bankKeeper,
		oraclesKeeper: oraclesKeeper,
		supplyKeeper:  supplyKeeper,
	}
}

func (k Keeper) Send(ctx sdk.Context, fromDid, toDid ixo.Did, amount sdk.Coins) sdk.Error {
	fromAddress := types.DidToAddr(fromDid)
	toAddress := types.DidToAddr(toDid)

	err := k.bankKeeper.SendCoins(ctx, fromAddress, toAddress, amount)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) Mint(ctx sdk.Context, oracleDid, toDid ixo.Did, amount sdk.Coins) sdk.Error {
	toAddress := types.DidToAddr(toDid)

	// Check if oracle exists
	if !k.oraclesKeeper.OracleExists(ctx, oracles.Oracle(oracleDid)) {
		return sdk.ErrInternal("oracle specified is not a registered oracle")
	}

	// Mint coins to module account
	err := k.supplyKeeper.MintCoins(ctx, types.ModuleName, amount)
	if err != nil {
		return err
	}

	// Send minted tokens to recipient
	err = k.supplyKeeper.SendCoinsFromModuleToAccount(ctx,
		types.ModuleName, toAddress, amount)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) Burn(ctx sdk.Context, oracleDid, fromDid ixo.Did, amount sdk.Coins) sdk.Error {
	fromAddress := types.DidToAddr(fromDid)

	// Check if oracle exists
	if !k.oraclesKeeper.OracleExists(ctx, oracles.Oracle(oracleDid)) {
		return sdk.ErrInternal("oracle specified is not a registered oracle")
	}

	// Take tokens to burn from account
	err := k.supplyKeeper.SendCoinsFromAccountToModule(ctx,
		fromAddress, types.ModuleName, amount)
	if err != nil {
		return err
	}

	// Burn coins from module account
	err = k.supplyKeeper.BurnCoins(ctx, types.ModuleName, amount)
	if err != nil {
		return err
	}

	return nil
}

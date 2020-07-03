package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"github.com/ixofoundation/ixo-blockchain/x/oracles"
	"github.com/ixofoundation/ixo-blockchain/x/treasury/internal/types"
)

type Keeper struct {
	cdc           *codec.Codec
	storeKey      sdk.StoreKey
	bankKeeper    bank.Keeper
	oraclesKeeper oracles.Keeper
	supplyKeeper  supply.Keeper
	didKeeper     did.Keeper
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, bankKeeper bank.Keeper,
	oraclesKeeper oracles.Keeper, supplyKeeper supply.Keeper,
	didKeeper did.Keeper) Keeper {

	return Keeper{
		cdc:           cdc,
		storeKey:      key,
		bankKeeper:    bankKeeper,
		oraclesKeeper: oraclesKeeper,
		supplyKeeper:  supplyKeeper,
		didKeeper:     didKeeper,
	}
}

func (k Keeper) Send(ctx sdk.Context, fromDid, toDid did.Did, amount sdk.Coins) sdk.Error {
	// Get from address
	fromDidDoc, err := k.didKeeper.GetDidDoc(ctx, fromDid)
	if err != nil {
		return err
	}
	fromAddress := fromDidDoc.Address()

	// Get to address
	toDidDoc, err := k.didKeeper.GetDidDoc(ctx, toDid)
	if err != nil {
		return err
	}
	toAddress := toDidDoc.Address()

	err = k.bankKeeper.SendCoins(ctx, fromAddress, toAddress, amount)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) OracleTransfer(ctx sdk.Context, fromDid, toDid, oracleDid did.Did, amount sdk.Coins) sdk.Error {

	// Check if oracle exists
	if !k.oraclesKeeper.OracleExists(ctx, oracleDid) {
		return sdk.ErrInternal("oracle specified is not a registered oracle")
	}

	// Confirm that oracle has the required capabilities
	oracle := k.oraclesKeeper.MustGetOracle(ctx, oracleDid)
	for _, c := range amount {
		if !oracle.Capabilities.Includes(c.Denom) {
			return sdk.ErrInternal(fmt.Sprintf(
				"oracle does not have capability to send %s", c.Denom))
		}

		// Get capability by token name
		capability := oracle.Capabilities.MustGet(c.Denom)
		if !capability.Capabilities.Includes(oracles.TransferCap) {
			return sdk.ErrInternal(fmt.Sprintf(
				"oracle does not have capability to send %s", c.Denom))
		}
	}

	// Perform send
	return k.Send(ctx, fromDid, toDid, amount)
}

func (k Keeper) OracleMint(ctx sdk.Context, oracleDid, toDid did.Did, amount sdk.Coins) sdk.Error {
	// Get to address
	toDidDoc, err := k.didKeeper.GetDidDoc(ctx, toDid)
	if err != nil {
		return err
	}
	toAddress := toDidDoc.Address()

	// Check if oracle exists
	if !k.oraclesKeeper.OracleExists(ctx, oracleDid) {
		return sdk.ErrInternal("oracle specified is not a registered oracle")
	}

	// Confirm that oracle has the required capabilities
	oracle := k.oraclesKeeper.MustGetOracle(ctx, oracleDid)
	for _, c := range amount {
		if !oracle.Capabilities.Includes(c.Denom) {
			return sdk.ErrInternal(fmt.Sprintf(
				"oracle does not have capability to mint %s", c.Denom))
		}

		// Get capability by token name
		capability := oracle.Capabilities.MustGet(c.Denom)
		if !capability.Capabilities.Includes(oracles.MintCap) {
			return sdk.ErrInternal(fmt.Sprintf(
				"oracle does not have capability to mint %s", c.Denom))
		}
	}

	// Mint coins to module account
	err = k.supplyKeeper.MintCoins(ctx, types.ModuleName, amount)
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

func (k Keeper) OracleBurn(ctx sdk.Context, oracleDid, fromDid did.Did, amount sdk.Coins) sdk.Error {
	// Get from address
	fromDidDoc, err := k.didKeeper.GetDidDoc(ctx, fromDid)
	if err != nil {
		return err
	}
	fromAddress := fromDidDoc.Address()

	// Check if oracle exists
	if !k.oraclesKeeper.OracleExists(ctx, oracleDid) {
		return sdk.ErrInternal("oracle specified is not a registered oracle")
	}

	// Confirm that oracle has the required capabilities
	oracle := k.oraclesKeeper.MustGetOracle(ctx, oracleDid)
	for _, c := range amount {
		if !oracle.Capabilities.Includes(c.Denom) {
			return sdk.ErrInternal(fmt.Sprintf(
				"oracle does not have capability to burn %s", c.Denom))
		}

		// Get capability by token name
		capability := oracle.Capabilities.MustGet(c.Denom)
		if !capability.Capabilities.Includes(oracles.BurnCap) {
			return sdk.ErrInternal(fmt.Sprintf(
				"oracle does not have capability to burn %s", c.Denom))
		}
	}

	// Take tokens to burn from account
	err = k.supplyKeeper.SendCoinsFromAccountToModule(ctx,
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

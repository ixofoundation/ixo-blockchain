package simulation

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/ixofoundation/ixo-blockchain/x/nft/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/nft/types"
)

const (
	OpWeightedMsgTransferNFT     = " op_weighted_msg_transfer_nft"
	OpWeightedMsgEditNFTMetadata = "op_weighted_msg_edit_nft_metadata"
	OpWeightedMsgMintNFT         = "op_weighted_msg_mint_nft"
	OpWeightedMsgBurnNFT         = "op_weighted_msg_burn_nft"
)

func WeightedOperations(appParams simulation.AppParams, cdc *codec.Codec, ak types.AccountKeeper, k keeper.Keeper) simulation.WeightedOperations {

	var (
		weightedMsgTransferNFT     int
		weightedMsgEditNFTMetadata int
		weightedMsgMintNFT         int
		weightedMsgBurnNFT         int
	)

	appParams.GetOrGenerate(cdc, OpWeightedMsgTransferNFT, &weightedMsgTransferNFT, nil,
		func(_ *rand.Rand) {
			weightedMsgTransferNFT = 100
		},
	)
	appParams.GetOrGenerate(cdc, OpWeightedMsgEditNFTMetadata, &weightedMsgEditNFTMetadata, nil,
		func(_ *rand.Rand) {
			weightedMsgEditNFTMetadata = 100
		},
	)
	appParams.GetOrGenerate(cdc, OpWeightedMsgMintNFT, &weightedMsgMintNFT, nil,
		func(_ *rand.Rand) {
			weightedMsgMintNFT = 5
		},
	)
	appParams.GetOrGenerate(cdc, OpWeightedMsgBurnNFT, &weightedMsgBurnNFT, nil,
		func(_ *rand.Rand) {
			weightedMsgBurnNFT = 5
		},
	)

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightedMsgTransferNFT,
			SimulateMsgTransferNFT(ak, k),
		),
		// simulation.NewWeightedOperation(
		// 	weightedMsgEditNFTMetadata,
		// 	SimulateMsgEditNFTMetadata(ak, k),
		// ),
		// simulation.NewWeightedOperation(
		// 	weightedMsgMintNFT,
		// 	SimulateMsgMintNFT(ak, k),
		// ),
		// simulation.NewWeightedOperation(
		// 	weightedMsgBurnNFT,
		// 	SimulateMsgBurnNFT(ak, k),
		// ),
	}
}

// SimulateMsgTransferNFT simulates the transfer of an NFT
func SimulateMsgTransferNFT(ak types.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string,
	) (simulation.OperationMsg, []simulation.FutureOperation, error) {

		simAccount, _ := simulation.RandomAcc(r, accs)
		ownerAddr, denom, nftID := getRandomNFTFromOwner(ctx, k, r)
		if ownerAddr.Empty() {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}
		msg := types.NewMsgTransferNFT(
			ownerAddr,          // sender
			simAccount.Address, // recipient
			denom,
			nftID,
		)

		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		acc := ak.GetAccount(ctx, simAccount.Address)
		coins := acc.SpendableCoins(ctx.BlockTime())

		var (
			fees sdk.Coins
			err  error
		)
		coins, hasNeg := coins.SafeSub(coins)
		if !hasNeg {
			fees, err = simulation.RandomFees(r, ctx, coins)
			if err != nil {
				return simulation.NoOpMsg(types.ModuleName), nil, err
			}
		}

		ownerAcc, ok := simulation.FindAccount(accs, ownerAddr)
		ownerAuthAcc := ak.GetAccount(ctx, ownerAddr)
		if !ok {
			return simulation.NoOpMsg(types.ModuleName), nil, errors.New("could not find acc")
		}

		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{ownerAuthAcc.GetAccountNumber()},
			[]uint64{ownerAuthAcc.GetSequence()},
			ownerAcc.PrivKey,
		)

		_, _, err = app.Deliver(tx)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// SimulateMsgEditNFTMetadata simulates an edit metadata transaction
func SimulateMsgEditNFTMetadata(ak types.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string,
	) (simulation.OperationMsg, []simulation.FutureOperation, error) {

		ownerAddr, denom, nftID := getRandomNFTFromOwner(ctx, k, r)
		if ownerAddr.Empty() {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		acc := ak.GetAccount(ctx, ownerAddr)
		coins := acc.SpendableCoins(ctx.BlockTime())

		var (
			fees sdk.Coins
			err  error
		)
		coins, hasNeg := coins.SafeSub(coins)
		if !hasNeg {
			fees, err = simulation.RandomFees(r, ctx, coins)
			if err != nil {
				return simulation.NoOpMsg(types.ModuleName), nil, err
			}
		}

		msg := types.NewMsgEditNFTMetadata(
			ownerAddr,
			nftID,
			denom,
			simulation.RandStringOfLength(r, 45), // tokenURI
		)

		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		// find the acc in simulated acocunts
		simAcc, ok := simulation.FindAccount(accs, acc.GetAddress())
		if !ok {
			return simulation.NoOpMsg(types.ModuleName), nil, errors.New("account not found")
		}
		fmt.Println("here")

		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{acc.GetAccountNumber()},
			[]uint64{acc.GetSequence()},
			simAcc.PrivKey,
		)

		_, _, err = app.Deliver(tx)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// SimulateMsgMintNFT simulates a mint of an NFT
func SimulateMsgMintNFT(ak types.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string,
	) (simulation.OperationMsg, []simulation.FutureOperation, error) {
		a1, _ := simulation.RandomAcc(r, accs)
		a2, _ := simulation.RandomAcc(r, accs)

		msg := types.NewMsgMintNFT(
			a1.Address,                           // sender
			a2.Address,                           // recipient
			simulation.RandStringOfLength(r, 10), // nft ID
			simulation.RandStringOfLength(r, 10), // denom
			simulation.RandStringOfLength(r, 45), // tokenURI
		)

		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}
		acc := ak.GetAccount(ctx, a1.Address)
		coins := acc.SpendableCoins(ctx.BlockTime())

		var (
			fees sdk.Coins
			err  error
		)
		coins, hasNeg := coins.SafeSub(coins)
		if !hasNeg {
			fees, err = simulation.RandomFees(r, ctx, coins)
			if err != nil {
				return simulation.NoOpMsg(types.ModuleName), nil, err
			}
		}

		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{acc.GetAccountNumber()},
			[]uint64{acc.GetSequence()},
			a1.PrivKey,
		)

		_, _, err = app.Deliver(tx)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// SimulateMsgBurnNFT simulates a burn of an existing NFT
func SimulateMsgBurnNFT(ak types.AccountKeeper, k keeper.Keeper) simulation.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simulation.Account, chainID string,
	) (simulation.OperationMsg, []simulation.FutureOperation, error) {

		ownerAddr, denom, nftID := getRandomNFTFromOwner(ctx, k, r)
		if ownerAddr.Empty() {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		msg := types.NewMsgBurnNFT(ownerAddr, nftID, denom)

		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		acc := ak.GetAccount(ctx, ownerAddr)
		coins := acc.SpendableCoins(ctx.BlockTime())

		var (
			fees sdk.Coins
			err  error
		)
		coins, hasNeg := coins.SafeSub(coins)
		if !hasNeg {
			fees, err = simulation.RandomFees(r, ctx, coins)
			if err != nil {
				return simulation.NoOpMsg(types.ModuleName), nil, err
			}
		}

		// find the acc in simulated acocunts
		simAcc, ok := simulation.FindAccount(accs, ownerAddr)
		if !ok {
			return simulation.NoOpMsg(types.ModuleName), nil, errors.New("account not found, burnMsg")
		}

		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{acc.GetAccountNumber()},
			[]uint64{acc.GetSequence()},
			simAcc.PrivKey,
		)

		_, _, err = app.Deliver(tx)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}

func getRandomNFTFromOwner(ctx sdk.Context, k keeper.Keeper, r *rand.Rand) (address sdk.AccAddress, denom, nftID string) {
	owners := k.GetOwners(ctx)

	ownersLen := len(owners)
	if ownersLen == 0 {
		return nil, "", ""
	}

	// get random owner
	i := r.Intn(ownersLen)
	owner := owners[i]

	idCollectionsLen := len(owner.IDCollections)
	if idCollectionsLen == 0 {
		return nil, "", ""
	}

	// get random collection from owner's balance
	i = r.Intn(idCollectionsLen)
	idsCollection := owner.IDCollections[i] // nfts IDs
	denom = idsCollection.Denom

	idsLen := len(idsCollection.IDs)
	if idsLen == 0 {
		return nil, "", ""
	}

	// get random nft from collection
	i = r.Intn(idsLen)
	nftID = idsCollection.IDs[i]

	return owner.Address, denom, nftID
}

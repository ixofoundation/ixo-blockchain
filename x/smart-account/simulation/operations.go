package simulation

import (
	"context"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/ixofoundation/ixo-blockchain/v7/x/smart-account/authenticator"
	"github.com/ixofoundation/ixo-blockchain/v7/x/smart-account/types"
)

const (
	OpWeightMsgAddAuthenticator = "op_weight_msg_add_authenticator"

	DefaultWeightMsgAddAuthenticator = 50
)

// AccountKeeper subset for x/smart-account sim.
type AccountKeeper interface {
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI
}

// BankKeeper subset for x/smart-account sim.
type BankKeeper interface {
	SpendableCoins(context.Context, sdk.AccAddress) sdk.Coins
}

// WeightedOperations exports the per-Msg simulation operations for
// x/smart-account.
func WeightedOperations(
	appParams simtypes.AppParams,
	cdc interface{},
	txGen client.TxConfig,
	ak AccountKeeper,
	bk BankKeeper,
) simulation.WeightedOperations {
	var weight int
	appParams.GetOrGenerate(OpWeightMsgAddAuthenticator, &weight, nil,
		func(_ *rand.Rand) { weight = DefaultWeightMsgAddAuthenticator })

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(weight, SimulateMsgAddAuthenticator(txGen, ak, bk)),
	}
}

// SimulateMsgAddAuthenticator picks a random account and registers a fresh
// SignatureVerification authenticator under it. The authenticator's config
// is a brand-new secp256k1 pubkey so add-then-add doesn't collide.
func SimulateMsgAddAuthenticator(
	txGen client.TxConfig,
	ak AccountKeeper,
	bk BankKeeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msgType := sdk.MsgTypeURL(&types.MsgAddAuthenticator{})

		simAcc, _ := simtypes.RandomAcc(r, accs)
		account := ak.GetAccount(ctx, simAcc.Address)
		if account == nil {
			return simtypes.NoOpMsg(types.ModuleName, msgType, "account does not exist"), nil, nil
		}
		spendable := bk.SpendableCoins(ctx, simAcc.Address)
		if spendable.IsZero() {
			return simtypes.NoOpMsg(types.ModuleName, msgType, "no spendable coins"), nil, nil
		}

		pk := secp256k1.GenPrivKey().PubKey()
		msg := &types.MsgAddAuthenticator{
			Sender:            simAcc.Address.String(),
			AuthenticatorType: authenticator.SignatureVerificationType,
			Data:              pk.Bytes(),
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           txGen,
			Cdc:             nil,
			Msg:             msg,
			Context:         ctx,
			SimAccount:      simAcc,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
		}
		_, _, _ = simulation.GenAndDeliverTxWithRandFees(txCtx)
		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}

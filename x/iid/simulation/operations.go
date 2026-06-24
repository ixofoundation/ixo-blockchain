package simulation

import (
	"context"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/ixofoundation/ixo-blockchain/v8/x/iid/types"
)

const (
	OpWeightMsgCreateIidDocument = "op_weight_msg_create_iid_document"

	DefaultWeightMsgCreateIidDocument = 100
)

// AccountKeeper subset for x/iid sim.
type AccountKeeper interface {
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI
}

// BankKeeper subset for x/iid sim.
type BankKeeper interface {
	SpendableCoins(context.Context, sdk.AccAddress) sdk.Coins
}

// IidKeeper subset for x/iid sim — only used to skip when a DID for the
// random account already exists.
type IidKeeper interface {
	GetDidDocument(ctx sdk.Context, key []byte) (types.IidDocument, bool)
}

// WeightedOperations exports the per-Msg simulation operations for x/iid.
func WeightedOperations(
	appParams simtypes.AppParams,
	cdc interface{},
	txGen client.TxConfig,
	ak AccountKeeper,
	bk BankKeeper,
	k IidKeeper,
) simulation.WeightedOperations {
	var weight int
	appParams.GetOrGenerate(OpWeightMsgCreateIidDocument, &weight, nil,
		func(_ *rand.Rand) { weight = DefaultWeightMsgCreateIidDocument })

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(weight, SimulateMsgCreateIidDocument(txGen, ak, bk, k)),
	}
}

// SimulateMsgCreateIidDocument picks a random account and tries to register
// "did:ixo:<addr>" as its DID. If the DID is already registered the op
// returns a NoOpMsg so the simulator continues without aborting.
func SimulateMsgCreateIidDocument(
	txGen client.TxConfig,
	ak AccountKeeper,
	bk BankKeeper,
	k IidKeeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msgType := sdk.MsgTypeURL(&types.MsgCreateIidDocument{})

		simAcc, _ := simtypes.RandomAcc(r, accs)
		account := ak.GetAccount(ctx, simAcc.Address)
		if account == nil {
			return simtypes.NoOpMsg(types.ModuleName, msgType, "account does not exist"), nil, nil
		}
		spendable := bk.SpendableCoins(ctx, simAcc.Address)
		if spendable.IsZero() {
			return simtypes.NoOpMsg(types.ModuleName, msgType, "no spendable coins"), nil, nil
		}

		did := "did:ixo:" + simAcc.Address.String()
		if _, exists := k.GetDidDocument(ctx, []byte(did)); exists {
			return simtypes.NoOpMsg(types.ModuleName, msgType, "did already registered"), nil, nil
		}

		msg := &types.MsgCreateIidDocument{
			Id:     did,
			Signer: simAcc.Address.String(),
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

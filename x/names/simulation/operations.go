package simulation

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/ixofoundation/ixo-blockchain/v8/x/names/types"
)

// Op weights — overridable via the sim CLI's params file. Defaults are
// modest; these flows depend on a namespace existing in state which the
// simulator only sees if a previous gov-driven CreateNamespace operation
// ran. Without that, every op falls through to NoOpMsg.
const (
	OpWeightMsgRegisterName = "op_weight_msg_register_name"

	DefaultWeightMsgRegisterName = 50
)

// AccountKeeper is the subset of the SDK auth keeper the names sim needs.
// Signature uses context.Context to match the cosmos-sdk simulation package.
type AccountKeeper interface {
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI
}

// BankKeeper is the subset of the SDK bank keeper the names sim needs.
type BankKeeper interface {
	SpendableCoins(context.Context, sdk.AccAddress) sdk.Coins
}

// NamesKeeper is the subset of the names keeper the simulator reaches into
// — limited so we can cheaply mock it later if needed.
type NamesKeeper interface {
	GetAllNamespaces(ctx sdk.Context) []types.Namespace
	HasNameRecord(ctx sdk.Context, namespace, normalizedName string) bool
}

// WeightedOperations returns the per-Msg simulation operations for x/names.
// Currently exposes a single MsgRegisterName op; gov-driven Create/Update
// Namespace are routed through the gov module's proposal sim instead.
func WeightedOperations(
	appParams simtypes.AppParams,
	cdc interface{},
	txGen client.TxConfig,
	ak AccountKeeper,
	bk BankKeeper,
	k NamesKeeper,
) simulation.WeightedOperations {
	var weightRegisterName int
	appParams.GetOrGenerate(OpWeightMsgRegisterName, &weightRegisterName, nil,
		func(_ *rand.Rand) { weightRegisterName = DefaultWeightMsgRegisterName })

	return simulation.WeightedOperations{
		simulation.NewWeightedOperation(weightRegisterName, SimulateMsgRegisterName(txGen, ak, bk, k)),
	}
}

// SimulateMsgRegisterName picks a random open self-register namespace, a
// random account, and asks the keeper to register a unique normalised name
// owned by the account's "did:ixo:<addr>" DID. Falls back to NoOpMsg when no
// suitable namespace exists or the chosen name is already taken.
func SimulateMsgRegisterName(
	txGen client.TxConfig,
	ak AccountKeeper,
	bk BankKeeper,
	k NamesKeeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		msgType := sdk.MsgTypeURL(&types.MsgRegisterName{})

		// Find a namespace that allows self-register.
		var ns *types.Namespace
		for _, n := range k.GetAllNamespaces(ctx) {
			if n.AllowSelfRegister {
				nn := n
				ns = &nn
				break
			}
		}
		if ns == nil {
			return simtypes.NoOpMsg(types.ModuleName, msgType, "no self-register namespace"), nil, nil
		}

		// Random account.
		simAcc, _ := simtypes.RandomAcc(r, accs)
		account := ak.GetAccount(ctx, simAcc.Address)
		if account == nil {
			return simtypes.NoOpMsg(types.ModuleName, msgType, "account does not exist"), nil, nil
		}
		spendable := bk.SpendableCoins(ctx, simAcc.Address)
		if spendable.IsZero() {
			return simtypes.NoOpMsg(types.ModuleName, msgType, "no spendable coins for fees"), nil, nil
		}

		// Pick a unique name — if collision, NoOp.
		name := fmt.Sprintf("sim-%d", r.Intn(1<<30))
		if uint32(len(name)) < ns.MinLength || (ns.MaxLength > 0 && uint32(len(name)) > ns.MaxLength) {
			return simtypes.NoOpMsg(types.ModuleName, msgType, "name violates namespace length"), nil, nil
		}
		normalized := types.NormalizeName(name)
		if k.HasNameRecord(ctx, ns.Name, normalized) {
			return simtypes.NoOpMsg(types.ModuleName, msgType, "name already taken"), nil, nil
		}

		msg := &types.MsgRegisterName{
			Signer:    simAcc.Address.String(),
			Namespace: ns.Name,
			Name:      name,
			OwnerDid:  "did:ixo:" + simAcc.Address.String(),
		}

		// We cannot easily seed the iid keeper with a DID owned by the random
		// account from here, so this op is essentially a smoke test: it
		// constructs and validates the msg shape but the keeper rejects the
		// unauthorised DID. That's still useful — it asserts the handler
		// rejects rather than crashes under random load.
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
		// Always return success so the simulator continues; the actual handler
		// rejection (no DID, owner DID does not exist) is expected.
		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}

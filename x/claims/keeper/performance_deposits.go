package keeper

import (
	errorsmod "cosmossdk.io/errors"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-blockchain/v6/x/claims/types"
)

// --------------------------
// AGENT DEPOSIT BALANCES
// --------------------------
//
// One rolling balance per (collectionId, agentAddress). Funds live inside
// the collection's existing escrow account (collection.escrow_account),
// which now holds three logically-separate buckets:
//   1. intent funds (existing — pre-paid APPROVAL on MsgClaimIntent)
//   2. agent performance-deposit balances (this file)
//   3. locked dispute deposits (per Dispute record, see disputes.go)
//
// Accounting is kept in state; bank balance is the sum across all three.

// SetAgentDepositBalance persists a balance entry. Callers must not pass
// a zero-amount balance; use RemoveAgentDepositBalance for that case to
// keep the store free of empty rows.
func (k Keeper) SetAgentDepositBalance(ctx sdk.Context, b types.AgentDepositBalance) {
	key := types.AgentDepositBalanceKeyCreate(b.CollectionId, b.AgentAddress)
	k.Set(ctx, key, types.AgentDepositBalanceKey, b, k.Marshal)
}

// GetAgentDepositBalance returns the current balance for an agent on a
// collection. Returns ErrAgentDepositNotFound if none exists.
func (k Keeper) GetAgentDepositBalance(ctx sdk.Context, collectionId, agentAddress string) (types.AgentDepositBalance, error) {
	key := types.AgentDepositBalanceKeyCreate(collectionId, agentAddress)
	val, found := k.Get(ctx, key, types.AgentDepositBalanceKey, k.UnmarshalAgentDepositBalance)
	if !found {
		return types.AgentDepositBalance{}, errorsmod.Wrapf(types.ErrAgentDepositNotFound, "for collection %s agent %s", collectionId, agentAddress)
	}
	balance, ok := val.(types.AgentDepositBalance)
	if !ok {
		return types.AgentDepositBalance{}, errorsmod.Wrapf(types.ErrAgentDepositNotFound, "for collection %s agent %s", collectionId, agentAddress)
	}
	return balance, nil
}

// GetAgentDepositBalanceOrZero returns the balance entry if it exists,
// otherwise a fresh zero-amount entry seeded with the provided keys.
// Useful before adding to / draining from a balance without having to
// branch on existence at the call site.
func (k Keeper) GetAgentDepositBalanceOrZero(ctx sdk.Context, collectionId, agentAddress string) types.AgentDepositBalance {
	balance, err := k.GetAgentDepositBalance(ctx, collectionId, agentAddress)
	if err != nil {
		return types.AgentDepositBalance{
			CollectionId: collectionId,
			AgentAddress: agentAddress,
			Amount:       sdk.NewCoins(),
		}
	}
	return balance
}

// RemoveAgentDepositBalance deletes the KV entry. Should only be called
// once the balance has reached zero (e.g. fully withdrawn or fully slashed).
func (k Keeper) RemoveAgentDepositBalance(ctx sdk.Context, collectionId, agentAddress string) {
	key := types.AgentDepositBalanceKeyCreate(collectionId, agentAddress)
	k.Delete(ctx, key, types.AgentDepositBalanceKey)
}

// UnmarshalAgentDepositBalance decodes a stored balance and validates it
// is well-formed. Used by Keeper.Get.
func (k Keeper) UnmarshalAgentDepositBalance(value []byte) (interface{}, bool) {
	data := types.AgentDepositBalance{}
	if !k.Unmarshal(value, &data) {
		return data, false
	}
	return data, types.IsValidAgentDepositBalance(&data)
}

// GetAllAgentDepositBalances returns every balance entry (used by
// genesis export and the unfiltered query).
func (k Keeper) GetAllAgentDepositBalances(ctx sdk.Context) []types.AgentDepositBalance {
	var out []types.AgentDepositBalance
	iter := k.GetAll(ctx, types.AgentDepositBalanceKey)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var b types.AgentDepositBalance
		k.Unmarshal(iter.Value(), &b)
		out = append(out, b)
	}
	return out
}

// GetCollectionAgentDepositBalances returns every balance entry under a
// specific collection. Used by the per-collection query.
func (k Keeper) GetCollectionAgentDepositBalances(ctx sdk.Context, collectionId string) []types.AgentDepositBalance {
	var out []types.AgentDepositBalance
	prefix := []byte(collectionId + "/")
	iter := k.GetAll(ctx, append(types.AgentDepositBalanceKey, prefix...))
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var b types.AgentDepositBalance
		k.Unmarshal(iter.Value(), &b)
		out = append(out, b)
	}
	return out
}

// --------------------------
// HIGHER-LEVEL OPERATIONS
// --------------------------

// AddPerformanceDeposit moves `amount` from agentAddress's wallet into the
// collection's escrow account and increments the stored balance entry.
// Permitted whether or not the agent currently has active disputes —
// only withdrawal and new submissions are gated.
//
// Emits AgentDepositBalanceCreatedEvent on the first add for a given
// (collection, agent) pair, or AgentDepositBalanceUpdatedEvent on every
// subsequent top-up.
func (k Keeper) AddPerformanceDeposit(
	ctx sdk.Context,
	collectionId string,
	agentAddress sdk.AccAddress,
	amount sdk.Coins,
) (types.AgentDepositBalance, error) {
	if !amount.IsAllPositive() {
		return types.AgentDepositBalance{}, errorsmod.Wrap(types.ErrAgentDepositAmountInvalid, "amount must be strictly positive")
	}

	collection, err := k.GetCollection(ctx, collectionId)
	if err != nil {
		return types.AgentDepositBalance{}, err
	}

	escrow, err := sdk.AccAddressFromBech32(collection.EscrowAccount)
	if err != nil {
		return types.AgentDepositBalance{}, errorsmod.Wrapf(types.ErrInternalError, "invalid collection escrow address: %s", err)
	}

	// Move funds first, then update state. SendCoins is atomic on tx failure;
	// if it fails the state mutation below is rolled back too.
	if err := k.BankKeeper.SendCoins(ctx, agentAddress, escrow, amount); err != nil {
		return types.AgentDepositBalance{}, err
	}

	// Detect "first ever top-up" by checking whether a balance entry exists
	// before we mutate. Drives the Created-vs-Updated event branch — same
	// pattern as MemberBudget on MsgSetCollectionMembers.
	_, existingErr := k.GetAgentDepositBalance(ctx, collectionId, agentAddress.String())
	isNew := existingErr != nil

	balance := k.GetAgentDepositBalanceOrZero(ctx, collectionId, agentAddress.String())
	balance.Amount = balance.Amount.Add(amount...)

	// Roll WithdrawableAt forward to max(current, now + min_deposit_period).
	// Closes the in-same-tx deposit + submit + withdraw exploit: any top-up
	// extends the lock so the agent always has at least min_deposit_period
	// of skin in the game after their most recent deposit. min_deposit_period
	// == 0 keeps WithdrawableAt at the zero time and effectively disables
	// the lock (legacy behavior).
	if collection.MinDepositPeriod > 0 {
		newWithdrawableAt := ctx.BlockTime().Add(collection.MinDepositPeriod)
		if balance.WithdrawableAt == nil || balance.WithdrawableAt.Before(newWithdrawableAt) {
			balance.WithdrawableAt = &newWithdrawableAt
		}
	}

	k.SetAgentDepositBalance(ctx, balance)

	if isNew {
		if err := ctx.EventManager().EmitTypedEvent(&types.AgentDepositBalanceCreatedEvent{
			Balance: &balance,
		}); err != nil {
			return types.AgentDepositBalance{}, err
		}
	} else {
		if err := ctx.EventManager().EmitTypedEvent(&types.AgentDepositBalanceUpdatedEvent{
			Balance: &balance,
		}); err != nil {
			return types.AgentDepositBalance{}, err
		}
	}

	return balance, nil
}

// WithdrawPerformanceDeposit pulls funds from the collection escrow back
// to the agent. If `amount` is empty, withdraws the full current balance.
// Rejected if the agent has any OPEN dispute targeting them on this
// collection (the active-dispute index governs this; see disputes.go).
//
// Emits AgentDepositBalanceUpdatedEvent on partial withdrawal, or
// AgentDepositBalanceRemovedEvent when the balance drains to zero and the
// KV entry is deleted (event carries the final zero-amount balance for
// indexer archival).
func (k Keeper) WithdrawPerformanceDeposit(
	ctx sdk.Context,
	collectionId string,
	agentAddress sdk.AccAddress,
	amount sdk.Coins,
) (sdk.Coins, sdk.Coins, error) {
	if k.HasActiveDisputeAgainstAgent(ctx, collectionId, agentAddress.String()) {
		return nil, nil, errorsmod.Wrapf(types.ErrAgentDepositBalanceCannotWithdraw,
			"agent %s has open disputes on collection %s", agentAddress.String(), collectionId)
	}

	balance, err := k.GetAgentDepositBalance(ctx, collectionId, agentAddress.String())
	if err != nil {
		return nil, nil, err
	}

	// Minimum-deposit-period gate. Set on each AddPerformanceDeposit to
	// max(current, now + collection.min_deposit_period); reading it back
	// here ensures every top-up gets at least one min_deposit_period window
	// of lock-in before the agent can pull the funds. Nil / zero withdrawable_at
	// means "no lock" (legacy collections, or collections where
	// min_deposit_period is zero).
	if balance.WithdrawableAt != nil && ctx.BlockTime().Before(*balance.WithdrawableAt) {
		return nil, nil, errorsmod.Wrapf(types.ErrAgentDepositLocked,
			"locked until %s; now is %s", balance.WithdrawableAt.String(), ctx.BlockTime().String())
	}

	// Empty amount means "withdraw everything currently in balance".
	if len(amount) == 0 {
		amount = balance.Amount
	}

	// Withdrawal request must validate as coins, be positive, and not exceed balance.
	if err := amount.Sort().Validate(); err != nil {
		return nil, nil, errorsmod.Wrapf(types.ErrAgentDepositAmountInvalid, "%s", err)
	}
	if !amount.IsAllPositive() {
		return nil, nil, errorsmod.Wrap(types.ErrAgentDepositAmountInvalid, "amount must be strictly positive")
	}
	if !amount.IsAllLTE(balance.Amount) {
		return nil, nil, errorsmod.Wrapf(types.ErrAgentDepositAmountExceedsBalance,
			"want %s, have %s", amount, balance.Amount)
	}

	collection, err := k.GetCollection(ctx, collectionId)
	if err != nil {
		return nil, nil, err
	}
	escrow, err := sdk.AccAddressFromBech32(collection.EscrowAccount)
	if err != nil {
		return nil, nil, errorsmod.Wrapf(types.ErrInternalError, "invalid collection escrow address: %s", err)
	}

	if err := k.BankKeeper.SendCoins(ctx, escrow, agentAddress, amount); err != nil {
		return nil, nil, err
	}

	balance.Amount = balance.Amount.Sub(amount...)

	if balance.Amount.IsZero() {
		k.RemoveAgentDepositBalance(ctx, collectionId, agentAddress.String())
		if err := ctx.EventManager().EmitTypedEvent(&types.AgentDepositBalanceRemovedEvent{
			Balance: &balance,
		}); err != nil {
			return nil, nil, err
		}
	} else {
		k.SetAgentDepositBalance(ctx, balance)
		if err := ctx.EventManager().EmitTypedEvent(&types.AgentDepositBalanceUpdatedEvent{
			Balance: &balance,
		}); err != nil {
			return nil, nil, err
		}
	}

	return amount, balance.Amount, nil
}

// SlashAgentDepositBalance debits up-to `intended` from the agent's balance
// and moves the actually-slashed amount to `destination`. Returns the
// actual amount slashed, which may be less than `intended` if the balance
// was short. If balance hits zero, the entry is removed.
//
// Designed so multiple disputes against the same agent drain the balance
// in adjudication order; later disputes can come up empty if a prior one
// already took everything.
//
// Caller is responsible for choosing the destination (e.g. winner address
// for the 80% share, adjudicator payout for the 20%). For the typical
// 80/20 split call this twice with the pre-computed amounts.
//
// Emits AgentDepositBalanceUpdatedEvent on partial slash, or
// AgentDepositBalanceRemovedEvent when the slash drains the balance to
// zero and the KV entry is deleted. If the agent had no balance entry
// to start with, returns zero coins and emits no event.
func (k Keeper) SlashAgentDepositBalance(
	ctx sdk.Context,
	collectionId string,
	agentAddress sdk.AccAddress,
	intended sdk.Coins,
	destination sdk.AccAddress,
) (sdk.Coins, error) {
	balance, err := k.GetAgentDepositBalance(ctx, collectionId, agentAddress.String())
	if err != nil {
		// No balance means no slashable funds. Not an error — the dispute
		// resolution still proceeds, just with zero pay-out.
		return sdk.NewCoins(), nil
	}

	// Slash min(intended, balance) per coin.
	actual := minCoinsPerDenom(intended, balance.Amount)
	if actual.IsZero() {
		return actual, nil
	}

	collection, err := k.GetCollection(ctx, collectionId)
	if err != nil {
		return nil, err
	}
	escrow, err := sdk.AccAddressFromBech32(collection.EscrowAccount)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrInternalError, "invalid collection escrow address: %s", err)
	}

	if err := k.BankKeeper.SendCoins(ctx, escrow, destination, actual); err != nil {
		return nil, err
	}

	balance.Amount = balance.Amount.Sub(actual...)
	if balance.Amount.IsZero() {
		k.RemoveAgentDepositBalance(ctx, collectionId, agentAddress.String())
		if err := ctx.EventManager().EmitTypedEvent(&types.AgentDepositBalanceRemovedEvent{
			Balance: &balance,
		}); err != nil {
			return nil, err
		}
	} else {
		k.SetAgentDepositBalance(ctx, balance)
		if err := ctx.EventManager().EmitTypedEvent(&types.AgentDepositBalanceUpdatedEvent{
			Balance: &balance,
		}); err != nil {
			return nil, err
		}
	}

	return actual, nil
}

// HasAgentMetDepositRequirement returns true when the agent's balance on
// `collectionId` is greater than or equal to `required` for every denom in
// `required`. `required` may be empty/zero, in which case the gate is open.
func (k Keeper) HasAgentMetDepositRequirement(
	ctx sdk.Context,
	collectionId, agentAddress string,
	required sdk.Coins,
) bool {
	if required.IsZero() {
		return true
	}
	balance := k.GetAgentDepositBalanceOrZero(ctx, collectionId, agentAddress)
	return required.IsAllLTE(balance.Amount)
}

// --------------------------
// HELPERS
// --------------------------

// minCoinsPerDenom returns, for each denom present in both `a` and `b`, the
// smaller amount. Denoms that appear in only one input are dropped. The
// result is a valid sorted sdk.Coins (no zero amounts).
//
// Used for slash math: "take min(intended, available) per denom" so a
// short balance results in a partial — never negative — slash.
func minCoinsPerDenom(a, b sdk.Coins) sdk.Coins {
	out := sdk.NewCoins()
	for _, ca := range a {
		bv := b.AmountOf(ca.Denom)
		if bv.IsZero() {
			continue
		}
		amt := ca.Amount
		if bv.LT(amt) {
			amt = bv
		}
		if amt.IsPositive() {
			out = out.Add(sdk.NewCoin(ca.Denom, amt))
		}
	}
	return out
}

// Iterator returns the storetype iterator over balances (used by genesis).
func (k Keeper) GetAgentDepositBalancesIterator(ctx sdk.Context) storetypes.Iterator {
	return k.GetAll(ctx, types.AgentDepositBalanceKey)
}

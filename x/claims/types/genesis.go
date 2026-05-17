package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := gs.Params.Validate(); err != nil {
		return err
	}

	// Validate member budgets: catch obviously malformed entries that could
	// cause runtime issues (e.g., zero/short period triggering tight reset
	// loops, malformed addresses, duplicate keys).
	seen := make(map[string]bool, len(gs.MemberBudgets))
	for _, mb := range gs.MemberBudgets {
		if mb.CollectionId == "" {
			return fmt.Errorf("member budget has empty collection_id")
		}
		if _, err := sdk.AccAddressFromBech32(mb.MemberAddress); err != nil {
			return fmt.Errorf("member budget has invalid member_address %s: %w", mb.MemberAddress, err)
		}
		key := mb.CollectionId + "/" + mb.MemberAddress
		if seen[key] {
			return fmt.Errorf("duplicate member budget for collection %s member %s", mb.CollectionId, mb.MemberAddress)
		}
		seen[key] = true
		if mb.Period < MinMemberBudgetPeriod {
			return fmt.Errorf("member budget period must be at least %s for collection %s member %s", MinMemberBudgetPeriod, mb.CollectionId, mb.MemberAddress)
		}
		if mb.PeriodSpendLimit.IsZero() && len(mb.PeriodCw20SpendLimit) == 0 {
			return fmt.Errorf("member budget has no spend limits for collection %s member %s", mb.CollectionId, mb.MemberAddress)
		}
		if err := mb.PeriodSpendLimit.Validate(); err != nil {
			return fmt.Errorf("invalid period_spend_limit for collection %s member %s: %w", mb.CollectionId, mb.MemberAddress, err)
		}
		if err := mb.PeriodSpent.Validate(); err != nil {
			return fmt.Errorf("invalid period_spent for collection %s member %s: %w", mb.CollectionId, mb.MemberAddress, err)
		}
	}

	// Validate agent deposit balances
	seenBalance := make(map[string]bool, len(gs.AgentDepositBalances))
	for _, b := range gs.AgentDepositBalances {
		if b.CollectionId == "" {
			return fmt.Errorf("agent deposit balance has empty collection_id")
		}
		if _, err := sdk.AccAddressFromBech32(b.AgentAddress); err != nil {
			return fmt.Errorf("agent deposit balance has invalid agent_address %s: %w", b.AgentAddress, err)
		}
		key := b.CollectionId + "/" + b.AgentAddress
		if seenBalance[key] {
			return fmt.Errorf("duplicate agent deposit balance for collection %s agent %s", b.CollectionId, b.AgentAddress)
		}
		seenBalance[key] = true
		if err := b.Amount.Validate(); err != nil {
			return fmt.Errorf("invalid amount for collection %s agent %s: %w", b.CollectionId, b.AgentAddress, err)
		}
		if b.Amount.IsZero() {
			return fmt.Errorf("agent deposit balance for collection %s agent %s is zero (entry should be omitted)", b.CollectionId, b.AgentAddress)
		}
	}
	return nil
}

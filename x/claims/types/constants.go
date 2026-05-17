package types

import "time"

const (
	DefaultIxoAccount         = "ixo1kqmtxkggcqa9u34lnr6shy0euvclgatw4f9zz5"
	DefaultCollectionSequence = uint64(1)
	DefaultIntentSequence     = uint64(1)

	// EntityAdjudicatorRevenueAccountName is the entity-account name created
	// on an adjudicator entity when it first receives an adjudication payout.
	// Mirrors EntityOracleRevenueAccountName from the oracle-payment flow.
	// Only used when the adjudicator DID resolves to an entity with module
	// accounts; non-entity DIDs (e.g. did:ixo:wasm:...) are paid out to the
	// adjudicator_address directly instead.
	EntityAdjudicatorRevenueAccountName = "adjudicatorRevenue"
)

// MinMemberBudgetPeriod is the minimum allowed period for a member budget.
// Enforced to prevent griefing via tiny periods that could cause heavy work
// in the lazy-reset loop (which advances PeriodResetAt one period at a time
// to catch up to the current block time).
// const MinMemberBudgetPeriod = 4 * time.Minute // TEMP: lowered from 24*time.Hour for SDK testing
const MinMemberBudgetPeriod = 24 * time.Hour

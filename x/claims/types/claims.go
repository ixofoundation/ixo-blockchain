package types

import (
	"fmt"
	"time"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	entitytypes "github.com/ixofoundation/ixo-blockchain/v7/x/entity/types"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v7/x/iid/types"
)

// IsValidCollection tells if a Claim Collection is valid,
func IsValidCollection(collection *Collection) bool {
	if collection == nil {
		return false
	}
	if iidtypes.IsEmpty(collection.Id) {
		return false
	}
	_, err := sdk.AccAddressFromBech32(collection.Admin)
	if err != nil {
		return false
	}
	if !iidtypes.IsValidDID(collection.Entity) {
		return false
	}
	if !iidtypes.IsValidDID(collection.Protocol) {
		return false
	}
	return true
}

// IsValidClaim tells if a Claim is valid,
func IsValidClaim(claim *Claim) bool {
	if claim == nil {
		return false
	}
	if iidtypes.IsEmpty(claim.ClaimId) {
		return false
	}
	if !iidtypes.IsValidDID(claim.AgentDid) {
		return false
	}
	return true
}

// IsValidDispute tells if a Dispute is valid,
func IsValidDispute(dispute *Dispute) bool {
	if dispute == nil {
		return false
	}
	if iidtypes.IsEmpty(dispute.SubjectId) {
		return false
	}
	if iidtypes.IsEmpty(dispute.Data.Proof) {
		return false
	}
	if iidtypes.IsEmpty(dispute.Data.Uri) {
		return false
	}
	return true
}

// IsValidIntent tells if a Intent is valid,
func IsValidIntent(intent *Intent) bool {
	if intent == nil {
		return false
	}
	if iidtypes.IsEmpty(intent.AgentDid) {
		return false
	}
	if iidtypes.IsEmpty(intent.Id) {
		return false
	}
	_, err := sdk.AccAddressFromBech32(intent.AgentAddress)
	if err != nil {
		return false
	}
	if iidtypes.IsEmpty(intent.CollectionId) {
		return false
	}
	_, err = sdk.AccAddressFromBech32(intent.FromAddress)
	if err != nil {
		return false
	}
	_, err = sdk.AccAddressFromBech32(intent.EscrowAddress)
	return err == nil
}

func HasBalances(ctx sdk.Context, bankKeeper bankkeeper.Keeper, payerAddr sdk.AccAddress,
	requiredFunds sdk.Coins) bool {
	for _, coin := range requiredFunds {
		if !bankKeeper.HasBalance(ctx, payerAddr, coin) {
			return false
		}
	}

	return true
}

func (p Payment) Validate(allowOraclePayments bool) error {
	_, err := sdk.AccAddressFromBech32(p.Account)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "err %s", err)
	}

	if !allowOraclePayments && p.IsOraclePayment {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "oracle payments is only allowed for APPROVAL payments")
	}

	if p.Contract_1155Payment != nil {
		return fmt.Errorf("contract_1155_payment is deprecated, use cw1155_payment instead")
	}

	// no 0 amounts allowed, otherwise unnecessary 0 amount payments
	if err = ValidateCW20Payments(p.Cw20Payment, false); err != nil {
		return err
	}

	// no 0 amounts allowed, otherwise unnecessary 0 amount payments
	if err = p.Amount.Sort().Validate(); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "amounts not valid: (%s)", err)
	}

	// no 0 amounts allowed, otherwise unnecessary 0 amount payments
	if err = ValidateCW1155Payments(p.Cw1155Payment, false); err != nil {
		return err
	}

	return nil
}

func (p Payments) AccountsIsEntityAccounts(entity entitytypes.Entity) bool {
	if !entity.ContainsAccountAddress(p.Approval.Account) || !entity.ContainsAccountAddress(p.Submission.Account) || !entity.ContainsAccountAddress(p.Rejection.Account) || !entity.ContainsAccountAddress(p.Evaluation.Account) {
		return false
	}
	return true
}

func (p Payments) Validate() error {
	// if evaluation payment has cw20 payments, it is not allowed
	if len(p.Evaluation.Cw20Payment) > 1 {
		return ErrCollectionEvalCW20Error
	}
	// if evaluation payment has cw1155 payments, it is not allowed
	if len(p.Evaluation.Cw1155Payment) > 1 {
		return ErrCollectionEvalCW1155Error
	}
	// if approval is oracle payment then no cw1155 payments allowed
	if p.Approval.IsOraclePayment && len(p.Approval.Cw1155Payment) > 0 {
		return ErrCollectionApprovalCW1155Error
	}

	if err := p.Submission.Validate(false); err != nil {
		return err
	}
	if err := p.Evaluation.Validate(false); err != nil {
		return err
	}
	if err := p.Approval.Validate(true); err != nil {
		return err
	}
	if err := p.Rejection.Validate(false); err != nil {
		return err
	}

	return nil
}

// Creates a deep copy of Payment
func (p *Payment) Clone() *Payment {
	if p == nil {
		return nil
	}

	var contract1155Payment *Contract1155Payment
	if p.Contract_1155Payment != nil {
		derefedContract1155Payment := *p.Contract_1155Payment
		contract1155Payment = &derefedContract1155Payment
	}

	cloned := &Payment{
		Account:              p.Account,
		Amount:               p.Amount,
		Contract_1155Payment: contract1155Payment,
		TimeoutNs:            p.TimeoutNs,
		IsOraclePayment:      p.IsOraclePayment,
	}

	if p.Cw20Payment != nil {
		cloned.Cw20Payment = make([]*CW20Payment, len(p.Cw20Payment))
		// Deep copy the cw20_payment field (slice of CW20Payment)
		for i, cw20 := range p.Cw20Payment {
			if cw20 != nil {
				derefedCW20Payment := *cw20
				cloned.Cw20Payment[i] = &derefedCW20Payment
			} else {
				cloned.Cw20Payment[i] = nil
			}
		}
	} else {
		cloned.Cw20Payment = nil
	}
	if p.Cw1155Payment != nil {
		cloned.Cw1155Payment = make([]*CW1155Payment, len(p.Cw1155Payment))
		for i, cw1155 := range p.Cw1155Payment {
			if cw1155 != nil {
				derefedCW1155Payment := *cw1155
				cloned.Cw1155Payment[i] = &derefedCW1155Payment
			} else {
				cloned.Cw1155Payment[i] = nil
			}
		}
	} else {
		cloned.Cw1155Payment = nil
	}

	return cloned
}

// Helper to get module account key in form of claims_escrow_{collectionId}
func GetModuleAccountKeyEscrow(collectionId string) string {
	return ModuleName + "_escrow_" + collectionId
}

// Helper to get module account address
func GetModuleAccountAddressEscrow(collectionId string) sdk.AccAddress {
	return authtypes.NewModuleAddress(GetModuleAccountKeyEscrow(collectionId))
}

// IsCoinsInMaxConstraints checks if the provided coins are within the max constraints
func IsCoinsInMaxConstraints(coins sdk.Coins, maxCoins sdk.Coins) bool {
	maxCoinsMap := make(map[string]sdk.Coin)
	for _, maxCoin := range maxCoins {
		maxCoinsMap[maxCoin.Denom] = maxCoin
	}

	for _, coin := range coins {
		maxCoin, ok := maxCoinsMap[coin.Denom]
		if !ok || !coin.IsLTE(maxCoin) {
			return false
		}
	}
	return true
}

// Validate checks that the Coins are sorted, have positive amount or zero, with a valid and unique
// denomination (i.e no duplicates). Otherwise, it returns an error. Copied from sdk.Coins.Validate()
func ValidateCoinsAllowZero(coins sdk.Coins) error {
	switch len(coins) {
	case 0:
		return nil

	case 1:
		if err := sdk.ValidateDenom(coins[0].Denom); err != nil {
			return err
		}
		if coins[0].IsNegative() {
			return fmt.Errorf("coin %s amount is negative", coins[0])
		}
		return nil

	default:
		// check single coin case
		if err := ValidateCoinsAllowZero(sdk.Coins{coins[0]}); err != nil {
			return err
		}

		lowDenom := coins[0].Denom

		for _, coin := range coins[1:] {
			if err := sdk.ValidateDenom(coin.Denom); err != nil {
				return err
			}
			if coin.Denom < lowDenom {
				return fmt.Errorf("denomination %s is not sorted", coin.Denom)
			}
			if coin.Denom == lowDenom {
				return fmt.Errorf("duplicate denomination %s", coin.Denom)
			}
			if coin.IsNegative() {
				return fmt.Errorf("coin %s amount is negative", coin.Denom)
			}

			// we compare each coin against the last denom
			lowDenom = coin.Denom
		}

		return nil
	}
}

// Create a module account for entity id and name of account as fragment in form: did#name
func CreateNewCollectionEscrow(ctx sdk.Context, accKeeper AccountKeeper, collectionId string) (sdk.AccAddress, error) {
	address := GetModuleAccountAddressEscrow(collectionId)

	if accKeeper.GetAccount(ctx, address) != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "account already exists")
	}

	account := accKeeper.NewAccountWithAddress(ctx, address)
	accKeeper.SetAccount(ctx, account)

	return account.GetAddress(), nil
}

// IsValidAgentDepositBalance tells if an AgentDepositBalance is well-formed
// (has the keys to live in state and a sortable Coins amount).
func IsValidAgentDepositBalance(b *AgentDepositBalance) bool {
	if b == nil {
		return false
	}
	if iidtypes.IsEmpty(b.CollectionId) {
		return false
	}
	if _, err := sdk.AccAddressFromBech32(b.AgentAddress); err != nil {
		return false
	}
	if err := ValidateCoinsAllowZero(b.Amount.Sort()); err != nil {
		return false
	}
	return true
}

// CollectionDisputeConfig is the subset of Collection fields that govern
// dispute resolution + performance deposits. Pulled out into a separate
// helper so the same validation runs for both MsgCreateCollection (which
// inlines these fields) and MsgUpdateCollectionDisputeConfig.
type CollectionDisputeConfig struct {
	ServiceAgentDepositRequired sdk.Coins
	EvaluatorDepositRequired    sdk.Coins
	DisputeDepositAmount        sdk.Coins
	Adjudicators                []*AdjudicationDid
	PenaltyAmountPerDispute     sdk.Coins
	MinDepositPeriod            time.Duration
}

// ValidateCollectionDisputeConfig enforces the cross-field invariants for
// dispute/performance-deposit config:
//   - any non-empty deposit amount must validate as Coins
//   - if any deposit-required, disputer stake, or penalty is set, the
//     collection MUST have at least one adjudication_entity_did
//   - adjudicators entries must each have a valid DID and a reward_percentage
//     in [0, 100]; DIDs must be unique within the list
//   - penalty_amount_per_dispute (if set) must be ≤ each non-empty
//     deposit-required field
//   - min_deposit_period must be non-negative
//
// Callers should populate cfg with the values they intend to persist; this
// runs no IO and never touches state.
func ValidateCollectionDisputeConfig(cfg CollectionDisputeConfig) error {
	// 1. amounts well-formed
	if err := ValidateCoinsAllowZero(cfg.ServiceAgentDepositRequired.Sort()); err != nil {
		return errorsmod.Wrapf(ErrDisputeConfigInvalid, "service_agent_deposit_required: %s", err)
	}
	if err := ValidateCoinsAllowZero(cfg.EvaluatorDepositRequired.Sort()); err != nil {
		return errorsmod.Wrapf(ErrDisputeConfigInvalid, "evaluator_deposit_required: %s", err)
	}
	if err := ValidateCoinsAllowZero(cfg.DisputeDepositAmount.Sort()); err != nil {
		return errorsmod.Wrapf(ErrDisputeConfigInvalid, "dispute_deposit_amount: %s", err)
	}
	if err := ValidateCoinsAllowZero(cfg.PenaltyAmountPerDispute.Sort()); err != nil {
		return errorsmod.Wrapf(ErrDisputeConfigInvalid, "penalty_amount_per_dispute: %s", err)
	}

	// 2. if any dispute / deposit feature is configured, adjudicators must be set.
	anyConfigured := !cfg.ServiceAgentDepositRequired.IsZero() ||
		!cfg.EvaluatorDepositRequired.IsZero() ||
		!cfg.DisputeDepositAmount.IsZero() ||
		!cfg.PenaltyAmountPerDispute.IsZero() ||
		len(cfg.Adjudicators) > 0
	if anyConfigured && len(cfg.Adjudicators) == 0 {
		return errorsmod.Wrap(ErrDisputeConfigInvalid,
			"adjudicators is required when any deposit or penalty amount is set")
	}

	// 3. adjudicator entries must each have a valid DID, a reward_percentage
	//    in [0, 100], and DIDs must be unique within the list.
	seen := make(map[string]bool, len(cfg.Adjudicators))
	for _, a := range cfg.Adjudicators {
		if a == nil {
			return errorsmod.Wrap(ErrDisputeConfigInvalid,
				"adjudicators contains a nil entry")
		}
		if !iidtypes.IsValidDID(a.Did) {
			return errorsmod.Wrapf(ErrDisputeConfigInvalid,
				"adjudicators contains invalid DID: %s", a.Did)
		}
		if seen[a.Did] {
			return errorsmod.Wrapf(ErrDisputeConfigInvalid,
				"adjudicators contains duplicate DID: %s", a.Did)
		}
		seen[a.Did] = true
		if !a.RewardPercentage.IsNil() {
			if a.RewardPercentage.IsNegative() || a.RewardPercentage.GT(OneHundred) {
				return errorsmod.Wrapf(ErrDisputeConfigInvalid,
					"adjudicators[%s].reward_percentage must be in [0, 100]", a.Did)
			}
		}
	}

	// 4. penalty cap: if set, must be ≤ each non-empty *role* deposit-required
	//    field. The penalty represents what gets slashed from the LOSER's
	//    agent-deposit balance on AWARDED, so it is bounded by what is
	//    economically available there. On DISMISSED the disputer's stake
	//    (dispute_deposit_amount) IS the pot — the penalty isn't applied to
	//    it — so there's no penalty-vs-dispute_deposit relationship to
	//    enforce here.
	if !cfg.PenaltyAmountPerDispute.IsZero() {
		penalty := cfg.PenaltyAmountPerDispute.Sort()
		if !cfg.ServiceAgentDepositRequired.IsZero() {
			if !penalty.IsAllLTE(cfg.ServiceAgentDepositRequired.Sort()) {
				return errorsmod.Wrap(ErrDisputeConfigInvalid,
					"penalty_amount_per_dispute exceeds service_agent_deposit_required")
			}
		}
		if !cfg.EvaluatorDepositRequired.IsZero() {
			if !penalty.IsAllLTE(cfg.EvaluatorDepositRequired.Sort()) {
				return errorsmod.Wrap(ErrDisputeConfigInvalid,
					"penalty_amount_per_dispute exceeds evaluator_deposit_required")
			}
		}
	}

	// 6. min deposit period must be non-negative. Zero is allowed (no lock,
	//    legacy behavior). No upper bound enforced — the admin is trusted to
	//    pick a sensible value; a too-long period is a foot-gun, not a
	//    safety hazard.
	if cfg.MinDepositPeriod < 0 {
		return errorsmod.Wrap(ErrDisputeConfigInvalid,
			"min_deposit_period must be non-negative")
	}

	return nil
}

// SnapshotDisputeConfigFromCollection projects a Collection's persisted
// dispute-config fields into a CollectionDisputeConfig for validation /
// runtime use.
func SnapshotDisputeConfigFromCollection(c Collection) CollectionDisputeConfig {
	return CollectionDisputeConfig{
		ServiceAgentDepositRequired: c.ServiceAgentDepositRequired,
		EvaluatorDepositRequired:    c.EvaluatorDepositRequired,
		DisputeDepositAmount:        c.DisputeDepositAmount,
		Adjudicators:                c.Adjudicators,
		PenaltyAmountPerDispute:     c.PenaltyAmountPerDispute,
		MinDepositPeriod:            c.MinDepositPeriod,
	}
}

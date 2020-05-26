package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// --------------------------------------------- TestSubscriptionContent

var _ SubscriptionContent = TestSubscriptionContent{}

// TestSubscriptionContent Is identical to BlockSubscriptionContent but does
// not take into consideration the context in ShouldCharge() and started()
type TestSubscriptionContent struct {
	FeeContractId      uint64   `json:"fee_contract_id" yaml:"fee_contract_id"`
	PeriodsSoFar       sdk.Uint `json:"periods_so_far" yaml:"periods_so_far"`
	MaxPeriods         sdk.Uint `json:"max_periods" yaml:"max_periods"`
	PeriodsAccumulated sdk.Uint `json:"periods_accumulated" yaml:"periods_accumulated"`
	PeriodLength       int64    `json:"period_length" yaml:"period_length"`
	PeriodStartBlock   int64    `json:"period_start_block" yaml:"period_start_block"`
	PeriodEndBlock     int64    `json:"period_end_block" yaml:"period_end_block"`
}

func NewTestSubscriptionContent(feeContractId uint64, maxPeriods sdk.Uint,
	periodLength, periodStartBlock int64) TestSubscriptionContent {
	return TestSubscriptionContent{
		FeeContractId:      feeContractId,
		PeriodsSoFar:       sdk.ZeroUint(),
		MaxPeriods:         maxPeriods,
		PeriodsAccumulated: sdk.ZeroUint(),
		PeriodLength:       periodLength,
		PeriodStartBlock:   periodStartBlock,
		PeriodEndBlock:     periodStartBlock + periodLength,
	}
}

func (s TestSubscriptionContent) GetFeeContractId() uint64 {
	return s.FeeContractId
}

func (s TestSubscriptionContent) GetPeriodUnit() string {
	return BlockSubscriptionUnit
}

func (s TestSubscriptionContent) started(ctx sdk.Context) bool {
	return true
}

// Ended True if max number of periods has been reached
func (s TestSubscriptionContent) Ended() bool {
	return s.PeriodsSoFar.GTE(s.MaxPeriods)
}

// ShouldCharge True if started
func (s TestSubscriptionContent) ShouldCharge(ctx sdk.Context) bool {
	return s.started(ctx)
}

func (s TestSubscriptionContent) NextPeriod(periodPaid bool) {
	// Update periods so far (periodsAccumulated if period not paid)
	s.PeriodsSoFar = s.PeriodsSoFar.Add(sdk.OneUint())
	if !periodPaid {
		s.PeriodsAccumulated = s.PeriodsAccumulated.Add(sdk.OneUint())
	}

	// Update period start/end
	s.PeriodStartBlock = s.PeriodEndBlock
	s.PeriodEndBlock = s.PeriodStartBlock + s.PeriodLength
}

func (s TestSubscriptionContent) Validate() sdk.Error {
	if s.PeriodsSoFar.GT(s.MaxPeriods) {
		return ErrInvalidSubscriptionContent(DefaultCodespace, "periods so far is greater than max periods")
	} else if s.PeriodStartBlock > s.PeriodEndBlock {
		return ErrInvalidSubscriptionContent(DefaultCodespace, "start time is after end time")
	} else if s.PeriodLength <= 0 {
		return ErrInvalidSubscriptionContent(DefaultCodespace, "period length must be greater than zero")
	} else if s.PeriodStartBlock+s.PeriodLength != s.PeriodEndBlock {
		return ErrInvalidSubscriptionContent(DefaultCodespace, "period start + period length != period end")
	}
	return nil
}

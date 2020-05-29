package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

const (
	BlockPeriodUnit = "block"
	TimePeriodUnit  = "time"
)

// --------------------------------------------- Subscription and Period

type Subscription struct {
	Id                 string   `json:"id" yaml:"id"`
	FeeContractId      string   `json:"fee_contract_id" yaml:"fee_contract_id"`
	PeriodsSoFar       sdk.Uint `json:"periods_so_far" yaml:"periods_so_far"`
	MaxPeriods         sdk.Uint `json:"max_periods" yaml:"max_periods"`
	PeriodsAccumulated sdk.Uint `json:"periods_accumulated" yaml:"periods_accumulated"`
	Period             Period   `json:"period" yaml:"period"`
}

func (s Subscription) Validate() sdk.Error {

	// Validate IDs
	if !IsValidSubscriptionId(s.Id) {
		return ErrInvalidId(DefaultCodespace, "subscription id invalid")
	} else if !IsValidFeeContractId(s.FeeContractId) {
		return ErrInvalidId(DefaultCodespace, "fee contract id invalid")
	}

	// Verify that periods so far <= max periods
	if s.PeriodsSoFar.GT(s.MaxPeriods) {
		return ErrInvalidPeriod(DefaultCodespace, "periods so far is greater than max periods")
	}

	// Validate period
	return s.Period.Validate()
}

func NewSubscription(id, feeContractId string, maxPeriods sdk.Uint,
	period Period) Subscription {
	return Subscription{
		Id:                 id,
		FeeContractId:      feeContractId,
		PeriodsSoFar:       sdk.ZeroUint(),
		MaxPeriods:         maxPeriods,
		PeriodsAccumulated: sdk.ZeroUint(),
		Period:             period,
	}
}

func (s Subscription) started(ctx sdk.Context) bool {
	return !s.PeriodsSoFar.IsZero() || s.Period.periodStarted(ctx)
}

// Ended True if max number of periods has been reached
func (s Subscription) Ended() bool {
	return s.PeriodsSoFar.GTE(s.MaxPeriods)
}

// NextPeriod Proceed to the next period
func (s Subscription) NextPeriod(periodPaid bool) {

	// Update periods so far (periodsAccumulated if period not paid)
	s.PeriodsSoFar = s.PeriodsSoFar.Add(sdk.OneUint())
	if !periodPaid {
		s.PeriodsAccumulated = s.PeriodsAccumulated.Add(sdk.OneUint())
	}

	// Advance period to next period
	s.Period.nextPeriod()
}

// ShouldCharge True if there's accumulated periods or the period indicates
// that we can charge. In any case, the subscription must have started.
func (s Subscription) ShouldCharge(ctx sdk.Context) bool {
	if !s.started(ctx) {
		return false
	}
	return !s.PeriodsAccumulated.IsZero() || s.Period.periodEnded(ctx)
}

type Period interface {
	GetPeriodUnit() string
	Validate() sdk.Error
	periodStarted(ctx sdk.Context) bool
	periodEnded(ctx sdk.Context) bool
	nextPeriod()
}

var _, _ Period = BlockPeriod{}, TimePeriod{}

// --------------------------------------------- BlockPeriod

type BlockPeriod struct {
	PeriodLength     int64 `json:"period_length" yaml:"period_length"`
	PeriodStartBlock int64 `json:"period_start_block" yaml:"period_start_block"`
}

func NewBlockPeriod(periodLength, periodStartBlock int64) BlockPeriod {
	return BlockPeriod{
		PeriodLength:     periodLength,
		PeriodStartBlock: periodStartBlock,
	}
}

func (s BlockPeriod) periodEndBlock() int64 {
	return s.PeriodStartBlock + s.PeriodLength
}

func (s BlockPeriod) GetPeriodUnit() string {
	return BlockPeriodUnit
}

func (s BlockPeriod) Validate() sdk.Error {

	// Validate period-related values
	if s.PeriodStartBlock > s.periodEndBlock() {
		return ErrInvalidPeriod(DefaultCodespace, "start time is after end time")
	} else if s.PeriodLength <= 0 {
		return ErrInvalidPeriod(DefaultCodespace, "period length must be greater than zero")
	} else if s.PeriodStartBlock+s.PeriodLength != s.periodEndBlock() {
		return ErrInvalidPeriod(DefaultCodespace, "period start + period length != period end")
	}

	return nil
}

func (s BlockPeriod) periodStarted(ctx sdk.Context) bool {
	return ctx.BlockHeight() > s.PeriodStartBlock
}

func (s BlockPeriod) periodEnded(ctx sdk.Context) bool {
	return ctx.BlockHeight() >= s.periodEndBlock()
}

func (s BlockPeriod) nextPeriod() {
	s.PeriodStartBlock = s.periodEndBlock()
}

// --------------------------------------------- TimePeriod

type TimePeriod struct {
	PeriodLength    time.Duration `json:"period_length" yaml:"period_length"`
	PeriodStartTime time.Time     `json:"period_start_time" yaml:"period_start_time"`
}

func NewTimePeriod(periodLength time.Duration, periodStartTime time.Time) TimePeriod {
	return TimePeriod{
		PeriodLength:    periodLength,
		PeriodStartTime: periodStartTime,
	}
}

func (s TimePeriod) periodEndTime() time.Time {
	return s.PeriodStartTime.Add(s.PeriodLength)
}

func (s TimePeriod) GetPeriodUnit() string {
	return TimePeriodUnit
}

func (s TimePeriod) Validate() sdk.Error {

	// Validate period-related values
	if s.PeriodStartTime.After(s.periodEndTime()) {
		return ErrInvalidPeriod(DefaultCodespace, "start time is after end time")
	} else if s.PeriodLength <= 0 {
		return ErrInvalidPeriod(DefaultCodespace, "period length cannot be zero")
	} else if !s.PeriodStartTime.Add(s.PeriodLength).Equal(s.periodEndTime()) {
		return ErrInvalidPeriod(DefaultCodespace, "period start + period length != period end")
	}

	return nil
}

func (s TimePeriod) periodStarted(ctx sdk.Context) bool {
	return ctx.BlockTime().After(s.PeriodStartTime)
}

func (s TimePeriod) periodEnded(ctx sdk.Context) bool {
	return ctx.BlockTime().After(s.periodEndTime())
}

func (s TimePeriod) nextPeriod() {
	s.PeriodStartTime = s.periodEndTime()
}

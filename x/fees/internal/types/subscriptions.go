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

// started True if not the first period, or the current period has started
func (s Subscription) started(ctx sdk.Context) bool {
	return !s.PeriodsSoFar.IsZero() || s.Period.periodStarted(ctx)
}

// MaxPeriodsReached True if max number of periods has been reached
func (s Subscription) MaxPeriodsReached() bool {
	return s.PeriodsSoFar.GTE(s.MaxPeriods)
}

// NextPeriod Proceed to the next period
func (s *Subscription) NextPeriod(periodPaid bool) {

	// Update periods so far (periodsAccumulated if period not paid)
	s.PeriodsSoFar = s.PeriodsSoFar.Add(sdk.OneUint())
	if !periodPaid {
		s.PeriodsAccumulated = s.PeriodsAccumulated.Add(sdk.OneUint())
	}

	// Advance period to next period
	s.Period = s.Period.nextPeriod()
}

// ShouldCharge True if the subscription has started and
//  (A) the max no. of periods has not been reached and the period has ended, or
//  (B) the max no. of periods has been reached but we have accumulated periods
// This means that accumulated periods only get tackled once the max number
// of periods has been reached.
func (s Subscription) ShouldCharge(ctx sdk.Context) bool {
	if !s.started(ctx) {
		return false
	} else if !s.MaxPeriodsReached() {
		return s.Period.periodEnded(ctx)
	} else {
		return !s.PeriodsAccumulated.IsZero()
	}
}

// IsComplete True if we have reached the max number of periods and there are
// no accumulated periods
func (s Subscription) IsComplete() bool {
	return s.MaxPeriodsReached() && s.PeriodsAccumulated.IsZero()
	// equivalent to s.MaxPeriodsReached() && !s.ShouldCharge(ctx)
}

type Period interface {
	GetPeriodUnit() string
	Validate() sdk.Error
	periodStarted(ctx sdk.Context) bool
	periodEnded(ctx sdk.Context) bool
	nextPeriod() Period
}

// --------------------------------------------- BlockPeriod

var _ Period = BlockPeriod{}

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

func (p BlockPeriod) periodEndBlock() int64 {
	return p.PeriodStartBlock + p.PeriodLength
}

func (p BlockPeriod) GetPeriodUnit() string {
	return BlockPeriodUnit
}

func (p BlockPeriod) Validate() sdk.Error {

	// Validate period-related values
	if p.PeriodStartBlock > p.periodEndBlock() {
		return ErrInvalidPeriod(DefaultCodespace, "start time is after end time")
	} else if p.PeriodLength <= 0 {
		return ErrInvalidPeriod(DefaultCodespace, "period length must be greater than zero")
	} else if p.PeriodStartBlock+p.PeriodLength != p.periodEndBlock() {
		return ErrInvalidPeriod(DefaultCodespace, "period start + period length != period end")
	}

	return nil
}

func (p BlockPeriod) periodStarted(ctx sdk.Context) bool {
	return ctx.BlockHeight() > p.PeriodStartBlock
}

func (p BlockPeriod) periodEnded(ctx sdk.Context) bool {
	return ctx.BlockHeight() >= p.periodEndBlock()
}

func (p BlockPeriod) nextPeriod() Period {
	p.PeriodStartBlock = p.periodEndBlock()
	return p
}

// --------------------------------------------- TimePeriod

var _ Period = TimePeriod{}

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

func (p TimePeriod) periodEndTime() time.Time {
	return p.PeriodStartTime.Add(p.PeriodLength)
}

func (p TimePeriod) GetPeriodUnit() string {
	return TimePeriodUnit
}

func (p TimePeriod) Validate() sdk.Error {

	// Validate period-related values
	if p.PeriodStartTime.After(p.periodEndTime()) {
		return ErrInvalidPeriod(DefaultCodespace, "start time is after end time")
	} else if p.PeriodLength <= 0 {
		return ErrInvalidPeriod(DefaultCodespace, "period length cannot be zero")
	} else if !p.PeriodStartTime.Add(p.PeriodLength).Equal(p.periodEndTime()) {
		return ErrInvalidPeriod(DefaultCodespace, "period start + period length != period end")
	}

	return nil
}

func (p TimePeriod) periodStarted(ctx sdk.Context) bool {
	return ctx.BlockTime().After(p.PeriodStartTime)
}

func (p TimePeriod) periodEnded(ctx sdk.Context) bool {
	return ctx.BlockTime().After(p.periodEndTime())
}

func (p TimePeriod) nextPeriod() Period {
	p.PeriodStartTime = p.periodEndTime()
	return p
}

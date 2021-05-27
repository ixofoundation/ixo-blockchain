package types

import (
	"fmt"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"
	"time"
)

const (
	BlockPeriodUnit = "block"
	TimePeriodUnit  = "time"
)

// --------------------------------------------- Subscription and Period

//type Subscription struct {
//	Id                 string   `json:"id" yaml:"id"`
//	PaymentContractId  string   `json:"payment_contract_id" yaml:"payment_contract_id"`
//	PeriodsSoFar       sdk.Uint `json:"periods_so_far" yaml:"periods_so_far"`
//	MaxPeriods         sdk.Uint `json:"max_periods" yaml:"max_periods"`
//	PeriodsAccumulated sdk.Uint `json:"periods_accumulated" yaml:"periods_accumulated"`
//	Period             Period   `json:"period" yaml:"period"`
//}

func (s *Subscription) GetPeriod() Period {
	period, ok := s.Period.GetCachedValue().(Period)
	if !ok {
		return nil
	}
	return period
}

func (s *Subscription) SetPeriod(period Period) error {
	m, ok := period.(proto.Message)
	if !ok {
		return fmt.Errorf("can't proto marshal %T", m)
	}

	any, err := codectypes.NewAnyWithValue(m)
	if err != nil {
		return err
	}

	s.Period = any
	return nil
}

func (s Subscription) Validate() error {

	// Validate IDs
	if !IsValidSubscriptionId(s.Id) {
		return sdkerrors.Wrap(ErrInvalidId, "subscription ID invalid")
	} else if !IsValidPaymentContractId(s.PaymentContractId) {
		return sdkerrors.Wrap(ErrInvalidId, "payment contract ID invalid")
	}

	// Verify that periods so far <= max periods
	if s.PeriodsSoFar.GT(s.MaxPeriods) {
		return sdkerrors.Wrap(ErrInvalidPeriod, "periods so far is greater than max periods")
	}

	// Validate period
	period := s.GetPeriod()
	if period == nil {
		return sdkerrors.Wrap(ErrInvalidPeriod, "missing period")
	}
	return period.Validate()
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (s Subscription) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var period Period
	return unpacker.UnpackAny(s.Period, &period)
}

func NewSubscription(id, contractId string, maxPeriods sdk.Uint, period Period) Subscription {
	subscription := Subscription{
		Id:                 id,
		PaymentContractId:  contractId,
		PeriodsSoFar:       sdk.ZeroUint(),
		MaxPeriods:         maxPeriods,
		PeriodsAccumulated: sdk.ZeroUint(),
	}

	err := subscription.SetPeriod(period)
	if err != nil {
		panic(err)
	}

	return subscription
}

// started True if not the first period, or the current period has started
func (s Subscription) started(ctx sdk.Context) bool {
	return !s.PeriodsSoFar.IsZero() || s.GetPeriod().periodStarted(ctx)
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
	err := s.SetPeriod(s.GetPeriod().nextPeriod())
	if err != nil {
		panic(err)
	}
}

// ShouldEffect True if the subscription has started and
//  (A) the max no. of periods has not been reached and the period has ended, or
//  (B) the max no. of periods has been reached but we have accumulated periods
// This means that accumulated periods only get tackled once the max number
// of periods has been reached.
func (s Subscription) ShouldEffect(ctx sdk.Context) bool {
	if !s.started(ctx) {
		return false
	} else if !s.MaxPeriodsReached() {
		return s.GetPeriod().periodEnded(ctx)
	} else {
		return !s.PeriodsAccumulated.IsZero()
	}
}

// IsComplete True if we have reached the max number of periods and there are
// no accumulated periods
func (s Subscription) IsComplete() bool {
	return s.MaxPeriodsReached() && s.PeriodsAccumulated.IsZero()
	// equivalent to s.MaxPeriodsReached() && !s.ShouldEffect(ctx)
}

type Period interface {
	proto.Message

	GetPeriodUnit() string
	Validate() error
	periodStarted(ctx sdk.Context) bool
	periodEnded(ctx sdk.Context) bool
	nextPeriod() Period
}

// --------------------------------------------- BlockPeriod

var _ Period = &BlockPeriod{}

//type BlockPeriod struct {
//	PeriodLength     int64 `json:"period_length" yaml:"period_length"`
//	PeriodStartBlock int64 `json:"period_start_block" yaml:"period_start_block"`
//}

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

func (p BlockPeriod) Validate() error {

	// Validate period-related values
	if p.PeriodStartBlock > p.periodEndBlock() {
		return sdkerrors.Wrap(ErrInvalidPeriod, "start time is after end time")
	} else if p.PeriodLength <= 0 {
		return sdkerrors.Wrap(ErrInvalidPeriod, "period length must be greater than zero")
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
	return &p
}

// --------------------------------------------- TimePeriod

var _ Period = &TimePeriod{}

//type TimePeriod struct {
//	PeriodDurationNs time.Duration `json:"period_duration_ns" yaml:"period_duration_ns"`
//	PeriodStartTime  time.Time     `json:"period_start_time" yaml:"period_start_time"`
//}

func NewTimePeriod(periodDurationNs time.Duration, periodStartTime time.Time) TimePeriod {
	return TimePeriod{
		PeriodDurationNs: periodDurationNs,
		PeriodStartTime:  periodStartTime,
	}
}

func (p TimePeriod) periodEndTime() time.Time {
	return p.PeriodStartTime.Add(p.PeriodDurationNs)
}

func (p TimePeriod) GetPeriodUnit() string {
	return TimePeriodUnit
}

func (p TimePeriod) Validate() error {

	// Validate period-related values
	if p.PeriodStartTime.After(p.periodEndTime()) {
		return sdkerrors.Wrap(ErrInvalidPeriod, "start time is after end time")
	} else if p.PeriodDurationNs <= 0 {
		return sdkerrors.Wrap(ErrInvalidPeriod, "period duration cannot be zero")
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
	return &p
}

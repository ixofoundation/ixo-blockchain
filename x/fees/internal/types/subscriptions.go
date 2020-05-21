package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"
)

const (
	BlockSubscriptionUnit = "block"
	TimeSubscriptionUnit  = "time"
)

// --------------------------------------------- Subscription and SubscriptionContent

type Subscription struct {
	Id      uint64              `json:"id" yaml:"id"`
	Content SubscriptionContent `json:"content" yaml:"content"`
}

func (s Subscription) Validate() sdk.Error {
	return s.Content.Validate()
}

func NewSubscription(id uint64, content SubscriptionContent) Subscription {
	return Subscription{
		Id:      id,
		Content: content,
	}
}

type SubscriptionContent interface {
	GetFeeContractId() uint64
	GetPeriodUnit() string
	Started(ctx sdk.Context) bool
	Ended() bool
	ShouldCharge(ctx sdk.Context) bool
	NextPeriod(periodPaid bool)
	Validate() sdk.Error
}

var _, _ SubscriptionContent = BlockSubscriptionContent{}, TimeSubscriptionContent{}

// --------------------------------------------- BlockSubscriptionContent

type BlockSubscriptionContent struct {
	FeeContractId      uint64   `json:"fee_contract_id" yaml:"fee_contract_id"`
	PeriodsSoFar       sdk.Uint `json:"periods_so_far" yaml:"periods_so_far"`
	MaxPeriods         sdk.Uint `json:"max_periods" yaml:"max_periods"`
	PeriodsAccumulated sdk.Uint `json:"periods_accumulated" yaml:"periods_accumulated"`
	PeriodLength       int64    `json:"period_length" yaml:"period_length"`
	PeriodStartBlock   int64    `json:"period_start_block" yaml:"period_start_block"`
	PeriodEndBlock     int64    `json:"period_end_block" yaml:"period_end_block"`
}

func NewBlockSubscriptionContent(feeContractId uint64, maxPeriods sdk.Uint,
	periodLength, periodStartBlock int64) BlockSubscriptionContent {
	return BlockSubscriptionContent{
		FeeContractId:      feeContractId,
		PeriodsSoFar:       sdk.ZeroUint(),
		MaxPeriods:         maxPeriods,
		PeriodsAccumulated: sdk.ZeroUint(),
		PeriodLength:       periodLength,
		PeriodStartBlock:   periodStartBlock,
		PeriodEndBlock:     periodStartBlock + periodLength,
	}
}

func (s BlockSubscriptionContent) GetFeeContractId() uint64 {
	return s.FeeContractId
}

func (s BlockSubscriptionContent) GetPeriodUnit() string {
	return BlockSubscriptionUnit
}

//Started True if height has passed start block, or if this is not the first period
func (s BlockSubscriptionContent) Started(ctx sdk.Context) bool {
	return !s.PeriodsSoFar.IsZero() || ctx.BlockHeight() > s.PeriodStartBlock
}

//Ended True if max number of periods has been reached
func (s BlockSubscriptionContent) Ended() bool {
	return s.PeriodsSoFar.GTE(s.MaxPeriods)
}

//ShouldCharge True if end of period reached or there's accumulated periods
func (s BlockSubscriptionContent) ShouldCharge(ctx sdk.Context) bool {
	return !s.PeriodsAccumulated.IsZero() || ctx.BlockHeight() >= s.PeriodEndBlock
}

func (s BlockSubscriptionContent) NextPeriod(periodPaid bool) {
	// Update periods so far (periodsAccumulated if period not paid)
	s.PeriodsSoFar = s.PeriodsSoFar.Add(sdk.OneUint())
	if !periodPaid {
		s.PeriodsAccumulated = s.PeriodsAccumulated.Add(sdk.OneUint())
	}

	// Update period start/end
	s.PeriodStartBlock = s.PeriodEndBlock
	s.PeriodEndBlock = s.PeriodStartBlock + s.PeriodLength
}

func (s BlockSubscriptionContent) Validate() sdk.Error {
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

// --------------------------------------------- TimeSubscriptionContent

type TimeSubscriptionContent struct {
	FeeContractId      uint64        `json:"fee_contract_id" yaml:"fee_contract_id"`
	PeriodsSoFar       sdk.Uint      `json:"periods_so_far" yaml:"periods_so_far"`
	MaxPeriods         sdk.Uint      `json:"max_periods" yaml:"max_periods"`
	PeriodsAccumulated sdk.Uint      `json:"periods_accumulated" yaml:"periods_accumulated"`
	PeriodLength       time.Duration `json:"period_length" yaml:"period_length"`
	PeriodStartTime    time.Time     `json:"period_start_time" yaml:"period_start_time"`
	PeriodEndTime      time.Time     `json:"period_end_time" yaml:"period_end_time"`
}

func NewTimeSubscriptionContent(feeContractId uint64, maxPeriods sdk.Uint,
	periodLength time.Duration, periodStartTime time.Time) TimeSubscriptionContent {
	return TimeSubscriptionContent{
		FeeContractId:      feeContractId,
		PeriodsSoFar:       sdk.ZeroUint(),
		MaxPeriods:         maxPeriods,
		PeriodsAccumulated: sdk.ZeroUint(),
		PeriodLength:       periodLength,
		PeriodStartTime:    periodStartTime,
		PeriodEndTime:      periodStartTime.Add(periodLength),
	}
}

func (s TimeSubscriptionContent) GetFeeContractId() uint64 {
	return s.FeeContractId
}

func (s TimeSubscriptionContent) GetPeriodUnit() string {
	return TimeSubscriptionUnit
}

//Started True if height has passed start block, or if this is not the first period
func (s TimeSubscriptionContent) Started(ctx sdk.Context) bool {
	return !s.PeriodsSoFar.IsZero() || ctx.BlockTime().After(s.PeriodStartTime)
}

//Ended True if max number of periods has been reached
func (s TimeSubscriptionContent) Ended() bool {
	return s.PeriodsSoFar.GTE(s.MaxPeriods)
}

//ShouldCharge True if end of period reached or there's accumulated periods
func (s TimeSubscriptionContent) ShouldCharge(ctx sdk.Context) bool {
	return !s.PeriodsAccumulated.IsZero() || ctx.BlockTime().After(s.PeriodEndTime)
}

//NextPeriod Proceed to the next period
func (s TimeSubscriptionContent) NextPeriod(periodPaid bool) {
	// Update periods so far (periodsAccumulated if period not paid)
	s.PeriodsSoFar = s.PeriodsSoFar.Add(sdk.OneUint())
	if !periodPaid {
		s.PeriodsAccumulated = s.PeriodsAccumulated.Add(sdk.OneUint())
	}

	// Update period start/end
	s.PeriodStartTime = s.PeriodEndTime
	s.PeriodEndTime = s.PeriodStartTime.Add(s.PeriodLength)
}

func (s TimeSubscriptionContent) Validate() sdk.Error {
	if s.PeriodsSoFar.GT(s.MaxPeriods) {
		return ErrInvalidSubscriptionContent(DefaultCodespace, "periods so far is greater than max periods")
	} else if s.PeriodStartTime.After(s.PeriodEndTime) {
		return ErrInvalidSubscriptionContent(DefaultCodespace, "start time is after end time")
	} else if s.PeriodLength <= 0 {
		return ErrInvalidSubscriptionContent(DefaultCodespace, "period length cannot be zero")
	} else if !s.PeriodStartTime.Add(s.PeriodLength).Equal(s.PeriodEndTime) {
		return ErrInvalidSubscriptionContent(DefaultCodespace, "period start + period length != period end")
	}
	return nil
}

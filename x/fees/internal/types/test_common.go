package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// --------------------------------------------- TestPeriod

var _ Period = TestPeriod{}

// TestPeriod Is identical to BlockPeriod but does
// not take into consideration the context in periodEnded() and periodStarted()
type TestPeriod struct {
	PeriodLength     int64 `json:"period_length" yaml:"period_length"`
	PeriodStartBlock int64 `json:"period_start_block" yaml:"period_start_block"`
}

func NewTestPeriod(periodLength, periodStartBlock int64) TestPeriod {
	return TestPeriod{
		PeriodLength:     periodLength,
		PeriodStartBlock: periodStartBlock,
	}
}

func (s TestPeriod) periodEndBlock() int64 {
	return s.PeriodStartBlock + s.PeriodLength
}

func (s TestPeriod) GetPeriodUnit() string {
	return BlockPeriodUnit
}

func (s TestPeriod) Validate() sdk.Error {
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

func (s TestPeriod) periodStarted(ctx sdk.Context) bool {
	return true
}

func (s TestPeriod) periodEnded(ctx sdk.Context) bool {
	return true
}

func (s TestPeriod) nextPeriod() {
	s.PeriodStartBlock = s.periodEndBlock()
}

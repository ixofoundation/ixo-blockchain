package apptesting

import (
	"slices"

	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AssertEventEmitted fails the test unless ctx's event manager has emitted
// exactly numEventsExpected events of type eventTypeExpected.
//
// Use this after invoking a msg handler / keeper method to verify that the
// module is emitting the events its proto contract advertises.
func (s *KeeperTestHelper) AssertEventEmitted(ctx sdk.Context, eventTypeExpected string, numEventsExpected int) {
	allEvents := ctx.EventManager().Events()
	count := 0
	for _, ev := range allEvents {
		if ev.Type == eventTypeExpected {
			count++
		}
	}
	s.Require().Equal(numEventsExpected, count,
		"expected %d events of type %q, got %d", numEventsExpected, eventTypeExpected, count)
}

// FindEvent returns the first abci.Event matching `name`, or a zero-value
// event if not found.
func (s *KeeperTestHelper) FindEvent(events []abci.Event, name string) abci.Event {
	idx := slices.IndexFunc(events, func(e abci.Event) bool { return e.Type == name })
	if idx == -1 {
		return abci.Event{}
	}
	return events[idx]
}

// ExtractAttributes flattens an abci.Event's attributes into a string map for
// readable assertions in tests.
func (s *KeeperTestHelper) ExtractAttributes(event abci.Event) map[string]string {
	attrs := make(map[string]string, len(event.Attributes))
	for _, a := range event.Attributes {
		attrs[a.Key] = a.Value
	}
	return attrs
}

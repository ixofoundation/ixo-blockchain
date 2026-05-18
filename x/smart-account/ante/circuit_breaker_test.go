package ante_test

import (
	"errors"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/ixofoundation/ixo-blockchain/v6/app/apptesting"
	smartaccountante "github.com/ixofoundation/ixo-blockchain/v6/x/smart-account/ante"
)

// CircuitBreakerTestSuite covers the dispatch decision in
// CircuitBreakerDecorator: when smart accounts are inactive (or the tx is
// missing the authenticator extension) the decorator should route to the
// "classic" / original ante chain; otherwise it routes to the authenticator
// chain.
type CircuitBreakerTestSuite struct {
	apptesting.KeeperTestHelper
}

func TestCircuitBreakerTestSuite(t *testing.T) {
	suite.Run(t, new(CircuitBreakerTestSuite))
}

func (s *CircuitBreakerTestSuite) SetupTest() { s.Setup() }

// minimalTx is a no-op sdk.Tx — has no extension options, so the circuit
// breaker treats it as a non-authenticator tx and routes to the classic chain.
type minimalTx struct{ msgs []sdk.Msg }

func (t *minimalTx) GetMsgs() []sdk.Msg                              { return t.msgs }
func (t *minimalTx) GetMsgsV2() ([]protoreflect.ProtoMessage, error) { return nil, nil }
func (t *minimalTx) ValidateBasic() error                            { return nil }

func (s *CircuitBreakerTestSuite) TestCircuitBreaker_RoutesToClassicByDefault() {
	s.SetupTest()

	classicCalled := false
	authCalled := false
	classic := func(ctx sdk.Context, _ sdk.Tx, _ bool) (sdk.Context, error) {
		classicCalled = true
		return ctx, nil
	}
	auth := func(ctx sdk.Context, _ sdk.Tx, _ bool) (sdk.Context, error) {
		authCalled = true
		return ctx, nil
	}

	dec := smartaccountante.NewCircuitBreakerDecorator(s.App.SmartAccountKeeper, auth, classic)
	tx := &minimalTx{}

	_, err := dec.AnteHandle(s.Ctx, tx, false, func(ctx sdk.Context, _ sdk.Tx, _ bool) (sdk.Context, error) {
		return ctx, nil
	})
	s.Require().NoError(err)
	s.Require().True(classicCalled, "tx with no authenticator extension must take the classic path")
	s.Require().False(authCalled)
}

func (s *CircuitBreakerTestSuite) TestCircuitBreaker_RoutesToClassicWhenSmartAccountsInactive() {
	s.SetupTest()
	// Disable smart accounts module-wide via SetActiveState(false).
	s.App.SmartAccountKeeper.SetActiveState(s.Ctx, false)

	classicCalled := false
	classic := func(ctx sdk.Context, _ sdk.Tx, _ bool) (sdk.Context, error) {
		classicCalled = true
		return ctx, nil
	}
	auth := func(ctx sdk.Context, _ sdk.Tx, _ bool) (sdk.Context, error) {
		return ctx, errors.New("auth flow should NOT be reached when smart accounts are off")
	}

	dec := smartaccountante.NewCircuitBreakerDecorator(s.App.SmartAccountKeeper, auth, classic)
	tx := &minimalTx{}
	_, err := dec.AnteHandle(s.Ctx, tx, false, func(ctx sdk.Context, _ sdk.Tx, _ bool) (sdk.Context, error) {
		return ctx, nil
	})
	s.Require().NoError(err)
	s.Require().True(classicCalled)

	// Restore for cleanliness.
	s.App.SmartAccountKeeper.SetActiveState(s.Ctx, true)
}

// TestIsCircuitBreakActive_DirectAPI exercises the helper independent of
// the decorator wiring. With a non-extension tx and an active module, the
// circuit breaker should still be active because the extension is missing.
func (s *CircuitBreakerTestSuite) TestIsCircuitBreakActive_DirectAPI() {
	s.SetupTest()
	// Module active, but tx lacks extension → circuit breaker still active.
	active, opts := smartaccountante.IsCircuitBreakActive(s.Ctx, &minimalTx{}, s.App.SmartAccountKeeper)
	s.Require().True(active)
	s.Require().Nil(opts)
}

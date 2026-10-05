package claims_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/ixofoundation/ixo-blockchain/v7/app/apptesting"
	"github.com/ixofoundation/ixo-blockchain/v7/x/claims"
	"github.com/ixofoundation/ixo-blockchain/v7/x/claims/types"
)

type ABCITestSuite struct {
	apptesting.KeeperTestHelper
}

func TestABCITestSuite(t *testing.T) { suite.Run(t, new(ABCITestSuite)) }

func (s *ABCITestSuite) SetupTest() {
	s.Setup()
}

func (s *ABCITestSuite) TestEndBlockerRemovesExpiredIntentWhenCW1155RefundFails() {
	agent := apptesting.RandomAccountAddress()
	from := apptesting.RandomAccountAddress()
	escrow := apptesting.RandomAccountAddress()
	contract := apptesting.RandomAccountAddress()
	expiredAt := s.Ctx.BlockTime().Add(-time.Second)

	intent := types.Intent{
		Id:            "intent-1",
		AgentAddress:  agent.String(),
		CollectionId:  "collection-1",
		FromAddress:   from.String(),
		EscrowAddress: escrow.String(),
		ExpireAt:      &expiredAt,
		Cw1155Payment: []*types.CW1155Payment{{
			Address: contract.String(),
			Amount:  1,
		}},
		Cw1155IntentPayment: []*types.CW1155IntentPayment{{
			Address: contract.String(),
			Tokens: []*types.CW1155IntentPaymentToken{{
				TokenId: "token-1",
				Amount:  1,
			}},
		}},
	}
	s.App.ClaimsKeeper.SetIntent(s.Ctx, intent)

	s.Require().NotPanics(func() {
		claims.EndBlocker(s.Ctx, s.App.ClaimsKeeper)
	})

	_, err := s.App.ClaimsKeeper.GetIntent(s.Ctx, agent.String(), "collection-1", "intent-1")
	s.Require().ErrorIs(err, types.ErrIntentNotFound)
}

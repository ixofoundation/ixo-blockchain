package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"

	"github.com/ixofoundation/ixo-blockchain/v8/app/apptesting"
	"github.com/ixofoundation/ixo-blockchain/v8/x/smart-account/authenticator"
)

// TestAddAuthenticator_Then_Get exercises the full add → list → remove
// lifecycle for the SignatureVerification authenticator type, which is
// registered by the app at startup.
func (s *KeeperTestSuite) TestAddAuthenticator_LifeCycle() {
	s.SetupTest()
	acc := apptesting.RandomAccountAddress()

	// SignatureVerification accepts any 33-byte secp256k1 pubkey as config.
	pk := secp256k1.GenPrivKey().PubKey()
	config := pk.Bytes()

	id, err := s.App.SmartAccountKeeper.AddAuthenticator(s.Ctx, acc, authenticator.SignatureVerificationType, config)
	s.Require().NoError(err)
	s.Require().Greater(id, uint64(0))

	// Query: account now has one authenticator.
	auths, err := s.App.SmartAccountKeeper.GetAuthenticatorDataForAccount(s.Ctx, acc)
	s.Require().NoError(err)
	s.Require().Len(auths, 1)
	s.Require().Equal(authenticator.SignatureVerificationType, auths[0].Type)

	// Remove and re-query.
	s.Require().NoError(s.App.SmartAccountKeeper.RemoveAuthenticator(s.Ctx, acc, id))
	auths, err = s.App.SmartAccountKeeper.GetAuthenticatorDataForAccount(s.Ctx, acc)
	s.Require().NoError(err)
	s.Require().Len(auths, 0)
}

// TestAddAuthenticator_UnknownType: registering an authenticator type that
// the AuthenticatorManager doesn't know about returns an error.
func (s *KeeperTestSuite) TestAddAuthenticator_UnknownType() {
	s.SetupTest()
	acc := apptesting.RandomAccountAddress()
	_, err := s.App.SmartAccountKeeper.AddAuthenticator(s.Ctx, acc, "DefinitelyNotRegistered", nil)
	s.Require().ErrorContains(err, "is not registered")
}

// TestAuthenticatorIDCounter: each Add increments the global counter, and
// the next Add returns a fresh id even after a Remove.
func (s *KeeperTestSuite) TestAuthenticatorIDCounter_Monotonic() {
	s.SetupTest()
	acc := apptesting.RandomAccountAddress()
	pk := secp256k1.GenPrivKey().PubKey().Bytes()

	id1, err := s.App.SmartAccountKeeper.AddAuthenticator(s.Ctx, acc, authenticator.SignatureVerificationType, pk)
	s.Require().NoError(err)
	id2, err := s.App.SmartAccountKeeper.AddAuthenticator(s.Ctx, acc, authenticator.SignatureVerificationType, pk)
	s.Require().NoError(err)
	s.Require().Greater(id2, id1, "ids must be monotonic across the same account")

	s.Require().NoError(s.App.SmartAccountKeeper.RemoveAuthenticator(s.Ctx, acc, id1))
	id3, err := s.App.SmartAccountKeeper.AddAuthenticator(s.Ctx, acc, authenticator.SignatureVerificationType, pk)
	s.Require().NoError(err)
	s.Require().Greater(id3, id2, "removed ids are NOT reused")
}

// TestSetActiveState toggles the module-wide circuit-breaker flag.
func (s *KeeperTestSuite) TestSetActiveState() {
	s.SetupTest()
	s.App.SmartAccountKeeper.SetActiveState(s.Ctx, false)
	s.App.SmartAccountKeeper.SetActiveState(s.Ctx, true)
}

package keeper_test

import (
	"github.com/ixofoundation/ixo-blockchain/v6/x/token/types"
)

// PauseToken doesn't call into the wasm contract — it only flips the
// Paused flag on the stored Token. Easy to test in-process.
func (s *KeeperTestSuite) TestMsgPauseToken() {
	s.SetupTest()
	minter, t := s.seedToken("pausable")

	resp, err := s.msgServer.PauseToken(s.goCtx(), &types.MsgPauseToken{
		Minter:          minter.String(),
		ContractAddress: t.ContractAddress,
		Paused:          true,
	})
	s.Require().NoError(err)
	s.Require().NotNil(resp)
	got, _ := s.App.TokenKeeper.GetToken(s.Ctx, minter.String(), t.ContractAddress)
	s.Require().True(got.Paused)

	// unpause
	_, err = s.msgServer.PauseToken(s.goCtx(), &types.MsgPauseToken{
		Minter: minter.String(), ContractAddress: t.ContractAddress, Paused: false,
	})
	s.Require().NoError(err)
	got, _ = s.App.TokenKeeper.GetToken(s.Ctx, minter.String(), t.ContractAddress)
	s.Require().False(got.Paused)
}

func (s *KeeperTestSuite) TestMsgPauseToken_NonExistentToken() {
	s.SetupTest()
	_, err := s.msgServer.PauseToken(s.goCtx(), &types.MsgPauseToken{
		Minter: "ixo1deadbeef", ContractAddress: "ixo1nope",
	})
	s.Require().ErrorContains(err, "token not found")
}

// StopToken sets Stopped=true and is also wasm-free.
func (s *KeeperTestSuite) TestMsgStopToken() {
	s.SetupTest()
	minter, t := s.seedToken("stoppable")

	_, err := s.msgServer.StopToken(s.goCtx(), &types.MsgStopToken{
		Minter:          minter.String(),
		ContractAddress: t.ContractAddress,
	})
	s.Require().NoError(err)

	got, _ := s.App.TokenKeeper.GetToken(s.Ctx, minter.String(), t.ContractAddress)
	s.Require().True(got.Stopped)
}

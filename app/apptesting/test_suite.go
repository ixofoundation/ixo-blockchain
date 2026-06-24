// Package apptesting provides a reusable testify-suite base struct
// (KeeperTestHelper) that every ixo-blockchain unit / keeper test should embed.
//
// Pattern is modelled after Osmosis's app/apptesting and Juno's testutil
// packages: each test gets a fresh in-memory IxoApp, a default sdk.Context,
// pre-funded TestAccs, a gRPC query helper, and a small set of fixture helpers
// (FundAcc, SetupValidator, AssertEventEmitted, BeginNewBlock, etc.).
package apptesting

import (
	"fmt"
	"os"
	"time"

	coreheader "cosmossdk.io/core/header"
	storemetrics "cosmossdk.io/store/metrics"
	"cosmossdk.io/store/rootmulti"
	storetypes "cosmossdk.io/store/types"
	abci "github.com/cometbft/cometbft/abci/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/stretchr/testify/suite"

	"cosmossdk.io/log"

	"github.com/ixofoundation/ixo-blockchain/v8/app"
)

// KeeperTestHelper is meant to be embedded into per-module test suites.
//
//	type FooTestSuite struct {
//	    apptesting.KeeperTestHelper
//	}
//	func (s *FooTestSuite) SetupTest() { s.Setup() }
type KeeperTestHelper struct {
	suite.Suite

	App         *app.IxoApp
	Ctx         sdk.Context
	QueryHelper *baseapp.QueryServiceTestHelper
	TestAccs    []sdk.AccAddress
}

var (
	defaultTestStartTime = time.Now().UTC()
	baseTestAccts        = []sdk.AccAddress{}
)

func init() {
	baseTestAccts = CreateRandomAccounts(3)
}

// Setup spins up a fresh IxoApp on an in-memory DB and configures all default
// test fixtures (context, query helper, three pre-generated test accounts).
// Validator signing-info is set so that subsequent ABCI calls don't panic.
func (s *KeeperTestHelper) Setup() {
	dir, err := os.MkdirTemp("", "ixod-test-home")
	if err != nil {
		panic(fmt.Sprintf("failed creating temporary directory: %v", err))
	}
	s.T().Cleanup(func() { _ = os.RemoveAll(dir) })

	if app.IsDebugLogEnabled() {
		s.App = app.SetupWithCustomHome(false, dir, s.T())
	} else {
		s.App = app.SetupWithCustomHome(false, dir)
	}

	s.setupGeneral()

	// Manually set validator signing info, otherwise certain slashing-touching
	// flows (e.g. begin-blocker for downtime detection) panic on a missing key.
	vals, err := s.App.StakingKeeper.GetAllValidators(s.Ctx)
	s.Require().NoError(err)
	for _, val := range vals {
		consAddr, err := val.GetConsAddr()
		s.Require().NoError(err)
		signingInfo := slashingtypes.NewValidatorSigningInfo(
			sdk.ConsAddress(consAddr),
			s.Ctx.BlockHeight(),
			0,
			time.Unix(0, 0),
			false,
			0,
		)
		s.Require().NoError(s.App.SlashingKeeper.SetValidatorSigningInfo(s.Ctx, consAddr, signingInfo))
	}
}

func (s *KeeperTestHelper) setupGeneral() {
	s.Ctx = s.App.BaseApp.NewContextLegacy(false, cmtproto.Header{
		Height:  1,
		ChainID: app.TestChainID,
		Time:    defaultTestStartTime,
	})
	s.QueryHelper = &baseapp.QueryServiceTestHelper{
		GRPCQueryRouter: s.App.GRPCQueryRouter(),
		Ctx:             s.Ctx,
	}
	s.TestAccs = append([]sdk.AccAddress{}, baseTestAccts...)
}

// SetupTestForInitGenesis returns an app with InitChain skipped — useful when
// a test needs to assert pre-genesis state or call InitGenesis manually.
func (s *KeeperTestHelper) SetupTestForInitGenesis() {
	s.App = app.Setup(true)
	s.Ctx = s.App.BaseApp.NewContextLegacy(true, cmtproto.Header{ChainID: app.TestChainID})
}

// CreateTestContext builds a stand-alone sdk.Context backed by an empty multi
// store. Useful when testing pure helpers that don't need a full app.
func (s *KeeperTestHelper) CreateTestContext() sdk.Context {
	ctx, _ := s.CreateTestContextWithMultiStore()
	return ctx
}

// CreateTestContextWithMultiStore returns the context plus its multi store so
// callers can mount their own kv stores under known keys.
func (s *KeeperTestHelper) CreateTestContextWithMultiStore() (sdk.Context, storetypes.CommitMultiStore) {
	db := dbm.NewMemDB()
	logger := log.NewNopLogger()
	ms := rootmulti.NewStore(db, logger, storemetrics.NewNoOpMetrics())
	return sdk.NewContext(ms, cmtproto.Header{}, false, logger), ms
}

// Commit finalises the current block, commits the multi store, and advances
// the suite's context one block forward (height+1, time+1s).
func (s *KeeperTestHelper) Commit() {
	_, err := s.App.FinalizeBlock(&abci.RequestFinalizeBlock{
		Height: s.Ctx.BlockHeight(),
		Time:   s.Ctx.BlockTime(),
	})
	s.Require().NoError(err)
	_, err = s.App.Commit()
	s.Require().NoError(err)

	header := s.Ctx.BlockHeader()
	header.Time = s.Ctx.BlockTime().Add(time.Second)
	header.Height++

	s.Ctx = s.App.BaseApp.NewUncachedContext(false, header).WithHeaderInfo(coreheader.Info{
		Height: header.Height,
		Time:   header.Time,
	})
}

// BeginNewBlock advances the chain by one block, calling BeginBlocker.
func (s *KeeperTestHelper) BeginNewBlock() {
	header := cmtproto.Header{
		Height:  s.Ctx.BlockHeight() + 1,
		Time:    s.Ctx.BlockTime().Add(5 * time.Second),
		ChainID: app.TestChainID,
	}
	s.Ctx = s.Ctx.WithBlockHeight(header.Height).WithBlockTime(header.Time)
	_, err := s.App.BeginBlocker(s.Ctx)
	s.Require().NoError(err)
	s.Ctx = s.App.NewContextLegacy(false, header)
}

// EndBlock invokes the EndBlocker on the suite's current context.
func (s *KeeperTestHelper) EndBlock() {
	_, err := s.App.EndBlocker(s.Ctx)
	s.Require().NoError(err)
}

// RunMsg dispatches a sdk.Msg through the app's MsgServiceRouter — useful when
// you want to exercise the full message flow (validation, ante is skipped) but
// don't want to construct a transaction.
func (s *KeeperTestHelper) RunMsg(msg sdk.Msg) (*sdk.Result, error) {
	router := s.App.GetBaseApp().MsgServiceRouter()
	if handler := router.Handler(msg); handler != nil {
		return handler(s.Ctx, msg)
	}
	s.FailNow("msg %T could not be routed", msg)
	return nil, fmt.Errorf("no handler for msg %T", msg)
}

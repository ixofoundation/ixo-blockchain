package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/stretchr/testify/suite"

	"github.com/ixofoundation/ixo-blockchain/v6/app/apptesting"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v6/x/iid/types"
	"github.com/ixofoundation/ixo-blockchain/v6/x/names/keeper"
	"github.com/ixofoundation/ixo-blockchain/v6/x/names/types"
)

// KeeperTestSuite covers x/names keeper, msg_server, query server, and genesis
// behaviour against a real in-memory IxoApp. The pattern follows Osmosis's
// per-module suite (one suite struct per module, embed apptesting.KeeperTestHelper).
type KeeperTestSuite struct {
	apptesting.KeeperTestHelper

	msgServer   types.MsgServer
	queryClient types.QueryClient
	queryServer types.QueryServer
	authority   string
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (s *KeeperTestSuite) SetupTest() {
	s.Setup()
	s.msgServer = keeper.NewMsgServerImpl(s.App.NamesKeeper)
	s.queryServer = s.App.NamesKeeper

	// The IxoApp already registers x/names's QueryServer with the gRPC router
	// at module-init time, so we MUST NOT call types.RegisterQueryServer here:
	// it would panic with "service has already been registered". The
	// queryClient routes through s.QueryHelper, which wraps the app's router.
	s.queryClient = types.NewQueryClient(s.QueryHelper)

	// Authority is the gov module account, mirroring app keeper wiring.
	s.authority = authtypes.NewModuleAddress(govtypes.ModuleName).String()
}

// ----------------------------------------------------------------------------
// Test fixtures
// ----------------------------------------------------------------------------

// goCtx returns a context.Context view of the suite's sdk.Context. The cosmos
// sdk's Context already satisfies context.Context, so this is a simple cast.
func (s *KeeperTestSuite) goCtx() context.Context { return s.Ctx }

// addrAndDID returns a fresh bech32 address plus a DID document already
// persisted in the iid keeper such that the address controls the DID.
func (s *KeeperTestSuite) addrAndDID() (sdk.AccAddress, string) {
	addr := apptesting.RandomAccountAddress()
	did := "did:ixo:" + addr.String()
	doc := iidtypes.IidDocument{
		Id:         did,
		Controller: []string{addr.String()},
	}
	s.App.IidKeeper.SetDidDocument(s.Ctx, []byte(did), doc)
	return addr, did
}

// onlyDID stores a DID document not controlled by the signer — used to
// exercise authorisation failures.
func (s *KeeperTestSuite) onlyDID() string {
	other := apptesting.RandomAccountAddress()
	did := "did:ixo:" + other.String()
	doc := iidtypes.IidDocument{
		Id:         did,
		Controller: []string{other.String()},
	}
	s.App.IidKeeper.SetDidDocument(s.Ctx, []byte(did), doc)
	return did
}

// defaultNamespace returns a Namespace ready to be passed to CreateNamespace.
// Self-registration is enabled, no registrars, generous length bounds.
func (s *KeeperTestSuite) defaultNamespace(name string) types.Namespace {
	return types.Namespace{
		Name:              name,
		Description:       "test namespace",
		AllowSelfRegister: true,
		MinLength:         3,
		MaxLength:         32,
	}
}

// registrarNamespace returns a Namespace where only the supplied registrar
// addresses can register/transfer/etc. Self-register is off.
func (s *KeeperTestSuite) registrarNamespace(name string, registrars ...string) types.Namespace {
	return types.Namespace{
		Name:                   name,
		Description:            "registrar namespace",
		AllowSelfRegister:      false,
		AllowRegistrarOverride: true,
		RegistrarAccounts:      registrars,
		MinLength:              3,
		MaxLength:              32,
	}
}

// mustCreateNamespace stores a namespace via SetNamespace (bypassing
// authority checks). Helpful for arranging fixtures inside tests.
func (s *KeeperTestSuite) mustCreateNamespace(ns types.Namespace) {
	s.App.NamesKeeper.SetNamespace(s.Ctx, ns)
}

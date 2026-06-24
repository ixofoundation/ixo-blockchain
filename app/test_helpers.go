package app

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"cosmossdk.io/log"
	sdkmath "cosmossdk.io/math"
	pruningtypes "cosmossdk.io/store/pruning/types"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/crypto/secp256k1"
	tmtypes "github.com/cometbft/cometbft/types"
	cosmosdb "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/baseapp"
	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client/flags"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cosmoserver "github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/testutil/mock"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sims "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module/testutil"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/ixofoundation/ixo-blockchain/v8/lib/ixo"
)

const TestChainID = "ixo-test-1"

// NewTestNetworkFixture returns a new ixo simapp AppConstructor for network simulation tests
func NewTestNetworkFixture() network.TestFixture {
	dir, err := os.MkdirTemp("", DefaultNodeHome)
	if err != nil {
		panic(fmt.Sprintf("failed creating temporary directory: %v", err))
	}
	defer os.RemoveAll(dir)

	app := NewIxoApp(log.NewNopLogger(), cosmosdb.NewMemDB(), nil, true, map[int64]bool{}, DefaultNodeHome, sims.EmptyAppOptions{}, EmptyWasmOpts)

	appCtr := func(val network.ValidatorI) servertypes.Application {
		return NewIxoApp(
			val.GetCtx().Logger, cosmosdb.NewMemDB(), nil, true,
			map[int64]bool{}, DefaultNodeHome, sims.EmptyAppOptions{}, EmptyWasmOpts,
			bam.SetPruning(pruningtypes.NewPruningOptionsFromString(val.GetAppConfig().Pruning)),
			bam.SetMinGasPrices(val.GetAppConfig().MinGasPrices),
			bam.SetChainID(val.GetCtx().Viper.GetString(flags.FlagChainID)),
		)
	}

	return network.TestFixture{
		AppConstructor: appCtr,
		GenesisState:   app.BasicModuleManager.DefaultGenesis(app.appCodec),
		EncodingConfig: testutil.TestEncodingConfig{
			InterfaceRegistry: app.InterfaceRegistry(),
			Codec:             app.AppCodec(),
			TxConfig:          app.TxConfig(),
			Amino:             app.LegacyAmino(),
		},
	}
}

// debugAppOptions implements servertypes.AppOptions and forces trace mode on.
type debugAppOptions struct{}

func (debugAppOptions) Get(o string) any {
	if o == cosmoserver.FlagTrace {
		return true
	}
	return nil
}

// IsDebugLogEnabled returns true when the IXO_KEEPER_DEBUG env var is set.
// When enabled, app.Setup uses a logger that prints to stdout and a tracing
// app-options shim. Useful when chasing keeper bugs in tests.
func IsDebugLogEnabled() bool {
	return os.Getenv("IXO_KEEPER_DEBUG") != ""
}

// Setup initialises a new IxoApp with an in-memory DB and a single bonded
// validator. Use this in test suites instead of constructing NewIxoApp by hand.
func Setup(isCheckTx bool) *IxoApp {
	return SetupWithCustomHome(isCheckTx, DefaultNodeHome)
}

// SetupWithCustomHome behaves like Setup but lets the caller pin the home dir.
func SetupWithCustomHome(isCheckTx bool, dir string, t ...*testing.T) *IxoApp {
	return SetupWithCustomHomeAndChainID(isCheckTx, dir, TestChainID, t...)
}

// SetupWithCustomHomeAndChainID is the most-flexible test setup helper.
func SetupWithCustomHomeAndChainID(isCheckTx bool, dir, chainID string, t ...*testing.T) *IxoApp {
	db := cosmosdb.NewMemDB()

	var (
		l       log.Logger
		appOpts servertypes.AppOptions
	)
	if IsDebugLogEnabled() {
		appOpts = debugAppOptions{}
	} else {
		appOpts = sims.EmptyAppOptions{}
	}
	if len(t) > 0 {
		l = log.NewTestLogger(t[0])
	} else {
		l = log.NewNopLogger()
	}

	app := NewIxoApp(
		l,
		db,
		nil,
		true,
		map[int64]bool{},
		dir,
		appOpts,
		EmptyWasmOpts,
		baseapp.SetChainID(chainID),
	)

	if !isCheckTx {
		genesisState := GenesisStateWithValSet(app)
		stateBytes, err := json.Marshal(genesisState)
		if err != nil {
			panic(err)
		}

		_, err = app.InitChain(
			&abci.RequestInitChain{
				Validators:      []abci.ValidatorUpdate{},
				ConsensusParams: sims.DefaultConsensusParams,
				AppStateBytes:   stateBytes,
				ChainId:         chainID,
				Time:            time.Now().UTC(),
			},
		)
		if err != nil {
			panic(err)
		}
	}

	return app
}

// GenesisStateWithValSet builds a default genesis state with a single
// pre-bonded validator and a single funded delegator account. The bond denom
// is `uixo` (ixo.IxoNativeToken) — staking and bank state are populated to
// satisfy invariants on InitChain.
func GenesisStateWithValSet(app *IxoApp) GenesisState {
	privVal := mock.NewPV()
	pubKey, err := privVal.GetPubKey()
	if err != nil {
		panic(err)
	}
	validator := tmtypes.NewValidator(pubKey, 1)
	valSet := tmtypes.NewValidatorSet([]*tmtypes.Validator{validator})

	senderPrivKey := secp256k1.GenPrivKey()
	acc := authtypes.NewBaseAccountWithAddress(senderPrivKey.PubKey().Address().Bytes())

	genesisState := app.DefaultGenesis()

	// auth: register the genesis account
	authGenesis := authtypes.NewGenesisState(authtypes.DefaultParams(), []authtypes.GenesisAccount{acc})
	genesisState[authtypes.ModuleName] = app.appCodec.MustMarshalJSON(authGenesis)

	// staking: register the validator with self-delegation
	bondDenom := ixo.IxoNativeToken
	bondAmt := sdk.DefaultPowerReduction
	validators := make([]stakingtypes.Validator, 0, len(valSet.Validators))
	delegations := make([]stakingtypes.Delegation, 0, len(valSet.Validators))
	for _, val := range valSet.Validators {
		pk, err := cryptocodec.FromCmtPubKeyInterface(val.PubKey)
		if err != nil {
			panic(err)
		}
		pkAny, err := codectypes.NewAnyWithValue(pk)
		if err != nil {
			panic(err)
		}
		v := stakingtypes.Validator{
			OperatorAddress:   sdk.ValAddress(val.Address).String(),
			ConsensusPubkey:   pkAny,
			Jailed:            false,
			Status:            stakingtypes.Bonded,
			Tokens:            bondAmt,
			DelegatorShares:   sdkmath.LegacyOneDec(),
			Description:       stakingtypes.Description{},
			UnbondingHeight:   int64(0),
			UnbondingTime:     time.Unix(0, 0).UTC(),
			Commission:        stakingtypes.NewCommission(sdkmath.LegacyZeroDec(), sdkmath.LegacyZeroDec(), sdkmath.LegacyZeroDec()),
			MinSelfDelegation: sdkmath.ZeroInt(),
		}
		validators = append(validators, v)
		delegations = append(delegations, stakingtypes.NewDelegation(acc.GetAddress().String(), sdk.ValAddress(val.Address).String(), sdkmath.LegacyOneDec()))
	}

	stakingParams := stakingtypes.DefaultParams()
	stakingParams.BondDenom = bondDenom
	stakingGenesis := stakingtypes.NewGenesisState(stakingParams, validators, delegations)
	genesisState[stakingtypes.ModuleName] = app.appCodec.MustMarshalJSON(stakingGenesis)

	// bank: add the bonded pool balance and total supply for the bond denom
	balances := []banktypes.Balance{
		{
			Address: authtypes.NewModuleAddress(stakingtypes.BondedPoolName).String(),
			Coins:   sdk.Coins{sdk.NewCoin(bondDenom, bondAmt)},
		},
	}
	totalSupply := sdk.NewCoins(sdk.NewCoin(bondDenom, bondAmt))
	bankGenesis := banktypes.NewGenesisState(
		banktypes.DefaultGenesisState().Params,
		balances,
		totalSupply,
		[]banktypes.Metadata{},
		[]banktypes.SendEnabled{},
	)
	genesisState[banktypes.ModuleName] = app.appCodec.MustMarshalJSON(bankGenesis)

	return genesisState
}

package app

import (
	"fmt"
	"os"

	"cosmossdk.io/log"
	pruningtypes "cosmossdk.io/store/pruning/types"
	cosmosdb "github.com/cosmos/cosmos-db"
	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client/flags"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sims "github.com/cosmos/cosmos-sdk/testutil/sims"
	"github.com/cosmos/cosmos-sdk/types/module/testutil"
)

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

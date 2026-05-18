package names

import (
	"context"
	"encoding/json"
	"fmt"

	"cosmossdk.io/core/appmodule"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"github.com/ixofoundation/ixo-blockchain/v6/x/names/keeper"
	"github.com/ixofoundation/ixo-blockchain/v6/x/names/simulation"
	"github.com/ixofoundation/ixo-blockchain/v6/x/names/types"
)

var (
	_ module.AppModuleBasic   = AppModuleBasic{}
	_ module.HasGenesisBasics = AppModuleBasic{}

	_ appmodule.AppModule        = AppModule{}
	_ module.HasConsensusVersion = AppModule{}
	_ module.HasGenesis          = AppModule{}
	_ module.HasServices         = AppModule{}
)

// AppModuleBasic is the standalone basic-info module surface.
type AppModuleBasic struct{}

func (AppModuleBasic) Name() string { return types.ModuleName }

func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterLegacyAminoCodec(cdc)
}

func (AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(reg)
}

func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesisState())
}

func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	var gs types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &gs); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}
	return gs.Validate()
}

func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	if err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx)); err != nil {
		return
	}
}

// GetTxCmd returns nil; CLI is provided via autocli.
func (AppModuleBasic) GetTxCmd() *cobra.Command { return nil }

// GetQueryCmd returns nil; CLI is provided via autocli.
func (AppModuleBasic) GetQueryCmd() *cobra.Command { return nil }

// AppModule wires the names module into the app.
type AppModule struct {
	AppModuleBasic
	keeper        keeper.Keeper
	accountKeeper simulation.AccountKeeper
	bankKeeper    simulation.BankKeeper
}

func NewAppModule(k keeper.Keeper, ak simulation.AccountKeeper, bk simulation.BankKeeper) AppModule {
	return AppModule{keeper: k, accountKeeper: ak, bankKeeper: bk}
}

func (AppModule) IsAppModule() {}

func (AppModule) IsOnePerModuleType() {}

func (am AppModule) Name() string { return am.AppModuleBasic.Name() }

func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

func (AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) {
	var gs types.GenesisState
	cdc.MustUnmarshalJSON(data, &gs)
	am.keeper.InitGenesis(ctx, gs)
}

func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(am.keeper.ExportGenesis(ctx))
}

func (AppModule) ConsensusVersion() uint64 { return 1 }

// AppModuleSimulation: registers the store decoder and exposes empty
// proposal/message simulation hooks. The names module is governance-driven
// (CreateNamespace / UpdateNamespace are authority-only) so per-Msg
// WeightedOperations are intentionally empty here; self-register flows are
// covered by keeper-level tests.

// ProposalContents returns no governance proposal content for x/names.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent { //nolint:staticcheck
	return nil
}

// ProposalMsgs returns no governance proposal messages for x/names.
func (AppModule) ProposalMsgs(_ module.SimulationState) []simtypes.WeightedProposalMsg {
	return nil
}

// RegisterStoreDecoder wires the store-determinism decoder into the simulator.
func (AppModule) RegisterStoreDecoder(sdr simtypes.StoreDecoderRegistry) {
	sdr[types.StoreKey] = simulation.NewDecodeStore(nil)
}

// WeightedOperations returns the per-Msg simulation operations for x/names.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	return simulation.WeightedOperations(simState.AppParams, simState.Cdc, simState.TxConfig, am.accountKeeper, am.bankKeeper, am.keeper)
}

// GenerateGenesisState seeds the SimulationState with the module's default
// genesis. Property tests can later replace this with randomized state.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(types.DefaultGenesisState())
}

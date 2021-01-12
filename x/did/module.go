package did

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/gov/types"

	//"github.com/cosmos/cosmos-sdk/client/flags"
	//"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	//"github.com/ixofoundation/ixo-blockchain/x/did/internal/types"

	"github.com/cosmos/cosmos-sdk/client"
	//"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/ixofoundation/ixo-blockchain/x/did/client/cli"
	"github.com/ixofoundation/ixo-blockchain/x/did/client/rest"
	"github.com/ixofoundation/ixo-blockchain/x/did/internal/keeper"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModuleBasic defines the basic application module used by the did module.
type AppModuleBasic struct{}

// Name returns the did module's name.
func (AppModuleBasic) Name() string {
	return ModuleName
}

// RegisterLegacyAminoCodec registers the did module's types for the given codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	RegisterLegacyAminoCodec(cdc)
}

// RegisterInterfaces registers interfaces and implementations of the did module.
func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

//func (AppModuleBasic) DefaultGenesis() json.RawMessage {
//	return ModuleCdc.MustMarshalJSON(DefaultGenesisState())
//}

// DefaultGenesis returns default genesis state as raw bytes for the did
// module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONMarshaler) json.RawMessage {
	//return ModuleCdc.MustMarshalJSON(DefaultGenesisState())
	return cdc.MustMarshalJSON(DefaultGenesisState())
}

//func (AppModuleBasic) ValidateGenesis(bz json.RawMessage) error {
//	var data GenesisState
//	err := ModuleCdc.UnmarshalJSON(bz, &data)
//	if err != nil {
//		return err
//	}
//	return ValidateGenesis(data)
//}

// ValidateGenesis performs genesis state validation for the did module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONMarshaler, config client.TxEncodingConfig, bz json.RawMessage) error {
	var data GenesisState
	err := cdc.UnmarshalJSON(bz, &data) //ModuleCdc.UnmarshalJSON(bz, &data)
	if err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", ModuleName, err)
	}

	return ValidateGenesis(data)
}

// RegisterRESTRoutes registers the REST routes for the did module.
func (AppModuleBasic) RegisterRESTRoutes(clientCtx client.Context, rtr *mux.Router) {
	rest.RegisterRoutes(clientCtx, rtr)
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the did module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx)) //tx.RegisterGRPCGatewayRoutes()
}

// GetTxCmd returns the root tx command for the did module.
func (AppModuleBasic) GetTxCmd(/*cdc *codec.Codec*/) *cobra.Command {
	didTxCmd := &cobra.Command{
		Use:                        ModuleName,
		Short:                      "did transaction sub commands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	didTxCmd.AddCommand(
		cli.GetCmdAddDidDoc(),
		cli.GetCmdAddCredential(),
	)

	//didTxCmd.AddCommand(flags.PostCommands(
	//	cli.GetCmdAddDidDoc(cdc)),
	//	cli.GetCmdAddCredential(cdc),
	//)...)

	return didTxCmd
}

// GetQueryCmd returns the root query command for the did module.
func (AppModuleBasic) GetQueryCmd(/*cdc *codec.Codec*/) *cobra.Command {
	didQueryCmd := &cobra.Command{
		Use:                        ModuleName,
		Short:                      "did query sub commands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	didQueryCmd.AddCommand(
		cli.GetCmdAddressFromBase58Pubkey(),
		cli.GetCmdAddressFromDid(),
		cli.GetCmdIxoDidFromMnemonic(),
		cli.GetCmdDidDoc(),
		cli.GetCmdAllDids(),
		cli.GetCmdAllDidDocs(),
	)

	//didQueryCmd.AddCommand(flags.GetCommands(
	//	cli.GetCmdAddressFromBase58Pubkey(),
	//	cli.GetCmdAddressFromDid(cdc),
	//	cli.GetCmdIxoDidFromMnemonic(),
	//	cli.GetCmdDidDoc(cdc),
	//	cli.GetCmdAllDids(cdc),
	//	cli.GetCmdAllDidDocs(cdc),
	//)...)

	return didQueryCmd
}
//____________________________________________________________________________

// AppModule implements an application module for the did module.
type AppModule struct {
	AppModuleBasic

	keeper keeper.Keeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(keeper Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         keeper,
	}
}

// Name returns the did module's name.
func (AppModule) Name() string {
	return ModuleName
}

// TODO what is this?
// RegisterInvariants performs a no-op.
func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// Route returns the message routing key for the did module.
func (am AppModule) Route() sdk.Route {
	return sdk.NewRoute(RouterKey, NewHandler(am.keeper)) //RouterKey
}

//func (am AppModule) NewHandler() sdk.Handler {
//	return NewHandler(am.keeper)
//}

// QuerierRoute returns the did module's querier route name.
func (AppModule) QuerierRoute() string {
	return QuerierRoute
}

// LegacyQuerierHandler returns the did module sdk.Querier.
func (am AppModule) LegacyQuerierHandler(cdc *codec.LegacyAmino) sdk.Querier {
	return keeper.NewQuerier(am.keeper, cdc)
}

// RegisterServices registers a GRPC query service to respond to the
// module-specific GRPC queries.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	//types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
	//types.RegisterMsgServer(cfg.MsgServer(), am.keeper)
	// TODO what does this do?
}

//func (am AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
//	var genesisState GenesisState
//	ModuleCdc.MustUnmarshalJSON(data, &genesisState)
//	InitGenesis(ctx, am.keeper, genesisState)
//	return []abci.ValidatorUpdate{}
//}

// InitGenesis performs genesis initialization for the did module. It returns
// no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONMarshaler, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState GenesisState
	cdc.MustUnmarshalJSON(data, &genesisState)
	InitGenesis(ctx, am.keeper, genesisState)
	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the exported genesis state as raw bytes for the did
// module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONMarshaler) json.RawMessage {
	gs := ExportGenesis(ctx, am.keeper)
	return cdc.MustMarshalJSON(gs) //ModuleCdc.MustMarshalJSON(gs)
}

// BeginBlock returns the begin blocker for the did module.
func (am AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {
}

// EndBlock returns the end blocker for the did module. It returns no validator
// updates.
func (AppModule) EndBlock(_ sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

package payments

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client/flags"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/ixofoundation/ixo-blockchain/x/payments/client/cli"

	//"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/ixofoundation/ixo-blockchain/x/payments/client/rest"
	"github.com/ixofoundation/ixo-blockchain/x/payments/internal/keeper"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

type AppModuleBasic struct{}

func (AppModuleBasic) Name() string {
	return ModuleName
}

func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	RegisterCodec(cdc)
}

func (AppModuleBasic) RegisterInterfaces(codectypes.InterfaceRegistry) {}

func (AppModuleBasic) DefaultGenesis(marshaler codec.JSONMarshaler) json.RawMessage {
	return ModuleCdc.MustMarshalJSON(DefaultGenesisState())
}

func (AppModuleBasic) ValidateGenesis(m codec.JSONMarshaler, enc client.TxEncodingConfig, bz json.RawMessage) error {
	var data GenesisState
	err := ModuleCdc.UnmarshalJSON(bz, &data)
	if err != nil {
		return err
	}
	return ValidateGenesis(data)
}

func (AppModuleBasic) RegisterRESTRoutes(ctx client.Context, rtr *mux.Router) {
	rest.RegisterRoutes(ctx, rtr)
}

func (AppModuleBasic) RegisterGRPCGatewayRoutes(client.Context, *runtime.ServeMux) {}

func (AppModuleBasic) GetTxCmd(/*cdc *codec.Codec*/) *cobra.Command {
	//paymentsTxCmd := &cobra.Command{
	//	Use:                        ModuleName,
	//	Short:                      "payments transaction sub commands",
	//	DisableFlagParsing:         true,
	//	SuggestionsMinimumDistance: 2,
	//	RunE:                       client.ValidateCmd,
	//}
	//
	//paymentsTxCmd.AddCommand(flags.PostCommands(
	//	cli.GetCmdCreatePaymentTemplate(cdc),
	//	cli.GetCmdCreatePaymentContract(cdc),
	//	cli.GetCmdCreateSubscription(cdc),
	//	cli.GetCmdSetPaymentContractAuthorisation(cdc),
	//	cli.GetCmdGrantPaymentDiscount(cdc),
	//	cli.GetCmdRevokePaymentDiscount(cdc),
	//	cli.GetCmdEffectPayment(cdc),
	//)...)
	//
	//return paymentsTxCmd
}

func (AppModuleBasic) GetQueryCmd(/*cdc *codec.Codec*/) *cobra.Command {
	//paymentsQueryCmd := &cobra.Command{
	//	Use:                        ModuleName,
	//	Short:                      "payments query sub commands",
	//	DisableFlagParsing:         true,
	//	SuggestionsMinimumDistance: 2,
	//	RunE:                       client.ValidateCmd,
	//}
	//
	//paymentsQueryCmd.AddCommand(flags.GetCommands(
	//	cli.GetCmdPaymentTemplate(cdc),
	//	cli.GetCmdPaymentContract(cdc),
	//	cli.GetCmdSubscription(cdc),
	//)...)
	//
	//return paymentsQueryCmd
}

type AppModule struct {
	AppModuleBasic
	keeper     keeper.Keeper
	bankKeeper bank.Keeper
}

func NewAppModule(keeper Keeper, bankKeeper bank.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         keeper,
		bankKeeper:     bankKeeper,
	}
}

func (AppModule) Name() string {
	return ModuleName
}

func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {
}

// TODO Populate
func (am AppModule) Route() sdk.Route {
	return sdk.Route{} //RouterKey
}

func (am AppModule) NewHandler() sdk.Handler {
	return NewHandler(am.keeper, am.bankKeeper)
}

func (AppModule) QuerierRoute() string {
	return QuerierRoute
}

func (am AppModule) LegacyQuerierHandler(*codec.LegacyAmino) sdk.Querier {
	return NewQuerier(am.keeper)
}

func (am AppModule) RegisterServices(module.Configurator) {}

func (am AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	//var genesisState GenesisState
	//ModuleCdc.MustUnmarshalJSON(data, &genesisState)
	//InitGenesis(ctx, am.keeper, genesisState)
	return []abci.ValidatorUpdate{}
}

func (am AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	gs := ExportGenesis(ctx, am.keeper)
	return ModuleCdc.MustMarshalJSON(gs)
}

func (am AppModule) BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock) {
}

func (am AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return EndBlocker(ctx, am.keeper)
}

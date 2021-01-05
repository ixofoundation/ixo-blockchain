package project
//
//import (
//	"encoding/json"
//	"github.com/cosmos/cosmos-sdk/client/flags"
//
//	"github.com/cosmos/cosmos-sdk/client"
//	//"github.com/cosmos/cosmos-sdk/client/context"
//	"github.com/cosmos/cosmos-sdk/codec"
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	"github.com/cosmos/cosmos-sdk/types/module"
//	"github.com/cosmos/cosmos-sdk/x/bank"
//	"github.com/gorilla/mux"
//	"github.com/spf13/cobra"
//	abci "github.com/tendermint/tendermint/abci/types"
//
//	"github.com/ixofoundation/ixo-blockchain/x/payments"
//	"github.com/ixofoundation/ixo-blockchain/x/project/client/cli"
//	"github.com/ixofoundation/ixo-blockchain/x/project/client/rest"
//	"github.com/ixofoundation/ixo-blockchain/x/project/internal/keeper"
//)
//
//var (
//	_ module.AppModule      = AppModule{}
//	_ module.AppModuleBasic = AppModuleBasic{}
//)
//
//type AppModuleBasic struct{}
//
//func (AppModuleBasic) Name() string {
//	return ModuleName
//}
//
//func (AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
//	RegisterCodec(cdc)
//}
//
//func (AppModuleBasic) DefaultGenesis() json.RawMessage {
//	return ModuleCdc.MustMarshalJSON(DefaultGenesisState())
//}
//
//func (AppModuleBasic) ValidateGenesis(bz json.RawMessage) error {
//	var data GenesisState
//	err := ModuleCdc.UnmarshalJSON(bz, &data)
//	if err != nil {
//		return err
//	}
//	return ValidateGenesis(data)
//}
//
//func (AppModuleBasic) RegisterRESTRoutes(ctx context.CLIContext, rtr *mux.Router) {
//	rest.RegisterRoutes(ctx, rtr)
//}
//
//func (AppModuleBasic) GetTxCmd(cdc *codec.Codec) *cobra.Command {
//	projectTxCmd := &cobra.Command{
//		Use:                        ModuleName,
//		Short:                      "project transaction sub commands",
//		DisableFlagParsing:         true,
//		SuggestionsMinimumDistance: 2,
//		RunE:                       client.ValidateCmd,
//	}
//
//	projectTxCmd.AddCommand(flags.PostCommands(
//		cli.GetCmdCreateProject(cdc),
//		cli.GetCmdCreateAgent(cdc),
//		cli.GetCmdUpdateProjectStatus(cdc),
//		cli.GetCmdUpdateAgent(cdc),
//		cli.GetCmdCreateClaim(cdc),
//		cli.GetCmdCreateEvaluation(cdc),
//		cli.GetCmdWithdrawFunds(cdc),
//	)...)
//
//	return projectTxCmd
//}
//
//func (AppModuleBasic) GetQueryCmd(cdc *codec.Codec) *cobra.Command {
//	projectQueryCmd := &cobra.Command{
//		Use:                        ModuleName,
//		Short:                      "project query sub commands",
//		DisableFlagParsing:         true,
//		SuggestionsMinimumDistance: 2,
//		RunE:                       client.ValidateCmd,
//	}
//
//	projectQueryCmd.AddCommand(flags.GetCommands(
//		cli.GetCmdProjectDoc(cdc),
//		cli.GetCmdProjectAccounts(cdc),
//		cli.GetCmdProjectTxs(cdc),
//		cli.GetParamsRequestHandler(cdc),
//	)...)
//
//	return projectQueryCmd
//}
//
//type AppModule struct {
//	AppModuleBasic
//	keeper         keeper.Keeper
//	paymentsKeeper payments.Keeper
//	bankKeeper     bank.Keeper
//}
//
//func NewAppModule(keeper Keeper, paymentsKeeper payments.Keeper,
//	bankKeeper bank.Keeper) AppModule {
//
//	return AppModule{
//		AppModuleBasic: AppModuleBasic{},
//		keeper:         keeper,
//		paymentsKeeper: paymentsKeeper,
//		bankKeeper:     bankKeeper,
//	}
//}
//
//func (AppModule) Name() string {
//	return ModuleName
//}
//
//func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {}
//
//func (AppModule) Route() string {
//	return RouterKey
//}
//
//func (am AppModule) NewHandler() sdk.Handler {
//	return NewHandler(am.keeper, am.paymentsKeeper, am.bankKeeper)
//}
//
//func (AppModule) QuerierRoute() string {
//	return QuerierRoute
//}
//
//func (am AppModule) NewQuerierHandler() sdk.Querier {
//	return keeper.NewQuerier(am.keeper)
//}
//
//func (am AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
//	var genesisState GenesisState
//	ModuleCdc.MustUnmarshalJSON(data, &genesisState)
//	InitGenesis(ctx, am.keeper, genesisState)
//	return []abci.ValidatorUpdate{}
//}
//
//func (am AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
//	gs := ExportGenesis(ctx, am.keeper)
//	return ModuleCdc.MustMarshalJSON(gs)
//}
//
//func (am AppModule) BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock) {
//}
//
//func (AppModule) EndBlock(_ sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
//	return []abci.ValidatorUpdate{}
//}

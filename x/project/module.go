package project

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	abciTypes "github.com/tendermint/tendermint/abci/types"

	"github.com/ixofoundation/ixo-cosmos/x/contracts"
	"github.com/ixofoundation/ixo-cosmos/x/fees"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
	"github.com/ixofoundation/ixo-cosmos/x/params"
	"github.com/ixofoundation/ixo-cosmos/x/project/client/cli"
	"github.com/ixofoundation/ixo-cosmos/x/project/client/rest"
	"github.com/ixofoundation/ixo-cosmos/x/project/internal/keeper"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

type AppModuleBasic struct{}

func (AppModuleBasic) Name() string {
	return ModuleName
}

func (AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
	Registercodec(cdc)
}

func (AppModuleBasic) DefaultGenesis() json.RawMessage {
	return nil
}

func (AppModuleBasic) ValidateGenesis(bz json.RawMessage) error {
	return nil
}

func (AppModuleBasic) RegisterRESTRoutes(ctx context.CLIContext, rtr *mux.Router) {
	rest.RegisterRoutes(ctx, rtr)
}

func (AppModuleBasic) GetTxCmd(cdc *codec.Codec) *cobra.Command {
	projectTxCmd := &cobra.Command{
		Use:                        ModuleName,
		Short:                      "Project transaction sub commands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	projectTxCmd.AddCommand(client.PostCommands(
		cli.CreateProjectCmd(cdc),
		cli.CreateAgentCmd(cdc),
		cli.UpdateProjectStatusCmd(cdc),
		cli.UpdateAgentCmd(cdc),
		cli.CreateClaimCmd(cdc),
		cli.CreateEvaluationCmd(cdc),
		cli.WithDrawFundsCmd(cdc),
	)...)

	return projectTxCmd
}

func (AppModuleBasic) GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	projectQueryCmd := &cobra.Command{
		Use:                        ModuleName,
		Short:                      "project query sub commands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	projectQueryCmd.AddCommand(client.GetCommands(
		cli.GetProjectDocCmd(cdc),
		cli.GetProjectAccountsCmd(cdc),
		cli.GetProjectTxsCmd(cdc),
	)...)

	return projectQueryCmd
}

type AppModule struct {
	AppModuleBasic
	keeper         keeper.Keeper
	feesKeeper     fees.Keeper
	contractKeeper contracts.Keeper
	bankKeeper     bank.Keeper
	paramsKeeper   params.Keeper
	ethClient      ixo.EthClient
}

func NewAppModule(keeper Keeper, feesKeeper fees.Keeper, contractKeeper contracts.Keeper,
	bankKeeper bank.Keeper, paramsKeeper params.Keeper, ethClient ixo.EthClient) AppModule {

	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         keeper,
		feesKeeper:     feesKeeper,
		contractKeeper: contractKeeper,
		bankKeeper:     bankKeeper,
		paramsKeeper:   paramsKeeper,
		ethClient:      ethClient,
	}
}

func (AppModule) Name() string {
	return ModuleName
}

func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {}

func (AppModule) Route() string {
	return RouterKey
}

func (am AppModule) NewHandler() sdk.Handler {
	return NewHandler(am.keeper, am.feesKeeper, am.contractKeeper, am.bankKeeper, am.paramsKeeper, am.ethClient)
}

func (AppModule) QuerierRoute() string {
	return QuerierRoute
}

func (am AppModule) NewQuerierHandler() sdk.Querier {
	return keeper.NewQuerier(am.keeper)
}

func (am AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abciTypes.ValidatorUpdate {
	return []abciTypes.ValidatorUpdate{}
}

func (am AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	return nil
}

func (am AppModule) BeginBlock(ctx sdk.Context, req abciTypes.RequestBeginBlock) {
}

func (AppModule) EndBlock(_ sdk.Context, _ abciTypes.RequestEndBlock) []abciTypes.ValidatorUpdate {
	return []abciTypes.ValidatorUpdate{}
}

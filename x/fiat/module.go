package fiat

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	abciTypes "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/ixofoundation/ixo-cosmos/codec"
	"github.com/ixofoundation/ixo-cosmos/x/fiat/client/cli"
	"github.com/ixofoundation/ixo-cosmos/x/fiat/client/rest"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

type AppModuleBasic struct{}

func (AppModuleBasic) Name() string {
	return ModuleName
}

func (AppModuleBasic) RegisterCodec(cdc *codec.Codec) { RegisterCodec(cdc) }

func (AppModuleBasic) DefaultGenesis() json.RawMessage {
	return ModuleCdc.MustMarshalJSON(DefaultGenesisState())
}

func (AppModuleBasic) ValidateGenesis(bz json.RawMessage) error {
	var data GenesisState
	err := ModuleCdc.UnmarshalJSON(bz, &data)
	if err != nil {
		return err
	}
	return ValidateGenesis(data)
}

func (AppModuleBasic) RegisterRESTRoutes(ctx context.CLIContext, rtr *mux.Router) {
	rest.RegisterRoutes(ctx, rtr)
}

func (AppModuleBasic) GetTxCmd(cdc *codec.Codec) *cobra.Command {
	fiatTxCmd := &cobra.Command{
		Use:                        ModuleName,
		Short:                      "fiat transaction sub commands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	fiatTxCmd.AddCommand(client.PostCommands(
		cli.IssueFiatCmd(cdc),
		cli.RedeemFiatCmd(cdc),
		cli.SendFiatCmd(cdc),
	)...)

	return fiatTxCmd
}

func (AppModuleBasic) GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	fiatQueryCmd := &cobra.Command{
		Use:                        ModuleName,
		Short:                      "fiat query sub commands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	fiatQueryCmd.AddCommand(client.GetCommands(
		cli.GetFiatAccountCmd(cdc),
		cli.GetAllFiatAccountsCmd(cdc),
	)...)

	return fiatQueryCmd
}

type AppModule struct {
	AppModuleBasic
	keeper Keeper
}

func NewAppModule(keeper Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         keeper,
	}
}

func (AppModule) Name() string {
	return ModuleName
}

func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {

}

func (AppModule) Route() string { return RouterKey }

func (am AppModule) NewHandler() sdk.Handler { return NewHandler(am.keeper) }

func (am AppModule) QuerierRoute() string { return QuerierRoute }

func (am AppModule) NewQuerierHandler() sdk.Querier { return NewQuerier(am.keeper) }

func (am AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abciTypes.ValidatorUpdate {
	var genesisState GenesisState

	_ = ModuleCdc.UnmarshalJSON(data, &genesisState)
	InitGenesis(ctx, am.keeper, genesisState)

	return []abciTypes.ValidatorUpdate{}
}

func (am AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	gs := ExportGenesis(ctx, am.keeper)

	return ModuleCdc.MustMarshalJSON(gs)
}

func (AppModule) BeginBlock(sdk.Context, abciTypes.RequestBeginBlock) {}

func (AppModule) EndBlock(_ sdk.Context, _ abciTypes.RequestEndBlock) []abciTypes.ValidatorUpdate {
	return []abciTypes.ValidatorUpdate{}
}

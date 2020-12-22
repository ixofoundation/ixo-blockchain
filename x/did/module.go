package did

import (
	"encoding/json"
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

type AppModuleBasic struct{}

func (AppModuleBasic) Name() string {
	return ModuleName
}

func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	RegisterLegacyAminoCodec(cdc)
}

func (AppModuleBasic) RegisterInterfaces(codectypes.InterfaceRegistry) {
}

//func (AppModuleBasic) DefaultGenesis() json.RawMessage {
//	return ModuleCdc.MustMarshalJSON(DefaultGenesisState())
//}

func (AppModuleBasic) DefaultGenesis(codec.JSONMarshaler) json.RawMessage {
	return ModuleCdc.MustMarshalJSON(DefaultGenesisState())
}

//func (AppModuleBasic) ValidateGenesis(bz json.RawMessage) error {
//	var data GenesisState
//	err := ModuleCdc.UnmarshalJSON(bz, &data)
//	if err != nil {
//		return err
//	}
//	return ValidateGenesis(data)
//}

func (AppModuleBasic) ValidateGenesis(cd codec.JSONMarshaler, enc client.TxEncodingConfig, bz json.RawMessage) error {
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

func (AppModuleBasic) RegisterGRPCGatewayRoutes(client.Context, *runtime.ServeMux) {
	//tx.RegisterGRPCGatewayRoutes()
}

func (AppModuleBasic) GetTxCmd(/*cdc *codec.Codec*/) *cobra.Command {
	didTxCmd := &cobra.Command{
		Use:                        ModuleName,
		Short:                      "did transaction sub commands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	didTxCmd.AddCommand(cli.GetCmdAddDidDoc())
	didTxCmd.AddCommand(cli.GetCmdAddCredential())

	//didTxCmd.AddCommand(flags.PostCommands(
	//	cli.GetCmdAddDidDoc(cdc)),
	//	cli.GetCmdAddCredential(cdc),
	//)...)

	return didTxCmd
}

func (AppModuleBasic) GetQueryCmd(/*cdc *codec.Codec*/) *cobra.Command {
	didQueryCmd := &cobra.Command{
		Use:                        ModuleName,
		Short:                      "did query sub commands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// TODO remove LegacyAmino from cli functions below and add them to command
	didQueryCmd.AddCommand(cli.GetCmdAddressFromBase58Pubkey())
	//didQueryCmd.AddCommand(cli.GetCmdAddressFromDid())
	didQueryCmd.AddCommand(cli.GetCmdIxoDidFromMnemonic())
	//didQueryCmd.AddCommand(cli.GetCmdDidDoc())
	//didQueryCmd.AddCommand(cli.GetCmdAllDids())
	//didQueryCmd.AddCommand(cli.GetCmdAllDidDocs())

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

type AppModule struct {
	AppModuleBasic
	keeper keeper.Keeper
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

// TODO Populate
func (am AppModule) Route() sdk.Route {
	return sdk.Route{} //RouterKey
}

func (am AppModule) NewHandler() sdk.Handler { return NewHandler(am.keeper) }

func (AppModule) QuerierRoute() string { return QuerierRoute }

func (am AppModule) LegacyQuerierHandler(cdc *codec.LegacyAmino) sdk.Querier {
	return NewQuerier(am.keeper, cdc)
}

func (AppModule) RegisterServices(module.Configurator) {}

//func (am AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
//	var genesisState GenesisState
//	ModuleCdc.MustUnmarshalJSON(data, &genesisState)
//	InitGenesis(ctx, am.keeper, genesisState)
//	return []abci.ValidatorUpdate{}
//}

func (AppModule) InitGenesis(sdk.Context, codec.JSONMarshaler, json.RawMessage) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONMarshaler) json.RawMessage {
	gs := ExportGenesis(ctx, am.keeper)
	return ModuleCdc.MustMarshalJSON(gs)
}

func (am AppModule) BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock) {
}

func (AppModule) EndBlock(_ sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

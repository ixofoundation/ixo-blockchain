package entity

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/ixofoundation/ixo-blockchain/x/entity/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/entity/types"
	entitycontracts "github.com/ixofoundation/ixo-blockchain/x/entity/types/contracts"
)

const (
	EntityNftContractName   = "entity_nft"
	EntityNftContractSymbol = "entity"
)

// NewParamChangeProposalHandler creates a new governance Handler for a ParamChangeProposal
func NewEntityParamChangeProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.InitializeNftContract:
			return handleTokenParameterChangeProposal(ctx, k, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized param proposal content type: %T", c)
		}
	}
}

func handleTokenParameterChangeProposal(ctx sdk.Context, k keeper.Keeper, p *types.InitializeNftContract) error {
	fmt.Printf("propsal handeler =============\n%+v\n", *p)
	fmt.Println("Supspace", k.ParamSpace.Name(), k.ParamSpace.HasKeyTable())
	var xx types.Params
	k.ParamSpace.GetParamSetIfExists(ctx, &xx)
	fmt.Printf("%+v\n", xx)

	// if err := msg.ValidateBasic(); err != nil {
	// 	return nil, err
	// }
	// ctx := sdk.UnwrapSDKContext(goCtx)

	adminAddr := authtypes.NewModuleAddress(types.NftModuleAddress())
	fmt.Println("================ehere1:")

	senderAddr, err := sdk.AccAddressFromBech32(p.NftMinterAddress)
	if err != nil {
		fmt.Println("================error1:", err)

		return err
	}
	// var adminAddr sdk.AccAddress
	// if msg.Admin != "" {
	// 	if adminAddr, err = sdk.AccAddressFromBech32(msg.Admin); err != nil {
	// 		return nil, sdkerrors.Wrap(err, "admin")
	// 	}
	// }

	// ctx.EventManager().EmitEvent(sdk.NewEvent(
	// 	sdk.EventTypeMessage,
	// 	sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
	// 	sdk.NewAttribute(sdk.AttributeKeySender, p.),
	// ))

	initiateNftContractMsg := entitycontracts.InitiateNftContract{
		Name:   EntityNftContractName,
		Symbol: EntityNftContractSymbol,
		Minter: adminAddr.String(),
	}

	fmt.Println("================ehere2:")
	encodedInitiateNftContractMsg, err := initiateNftContractMsg.Marshal()

	if err != nil {
		fmt.Println("================error2:", err)

		return err
	}

	deposit := sdk.NewCoins(sdk.NewCoin("uixo", sdk.ZeroInt()))
	fmt.Println("================ehere3:")

	contractAddr, _, err := k.WasmKeeper.Instantiate(ctx, p.NftContractCodeId, senderAddr, adminAddr, encodedInitiateNftContractMsg, "initiate_entity_nft_contract", deposit)
	if err != nil {
		fmt.Println("================error3:", err)

		return err
	}

	fmt.Println(contractAddr)

	xx.NftContractAddress = contractAddr.String()
	xx.NftContractMinter = initiateNftContractMsg.Minter

	k.ParamSpace.SetParamSet(ctx, &xx)

	var yy types.Params
	k.ParamSpace.GetParamSetIfExists(ctx, &yy)
	fmt.Println("address after sent ======", contractAddr)
	fmt.Println(yy)

	// for _, c := range p.Changes {
	// 	ss, ok := k.GetParams()
	// 	if !ok {
	// 		return sdkerrors.Wrap(proposal.ErrUnknownSubspace, c.Subspace)
	// 	}

	// 	// k.Logger(ctx).Info(
	// 	// 	fmt.Sprintf("attempt to set new parameter value; key: %s, value: %s", c.Key, c.Value),
	// 	// )

	// 	if err := ss.Update(ctx, []byte(c.Key), []byte(c.Value)); err != nil {
	// 		return sdkerrors.Wrapf(proposal.ErrSettingParameter, "key: %s, value: %s, err: %s", c.Key, c.Value, err.Error())
	// 	}
	// }

	return nil
}

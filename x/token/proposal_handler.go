package token

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/ixofoundation/ixo-blockchain/x/token/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/token/types"
)

const (
	TokenNftContractName   = "token_nft"
	TokenNftContractSymbol = "token"
)

// NewParamChangeProposalHandler creates a new governance Handler for a ParamChangeProposal
func NewTokenParamChangeProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.SetTokenContractCodes:
			return handleTokenParameterChangeProposal(ctx, k, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized param proposal content type: %T", c)
		}
	}
}

func handleTokenParameterChangeProposal(ctx sdk.Context, k keeper.Keeper, p *types.SetTokenContractCodes) error {
	fmt.Printf("token propsal handeler =============\n%+v\n", *p)
	fmt.Println("Supspace", k.ParamSpace.Name(), k.ParamSpace.HasKeyTable())
	var xx types.Params
	k.ParamSpace.GetParamSetIfExists(ctx, &xx)
	fmt.Printf("%+v\n", xx)

	// if err := msg.ValidateBasic(); err != nil {
	// 	return nil, err
	// }
	// ctx := sdk.UnwrapSDKContext(goCtx)

	// adminAddr := authtypes.NewModuleAddress(types.NftModuleAddress())

	// senderAddr, err := sdk.AccAddressFromBech32(p.NftMinterAddress)
	// if err != nil {
	// 	return nil
	// }
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

	// initiateNftContractMsg := tokencontracts.InitiateNftContract{
	// 	Name:   TokenNftContractName,
	// 	Symbol: TokenNftContractSymbol,
	// 	Minter: adminAddr.String(),
	// }

	// encodedInitiateNftContractMsg, err := initiateNftContractMsg.Marshal()
	// if err != nil {
	// 	return nil
	// }

	// deposit := sdk.NewCoins(sdk.NewCoin("uixo", sdk.ZeroInt()))

	// contractAddr, _, err := k.WasmKeeper.Instantiate(ctx, p.NftContractCodeId, senderAddr, adminAddr, encodedInitiateNftContractMsg, "initiate_token_nft_contract", deposit)
	// if err != nil {
	// 	return err
	// }

	xx.Cw20ContractCode = strconv.FormatUint(p.Cw20ContractCode, 10)
	xx.Cw721ContractCode = strconv.FormatUint(p.Cw721ContractCode, 10)
	xx.Ixo1155ContractCode = strconv.FormatUint(p.Ixo1155ContractCode, 10)

	k.ParamSpace.SetParamSet(ctx, &xx)

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

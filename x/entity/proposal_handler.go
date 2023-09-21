package entity

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/ixofoundation/ixo-blockchain/v2/x/entity/keeper"
	"github.com/ixofoundation/ixo-blockchain/v2/x/entity/types"
	nft "github.com/ixofoundation/ixo-blockchain/v2/x/entity/types/contracts"
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
	var xx types.Params
	k.ParamSpace.GetParamSetIfExists(ctx, &xx)
	adminAddr := authtypes.NewModuleAddress(types.NftModuleAddress())

	senderAddr, err := sdk.AccAddressFromBech32(p.NftMinterAddress)
	if err != nil {
		return err
	}

	initiateNftContractMsg := nft.InitiateNftContract{
		Name:   EntityNftContractName,
		Symbol: EntityNftContractSymbol,
		Minter: adminAddr.String(),
	}

	encodedInitiateNftContractMsg, err := nft.Marshal(initiateNftContractMsg)
	if err != nil {
		return err
	}

	contractAddr, _, err := k.WasmKeeper.Instantiate(ctx, p.NftContractCodeId, senderAddr, adminAddr, encodedInitiateNftContractMsg, "initiate_entity_nft_contract", sdk.NewCoins(sdk.NewCoin("uixo", sdk.ZeroInt())))
	if err != nil {
		// return nil as still want proposal to pass even though contractCode doest'n exist yet (for entity module bootstrap purposes)
		// if error it means proposal is just like empty proposal, look into returning error in future
		return nil
	}

	xx.NftContractAddress = contractAddr.String()
	xx.NftContractMinter = initiateNftContractMsg.Minter

	k.ParamSpace.SetParamSet(ctx, &xx)

	return nil
}

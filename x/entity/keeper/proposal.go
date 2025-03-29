package keeper

import (
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	"github.com/ixofoundation/ixo-blockchain/v5/x/entity/types"
	nft "github.com/ixofoundation/ixo-blockchain/v5/x/entity/types/contracts"
)

const (
	EntityNftContractName   = "entity_nft"
	EntityNftContractSymbol = "entity"
)

// NewEntityProposalHandler creates a new governance Handler for a ParamChangeProposal
func NewEntityProposalHandler(k Keeper) govtypesv1.Handler {
	return func(ctx sdk.Context, content govtypesv1.Content) error {
		switch c := content.(type) {
		case *types.InitializeNftContract:
			return k.handleNftInitializeContractProposal(ctx, c)

		default:
			return errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized param proposal content type: %T", c)
		}
	}
}

func (k Keeper) handleNftInitializeContractProposal(ctx sdk.Context, p *types.InitializeNftContract) error {
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

	contractAddr, _, err := k.WasmKeeper.Instantiate(ctx, p.NftContractCodeId, senderAddr, adminAddr, encodedInitiateNftContractMsg, "initiate_entity_nft_contract", sdk.NewCoins(sdk.NewCoin("uixo", math.ZeroInt())))
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

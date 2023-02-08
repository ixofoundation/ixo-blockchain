package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	iidtypes "github.com/ixofoundation/ixo-blockchain/x/iid/types"
	"github.com/ixofoundation/ixo-blockchain/x/token/types"
	"github.com/ixofoundation/ixo-blockchain/x/token/types/contracts/cw20"
	"github.com/ixofoundation/ixo-blockchain/x/token/types/contracts/cw721"
	"github.com/ixofoundation/ixo-blockchain/x/token/types/contracts/ixo1155"
)

type msgServer struct {
	Keeper Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the module MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(k Keeper) types.MsgServer {
	return &msgServer{
		Keeper: k,
	}
}

func (s msgServer) SetupMinter(goCtx context.Context, msg *types.MsgSetupMinter) (*types.MsgSetupMinterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := s.Keeper.GetParams(ctx)

	minterDidDoc, found := s.Keeper.IidKeeper.GetDidDocument(ctx, []byte(msg.MinterDid.Did()))
	if !found {
		return nil, sdkerrors.Wrapf(iidtypes.ErrDidDocumentNotFound, "did document for minter not found")
	}

	minterAddress, err := minterDidDoc.GetVerificationMethodBlockchainAddress(msg.MinterDid.String())
	if err != nil {
		return nil, err
	}

	var codeId uint64
	var label string
	var contractType types.ContractType
	var encodedInitiateMessage []byte

	switch contractInfo := msg.ContractInfo.(type) {
	case *types.MsgSetupMinter_Cw20:
		codeId = params.Cw20ContractCode
		label = fmt.Sprintf("%s-cw20-contract", msg.MinterDid.String())
		contractType = types.ContractType_CW20
		// uint := uint(contractInfo.Cw20.Cap)
		encodedInitiateMessage, err = cw20.InstantiateMsg{
			Name:     msg.Name,
			Symbol:   contractInfo.Cw20.Symbol,
			Decimals: contractInfo.Cw20.Decimals,
			InitialBalances: func() (coins []cw20.Cw20Coin) {
				for _, bal := range contractInfo.Cw20.InitialBalances {
					coins = append(coins, cw20.Cw20Coin{Address: bal.Address, Amount: bal.Amount})
				}
				return
			}(),
			Mint: cw20.MinterResponse{
				Minter: msg.MinterAddress,
				// Cap:    &contractInfo.Cw20.Cap,
			},
		}.Marshal()

	case *types.MsgSetupMinter_Cw721:
		codeId = params.Cw721ContractCode
		label = fmt.Sprintf("%s-cw721-contract", msg.MinterDid.String())
		contractType = types.ContractType_CW721
		encodedInitiateMessage, err = cw721.InitiateNftContract{
			Name:   msg.Name,
			Symbol: contractInfo.Cw721.Symbol,
			Minter: minterAddress.String(),
		}.Marshal()

	case *types.MsgSetupMinter_Cw1155:
		codeId = params.Ixo1155ContractCode
		label = fmt.Sprintf("%s-cw1155-contract", msg.MinterDid.String())
		contractType = types.ContractType_IXO1155
		encodedInitiateMessage, err = ixo1155.InstantiateMsg{
			Minter: minterAddress.String(),
		}.Marshal()

	default:
		return &types.MsgSetupMinterResponse{}, sdkerrors.ErrInvalidType.Wrap("invalid contract provided")
	}
	if err != nil {
		return &types.MsgSetupMinterResponse{}, err
	}

	contractAddr, _, err := s.Keeper.WasmKeeper.Instantiate(
		ctx,
		codeId,
		minterAddress,
		minterAddress,
		encodedInitiateMessage,
		label,
		sdk.NewCoins(sdk.NewCoin("uixo", sdk.ZeroInt())),
	)
	if err != nil {
		return nil, err
	}

	tokenMinter := types.TokenMinter{
		MinterDid:       msg.MinterDid,
		MinterAddress:   minterAddress.String(),
		ContractAddress: contractAddr.String(),
		ContractType:    contractType,
		Name:            msg.Name,
		Description:     msg.Description,
	}
	s.Keeper.SetMinter(ctx, tokenMinter)

	if err := ctx.EventManager().EmitTypedEvent(&tokenMinter); err != nil {
		s.Keeper.IidKeeper.Logger(ctx).Error("failed to emit TokerMinterEvent", err)
	}

	return &types.MsgSetupMinterResponse{}, nil
}

func (s msgServer) MintToken(goCtx context.Context, msg *types.MsgMintToken) (*types.MsgMintTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	tokenMinter, err := s.Keeper.GetMinterContract(ctx, msg.MinterDid, msg.ContractAddress)
	if err != nil {
		return nil, err
	}

	minterDidDoc, found := s.Keeper.IidKeeper.GetDidDocument(ctx, []byte(msg.MinterDid.Did()))
	if !found {
		return nil, sdkerrors.Wrapf(iidtypes.ErrDidDocumentNotFound, "did document for minter not found")
	}

	minterAddress, err := minterDidDoc.GetVerificationMethodBlockchainAddress(msg.MinterDid.String())
	if err != nil {
		return nil, err
	}

	ownerDidDoc, found := s.Keeper.IidKeeper.GetDidDocument(ctx, []byte(msg.OwnerDid.Did()))
	if !found {
		return nil, sdkerrors.Wrapf(iidtypes.ErrDidDocumentNotFound, "did document for owner not found")
	}

	ownerAddress, err := ownerDidDoc.GetVerificationMethodBlockchainAddress(msg.OwnerDid.String())
	if err != nil {
		return nil, err
	}

	var encodedMintMessage []byte

	switch mintInfo := msg.MintContract.(type) {
	case *types.MsgMintToken_Cw20:
		if tokenMinter.ContractType != types.ContractType_CW20 {
			return nil, sdkerrors.ErrInvalidType.Wrap("selected contract is not a cw20 contract type")
		}
		encodedMintMessage, err = cw20.Mint{
			Recipient: ownerAddress.String(),
			Amount:    uint(mintInfo.Cw20.Amount),
		}.Marshal()

	case *types.MsgMintToken_Cw721:
		if tokenMinter.ContractType != types.ContractType_CW721 {
			return nil, sdkerrors.ErrInvalidType.Wrap("selected contract is not a cw721 contract type")
		}

		uri := func() string {
			switch uri := mintInfo.Cw721.TokenUri.(type) {
			case *types.MintCw721_Uri:
				return uri.Uri
			case *types.MintCw721_Image:
				return uri.Image
			default:
				return ""
			}
		}()

		encodedMintMessage, err = cw721.WasmMsgMint{
			Mint: cw721.Mint{
				TokenId:  msg.OwnerDid.Did(),
				Owner:    ownerAddress.String(),
				TokenUri: uri,
			},
		}.Marshal()

	case *types.MsgMintToken_Cw1155:
		if tokenMinter.ContractType != types.ContractType_IXO1155 {
			return nil, sdkerrors.ErrInvalidType.Wrap("selected contract is not a cw1155 contract type")
		}
		encodedMintMessage, err = ixo1155.Mint{
			To:      ownerAddress.String(),
			TokenId: ixo1155.TokenId(msg.OwnerDid.Did()),
			Value:   mintInfo.Cw1155.Value,
			Msg:     []byte{},
		}.Marshal()

	default:
		// err := sdkerrors.Wrapf(ErrInvalidInput, "verification material not set for verification method %s", v.Method.Id)
		return nil, sdkerrors.ErrInvalidType.Wrap("invalid contract provided")
	}
	if err != nil {
		return nil, err
	}

	contractAddress, err := sdk.AccAddressFromBech32(tokenMinter.ContractAddress)
	if err != nil {
		return nil, err
	}

	_, err = s.Keeper.WasmKeeper.Execute(
		ctx,
		contractAddress,
		minterAddress,
		encodedMintMessage,
		sdk.NewCoins(sdk.NewCoin("uixo", sdk.ZeroInt())),
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgMintTokenResponse{}, nil
}

func (s msgServer) TransferToken(goCtx context.Context, msg *types.MsgTransferToken) (*types.MsgTransferTokenResponse, error) {
	return &types.MsgTransferTokenResponse{}, nil

}

package keeper

import (
	"context"
	"fmt"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	iidkeeper "github.com/ixofoundation/ixo-blockchain/x/iid/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/token/types"
	"github.com/ixofoundation/ixo-blockchain/x/token/types/contracts/cw20"
	"github.com/ixofoundation/ixo-blockchain/x/token/types/contracts/cw721"
	"github.com/ixofoundation/ixo-blockchain/x/token/types/contracts/ixo1155"
)

type msgServer struct {
	Keeper Keeper
}

func (s msgServer) wasmKeeper() wasmtypes.ContractOpsKeeper { return s.Keeper.WasmKeeper }
func (s msgServer) iidKeeper() iidkeeper.Keeper             { return s.Keeper.IidKeeper }

// NewMsgServerImpl returns an implementation of the project MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(k Keeper) types.MsgServer {
	return &msgServer{
		Keeper: k,
	}
}

func (s msgServer) SetupMinter(goCtx context.Context, msg *types.MsgSetupMinter) (*types.MsgSetupMinterResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := s.Keeper.GetParams(ctx)

	// s.wasmKeeper().
	minterDidDoc, found := s.iidKeeper().GetDidDocument(ctx, []byte(msg.MinterDid.Did()))
	if !found {
		return &types.MsgSetupMinterResponse{}, nil
	}

	minterAddress, err := minterDidDoc.GetVerificationMethodBlockchainAddress(msg.MinterDid.String())
	if err != nil {
		return &types.MsgSetupMinterResponse{}, err
	}

	var codeId uint64
	var label string
	var contractType types.ContractType
	var encodedInitiateMessage []byte

	switch contractInfo := msg.ContractInfo.(type) {
	case *types.MsgSetupMinter_Cw20:
		codeId = params.GetCw20ContractCode()
		label = fmt.Sprintf("%s-cw20-contract", msg.MinterDid.String())
		contractType = types.ContractType_CW20
		// uint := uint(contractInfo.Cw20.Cap)
		encodedInitiateMessage, err = cw20.InstantiateMsg{
			Name:     msg.Name,
			Symbol:   contractInfo.Cw20.Symbol,
			Decimals: contractInfo.Cw20.Decimals,
			InitialBalances: func() (coins []cw20.Cw20Coin) {
				for _, bal := range contractInfo.Cw20.InstialBalances {
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
		codeId = params.GetCw721ContractCode()
		label = fmt.Sprintf("%s-cw721-contract", msg.MinterDid.String())
		contractType = types.ContractType_CW721
		encodedInitiateMessage, err = cw721.InitiateNftContract{
			Name:   msg.Name,
			Symbol: contractInfo.Cw721.Symbol,
			Minter: minterAddress.String(),
		}.Marshal()

	case *types.MsgSetupMinter_Cw1155:
		codeId = params.GetIxo1155ContractCode()
		label = fmt.Sprintf("%s-cw1155-contract", msg.MinterDid.String())
		contractType = types.ContractType_IXO1155
		encodedInitiateMessage, err = ixo1155.InstantiateMsg{
			Minter: minterAddress.String(),
		}.Marshal()
	default:
		// err := sdkerrors.Wrapf(ErrInvalidInput, "verification material not set for verification method %s", v.Method.Id)
		return &types.MsgSetupMinterResponse{}, sdkerrors.ErrInvalidType.Wrap("invalid contract provided")
	}

	if err != nil {
		return &types.MsgSetupMinterResponse{}, err
	}

	contractAddr, _, err := s.wasmKeeper().Instantiate(
		ctx,
		codeId,
		minterAddress,
		minterAddress,
		encodedInitiateMessage,
		label,
		sdk.NewCoins(sdk.NewCoin("uixo", sdk.ZeroInt())),
	)

	if err != nil {
		return &types.MsgSetupMinterResponse{}, err
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
		s.iidKeeper().Logger(ctx).Error("failed to emit TokerMinterEvent", err)
	}

	return &types.MsgSetupMinterResponse{}, nil
}

func (s msgServer) MintToken(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	tokenMinter, err := s.Keeper.GetMinterContract(ctx, msg.MinterDid, msg.ContractAddress)
	if err != nil {
		return &types.MsgMintResponse{}, err
	}

	minterDidDoc, found := s.iidKeeper().GetDidDocument(ctx, []byte(msg.MinterDid.Did()))
	if !found {
		return &types.MsgMintResponse{}, nil
	}

	minterAddress, err := minterDidDoc.GetVerificationMethodBlockchainAddress(msg.MinterDid.String())
	if err != nil {
		return &types.MsgMintResponse{}, err
	}

	ownerDidDoc, found := s.iidKeeper().GetDidDocument(ctx, []byte(msg.OwnerDid.Did()))
	if !found {
		return &types.MsgMintResponse{}, nil
	}

	ownerAddress, err := ownerDidDoc.GetVerificationMethodBlockchainAddress(msg.OwnerDid.String())
	if err != nil {
		return &types.MsgMintResponse{}, err
	}

	var encodedMintMessage []byte

	switch mintInfo := msg.MintContract.(type) {
	case *types.MsgMint_Cw20:
		if tokenMinter.ContractType != types.ContractType_CW20 {
			return &types.MsgMintResponse{}, sdkerrors.ErrInvalidType.Wrap("selected contract is not a cw20 contract type")
		}
		encodedMintMessage, err = cw20.Mint{
			Recipient: ownerAddress.String(),
			Amount:    uint(mintInfo.Cw20.Amount),
		}.Marshal()

	case *types.MsgMint_Cw721:
		if tokenMinter.ContractType != types.ContractType_CW721 {
			return &types.MsgMintResponse{}, sdkerrors.ErrInvalidType.Wrap("selected contract is not a cw721 contract type")
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

	case *types.MsgMint_Cw1155:
		if tokenMinter.ContractType != types.ContractType_IXO1155 {
			return &types.MsgMintResponse{}, sdkerrors.ErrInvalidType.Wrap("selected contract is not a cw1155 contract type")
		}
		encodedMintMessage, err = ixo1155.Mint{
			To:      ownerAddress.String(),
			TokenId: ixo1155.TokenId(msg.OwnerDid.Did()),
			Value:   mintInfo.Cw1155.Value,
			Msg:     []byte{},
		}.Marshal()

	default:
		// err := sdkerrors.Wrapf(ErrInvalidInput, "verification material not set for verification method %s", v.Method.Id)
		return &types.MsgMintResponse{}, sdkerrors.ErrInvalidType.Wrap("invalid contract provided")
	}

	if err != nil {
		return &types.MsgMintResponse{}, err
	}

	contractAddressBytes, err := sdk.AccAddressFromBech32(tokenMinter.ContractAddress)
	if err != nil {
		return nil, err
	}

	s.wasmKeeper().Execute(
		ctx,
		contractAddressBytes,
		minterAddress,
		encodedMintMessage,
		sdk.NewCoins(sdk.NewCoin("uixo", sdk.ZeroInt())),
	)

	return &types.MsgMintResponse{}, nil
}

func (s msgServer) TransferToken(goCtx context.Context, msg *types.MsgTransferToken) (*types.MsgTransferTokenResponse, error) {
	return s.Keeper.TransferToken(sdk.UnwrapSDKContext(goCtx), msg)

}

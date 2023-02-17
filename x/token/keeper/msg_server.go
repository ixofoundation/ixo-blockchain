package keeper

import (
	"context"
	"crypto/md5"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	iidtypes "github.com/ixofoundation/ixo-blockchain/x/iid/types"
	"github.com/ixofoundation/ixo-blockchain/x/token/types"
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

func (s msgServer) CreateToken(goCtx context.Context, msg *types.MsgCreateToken) (*types.MsgCreateTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := s.Keeper.GetParams(ctx)

	_, found := s.Keeper.IidKeeper.GetDidDocument(ctx, []byte(msg.Class.Did()))
	if !found {
		return nil, sdkerrors.Wrapf(iidtypes.ErrDidDocumentNotFound, "class did document not found for %s", msg.Class.Did())
	}

	if s.Keeper.CheckTokensDuplicateName(ctx, msg.Name) {
		return nil, sdkerrors.Wrapf(types.ErrTokenNameDuplicate, msg.Name)
	}

	minterDidDoc, found := s.Keeper.IidKeeper.GetDidDocument(ctx, []byte(msg.MinterDid.Did()))
	if !found {
		return nil, sdkerrors.Wrapf(iidtypes.ErrDidDocumentNotFound, "did document for minter not found")
	}

	minterAddress, err := minterDidDoc.GetVerificationMethodBlockchainAddress(msg.MinterDid.String())
	if err != nil {
		return nil, err
	}

	encodedInitiateMessage, err := ixo1155.Marshal(ixo1155.InstantiateMsg{
		Minter: minterAddress.String(),
	})
	if err != nil {
		return nil, err
	}

	contractAddr, _, err := s.Keeper.WasmKeeper.Instantiate(
		ctx,
		params.Ixo1155ContractCode,
		minterAddress,
		minterAddress,
		encodedInitiateMessage,
		fmt.Sprintf("%s-ixo1155-contract", msg.MinterDid.String()),
		sdk.NewCoins(sdk.NewCoin("uixo", sdk.ZeroInt())),
	)
	if err != nil {
		return nil, err
	}

	token := types.Token{
		MinterDid:       msg.MinterDid,
		MinterAddress:   minterAddress.String(),
		ContractAddress: contractAddr.String(),
		Class:           msg.Class.Did(),
		Name:            msg.Name,
		Description:     msg.Description,
		Image:           msg.Image,
		Type:            msg.TokenType,
		Cap:             msg.Cap,
		Supply:          sdk.ZeroUint(),
		Paused:          false,
		Deactivated:     false,
	}
	s.Keeper.SetToken(ctx, token)

	if err := ctx.EventManager().EmitTypedEvent(&types.TokenCreatedEvent{
		ContractAddress: contractAddr.String(),
		Minter:          msg.MinterDid.Did(),
	}); err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to emit createToken event")
	}

	return &types.MsgCreateTokenResponse{}, nil
}

func (s msgServer) MintToken(goCtx context.Context, msg *types.MsgMintToken) (*types.MsgMintTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	token, err := s.Keeper.GetToken(ctx, msg.MinterDid, msg.ContractAddress)
	if err != nil {
		return nil, err
	}

	if token.Paused {
		return nil, types.ErrTokenPausedIncorrect
	}

	if token.Deactivated {
		return nil, types.ErrTokenDeactivatedIncorrect
	}

	minterDidDoc, found := s.Keeper.IidKeeper.GetDidDocument(ctx, []byte(msg.MinterDid.Did()))
	if !found {
		return nil, sdkerrors.Wrapf(iidtypes.ErrDidDocumentNotFound, "did document for minter not found %s", msg.MinterDid.Did())
	}

	minterAddress, err := minterDidDoc.GetVerificationMethodBlockchainAddress(msg.MinterDid.String())
	if err != nil {
		return nil, err
	}

	ownerDidDoc, found := s.Keeper.IidKeeper.GetDidDocument(ctx, []byte(msg.OwnerDid.Did()))
	if !found {
		return nil, sdkerrors.Wrapf(iidtypes.ErrDidDocumentNotFound, "did document for owner not found %s", msg.OwnerDid.Did())
	}

	ownerAddress, err := ownerDidDoc.GetVerificationMethodBlockchainAddress(msg.OwnerDid.String())
	if err != nil {
		return nil, err
	}

	contractAddress, err := sdk.AccAddressFromBech32(token.ContractAddress)
	if err != nil {
		return nil, err
	}

	var amounts = sdk.NewUint(0)
	var batches []types.MintBatchData

	// valiodation checks, check name is same as token and amounts dont go more than toen cap
	for _, batch := range msg.MintBatch {
		amounts = amounts.Add(batch.Amount)

		if token.Name != batch.Name {
			return nil, sdkerrors.Wrapf(types.ErrTokenNameIncorrect, "token name is not same as class token name %s", batch.Name)
		}

		// generate token id and uri
		tokenId := fmt.Sprintf("%x", md5.Sum([]byte(batch.Name+batch.Index)))
		tokenUri := types.TokenUriBase + tokenId

		batches = append(batches, types.MintBatchData{
			Id:         tokenId,
			Uri:        tokenUri,
			Name:       batch.Name,
			Index:      batch.Index,
			Collection: batch.Collection,
			TokenData:  batch.TokenData,
			Amount:     batch.Amount,
		})
	}

	// skip check if cap zero(unlimited), otherwise check if supply plus new amount is less than cap
	if !token.Cap.IsZero() && token.Supply.Add(amounts).GT(token.Cap) {
		return nil, sdkerrors.Wrapf(types.ErrTokenAmountIncorrect, "amounts %s plus current supply %s is greater than token cap %s", amounts.String(), token.Supply.String(), token.Cap.String())
	}

	encodedMintMessage, err := ixo1155.Marshal(ixo1155.WasmMsgBatchMint{
		BatchMint: ixo1155.BatchMint{
			To:    ownerAddress.String(),
			Batch: types.Map(batches, func(b types.MintBatchData) ixo1155.Batch { return b.GetWasmMintBatch() }),
		},
	})
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

	// Add new minted tokens to token supply and persist
	token.UpdateSupply(sdk.Int(amounts))
	s.Keeper.SetToken(ctx, token)

	// create and persist new token properties
	for _, batch := range batches {
		tokenProperties := types.TokenProperties{
			Id:         batch.Id,
			Index:      batch.Index,
			Collection: batch.Collection,
			TokenData:  batch.TokenData,
		}
		s.Keeper.SetTokenProperties(ctx, tokenProperties)
	}

	if err := ctx.EventManager().EmitTypedEvents(
		&types.TokenMintedEvent{
			ContractAddress: contractAddress.String(),
			Minter:          msg.MinterDid.Did(),
			Owner:           msg.OwnerDid.Did(),
			Batches:         types.Map(batches, func(b types.MintBatchData) *types.TokenMintedBatch { return b.GetTokenMintedEventBatch() }),
		},
		// TokenUpdatedEvent event since token supply has been update
		&types.TokenUpdatedEvent{
			ContractAddress: token.ContractAddress,
			Signer:          msg.MinterDid.Did(),
		},
	); err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to emit mint token events")
	}

	return &types.MsgMintTokenResponse{}, nil
}

func (s msgServer) TransferToken(goCtx context.Context, msg *types.MsgTransferToken) (*types.MsgTransferTokenResponse, error) {
	return &types.MsgTransferTokenResponse{}, nil

}

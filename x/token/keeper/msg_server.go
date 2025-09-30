package keeper

import (
	"context"
	"crypto/md5"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v6/x/iid/types"
	"github.com/ixofoundation/ixo-blockchain/v6/x/token/types"
	"github.com/ixofoundation/ixo-blockchain/v6/x/token/types/contracts/ixo1155"
)

// msgServer provides a way to reference keeper pointer in the message server interface.
type msgServer struct {
	Keeper *Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an instance of MsgServer for the provided keeper.
func NewMsgServerImpl(k *Keeper) types.MsgServer {
	return &msgServer{
		Keeper: k,
	}
}

func (s msgServer) CreateToken(goCtx context.Context, msg *types.MsgCreateToken) (*types.MsgCreateTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := s.Keeper.GetParams(ctx)

	_, found := s.Keeper.IidKeeper.GetDidDocument(ctx, []byte(msg.Class.Did()))
	if !found {
		return nil, errorsmod.Wrapf(iidtypes.ErrDidDocumentNotFound, "class did document not found for %s", msg.Class.Did())
	}

	if s.Keeper.CheckTokensDuplicateName(ctx, msg.Name) {
		return nil, errorsmod.Wrapf(types.ErrTokenNameDuplicate, msg.Name)
	}

	minter, err := sdk.AccAddressFromBech32(msg.Minter)
	if err != nil {
		return nil, err
	}

	encodedInitiateMessage, err := ixo1155.Marshal(ixo1155.InstantiateMsg{
		Minter: minter.String(),
	})
	if err != nil {
		return nil, err
	}

	contractAddr, _, err := s.Keeper.WasmKeeper.Instantiate(
		ctx,
		params.Ixo1155ContractCode,
		minter,
		minter,
		encodedInitiateMessage,
		fmt.Sprintf("%s-ixo1155-contract", msg.Minter),
		sdk.NewCoins(sdk.NewCoin("uixo", math.ZeroInt())),
	)
	if err != nil {
		return nil, err
	}

	token := types.Token{
		Minter:          msg.Minter,
		ContractAddress: contractAddr.String(),
		Class:           msg.Class.Did(),
		Name:            msg.Name,
		Description:     msg.Description,
		Image:           msg.Image,
		Type:            msg.TokenType,
		Cap:             msg.Cap,
		Supply:          math.ZeroUint(),
		Paused:          false,
		Stopped:         false,
	}
	s.Keeper.SetToken(ctx, token)

	if err := ctx.EventManager().EmitTypedEvent(&types.TokenCreatedEvent{
		Token: &token,
	}); err != nil {
		return nil, errorsmod.Wrapf(err, "failed to emit create Token event")
	}

	return &types.MsgCreateTokenResponse{}, nil
}

func (s msgServer) MintToken(goCtx context.Context, msg *types.MsgMintToken) (*types.MsgMintTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	token, err := s.Keeper.GetToken(ctx, msg.Minter, msg.ContractAddress)
	if err != nil {
		return nil, err
	}

	if token.Paused {
		return nil, types.ErrTokenPaused
	}

	if token.Stopped {
		return nil, types.ErrTokenStopped
	}

	minterAddress, err := sdk.AccAddressFromBech32(msg.Minter)
	if err != nil {
		return nil, err
	}

	ownerAddress, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	contractAddress, err := sdk.AccAddressFromBech32(token.ContractAddress)
	if err != nil {
		return nil, err
	}

	amounts := math.NewUint(0)
	var batches []types.MintBatchData

	// validation checks, check name is same as token and amounts dont go more than then cap
	for _, batch := range msg.MintBatch {
		amounts = amounts.Add(batch.Amount)

		if token.Name != batch.Name {
			return nil, errorsmod.Wrapf(types.ErrTokenNameIncorrect, "token name is not same as class token name %s", batch.Name)
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
		return nil, errorsmod.Wrapf(types.ErrTokenAmountIncorrect, "amounts %s plus current supply %s is greater than token cap %s", amounts.String(), token.Supply.String(), token.Cap.String())
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
		sdk.NewCoins(sdk.NewCoin("uixo", math.ZeroInt())),
	)
	if err != nil {
		return nil, err
	}

	// Add new minted tokens to token supply and persist
	token.Supply = token.Supply.Add(amounts)
	s.Keeper.SetToken(ctx, token)

	// create and persist new token properties
	for _, batch := range batches {
		tokenProperties := types.TokenProperties{
			Id:         batch.Id,
			Index:      batch.Index,
			Collection: batch.Collection,
			TokenData:  batch.TokenData,
			Name:       batch.Name,
		}
		s.Keeper.SetTokenProperties(ctx, tokenProperties)
		if err := ctx.EventManager().EmitTypedEvents(
			&types.TokenMintedEvent{
				ContractAddress: contractAddress.String(),
				Minter:          minterAddress.String(),
				Owner:           ownerAddress.String(),
				Amount:          batch.Amount,
				TokenProperties: &tokenProperties,
			},
		); err != nil {
			return nil, errorsmod.Wrapf(err, "failed to emit tokenMintEvent")
		}
	}

	if err := ctx.EventManager().EmitTypedEvents(
		&types.TokenUpdatedEvent{
			Token: &token,
		},
	); err != nil {
		return nil, errorsmod.Wrapf(err, "failed to emit tokenUpdatedEvent")
	}

	return &types.MsgMintTokenResponse{}, nil
}

func (s msgServer) TransferToken(goCtx context.Context, msg *types.MsgTransferToken) (*types.MsgTransferTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	recipientAddress, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, err
	}

	ownerAddress, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	// Get Token for the first token in tokens field
	_, token, err := s.Keeper.GetTokenById(ctx, msg.Tokens[0].Id)
	if err != nil {
		return nil, err
	}

	contractAddress, err := sdk.AccAddressFromBech32(token.ContractAddress)
	if err != nil {
		return nil, err
	}

	encodedTransferMessage, err := ixo1155.Marshal(ixo1155.WasmBatchSendFrom{
		BatchSendFrom: ixo1155.BatchSendFrom{
			From:  ownerAddress.String(),
			To:    recipientAddress.String(),
			Batch: types.Map(msg.Tokens, func(b *types.TokenBatch) ixo1155.Batch { return b.GetWasmTransferBatch() }),
		},
	})
	if err != nil {
		return nil, err
	}

	_, err = s.Keeper.WasmKeeper.Execute(
		ctx,
		contractAddress,
		ownerAddress,
		encodedTransferMessage,
		sdk.NewCoins(sdk.NewCoin("uixo", math.ZeroInt())),
	)
	if err != nil {
		return nil, err
	}

	if err := ctx.EventManager().EmitTypedEvent(
		&types.TokenTransferredEvent{
			Recipient: recipientAddress.String(),
			Owner:     ownerAddress.String(),
			Tokens:    msg.Tokens,
		},
	); err != nil {
		return nil, errorsmod.Wrapf(err, "failed to emit transfer token event")
	}

	return &types.MsgTransferTokenResponse{}, nil
}

func (s msgServer) RetireToken(goCtx context.Context, msg *types.MsgRetireToken) (*types.MsgRetireTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get Token for the first token in tokens field
	_, token, err := s.Keeper.GetTokenById(ctx, msg.Tokens[0].Id)
	if err != nil {
		return nil, err
	}

	ownerAddress, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	contractAddress, err := sdk.AccAddressFromBech32(token.ContractAddress)
	if err != nil {
		return nil, err
	}

	encodedBurnMessage, err := ixo1155.Marshal(ixo1155.WasmMsgBatchBurn{
		BatchBurn: ixo1155.BatchBurn{
			From:  ownerAddress.String(),
			Batch: types.Map(msg.Tokens, func(b *types.TokenBatch) ixo1155.Batch { return b.GetWasmTransferBatch() }),
		},
	})
	if err != nil {
		return nil, err
	}

	_, err = s.Keeper.WasmKeeper.Execute(
		ctx,
		contractAddress,
		ownerAddress,
		encodedBurnMessage,
		sdk.NewCoins(sdk.NewCoin("uixo", math.ZeroInt())),
	)
	if err != nil {
		return nil, err
	}

	// Update Token retired tokens and persist
	var retiredTokens []*types.TokensRetired
	for _, batch := range msg.Tokens {
		retiredTokens = append(retiredTokens, &types.TokensRetired{
			Id:           batch.Id,
			Amount:       batch.Amount,
			Reason:       msg.Reason,
			Jurisdiction: msg.Jurisdiction,
			Owner:        msg.Owner,
		})
	}
	token.Retired = append(token.Retired, retiredTokens...)
	s.Keeper.SetToken(ctx, *token)

	if err := ctx.EventManager().EmitTypedEvents(
		&types.TokenRetiredEvent{
			Owner:  ownerAddress.String(),
			Tokens: msg.Tokens,
		},
		&types.TokenUpdatedEvent{
			Token: token,
		},
	); err != nil {
		return nil, errorsmod.Wrapf(err, "failed to emit retire token event")
	}

	return &types.MsgRetireTokenResponse{}, nil
}

func (s msgServer) TransferCredit(goCtx context.Context, msg *types.MsgTransferCredit) (*types.MsgTransferCreditResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get Token for the first token in tokens field
	_, token, err := s.Keeper.GetTokenById(ctx, msg.Tokens[0].Id)
	if err != nil {
		return nil, err
	}

	ownerAddress, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	contractAddress, err := sdk.AccAddressFromBech32(token.ContractAddress)
	if err != nil {
		return nil, err
	}

	encodedBurnMessage, err := ixo1155.Marshal(ixo1155.WasmMsgBatchBurn{
		BatchBurn: ixo1155.BatchBurn{
			From:  ownerAddress.String(),
			Batch: types.Map(msg.Tokens, func(b *types.TokenBatch) ixo1155.Batch { return b.GetWasmTransferBatch() }),
		},
	})
	if err != nil {
		return nil, err
	}

	_, err = s.Keeper.WasmKeeper.Execute(
		ctx,
		contractAddress,
		ownerAddress,
		encodedBurnMessage,
		sdk.NewCoins(sdk.NewCoin("uixo", math.ZeroInt())),
	)
	if err != nil {
		return nil, err
	}

	// Update Token transferred tokens and persist
	var transferredCredits []*types.CreditsTransferred
	for _, batch := range msg.Tokens {
		transferredCredits = append(transferredCredits, &types.CreditsTransferred{
			Id:              batch.Id,
			Amount:          batch.Amount,
			Reason:          msg.Reason,
			Jurisdiction:    msg.Jurisdiction,
			Owner:           msg.Owner,
			AuthorizationId: msg.AuthorizationId,
		})
	}
	token.Transferred = append(token.Transferred, transferredCredits...)
	s.Keeper.SetToken(ctx, *token)

	if err := ctx.EventManager().EmitTypedEvents(
		&types.CreditsTransferredEvent{
			Owner:  ownerAddress.String(),
			Tokens: msg.Tokens,
		},
		&types.TokenUpdatedEvent{
			Token: token,
		},
	); err != nil {
		return nil, errorsmod.Wrapf(err, "failed to emit itmos transfer token event")
	}

	return &types.MsgTransferCreditResponse{}, nil
}

func (s msgServer) CancelToken(goCtx context.Context, msg *types.MsgCancelToken) (*types.MsgCancelTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get Token for the first token in tokens field
	_, token, err := s.Keeper.GetTokenById(ctx, msg.Tokens[0].Id)
	if err != nil {
		return nil, err
	}

	ownerAddress, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return nil, err
	}

	contractAddress, err := sdk.AccAddressFromBech32(token.ContractAddress)
	if err != nil {
		return nil, err
	}

	encodedBurnMessage, err := ixo1155.Marshal(ixo1155.WasmMsgBatchBurn{
		BatchBurn: ixo1155.BatchBurn{
			From:  ownerAddress.String(),
			Batch: types.Map(msg.Tokens, func(b *types.TokenBatch) ixo1155.Batch { return b.GetWasmTransferBatch() }),
		},
	})
	if err != nil {
		return nil, err
	}

	_, err = s.Keeper.WasmKeeper.Execute(
		ctx,
		contractAddress,
		ownerAddress,
		encodedBurnMessage,
		sdk.NewCoins(sdk.NewCoin("uixo", math.ZeroInt())),
	)
	if err != nil {
		return nil, err
	}

	// Update Token cancelled and remove amount from supply tokens and persist
	var cancelledTokens []*types.TokensCancelled
	amount := math.NewUint(0)

	for _, batch := range msg.Tokens {
		amount = amount.Add(batch.Amount)
		cancelledTokens = append(cancelledTokens, &types.TokensCancelled{
			Id:     batch.Id,
			Amount: batch.Amount,
			Reason: msg.Reason,
			Owner:  msg.Owner,
		})
	}

	token.Supply = token.Supply.Sub(amount)
	token.Cancelled = append(token.Cancelled, cancelledTokens...)
	s.Keeper.SetToken(ctx, *token)

	if err := ctx.EventManager().EmitTypedEvents(
		&types.TokenRetiredEvent{
			Owner:  ownerAddress.String(),
			Tokens: msg.Tokens,
		},
		// TokenUpdatedEvent event since token supply has been update
		&types.TokenUpdatedEvent{
			Token: token,
		},
	); err != nil {
		return nil, errorsmod.Wrapf(err, "failed to emit retire token event")
	}
	return &types.MsgCancelTokenResponse{}, nil
}

func (s msgServer) PauseToken(goCtx context.Context, msg *types.MsgPauseToken) (*types.MsgPauseTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	token, err := s.Keeper.GetToken(ctx, msg.Minter, msg.ContractAddress)
	if err != nil {
		return nil, err
	}

	// Update Token paused and persist
	token.Paused = msg.Paused
	s.Keeper.SetToken(ctx, token)

	if err := ctx.EventManager().EmitTypedEvents(
		&types.TokenPausedEvent{
			ContractAddress: token.ContractAddress,
			Minter:          token.Minter,
			Paused:          msg.Paused,
		},
		// TokenUpdatedEvent event since token supply has been update
		&types.TokenUpdatedEvent{
			Token: &token,
		},
	); err != nil {
		return nil, errorsmod.Wrapf(err, "failed to emit pause token event")
	}
	return &types.MsgPauseTokenResponse{}, nil
}

func (s msgServer) StopToken(goCtx context.Context, msg *types.MsgStopToken) (*types.MsgStopTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	token, err := s.Keeper.GetToken(ctx, msg.Minter, msg.ContractAddress)
	if err != nil {
		return nil, err
	}

	// Update Token stopped and persist
	token.Stopped = true
	s.Keeper.SetToken(ctx, token)

	if err := ctx.EventManager().EmitTypedEvents(
		&types.TokenStoppedEvent{
			ContractAddress: token.ContractAddress,
			Minter:          token.Minter,
			Stopped:         true,
		},
		// TokenUpdatedEvent event since token supply has been update
		&types.TokenUpdatedEvent{
			Token: &token,
		},
	); err != nil {
		return nil, errorsmod.Wrapf(err, "failed to emit stop token event")
	}
	return &types.MsgStopTokenResponse{}, nil
}

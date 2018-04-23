package project

import (
	"fmt"
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "project" type messages.
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case SetTrendMsg:
			return handleSetTrendMsg(ctx, k, msg)
		case QuizMsg:
			return handleQuizMsg(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized project Msg type: %v", reflect.TypeOf(msg).Name())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// Handle QuizMsg This is the engine of your module
func handleSetTrendMsg(ctx sdk.Context, k Keeper, msg SetTrendMsg) sdk.Result {
	k.setTrend(ctx, msg.Project)
	return sdk.Result{}
}

// Handle QuizMsg This is the engine of your module
func handleQuizMsg(ctx sdk.Context, k Keeper, msg QuizMsg) sdk.Result {

	correct := k.CheckTrend(ctx, msg.ProjectAnswer)

	if !correct {
		return ErrIncorrectProjectAnswer(msg.ProjectAnswer).Result()
	}

	if ctx.IsCheckTx() {
		return sdk.Result{} // TODO
	}

	bonusCoins := sdk.Coins{{msg.ProjectAnswer, 69}}

	_, err := k.ck.AddCoins(ctx, msg.Sender, bonusCoins)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{}
}

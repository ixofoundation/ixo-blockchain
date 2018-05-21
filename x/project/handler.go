package project

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(k ProjectKeeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		fmt.Println("Handler")
		fmt.Println(msg)
		switch msg := msg.(type) {
		case CreateProjectMsg:
			return handleCreateProjectMsg(ctx, k, msg)
		case CreateAgentMsg:
			return handleCreateAgentMsg(ctx, k, msg)
		case UpdateAgentMsg:
			return handleUpdateAgentMsg(ctx, k, msg)
		case CreateClaimMsg:
			return handleCreateClaimMsg(ctx, k, msg)
		case CreateEvaluationMsg:
			return handleCreateEvaluationMsg(ctx, k, msg)
		default:
			return sdk.ErrUnknownRequest("No match for message type.").Result()
		}
	}
}

func handleCreateProjectMsg(ctx sdk.Context, k ProjectKeeper, msg CreateProjectMsg) sdk.Result {
	fmt.Println("Handler")
	fmt.Println(msg)
	fmt.Println(msg.ProjectDoc)
	newProjectDoc := msg.ProjectDoc
	projectDoc, err := k.AddProjectDoc(ctx, newProjectDoc)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Code: sdk.CodeOK,
		Data: k.pm.encodeProject(projectDoc),
	}
}

func handleCreateAgentMsg(ctx sdk.Context, k ProjectKeeper, msg CreateAgentMsg) sdk.Result {
	return sdk.Result{
		Code: sdk.CodeOK,
		Data: []byte("Action complete"),
	}
}
func handleUpdateAgentMsg(ctx sdk.Context, k ProjectKeeper, msg UpdateAgentMsg) sdk.Result {
	return sdk.Result{
		Code: sdk.CodeOK,
		Data: []byte("Action complete"),
	}
}
func handleCreateClaimMsg(ctx sdk.Context, k ProjectKeeper, msg CreateClaimMsg) sdk.Result {
	return sdk.Result{
		Code: sdk.CodeOK,
		Data: []byte("Action complete"),
	}
}
func handleCreateEvaluationMsg(ctx sdk.Context, k ProjectKeeper, msg CreateEvaluationMsg) sdk.Result {
	return sdk.Result{
		Code: sdk.CodeOK,
		Data: []byte("Action complete"),
	}
}

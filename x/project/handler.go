package project

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(k ProjectKeeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case AddProjectMsg:
			return handleAddProjectDocMsg(ctx, k, msg)
		case GetProjectMsg:
			return handleGetProjectDocMsg(ctx, k, msg)
		default:
			return sdk.ErrUnknownRequest("No match for message type.").Result()
		}
	}
}

func handleAddProjectDocMsg(ctx sdk.Context, k ProjectKeeper, msg AddProjectMsg) sdk.Result {
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

func handleGetProjectDocMsg(ctx sdk.Context, k ProjectKeeper, msg GetProjectMsg) sdk.Result {
	projectDoc := k.GetProjectDoc(ctx, msg.Did)

	return sdk.Result{
		Code: sdk.CodeOK,
		Data: k.pm.encodeProject(projectDoc),
	}
}

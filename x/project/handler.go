package project

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

//NewHandler handle project requests
func NewHandler() sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		fmt.Println("In Project Handler: *****************************************************")
		
		return sdk.Result{}
	}
}

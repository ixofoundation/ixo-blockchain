package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
)

func CheckNotEmpty(value string, name string) (valid bool, err sdk.Error) {
	if strings.TrimSpace(value) == "" {
		return false, sdk.ErrUnknownRequest(name + " is empty.")
	} else {
		return true, nil
	}
}

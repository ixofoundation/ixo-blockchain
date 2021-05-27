package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strings"
)

func CheckNotEmpty(value string, name string) (valid bool, err error) {
	if strings.TrimSpace(value) == "" {
		return false, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "%s is empty", name)
	} else {
		return true, nil
	}
}

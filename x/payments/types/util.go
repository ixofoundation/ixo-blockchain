package types

import (
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func CheckNotEmpty(value string, name string) (valid bool, err error) {
	if strings.TrimSpace(value) == "" {
		return false, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "%s is empty", name)
	} else {
		return true, nil
	}
}

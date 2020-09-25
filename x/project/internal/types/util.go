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

func withoutQuotes(s string) string {
	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}
	return s
}

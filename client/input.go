package ixoClient

import (
	"bufio"
	"os"
	"strings"

	"github.com/bgentry/speakeasy"
	isatty "github.com/mattn/go-isatty"
	"github.com/pkg/errors"
)

// GetString will prompt for a string
func GetString(prompt string, minLength int, buf *bufio.Reader) (str string, err error) {
	if inputIsTty() {
		str, err = speakeasy.Ask(prompt)
	} else {
		str, err = readLineFromBuf(buf)
	}
	if err != nil {
		return "", err
	}
	if len(str) < minLength {
		return "", errors.Errorf("Must be at least %d characters", minLength)
	}

	return str, nil
}

// inputIsTty returns true iff we have an interactive prompt,
// where we can disable echo and request to repeat the password.
// If false, we can optimize for piped input from another command
func inputIsTty() bool {
	return isatty.IsTerminal(os.Stdin.Fd()) || isatty.IsCygwinTerminal(os.Stdin.Fd())
}

// readLineFromBuf reads one line from stdin.
// Subsequent calls reuse the same buffer, so we don't lose
// any input when reading a password twice (to verify)
func readLineFromBuf(buf *bufio.Reader) (string, error) {
	pass, err := buf.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(pass), nil
}

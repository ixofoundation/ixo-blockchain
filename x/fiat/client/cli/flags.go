package cli

import (
	flag "github.com/spf13/pflag"
)

// noLint
const (
	FlagTo      = "to"
	FlagFrom    = "from"
	FlagAddress = "address"
)

var (
	fsTo      = flag.NewFlagSet("", flag.ContinueOnError)
	fsAddress = flag.NewFlagSet("", flag.ContinueOnError)
	fsFrom    = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	fsTo.String(FlagTo, "", "Address to send coins")
	fsAddress.String(FlagAddress, "", "Address to query")
	fsFrom.String(FlagFrom, "", "address of buyer account")
}

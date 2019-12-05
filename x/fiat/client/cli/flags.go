package cli

import (
	flag "github.com/spf13/pflag"
)

// noLint
const (
	FlagTo                 = "to"
	FlagFrom               = "from"
	FlagPegHash            = "pegHash"
)

var (
	fsTo                 = flag.NewFlagSet("", flag.ContinueOnError)
	fsPegHash            = flag.NewFlagSet("", flag.ContinueOnError)
	fsFrom               = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	fsTo.String(FlagTo, "", "Address to send coins")
	fsPegHash.String(FlagPegHash, "", "Peg Hash to be negotiated ")
	fsFrom.String(FlagFrom, "", "address of buyer account")
}

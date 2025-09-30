package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/ixofoundation/ixo-blockchain/v6/x/token/client/cli"
)

var ProposalHandler = govclient.NewProposalHandler(cli.NewCmdUpdateTokenParamsProposal)

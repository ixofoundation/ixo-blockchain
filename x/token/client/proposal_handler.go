package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/ixofoundation/ixo-blockchain/x/token/client/cli"
	"github.com/ixofoundation/ixo-blockchain/x/token/client/rest"
)

var ProposalHandler = govclient.NewProposalHandler(cli.NewCmdUpdateTokenParamsProposal, rest.ProposalRESTHandler)

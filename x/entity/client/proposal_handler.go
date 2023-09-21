package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/ixofoundation/ixo-blockchain/v2/x/entity/client/cli"
	"github.com/ixofoundation/ixo-blockchain/v2/x/entity/client/rest"
)

var ProposalHandler = govclient.NewProposalHandler(cli.NewCmdUpdateEntityParamsProposal, rest.ProposalRESTHandler)

package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/ixofoundation/ixo-blockchain/v5/x/entity/client/cli"
)

var ProposalHandler = govclient.NewProposalHandler(cli.NewCmdUpdateEntityParamsProposal)

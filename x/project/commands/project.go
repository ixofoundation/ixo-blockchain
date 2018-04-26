
package project

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/wire"
)

//CreateProjectCmd to create Project
func CreateProjectCmd(cdc *wire.Codec) *cobra.Command {
	cmdr := commander{
		cdc,
	}
	return &cobra.Command{
		Use:   "create project",
		Short: "Creates project",
		RunE:  cmdr.createProjectCmd,
	}
}

type commander struct {
	cdc	*wire.Codec
}

func (c commander) createProjectCmd(cmd *cobra.Command, args []string) error {

	fmt.Println(args[0]);
	
	if len(args) != 1 || len(args[0]) == 0 {
		return errors.New("You must provide an account name")
	}

	return nil
}

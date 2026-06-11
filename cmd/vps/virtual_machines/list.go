package virtual_machines

import (
	"context"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
	"log"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get virtual machine list",
	Long:  `This endpoint retrieves a list of all available virtual machines.`,
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSGetVirtualMachinesV1WithResponse(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

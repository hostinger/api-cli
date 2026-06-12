package virtual_machines

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get virtual machines",
	Long:  "Retrieve all available virtual machines.\n\nUse this endpoint to view available VPS instances.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSGetVirtualMachinesV1WithResponse(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

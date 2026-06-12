package templates

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get templates",
	Long:  "Retrieve available OS templates for virtual machines.\n\nUse this endpoint to view operating system options before creating or recreating VPS instances.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSGetTemplatesV1WithResponse(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

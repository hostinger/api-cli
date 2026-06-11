package templates

import (
	"context"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
	"log"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get template list",
	Long:  `This endpoint retrieves a list of available OS templates for virtual machines.`,
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSGetTemplatesV1WithResponse(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

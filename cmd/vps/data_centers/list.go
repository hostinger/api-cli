package data_centers

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get data center list",
	Long:  "Retrieve all available data centers.\n\nUse this endpoint to view location options before deploying VPS instances.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSGetDataCenterListV1WithResponse(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

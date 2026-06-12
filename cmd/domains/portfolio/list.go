package portfolio

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get domain list",
	Long:  "Retrieve all domains associated with your account.\n\nUse this endpoint to view user's domain portfolio.",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().DomainsGetDomainListV1WithResponse(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

package domains

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DeleteWebsiteSubdomainCmd = &cobra.Command{
	Use:   "delete-website-subdomain <username> <domain> <subdomain>",
	Short: "Delete website subdomain",
	Long:  "Delete an existing subdomain from the selected website.\n\nUse this endpoint to remove subdomains that are no longer needed.",
	Args:  cobra.MatchAll(cobra.ExactArgs(3)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingDeleteWebsiteSubdomainV1WithResponse(context.TODO(), args[0], args[1], args[2])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

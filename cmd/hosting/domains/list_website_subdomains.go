package domains

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListWebsiteSubdomainsCmd = &cobra.Command{
	Use:   "list-website-subdomains <username> <domain>",
	Short: "List website subdomains",
	Long:  "Retrieve all subdomains created under the selected website.\n\nUse this endpoint to inspect subdomain configuration for a specific website,\nincluding the parent domain and root directory assigned to each subdomain.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingListWebsiteSubdomainsV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

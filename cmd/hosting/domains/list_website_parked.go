package domains

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListWebsiteParkedCmd = &cobra.Command{
	Use:   "list-website-parked <username> <domain>",
	Short: "List website parked domains",
	Long:  "Retrieve all parked or alias domains created under the selected website.\n\nUse this endpoint to inspect parked domain configuration for a specific website,\nincluding the parent domain and root directory assigned to each parked domain.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingListWebsiteParkedDomainsV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

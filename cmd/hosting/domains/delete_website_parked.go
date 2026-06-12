package domains

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DeleteWebsiteParkedCmd = &cobra.Command{
	Use:   "delete-website-parked <username> <domain> <parked-domain>",
	Short: "Delete website parked domain",
	Long:  "Delete an existing parked or alias domain from the selected website.\n\nUse this endpoint to remove parked domains that are no longer needed.",
	Args:  cobra.MatchAll(cobra.ExactArgs(3)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingDeleteWebsiteParkedDomainV1WithResponse(context.TODO(), args[0], args[1], args[2])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

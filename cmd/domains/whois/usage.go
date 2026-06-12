package whois

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var UsageCmd = &cobra.Command{
	Use:   "usage <whois-id>",
	Short: "Get WHOIS profile usage",
	Long:  "Retrieve domain list where provided WHOIS contact profile is used.\n\nUse this endpoint to view which domains use specific contact profiles.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().DomainsGetWHOISProfileUsageV1WithResponse(context.TODO(), utils.StringToInt(args[0]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

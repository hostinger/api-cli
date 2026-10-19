package cdn

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var StatusCmd = &cobra.Command{
	Use:   "status <username> <domain>",
	Short: "Get CDN status",
	Long:  "Returns whether the Hostinger CDN is enabled for the website domain. Works for the main\ndomain, parked domains and subdomains of the account.\n\nThe domain must be the exact website name as listed on the account, without a `www.` prefix.\nReturns 404 when the domain is not a website on the account.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingGetCDNStatusV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

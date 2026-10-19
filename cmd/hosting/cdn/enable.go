package cdn

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var EnableCmd = &cobra.Command{
	Use:   "enable <username> <domain>",
	Short: "Enable CDN",
	Long:  "Turns on the Hostinger CDN for the website domain, which speeds up page loads and reduces the\nload on the server. Works for the main domain, parked domains and subdomains of the account.\n\nEnabling changes the domain's DNS records and usually takes a few seconds. After that, DNS\npropagation can take a few more minutes before visitors are served through the CDN.\n\nRepeating the call is safe, for example after a timeout. When CDN is already enabled, it\napplies the DNS change again and refreshes the CDN SSL certificate.\n\nReturns 422 when the account is disabled, while the account is being moved to another\nserver (retry after the move finishes), or when the CDN provider rejects the change.\nReturns 404 when the domain is not a website on the account.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingEnableCDNV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

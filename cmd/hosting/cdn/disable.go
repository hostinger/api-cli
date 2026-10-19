package cdn

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DisableCmd = &cobra.Command{
	Use:   "disable <username> <domain>",
	Short: "Disable CDN",
	Long:  "Turns off the Hostinger CDN for the website domain. Once DNS propagation completes, the\nwebsite content is served directly from the origin server. Repeating the call is safe; when\nCDN is already disabled, it applies the DNS change again.\n\nReturns 422 for platform-managed domains (for example `*.hostingersite.com` and\n`*.app-preview.com`), where CDN cannot be disabled, even when CDN is already off. Also returns\n422 while the account is being moved to another server (retry after the move finishes), or\nwhen the CDN provider rejects the change. Returns 404 when the domain is not a website on the\naccount.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingDisableCDNV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

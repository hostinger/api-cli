package ssl

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ReinstallWebsiteCmd = &cobra.Command{
	Use:   "reinstall-website <website_uid> <domain>",
	Short: "Reinstall website SSL",
	Long:  "Replaces the Let's Encrypt certificate of the domain: the current platform certificate, when\none is recorded, is revoked and removed, then a new setup starts in the background. Returns at\nonce; `Get website SSL status` reports `installing` while it runs, then `active` or `failed`.\n\nReturns 422 for free subdomains, when a certificate process is recorded for the domain (a\nfailed setup counts until it is cleaned up), or when the domain hit its limit of three setups\nper seven days. Returns 429 when the same domain was requested less than a minute ago, and 403\nwhen the website is suspended or locked, and 404 when the website or the domain does not exist.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().AgencyHostingReinstallWebsiteSSLV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

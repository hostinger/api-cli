package ssl

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var UninstallWebsiteCmd = &cobra.Command{
	Use:   "uninstall-website <website_uid> <domain>",
	Short: "Uninstall website SSL",
	Long:  "Removes the platform-issued Let's Encrypt certificate of the domain: the certificate is revoked\nand deleted before the response, so the domain is no longer served with a platform certificate\nuntil a new setup completes. Also succeeds when the domain has no platform certificate to\nremove. Uploaded (custom) certificates are not affected.\n\nReturns 422 when a certificate process is recorded for the domain (a failed setup counts until\nit is cleaned up), 429 when the same domain was requested less than a minute ago, and 403 when\nthe website is suspended or locked, and 404 when the website or the domain does not exist.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().AgencyHostingUninstallWebsiteSSLV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

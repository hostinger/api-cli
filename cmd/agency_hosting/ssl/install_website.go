package ssl

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var InstallWebsiteCmd = &cobra.Command{
	Use:   "install-website <website_uid> <domain>",
	Short: "Install website SSL",
	Long:  "Starts a Let's Encrypt certificate setup for the domain and returns at once; the setup runs in\nthe background. `Get website SSL status` reports `installing` while it runs, then `active` or\n`failed`; the `ssl_setup` entry of `List website processes` shows the same progress.\n\nReturns 422 when the domain already has a platform certificate that is not expired, when a\ncertificate process is recorded for the domain (a failed setup counts until it is cleaned up),\nor when the domain hit its limit of three setups per seven days. Returns 429 when the same\ndomain was requested less than a minute ago, 403 when the website is suspended or locked, and\n404 when the website or the domain does not exist.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().AgencyHostingInstallWebsiteSSLV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

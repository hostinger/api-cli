package ssl

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var WebsiteStatusCmd = &cobra.Command{
	Use:   "website-status <website_uid> <domain>",
	Short: "Get website SSL status",
	Long:  "Returns the SSL state of one domain of an Agency Plan website: the certificate `status`,\nwhether the certificate was uploaded by the customer, and when it stops being valid.\n\n`installing` means a certificate setup is running or retrying; the `ssl_setup` entry of\n`List website processes` shows the same progress. `active` means a valid certificate is in\nplace: uploaded by the customer, issued by the platform, or a lifetime certificate bought for\nthe domain. `failed` means the last setup gave up and no valid certificate is in place.\n`expired` means the certificate has run out. `not_installed` means the domain has no\ncertificate and no setup process. Returns 404 when the website or the domain does not exist.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().AgencyHostingGetWebsiteSSLStatusV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

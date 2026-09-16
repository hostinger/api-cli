package ssl

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var StatusCmd = &cobra.Command{
	Use:   "status <username> <domain>",
	Short: "Get SSL status",
	Long:  "Returns the SSL state of the website: the certificate `status` and `provider`, whether the\ncertificate is a lifetime one managed by the platform, whether HTTP requests are redirected to\nHTTPS, when the certificate stops being valid and the last installation error.\n\n`installing` and `waiting_for_retry` mean an installation is in progress. `failed` means the\nlast installation gave up, or the website was not updated for 60 minutes while `installing`;\n`last_error` holds the reason when it is a known message, otherwise it is null. `expired`\nmeans the assigned certificate's validity has ended. `not_installed` means no certificate is\nassigned. Free subdomains use a platform-managed certificate: with no installation recorded\nthey report `active` with `provider` and `expires_at` null.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingGetSSLStatusV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

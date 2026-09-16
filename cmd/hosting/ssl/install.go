package ssl

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var InstallCmd = &cobra.Command{
	Use:   "install <username> <domain>",
	Short: "Install SSL",
	Long:  "Requests a lifetime SSL certificate for the website. The installation runs in the background;\n`Get SSL status` reports `active` or `failed` when it ends. An `active` lifetime certificate\ndoes not block the request: a new installation is requested, which is how a certificate is\nreinstalled.\n\nReturns 422 for free subdomains (their certificate is managed by the platform), while an\ninstallation is `installing` or `waiting_for_retry`, when the website's certificate was\nrevoked (it cannot be reissued), and when an uploaded custom certificate is installed; that\none has to be uninstalled first.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingInstallSSLV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

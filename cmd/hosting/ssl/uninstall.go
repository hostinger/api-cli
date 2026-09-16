package ssl

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var UninstallCmd = &cobra.Command{
	Use:   "uninstall <username> <domain>",
	Short: "Uninstall SSL",
	Long:  "Removes the SSL certificate assigned to the website, turns the HTTPS redirect off and cancels\na pending installation retry. The website serves plain HTTP until a new installation\ncompletes. `Get SSL status` reports `not_installed` as soon as the call returns; the call also\nsucceeds when no certificate is assigned, so repeating it is safe.\n\nReturns 422 for free subdomains (their certificate is managed by the platform) and while an\ninstallation is `installing`.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().HostingUninstallSSLV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

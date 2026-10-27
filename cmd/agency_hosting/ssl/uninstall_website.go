package ssl

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var UninstallWebsiteCmd = &cobra.Command{
	Use:     "uninstall-website <website_uid> <domain>",
	Short:   "Uninstall website SSL",
	Long:    "Removes the platform-issued Let's Encrypt certificate of the domain: the certificate is revoked\nand deleted before the response, so the domain is no longer served with a platform certificate\nuntil a new setup completes. Also succeeds when the domain has no platform certificate to\nremove. Uploaded (custom) certificates are not affected.\n\nReturns 422 when a certificate process is recorded for the domain (a failed setup counts until\nit is cleaned up), 429 when the same domain was requested less than a minute ago, and 403 when\nthe website is suspended or locked, and 404 when the website or the domain does not exist.",
	Example: "  hostinger agency-hosting ssl uninstall-website zpwlGlp19 mydomain.tld",
	ValidArgsFunction: utils.CompleteArgs(
		nil,
		&utils.ArgSource{
			Value: "domain",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().DomainsGetDomainListV1WithResponse(ctx)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		// Last step before the request: every input check has passed. Name
		// the target, which may have come from the link.
		if err := utils.Confirm(cmd, "Uninstall website SSL", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().AgencyHostingUninstallWebsiteSSLV1WithResponse(ctx, args[0], args[1])
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		}
		respBody, status, err := call(cmd.Context())
		if err != nil {
			return err
		}
		return output.Format(cmd, respBody, status)
	},
}

func init() {
	utils.RegisterArgs(UninstallWebsiteCmd, utils.StringArg("website_uid"), utils.StringArg("domain"))
	UninstallWebsiteCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}

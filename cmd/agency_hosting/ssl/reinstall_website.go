package ssl

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ReinstallWebsiteCmd = &cobra.Command{
	Use:     "reinstall-website <website_uid> <domain>",
	Short:   "Reinstall website SSL",
	Long:    "Replaces the Let's Encrypt certificate of the domain: the current platform certificate, when\none is recorded, is revoked and removed, then a new setup starts in the background. Returns at\nonce; `Get website SSL status` reports `installing` while it runs, then `active` or `failed`.\n\nReturns 422 for free subdomains, when a certificate process is recorded for the domain (a\nfailed setup counts until it is cleaned up), or when the domain hit its limit of three setups\nper seven days. Returns 429 when the same domain was requested less than a minute ago, and 403\nwhen the website is suspended or locked, and 404 when the website or the domain does not exist.",
	Example: "  hostinger agency-hosting ssl reinstall-website zpwlGlp19 mydomain.tld",
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
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().AgencyHostingReinstallWebsiteSSLV1WithResponse(ctx, args[0], args[1])
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
	utils.RegisterArgs(ReinstallWebsiteCmd, utils.StringArg("website_uid"), utils.StringArg("domain"))
}

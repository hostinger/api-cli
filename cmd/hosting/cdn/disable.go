package cdn

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var DisableCmd = &cobra.Command{
	Use:     "disable [domain]",
	Short:   "Disable CDN",
	Long:    "Turns off the Hostinger CDN for the website domain. Once DNS propagation completes, the\nwebsite content is served directly from the origin server. Repeating the call is safe; when\nCDN is already disabled, it applies the DNS change again.\n\nReturns 422 for platform-managed domains (for example `*.hostingersite.com` and\n`*.app-preview.com`), where CDN cannot be disabled, even when CDN is already off. Also returns\n422 while the account is being moved to another server (retry after the move finishes), or\nwhen the CDN provider rejects the change. Returns 404 when the domain is not a website on the\naccount.\n\nThe website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger hosting cdn disable mydomain.tld",
	ValidArgsFunction: utils.CompleteArgs(&utils.ArgSource{
		Value: "domain",
		Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
			r, err := api.CompletionRequest().HostingListWebsitesV1WithResponse(ctx, nil)
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		},
	}),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		username, domain, args, err := site.Resolve(cmd, args, 0)
		if err != nil {
			return api.CommandError(err)
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingDisableCDNV1WithResponse(ctx, username, domain)
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
	utils.RegisterScopedArgs(DisableCmd, 2, utils.StringArg("username"), utils.StringArg("domain"))
}

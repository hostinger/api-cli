package cache

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var StatusCmd = &cobra.Command{
	Use:     "status [domain]",
	Short:   "Get cache status",
	Long:    "Returns the caching state of the website: whether server-side caching is enabled and whether\ndevelopment (cacheless) mode is on. Works for the main domain, addon websites, parked domains\nand subdomains of the account; each reports its own settings and does not inherit them from\nits parent website.\n\n`is_enabled` is the setting changed by `Toggle website cache` and reflects a toggle\nimmediately. `is_cacheless_mode_enabled` is true while development mode is on; caching is then\nbypassed even when `is_enabled` is true. `Toggle cacheless mode` is applied in the background,\nso this flag can take a few minutes to follow it.\n\nThe domain must be the exact website name as listed on the account, without a `www.` prefix.\nReturns 404 when the domain is not a website on the account.\n\nThe website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger hosting cache status mydomain.tld",
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
			r, err := api.Request().HostingGetCacheStatusV1WithResponse(ctx, username, domain)
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
	utils.RegisterScopedArgs(StatusCmd, 2, utils.StringArg("username"), utils.StringArg("domain"))
}

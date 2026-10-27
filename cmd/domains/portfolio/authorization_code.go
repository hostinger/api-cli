package portfolio

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var AuthorizationCodeCmd = &cobra.Command{
	Use:     "authorization-code [domain]",
	Short:   "Get domain authorization code",
	Long:    "Retrieve the authorization (EPP) code for a specified domain so it can be transferred\naway from Hostinger to another registrar.\n\nRequesting a new code invalidates any code retrieved previously.\n\nUse this endpoint to obtain the code required to transfer a domain to another registrar.\n\nThe domain can be given as [domain] or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link` when that website's domain belongs to the account, default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger domains portfolio authorization-code mydomain.tld",
	ValidArgsFunction: utils.CompleteArgs(&utils.ArgSource{
		Value: "domain",
		Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
			r, err := api.CompletionRequest().DomainsGetDomainListV1WithResponse(ctx)
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		},
	}),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		domain, args, err := site.ResolveDomain(cmd, args, 0, "domains")
		if err != nil {
			return api.CommandError(err)
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().DomainsGetDomainAuthorizationCodeV1WithResponse(ctx, domain)
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
	utils.RegisterScopedArgs(AuthorizationCodeCmd, 1, utils.StringArg("domain"))
}

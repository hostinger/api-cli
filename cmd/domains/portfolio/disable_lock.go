package portfolio

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var DisableLockCmd = &cobra.Command{
	Use:     "disable-lock [domain]",
	Short:   "Disable domain lock",
	Long:    "Disable domain lock for the domain.\n\nDomain lock needs to be disabled before transferring the domain to another registrar.\n\nUse this endpoint to prepare domains for transfer to other registrars.\n\nThe domain can be given as [domain] or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link` when that website's domain belongs to the account, default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger domains portfolio disable-lock mydomain.tld",
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
		// A destructive command acts without a question only on a target that was typed.
		typed := len(args) > 0
		domain, args, err := site.ResolveDomain(cmd, args, 0, "domains")
		if err != nil {
			return api.CommandError(err)
		}
		// Last step before the request: every input check has passed. Name
		// the target, which may have come from the link.
		if err := utils.ConfirmScoped(cmd, "Disable domain lock", append([]string{domain}, args...), typed); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().DomainsDisableDomainLockV1WithResponse(ctx, domain)
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
	utils.RegisterScopedArgs(DisableLockCmd, 1, utils.StringArg("domain"))
	DisableLockCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}

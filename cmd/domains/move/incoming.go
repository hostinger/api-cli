package move

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var IncomingCmd = &cobra.Command{
	Use:     "incoming [domain]",
	Short:   "Get incoming domain move",
	Long:    "Retrieve the incoming move for a specified domain.\n\nReturns 404 when no account is moving this domain to you.\n\nUse this endpoint to check whether a domain addressed to you is still waiting to be accepted.\n\nThe domain can be given as [domain] or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link` when that website's domain belongs to the account, default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger domains move incoming mydomain.tld",
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
		params := incomingParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().DomainsGetIncomingDomainMoveV1WithResponse(ctx, domain, params)
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
	utils.RegisterScopedArgs(IncomingCmd, 1, utils.StringArg("domain"))
	IncomingCmd.Flags().BoolP("force-sync", "", false, "Re-check the move against the registry before responding. Only has an effect while the move is in the `activating` status.")
	IncomingCmd.RegisterFlagCompletionFunc("force-sync", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
}

func incomingParams(cmd *cobra.Command) *client.DomainsGetIncomingDomainMoveV1Params {
	params := &client.DomainsGetIncomingDomainMoveV1Params{}
	if cmd.Flags().Changed("force-sync") {
		v, _ := cmd.Flags().GetBool("force-sync")
		params.ForceSync = &v
	}
	return params
}

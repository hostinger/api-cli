package websites

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:     "delete [domain]",
	Short:   "Delete website",
	Long:    "This endpoint permanently removes a website and all of its data. This action\ncannot be undone. Before calling it, make sure the user understands the\nconsequences and explicitly confirms that they want to proceed.\n\nAll website files, databases and related configuration will be removed.\nThe hosting plan itself is kept, so a new website can be created on it afterwards.\n\nSupported websites: main and addon domain websites on web hosting plans, and\nWebsite Builder websites. Parked domains and subdomains cannot be deleted with\nthis endpoint. The domain must be the exact website domain, not a preview\ndomain or an alias.\n\nReturns 404 when the domain does not exist or does not belong to the\nauthenticated client.\n\nWebsite removal is processed asynchronously and can take a few minutes to\ncomplete. The response returns before the removal finishes.\n\nThe website can be given as [domain] or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger hosting websites delete mydomain.tld",
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
		// A destructive command acts without a question only on a target that was typed.
		typed := len(args) > 0
		domain, args, err := site.ResolveDomain(cmd, args, 0, "hosting")
		if err != nil {
			return api.CommandError(err)
		}
		// Last step before the request: every input check has passed. Name
		// the target, which may have come from the link.
		if err := utils.ConfirmScoped(cmd, "Delete website", append([]string{domain}, args...), typed); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingDeleteWebsiteV1WithResponse(ctx, domain)
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
	utils.RegisterScopedArgs(DeleteCmd, 1, utils.StringArg("domain"))
	DeleteCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}

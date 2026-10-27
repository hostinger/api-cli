package nodejs

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var RestartApplicationCmd = &cobra.Command{
	Use:     "restart-application [domain]",
	Short:   "Restart Node.js application",
	Long:    "Restarts the Node.js server process for the website. Does not rebuild or redeploy the\napplication. Use it to apply environment or configuration changes, or to recover a hung\napplication.\n\nOnly applicable to server-side applications (Express, Next.js, NestJS, etc.). Static\nfront-end apps (React, Vue, Vite) have no persistent server process, so restarting them\nhas no effect. Returns success even when the website has no server process to restart.\n\nThe website can be given as [domain] (the account username is looked up), as <username> <domain>, or left out, in which case it comes from --domain, HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, the default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger hosting nodejs restart-application mydomain.tld",
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
		username, domain, args, err := site.Resolve(cmd, args, 0)
		if err != nil {
			return api.CommandError(err)
		}
		// Last step before the request: every input check has passed. Name
		// the target, which may have come from the link.
		if err := utils.ConfirmScoped(cmd, "Restart Node.js application", append([]string{domain}, args...), typed); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingRestartNodeJsApplicationV1WithResponse(ctx, username, domain)
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
	utils.RegisterScopedArgs(RestartApplicationCmd, 2, utils.StringArg("username"), utils.StringArg("domain"))
	RestartApplicationCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}

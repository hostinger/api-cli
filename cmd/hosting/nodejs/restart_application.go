package nodejs

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var RestartApplicationCmd = &cobra.Command{
	Use:     "restart-application <username> <domain>",
	Short:   "Restart Node.js application",
	Long:    "Restarts the Node.js server process for the website. Does not rebuild or redeploy the\napplication. Use it to apply environment or configuration changes, or to recover a hung\napplication.\n\nOnly applicable to server-side applications (Express, Next.js, NestJS, etc.). Static\nfront-end apps (React, Vue, Vite) have no persistent server process, so restarting them\nhas no effect. Returns success even when the website has no server process to restart.",
	Example: "  hostinger hosting nodejs restart-application u123456789 mydomain.tld",
	ValidArgsFunction: utils.CompleteArgs(
		&utils.ArgSource{
			Value: "username",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().HostingListWebsitesV1WithResponse(ctx, nil)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
		&utils.ArgSource{
			Value: "domain",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().HostingListWebsitesV1WithResponse(ctx, nil)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		// Last step before the request: every input check has passed.
		if err := utils.Confirm(cmd, "Restart Node.js application", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingRestartNodeJsApplicationV1WithResponse(ctx, args[0], args[1])
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
	utils.RegisterArgs(RestartApplicationCmd, utils.StringArg("username"), utils.StringArg("domain"))
	RestartApplicationCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}

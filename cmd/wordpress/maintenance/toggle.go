package maintenance

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ToggleCmd = &cobra.Command{
	Use:     "toggle [username] <software>",
	Short:   "Toggle maintenance mode",
	Long:    "Enable or disable maintenance mode for the specified WordPress installation,\nbased on the `enabled` flag.\n\nProvide the WordPress installation (software) identifier in the path. It can\nbe obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).\n\nThe hosting account can be given as [username] or left out, in which case it comes from --username, HOSTINGER_USERNAME, the website named by --domain or HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger wordpress maintenance toggle u123456789 1232456789 --enabled=true",
	ValidArgsFunction: utils.CompleteArgs(&utils.ArgSource{
		Value: "username",
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
		username, args, err := site.ResolveUsername(cmd, args, 1)
		if err != nil {
			return api.CommandError(err)
		}
		body, err := toggleBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingToggleMaintenanceModeV1WithBodyWithResponse(ctx, username, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterScopedArgs(ToggleCmd, 1, utils.StringArg("username"), utils.StringArg("software"))
	ToggleCmd.Flags().BoolP("enabled", "", false, "Enable (true) or disable (false) maintenance mode for the WordPress installation.")
	ToggleCmd.MarkFlagRequired("enabled")
	ToggleCmd.RegisterFlagCompletionFunc("enabled", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
}

func toggleBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	enabledVal, _ := cmd.Flags().GetBool("enabled")
	body["enabled"] = enabledVal
	return body, nil
}

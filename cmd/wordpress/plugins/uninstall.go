package plugins

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var UninstallCmd = &cobra.Command{
	Use:     "uninstall <username> <software>",
	Short:   "Uninstall WordPress plugins",
	Long:    "Uninstall one or more plugins from a WordPress installation.\n\nProvide the WordPress installation (software) identifier in the path. It can\nbe obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).\n\nThis operation is asynchronous: a successful response only means the uninstall\njob has been queued.",
	Example: "  hostinger wordpress plugins uninstall u123456789 1232456789 --plugins akismet,hello-dolly",
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
		nil,
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := uninstallBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		// Last step before the request: every input check has passed.
		if err := utils.Confirm(cmd, "Uninstall WordPress plugins", args); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingUninstallWordPressPluginsV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(UninstallCmd, utils.StringArg("username"), utils.StringArg("software"))
	UninstallCmd.Flags().StringSliceP("plugins", "", nil, "Slugs of the installed plugins to uninstall.")
	UninstallCmd.MarkFlagRequired("plugins")
	UninstallCmd.RegisterFlagCompletionFunc("plugins", cobra.NoFileCompletions)
	UninstallCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}

func uninstallBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	pluginsVal, _ := cmd.Flags().GetStringSlice("plugins")
	body["plugins"] = pluginsVal
	return body, nil
}

package object_cache

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ToggleMemcachedCmd = &cobra.Command{
	Use:     "toggle-memcached <username> <software>",
	Short:   "Toggle Memcached object cache",
	Long:    "Activate or deactivate the Memcached object cache for the specified WordPress\ninstallation, based on the `enabled` flag.\n\nProvide the WordPress installation (software) identifier in the path. It can\nbe obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).",
	Example: "  hostinger wordpress object-cache toggle-memcached u123456789 1232456789 --enabled=true",
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
		body, err := toggleMemcachedBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingToggleMemcachedObjectCacheV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(ToggleMemcachedCmd, utils.StringArg("username"), utils.StringArg("software"))
	ToggleMemcachedCmd.Flags().BoolP("enabled", "", false, "Activate (true) or deactivate (false) the Memcached object cache for the WordPress installation.")
	ToggleMemcachedCmd.MarkFlagRequired("enabled")
	ToggleMemcachedCmd.RegisterFlagCompletionFunc("enabled", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
}

func toggleMemcachedBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	enabledVal, _ := cmd.Flags().GetBool("enabled")
	body["enabled"] = enabledVal
	return body, nil
}

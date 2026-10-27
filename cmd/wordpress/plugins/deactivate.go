package plugins

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

var DeactivateCmd = &cobra.Command{
	Use:     "deactivate [username] <software>",
	Short:   "Deactivate WordPress plugin",
	Long:    "Deactivate an installed plugin on a WordPress installation.\n\nProvide the WordPress installation (software) identifier in the path. It can\nbe obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).\n\nThis operation is asynchronous: a successful response only means the\ndeactivation job has been queued.\n\nThe hosting account can be given as [username] or left out, in which case it comes from --username, HOSTINGER_USERNAME, the website named by --domain or HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger wordpress plugins deactivate u123456789 1232456789 --plugin akismet",
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
		// A destructive command acts without a question only on a target that was typed.
		typed := len(args) > 1
		username, args, err := site.ResolveUsername(cmd, args, 1)
		if err != nil {
			return api.CommandError(err)
		}
		body, err := deactivateBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		// Last step before the request: every input check has passed. Name
		// the target, which may have come from the link.
		if err := utils.ConfirmScoped(cmd, "Deactivate WordPress plugin", append([]string{username}, args...), typed); err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingDeactivateWordPressPluginV1WithBodyWithResponse(ctx, username, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterScopedArgs(DeactivateCmd, 1, utils.StringArg("username"), utils.StringArg("software"))
	DeactivateCmd.Flags().StringP("plugin", "", "", "Slug of the installed plugin to deactivate.")
	DeactivateCmd.MarkFlagRequired("plugin")
	DeactivateCmd.RegisterFlagCompletionFunc("plugin", cobra.NoFileCompletions)
	DeactivateCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
}

func deactivateBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	pluginVal, _ := cmd.Flags().GetString("plugin")
	body["plugin"] = pluginVal
	return body, nil
}

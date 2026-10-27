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

var InstallCmd = &cobra.Command{
	Use:     "install [username] <software>",
	Short:   "Install WordPress plugins",
	Long:    "Install one or more plugins on an existing WordPress installation.\n\nProvide the WordPress installation (software) identifier in the path. It can\nbe obtained from GET /api/hosting/v1/wordpress/installations (the `id`\nfield). Use GET /api/hosting/v1/wordpress/plugins to discover the plugin\nslugs available for installation.\n\nThis operation is asynchronous: a successful response only means the install\njob has been queued, not that the plugins are ready.\n\nThe hosting account can be given as [username] or left out, in which case it comes from --username, HOSTINGER_USERNAME, the website named by --domain or HOSTINGER_DOMAIN, the .hostinger/site.json written by `hostinger up` or `hostinger link`, default_domain in the config file, or a prompt when run in a terminal.",
	Example: "  hostinger wordpress plugins install u123456789 1232456789 --plugins akismet,hello-dolly",
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
		body, err := installBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingInstallWordPressPluginsV1WithBodyWithResponse(ctx, username, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterScopedArgs(InstallCmd, 1, utils.StringArg("username"), utils.StringArg("software"))
	InstallCmd.Flags().StringSliceP("plugins", "", nil, "Plugin slugs to install. Use GET /api/hosting/v1/wordpress/plugins to discover available slugs.")
	InstallCmd.MarkFlagRequired("plugins")
	InstallCmd.RegisterFlagCompletionFunc("plugins", cobra.NoFileCompletions)
}

func installBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	pluginsVal, _ := cmd.Flags().GetStringSlice("plugins")
	body["plugins"] = pluginsVal
	return body, nil
}

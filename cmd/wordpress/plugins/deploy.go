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

var DeployCmd = &cobra.Command{
	Use:     "deploy <username> <domain>",
	Short:   "Deploy WordPress plugin",
	Long:    "Deploy a WordPress plugin from an already uploaded directory.\n\nThis endpoint allows you to deploy a WordPress plugin that has been uploaded to the website's directory.\nThe plugin will be activated and made available in the WordPress admin panel.",
	Example: "  hostinger wordpress plugins deploy u123456789 mydomain.tld --plugin-path my-plugin-new --slug my-plugin",
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
				r, err := api.CompletionRequest().DomainsGetDomainListV1WithResponse(ctx)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := deployBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingDeployWordPressPluginV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(DeployCmd, utils.StringArg("username"), utils.StringArg("domain"))
	DeployCmd.Flags().StringP("plugin-path", "", "", "Relative path to the plugin directory from wp-content/plugins")
	DeployCmd.Flags().StringP("slug", "", "", "Slug of the plugin")
	DeployCmd.MarkFlagRequired("plugin-path")
	DeployCmd.MarkFlagRequired("slug")
	DeployCmd.RegisterFlagCompletionFunc("plugin-path", cobra.NoFileCompletions)
	DeployCmd.RegisterFlagCompletionFunc("slug", cobra.NoFileCompletions)
}

func deployBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	pluginPathVal, _ := cmd.Flags().GetString("plugin-path")
	body["plugin_path"] = pluginPathVal
	slugVal, _ := cmd.Flags().GetString("slug")
	body["slug"] = slugVal
	return body, nil
}

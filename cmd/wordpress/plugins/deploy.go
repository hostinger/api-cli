package plugins

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DeployCmd = &cobra.Command{
	Use:   "deploy <username> <domain>",
	Short: "Deploy WordPress plugin",
	Long:  "Deploy a WordPress plugin from an already uploaded directory.\n\nThis endpoint allows you to deploy a WordPress plugin that has been uploaded to the website's directory.\nThe plugin will be activated and made available in the WordPress admin panel.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(deployBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingDeployWordPressPluginV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	DeployCmd.Flags().StringP("plugin-path", "", "", "Relative path to the plugin directory from wp-content/plugins")
	DeployCmd.Flags().StringP("slug", "", "", "Slug of the plugin")
	DeployCmd.MarkFlagRequired("plugin-path")
	DeployCmd.MarkFlagRequired("slug")
}

func deployBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	pluginPathVal, _ := cmd.Flags().GetString("plugin-path")
	body["plugin_path"] = pluginPathVal
	slugVal, _ := cmd.Flags().GetString("slug")
	body["slug"] = slugVal
	return body
}

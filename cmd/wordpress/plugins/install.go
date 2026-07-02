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

var InstallCmd = &cobra.Command{
	Use:   "install <username> <software>",
	Short: "Install WordPress plugins",
	Long:  "Install one or more plugins on an existing WordPress installation.\n\nProvide the WordPress installation (software) identifier in the path. It can\nbe obtained from GET /api/hosting/v1/wordpress/installations (the `id`\nfield). Use GET /api/hosting/v1/wordpress/plugins to discover the plugin\nslugs available for installation.\n\nThis operation is asynchronous: a successful response only means the install\njob has been queued, not that the plugins are ready.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(installBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingInstallWordPressPluginsV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	InstallCmd.Flags().StringSliceP("plugins", "", nil, "Plugin slugs to install. Use GET /api/hosting/v1/wordpress/plugins to discover available slugs.")
	InstallCmd.MarkFlagRequired("plugins")
}

func installBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	pluginsVal, _ := cmd.Flags().GetStringSlice("plugins")
	body["plugins"] = pluginsVal
	return body
}

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

var ActivateCmd = &cobra.Command{
	Use:   "activate <username> <software>",
	Short: "Activate WordPress plugin",
	Long:  "Activate an installed plugin on a WordPress installation.\n\nProvide the WordPress installation (software) identifier in the path. It can\nbe obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).\n\nThis operation is asynchronous: a successful response only means the activation\njob has been queued.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(activateBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingActivateWordPressPluginV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ActivateCmd.Flags().StringP("plugin", "", "", "Slug of the installed plugin to activate.")
	ActivateCmd.MarkFlagRequired("plugin")
}

func activateBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	pluginVal, _ := cmd.Flags().GetString("plugin")
	body["plugin"] = pluginVal
	return body
}

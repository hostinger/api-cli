package object_cache

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ToggleMemcachedCmd = &cobra.Command{
	Use:   "toggle-memcached <username> <software>",
	Short: "Toggle Memcached object cache",
	Long:  "Activate or deactivate the Memcached object cache for the specified WordPress\ninstallation, based on the `enabled` flag.\n\nProvide the WordPress installation (software) identifier in the path. It can\nbe obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(toggleMemcachedBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingToggleMemcachedObjectCacheV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ToggleMemcachedCmd.Flags().BoolP("enabled", "", false, "Activate (true) or deactivate (false) the Memcached object cache for the WordPress installation.")
	ToggleMemcachedCmd.MarkFlagRequired("enabled")
}

func toggleMemcachedBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	enabledVal, _ := cmd.Flags().GetBool("enabled")
	body["enabled"] = enabledVal
	return body
}

package maintenance

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ToggleCmd = &cobra.Command{
	Use:   "toggle <username> <software>",
	Short: "Toggle maintenance mode",
	Long:  "Enable or disable maintenance mode for the specified WordPress installation,\nbased on the `enabled` flag.\n\nProvide the WordPress installation (software) identifier in the path. It can\nbe obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(toggleBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingToggleMaintenanceModeV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ToggleCmd.Flags().BoolP("enabled", "", false, "Enable (true) or disable (false) maintenance mode for the WordPress installation.")
	ToggleCmd.MarkFlagRequired("enabled")
}

func toggleBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	enabledVal, _ := cmd.Flags().GetBool("enabled")
	body["enabled"] = enabledVal
	return body
}

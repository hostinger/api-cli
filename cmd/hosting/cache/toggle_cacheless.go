package cache

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ToggleCachelessCmd = &cobra.Command{
	Use:   "toggle-cacheless <username> <domain>",
	Short: "Toggle cacheless mode",
	Long:  "Turns development (cacheless) mode on or off, based on the enabled flag. When enabled, nothing\nis cached, effectively turning off all caching for the website; use it while actively developing,\ntesting changes, debugging issues, or when real-time updates must be visible. Disable it after\nfinishing development work to restore the performance benefits of caching.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(toggleCachelessBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingToggleCachelessModeV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ToggleCachelessCmd.Flags().BoolP("enabled", "", false, "Turn development (cacheless) mode on (true) or off (false) for the website.")
	ToggleCachelessCmd.MarkFlagRequired("enabled")
}

func toggleCachelessBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	enabledVal, _ := cmd.Flags().GetBool("enabled")
	body["enabled"] = enabledVal
	return body
}

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

var ToggleWebsiteCmd = &cobra.Command{
	Use:   "toggle-website <username> <domain>",
	Short: "Toggle website cache",
	Long:  "Turns server-side caching for the website on or off, based on the enabled flag. Enable it for\nfaster page loads, reduced server load, and improved user experience; recommended for production\nwebsites. Disabling may impact performance; to temporarily bypass caching while developing or\ndebugging, prefer toggling cacheless mode instead.\n\nDoes nothing if caching is already in the requested state.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(toggleWebsiteBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingToggleWebsiteCacheV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ToggleWebsiteCmd.Flags().BoolP("enabled", "", false, "Turn server-side caching on (true) or off (false) for the website.")
	ToggleWebsiteCmd.MarkFlagRequired("enabled")
}

func toggleWebsiteBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	enabledVal, _ := cmd.Flags().GetBool("enabled")
	body["enabled"] = enabledVal
	return body
}

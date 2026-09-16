package ssl

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ToggleHttpsRedirectCmd = &cobra.Command{
	Use:   "toggle-https-redirect <username> <domain>",
	Short: "Toggle HTTPS redirect",
	Long:  "Turns the HTTP to HTTPS redirect of the website on or off, based on `is_enabled`. Does\nnothing when the redirect is already in the requested state. Turning it on requires an\ninstalled certificate (`status` `active` or `expired` on `Get SSL status`) and returns 422\nwhen there is none; turning it off is always accepted.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(toggleHttpsRedirectBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingToggleHTTPSRedirectV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ToggleHttpsRedirectCmd.Flags().BoolP("is-enabled", "", false, "Turn the HTTP to HTTPS redirect on (true) or off (false) for the website.")
	ToggleHttpsRedirectCmd.MarkFlagRequired("is-enabled")
}

func toggleHttpsRedirectBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	isEnabledVal, _ := cmd.Flags().GetBool("is-enabled")
	body["is_enabled"] = isEnabledVal
	return body
}

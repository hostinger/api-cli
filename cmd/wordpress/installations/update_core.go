package installations

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var UpdateCoreCmd = &cobra.Command{
	Use:   "update-core <username> <software>",
	Short: "Update WordPress core",
	Long:  "Update the WordPress core for the specified installation (minor update or a\nspecific version).\n\nProvide the WordPress installation (software) identifier in the path. It can\nbe obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).\n\nThis operation is asynchronous: a successful response only means the update\njob has been queued.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(updateCoreBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingUpdateWordPressCoreV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	UpdateCoreCmd.Flags().BoolP("minor", "", false, "Update the minor version only.")
	UpdateCoreCmd.Flags().StringP("version", "", "", "Update to a specific WordPress core version.")
}

func updateCoreBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("minor") {
		v, _ := cmd.Flags().GetBool("minor")
		body["minor"] = v
	}
	if cmd.Flags().Changed("version") {
		v, _ := cmd.Flags().GetString("version")
		body["version"] = v
	}
	return body
}

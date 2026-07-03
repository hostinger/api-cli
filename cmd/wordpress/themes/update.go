package themes

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:   "update <username> <software>",
	Short: "Update WordPress themes",
	Long:  "Update one or more installed themes to their latest version on a WordPress\ninstallation.\n\nProvide the WordPress installation (software) identifier in the path. It can\nbe obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).\n\nThis operation is asynchronous: a successful response only means the update job\nhas been queued.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(updateBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingUpdateWordPressThemesV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	UpdateCmd.Flags().StringSliceP("themes", "", nil, "Slugs of the installed themes to update to their latest version.")
	UpdateCmd.MarkFlagRequired("themes")
}

func updateBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	themesVal, _ := cmd.Flags().GetStringSlice("themes")
	body["themes"] = themesVal
	return body
}

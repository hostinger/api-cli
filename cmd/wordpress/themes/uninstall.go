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

var UninstallCmd = &cobra.Command{
	Use:   "uninstall <username> <software>",
	Short: "Uninstall WordPress themes",
	Long:  "Uninstall one or more themes from a WordPress installation.\n\nProvide the WordPress installation (software) identifier in the path. It can\nbe obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).\n\nThis operation is asynchronous: a successful response only means the uninstall\njob has been queued.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(uninstallBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingUninstallWordPressThemesV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	UninstallCmd.Flags().StringSliceP("themes", "", nil, "Slugs of the installed themes to uninstall.")
	UninstallCmd.MarkFlagRequired("themes")
}

func uninstallBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	themesVal, _ := cmd.Flags().GetStringSlice("themes")
	body["themes"] = themesVal
	return body
}

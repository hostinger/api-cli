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

var DeleteCmd = &cobra.Command{
	Use:   "delete <username> <software>",
	Short: "Delete WordPress installation",
	Long:  "Delete the specified WordPress installation, with optional file and database\nremoval. This removes all associated components including plugins, themes,\nstaging websites and any other related data.\n\nProvide the WordPress installation (software) identifier in the path. It can\nbe obtained from GET /api/hosting/v1/wordpress/installations (the `id` field).",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(deleteBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().HostingDeleteWordPressInstallationV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	DeleteCmd.Flags().BoolP("delete-database", "", false, "Delete the installation database.")
	DeleteCmd.Flags().BoolP("delete-files", "", false, "Delete installation files from disk.")
}

func deleteBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("delete-database") {
		v, _ := cmd.Flags().GetBool("delete-database")
		body["delete_database"] = v
	}
	if cmd.Flags().Changed("delete-files") {
		v, _ := cmd.Flags().GetBool("delete-files")
		body["delete_files"] = v
	}
	return body
}

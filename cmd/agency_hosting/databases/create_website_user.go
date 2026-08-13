package databases

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var CreateWebsiteUserCmd = &cobra.Command{
	Use:   "create-website-user <website_uid> <database_name>",
	Short: "Create website database user",
	Long:  "Creates a user for an existing database on an Agency Plan website.\n\nEach database supports a single non-system user; creating a user for a database that already has one fails.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(createWebsiteUserBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().AgencyHostingCreateWebsiteDatabaseUserV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateWebsiteUserCmd.Flags().StringP("database-user", "", "", "Database username to create (alphanumeric and underscores).")
	CreateWebsiteUserCmd.Flags().StringP("host", "", "", "Host the user connects from (IPv4, IPv6, % wildcard, or localhost). Defaults to localhost.")
	CreateWebsiteUserCmd.Flags().StringP("password", "", "", "Password for the database user (requires mixed case, letters, and numbers).")
	CreateWebsiteUserCmd.MarkFlagRequired("database-user")
	CreateWebsiteUserCmd.MarkFlagRequired("password")
}

func createWebsiteUserBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	databaseUserVal, _ := cmd.Flags().GetString("database-user")
	body["database_user"] = databaseUserVal
	if cmd.Flags().Changed("host") {
		v, _ := cmd.Flags().GetString("host")
		body["host"] = v
	}
	passwordVal, _ := cmd.Flags().GetString("password")
	body["password"] = passwordVal
	return body
}

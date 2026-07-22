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

var CreatePlanWebsiteUserCmd = &cobra.Command{
	Use:   "create-plan-website-user <website_uid> <database_name>",
	Short: "Create Agency Plan website database user",
	Long:  "Creates a user for an existing database on an Agency Plan website.\n\nEach database supports a single non-system user; creating a user for a database that already has one fails.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(createPlanWebsiteUserBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().AgencyHostingCreateAgencyPlanWebsiteDatabaseUserV1WithBodyWithResponse(context.TODO(), args[0], args[1], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreatePlanWebsiteUserCmd.Flags().StringP("database-user", "", "", "Database username to create (alphanumeric and underscores).")
	CreatePlanWebsiteUserCmd.Flags().StringP("host", "", "", "Host the user connects from (IPv4, IPv6, % wildcard, or localhost). Defaults to localhost.")
	CreatePlanWebsiteUserCmd.Flags().StringP("password", "", "", "Password for the database user (requires mixed case, letters, and numbers).")
	CreatePlanWebsiteUserCmd.MarkFlagRequired("database-user")
	CreatePlanWebsiteUserCmd.MarkFlagRequired("password")
}

func createPlanWebsiteUserBody(cmd *cobra.Command) map[string]any {
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

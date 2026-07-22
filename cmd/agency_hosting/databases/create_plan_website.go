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

var CreatePlanWebsiteCmd = &cobra.Command{
	Use:   "create-plan-website <website_uid>",
	Short: "Create Agency Plan website database",
	Long:  "Creates a MySQL database with a dedicated user for an Agency Plan website.\n\nThe database name, username, and password must all be provided by the caller.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(createPlanWebsiteBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().AgencyHostingCreateAgencyPlanWebsiteDatabaseV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreatePlanWebsiteCmd.Flags().StringP("database-name", "", "", "Database name to create (alphanumeric characters).")
	CreatePlanWebsiteCmd.Flags().StringP("database-user", "", "", "Database username to create alongside the database (alphanumeric characters).")
	CreatePlanWebsiteCmd.Flags().StringP("password", "", "", "Password for the database user (requires mixed case, letters, and numbers).")
	CreatePlanWebsiteCmd.MarkFlagRequired("database-name")
	CreatePlanWebsiteCmd.MarkFlagRequired("database-user")
	CreatePlanWebsiteCmd.MarkFlagRequired("password")
}

func createPlanWebsiteBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	databaseNameVal, _ := cmd.Flags().GetString("database-name")
	body["database_name"] = databaseNameVal
	databaseUserVal, _ := cmd.Flags().GetString("database-user")
	body["database_user"] = databaseUserVal
	passwordVal, _ := cmd.Flags().GetString("password")
	body["password"] = passwordVal
	return body
}

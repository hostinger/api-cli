package databases

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateWebsiteCmd = &cobra.Command{
	Use:               "create-website <website_uid>",
	Short:             "Create website database",
	Long:              "Creates a MySQL database with a dedicated user for an Agency Plan website.\n\nThe database name, username, and password must all be provided by the caller.",
	Example:           "  hostinger agency-hosting databases create-website zpwlGlp19 --database-name mydatabase --database-user myuser --password 'Str0ngP@ssword!'",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := createWebsiteBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().AgencyHostingCreateWebsiteDatabaseV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		}
		respBody, status, err := call(cmd.Context())
		if err != nil {
			return err
		}
		return output.Format(cmd, respBody, status)
	},
}

func init() {
	utils.RegisterArgs(CreateWebsiteCmd, utils.StringArg("website_uid"))
	CreateWebsiteCmd.Flags().StringP("database-name", "", "", "Database name to create (alphanumeric characters).")
	CreateWebsiteCmd.Flags().StringP("database-user", "", "", "Database username to create alongside the database (alphanumeric characters).")
	CreateWebsiteCmd.Flags().StringP("password", "", "", "Password for the database user (requires mixed case, letters, and numbers).")
	CreateWebsiteCmd.MarkFlagRequired("database-name")
	CreateWebsiteCmd.MarkFlagRequired("database-user")
	CreateWebsiteCmd.MarkFlagRequired("password")
	CreateWebsiteCmd.RegisterFlagCompletionFunc("database-name", cobra.NoFileCompletions)
	CreateWebsiteCmd.RegisterFlagCompletionFunc("database-user", cobra.NoFileCompletions)
	CreateWebsiteCmd.RegisterFlagCompletionFunc("password", cobra.NoFileCompletions)
}

func createWebsiteBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	databaseNameVal, _ := cmd.Flags().GetString("database-name")
	body["database_name"] = databaseNameVal
	databaseUserVal, _ := cmd.Flags().GetString("database-user")
	body["database_user"] = databaseUserVal
	passwordVal, _ := cmd.Flags().GetString("password")
	body["password"] = passwordVal
	return body, nil
}

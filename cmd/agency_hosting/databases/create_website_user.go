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

var CreateWebsiteUserCmd = &cobra.Command{
	Use:               "create-website-user <website_uid> <database_name>",
	Short:             "Create website database user",
	Long:              "Creates a user for an existing database on an Agency Plan website.\n\nEach database supports a single non-system user; creating a user for a database that already has one fails.",
	Example:           "  hostinger agency-hosting databases create-website-user zpwlGlp19 my_database --database-user my_user --password 'Str0ngP@ssword!'",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := createWebsiteUserBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().AgencyHostingCreateWebsiteDatabaseUserV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(CreateWebsiteUserCmd, utils.StringArg("website_uid"), utils.StringArg("database_name"))
	CreateWebsiteUserCmd.Flags().StringP("database-user", "", "", "Database username to create (alphanumeric and underscores).")
	CreateWebsiteUserCmd.Flags().StringP("host", "", "", "Host the user connects from (IPv4, IPv6, % wildcard, or localhost). Defaults to localhost.")
	CreateWebsiteUserCmd.Flags().StringP("password", "", "", "Password for the database user (requires mixed case, letters, and numbers).")
	CreateWebsiteUserCmd.MarkFlagRequired("database-user")
	CreateWebsiteUserCmd.MarkFlagRequired("password")
	CreateWebsiteUserCmd.RegisterFlagCompletionFunc("database-user", cobra.NoFileCompletions)
	CreateWebsiteUserCmd.RegisterFlagCompletionFunc("host", cobra.NoFileCompletions)
	CreateWebsiteUserCmd.RegisterFlagCompletionFunc("password", cobra.NoFileCompletions)
}

func createWebsiteUserBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	databaseUserVal, _ := cmd.Flags().GetString("database-user")
	body["database_user"] = databaseUserVal
	if cmd.Flags().Changed("host") {
		v, _ := cmd.Flags().GetString("host")
		body["host"] = v
	}
	passwordVal, _ := cmd.Flags().GetString("password")
	body["password"] = passwordVal
	return body, nil
}

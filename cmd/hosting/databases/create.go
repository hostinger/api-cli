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

var CreateCmd = &cobra.Command{
	Use:     "create <username>",
	Short:   "Create account database",
	Long:    "Creates a database with a database user and password for the specified account.\n\nThe database name and user are automatically prefixed with the account username when needed.",
	Example: "  hostinger hosting databases create u123456789 --name u123456789_test_db --password 'Str0ngP@ssword!' --user u123456789_admin --website-domain example.com",
	ValidArgsFunction: utils.CompleteArgs(
		&utils.ArgSource{
			Value: "username",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().HostingListWebsitesV1WithResponse(ctx, nil)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := createBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingCreateAccountDatabaseV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(CreateCmd, utils.StringArg("username"))
	CreateCmd.Flags().StringP("name", "", "", "Database name. If the account username prefix is omitted, it is added automatically.")
	CreateCmd.Flags().StringP("password", "", "", "Database user password.")
	CreateCmd.Flags().StringP("user", "", "", "Database user. If the account username prefix is omitted, it is added automatically.")
	CreateCmd.Flags().StringP("website-domain", "", "", "Website domain assigned to the database.")
	CreateCmd.MarkFlagRequired("name")
	CreateCmd.MarkFlagRequired("password")
	CreateCmd.MarkFlagRequired("user")
	CreateCmd.MarkFlagRequired("website-domain")
	CreateCmd.RegisterFlagCompletionFunc("name", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("password", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("user", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("website-domain", cobra.NoFileCompletions)
}

func createBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	nameVal, _ := cmd.Flags().GetString("name")
	body["name"] = nameVal
	passwordVal, _ := cmd.Flags().GetString("password")
	body["password"] = passwordVal
	userVal, _ := cmd.Flags().GetString("user")
	body["user"] = userVal
	websiteDomainVal, _ := cmd.Flags().GetString("website-domain")
	body["website_domain"] = websiteDomainVal
	return body, nil
}

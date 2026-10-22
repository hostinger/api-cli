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

var ChangePasswordCmd = &cobra.Command{
	Use:     "change-password <username> <name>",
	Short:   "Change database password",
	Long:    "Changes the password for the specified database user.\n\nThe database name must be the full name returned by the list databases endpoint.\nThe password must also be updated in any website configuration that uses this database.",
	Example: "  hostinger hosting databases change-password u123456789 u123456789_test_db --password 'Secu4ePa$$wor!D'",
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
		nil,
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := changePasswordBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().HostingChangeDatabasePasswordV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(ChangePasswordCmd, utils.StringArg("username"), utils.StringArg("name"))
	ChangePasswordCmd.Flags().StringP("password", "", "", "New database user password.")
	ChangePasswordCmd.MarkFlagRequired("password")
	ChangePasswordCmd.RegisterFlagCompletionFunc("password", cobra.NoFileCompletions)
}

func changePasswordBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	passwordVal, _ := cmd.Flags().GetString("password")
	body["password"] = passwordVal
	return body, nil
}

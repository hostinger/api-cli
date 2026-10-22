package mailboxes

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ChangeMailboxPasswordCmd = &cobra.Command{
	Use:               "change-mailbox-password <mailbox-id>",
	Short:             "Change mailbox password",
	Long:              "Change the password of a mailbox.",
	Example:           "  hostinger mail mailboxes change-mailbox-password AC1a2b3c4d5e6f7g --password 'SecurePassword123!'",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := changeMailboxPasswordBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().MailChangeMailboxPasswordV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(ChangeMailboxPasswordCmd, utils.StringArg("mailbox-id"))
	ChangeMailboxPasswordCmd.Flags().StringP("password", "", "", "New mailbox password. Minimum 8 characters with uppercase, lowercase, number and special character; must not be a commonly used password.")
	ChangeMailboxPasswordCmd.MarkFlagRequired("password")
	ChangeMailboxPasswordCmd.RegisterFlagCompletionFunc("password", cobra.NoFileCompletions)
}

func changeMailboxPasswordBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	passwordVal, _ := cmd.Flags().GetString("password")
	body["password"] = passwordVal
	return body, nil
}

package mailboxes

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateMailboxCmd = &cobra.Command{
	Use:     "create-mailbox [order-id]",
	Short:   "Create mailbox",
	Long:    "Create a mailbox under the given mail order. The full email address is\ncomposed from the given local part and the domain of the order.\n\nThe order can be given as [order-id] or left out, in which case it comes from --order-id, HOSTINGER_ORDER_ID, the .hostinger/site.json written by `hostinger up` or `hostinger link` (hosting orders), or a prompt when run in a terminal (the only order is used without asking).",
	Example: "  hostinger mail mailboxes create-mailbox OR1a2b3c4d5e6f7g --local-part john.doe --password 'SecurePassword123!'",
	ValidArgsFunction: utils.CompleteArgs(&utils.ArgSource{
		Value: "id",
		Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
			r, err := api.CompletionRequest().MailListOrdersV1WithResponse(ctx, nil)
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		},
	}),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		orderID, args, err := site.ResolveOrder(cmd, args, 0, "mail")
		if err != nil {
			return api.CommandError(err)
		}
		body, err := createMailboxBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().MailCreateMailboxV1WithBodyWithResponse(ctx, orderID, "application/json", bytes.NewReader(payload))
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
	utils.RegisterScopedArgs(CreateMailboxCmd, 1, utils.StringArg("order-id"))
	CreateMailboxCmd.Flags().StringP("local-part", "", "", "Local part of the mailbox address (the part before the @). The domain is taken from the order. Must start and end with a letter or digit; single dots, underscores and hyphens are allowed in between.")
	CreateMailboxCmd.Flags().StringP("password", "", "", "Mailbox password. Minimum 8 characters with uppercase, lowercase, number and special character.")
	CreateMailboxCmd.MarkFlagRequired("local-part")
	CreateMailboxCmd.MarkFlagRequired("password")
	CreateMailboxCmd.RegisterFlagCompletionFunc("local-part", cobra.NoFileCompletions)
	CreateMailboxCmd.RegisterFlagCompletionFunc("password", cobra.NoFileCompletions)
}

func createMailboxBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	localPartVal, _ := cmd.Flags().GetString("local-part")
	body["local_part"] = localPartVal
	passwordVal, _ := cmd.Flags().GetString("password")
	body["password"] = passwordVal
	return body, nil
}

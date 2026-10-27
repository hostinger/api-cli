package api_tokens

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

var CreateCmd = &cobra.Command{
	Use:     "create [order-id]",
	Short:   "Create API token",
	Long:    "Create an API token for the given mail order. The token grants access\nto the Hostinger Email API (https://api.mail.hostinger.com/), where\nyou can provision and manage the mailboxes it is scoped to.\n\nThe plaintext token is returned only in this response, never again.\nA maximum of 10 tokens can exist per order. Use\n`scope.has_all_mailboxes` to cover all current and future mailboxes,\nor list specific mailboxes in `scope.mailbox_ids`.\n\nThe order can be given as [order-id] or left out, in which case it comes from --order-id, HOSTINGER_ORDER_ID, the .hostinger/site.json written by `hostinger up` or `hostinger link` (hosting orders), or a prompt when run in a terminal (the only order is used without asking).\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --scope\n    {\n      \"has_all_mailboxes\": false,\n      \"mailbox_ids\": [\n        \"AC1a2b3c4d5e6f7g\"\n      ]\n    }",
	Example: "  hostinger mail api-tokens create OR1a2b3c4d5e6f7g --name 'CRM integration' --scope @scope.json",
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
		body, err := createBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().MailCreateAPITokenV1WithBodyWithResponse(ctx, orderID, "application/json", bytes.NewReader(payload))
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
	utils.RegisterScopedArgs(CreateCmd, 1, utils.StringArg("order-id"))
	CreateCmd.Flags().StringP("name", "", "", "Human-readable label for this token")
	CreateCmd.Flags().StringP("scope", "", "", "Mailbox scope this token can access (JSON, @file or -; shape in the help above)")
	CreateCmd.MarkFlagRequired("name")
	CreateCmd.MarkFlagRequired("scope")
	CreateCmd.RegisterFlagCompletionFunc("name", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("scope", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
}

func createBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	nameVal, _ := cmd.Flags().GetString("name")
	body["name"] = nameVal
	scopeVal, _ := cmd.Flags().GetString("scope")
	scopeValJSON, err := utils.JSONValue(scopeVal, "scope")
	if err != nil {
		return nil, err
	}
	body["scope"] = scopeValJSON
	return body, nil
}

package api_tokens

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create <order-id>",
	Short: "Create API token",
	Long:  "Create an API token for the given mail order. The token grants access\nto the [Hostinger Email API](https://api.mail.hostinger.com/), where\nyou can provision and manage the mailboxes it is scoped to.\n\nThe plaintext token is returned only in this response, never again.\nA maximum of 10 tokens can exist per order. Use\n`scope.has_all_mailboxes` to cover all current and future mailboxes,\nor list specific mailboxes in `scope.mailbox_ids`.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(createBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().MailCreateAPITokenV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateCmd.Flags().StringP("name", "", "", "Human-readable label for this token")
	CreateCmd.Flags().StringP("scope", "", "", "Mailbox scope this token can access (JSON)")
	CreateCmd.MarkFlagRequired("name")
	CreateCmd.MarkFlagRequired("scope")
}

func createBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	nameVal, _ := cmd.Flags().GetString("name")
	body["name"] = nameVal
	scopeVal, _ := cmd.Flags().GetString("scope")
	body["scope"] = utils.JSONValue(scopeVal, "scope")
	return body
}

package mailboxes

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var CreateMailboxCmd = &cobra.Command{
	Use:   "create-mailbox <order-id>",
	Short: "Create mailbox",
	Long:  "Create a mailbox under the given mail order. The full email address is\ncomposed from the given local part and the domain of the order.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(createMailboxBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().MailCreateMailboxV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateMailboxCmd.Flags().StringP("local-part", "", "", "Local part of the mailbox address (the part before the @). The domain is taken from the order. Must start and end with a letter or digit; single dots, underscores and hyphens are allowed in between.")
	CreateMailboxCmd.Flags().StringP("password", "", "", "Mailbox password. Minimum 8 characters with uppercase, lowercase, number and special character.")
	CreateMailboxCmd.MarkFlagRequired("local-part")
	CreateMailboxCmd.MarkFlagRequired("password")
}

func createMailboxBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	localPartVal, _ := cmd.Flags().GetString("local-part")
	body["local_part"] = localPartVal
	passwordVal, _ := cmd.Flags().GetString("password")
	body["password"] = passwordVal
	return body
}

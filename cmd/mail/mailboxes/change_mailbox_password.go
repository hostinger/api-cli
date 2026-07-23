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

var ChangeMailboxPasswordCmd = &cobra.Command{
	Use:   "change-mailbox-password <mailbox-id>",
	Short: "Change mailbox password",
	Long:  "Change the password of a mailbox.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(changeMailboxPasswordBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().MailChangeMailboxPasswordV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ChangeMailboxPasswordCmd.Flags().StringP("password", "", "", "New mailbox password. Minimum 8 characters with uppercase, lowercase, number and special character; must not be a commonly used password.")
	ChangeMailboxPasswordCmd.MarkFlagRequired("password")
}

func changeMailboxPasswordBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	passwordVal, _ := cmd.Flags().GetString("password")
	body["password"] = passwordVal
	return body
}

package autoreplies

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create <mailbox-id>",
	Short: "Create autoreply",
	Long:  "Create an automatic reply for the given mailbox. A mailbox can have\nonly one autoreply. Omit `starts_at` to activate the autoreply\nimmediately and omit `ends_at` to keep it active indefinitely.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(createBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().MailCreateAutoreplyV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateCmd.Flags().StringP("body", "", "", "Body of the automatic reply")
	CreateCmd.Flags().StringP("display-name", "", "", "Sender display name used for the reply")
	CreateCmd.Flags().StringP("ends-at", "", "", "When the autoreply stops. Omit for an indefinite autoreply.")
	CreateCmd.Flags().StringP("starts-at", "", "", "When the autoreply becomes active. Defaults to now.")
	CreateCmd.Flags().StringP("subject", "", "", "Subject of the automatic reply")
	CreateCmd.MarkFlagRequired("body")
	CreateCmd.MarkFlagRequired("subject")
}

func createBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	bodyVal, _ := cmd.Flags().GetString("body")
	body["body"] = bodyVal
	if cmd.Flags().Changed("display-name") {
		v, _ := cmd.Flags().GetString("display-name")
		body["display_name"] = v
	}
	if cmd.Flags().Changed("ends-at") {
		v, _ := cmd.Flags().GetString("ends-at")
		body["ends_at"] = v
	}
	if cmd.Flags().Changed("starts-at") {
		v, _ := cmd.Flags().GetString("starts-at")
		body["starts_at"] = v
	}
	subjectVal, _ := cmd.Flags().GetString("subject")
	body["subject"] = subjectVal
	return body
}

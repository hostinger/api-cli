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

var UpdateCmd = &cobra.Command{
	Use:   "update <autoreply-id>",
	Short: "Update autoreply",
	Long:  "Replace the autoreply with the given content and schedule. Omitted\noptional fields are cleared: omit `starts_at` to activate the\nautoreply immediately and omit `ends_at` to keep it active\nindefinitely.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(updateBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().MailUpdateAutoreplyV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	UpdateCmd.Flags().StringP("body", "", "", "Body of the automatic reply")
	UpdateCmd.Flags().StringP("display-name", "", "", "Sender display name used for the reply")
	UpdateCmd.Flags().StringP("ends-at", "", "", "When the autoreply stops. Omit for an indefinite autoreply.")
	UpdateCmd.Flags().StringP("starts-at", "", "", "When the autoreply becomes active. Defaults to now.")
	UpdateCmd.Flags().StringP("subject", "", "", "Subject of the automatic reply")
	UpdateCmd.MarkFlagRequired("body")
	UpdateCmd.MarkFlagRequired("subject")
}

func updateBody(cmd *cobra.Command) map[string]any {
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

package webhooks

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

var UpdateCmd = &cobra.Command{
	Use:   "update <webhook-id>",
	Short: "Update webhook",
	Long:  "Partially update a webhook. Only the fields included in the request\nbody are changed; omitted fields retain their current values. Pass\n`\"description\": null` to clear the description.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "status", []string{"active", "disabled", "paused"})
		payload, err := json.Marshal(updateBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().MailUpdateWebhookV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	UpdateCmd.Flags().StringP("description", "", "", "New description, or null to clear it")
	UpdateCmd.Flags().StringSliceP("events", "", nil, "Replaces the full list of subscribed events (one of: message.received)")
	UpdateCmd.Flags().StringP("name", "", "", "New human-readable name for the webhook")
	UpdateCmd.Flags().StringP("status", "", "", "New status for the webhook (one of: active, disabled, paused)")
	UpdateCmd.Flags().StringP("url", "", "", "New URL to deliver events to")
}

func updateBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("description") {
		v, _ := cmd.Flags().GetString("description")
		body["description"] = v
	}
	if cmd.Flags().Changed("events") {
		v, _ := cmd.Flags().GetStringSlice("events")
		body["events"] = v
	}
	if cmd.Flags().Changed("name") {
		v, _ := cmd.Flags().GetString("name")
		body["name"] = v
	}
	if cmd.Flags().Changed("status") {
		v, _ := cmd.Flags().GetString("status")
		body["status"] = v
	}
	if cmd.Flags().Changed("url") {
		v, _ := cmd.Flags().GetString("url")
		body["url"] = v
	}
	return body
}

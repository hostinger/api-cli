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

var CreateCmd = &cobra.Command{
	Use:   "create <mailbox-id>",
	Short: "Create webhook",
	Long:  "Create a webhook for the given mailbox. The generated secret is\nreturned only in this response and is sent as a bearer token with\nevery delivery.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		utils.EnumCheck(cmd, "status", []string{"active", "disabled", "paused"})
		payload, err := json.Marshal(createBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().MailCreateWebhookV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateCmd.Flags().StringP("description", "", "", "Optional description of the webhook's purpose")
	CreateCmd.Flags().StringSliceP("events", "", nil, "Events that trigger this webhook (one of: message.received)")
	CreateCmd.Flags().StringP("name", "", "", "Human-readable name for this webhook")
	CreateCmd.Flags().StringP("status", "", "active", "Initial status of the webhook (one of: active, disabled, paused)")
	CreateCmd.Flags().StringP("url", "", "", "Publicly reachable URL that receives the webhook POST requests")
	CreateCmd.MarkFlagRequired("events")
	CreateCmd.MarkFlagRequired("name")
	CreateCmd.MarkFlagRequired("url")
}

func createBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("description") {
		v, _ := cmd.Flags().GetString("description")
		body["description"] = v
	}
	eventsVal, _ := cmd.Flags().GetStringSlice("events")
	body["events"] = eventsVal
	nameVal, _ := cmd.Flags().GetString("name")
	body["name"] = nameVal
	if cmd.Flags().Changed("status") {
		v, _ := cmd.Flags().GetString("status")
		body["status"] = v
	}
	urlVal, _ := cmd.Flags().GetString("url")
	body["url"] = urlVal
	return body
}

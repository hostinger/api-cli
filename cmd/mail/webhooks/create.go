package webhooks

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:               "create <mailbox-id>",
	Short:             "Create webhook",
	Long:              "Create a webhook for the given mailbox. The generated secret is\nreturned only in this response and is sent as a bearer token with\nevery delivery.",
	Example:           "  hostinger mail webhooks create AC1a2b3c4d5e6f7g --events message.received --name 'New message notifier' --url https://example.com/webhooks/incoming",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "status", []string{"active", "disabled", "paused"}); err != nil {
			return err
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
			r, err := api.Request().MailCreateWebhookV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(CreateCmd, utils.StringArg("mailbox-id"))
	CreateCmd.Flags().StringP("description", "", "", "Optional description of the webhook's purpose")
	CreateCmd.Flags().StringSliceP("events", "", nil, "Events that trigger this webhook (one of: message.received)")
	CreateCmd.Flags().StringP("name", "", "", "Human-readable name for this webhook")
	CreateCmd.Flags().StringP("status", "", "active", "Initial status of the webhook (one of: active, disabled, paused)")
	CreateCmd.Flags().StringP("url", "", "", "Publicly reachable URL that receives the webhook POST requests")
	CreateCmd.MarkFlagRequired("events")
	CreateCmd.MarkFlagRequired("name")
	CreateCmd.MarkFlagRequired("url")
	CreateCmd.RegisterFlagCompletionFunc("description", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("events", cobra.FixedCompletions([]cobra.Completion{"message.received"}, cobra.ShellCompDirectiveNoFileComp))
	CreateCmd.RegisterFlagCompletionFunc("name", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("status", cobra.FixedCompletions([]cobra.Completion{"active", "disabled", "paused"}, cobra.ShellCompDirectiveNoFileComp))
	CreateCmd.RegisterFlagCompletionFunc("url", cobra.NoFileCompletions)
}

func createBody(cmd *cobra.Command) (map[string]any, error) {
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
	return body, nil
}

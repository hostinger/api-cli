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

var UpdateCmd = &cobra.Command{
	Use:               "update <webhook-id>",
	Short:             "Update webhook",
	Long:              "Partially update a webhook. Only the fields included in the request\nbody are changed; omitted fields retain their current values. Pass\n`\"description\": null` to clear the description.",
	Example:           "  hostinger mail webhooks update 019683f8-1234-7abc-8def-0123456789ab",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.EnumCheck(cmd, "status", []string{"active", "disabled", "paused"}); err != nil {
			return err
		}
		body, err := updateBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().MailUpdateWebhookV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(UpdateCmd, utils.StringArg("webhook-id"))
	UpdateCmd.Flags().StringP("description", "", "", "New description, or null to clear it")
	UpdateCmd.Flags().StringSliceP("events", "", nil, "Replaces the full list of subscribed events (one of: message.received)")
	UpdateCmd.Flags().StringP("name", "", "", "New human-readable name for the webhook")
	UpdateCmd.Flags().StringP("status", "", "", "New status for the webhook (one of: active, disabled, paused)")
	UpdateCmd.Flags().StringP("url", "", "", "New URL to deliver events to")
	UpdateCmd.RegisterFlagCompletionFunc("description", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("events", cobra.FixedCompletions([]cobra.Completion{"message.received"}, cobra.ShellCompDirectiveNoFileComp))
	UpdateCmd.RegisterFlagCompletionFunc("name", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("status", cobra.FixedCompletions([]cobra.Completion{"active", "disabled", "paused"}, cobra.ShellCompDirectiveNoFileComp))
	UpdateCmd.RegisterFlagCompletionFunc("url", cobra.NoFileCompletions)
}

func updateBody(cmd *cobra.Command) (map[string]any, error) {
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
	return body, nil
}

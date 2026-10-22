package autoreplies

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
	Short:             "Create autoreply",
	Long:              "Create an automatic reply for the given mailbox. A mailbox can have\nonly one autoreply. Omit `starts_at` to activate the autoreply\nimmediately and omit `ends_at` to keep it active indefinitely.",
	Example:           "  hostinger mail autoreplies create AC1a2b3c4d5e6f7g --body 'I am on vacation until August 1st.' --subject 'Out of office'",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := createBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().MailCreateAutoreplyV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	CreateCmd.Flags().StringP("body", "", "", "Body of the automatic reply")
	CreateCmd.Flags().StringP("display-name", "", "", "Sender display name used for the reply")
	CreateCmd.Flags().StringP("ends-at", "", "", "When the autoreply stops. Omit for an indefinite autoreply.")
	CreateCmd.Flags().StringP("starts-at", "", "", "When the autoreply becomes active. Defaults to now.")
	CreateCmd.Flags().StringP("subject", "", "", "Subject of the automatic reply")
	CreateCmd.MarkFlagRequired("body")
	CreateCmd.MarkFlagRequired("subject")
	CreateCmd.RegisterFlagCompletionFunc("body", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("display-name", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("ends-at", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("starts-at", cobra.NoFileCompletions)
	CreateCmd.RegisterFlagCompletionFunc("subject", cobra.NoFileCompletions)
}

func createBody(cmd *cobra.Command) (map[string]any, error) {
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
	return body, nil
}

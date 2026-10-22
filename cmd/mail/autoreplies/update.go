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

var UpdateCmd = &cobra.Command{
	Use:               "update <autoreply-id>",
	Short:             "Update autoreply",
	Long:              "Replace the autoreply with the given content and schedule. Omitted\noptional fields are cleared: omit `starts_at` to activate the\nautoreply immediately and omit `ends_at` to keep it active\nindefinitely.",
	Example:           "  hostinger mail autoreplies update AR1a2b3c4d5e6f7g --body 'I am on vacation until August 1st.' --subject 'Out of office'",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := updateBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().MailUpdateAutoreplyV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(UpdateCmd, utils.StringArg("autoreply-id"))
	UpdateCmd.Flags().StringP("body", "", "", "Body of the automatic reply")
	UpdateCmd.Flags().StringP("display-name", "", "", "Sender display name used for the reply")
	UpdateCmd.Flags().StringP("ends-at", "", "", "When the autoreply stops. Omit for an indefinite autoreply.")
	UpdateCmd.Flags().StringP("starts-at", "", "", "When the autoreply becomes active. Defaults to now.")
	UpdateCmd.Flags().StringP("subject", "", "", "Subject of the automatic reply")
	UpdateCmd.MarkFlagRequired("body")
	UpdateCmd.MarkFlagRequired("subject")
	UpdateCmd.RegisterFlagCompletionFunc("body", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("display-name", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("ends-at", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("starts-at", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("subject", cobra.NoFileCompletions)
}

func updateBody(cmd *cobra.Command) (map[string]any, error) {
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

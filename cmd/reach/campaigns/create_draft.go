package campaigns

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateDraftCmd = &cobra.Command{
	Use:               "create-draft <profile-uuid>",
	Short:             "Create a draft campaign",
	Long:              "Create a campaign in a profile.\n\nThe campaign is created as a draft, so nothing is sent and no contact is touched. It has no\naudience yet either - targeting and scheduling are not part of this request, the draft is\nfinished and sent from the Reach interface.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --metadata\n    {\n      \"preheader\": \"Our biggest deals of the year\",\n      \"source\": \"api\"\n    }",
	Example:           "  hostinger reach campaigns create-draft 550e8400-e09b-41d4-a716-400055000000 --sender-email marketing@example.com --sender-name 'Marketing Team'",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := createDraftBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachCreateADraftCampaignV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(CreateDraftCmd, utils.StringArg("profile-uuid"))
	CreateDraftCmd.Flags().StringP("metadata", "", "", "Extra campaign fields. Any key outside the listed ones is rejected. (JSON, @file or -; shape in the help above)")
	CreateDraftCmd.Flags().StringP("sender-email", "", "", "From address of the campaign. Its domain has to be verified on the profile before the campaign can be sent.")
	CreateDraftCmd.Flags().StringP("sender-name", "", "", "From name shown to the recipients.")
	CreateDraftCmd.Flags().StringP("subject", "", "", "Subject line of the email.")
	CreateDraftCmd.Flags().StringP("template-uuid", "", "", "Template to send, as returned by the template endpoints. Can be left out and attached later, but the campaign cannot be sent without one.")
	CreateDraftCmd.Flags().StringP("title", "", "", "Name the campaign is listed under. Not shown to the recipients.")
	CreateDraftCmd.MarkFlagRequired("sender-email")
	CreateDraftCmd.MarkFlagRequired("sender-name")
	CreateDraftCmd.RegisterFlagCompletionFunc("metadata", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	CreateDraftCmd.RegisterFlagCompletionFunc("sender-email", cobra.NoFileCompletions)
	CreateDraftCmd.RegisterFlagCompletionFunc("sender-name", cobra.NoFileCompletions)
	CreateDraftCmd.RegisterFlagCompletionFunc("subject", cobra.NoFileCompletions)
	CreateDraftCmd.RegisterFlagCompletionFunc("template-uuid", cobra.NoFileCompletions)
	CreateDraftCmd.RegisterFlagCompletionFunc("title", cobra.NoFileCompletions)
}

func createDraftBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("metadata") {
		metadataVal, _ := cmd.Flags().GetString("metadata")
		metadataValJSON, err := utils.JSONValue(metadataVal, "metadata")
		if err != nil {
			return nil, err
		}
		body["metadata"] = metadataValJSON
	}
	senderEmailVal, _ := cmd.Flags().GetString("sender-email")
	body["sender_email"] = senderEmailVal
	senderNameVal, _ := cmd.Flags().GetString("sender-name")
	body["sender_name"] = senderNameVal
	if cmd.Flags().Changed("subject") {
		v, _ := cmd.Flags().GetString("subject")
		body["subject"] = v
	}
	if cmd.Flags().Changed("template-uuid") {
		v, _ := cmd.Flags().GetString("template-uuid")
		body["template_uuid"] = v
	}
	if cmd.Flags().Changed("title") {
		v, _ := cmd.Flags().GetString("title")
		body["title"] = v
	}
	return body, nil
}

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

var UpdateCmd = &cobra.Command{
	Use:     "update <profile-uuid> <campaign-uuid>",
	Short:   "Update campaign",
	Long:    "Update a campaign's content, sender, template and audience.\n\nOnly `draft`, `scheduled` and `failed` campaigns can be updated; a campaign that is sending or\nalready sent is rejected. Updating a `failed` campaign moves it back to `draft`. Audience changes\nreplace rather than merge: `all_contacts: true` targets every contact of the profile, while\n`segment_uuids` replaces the segments currently targeted. Leave both out to keep the audience as\nit is. This does not send the campaign - use the send endpoint for that.\n\nJSON flags take inline JSON, @file.json, or - to read stdin. Example values:\n\n  --metadata\n    {\n      \"preheader\": \"Our biggest deals of the year\",\n      \"source\": \"api\"\n    }",
	Example: "  hostinger reach campaigns update 550e8400-e09b-41d4-a716-400055000000 550e8400-e09b-41d4-a716-400055000000",
	ValidArgsFunction: utils.CompleteArgs(
		nil,
		&utils.ArgSource{
			Value: "uuid",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().ReachListCampaignsV1WithResponse(ctx, args[0], nil)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
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
			r, err := api.Request().ReachUpdateCampaignV1WithBodyWithResponse(ctx, args[0], args[1], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(UpdateCmd, utils.StringArg("profile-uuid"), utils.StringArg("campaign-uuid"))
	UpdateCmd.Flags().BoolP("all-contacts", "", false, "`true` targets every contact of the profile, `false` targets only `segment_uuids`. Required when `segment_uuids` or `excluded_segment_uuids` is sent.")
	UpdateCmd.Flags().StringSliceP("excluded-segment-uuids", "", nil, "Segments to leave out. Only allowed together with `all_contacts: true`.")
	UpdateCmd.Flags().StringP("metadata", "", "", "Extra campaign fields. Any key outside the listed ones is rejected. (JSON, @file or -; shape in the help above)")
	UpdateCmd.Flags().StringSliceP("segment-uuids", "", nil, "Segments to target. Replaces the segments currently targeted.")
	UpdateCmd.Flags().StringP("sender-email", "", "", "From address of the campaign. Its domain has to be verified on the profile before the campaign can be sent.")
	UpdateCmd.Flags().StringP("sender-name", "", "", "From name shown to the recipients.")
	UpdateCmd.Flags().StringP("subject", "", "", "Subject line of the email.")
	UpdateCmd.Flags().StringP("template-uuid", "", "", "Template to send, as returned by the template endpoints.")
	UpdateCmd.Flags().StringP("title", "", "", "Name the campaign is listed under. Not shown to the recipients.")
	UpdateCmd.RegisterFlagCompletionFunc("all-contacts", cobra.FixedCompletions([]cobra.Completion{"true", "false"}, cobra.ShellCompDirectiveNoFileComp))
	UpdateCmd.RegisterFlagCompletionFunc("excluded-segment-uuids", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("metadata", cobra.FixedCompletions(nil, cobra.ShellCompDirectiveDefault))
	UpdateCmd.RegisterFlagCompletionFunc("segment-uuids", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("sender-email", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("sender-name", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("subject", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("template-uuid", cobra.NoFileCompletions)
	UpdateCmd.RegisterFlagCompletionFunc("title", cobra.NoFileCompletions)
}

func updateBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	if cmd.Flags().Changed("all-contacts") {
		v, _ := cmd.Flags().GetBool("all-contacts")
		body["all_contacts"] = v
	}
	if cmd.Flags().Changed("excluded-segment-uuids") {
		v, _ := cmd.Flags().GetStringSlice("excluded-segment-uuids")
		body["excluded_segment_uuids"] = v
	}
	if cmd.Flags().Changed("metadata") {
		metadataVal, _ := cmd.Flags().GetString("metadata")
		metadataValJSON, err := utils.JSONValue(metadataVal, "metadata")
		if err != nil {
			return nil, err
		}
		body["metadata"] = metadataValJSON
	}
	if cmd.Flags().Changed("segment-uuids") {
		v, _ := cmd.Flags().GetStringSlice("segment-uuids")
		body["segment_uuids"] = v
	}
	if cmd.Flags().Changed("sender-email") {
		v, _ := cmd.Flags().GetString("sender-email")
		body["sender_email"] = v
	}
	if cmd.Flags().Changed("sender-name") {
		v, _ := cmd.Flags().GetString("sender-name")
		body["sender_name"] = v
	}
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

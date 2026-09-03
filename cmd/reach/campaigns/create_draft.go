package campaigns

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

var CreateDraftCmd = &cobra.Command{
	Use:   "create-draft <profile-uuid>",
	Short: "Create a draft campaign",
	Long:  "Create a campaign in a profile.\n\nThe campaign is created as a draft, so nothing is sent and no contact is touched. It has no\naudience yet either - targeting and scheduling are not part of this request, the draft is\nfinished and sent from the Reach interface.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(createDraftBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().ReachCreateADraftCampaignV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateDraftCmd.Flags().StringP("metadata", "", "", "Extra campaign fields. Any key outside the listed ones is rejected. (JSON)")
	CreateDraftCmd.Flags().StringP("sender-email", "", "", "From address of the campaign. Its domain has to be verified on the profile before\nthe campaign can be sent.")
	CreateDraftCmd.Flags().StringP("sender-name", "", "", "From name shown to the recipients.")
	CreateDraftCmd.Flags().StringP("subject", "", "", "Subject line of the email.")
	CreateDraftCmd.Flags().StringP("template-uuid", "", "", "Template to send, as returned by the template endpoints. Can be left out and\nattached later, but the campaign cannot be sent without one.")
	CreateDraftCmd.Flags().StringP("title", "", "", "Name the campaign is listed under. Not shown to the recipients.")
	CreateDraftCmd.MarkFlagRequired("sender-email")
	CreateDraftCmd.MarkFlagRequired("sender-name")
}

func createDraftBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	if cmd.Flags().Changed("metadata") {
		v, _ := cmd.Flags().GetString("metadata")
		body["metadata"] = utils.JSONValue(v, "metadata")
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
	return body
}

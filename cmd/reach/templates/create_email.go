package templates

import (
	"bytes"
	"context"
	"encoding/json"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var CreateEmailCmd = &cobra.Command{
	Use:   "create-email <profile-uuid>",
	Short: "Create an email template",
	Long:  "Create an email template in a profile.\n\nThe template holds the HTML body a campaign reuses, so it can be created before any\ncampaign exists. Only the template metadata comes back - keep the returned `uuid` to\nreference it as the `template_uuid` of a campaign.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		payload, err := json.Marshal(createEmailBody(cmd))
		if err != nil {
			log.Fatal(err)
		}
		r, err := api.Request().ReachCreateAnEmailTemplateV1WithBodyWithResponse(context.TODO(), args[0], "application/json", bytes.NewReader(payload))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	CreateEmailCmd.Flags().StringP("template-content", "", "", "The email body as HTML. It is sanitised before it is stored, so the saved template\ncan differ from what was sent - inline any styles the email clients need and keep\nthe markup self-contained.")
	CreateEmailCmd.Flags().StringP("title", "", "", "Name the template is listed under. Not shown to the recipients.")
	CreateEmailCmd.MarkFlagRequired("template-content")
}

func createEmailBody(cmd *cobra.Command) map[string]any {
	body := map[string]any{}
	templateContentVal, _ := cmd.Flags().GetString("template-content")
	body["template_content"] = templateContentVal
	if cmd.Flags().Changed("title") {
		v, _ := cmd.Flags().GetString("title")
		body["title"] = v
	}
	return body
}

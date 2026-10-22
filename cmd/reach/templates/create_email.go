package templates

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateEmailCmd = &cobra.Command{
	Use:               "create-email <profile-uuid>",
	Short:             "Create an email template",
	Long:              "Create an email template in a profile.\n\nThe template holds the HTML body a campaign reuses, so it can be created before any\ncampaign exists. Only the template metadata comes back - keep the returned `uuid` to\nreference it as the `template_uuid` of a campaign.",
	Example:           "  hostinger reach templates create-email 550e8400-e09b-41d4-a716-400055000000 --template-content '<html><body><h1>Email Template</h1><p>Hello, traveler!</p></body></html>'",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		body, err := createEmailBody(cmd)
		if err != nil {
			return err
		}
		payload, err := json.Marshal(body)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachCreateAnEmailTemplateV1WithBodyWithResponse(ctx, args[0], "application/json", bytes.NewReader(payload))
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
	utils.RegisterArgs(CreateEmailCmd, utils.StringArg("profile-uuid"))
	CreateEmailCmd.Flags().StringP("template-content", "", "", "The email body as HTML. It is sanitised before it is stored, so the saved template can differ from what was sent - inline any styles the email clients need and keep the markup self-contained.")
	CreateEmailCmd.Flags().StringP("title", "", "", "Name the template is listed under. Not shown to the recipients.")
	CreateEmailCmd.MarkFlagRequired("template-content")
	CreateEmailCmd.RegisterFlagCompletionFunc("template-content", cobra.NoFileCompletions)
	CreateEmailCmd.RegisterFlagCompletionFunc("title", cobra.NoFileCompletions)
}

func createEmailBody(cmd *cobra.Command) (map[string]any, error) {
	body := map[string]any{}
	templateContentVal, _ := cmd.Flags().GetString("template-content")
	body["template_content"] = templateContentVal
	if cmd.Flags().Changed("title") {
		v, _ := cmd.Flags().GetString("title")
		body["title"] = v
	}
	return body, nil
}

package templates

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var EmailCmd = &cobra.Command{
	Use:               "email <profile-uuid> <template-uuid>",
	Short:             "Get email template",
	Long:              "Get one email template of the profile with its full HTML in `template_content`.\n\nA campaign's email content is its template: to read what a campaign sends, get the campaign\nand pass its `template_uuid` here. The HTML can be large, so fetch it only when the content\nitself is needed, never to browse or pick templates - the list email templates endpoint does\nthat without content. Personalization tags such as `{{name}}` come back unrendered.",
	Example:           "  hostinger reach templates email 550e8400-e09b-41d4-a716-400055000000 550e8400-e29b-41d4-a716-446655440000",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachGetEmailTemplateV1WithResponse(ctx, args[0], args[1])
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
	utils.RegisterArgs(EmailCmd, utils.StringArg("profile-uuid"), utils.StringArg("template-uuid"))
}

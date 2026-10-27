package templates

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListEmailCmd = &cobra.Command{
	Use:               "list-email <profile-uuid>",
	Short:             "List email templates",
	Long:              "Get a list of the email templates in a profile, most recently updated first.\n\nTemplates are the reusable email bodies a campaign is built from. The list is not paginated\nand only the metadata is returned, without the HTML content. Use the `uuid` of a template\nas the `template_uuid` when creating a campaign. To read one template's HTML use the get\nemail template endpoint, and only when the content itself is needed.",
	Example:           "  hostinger reach templates list-email 550e8400-e09b-41d4-a716-400055000000",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachListEmailTemplatesV1WithResponse(ctx, args[0])
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
	utils.RegisterArgs(ListEmailCmd, utils.StringArg("profile-uuid"))
}

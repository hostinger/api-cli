package templates

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListEmailCmd = &cobra.Command{
	Use:   "list-email <profile-uuid>",
	Short: "List email templates",
	Long:  "Get a list of the email templates in a profile, most recently updated first.\n\nTemplates are the reusable email bodies a campaign is built from. The list is not paginated\nand only the metadata is returned - the template content itself is not exposed. Use the\n`uuid` of a template as the `template_uuid` when creating a campaign.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().ReachListEmailTemplatesV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

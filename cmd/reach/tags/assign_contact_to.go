package tags

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var AssignContactToCmd = &cobra.Command{
	Use:   "assign-contact-to <profile-uuid> <tag-uuid> <contact-uuid>",
	Short: "Assign a contact to a tag",
	Long:  "Assign a tag to a single contact.\n\nUnlike the bulk endpoint this is applied immediately rather than queued. Assigning a tag\nthe contact already carries succeeds without duplicating it.",
	Args:  cobra.MatchAll(cobra.ExactArgs(3)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().ReachAssignAContactToATagV1WithResponse(context.TODO(), args[0], uuid.MustParse(args[1]), uuid.MustParse(args[2]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

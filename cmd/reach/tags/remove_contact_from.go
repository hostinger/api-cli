package tags

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var RemoveContactFromCmd = &cobra.Command{
	Use:   "remove-contact-from <profile-uuid> <tag-uuid> <contact-uuid>",
	Short: "Remove a contact from a tag",
	Long:  "Remove a tag from a single contact.\n\nUnlike the bulk endpoint this is applied immediately rather than queued. Neither the tag\nnor the contact is deleted.",
	Args:  cobra.MatchAll(cobra.ExactArgs(3)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().ReachRemoveAContactFromATagV1WithResponse(context.TODO(), args[0], uuid.MustParse(args[1]), uuid.MustParse(args[2]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

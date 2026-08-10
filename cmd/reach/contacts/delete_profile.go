package contacts

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DeleteProfileCmd = &cobra.Command{
	Use:   "delete-profile <profile-uuid> <contact-uuid>",
	Short: "Delete a profile contact",
	Long:  "Permanently delete a contact from a profile.\n\nThe contact is removed together with its custom field values and tag assignments.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().ReachDeleteAProfileContactV1WithResponse(context.TODO(), args[0], uuid.MustParse(args[1]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

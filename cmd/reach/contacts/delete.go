package contacts

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete <uuid>",
	Short: "Delete a contact",
	Long:  "Delete a contact with the specified UUID.\n\nThis endpoint permanently removes a contact from the email marketing system.\n\n**Deprecated.** This endpoint cannot target a profile, so it always falls back to the\nclient's default profile and cannot delete contacts of any other profile. Use\n`DELETE /api/reach/v1/profiles/{profileUuid}/contacts/{contactUuid}` instead.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().ReachDeleteAContactV1WithResponse(context.TODO(), uuid.MustParse(args[0]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

package contact_fields

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete <profile-uuid> <field-uuid>",
	Short: "Delete a contact field",
	Long:  "Delete a custom contact field.\n\nEvery value contacts hold for the field is deleted with it, and for the choice types so\nare its options. The contacts themselves are not affected.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().ReachDeleteAContactFieldV1WithResponse(context.TODO(), args[0], uuid.MustParse(args[1]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

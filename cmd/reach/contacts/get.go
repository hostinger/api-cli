package contacts

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:   "get <profile-uuid> <contact-uuid>",
	Short: "Get contact details",
	Long:  "Get the full details of a single contact.\n\nAlongside the contact's own attributes this returns the tags assigned to it and the\nvalues it holds for the profile's custom contact fields.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().ReachGetContactDetailsV1WithResponse(context.TODO(), args[0], uuid.MustParse(args[1]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

package contact_fields

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list <profile-uuid>",
	Short: "List contact fields",
	Long:  "Get the custom contact fields defined in a profile.\n\nCustom fields let you store your own attributes on contacts. The returned uuids are what\nyou pass to the contact update endpoint to set values, and choice fields also list the\noptions available to pick from.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().ReachListContactFieldsV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

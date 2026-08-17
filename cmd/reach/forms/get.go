package forms

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:   "get <profile-uuid> <form-uuid>",
	Short: "Get form details",
	Long:  "Get a single form with the URL of its hosted template and the tags it applies to the contacts\nit captures.\n\nThere is no ready-made embed snippet in the response - either serve the template HTML yourself\nor build your own embed around the form uuid.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().ReachGetFormDetailsV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

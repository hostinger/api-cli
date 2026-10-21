package forms

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DuplicateCmd = &cobra.Command{
	Use:   "duplicate <profile-uuid> <form-uuid>",
	Short: "Duplicate form",
	Long:  "Copy a form into a new draft form with the same name and design.\n\nThe copy shares the source's automatic `form:{name}` tag; other submitter tags are not copied. The\ncopy is not live until its status is set to `active`.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().ReachDuplicateFormV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

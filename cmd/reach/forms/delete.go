package forms

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete <profile-uuid> <form-uuid>",
	Short: "Delete form",
	Long:  "Permanently delete a form together with its template.\n\nA form that has already captured submissions cannot be deleted, so that the contacts it collected\nare never silently discarded - pause the form instead to stop it collecting new ones. Views alone\ndo not block deletion.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().ReachDeleteFormV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

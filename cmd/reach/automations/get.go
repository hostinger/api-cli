package automations

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:   "get <profile-uuid> <automation-uuid>",
	Short: "Get automation details",
	Long:  "Get a single automation with the counts of contacts that entered it, are moving through it,\nfinished it or failed on the way.\n\nThis describes the automation itself. To see the workflow it runs, use the steps endpoint.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().ReachGetAutomationDetailsV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

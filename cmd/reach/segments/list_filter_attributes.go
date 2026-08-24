package segments

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListFilterAttributesCmd = &cobra.Command{
	Use:   "list-filter-attributes <profile-uuid>",
	Short: "List segment filter attributes",
	Long:  "List every attribute a segment condition can filter on, with the operators each attribute\naccepts, the value format they expect and, where the value is constrained, the allowed\nvalues.\n\nThe list is profile specific: it includes the profile's custom contact fields, its tags and\nits 20 most recently published campaigns, so the valid attributes cannot be hardcoded. Read\nit before creating or updating a segment to discover the valid `attribute`, `operator` and\n`value` combinations.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().ReachListSegmentFilterAttributesV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

package segments

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ProfileCmd = &cobra.Command{
	Use:   "profile <profile-uuid> <segment-uuid>",
	Short: "Get profile segment details",
	Long:  "Get a single segment of a profile, including the conditions that define it.\n\nTo retrieve the contacts currently matching those conditions, use the segment contacts\nendpoint instead.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().ReachGetProfileSegmentDetailsV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

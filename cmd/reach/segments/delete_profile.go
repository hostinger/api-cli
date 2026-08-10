package segments

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var DeleteProfileCmd = &cobra.Command{
	Use:   "delete-profile <profile-uuid> <segment-uuid>",
	Short: "Delete a profile segment",
	Long:  "Delete a segment.\n\nOnly the segment definition is removed. The contacts that matched it are left untouched.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().ReachDeleteAProfileSegmentV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

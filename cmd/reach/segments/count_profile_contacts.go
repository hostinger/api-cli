package segments

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var CountProfileContactsCmd = &cobra.Command{
	Use:   "count-profile-contacts <profile-uuid> <segment-uuid>",
	Short: "Count profile segment contacts",
	Long:  "Count the contacts currently matching a segment without listing them.\n\nCheaper than paging through the segment contacts endpoint when only the size is needed.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().ReachCountProfileSegmentContactsV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

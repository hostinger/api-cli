package segments

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:   "get <segment-uuid>",
	Short: "Get segment details",
	Long:  "Get details of a specific segment.\n\nThis endpoint retrieves information about a single segment identified by UUID.\nSegments are used to organize and group contacts based on specific criteria.\n\n**Deprecated.** This endpoint cannot target a profile, so it always falls back to\nthe client's default profile and cannot read segments of any other profile. Use\n`GET /api/reach/v1/profiles/{profileUuid}/segmentation/segments/{segmentUuid}` instead.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().ReachGetSegmentDetailsV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

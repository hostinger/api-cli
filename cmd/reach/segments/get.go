package segments

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:     "get <segment-uuid>",
	Short:   "Get segment details",
	Long:    "Get details of a specific segment.\n\nThis endpoint retrieves information about a single segment identified by UUID.\nSegments are used to organize and group contacts based on specific criteria.\n\nDeprecated. This endpoint cannot target a profile, so it always falls back to\nthe client's default profile and cannot read segments of any other profile. Use\n`GET /api/reach/v1/profiles/{profileUuid}/segmentation/segments/{segmentUuid}` instead.",
	Example: "  hostinger reach segments get 550e8400-e09b-41d4-a716-400055000000",
	ValidArgsFunction: utils.CompleteArgs(
		&utils.ArgSource{
			Value: "uuid",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().ReachListSegmentsV1WithResponse(ctx)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachGetSegmentDetailsV1WithResponse(ctx, args[0])
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		}
		respBody, status, err := call(cmd.Context())
		if err != nil {
			return err
		}
		return output.Format(cmd, respBody, status)
	},
}

func init() {
	utils.RegisterArgs(GetCmd, utils.StringArg("segment-uuid"))
}

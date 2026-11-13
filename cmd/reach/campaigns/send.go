package campaigns

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var SendCmd = &cobra.Command{
	Use:     "send <profile-uuid> <campaign-uuid>",
	Short:   "Send campaign",
	Long:    "Send a campaign now to the audience stored on it.\n\nThis is irreversible: once accepted, the emails go out and cannot be recalled. Confirm with the\ncustomer before calling. The campaign needs a subject, sender, template and audience first - set\nthem with the update endpoint. Delivery is asynchronous, so success means accepted, not delivered:\nthe campaign moves to `sending` and later to `publish`, or to `failed`. Read the campaign to follow\nit. The request is rejected when the campaign has no recipients, its sender domain is not verified\nor is blocked, the plan's sending limits would be exceeded, or the campaign is already being sent.",
	Example: "  hostinger reach campaigns send 550e8400-e09b-41d4-a716-400055000000 550e8400-e09b-41d4-a716-400055000000",
	ValidArgsFunction: utils.CompleteArgs(
		nil,
		&utils.ArgSource{
			Value: "uuid",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().ReachListCampaignsV1WithResponse(ctx, args[0], nil)
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
			r, err := api.Request().ReachSendCampaignV1WithResponse(ctx, args[0], args[1])
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
	utils.RegisterArgs(SendCmd, utils.StringArg("profile-uuid"), utils.StringArg("campaign-uuid"))
}

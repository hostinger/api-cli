package campaigns

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var PerformanceCmd = &cobra.Command{
	Use:     "performance <profile-uuid> <campaign-uuid>",
	Short:   "Get campaign performance",
	Long:    "Get the performance of a campaign: delivery, opens, clicks and unsubscribes, with the\nmatching rates.\n\nEvery count is unique contacts rather than raw events, so a contact who opens the same email\nfive times is counted once.",
	Example: "  hostinger reach campaigns performance 550e8400-e09b-41d4-a716-400055000000 550e8400-e09b-41d4-a716-400055000000",
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
			r, err := api.Request().ReachGetCampaignPerformanceV1WithResponse(ctx, args[0], args[1])
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
	utils.RegisterArgs(PerformanceCmd, utils.StringArg("profile-uuid"), utils.StringArg("campaign-uuid"))
}

package automations

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:     "get <profile-uuid> <automation-uuid>",
	Short:   "Get automation details",
	Long:    "Get a single automation with the counts of contacts that entered it, are moving through it,\nfinished it or failed on the way.\n\nThis describes the automation itself. To see the workflow it runs, use the steps endpoint.",
	Example: "  hostinger reach automations get 550e8400-e09b-41d4-a716-400055000000 550e8400-e09b-41d4-a716-400055000000",
	ValidArgsFunction: utils.CompleteArgs(
		nil,
		&utils.ArgSource{
			Value: "uuid",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().ReachListAutomationsV1WithResponse(ctx, args[0], nil)
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
			r, err := api.Request().ReachGetAutomationDetailsV1WithResponse(ctx, args[0], args[1])
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
	utils.RegisterArgs(GetCmd, utils.StringArg("profile-uuid"), utils.StringArg("automation-uuid"))
}

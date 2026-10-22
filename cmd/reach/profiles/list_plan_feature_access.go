package profiles

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListPlanFeatureAccessCmd = &cobra.Command{
	Use:               "list-plan-feature-access <profile-uuid>",
	Short:             "List plan feature access",
	Long:              "List which plan features the profile can use.\n\nThis is the feature lock matrix, not a usage quota. `available` means the feature can be\nused right now and `locked` means it is not part of the base plan, so an upgrade is needed.\nFor remaining emails, recipients and AI credits use the limits endpoint instead.\n\nWorth checking before building something that cannot be activated afterwards, such as an\nautomation on a plan without automation activation.",
	Example:           "  hostinger reach profiles list-plan-feature-access 550e8400-e09b-41d4-a716-400055000000",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachListPlanFeatureAccessV1WithResponse(ctx, args[0])
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
	utils.RegisterArgs(ListPlanFeatureAccessCmd, utils.StringArg("profile-uuid"))
}

package profiles

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var RemainingPlanLimitsCmd = &cobra.Command{
	Use:               "remaining-plan-limits <profile-uuid>",
	Short:             "Get remaining plan limits",
	Long:              "Get how much of the plan is left for the current period.\n\nTwo things to keep in mind before you build alerting on this. The period is a calendar month\nrather than a billing anniversary, so the counters reset on the 1st no matter when the\nsubscription started. And usage is tracked per order, so every profile on the same order shares\none pool and reports the same numbers here. Only the current period is available, past usage is\nnot kept.",
	Example:           "  hostinger reach profiles remaining-plan-limits 550e8400-e09b-41d4-a716-400055000000",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().ReachGetRemainingPlanLimitsV1WithResponse(ctx, args[0])
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
	utils.RegisterArgs(RemainingPlanLimitsCmd, utils.StringArg("profile-uuid"))
}

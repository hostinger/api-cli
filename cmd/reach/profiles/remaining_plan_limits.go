package profiles

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var RemainingPlanLimitsCmd = &cobra.Command{
	Use:   "remaining-plan-limits <profile-uuid>",
	Short: "Get remaining plan limits",
	Long:  "Get how much of the plan is left for the current period.\n\nTwo things to keep in mind before you build alerting on this. The period is a calendar month\nrather than a billing anniversary, so the counters reset on the 1st no matter when the\nsubscription started. And usage is tracked per order, so every profile on the same order shares\none pool and reports the same numbers here. Only the current period is available, past usage is\nnot kept.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().ReachGetRemainingPlanLimitsV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

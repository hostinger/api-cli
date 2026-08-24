package profiles

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListPlanFeatureAccessCmd = &cobra.Command{
	Use:   "list-plan-feature-access <profile-uuid>",
	Short: "List plan feature access",
	Long:  "List which plan features the profile can use.\n\nThis is the feature lock matrix, not a usage quota. `available` means the feature can be\nused right now and `locked` means it is not part of the base plan, so an upgrade is needed.\nFor remaining emails, recipients and AI credits use the limits endpoint instead.\n\nWorth checking before building something that cannot be activated afterwards, such as an\nautomation on a plan without automation activation.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().ReachListPlanFeatureAccessV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

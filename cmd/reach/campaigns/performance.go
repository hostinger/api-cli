package campaigns

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var PerformanceCmd = &cobra.Command{
	Use:   "performance <profile-uuid> <campaign-uuid>",
	Short: "Get campaign performance",
	Long:  "Get the performance of a campaign: delivery, opens, clicks and unsubscribes, with the\nmatching rates.\n\nEvery count is unique contacts rather than raw events, so a contact who opens the same email\nfive times is counted once.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().ReachGetCampaignPerformanceV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

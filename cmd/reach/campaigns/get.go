package campaigns

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
	Use:   "get <profile-uuid> <campaign-uuid>",
	Short: "Get campaign details",
	Long:  "Get a single campaign with its sender, subject, template reference, targeting and delivery\nprogress.\n\nThis describes how the campaign was set up and how far it has got. For opens, clicks and\nunsubscribes use the campaign statistics endpoint.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().ReachGetCampaignDetailsV1WithResponse(context.TODO(), args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

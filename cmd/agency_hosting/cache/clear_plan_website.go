package cache

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ClearPlanWebsiteCmd = &cobra.Command{
	Use:   "clear-plan-website <website_uid>",
	Short: "Clear Agency Plan website cache",
	Long:  "Clears cache for all domains associated with an Agency Plan website, including its preview domain.\n\nThis operation clears all cache types for the website.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().AgencyHostingClearAgencyPlanWebsiteCacheV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

package wordpress

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var PlanWebsiteSettingsCmd = &cobra.Command{
	Use:   "plan-website-settings <website_uid>",
	Short: "Get Agency Plan website WordPress settings",
	Long:  "Returns the current WordPress settings for an Agency Plan website: installed core version,\nLiteSpeed Cache plugin status, object cache status, and maintenance mode status.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().AgencyHostingGetAgencyPlanWebsiteWordPressSettingsV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

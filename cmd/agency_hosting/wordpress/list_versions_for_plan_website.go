package wordpress

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var ListVersionsForPlanWebsiteCmd = &cobra.Command{
	Use:   "list-versions-for-plan-website <website_uid>",
	Short: "List available WordPress versions for an Agency Plan website",
	Long:  "Lists the WordPress core versions available for installation on an Agency Plan website.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().AgencyHostingListAvailableWordPressVersionsForAnAgencyPlanWebsiteV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

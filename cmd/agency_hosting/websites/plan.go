package websites

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/spf13/cobra"
)

var PlanCmd = &cobra.Command{
	Use:   "plan <website_uid>",
	Short: "Get Agency Plan website details",
	Long:  "Retrieves detailed information about a specific Agency Plan website, including configuration,\nstatus, metadata, hosting plan details, and resource quotas.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().AgencyHostingGetAgencyPlanWebsiteDetailsV1WithResponse(context.TODO(), args[0])
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

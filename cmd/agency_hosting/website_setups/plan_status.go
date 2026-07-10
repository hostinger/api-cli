package website_setups

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var PlanStatusCmd = &cobra.Command{
	Use:   "plan-status <order_id> <setup_uuid>",
	Short: "Get Agency Plan website setup status",
	Long:  "Returns the current status of an Agency Plan website setup started via the setups\nendpoint.\n\nPoll this endpoint using the `setup_uuid` returned from the provisioning request until\n`status` becomes `completed`, at which point `website_uid` identifies the new website.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().AgencyHostingGetAgencyPlanWebsiteSetupStatusV1WithResponse(context.TODO(), utils.StringToInt(args[0]), uuid.MustParse(args[1]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

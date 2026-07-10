package datacenters

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListForPlanOrderCmd = &cobra.Command{
	Use:   "list-for-plan-order <order_id>",
	Short: "List available datacenters for an Agency Plan order",
	Long:  "Lists the datacenters available for provisioning a new website on the given Agency Plan\nhosting order.\n\nEach datacenter includes a `pinger_url` you can ping from the client to measure round-trip\nlatency; comparing the results across datacenters lets you pick the nearest one (lowest\nping) before choosing its `code` as the `datacenter_code` when creating a website setup.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().AgencyHostingListAvailableDatacentersForAnAgencyPlanOrderV1WithResponse(context.TODO(), utils.StringToInt(args[0]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

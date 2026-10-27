package datacenters

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:               "list [order-id]",
	Short:             "List available datacenters",
	Long:              "Lists the datacenters available for provisioning a new website on the given Agency Plan\nhosting order.\n\nEach datacenter includes a `pinger_url` you can ping from the client to measure round-trip\nlatency; comparing the results across datacenters lets you pick the nearest one (lowest\nping) before choosing its `code` as the `datacenter_code` when creating a website setup.\n\nThe order can be given as [order-id] or left out, in which case it comes from --order-id, HOSTINGER_ORDER_ID, the .hostinger/site.json written by `hostinger up` or `hostinger link` (hosting orders), or a prompt when run in a terminal (the only order is used without asking).",
	Example:           "  hostinger agency-hosting datacenters list 123456",
	ValidArgsFunction: cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		orderID, args, err := site.ResolveOrder(cmd, args, 0, "agency-hosting")
		if err != nil {
			return api.CommandError(err)
		}
		orderIDInt, err := utils.ScopeInt("order ID", orderID)
		if err != nil {
			return err
		}
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().AgencyHostingListAvailableDatacentersV1WithResponse(ctx, orderIDInt)
			if err != nil {
				return nil, 0, err
			}
			return r.Body, r.StatusCode(), nil
		}
		respBody, status, err := call(cmd.Context())
		if err != nil {
			return err
		}
		return output.Format(cmd, respBody, status)
	},
}

func init() {
	utils.RegisterScopedArgs(ListCmd, 1, utils.IntArg("order_id"))
}

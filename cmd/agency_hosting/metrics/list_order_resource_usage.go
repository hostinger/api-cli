package metrics

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/site"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListOrderResourceUsageCmd = &cobra.Command{
	Use:               "list-order-resource-usage [order-id]",
	Short:             "List order resource usage metrics",
	Long:              "Returns aggregated CPU, memory, and process usage for the Agency Plan order\nover the selected time frame, plus the plan quotas and a per-website\nbreakdown. Each website is identified by uid. Suspended and deleted websites\nare excluded from both the order totals and the per-website breakdown.\nValues may be up to one hour stale. Disk and inode usage are on the\ndisk-usage-metrics endpoint.\n\nThe order can be given as [order-id] or left out, in which case it comes from --order-id, HOSTINGER_ORDER_ID, the .hostinger/site.json written by `hostinger up` or `hostinger link` (hosting orders), or a prompt when run in a terminal (the only order is used without asking).",
	Example:           "  hostinger agency-hosting metrics list-order-resource-usage 123456",
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
		params := listOrderResourceUsageParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().AgencyHostingListOrderResourceUsageMetricsV1WithResponse(ctx, orderIDInt, params)
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
	utils.RegisterScopedArgs(ListOrderResourceUsageCmd, 1, utils.IntArg("order_id"))
	ListOrderResourceUsageCmd.Flags().IntP("time-frame-hours", "", 24, "Length of the window in hours, ending now. Bucket size grows with the window. (one of: 1, 24, 168, 336, 720)")
	ListOrderResourceUsageCmd.RegisterFlagCompletionFunc("time-frame-hours", cobra.FixedCompletions([]cobra.Completion{"1", "24", "168", "336", "720"}, cobra.ShellCompDirectiveNoFileComp))
}

func listOrderResourceUsageParams(cmd *cobra.Command) *client.AgencyHostingListOrderResourceUsageMetricsV1Params {
	params := &client.AgencyHostingListOrderResourceUsageMetricsV1Params{}
	if cmd.Flags().Changed("time-frame-hours") {
		v, _ := cmd.Flags().GetInt("time-frame-hours")
		e := client.AgencyHostingListOrderResourceUsageMetricsV1ParamsTimeFrameHours(v)
		params.TimeFrameHours = &e
	}
	return params
}

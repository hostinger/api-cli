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

var ListPlanOrderDiskUsageCmd = &cobra.Command{
	Use:               "list-plan-order-disk-usage [order-id]",
	Short:             "List Agency Plan order disk usage metrics",
	Long:              "Returns aggregated disk and inode usage for the Agency Plan order over the\nselected time frame, plus the plan quotas. Figures cover the whole order\naccount. Values may be up to one hour stale. CPU, memory, and process usage\nare on the resource-usage-metrics endpoint.\n\nThe order can be given as [order-id] or left out, in which case it comes from --order-id, HOSTINGER_ORDER_ID, the .hostinger/site.json written by `hostinger up` or `hostinger link` (hosting orders), or a prompt when run in a terminal (the only order is used without asking).",
	Example:           "  hostinger agency-hosting metrics list-plan-order-disk-usage 123456",
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
		params := listPlanOrderDiskUsageParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().AgencyHostingListAgencyPlanOrderDiskUsageMetricsV1WithResponse(ctx, orderIDInt, params)
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
	utils.RegisterScopedArgs(ListPlanOrderDiskUsageCmd, 1, utils.IntArg("order_id"))
	ListPlanOrderDiskUsageCmd.Flags().IntP("time-frame-days", "", 1, "Length of the window in days, ending now. Bucket size grows with the window. (one of: 1, 7, 14, 30)")
	ListPlanOrderDiskUsageCmd.RegisterFlagCompletionFunc("time-frame-days", cobra.FixedCompletions([]cobra.Completion{"1", "7", "14", "30"}, cobra.ShellCompDirectiveNoFileComp))
}

func listPlanOrderDiskUsageParams(cmd *cobra.Command) *client.AgencyHostingListAgencyPlanOrderDiskUsageMetricsV1Params {
	params := &client.AgencyHostingListAgencyPlanOrderDiskUsageMetricsV1Params{}
	if cmd.Flags().Changed("time-frame-days") {
		v, _ := cmd.Flags().GetInt("time-frame-days")
		e := client.AgencyHostingListAgencyPlanOrderDiskUsageMetricsV1ParamsTimeFrameDays(v)
		params.TimeFrameDays = &e
	}
	return params
}

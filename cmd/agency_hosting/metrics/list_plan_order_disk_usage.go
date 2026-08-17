package metrics

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ListPlanOrderDiskUsageCmd = &cobra.Command{
	Use:   "list-plan-order-disk-usage <order_id>",
	Short: "List Agency Plan order disk usage metrics",
	Long:  "Returns aggregated disk and inode usage for the Agency Plan order over the\nselected time frame, plus the plan quotas. Figures cover the whole order\naccount. Values may be up to one hour stale. CPU, memory, and process usage\nare on the resource-usage-metrics endpoint.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().AgencyHostingListAgencyPlanOrderDiskUsageMetricsV1WithResponse(context.TODO(), utils.StringToInt(args[0]), listPlanOrderDiskUsageParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListPlanOrderDiskUsageCmd.Flags().IntP("time-frame-days", "", 1, "Length of the window in days, ending now. Bucket size grows with the window. (one of: 1, 7, 14, 30)")
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

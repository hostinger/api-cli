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

var ListOrderResourceUsageCmd = &cobra.Command{
	Use:   "list-order-resource-usage <order_id>",
	Short: "List order resource usage metrics",
	Long:  "Returns aggregated CPU, memory, and process usage for the Agency Plan order\nover the selected time frame, plus the plan quotas and a per-website\nbreakdown. Each website is identified by uid. Suspended and deleted websites\nare excluded from both the order totals and the per-website breakdown.\nValues may be up to one hour stale. Disk and inode usage are on the\ndisk-usage-metrics endpoint.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().AgencyHostingListOrderResourceUsageMetricsV1WithResponse(context.TODO(), utils.StringToInt(args[0]), listOrderResourceUsageParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	ListOrderResourceUsageCmd.Flags().IntP("time-frame-hours", "", 24, "Length of the window in hours, ending now. Bucket size grows with the window. (one of: 1, 24, 168, 336, 720)")
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

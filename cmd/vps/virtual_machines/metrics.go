package virtual_machines

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var MetricsCmd = &cobra.Command{
	Use:   "metrics <virtual-machine-id>",
	Short: "Get metrics",
	Long:  "Retrieve historical metrics for a specified virtual machine.\n\nIt includes the following metrics: \n- CPU usage\n- Memory usage\n- Disk usage\n- Network usage\n- Uptime\n\nUse this endpoint to monitor VPS performance and resource utilization over time.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSGetMetricsV1WithResponse(context.TODO(), utils.StringToInt(args[0]), metricsParams(cmd))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

func init() {
	MetricsCmd.Flags().StringP("date-from", "", "", "")
	MetricsCmd.Flags().StringP("date-to", "", "", "")
	MetricsCmd.MarkFlagRequired("date-from")
	MetricsCmd.MarkFlagRequired("date-to")
}

func metricsParams(cmd *cobra.Command) *client.VPSGetMetricsV1Params {
	params := &client.VPSGetMetricsV1Params{}
	dateFromVal, _ := cmd.Flags().GetString("date-from")
	params.DateFrom = utils.StringToTime(dateFromVal)
	dateToVal, _ := cmd.Flags().GetString("date-to")
	params.DateTo = utils.StringToTime(dateToVal)
	return params
}

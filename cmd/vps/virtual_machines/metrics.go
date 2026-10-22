package virtual_machines

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/client"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var MetricsCmd = &cobra.Command{
	Use:     "metrics <virtual-machine-id>",
	Short:   "Get metrics",
	Long:    "Retrieve historical metrics for a specified virtual machine.\n\nIt includes the following metrics:\n- CPU usage\n- Memory usage\n- Disk usage\n- Network usage\n- Uptime\n\nUse this endpoint to monitor VPS performance and resource utilization over time.",
	Example: "  hostinger vps virtual-machines metrics 1268054 --date-from 2025-05-01T00:00:00Z --date-to 2025-06-01T00:00:00Z",
	ValidArgsFunction: utils.CompleteArgs(
		&utils.ArgSource{
			Value: "id",
			Fetch: func(ctx context.Context, args []string) ([]byte, int, error) {
				r, err := api.CompletionRequest().VPSGetVirtualMachinesV1WithResponse(ctx)
				if err != nil {
					return nil, 0, err
				}
				return r.Body, r.StatusCode(), nil
			},
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		args = utils.PromptedArgs(cmd, args)
		if err := utils.FormatCheck(cmd, "date-from", "date-time"); err != nil {
			return err
		}
		if err := utils.FormatCheck(cmd, "date-to", "date-time"); err != nil {
			return err
		}
		params := metricsParams(cmd)
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().VPSGetMetricsV1WithResponse(ctx, utils.StringToInt(args[0]), params)
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
	utils.RegisterArgs(MetricsCmd, utils.IntArg("virtual-machine-id"))
	MetricsCmd.Flags().StringP("date-from", "", "", "")
	MetricsCmd.Flags().StringP("date-to", "", "", "")
	MetricsCmd.MarkFlagRequired("date-from")
	MetricsCmd.MarkFlagRequired("date-to")
	MetricsCmd.RegisterFlagCompletionFunc("date-from", cobra.NoFileCompletions)
	MetricsCmd.RegisterFlagCompletionFunc("date-to", cobra.NoFileCompletions)
}

func metricsParams(cmd *cobra.Command) *client.VPSGetMetricsV1Params {
	params := &client.VPSGetMetricsV1Params{}
	dateFromVal, _ := cmd.Flags().GetString("date-from")
	params.DateFrom = utils.StringToTime(dateFromVal)
	dateToVal, _ := cmd.Flags().GetString("date-to")
	params.DateTo = utils.StringToTime(dateToVal)
	return params
}

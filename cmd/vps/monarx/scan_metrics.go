package monarx

import (
	"context"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ScanMetricsCmd = &cobra.Command{
	Use:     "scan-metrics <virtual-machine-id>",
	Short:   "Get scan metrics",
	Long:    "Retrieve scan metrics for the Monarx (https://www.monarx.com/) malware scanner\ninstalled on a specified virtual machine.\n\nThe scan metrics provide detailed information about malware scans performed\nby Monarx, including number of scans, detected threats, and other relevant\nstatistics. This information is useful for monitoring security status of the\nvirtual machine and assessing effectiveness of the malware scanner.\n\nUse this endpoint to monitor VPS security scan results and threat detection.",
	Example: "  hostinger vps monarx scan-metrics 1268054",
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
		call := func(ctx context.Context) ([]byte, int, error) {
			r, err := api.Request().VPSGetScanMetricsV1WithResponse(ctx, utils.StringToInt(args[0]))
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
	utils.RegisterArgs(ScanMetricsCmd, utils.IntArg("virtual-machine-id"))
}

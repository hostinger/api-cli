package monarx

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var ScanMetricsCmd = &cobra.Command{
	Use:   "scan-metrics <virtual-machine-id>",
	Short: "Get scan metrics",
	Long:  "Retrieve scan metrics for the [Monarx](https://www.monarx.com/) malware scanner\ninstalled on a specified virtual machine.\n\nThe scan metrics provide detailed information about malware scans performed\nby Monarx, including number of scans, detected threats, and other relevant\nstatistics. This information is useful for monitoring security status of the\nvirtual machine and assessing effectiveness of the malware scanner.\n\nUse this endpoint to monitor VPS security scan results and threat detection.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSGetScanMetricsV1WithResponse(context.TODO(), utils.StringToInt(args[0]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

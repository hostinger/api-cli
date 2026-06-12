package recovery

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var StopCmd = &cobra.Command{
	Use:   "stop <virtual-machine-id>",
	Short: "Stop recovery mode",
	Long:  "Stop recovery mode for a specified virtual machine.\n\nIf virtual machine is not in recovery mode, this operation will fail.\n\nUse this endpoint to exit system rescue mode and return VPS to normal operation.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSStopRecoveryModeV1WithResponse(context.TODO(), utils.StringToInt(args[0]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

package virtual_machines

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
	Short: "Stop virtual machine",
	Long:  "Stop a specified virtual machine.\n\nIf the virtual machine is already stopped, the request will still be processed without any effect.\n\nUse this endpoint to power off running VPS instances.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSStopVirtualMachineV1WithResponse(context.TODO(), utils.StringToInt(args[0]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

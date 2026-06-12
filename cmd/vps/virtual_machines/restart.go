package virtual_machines

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var RestartCmd = &cobra.Command{
	Use:   "restart <virtual-machine-id>",
	Short: "Restart virtual machine",
	Long:  "Restart a specified virtual machine by fully stopping and starting it.\n\nIf the virtual machine was stopped, it will be started.\n\nUse this endpoint to reboot VPS instances.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSRestartVirtualMachineV1WithResponse(context.TODO(), utils.StringToInt(args[0]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

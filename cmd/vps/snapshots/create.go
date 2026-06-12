package snapshots

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create <virtual-machine-id>",
	Short: "Create snapshot",
	Long:  "Create a snapshot of a specified virtual machine.\n\nA snapshot captures the state and data of the virtual machine at a specific point in time, \nallowing users to restore the virtual machine to that state if needed. \nThis operation is useful for backup purposes, system recovery, \nand testing changes without affecting the current state of the virtual machine.\n\n**Creating new snapshot will overwrite the existing snapshot!**\n\nUse this endpoint to capture VPS state for backup and recovery purposes.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSCreateSnapshotV1WithResponse(context.TODO(), utils.StringToInt(args[0]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

package snapshots

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var RestoreCmd = &cobra.Command{
	Use:   "restore <virtual-machine-id>",
	Short: "Restore snapshot",
	Long:  "Restore a specified virtual machine to a previous state using a snapshot.\n\nRestoring from a snapshot allows users to revert the virtual machine to that state,\nwhich is useful for system recovery, undoing changes, or testing.\n\nUse this endpoint to revert VPS instances to previous saved states.",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSRestoreSnapshotV1WithResponse(context.TODO(), utils.StringToInt(args[0]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

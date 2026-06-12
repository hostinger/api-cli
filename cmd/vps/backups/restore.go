package backups

import (
	"context"
	"log"

	"github.com/hostinger/api-cli/api"
	"github.com/hostinger/api-cli/output"
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var RestoreCmd = &cobra.Command{
	Use:   "restore <virtual-machine-id> <backup-id>",
	Short: "Restore backup",
	Long:  "Restore a backup for a specified virtual machine.\n\nThe system will then initiate the restore process, which may take some time depending on the size of the backup.\n\n**All data on the virtual machine will be overwritten with the data from the backup.**\n\nUse this endpoint to recover VPS data from backup points.",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {
		r, err := api.Request().VPSRestoreBackupV1WithResponse(context.TODO(), utils.StringToInt(args[0]), utils.StringToInt(args[1]))
		if err != nil {
			log.Fatal(err)
		}

		output.Format(cmd, r.Body, r.StatusCode())
	},
}

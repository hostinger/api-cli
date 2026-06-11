package backups

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "backups",
	Short: "VM backup management",
	Long:  `Safeguard your data by managing backups. You can list available backups or restore a virtual machine from a backup.`,
}

func init() {
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(RestoreCmd)
}

package backups

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "backups",
	Short: "Backups commands",
}

func init() {
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(RestoreCmd)
}

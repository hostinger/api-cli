package snapshots

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "snapshots",
	Short: "Snapshot commands",
}

func init() {
	GroupCmd.AddCommand(GetCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(RestoreCmd)
}

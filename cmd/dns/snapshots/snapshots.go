package snapshots

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "snapshots",
	Short: "DNS zone snapshot management",
	Long:  ``,
}

func init() {
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(GetCmd)
	GroupCmd.AddCommand(RestoreCmd)
}

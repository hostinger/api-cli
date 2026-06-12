package snapshots

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "snapshots",
	Short: "Snapshots commands",
}

func init() {
	GroupCmd.AddCommand(CreateCmd)
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(GetCmd)
	GroupCmd.AddCommand(RestoreCmd)
}

package snapshots

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "snapshots",
	Short: "Snapshot commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(GetCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(RestoreCmd)
}

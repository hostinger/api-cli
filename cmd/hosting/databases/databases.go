package databases

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "databases",
	Short: "Databases commands",
}

func init() {
	GroupCmd.AddCommand(ChangePasswordCmd)
	GroupCmd.AddCommand(CreateCmd)
	GroupCmd.AddCommand(CreateRemoteConnectionCmd)
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(DeleteRemoteConnectionCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(ListRemoteConnectionsCmd)
	GroupCmd.AddCommand(PhpmyadminLinkCmd)
	GroupCmd.AddCommand(RepairCmd)
}

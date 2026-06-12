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
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(PhpmyadminLinkCmd)
	GroupCmd.AddCommand(RepairCmd)
}

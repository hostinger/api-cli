package contacts

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "contacts",
	Short: "Contacts commands",
}

func init() {
	GroupCmd.AddCommand(CreateCmd)
	GroupCmd.AddCommand(CreateBulkCmd)
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(ListGroupsCmd)
}

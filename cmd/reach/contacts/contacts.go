package contacts

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "contacts",
	Short: "Contacts commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(CreateCmd)
	GroupCmd.AddCommand(CreateBulkCmd)
	GroupCmd.AddCommand(CreateInBulkCmd)
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(DeleteProfileCmd)
	GroupCmd.AddCommand(GetCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(ListGroupsCmd)
	GroupCmd.AddCommand(ListProfileCmd)
	GroupCmd.AddCommand(UpdateCmd)
}

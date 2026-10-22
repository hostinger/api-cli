package tags

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "tags",
	Short: "Tags commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(AssignContactToCmd)
	GroupCmd.AddCommand(AssignContactsToCmd)
	GroupCmd.AddCommand(CreateOrFindCmd)
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(ListProfileCmd)
	GroupCmd.AddCommand(RemoveContactFromCmd)
	GroupCmd.AddCommand(RemoveContactsFromCmd)
	GroupCmd.AddCommand(RenameCmd)
}

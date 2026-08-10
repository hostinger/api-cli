package tags

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "tags",
	Short: "Tags commands",
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

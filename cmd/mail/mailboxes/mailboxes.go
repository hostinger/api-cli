package mailboxes

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "mailboxes",
	Short: "Mailboxes commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(ChangeMailboxPasswordCmd)
	GroupCmd.AddCommand(CreateMailboxCmd)
	GroupCmd.AddCommand(DeleteMailboxCmd)
	GroupCmd.AddCommand(ListCmd)
}

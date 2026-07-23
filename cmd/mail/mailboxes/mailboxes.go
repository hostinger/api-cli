package mailboxes

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "mailboxes",
	Short: "Mailboxes commands",
}

func init() {
	GroupCmd.AddCommand(CreateMailboxCmd)
	GroupCmd.AddCommand(MailboxCmd)
}

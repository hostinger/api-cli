package mail

import (
	"github.com/hostinger/api-cli/cmd/mail/mailboxes"
	"github.com/hostinger/api-cli/cmd/mail/orders"

	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "mail",
	Short: "Mail commands",
}

func init() {
	GroupCmd.AddCommand(mailboxes.GroupCmd)
	GroupCmd.AddCommand(orders.GroupCmd)
}

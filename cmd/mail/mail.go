package mail

import (
	"github.com/hostinger/api-cli/cmd/mail/aliases"
	"github.com/hostinger/api-cli/cmd/mail/api_tokens"
	"github.com/hostinger/api-cli/cmd/mail/autoreplies"
	"github.com/hostinger/api-cli/cmd/mail/catchalls"
	"github.com/hostinger/api-cli/cmd/mail/forwarders"
	"github.com/hostinger/api-cli/cmd/mail/logs"
	"github.com/hostinger/api-cli/cmd/mail/mailboxes"
	"github.com/hostinger/api-cli/cmd/mail/orders"
	"github.com/hostinger/api-cli/cmd/mail/webhooks"

	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "mail",
	Short: "Mail commands",
}

func init() {
	GroupCmd.AddCommand(aliases.GroupCmd)
	GroupCmd.AddCommand(api_tokens.GroupCmd)
	GroupCmd.AddCommand(autoreplies.GroupCmd)
	GroupCmd.AddCommand(catchalls.GroupCmd)
	GroupCmd.AddCommand(forwarders.GroupCmd)
	GroupCmd.AddCommand(logs.GroupCmd)
	GroupCmd.AddCommand(mailboxes.GroupCmd)
	GroupCmd.AddCommand(orders.GroupCmd)
	GroupCmd.AddCommand(webhooks.GroupCmd)
}

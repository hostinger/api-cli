package logs

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "logs",
	Short: "Logs commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(ListAccessCmd)
	GroupCmd.AddCommand(ListActionCmd)
	GroupCmd.AddCommand(ListInboundCmd)
	GroupCmd.AddCommand(ListMailboxActionCmd)
	GroupCmd.AddCommand(ListOutboundCmd)
}

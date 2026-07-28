package logs

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "logs",
	Short: "Logs commands",
}

func init() {
	GroupCmd.AddCommand(ListAccessCmd)
	GroupCmd.AddCommand(ListActionCmd)
	GroupCmd.AddCommand(ListInboundCmd)
	GroupCmd.AddCommand(ListMailboxActionCmd)
	GroupCmd.AddCommand(ListOutboundCmd)
}

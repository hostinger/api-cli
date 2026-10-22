package move

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "move",
	Short: "Move commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(AcceptIncomingCmd)
	GroupCmd.AddCommand(CancelOutgoingCmd)
	GroupCmd.AddCommand(IncomingCmd)
	GroupCmd.AddCommand(IncomingListCmd)
	GroupCmd.AddCommand(OutgoingCmd)
	GroupCmd.AddCommand(OutgoingListCmd)
	GroupCmd.AddCommand(RejectIncomingCmd)
	GroupCmd.AddCommand(StartOutgoingCmd)
}

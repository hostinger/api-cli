package move

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "move",
	Short: "Move commands",
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

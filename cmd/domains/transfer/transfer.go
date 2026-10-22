package transfer

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Transfer commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(ClaimFreeCmd)
	GroupCmd.AddCommand(GetCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(StartCmd)
}

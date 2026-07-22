package transfer

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Transfer commands",
}

func init() {
	GroupCmd.AddCommand(GetCmd)
	GroupCmd.AddCommand(ListCmd)
}

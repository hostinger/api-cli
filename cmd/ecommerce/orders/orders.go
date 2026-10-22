package orders

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "orders",
	Short: "Orders commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(CancelCmd)
	GroupCmd.AddCommand(FulfilCmd)
	GroupCmd.AddCommand(ListStoreCmd)
	GroupCmd.AddCommand(RetrieveCmd)
}

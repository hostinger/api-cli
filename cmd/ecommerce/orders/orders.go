package orders

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "orders",
	Short: "Orders commands",
}

func init() {
	GroupCmd.AddCommand(CancelCmd)
	GroupCmd.AddCommand(FulfilCmd)
	GroupCmd.AddCommand(ListStoreCmd)
	GroupCmd.AddCommand(RetrieveCmd)
}

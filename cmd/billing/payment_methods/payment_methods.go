package payment_methods

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "payment-methods",
	Short: "Payment method management",
	Long:  ``,
}

func init() {
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(SetDefaultCmd)
}

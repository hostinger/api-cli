package payment_methods

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "payment-methods",
	Short: "Payment methods commands",
}

func init() {
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(SetDefaultCmd)
}

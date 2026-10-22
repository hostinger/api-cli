package payment_methods

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "payment-methods",
	Short: "Payment methods commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(DeleteCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(SetDefaultCmd)
}

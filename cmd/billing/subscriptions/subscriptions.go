package subscriptions

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "subscriptions",
	Short: "Subscription management",
	Long:  ``,
}

func init() {
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(EnableAutoRenewalCmd)
	GroupCmd.AddCommand(DisableAutoRenewalCmd)
}

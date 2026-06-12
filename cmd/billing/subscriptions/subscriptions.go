package subscriptions

import (
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "subscriptions",
	Short: "Subscriptions commands",
}

func init() {
	GroupCmd.AddCommand(DisableAutoRenewalCmd)
	GroupCmd.AddCommand(EnableAutoRenewalCmd)
	GroupCmd.AddCommand(ListCmd)
}

package subscriptions

import (
	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "subscriptions",
	Short: "Subscriptions commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(DisableAutoRenewalCmd)
	GroupCmd.AddCommand(EnableAutoRenewalCmd)
	GroupCmd.AddCommand(ListCmd)
	GroupCmd.AddCommand(RenewCmd)
}

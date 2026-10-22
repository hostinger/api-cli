package billing

import (
	"github.com/hostinger/api-cli/cmd/billing/catalog"
	"github.com/hostinger/api-cli/cmd/billing/orders"
	"github.com/hostinger/api-cli/cmd/billing/payment_methods"
	"github.com/hostinger/api-cli/cmd/billing/subscriptions"

	"github.com/hostinger/api-cli/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:   "billing",
	Short: "Billing commands",
	Args:  cobra.ArbitraryArgs,
	RunE:  utils.GroupRunE,
}

func init() {
	GroupCmd.AddCommand(catalog.GroupCmd)
	GroupCmd.AddCommand(orders.GroupCmd)
	GroupCmd.AddCommand(payment_methods.GroupCmd)
	GroupCmd.AddCommand(subscriptions.GroupCmd)
}
